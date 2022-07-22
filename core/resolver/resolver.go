/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This resolver is based on the built in grpc-go dns resolver.
// The biggest issue with the built in resolver is that it doesn't
// re-resolve until a fatal error or a trigger to `ResolveNow`. This
// isn't very sufficient and we want to periodically re-query DNS to
// refresh our view of the world so as new routes are added, we pick
// them up.
// This package is mostly based on the dns_resolver.go from grpc-go, just
// forked and modified for our needs. You shouldn't need to use this directly
// but our grpcclient uses it automatically.
package resolver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"sync"
	"time"

	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(NewBuilder())
}

const (
	defaultPort = "443"
	SchemeName  = "ps-dns"
)

var (
	errMissingAddr = errors.New("dns resolver: missing address")

	// Addresses ending with a colon that is supposed to be the separator
	// between host and port is not allowed.  E.g. "::" is a valid address as
	// it is an IPv6 address (host only) and "[::]:" is invalid as it ends with
	// a colon as the host and port separator
	errEndsWithColon = errors.New("dns resolver: missing port after port-separator colon")

	errNegativeRefresh = errors.New("dns resolver: negative refresh interval")
)

var catchallPrefix = netip.MustParsePrefix("0.0.0.0/0")

func NewBuilder() resolver.Builder {
	return &builder{}
}

type builder struct{}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, bopts resolver.BuildOptions) (resolver.Resolver, error) {
	opts := target.URL.Query()
	refreshInterval := forceDuration(getOr(opts, "refresh", "0s"))
	prefix := netip.MustParsePrefix(getOr(opts, "prefix", "0.0.0.0/0"))
	resolverAddr := getOr(opts, "resolver", "")

	if refreshInterval < 0 {
		return nil, errNegativeRefresh
	}

	// fall back to normal DNS resolver if no refreshing or filtering is wanted
	if refreshInterval == 0 && prefix == catchallPrefix && resolverAddr == "" {
		return resolver.Get("dns").Build(target, cc, bopts)
	}

	endpoint := target.URL.Path[1:]

	host, port, err := parseTarget(endpoint, defaultPort)
	if err != nil {
		return nil, err
	}

	// IP address.
	if ipAddr, ok := formatIP(host); ok {
		cc.UpdateState(resolver.State{
			Addresses: []resolver.Address{
				{Addr: ipAddr + ":" + port},
			},
		})
		return deadResolver{}, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	d := &dnsResolver{
		ctx:      ctx,
		cancel:   cancel,
		cc:       cc,
		rn:       make(chan struct{}, 1),
		host:     host,
		port:     port,
		refresh:  refreshInterval,
		prefix:   prefix,
		resolver: net.DefaultResolver,
	}

	if resolverAddr != "" {
		d.resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				var dialer net.Dialer
				return dialer.DialContext(ctx, network, resolverAddr)
			},
		}
	}

	d.wg.Add(1)
	go d.watcher()
	return d, nil
}

func (b *builder) Scheme() string {
	return SchemeName
}

type dnsResolver struct {
	ctx    context.Context
	cancel context.CancelFunc
	cc     resolver.ClientConn
	rn     chan struct{}
	wg     sync.WaitGroup

	host    string
	port    string
	refresh time.Duration
	prefix  netip.Prefix

	resolver *net.Resolver
}

func (r *dnsResolver) ResolveNow(resolver.ResolveNowOptions) {
	select {
	case r.rn <- struct{}{}:
	default:
	}
}

func (r *dnsResolver) Close() {
	r.cancel()
	r.wg.Wait()
}

func (r *dnsResolver) watcher() {
	defer r.wg.Done()
	var timer *time.Timer
	for {
		state, err := r.lookup()
		if err != nil {
			r.cc.ReportError(err)
		} else {
			r.cc.UpdateState(*state)
		}

		if timer == nil {
			timer = time.NewTimer(r.refresh)
		} else {
			timer.Reset(r.refresh)
		}

		select {
		case <-r.ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return
		case <-r.rn:
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
		}
	}
}

func handleDNSError(err error, lookupType string) error {
	if dnsErr, ok := err.(*net.DNSError); ok && !dnsErr.IsTimeout && !dnsErr.IsTemporary {
		// Timeouts and temporary errors should be communicated to gRPC to
		// attempt another DNS query (with backoff).  Other errors should be
		// suppressed (they may represent the absence of a TXT record).
		return nil
	}
	if err != nil {
		err = fmt.Errorf("dns: %v record lookup error: %v", lookupType, err)
		// logger.Info(err)
	}
	return err
}

func (r *dnsResolver) lookup() (*resolver.State, error) {
	addrs, err := r.lookupHost()
	if err != nil {
		return nil, err
	}

	return &resolver.State{
		Addresses: addrs,
	}, nil
}

func (r *dnsResolver) lookupHost() ([]resolver.Address, error) {
	addrs, err := r.resolver.LookupHost(r.ctx, r.host)
	if err != nil {
		err = handleDNSError(err, "A")
		return nil, err
	}
	// Allocate assuming all addresses match, but if they don't, we just overallocated
	// once rather than doing two passes to check out many pass first. Numbers here should
	// be so minimal it doesn't really matter. Dozens at most.
	newAddrs := make([]resolver.Address, 0, len(addrs))
	for _, a := range addrs {
		ip, err := netip.ParseAddr(a)
		if err != nil {
			return nil, fmt.Errorf("dns: error parsing A record IP address %v: %w", a, err)
		}
		if r.prefix == catchallPrefix || r.prefix.Contains(ip) {
			addr := formatAddr(ip) + ":" + r.port
			newAddrs = append(newAddrs, resolver.Address{Addr: addr})
		}
	}

	// Well, there are no IPs that match our prefix, but there are others that don't match,
	// so let's brute force and use those.
	if r.prefix != catchallPrefix && len(newAddrs) == 0 && len(addrs) > 0 {
		newAddrs = make([]resolver.Address, 0, len(addrs))
		for _, a := range addrs {
			// these have all already been checked in the loop prior, so this is safe.
			ip := netip.MustParseAddr(a)
			addr := formatAddr(ip) + ":" + r.port
			newAddrs = append(newAddrs, resolver.Address{Addr: addr})
		}
	}

	return newAddrs, nil
}

// formatIP returns ok = false if addr is not a valid textual representation of an IP address.
// If addr is an IPv4 address, return the addr and ok = true.
// If addr is an IPv6 address, return the addr enclosed in square brackets and ok = true.
func formatIP(addr string) (addrIP string, ok bool) {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return "", false
	}
	return formatAddr(ip), true
}

func formatAddr(addr netip.Addr) string {
	if addr.Is4() {
		return addr.String()
	}
	return "[" + addr.String() + "]"
}

// parseTarget takes the user input target string and default port, returns formatted host and port info.
// If target doesn't specify a port, set the port to be the defaultPort.
// If target is in IPv6 format and host-name is enclosed in square brackets, brackets
// are stripped when setting the host.
// examples:
// target: "www.google.com" defaultPort: "443" returns host: "www.google.com", port: "443"
// target: "ipv4-host:80" defaultPort: "443" returns host: "ipv4-host", port: "80"
// target: "[ipv6-host]" defaultPort: "443" returns host: "ipv6-host", port: "443"
// target: ":80" defaultPort: "443" returns host: "localhost", port: "80"
func parseTarget(target, defaultPort string) (host, port string, err error) {
	if target == "" {
		return "", "", errMissingAddr
	}
	if _, err := netip.ParseAddr(target); err == nil {
		// target is an IPv4 or IPv6(without brackets) address
		return target, defaultPort, nil
	}
	if host, port, err = net.SplitHostPort(target); err == nil {
		if port == "" {
			// If the port field is empty (target ends with colon), e.g. "[::1]:", this is an error.
			return "", "", errEndsWithColon
		}
		// target has port, i.e ipv4-host:port, [ipv6-host]:port, host-name:port
		if host == "" {
			// Keep consistent with net.Dial(): If the host is empty, as in ":80", the local system is assumed.
			host = "localhost"
		}
		return host, port, nil
	}
	if host, port, err = net.SplitHostPort(target + ":" + defaultPort); err == nil {
		// target doesn't have port
		return host, port, nil
	}
	return "", "", fmt.Errorf("invalid target address %v, error info: %v", target, err)
}

func getOr(v url.Values, key string, fallback string) string {
	if !v.Has(key) {
		return fallback
	}
	if val := v.Get(key); val != "" {
		return val
	}
	return fallback
}

func forceDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

// deadResolver is a resolver that does nothing.
type deadResolver struct{}

func (deadResolver) ResolveNow(resolver.ResolveNowOptions) {}

func (deadResolver) Close() {}

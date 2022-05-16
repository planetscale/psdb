package connset

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"
)

// ErrConnSetClosed is returned for any operation when the connset was
// closed with the Close() method.
var ErrConnSetClosed = errors.New("connset is closed")

// prevent importing "math" just for this constant, math.MaxUint16
const maxUint16 = 65535

type ConnSet struct {
	dialer DialerFn

	warming   map[string]struct{}
	warmingMu sync.Mutex

	conns   map[string]*trackedConn
	connsMu sync.Mutex

	closed  bool
	closeCh chan struct{}
}

type trackedConn struct {
	conn io.Closer
	c    uint16 // 65k connections are enough for anyone :)
	lu   time.Time
}

func (tc *trackedConn) incr() {
	if tc.c == maxUint16 {
		panic("uint16 overflow, 65k connections were, in fact, not enough")
	}
	tc.c++
	tc.lu = time.Now()
}

func (tc *trackedConn) decr() {
	if tc.c > 0 {
		// prevent underflow, but don't panic() since
		// this is a bug in code, but undercounting wouldn't cause
		// harm in a ConnSet. We just want to strictly prevent
		// the uint16 from wrapping around
		tc.c--
	}
	tc.lu = time.Now()
}

type DialerFn func(ctx context.Context, addr string) (io.Closer, error)

func New(dialer DialerFn, ttl time.Duration) *ConnSet {
	c := &ConnSet{
		dialer:  dialer,
		warming: make(map[string]struct{}),
		conns:   make(map[string]*trackedConn),
		closeCh: make(chan struct{}, 1),
	}
	if int64(ttl) >= 0 {
		go c.cleanup(ttl)
	}
	return c
}

// Dial is intended to pre-warm a connection without retrieving it or using it.
// It is intentionally not preventing concurrent dials from `Get()` since in all cases,
// Get shouldn't be stuck behind something like this. We also don't try to prevent
// concurrent and extraneous dialing, but we do intentionally want to prevent
// concurrent pre-warms via Dial for the same target address.
// Unlike Get(), Dial() is intentionally only an optimization, and succeeding or failing
// here isn't as important. We don't want to bombard a service with many concurrent dials
// while prewarming. Only one is necessary.
func (c *ConnSet) Dial(ctx context.Context, addr string) error {
	if c.closed {
		return ErrConnSetClosed
	}

	c.connsMu.Lock()
	if _, ok := c.conns[addr]; ok {
		// someone is already dialing or has dialed, so just bail,
		// we don't need to dial twice.
		c.connsMu.Unlock()
		return nil
	}
	c.connsMu.Unlock()

	c.warmingMu.Lock()
	if _, ok := c.warming[addr]; ok {
		// this address is already being warmed, so bail.
		c.warmingMu.Unlock()
		return nil
	}

	// add to the warming queue
	c.warming[addr] = struct{}{}
	c.warmingMu.Unlock()

	defer func() {
		c.warmingMu.Lock()
		delete(c.warming, addr)
		c.warmingMu.Unlock()
	}()

	// get without tracking to avoid needing to release
	if _, err := c.get(ctx, addr, false /* track */); err != nil {
		return err
	}
	return nil
}

// Get returns the connection for the given addr. Get is allowed to dial multiple times
// concurrently, with the hope that your dial may finish sooner and you get to use it.
// The first Get() that succeeds is the one that will be cached for use. This does allow
// multiple concurrent dials being used against the underlying service, with the first one winning
// and the others being closed immediately. Once a connection is established, it does not allow
// multiple concurrent dialers anymore since they should all use the same underlying cached
// connection.
func (c *ConnSet) Get(ctx context.Context, addr string) (io.Closer, error) {
	return c.get(ctx, addr, true /* track */)
}

func (c *ConnSet) get(ctx context.Context, addr string, track bool) (io.Closer, error) {
	if c.closed {
		return nil, ErrConnSetClosed
	}

	c.connsMu.Lock()
	tconn, ok := c.conns[addr]
	if ok {
		// we already dialed, return the existing conn
		if track {
			tconn.incr()
		}
		c.connsMu.Unlock()
		return tconn.conn, nil
	}
	c.connsMu.Unlock()

	conn, err := c.dialer(ctx, addr)
	if err != nil {
		return nil, err
	}

	// it's possible that two connections dialed at the same time, make sure to
	// only store one connection
	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	tconn, ok = c.conns[addr]
	if ok {
		// we already dialed (the first dial won the race), return the existing
		// conn and close the second dialed connection
		go conn.Close()
		if track {
			tconn.incr()
		}
		return tconn.conn, nil
	}

	// the first dial won the race, store the connection
	tconn = &trackedConn{conn, 0, time.Now()}
	if track {
		tconn.incr()
	}

	c.conns[addr] = tconn
	return tconn.conn, nil
}

// Releases the connection, effectively decrementing it's lease count
func (c *ConnSet) Release(addr string) error {
	if c.closed {
		return ErrConnSetClosed
	}

	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	tconn, ok := c.conns[addr]
	if !ok {
		return nil
	}
	tconn.decr()

	return nil
}

// Close closes all underlying connections
func (c *ConnSet) Close() {
	c.closed = true

	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	for _, tconn := range c.conns {
		go tconn.conn.Close()
	}

	c.conns = make(map[string]*trackedConn)
	c.closeCh <- struct{}{}
}

func (c *ConnSet) cleanup(ttl time.Duration) {
	freq := time.Duration(ttl.Milliseconds()/3) * time.Millisecond

	for {
		select {
		case <-c.closeCh:
			return
		case <-time.After(freq):
			if c.closed {
				return
			}

			now := time.Now()
			threshold := now.Add(-ttl)

			c.connsMu.Lock()
			for addr, tconn := range c.conns {
				if tconn.c == 0 && tconn.lu.Before(threshold) {
					delete(c.conns, addr)
					go tconn.conn.Close()
				}
			}
			c.connsMu.Unlock()
		}
	}
}

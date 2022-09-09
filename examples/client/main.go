package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bufbuild/connect-go"

	"github.com/planetscale/psdb/auth"
	psdbclient "github.com/planetscale/psdb/core/client"
	psdbv1beta1 "github.com/planetscale/psdb/types/psdb/v1beta1"
	"github.com/planetscale/psdb/types/psdb/v1beta1/psdbv1beta1connect"
)

var (
	flagAddr     = flag.String("addr", "127.0.0.1:8080", "rpc address to test")
	flagTLSCa    = flag.String("tls-ca", "testcerts/ca-cert.pem", "")
	flagUser     = flag.String("u", "xxxxxxxxxx", "")
	flagPassword = flag.String("p", "bar", "")
)

func main() {
	flag.Parse()

	tlsConfig := psdbclient.DefaultTLSConfig()
	if *flagTLSCa != "" {
		var err error
		tlsConfig, err = psdbclient.TLSConfigWithRoot(*flagTLSCa)
		if err != nil {
			panic(err)
		}
	}

	opts := []psdbclient.Option{
		psdbclient.WithTLSConfig(tlsConfig),
	}

	client := psdbclient.New(
		*flagAddr,
		psdbv1beta1connect.NewDatabaseClient,
		auth.NewBasicAuth(*flagUser, *flagPassword),
		opts...,
	)
	fmt.Println(client)

	fmt.Println(client.CreateSession(context.Background(), connect.NewRequest(&psdbv1beta1.CreateSessionRequest{})))

	pool := psdbclient.NewUnauthenticatedPool(
		psdbv1beta1connect.NewDatabaseClient,
		opts...,
	)

	client = pool.Get(*flagAddr)
	fmt.Println(client, pool.Get(*flagAddr))
	fmt.Println(client.CreateSession(context.Background(), connect.NewRequest(&psdbv1beta1.CreateSessionRequest{})))
	fmt.Println(pool.Len())
}

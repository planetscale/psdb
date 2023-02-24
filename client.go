package psdb

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/planetscale/psdb/auth"
	coreclient "github.com/planetscale/psdb/core/client"
	psdbv1 "github.com/planetscale/psdb/types/psdb/v1"
	"github.com/planetscale/psdb/types/psdb/v1/psdbv1connect"
)

type (
	TableCursor = psdbv1.TableCursor
	SyncStream  = connect.ServerStreamForClient[psdbv1.SyncResponse]
	SyncRequest = psdbv1.SyncRequest
)

const (
	Primary = psdbv1.TabletType_primary
	Replica = psdbv1.TabletType_replica
)

// Client is a PlanetScale Database client
type Client struct {
	core psdbv1connect.DatabaseClient
}

// New creates a new Client with the provided Config
func New(cfg Config) *Client {
	return &Client{
		core: coreclient.New(
			cfg.Host,
			psdbv1connect.NewDatabaseClient,
			auth.NewBasicAuth(cfg.User, cfg.Password),
		),
	}
}

func (c *Client) Sync(ctx context.Context, r *SyncRequest) (*SyncStream, error) {
	return c.core.Sync(ctx, connect.NewRequest(r))
}

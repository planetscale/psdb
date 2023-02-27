// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: psdb/v1/database.proto

package psdbv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/planetscale/psdb/types/psdb/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// DatabaseName is the fully-qualified name of the Database service.
	DatabaseName = "psdb.v1.Database"
)

// DatabaseClient is a client for the psdb.v1.Database service.
type DatabaseClient interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
	Execute(context.Context, *connect_go.Request[v1.ExecuteRequest]) (*connect_go.Response[v1.ExecuteResponse], error)
	StreamExecute(context.Context, *connect_go.Request[v1.ExecuteRequest]) (*connect_go.ServerStreamForClient[v1.ExecuteResponse], error)
	Prepare(context.Context, *connect_go.Request[v1.PrepareRequest]) (*connect_go.Response[v1.PrepareResponse], error)
	CloseSession(context.Context, *connect_go.Request[v1.CloseSessionRequest]) (*connect_go.Response[v1.CloseSessionResponse], error)
	Sync(context.Context, *connect_go.Request[v1.SyncRequest]) (*connect_go.ServerStreamForClient[v1.SyncResponse], error)
}

// NewDatabaseClient constructs a client for the psdb.v1.Database service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDatabaseClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DatabaseClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &databaseClient{
		createSession: connect_go.NewClient[v1.CreateSessionRequest, v1.CreateSessionResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/CreateSession",
			opts...,
		),
		execute: connect_go.NewClient[v1.ExecuteRequest, v1.ExecuteResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/Execute",
			opts...,
		),
		streamExecute: connect_go.NewClient[v1.ExecuteRequest, v1.ExecuteResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/StreamExecute",
			opts...,
		),
		prepare: connect_go.NewClient[v1.PrepareRequest, v1.PrepareResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/Prepare",
			opts...,
		),
		closeSession: connect_go.NewClient[v1.CloseSessionRequest, v1.CloseSessionResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/CloseSession",
			opts...,
		),
		sync: connect_go.NewClient[v1.SyncRequest, v1.SyncResponse](
			httpClient,
			baseURL+"/psdb.v1.Database/Sync",
			opts...,
		),
	}
}

// databaseClient implements DatabaseClient.
type databaseClient struct {
	createSession *connect_go.Client[v1.CreateSessionRequest, v1.CreateSessionResponse]
	execute       *connect_go.Client[v1.ExecuteRequest, v1.ExecuteResponse]
	streamExecute *connect_go.Client[v1.ExecuteRequest, v1.ExecuteResponse]
	prepare       *connect_go.Client[v1.PrepareRequest, v1.PrepareResponse]
	closeSession  *connect_go.Client[v1.CloseSessionRequest, v1.CloseSessionResponse]
	sync          *connect_go.Client[v1.SyncRequest, v1.SyncResponse]
}

// CreateSession calls psdb.v1.Database.CreateSession.
func (c *databaseClient) CreateSession(ctx context.Context, req *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return c.createSession.CallUnary(ctx, req)
}

// Execute calls psdb.v1.Database.Execute.
func (c *databaseClient) Execute(ctx context.Context, req *connect_go.Request[v1.ExecuteRequest]) (*connect_go.Response[v1.ExecuteResponse], error) {
	return c.execute.CallUnary(ctx, req)
}

// StreamExecute calls psdb.v1.Database.StreamExecute.
func (c *databaseClient) StreamExecute(ctx context.Context, req *connect_go.Request[v1.ExecuteRequest]) (*connect_go.ServerStreamForClient[v1.ExecuteResponse], error) {
	return c.streamExecute.CallServerStream(ctx, req)
}

// Prepare calls psdb.v1.Database.Prepare.
func (c *databaseClient) Prepare(ctx context.Context, req *connect_go.Request[v1.PrepareRequest]) (*connect_go.Response[v1.PrepareResponse], error) {
	return c.prepare.CallUnary(ctx, req)
}

// CloseSession calls psdb.v1.Database.CloseSession.
func (c *databaseClient) CloseSession(ctx context.Context, req *connect_go.Request[v1.CloseSessionRequest]) (*connect_go.Response[v1.CloseSessionResponse], error) {
	return c.closeSession.CallUnary(ctx, req)
}

// Sync calls psdb.v1.Database.Sync.
func (c *databaseClient) Sync(ctx context.Context, req *connect_go.Request[v1.SyncRequest]) (*connect_go.ServerStreamForClient[v1.SyncResponse], error) {
	return c.sync.CallServerStream(ctx, req)
}

// DatabaseHandler is an implementation of the psdb.v1.Database service.
type DatabaseHandler interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
	Execute(context.Context, *connect_go.Request[v1.ExecuteRequest]) (*connect_go.Response[v1.ExecuteResponse], error)
	StreamExecute(context.Context, *connect_go.Request[v1.ExecuteRequest], *connect_go.ServerStream[v1.ExecuteResponse]) error
	Prepare(context.Context, *connect_go.Request[v1.PrepareRequest]) (*connect_go.Response[v1.PrepareResponse], error)
	CloseSession(context.Context, *connect_go.Request[v1.CloseSessionRequest]) (*connect_go.Response[v1.CloseSessionResponse], error)
	Sync(context.Context, *connect_go.Request[v1.SyncRequest], *connect_go.ServerStream[v1.SyncResponse]) error
}

// NewDatabaseHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDatabaseHandler(svc DatabaseHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/psdb.v1.Database/CreateSession", connect_go.NewUnaryHandler(
		"/psdb.v1.Database/CreateSession",
		svc.CreateSession,
		opts...,
	))
	mux.Handle("/psdb.v1.Database/Execute", connect_go.NewUnaryHandler(
		"/psdb.v1.Database/Execute",
		svc.Execute,
		opts...,
	))
	mux.Handle("/psdb.v1.Database/StreamExecute", connect_go.NewServerStreamHandler(
		"/psdb.v1.Database/StreamExecute",
		svc.StreamExecute,
		opts...,
	))
	mux.Handle("/psdb.v1.Database/Prepare", connect_go.NewUnaryHandler(
		"/psdb.v1.Database/Prepare",
		svc.Prepare,
		opts...,
	))
	mux.Handle("/psdb.v1.Database/CloseSession", connect_go.NewUnaryHandler(
		"/psdb.v1.Database/CloseSession",
		svc.CloseSession,
		opts...,
	))
	mux.Handle("/psdb.v1.Database/Sync", connect_go.NewServerStreamHandler(
		"/psdb.v1.Database/Sync",
		svc.Sync,
		opts...,
	))
	return "/psdb.v1.Database/", mux
}

// UnimplementedDatabaseHandler returns CodeUnimplemented from all methods.
type UnimplementedDatabaseHandler struct{}

func (UnimplementedDatabaseHandler) CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.CreateSession is not implemented"))
}

func (UnimplementedDatabaseHandler) Execute(context.Context, *connect_go.Request[v1.ExecuteRequest]) (*connect_go.Response[v1.ExecuteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.Execute is not implemented"))
}

func (UnimplementedDatabaseHandler) StreamExecute(context.Context, *connect_go.Request[v1.ExecuteRequest], *connect_go.ServerStream[v1.ExecuteResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.StreamExecute is not implemented"))
}

func (UnimplementedDatabaseHandler) Prepare(context.Context, *connect_go.Request[v1.PrepareRequest]) (*connect_go.Response[v1.PrepareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.Prepare is not implemented"))
}

func (UnimplementedDatabaseHandler) CloseSession(context.Context, *connect_go.Request[v1.CloseSessionRequest]) (*connect_go.Response[v1.CloseSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.CloseSession is not implemented"))
}

func (UnimplementedDatabaseHandler) Sync(context.Context, *connect_go.Request[v1.SyncRequest], *connect_go.ServerStream[v1.SyncResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("psdb.v1.Database.Sync is not implemented"))
}

package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientApp struct {
	handler pb.SystemStatsServiceClient
	conn    *grpc.ClientConn
}

func NewClientApp(addr string, connTimeout time.Duration) (*ClientApp, error) {
	grpcClient, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: connTimeout,
		}),
	)
	if err != nil {
		return nil, err
	}

	handler := pb.NewSystemStatsServiceClient(grpcClient)
	return &ClientApp{
		handler: handler,
		conn:    grpcClient,
	}, nil
}

func (c *ClientApp) Run(ctx context.Context) error {
	stream, err := c.handler.GetSystemStats(ctx, &pb.EmptyRequest{})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if ctx.Err() != nil && errors.Is(ctx.Err(), context.Canceled) {
			return nil
		}
		if err != nil {
			return err
		}
		// print received stats
		fmt.Printf("Received:\n%v\n", resp)
	}
}

func (c *ClientApp) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

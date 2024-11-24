package app

import (
	"context"
	"errors"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GetSystemStatsResponseHandler func(*pb.SystemStatsResponse) error

type ClientApp struct {
	serviceClient pb.SystemStatsServiceClient
	handler       GetSystemStatsResponseHandler
	conn          *grpc.ClientConn
	log           logging.Logger
}

func NewClientApp(
	addr string,
	connTimeout time.Duration,
	handler GetSystemStatsResponseHandler,
	logger logging.Logger,
) (*ClientApp, error) {
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

	serviceClient := pb.NewSystemStatsServiceClient(grpcClient)
	return &ClientApp{
		serviceClient: serviceClient,
		handler:       handler,
		conn:          grpcClient,
		log:           logger,
	}, nil
}

func (c *ClientApp) Run(ctx context.Context) error {
	c.log.Info("getting system stats")
	stream, err := c.serviceClient.GetSystemStats(ctx, &pb.EmptyRequest{})
	if err != nil {
		return err
	}
	c.log.Info("stream opened")
	for {
		resp, err := stream.Recv()
		if ctx.Err() != nil && errors.Is(ctx.Err(), context.Canceled) {
			return nil
		}
		if err != nil {
			return err
		}
		// print received stats
		if err := c.handler(resp); err != nil {
			return err
		}
	}
}

func (c *ClientApp) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

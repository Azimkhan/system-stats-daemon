package app

import (
	"context"
	"net"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/core/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	lsn        net.Listener
}

func (s *Server) Serve() error {
	return s.grpcServer.Serve(s.lsn)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}

func NewServer(ctx context.Context, conf *config.Config) (*Server, error) {
	// gRPC server
	lsn, err := net.Listen("tcp", conf.Server.BindAddr)
	if err != nil {
		return nil, err
	}

	collectInterval := time.Duration(float64(conf.Stream.Interval.Nanoseconds()) / 2.5)
	// create stat service
	statService, err := service.NewStatService(
		[]string{"cpuloadavg", "diskio"},
		collectInterval,
	)
	if err != nil {
		return nil, err
	}
	// run stat service
	go func() {
		statService.Run(ctx)
	}()

	// create and register rpc handler
	handler := NewRPCHandler(ctx, statService, conf.Stream.InitialDelay, conf.Stream.Interval)
	grpcServer := grpc.NewServer()
	pb.RegisterSystemStatsServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	server := &Server{
		grpcServer: grpcServer,
		lsn:        lsn,
	}
	return server, nil
}

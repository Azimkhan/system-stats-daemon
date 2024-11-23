package app

import (
	"context"
	"net"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/core/service"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	lsn        net.Listener
	log        logging.Logger
}

func (s *Server) Serve() error {
	s.log.Info("starting gRPC server", "addr", s.lsn.Addr().String())
	return s.grpcServer.Serve(s.lsn)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}

func NewServer(
	ctx context.Context,
	serverConfig *config.ServerConfig,
	streamConfig *config.StreamingConfig,
	logger logging.Logger,
) (*Server, error) {
	// gRPC server
	lsn, err := net.Listen("tcp", serverConfig.BindAddr)
	if err != nil {
		return nil, err
	}

	collectInterval := time.Duration(float64(streamConfig.Interval.Nanoseconds()) / 2.5)
	// create stat service
	statService, err := service.NewStatService(
		[]string{"cpuloadavg", "diskio"},
		collectInterval,
		logger,
	)
	if err != nil {
		return nil, err
	}
	// run stat service
	go func() {
		statService.Run(ctx)
	}()

	// create and register rpc handler
	handler := NewRPCHandler(ctx, statService, streamConfig.InitialDelay, streamConfig.Interval, logger)
	grpcServer := grpc.NewServer()
	pb.RegisterSystemStatsServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	server := &Server{
		grpcServer: grpcServer,
		lsn:        lsn,
		log:        logger,
	}
	return server, nil
}

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

type ServerApp struct {
	ctx         context.Context
	grpcServer  *grpc.Server
	lsn         net.Listener
	log         logging.Logger
	statService *service.StatService
}

func (s *ServerApp) Serve() error {
	// run stat service
	go func() {
		s.statService.Run(s.ctx)
	}()

	// run gRPC server
	s.log.Info("starting gRPC server", "addr", s.lsn.Addr().String())
	return s.grpcServer.Serve(s.lsn)
}

func (s *ServerApp) Stop() {
	s.grpcServer.Stop()
}

func NewServerApp(
	ctx context.Context,
	stats []string,
	serverConfig *config.ServerConfig,
	streamConfig *config.StreamingConfig,
	logger logging.Logger,
) (*ServerApp, error) {
	collectInterval := time.Duration(float64(streamConfig.Interval.Nanoseconds()) / 2.5)
	// create stat service
	statService, err := service.NewStatService(
		stats,
		collectInterval,
		logger,
	)
	if err != nil {
		return nil, err
	}

	// init gRPC server
	lsn, err := net.Listen("tcp", serverConfig.BindAddr)
	if err != nil {
		return nil, err
	}
	handler := NewRPCHandler(ctx, statService, streamConfig.InitialDelay, streamConfig.Interval, logger)
	grpcServer := grpc.NewServer()
	pb.RegisterSystemStatsServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	server := &ServerApp{
		ctx:         ctx,
		grpcServer:  grpcServer,
		statService: statService,
		lsn:         lsn,
		log:         logger,
	}
	return server, nil
}

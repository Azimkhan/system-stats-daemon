package app

import (
	"context"
	"errors"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"github.com/Azimkhan/system-stats-daemon/internal/core/service"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
)

var ErrNoStats = errors.New("no stats available")

type RPCHandler struct {
	pb.UnsafeSystemStatsServiceServer
	ctx      context.Context
	service  *service.StatService
	delay    time.Duration
	interval time.Duration
	log      logging.Logger
}

func NewRPCHandler(
	ctx context.Context,
	service *service.StatService,
	delay time.Duration,
	interval time.Duration,
	logger logging.Logger,
) *RPCHandler {
	return &RPCHandler{
		ctx:      ctx,
		service:  service,
		delay:    delay,
		interval: interval,
		log:      logger.With("service", "gRPC"),
	}
}

func (s *RPCHandler) GetSystemStats(
	_ *pb.EmptyRequest,
	server pb.SystemStatsService_GetSystemStatsServer,
) error {
	s.log.Debug("client connected", "method", "GetSystemStats")
	// Initial delay
	time.Sleep(s.delay)

	// Send initial stats
	if err := s.sendStats(server); err != nil {
		return err
	}

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	// Send stats periodically
	for {
		select {
		case <-server.Context().Done():
			s.log.Debug("client disconnected")
			return nil
		case <-s.ctx.Done():
			s.log.Debug("server shutting down")
			return nil
		case <-ticker.C:
			if err := s.sendStats(server); err != nil {
				return err
			}
		}
	}
}

func (s *RPCHandler) sendStats(server pb.SystemStatsService_GetSystemStatsServer) error {
	s.log.Debug("sending stats")
	stats, err := s.service.GetStats()
	if err != nil {
		return err
	}
	if stats == nil {
		return ErrNoStats
	}
	response := s.mapResponse(stats)
	return server.Send(response)
}

func (s *RPCHandler) mapResponse(stats *core.Stats) (response *pb.SystemStatsResponse) {
	response = &pb.SystemStatsResponse{}
	if stats.CPULoadAverage != nil {
		rows := make([]*pb.CPULoadAverage, 0, len(stats.CPULoadAverage.Rows))
		for _, row := range stats.CPULoadAverage.Rows {
			rows = append(rows, &pb.CPULoadAverage{
				MinutesAgo:  row.MinutesAgo,
				AverageLoad: row.Value,
			})
		}
		response.CpuLoadAverage = rows
	}

	if stats.DiskLoad != nil {
		rows := make([]*pb.DiskLoad, 0, len(stats.DiskLoad.Rows))
		for _, row := range stats.DiskLoad.Rows {
			rows = append(rows, &pb.DiskLoad{
				Device:                row.Device,
				Throughput:            row.Throughput,
				TransactionsPerSecond: row.TPS,
			})
		}
		response.DiskLoad = rows
	}
	return
}

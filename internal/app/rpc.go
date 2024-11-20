package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"github.com/Azimkhan/system-stats-daemon/internal/core/service"
)

var ErrNoStats = errors.New("no stats available")

type RPCHandler struct {
	pb.UnsafeSystemStatsServiceServer
	ctx      context.Context
	service  *service.StatService
	delay    time.Duration
	interval time.Duration
}

func NewRPCHandler(
	ctx context.Context,
	service *service.StatService,
	delay time.Duration,
	interval time.Duration,
) *RPCHandler {
	return &RPCHandler{
		ctx:      ctx,
		service:  service,
		delay:    delay,
		interval: interval,
	}
}

func (s *RPCHandler) GetSystemStats(
	_ *pb.EmptyRequest,
	server pb.SystemStatsService_GetSystemStatsServer,
) error {
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
			fmt.Println("client done")
			return nil
		case <-s.ctx.Done():
			fmt.Println("server done")
			return nil
		case <-ticker.C:
			fmt.Println("sending stats")
			if err := s.sendStats(server); err != nil {
				return err
			}
		}
	}
}

func (s *RPCHandler) sendStats(server pb.SystemStatsService_GetSystemStatsServer) error {
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

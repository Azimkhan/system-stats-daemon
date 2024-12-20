package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"github.com/Azimkhan/system-stats-daemon/internal/core/diskio"
	"github.com/Azimkhan/system-stats-daemon/internal/core/loadaverage"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
)

type StatType string

type StatFiller interface {
	Fill(*core.Stats) error
}

type StatService struct {
	currentStats    *core.Stats
	lastErr         error
	fillers         []StatFiller
	rwMutex         *sync.RWMutex
	collectInterval time.Duration
	log             logging.Logger
}

func getFiller(statType string) (StatFiller, error) {
	switch statType {
	case "cpuloadavg":
		return loadaverage.NewCollector(), nil
	case "diskio":
		return diskio.NewCollector(), nil
	default:
		return nil, ErrInvalidStatType
	}
}

func NewStatService(stats []string, collectInterval time.Duration, logger logging.Logger) (*StatService, error) {
	fillers := make([]StatFiller, 0, len(stats))
	for _, stat := range stats {
		filler, err := getFiller(stat)
		if err != nil {
			return nil, err
		}
		fillers = append(fillers, filler)
	}
	return &StatService{
		fillers:         fillers,
		rwMutex:         &sync.RWMutex{},
		collectInterval: collectInterval,
		log:             logger.With("service", "stat"),
	}, nil
}

// Run periodically collects stats from all fillers.
func (s *StatService) Run(ctx context.Context) {
	s.log.Debug("starting stat collection service")
	ticker := time.NewTicker(s.collectInterval)
	defer ticker.Stop()
	for {
		// collect stats immediately and then periodically
		s.collectStats()
		select {
		case <-ticker.C:
			s.collectStats()
			if s.lastErr == nil {
				s.log.Debug("stats collected", "stats", s.currentStats)
			} else {
				s.log.Warn("failed to collect stats", "error", s.lastErr)
			}
		case <-ctx.Done():
			return
		}
	}
}

// collectStats collects stats from all fillers and updates the current stats.
func (s *StatService) collectStats() {
	stats := &core.Stats{}
	// collect multiple stats concurrently and consider error handling
	errorsChan := make(chan error, len(s.fillers))
	for _, filler := range s.fillers {
		go func(filler StatFiller) {
			errorsChan <- filler.Fill(stats)
		}(filler)
	}

	// collect errors
	errs := make([]error, 0, len(s.fillers))
	for i := 0; i < len(s.fillers); i++ {
		errs = append(errs, <-errorsChan)
	}

	// join errors
	err := errors.Join(errs...)

	// update stats
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if err != nil {
		s.lastErr = &ErrCollectStats{Err: err}
		s.currentStats = nil
		return
	}
	s.lastErr = nil
	s.currentStats = stats
}

// GetStats returns the current stats.
func (s *StatService) GetStats() (*core.Stats, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.currentStats, s.lastErr
}

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const gRPCAddr = ":50052"

type StatsResponseHandler struct {
	responses chan *pb.SystemStatsResponse
}

func (s *StatsResponseHandler) Handle(resp *pb.SystemStatsResponse) error {
	s.responses <- resp
	return nil
}

type MainTestSuite struct {
	suite.Suite
	logger logging.Logger
}

func (s *MainTestSuite) SetupSuite() {
	var err error
	s.logger, err = logging.NewLogger(&config.LoggingConfig{
		Level:   "debug",
		Handler: "json",
	})
	require.NoError(s.T(), err)
}

func (s *MainTestSuite) TestCollect() {
	testData := []struct {
		name     string
		stats    []string
		validate func(context.Context, *StatsResponseHandler)
	}{
		{
			name:  "collect single stat",
			stats: []string{"cpuloadavg"},
			validate: func(timeout context.Context, responseHandler *StatsResponseHandler) {
				select {
				case <-timeout.Done():
					require.Fail(s.T(), "timeout")
				case resp := <-responseHandler.responses:
					require.NotNil(s.T(), resp)
					require.NotNil(s.T(), resp.CpuLoadAverage)
					require.Nil(s.T(), resp.DiskLoad)
				}
			},
		},

		{
			name:  "collect multiple stats",
			stats: []string{"cpuloadavg", "diskio"},

			validate: func(timeout context.Context, responseHandler *StatsResponseHandler) {
				select {
				case <-timeout.Done():
					require.Fail(s.T(), "timeout")
				case resp := <-responseHandler.responses:
					require.NotNil(s.T(), resp)
					require.NotNil(s.T(), resp.CpuLoadAverage)
					require.NotNil(s.T(), resp.DiskLoad)
				}
			},
		},
	}

	for _, tt := range testData {
		s.T().Run(tt.name, func(_ *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			serverApp, err := app.NewServerApp(ctx, tt.stats, &config.ServerConfig{
				BindAddr: gRPCAddr,
			}, &config.StreamingConfig{
				InitialDelay: 2 * time.Second,
				Interval:     10 * time.Second,
			}, s.logger)
			require.NoError(s.T(), err)
			defer serverApp.Stop()

			go func() {
				err := serverApp.Serve()
				require.NoError(s.T(), err)
			}()

			responseHandler := &StatsResponseHandler{
				responses: make(chan *pb.SystemStatsResponse, 3),
			}
			defer close(responseHandler.responses)
			clientApp, err := app.NewClientApp(gRPCAddr, 5*time.Second, responseHandler.Handle, s.logger)
			require.NoError(s.T(), err)
			defer func() {
				require.NoError(s.T(), clientApp.Close())
			}()

			go func() {
				err := clientApp.Run(ctx)
				// require.NoError(s.T(), err)
				if err != nil {
					s.logger.Error("clientApp.Run", "error", err)
				}
			}()

			// avg1, avg5, avg15, err := cpuLoadAvg()
			require.NoError(s.T(), err)

			timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer timeoutCancel()

			tt.validate(timeoutCtx, responseHandler)
			cancel()
		})
	}
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

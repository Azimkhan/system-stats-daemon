package integration

import (
	"context"
	"github.com/Azimkhan/system-stats-daemon/gen/systemstats/pb"
	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
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
func verifyCpuLoadAvg(t *testing.T, expected *CpuLoadAvg, actual []*pb.CPULoadAverage) {
	require.InDelta(t, expected.Avg1, actual[0].AverageLoad, 0.01)
	require.InDelta(t, expected.Avg5, actual[1].AverageLoad, 0.01)
	require.InDelta(t, expected.Avg15, actual[2].AverageLoad, 0.01)
}

func verifyDiskLoad(t *testing.T, load []*DiskIO, load2 []*pb.DiskLoad) {
	var diskMap = make(map[string]*pb.DiskLoad)
	for _, d := range load2 {
		diskMap[d.Device] = d
	}
	for _, d := range load {
		// allow 5% error
		require.NotNil(t, diskMap[d.Device])
		require.InDelta(t, d.TPS/diskMap[d.Device].TransactionsPerSecond, 1, 0.05)
		require.InDelta(t, d.Throughput/diskMap[d.Device].Throughput, 1, 0.05)
	}

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
				currentCpuLoad, err := cpuLoadAvg()
				require.NoError(s.T(), err)

				select {
				case <-timeout.Done():
					require.Fail(s.T(), "timeout")
				case resp := <-responseHandler.responses:
					require.NotNil(s.T(), resp)
					require.NotNil(s.T(), resp.CpuLoadAverage)
					require.Nil(s.T(), resp.DiskLoad)
					verifyCpuLoadAvg(s.T(), currentCpuLoad, resp.CpuLoadAverage)
				}
			},
		},

		{
			name:  "collect multiple stats",
			stats: []string{"cpuloadavg", "diskio"},

			validate: func(timeout context.Context, responseHandler *StatsResponseHandler) {
				currentCpuLoad, err := cpuLoadAvg()
				require.NoError(s.T(), err)

				currentDiskLoad, err := diskIO()
				require.NoError(s.T(), err)

				select {
				case <-timeout.Done():
					require.Fail(s.T(), "timeout")
				case resp := <-responseHandler.responses:
					require.NotNil(s.T(), resp)
					require.NotNil(s.T(), resp.CpuLoadAverage)
					require.NotNil(s.T(), resp.DiskLoad)
					verifyCpuLoadAvg(s.T(), currentCpuLoad, resp.CpuLoadAverage)
					verifyDiskLoad(s.T(), currentDiskLoad, resp.DiskLoad)
				}
			},
		},
	}

	for _, tt := range testData {
		s.T().Run(tt.name, func(t *testing.T) {
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
				//require.NoError(s.T(), err)
				if err != nil {
					s.logger.Error("clientApp.Run", "error", err)
				}
			}()

			//avg1, avg5, avg15, err := cpuLoadAvg()
			require.NoError(s.T(), err)

			validationCtx, vCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer vCancel()

			tt.validate(validationCtx, responseHandler)
			cancel()
		})
	}

}

func (s *MainTestSuite) TestCollectErrors() {

}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

//go:build linux

package loadaverage

import (
	"testing"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"github.com/stretchr/testify/require"
)

func TestLoadAverageCollector_Collect(t *testing.T) {
	collector := Collector{
		executeCommand: func() ([]byte, error) {
			return []byte("0.15 0.11 0.09 1/411 3200877\n"), nil
		},
	}
	expected := &core.CPULoadAverage{
		Rows: []*core.CPULoadAverageRow{
			{
				MinutesAgo: 1,
				Value:      0.15,
			},
			{
				MinutesAgo: 5,
				Value:      0.11,
			},
			{
				MinutesAgo: 15,
				Value:      0.09,
			},
		},
	}
	res, err := collector.Collect()
	require.NoError(t, err)
	require.Equal(t, expected, res)
}

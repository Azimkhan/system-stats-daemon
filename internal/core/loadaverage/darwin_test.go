//go:build darwin

package loadaverage

import (
	"testing"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"github.com/stretchr/testify/require"
)

func TestLoadAverageCollector_Collect(t *testing.T) {
	collector := CollectorImpl{
		executeCommand: func() ([]byte, error) {
			return []byte("{ 2.96 4.09 3.86 }\n"), nil
		},
	}
	expected := &core.CPULoadAverage{
		Rows: []*core.CPULoadAverageRow{
			{
				MinutesAgo: 1,
				Value:      2.96,
			},
			{
				MinutesAgo: 5,
				Value:      4.09,
			},
			{
				MinutesAgo: 15,
				Value:      3.86,
			},
		},
	}
	res, err := collector.Collect()
	require.NoError(t, err)
	require.Equal(t, expected, res)
}

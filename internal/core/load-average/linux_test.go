//go:build linux

package load_average

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadAverageCollector_Collect(t *testing.T) {
	collector := LoadAverageCollectorImpl{
		command: func() ([]byte, error) {
			return []byte(""), nil
		},
	}
	res, err := collector.Collect()
	require.NoError(t, err)
}

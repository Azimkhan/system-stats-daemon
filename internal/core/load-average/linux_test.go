//go:build linux

package load_average

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadAverageCollector_Collect(t *testing.T) {
	collector := LoadAverageCollectorImpl{
		executeCommand: func() ([]byte, error) {
			return []byte(""), nil
		},
	}
	_, err := collector.Collect()
	require.NoError(t, err)
}

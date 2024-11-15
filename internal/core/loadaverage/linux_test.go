//go:build linux

package loadaverage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAverageCollector_Collect(t *testing.T) {
	collector := CollectorImpl{
		executeCommand: func() ([]byte, error) {
			return []byte(""), nil
		},
	}
	_, err := collector.Collect()
	require.NoError(t, err)
}

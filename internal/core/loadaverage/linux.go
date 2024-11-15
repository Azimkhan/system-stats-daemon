//go:build linux

package loadaverage

import "github.com/Azimkhan/system-stats-daemon/internal/core"

type CollectorImpl struct {
	executeCommand func() ([]byte, error)
}

// cat /proc/loadavg
// 3.77 3.08 2.71 1/903 3597
func (l *CollectorImpl) Collect() (*core.CPULoadAverage, error) {
	// TODO implement
	return nil, nil
}

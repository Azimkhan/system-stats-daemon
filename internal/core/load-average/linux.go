//go:build linux

package load_average

import "github.com/Azimkhan/system-stats-daemon/internal/core"

type LoadAverageCollectorImpl struct {
}

func (l *LoadAverageCollectorImpl) Collect() (*core.CPULoadAverage, error) {
	panic("not implemented")
}

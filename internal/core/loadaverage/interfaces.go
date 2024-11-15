package loadaverage

import "github.com/Azimkhan/system-stats-daemon/internal/core"

type Collector interface {
	Collect() (*core.CPULoadAverage, error)
}

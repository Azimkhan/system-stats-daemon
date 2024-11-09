package load_average

import "github.com/Azimkhan/system-stats-daemon/internal/core"

type LoadAverageCollector interface {
	Collect() (*core.CPULoadAverage, error)
}

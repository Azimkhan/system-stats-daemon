package loadaverage

import (
	"errors"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

var ErrorInvalidOutput = errors.New("invalid command output")

type Collector interface {
	Collect() (*core.CPULoadAverage, error)
}

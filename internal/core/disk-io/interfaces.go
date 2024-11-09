package disk_io

import (
	"errors"
	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

var (
	InvalidOutputError = errors.New("invalid command output")
)

type DiskIOCollector interface {
	Collect() (*core.DiskIO, error)
}

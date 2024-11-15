//go:build linux

package loadaverage

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

type CollectorImpl struct {
	executeCommand func() ([]byte, error)
}

func NewCollector() Collector {
	return &CollectorImpl{
		executeCommand: executeCommand,
	}
}

func executeCommand() ([]byte, error) {
	return exec.Command("cat", "/proc/loadavg").Output()
}

func (l *CollectorImpl) Collect() (*core.CPULoadAverage, error) {
	output, err := l.executeCommand()
	if err != nil {
		return nil, err
	}
	// 0.15 0.11 0.09 1/411 3200877
	rawStr := strings.Trim(string(output), " \n")
	parts := strings.Split(rawStr, " ")
	if len(parts) != 5 {
		return nil, ErrorInvalidOutput
	}
	rows := make([]*core.CPULoadAverageRow, 3)
	for i, v := range parts[:3] {
		minutesAgo := []int{1, 5, 15}[i]
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, err
		}
		rows[i] = &core.CPULoadAverageRow{
			MinutesAgo: minutesAgo,
			Value:      float32(f),
		}
	}
	return &core.CPULoadAverage{Rows: rows}, nil
}

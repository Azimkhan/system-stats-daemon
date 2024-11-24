//go:build darwin

package diskio

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

func NewCollector() *Collector {
	return &Collector{
		executeCommand: executeIostat,
	}
}

func executeIostat() ([]byte, error) {
	return exec.Command("iostat", "-d").Output()
}

type Collector struct {
	executeCommand func() ([]byte, error)
}

func (d *Collector) Collect() (*core.DiskIO, error) {
	output, err := d.executeCommand()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 3 {
		return nil, ErrorInvalidOutput
	}

	// get disk names
	diskNames := strings.Fields(lines[0])

	// get stats
	rows := make([]*core.DiskIORow, 0, len(diskNames))
	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		if len(fields) != len(diskNames)*3 {
			continue
		}

		for i := 0; i < len(diskNames); i++ {
			tps, err := strconv.ParseFloat(fields[i*3+1], 32)
			if err != nil {
				return nil, err
			}
			kps, err := strconv.ParseFloat(fields[i*3], 32)
			if err != nil {
				return nil, err
			}
			row := &core.DiskIORow{
				Device:     diskNames[i],
				TPS:        float32(tps),
				Throughput: float32(kps),
			}
			rows = append(rows, row)
		}
	}
	return &core.DiskIO{Rows: rows}, nil
}

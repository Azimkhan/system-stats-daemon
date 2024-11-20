//go:build linux

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
	return exec.Command("iostat", "-d", "-k").Output()
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
	if len(lines) < 4 {
		return nil, ErrorInvalidOutput
	}

	// header indexes
	headerIndexes := make(map[string]int)
	headerLine := strings.Fields(lines[2])
	for i, header := range headerLine {
		headerIndexes[header] = i
	}

	// get stats
	rows := make([]*core.DiskIORow, 0, len(lines)-3)
	for _, line := range lines[3:] {
		fields := strings.Fields(line)
		if len(fields) != len(headerLine) {
			continue
		}

		tps, err := strconv.ParseFloat(fields[headerIndexes["tps"]], 32)
		if err != nil {
			return nil, err
		}
		readThroughput, err := strconv.ParseFloat(fields[headerIndexes["kB_read/s"]], 32)
		if err != nil {
			return nil, err
		}
		writeThroughput, err := strconv.ParseFloat(fields[headerIndexes["kB_wrtn/s"]], 32)
		if err != nil {
			return nil, err
		}
		totalThroughput := readThroughput + writeThroughput
		row := &core.DiskIORow{
			Device:     fields[headerIndexes["Device"]],
			TPS:        float32(tps),
			Throughput: float32(totalThroughput),
		}
		rows = append(rows, row)
	}
	return &core.DiskIO{Rows: rows}, nil
}

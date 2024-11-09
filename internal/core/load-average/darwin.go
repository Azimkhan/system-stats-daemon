//go:build darwin

package load_average

import (
	"github.com/Azimkhan/system-stats-daemon/internal/core"
	"os/exec"
	"strconv"
	"strings"
)

func NewLoadAverageCollector() LoadAverageCollector {
	return &LoadAverageCollectorImpl{
		executeCommand: executeSysctl,
	}
}

type LoadAverageCollectorImpl struct {
	executeCommand func() ([]byte, error)
}

func executeSysctl() ([]byte, error) {
	return exec.Command("sysctl", "-n", "vm.loadavg").Output()
}

func (l *LoadAverageCollectorImpl) Collect() (*core.CPULoadAverage, error) {
	output, err := l.executeCommand()
	if err != nil {
		return nil, err
	}
	// parse a string like { 2.51 2.72 2.84 }\n
	myStr := strings.Trim(string(output), "{} \n")
	var rows [3]core.CPULoadAverageRow
	for i, v := range strings.Split(myStr, " ") {
		minutesAgo := []int{1, 5, 15}[i]
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, err
		}
		rows[i] = core.CPULoadAverageRow{
			MinutesAgo: minutesAgo,
			Value:      float32(f),
		}

	}
	return &core.CPULoadAverage{Rows: rows[:]}, nil
}

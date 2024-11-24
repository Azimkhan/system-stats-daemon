package integration

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"time"
)

type DiskIO struct {
	Device     string
	TPS        float32
	Throughput float32
}

type CpuLoadAvg struct {
	Avg1  float32
	Avg5  float32
	Avg15 float32
}

func diskIO() ([]*DiskIO, error) {
	bootTime, err := host.BootTime()
	timeUp := uint64(time.Now().Unix()) - bootTime

	finalStats, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}

	var res []*DiskIO
	for name, finalStat := range finalStats {

		transactions := finalStat.ReadCount + finalStat.WriteCount
		kbPerTransaction := float32((finalStat.ReadBytes + finalStat.WriteBytes) / transactions / 1000)
		tps := float32(transactions / timeUp)
		res = append(res, &DiskIO{
			Device:     name,
			TPS:        tps,
			Throughput: kbPerTransaction,
		})
	}
	return res, nil
}

func cpuLoadAvg() (*CpuLoadAvg, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return &CpuLoadAvg{
		Avg1:  float32(avg.Load1),
		Avg5:  float32(avg.Load5),
		Avg15: float32(avg.Load15),
	}, nil
}

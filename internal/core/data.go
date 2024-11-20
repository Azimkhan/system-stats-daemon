package core

type CPULoadAverageRow struct {
	MinutesAgo uint32
	Value      float32
}
type CPULoadAverage struct {
	Rows []*CPULoadAverageRow
}

type CPUCurrentUsage struct {
	User   float32
	Idle   float32
	System float32
}

type DiskIORow struct {
	Device     string
	TPS        float32
	Throughput float32
}
type DiskIO struct {
	Rows []*DiskIORow
}

type Stats struct {
	CPULoadAverage  *CPULoadAverage
	CPUCurrentUsage *CPUCurrentUsage
	DiskLoad        *DiskIO
}

package core

type CPULoadAverageRow struct {
	MinutesAgo int
	Value      float32
}
type CPULoadAverage struct {
	Rows []CPULoadAverageRow
}

type CPUCurrentUsage struct {
	User   float32
	Idle   float32
	System float32
}

type Stats struct {
	CPULoadAverage  *CPULoadAverage
	CPUCurrentUsage *CPUCurrentUsage
}

package loadaverage

import "github.com/Azimkhan/system-stats-daemon/internal/core"

func (l *Collector) Fill(s *core.Stats) error {
	la, err := l.Collect()
	if err != nil {
		return err
	}
	s.CPULoadAverage = la
	return nil
}

package diskio

import "github.com/Azimkhan/system-stats-daemon/internal/core"

func (d *Collector) Fill(s *core.Stats) error {
	diskIO, err := d.Collect()
	if err != nil {
		return err
	}
	s.DiskLoad = diskIO
	return nil
}

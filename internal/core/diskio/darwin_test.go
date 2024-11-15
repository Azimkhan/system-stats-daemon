//go:build darwin

package diskio

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

const exampleOutput = `disk0               disk4               disk5
    KB/t  tps  MB/s     KB/t  tps  MB/s     KB/t  tps  MB/s
   13.00  224  2.84     4.02    0  0.00     4.02    0  0.00`

func TestDiskIOCollectorImpl_Collect(t *testing.T) {
	type fields struct {
		executeCommand func() ([]byte, error)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.DiskIO
		wantErr error
	}{
		{
			name: "Normal Run",
			fields: fields{
				executeCommand: func() ([]byte, error) {
					return []byte(exampleOutput), nil
				},
			},
			want: &core.DiskIO{
				Rows: []core.DiskIORow{
					{
						Device:     "disk0",
						TPS:        224,
						Throughput: 13,
					},
					{
						Device:     "disk4",
						TPS:        0,
						Throughput: 4.02,
					},
					{
						Device:     "disk5",
						TPS:        0,
						Throughput: 4.02,
					},
				},
			},
		},
		{
			name: "Invalid Output",
			fields: fields{
				executeCommand: func() ([]byte, error) {
					return []byte("malformed output"), nil
				},
			},
			want:    nil,
			wantErr: ErrorInvalidOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CollectorImpl{
				executeCommand: tt.fields.executeCommand,
			}
			got, err := d.Collect()
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("Collect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

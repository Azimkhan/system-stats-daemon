//go:build linux

package diskio

import (
	"testing"

	"github.com/Azimkhan/system-stats-daemon/internal/core"
)

const exampleOutput = `Linux 6.10.4-linuxkit (af4940bdc5a0)    11/10/24        _aarch64_       (10 CPU)

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
vda              14.08        44.46       172.86      4560.25    1230589    4784364  126216936
vdb               0.11         7.88         0.00         0.00     218072          0          0`

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
				Rows: []*core.DiskIORow{
					{
						Device:     "vda",
						TPS:        14.08,
						Throughput: 217.32,
					},
					{
						Device:     "vdb",
						TPS:        0.11,
						Throughput: 7.88,
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
			if err != nil {
				require.ErrorIs(t, err, tt.wantErr)
			}
			require.Equal(t, got, tt.want)
		})
	}
}

package godis

import (
	"reflect"
	"testing"
)

func TestRedis_ClusterAddSlots(t *testing.T) {

	type args struct {
		slots []int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterAddSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterAddSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterAddSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterCountKeysInSlot(t *testing.T) {

	type args struct {
		slot int
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterCountKeysInSlot(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterCountKeysInSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterCountKeysInSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterDelSlots(t *testing.T) {

	type args struct {
		slots []int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterDelSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterDelSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterDelSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterFailover(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterFailover()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterFailover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterFailover() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterFlushSlots(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterFlushSlots()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterFlushSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterFlushSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterForget(t *testing.T) {

	type args struct {
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterForget(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterForget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterForget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterGetKeysInSlot(t *testing.T) {

	type args struct {
		slot  int
		count int
	}
	tests := []struct {
		name string

		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterGetKeysInSlot(tt.args.slot, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterGetKeysInSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterGetKeysInSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ClusterInfo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(&Option{
				Host: "localhost",
				Port: 7000,
			})
			got, err := r.ClusterInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestRedis_ClusterKeySlot(t *testing.T) {

	type args struct {
		key string
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterKeySlot(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterKeySlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterKeySlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterMeet(t *testing.T) {

	type args struct {
		ip   string
		port int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterMeet(tt.args.ip, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterMeet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterMeet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterNodes(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterNodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterReplicate(t *testing.T) {

	type args struct {
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterReplicate(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterReplicate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterReplicate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterReset(t *testing.T) {

	type args struct {
		resetType Reset
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterReset(tt.args.resetType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSaveConfig(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSaveConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSaveConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSaveConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotImporting(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSetSlotImporting(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotImporting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotImporting() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotMigrating(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSetSlotMigrating(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotMigrating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotMigrating() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotNode(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSetSlotNode(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotStable(t *testing.T) {

	type args struct {
		slot int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSetSlotStable(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotStable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotStable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSlaves(t *testing.T) {

	type args struct {
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ClusterSlaves(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSlaves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSlaves() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSlots(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ClusterSlots",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(&Option{
				Host: "localhost",
				Port: 7000,
			})
			got, err := r.ClusterSlots()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

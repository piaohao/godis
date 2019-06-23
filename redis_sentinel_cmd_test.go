package godis

import (
	"reflect"
	"testing"
)

func TestRedis_SentinelFailover(t *testing.T) {

	type args struct {
		masterName string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "SentinelFailover",

			args: args{
				masterName: "a",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelFailover(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelFailover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelFailover() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_SentinelGetMasterAddrByName(t *testing.T) {

	type args struct {
		masterName string
	}
	tests := []struct {
		name string

		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "SentinelGetMasterAddrByName",

			args: args{
				masterName: "a",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelGetMasterAddrByName(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelGetMasterAddrByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelGetMasterAddrByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelMasters(t *testing.T) {

	tests := []struct {
		name string

		want    []map[string]string
		wantErr bool
	}{
		{
			name: "append",

			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelMasters()
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelMasters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelMasters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelMonitor(t *testing.T) {

	type args struct {
		masterName string
		ip         string
		port       int
		quorum     int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "append",

			args: args{
				masterName: "a",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelMonitor(tt.args.masterName, tt.args.ip, tt.args.port, tt.args.quorum)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelMonitor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelRemove(t *testing.T) {

	type args struct {
		masterName string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "append",

			args: args{
				masterName: "a",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelRemove(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelRemove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelRemove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelReset(t *testing.T) {

	type args struct {
		pattern string
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "append",

			args: args{
				pattern: "a",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelReset(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelSet(t *testing.T) {

	type args struct {
		masterName   string
		parameterMap map[string]string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "append",

			args: args{
				masterName: "a",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelSet(tt.args.masterName, tt.args.parameterMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelSlaves(t *testing.T) {

	type args struct {
		masterName string
	}
	tests := []struct {
		name string

		args    args
		want    []map[string]string
		wantErr bool
	}{
		{
			name: "append",

			args: args{
				masterName: "a",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SentinelSlaves(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelSlaves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelSlaves() got = %v, want %v", got, tt.want)
			}
		})
	}
}

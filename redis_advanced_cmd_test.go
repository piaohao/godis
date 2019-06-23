package godis

import (
	"reflect"
	"testing"
)

func TestRedis_ConfigGet(t *testing.T) {

	type args struct {
		pattern string
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
			got, err := r.ConfigGet(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ConfigSet(t *testing.T) {

	type args struct {
		parameter string
		value     string
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
			got, err := r.ConfigSet(tt.args.parameter, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfigSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogGet(t *testing.T) {

	type args struct {
		entries []int64
	}
	tests := []struct {
		name string

		args    args
		want    []Slowlog
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
			got, err := r.SlowlogGet(tt.args.entries...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlowlogGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogLen(t *testing.T) {

	tests := []struct {
		name string

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
			got, err := r.SlowlogLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlowlogLen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogReset(t *testing.T) {

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
			got, err := r.SlowlogReset()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlowlogReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectEncoding(t *testing.T) {

	type args struct {
		str string
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
			got, err := r.ObjectEncoding(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectEncoding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectIdletime(t *testing.T) {

	type args struct {
		str string
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
			got, err := r.ObjectIdletime(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectIdletime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectIdletime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectRefcount(t *testing.T) {

	type args struct {
		str string
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
			got, err := r.ObjectRefcount(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectRefcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectRefcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package godis

import (
	"reflect"
	"testing"
	"time"
)

func TestRedis_Keys(t *testing.T) {

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
			got, err := r.Keys(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Exists(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Exists(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Del(t *testing.T) {
	TestRedis_Set(t)

	type args struct {
		key []string
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "del",

			args: args{
				key: []string{"godis"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Del(tt.args.key...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Del() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Blpop(t *testing.T) {

	type args struct {
		args []string
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
			got, err := r.Blpop(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Blpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_BlpopTimout(t *testing.T) {

	type args struct {
		timeout int
		keys    []string
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
			got, err := r.BlpopTimout(tt.args.timeout, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlpopTimout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlpopTimout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Brpop(t *testing.T) {

	type args struct {
		args []string
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
			got, err := r.Brpop(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Brpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Brpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_BrpopTimout(t *testing.T) {

	type args struct {
		timeout int
		keys    []string
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
			got, err := r.BrpopTimout(tt.args.timeout, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrpopTimout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BrpopTimout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Mget(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Mget(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Mset(t *testing.T) {

	type args struct {
		keysvalues []string
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
			got, err := r.Mset(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Msetnx(t *testing.T) {

	type args struct {
		keysvalues []string
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
			got, err := r.Msetnx(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Msetnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Msetnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Rename(t *testing.T) {

	type args struct {
		oldkey string
		newkey string
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
			got, err := r.Rename(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rename() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Renamenx(t *testing.T) {

	type args struct {
		oldkey string
		newkey string
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
			got, err := r.Renamenx(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renamenx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Renamenx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Brpoplpush(t *testing.T) {

	type args struct {
		source      string
		destination string
		timeout     int
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
			got, err := r.Brpoplpush(tt.args.source, tt.args.destination, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Brpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Brpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sdiff(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Sdiff(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sdiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sdiff() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sdiffstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
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
			got, err := r.Sdiffstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sdiffstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sdiffstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sinter(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Sinter(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sinter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sinter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sinterstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
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
			got, err := r.Sinterstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Smove(t *testing.T) {

	type args struct {
		srckey string
		dstkey string
		member string
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
			got, err := r.Smove(tt.args.srckey, tt.args.dstkey, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Smove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Smove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sort(t *testing.T) {

	type args struct {
		key               string
		sortingParameters []SortingParams
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
			got, err := r.Sort(tt.args.key, tt.args.sortingParameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SortMulti(t *testing.T) {

	type args struct {
		key               string
		dstkey            string
		sortingParameters []SortingParams
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
			got, err := r.SortMulti(tt.args.key, tt.args.dstkey, tt.args.sortingParameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortMulti() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SortMulti() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sunion(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Sunion(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sunion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sunion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sunionstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
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
			got, err := r.Sunionstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Watch(t *testing.T) {

	type args struct {
		keys []string
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
			got, err := r.Watch(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Watch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Unwatch(t *testing.T) {

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
			got, err := r.Unwatch()
			if (err != nil) != tt.wantErr {
				t.Errorf("Unwatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unwatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zinterstore(t *testing.T) {

	type args struct {
		dstkey string
		sets   []string
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
			got, err := r.Zinterstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZinterstoreWithParams(t *testing.T) {

	type args struct {
		dstkey string
		params ZParams
		sets   []string
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
			got, err := r.ZinterstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZinterstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZinterstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zunionstore(t *testing.T) {

	type args struct {
		dstkey string
		sets   []string
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
			got, err := r.Zunionstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZunionstoreWithParams(t *testing.T) {

	type args struct {
		dstkey string
		params ZParams
		sets   []string
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
			got, err := r.ZunionstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZunionstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZunionstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Publish(t *testing.T) {

	type args struct {
		channel string
		message string
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
			got, err := r.Publish(tt.args.channel, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Publish() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Subscribe(t *testing.T) {
	flushAll()

	type args struct {
		redisPubSub *RedisPubSub
		channels    []string
	}
	pubsub := &RedisPubSub{
		OnMessage: func(channel, message string) {
			t.Logf("receive message ,channel:%s,message:%s", channel, message)
		},
		OnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive subscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnUnsubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive unsubscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnPMessage: func(pattern string, channel, message string) {
			t.Logf("receive pmessage ,pattern:%s,channel:%s,message:%s", pattern, channel, message)
		},
		OnPSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive psubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPUnsubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive punsubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPong: func(channel string) {
			t.Logf("receive pong ,channel:%s", channel)
		},
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "pubsub",

			args: args{
				redisPubSub: pubsub,
				channels:    []string{"godis"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				r := NewRedis(option)
				r.Connect()
				if err := r.Subscribe(tt.args.redisPubSub, tt.args.channels...); (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
			redis := NewRedis(option)
			redis.Publish("godis", "publish a message to godis channel")
			redis.Close()
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
		})
	}
}

func TestRedis_Psubscribe(t *testing.T) {

	type args struct {
		redisPubSub *RedisPubSub
		patterns    []string
	}
	tests := []struct {
		name string

		args    args
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
			if err := r.Psubscribe(tt.args.redisPubSub, tt.args.patterns...); (err != nil) != tt.wantErr {
				t.Errorf("Psubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_RandomKey(t *testing.T) {

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
			got, err := r.RandomKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandomKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bitop(t *testing.T) {

	type args struct {
		op      BitOP
		destKey string
		srcKeys []string
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
			got, err := r.Bitop(tt.args.op, tt.args.destKey, tt.args.srcKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Scan(t *testing.T) {

	type args struct {
		cursor string
		params []ScanParams
	}
	tests := []struct {
		name string

		args    args
		want    *ScanResult
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
			got, err := r.Scan(tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pfmerge(t *testing.T) {

	type args struct {
		destkey    string
		sourcekeys []string
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
			got, err := r.Pfmerge(tt.args.destkey, tt.args.sourcekeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfmerge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfmerge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pfcount(t *testing.T) {
	type args struct {
		keys []string
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
			got, err := r.Pfcount(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

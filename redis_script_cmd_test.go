package godis

import (
	"reflect"
	"testing"
)

func TestRedis_Eval(t *testing.T) {
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	type args struct {
		script   string
		keyCount int
		params   []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "eval",
			args: args{
				script:   `return redis.call("get",KEYS[1])`,
				keyCount: 1,
				params:   []string{"godis"},
			},
			want:    "good",
			wantErr: false,
		},
		{
			name: "eval",
			args: args{
				script:   `return redis.call("set",KEYS[1],ARGV[1])`,
				keyCount: 1,
				params:   []string{"eval", "godis"},
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "eval",
			args: args{
				script:   `return redis.call("get",KEYS[1])`,
				keyCount: 1,
				params:   []string{"eval"},
			},
			want:    "godis",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Eval(tt.args.script, tt.args.keyCount, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_EvalByKeyArgs(t *testing.T) {
	TestRedis_Set(t)

	type args struct {
		script string
		keys   []string
		args   []string
	}
	tests := []struct {
		name string

		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "eval",

			args: args{
				script: `return redis.call("get",KEYS[1])`,
				keys:   []string{"godis"},
				args:   []string{},
			},
			want:    "good",
			wantErr: false,
		},
		{
			name: "eval",

			args: args{
				script: `return redis.call("set",KEYS[1],ARGV[1])`,
				keys:   []string{"eval"},
				args:   []string{"godis"},
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "eval",

			args: args{
				script: `return redis.call("get",KEYS[1])`,
				keys:   []string{"eval"},
				args:   []string{},
			},
			want:    "godis",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.EvalByKeyArgs(tt.args.script, tt.args.keys, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Evalsha(t *testing.T) {

	type args struct {
		sha1     string
		keyCount int
		params   []string
	}
	tests := []struct {
		name string

		args    args
		want    interface{}
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
			got, err := r.Evalsha(tt.args.sha1, tt.args.keyCount, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evalsha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evalsha() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ScriptExists(t *testing.T) {

	type args struct {
		sha1 []string
	}
	tests := []struct {
		name string

		args    args
		want    []bool
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
			got, err := r.ScriptExists(tt.args.sha1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ScriptLoad(t *testing.T) {

	type args struct {
		script string
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
			got, err := r.ScriptLoad(tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ScriptLoad() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package godis

import (
	"testing"
	"time"
)

func TestRedis_Auth(t *testing.T) {

	type args struct {
		password string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "auth1",

			args:    args{password: ""},
			want:    "",
			wantErr: true,
		},
		{
			name: "auth2",

			args:    args{password: "1234567"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Auth(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Ping(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "ping",

			want:    "PONG",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Quit(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "append",

			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Quit()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quit() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_FlushDB(t *testing.T) {
	redis := NewRedis(option)
	_, err := redis.Set("godis", "good")
	if err != nil {
		t.Errorf("expect no error,but %v", err)
		return
	}
	redis.Close()
	redis = NewRedis(option)
	reply, _ := redis.Get("godis")
	if reply != "good" {
		t.Errorf("want good,but %s", reply)
		return
	}
	redis.Close()
	redis = NewRedis(option)
	redis.Select(2)
	_, err = redis.Set("godis", "good")
	if err != nil {
		t.Errorf("expect no error,but %v", err)
		return
	}
	redis.Close()

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "flushdb",

			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.FlushDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlushDB() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
	redis = NewRedis(option)
	reply, _ = redis.Get("godis")
	if reply != "" {
		t.Errorf("want empty string , but %s", reply)
		return
	}
	redis.Close()

	redis = NewRedis(option)
	redis.Select(2)
	reply, _ = redis.Get("godis")
	if reply != "good" {
		t.Errorf("want good , but %s", reply)
		return
	}
	redis.Close()
}

func TestRedis_FlushAll(t *testing.T) {
	redis := NewRedis(option)
	_, err := redis.Set("godis", "good")
	if err != nil {
		t.Errorf("expect no error,but %v", err)
		return
	}
	redis.Close()
	redis = NewRedis(option)
	reply, _ := redis.Get("godis")
	if reply != "good" {
		t.Errorf("want good,but %s", reply)
		return
	}
	redis.Close()
	redis = NewRedis(option)
	redis.Select(2)
	_, err = redis.Set("godis", "good")
	if err != nil {
		t.Errorf("expect no error,but %v", err)
		return
	}
	redis.Close()

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "flushall",

			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.FlushAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlushAll() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
	redis = NewRedis(option)
	reply, _ = redis.Get("godis")
	if reply != "" {
		t.Errorf("want empty string , but %s", reply)
		return
	}
	redis.Close()

	redis = NewRedis(option)
	redis.Select(2)
	reply, _ = redis.Get("godis")
	if reply != "" {
		t.Errorf("want empty string , but %s", reply)
		return
	}
	redis.Close()
}

func TestRedis_DbSize(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Set("godis1", "good")
	redis.Close()

	tests := []struct {
		name string

		want    int64
		wantErr bool
	}{
		{
			name: "dbsize",

			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.DbSize()
			if (err != nil) != tt.wantErr {
				t.Errorf("DbSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DbSize() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Select(t *testing.T) {

	type args struct {
		index int
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
				index: 15,
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "append",

			args: args{
				index: 16,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Select(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Save(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "save",

			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Save()
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Bgsave(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "Bgsave",

			want:    "Background saving started",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Bgsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bgsave() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Bgrewriteaof(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Bgrewriteaof",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Bgrewriteaof()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgrewriteaof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_Lastsave(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{
			name:    "lastsave",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Lastsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lastsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("last save time is ", got)
			r.Close()
		})
	}
}

// ignore this case,cause it will shutdown redis
func _TestRedis_Shutdown(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "shutdown",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Shutdown()
			if (err != nil) != tt.wantErr {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shutdown() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Info(t *testing.T) {
	//sleep 2 second,cause pre test case crash redis
	time.Sleep(2 * time.Second)

	type args struct {
		section []string
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "info",
			wantErr: false,
		},
		{
			name: "info",
			args: args{
				section: []string{"stats"},
			},
			wantErr: false,
		},
		{
			name: "info",
			args: args{
				section: []string{"clients", "memory"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Info(tt.args.section...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_Slaveof(t *testing.T) {

	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{
			name: "slaveof",

			args: args{
				host: "localhost",
				port: 6379,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Slaveof(tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("Slaveof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_SlaveofNoOne(t *testing.T) {
	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name:    "SlaveofNoOne",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.SlaveofNoOne()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlaveofNoOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_GetDB(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "GetDB",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			if got := r.GetDB(); got != tt.want {
				t.Errorf("GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Debug(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	type args struct {
		params DebugParams
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "debug",
			args: args{
				params: *NewDebugParamsObject("godis"),
			},
			wantErr: false,
		},
		{
			name: "debug",
			args: args{
				params: *NewDebugParamsReload(),
			},
			wantErr: false,
		},
		{
			name: "debug",
			args: args{
				params: *NewDebugParamsSegfault(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Debug(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Debug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_ConfigResetStat(t *testing.T) {
	//sleep 2 second,cause pre test case crash redis
	time.Sleep(2 * time.Second)

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		{
			name: "ConfigResetStat",

			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ConfigResetStat()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigResetStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			r.Close()
		})
	}
}

func TestRedis_WaitReplicas(t *testing.T) {

	type args struct {
		replicas int
		timeout  int64
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "WaitReplicas",

			args: args{
				replicas: 1,
				timeout:  1,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.WaitReplicas(tt.args.replicas, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaitReplicas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WaitReplicas() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

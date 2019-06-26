package godis

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"
)

//var connectionHandler = newRedisClusterConnectionHandler([]string{"localhost:7000", "localhost:7001", "localhost:7002", "localhost:7003", "localhost:7004", "localhost:7005"},
//	0, 0, "", &PoolConfig{})
var clusterOption = &ClusterOption{
	Nodes:             []string{"localhost:7000", "localhost:7001", "localhost:7002", "localhost:7003", "localhost:7004", "localhost:7005"},
	ConnectionTimeout: 5 * time.Second,
	SoTimeout:         5 * time.Second,
	MaxAttempts:       0,
	Password:          "",
	PoolConfig: &PoolConfig{
		MaxTotal: 100,
	},
}

func TestRedisCluster_Append(t *testing.T) {
	NewRedisCluster(clusterOption).Del("godis")
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "append",
			args: args{
				key:   "godis",
				value: "good",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Append(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Append() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Bitcount(t *testing.T) {
	NewRedisCluster(clusterOption).Set("godis", "good")
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Bitcount",

			args: args{
				key: "godis",
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Bitcount(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_BitcountRange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.BitcountRange(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitcountRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BitcountRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Bitfield(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		arguments []string
	}
	tests := []struct {
		name string

		args    args
		want    []int64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Bitfield(tt.args.key, tt.args.arguments...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitfield() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bitfield() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Bitop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Bitpos(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		value  bool
		params []BitPosParams
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Bitpos(tt.args.key, tt.args.value, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitpos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitpos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Blpop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_BlpopTimout(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Brpop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_BrpopTimout(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Brpoplpush(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Decr(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Decr(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_DecrBy(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		decrement int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.DecrBy(tt.args.key, tt.args.decrement)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Del(t *testing.T) {
	NewRedisCluster(clusterOption).Set("godis", "good")
	reply, _ := NewRedisCluster(clusterOption).Get("godis")
	if reply != "good" {
		t.Errorf("want reply good,but %s", reply)
		return
	}
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		keys []string
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
				keys: []string{"godis"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Del(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Del() got = %v, want %v", got, tt.want)
			}
		})
	}
	reply, _ = NewRedisCluster(clusterOption).Get("godis")
	if reply != "" {
		t.Errorf("want reply empty string,but %s", reply)
		return
	}
}

func TestRedisCluster_Echo(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Echo(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Echo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Eval(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		script   string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Eval(tt.args.script, tt.args.keyCount, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Evalsha(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Exists(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Expire(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		seconds int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Expire(tt.args.key, tt.args.seconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Expire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ExpireAt(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key      string
		unixtime int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ExpireAt(tt.args.key, tt.args.unixtime)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpireAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExpireAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Geoadd(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		longitude float64
		latitude  float64
		member    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Geoadd(tt.args.key, tt.args.longitude, tt.args.latitude, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geoadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Geoadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_GeoaddByMap(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key                 string
		memberCoordinateMap map[string]GeoCoordinate
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.GeoaddByMap(tt.args.key, tt.args.memberCoordinateMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoaddByMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoaddByMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Geodist(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		member1 string
		member2 string
		unit    []GeoUnit
	}
	tests := []struct {
		name string

		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Geodist(tt.args.key, tt.args.member1, tt.args.member2, tt.args.unit...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geodist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Geodist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Geohash(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		members []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Geohash(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geohash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Geohash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Geopos(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name string

		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Geopos(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geopos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Geopos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Georadius(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		longitude float64
		latitude  float64
		radius    float64
		unit      GeoUnit
		param     []GeoRadiusParam
	}
	tests := []struct {
		name string

		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Georadius(tt.args.key, tt.args.longitude, tt.args.latitude, tt.args.radius, tt.args.unit, tt.args.param...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Georadius() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Georadius() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_GeoradiusByMember(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		member string
		radius float64
		unit   GeoUnit
		param  []GeoRadiusParam
	}
	tests := []struct {
		name string

		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.GeoradiusByMember(tt.args.key, tt.args.member, tt.args.radius, tt.args.unit, tt.args.param...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoradiusByMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoradiusByMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Get(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_GetSet(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.GetSet(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Getbit(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		offset int64
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Getbit(tt.args.key, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getbit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getbit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Getrange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key         string
		startOffset int64
		endOffset   int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Getrange(tt.args.key, tt.args.startOffset, tt.args.endOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hdel(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		fields []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hdel(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hdel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hdel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hexists(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hexists(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hexists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hexists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hget(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hget(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_HgetAll(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
	}
	tests := []struct {
		name string

		args    args
		want    map[string]string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.HgetAll(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HgetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HgetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_HincrBy(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
		value int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.HincrBy(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_HincrByFloat(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
		value float64
	}
	tests := []struct {
		name string

		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.HincrByFloat(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrByFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrByFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hkeys(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hkeys(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hkeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hkeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hlen(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hlen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hlen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hlen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hmget(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		fields []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hmget(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hmget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hmget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hmset(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key  string
		hash map[string]string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hmset(tt.args.key, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hmset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hmset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hscan(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hset(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hset(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hsetnx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		field string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hsetnx(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hsetnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hsetnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Hvals(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Hvals(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hvals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hvals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Incr(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	cluster.Del("godis")
	for i := 0; i < 10000; i++ {
		cluster.Incr("godis")
	}
	reply, _ := cluster.Get("godis")
	if reply != "10000" {
		t.Errorf("want 10000,but %s", reply)
	}
}

func TestRedisCluster_IncrBy(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		increment int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.IncrBy(tt.args.key, tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_IncrByFloat(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key       string
		increment float64
	}
	tests := []struct {
		name string

		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.IncrByFloat(tt.args.key, tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncrByFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncrByFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Keys(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Lindex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		index int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lindex(tt.args.key, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lindex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lindex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Linsert(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		where ListOption
		pivot string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Linsert(tt.args.key, tt.args.where, tt.args.pivot, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Linsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Linsert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Llen(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Llen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Llen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Llen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lpop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lpop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lpush(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		strings []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lpush(tt.args.key, tt.args.strings...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lpushx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key  string
		strs []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lpushx(tt.args.key, tt.args.strs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpushx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpushx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lrange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		stop  int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lrange(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lrem(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		count int64
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lrem(tt.args.key, tt.args.count, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lrem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lrem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Lset(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		index int64
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Lset(tt.args.key, tt.args.index, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Ltrim(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		stop  int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Ltrim(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ltrim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ltrim() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Mget(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Move(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		dbIndex int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Move(tt.args.key, tt.args.dbIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Mset(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Msetnx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Persist(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Persist(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Persist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Persist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Pexpire(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key          string
		milliseconds int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Pexpire(tt.args.key, tt.args.milliseconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pexpire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pexpire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_PexpireAt(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key                   string
		millisecondsTimestamp int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.PexpireAt(tt.args.key, tt.args.millisecondsTimestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("PexpireAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PexpireAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Pfadd(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key      string
		elements []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Pfadd(tt.args.key, tt.args.elements...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Pfcount(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Pfmerge(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Psetex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key          string
		milliseconds int64
		value        string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Psetex(tt.args.key, tt.args.milliseconds, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Psetex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Psetex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Psubscribe(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			if err := r.Psubscribe(tt.args.redisPubSub, tt.args.patterns...); (err != nil) != tt.wantErr {
				t.Errorf("Psubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisCluster_Pttl(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Pttl(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pttl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pttl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Publish(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_RandomKey(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Rename(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Renamenx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Rpop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Rpop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Rpoplpush(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		srckey string
		dstkey string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Rpoplpush(tt.args.srckey, tt.args.dstkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Rpush(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		strings []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Rpush(tt.args.key, tt.args.strings...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Rpushx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key  string
		strs []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Rpushx(tt.args.key, tt.args.strs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpushx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpushx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Sadd(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		members []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Sadd(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Scan(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Scard(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Scard(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Scard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ScriptExists(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key  string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ScriptExists(tt.args.key, tt.args.sha1...)
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

func TestRedisCluster_ScriptLoad(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ScriptLoad(tt.args.key, tt.args.script)
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

func TestRedisCluster_Sdiff(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Sdiffstore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Set(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Set(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_SetWithParams(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
		nxxx  string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.SetWithParams(tt.args.key, tt.args.value, tt.args.nxxx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_SetWithParamsAndTime(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
		nxxx  string
		expx  string
		time  int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.SetWithParamsAndTime(tt.args.key, tt.args.value, tt.args.nxxx, tt.args.expx, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetWithParamsAndTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetWithParamsAndTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Setbit(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		offset int64
		value  string
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Setbit(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setbit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setbit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_SetbitWithBool(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		offset int64
		value  bool
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.SetbitWithBool(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetbitWithBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetbitWithBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Setex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		seconds int
		value   string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Setex(tt.args.key, tt.args.seconds, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Setnx(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		value string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Setnx(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Setrange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		offset int64
		value  string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Setrange(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Sinter(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Sinterstore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Sismember(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Sismember(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sismember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sismember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Smembers(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Smembers(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Smembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Smembers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Smove(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Sort(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_SortMulti(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Spop(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Spop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Spop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_SpopBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		count int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.SpopBatch(tt.args.key, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpopBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpopBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Srandmember(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Srandmember(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Srandmember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Srandmember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_SrandmemberBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.SrandmemberBatch(tt.args.key, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SrandmemberBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SrandmemberBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Srem(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key     string
		members []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Srem(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Srem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Srem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Sscan(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Sscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Strlen(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Strlen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strlen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Strlen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Subscribe(t *testing.T) {
	NewRedisCluster(clusterOption).Del("godis")
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
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Subscribe",
			args: args{
				redisPubSub: pubsub,
				channels:    []string{"godis"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func(tmp struct {
				name    string
				args    args
				wantErr bool
			}) {
				r := NewRedisCluster(clusterOption)
				if err := r.Subscribe(tt.args.redisPubSub, tt.args.channels...); (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}
			}(tt)
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
			NewRedisCluster(clusterOption).Publish("godis", "publish a message to godis channel")
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
		})
	}
}

func TestRedisCluster_Substr(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int
		end   int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Substr(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Substr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Substr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Sunion(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Sunionstore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Ttl(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Ttl(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ttl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ttl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Type(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Type(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Type() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Type() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Unwatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Watch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Zadd(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		score  float64
		member string
		params []ZAddParams
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zadd(tt.args.key, tt.args.score, tt.args.member, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZaddByMap(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key          string
		scoreMembers map[string]float64
		params       []ZAddParams
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZaddByMap(tt.args.key, tt.args.scoreMembers, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZaddByMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZaddByMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zcard(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zcard(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zcard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zcard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zcount(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zcount(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zincrby(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		score  float64
		member string
		params []ZAddParams
	}
	tests := []struct {
		name string

		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zincrby(tt.args.key, tt.args.score, tt.args.member, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zincrby() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zincrby() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zinterstore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_ZinterstoreWithParams(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_Zlexcount(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zlexcount(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zlexcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zlexcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zrange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zrange(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByLex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByLex(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByLexBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByLexBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByLexBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByLexBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByScore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByScore(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByScoreBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByScoreBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByScoreWithScores(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByScoreWithScores(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeByScoreWithScoresBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeByScoreWithScoresBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreWithScoresBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreWithScoresBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrangeWithScores(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrangeWithScores(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zrank(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zrank(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zrem(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		member []string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zrem(tt.args.key, tt.args.member...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZremrangeByLex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		min string
		max string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZremrangeByLex(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZremrangeByRank(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZremrangeByRank(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByRank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZremrangeByScore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start string
		end   string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZremrangeByScore(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zrevrange(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zrevrange(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrevrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zrevrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeByLex(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		max string
		min string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeByLex(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeByLexBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		max    string
		min    string
		offset int
		count  int
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeByLexBatch(tt.args.key, tt.args.max, tt.args.min, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByLexBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByLexBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeByScore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		max string
		min string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeByScore(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeByScoreWithScores(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key string
		max string
		min string
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeByScoreWithScores(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScoreWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScoreWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeByScoreWithScoresBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		max    string
		min    string
		offset int
		count  int
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeByScoreWithScoresBatch(tt.args.key, tt.args.max, tt.args.min, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScoreWithScoresBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScoreWithScoresBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_ZrevrangeWithScores(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name string

		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.ZrevrangeWithScores(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zrevrank(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zrevrank(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrevrank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrevrank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zscan(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zscore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name string

		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
			got, err := r.Zscore(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zscore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zscore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCluster_Zunionstore(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func TestRedisCluster_ZunionstoreWithParams(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
	}
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
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisCluster(clusterOption)
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

func Test_newRedisClusterCommand(t *testing.T) {
	type args struct {
		maxAttempts       int
		connectionHandler *redisClusterConnectionHandler
	}
	tests := []struct {
		name string
		args args
		want *redisClusterCommand
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRedisClusterCommand(tt.args.maxAttempts, tt.args.connectionHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedisClusterCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRedisClusterConnectionHandler(t *testing.T) {
	type args struct {
		nodes             []string
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
		poolConfig        *PoolConfig
	}
	tests := []struct {
		name string
		args args
		want *redisClusterConnectionHandler
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRedisClusterConnectionHandler(tt.args.nodes, tt.args.connectionTimeout, tt.args.soTimeout, tt.args.password, tt.args.poolConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedisClusterConnectionHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRedisClusterHashTagUtil(t *testing.T) {
	tests := []struct {
		name string
		want *redisClusterHashTagUtil
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRedisClusterHashTagUtil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedisClusterHashTagUtil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRedisClusterInfoCache(t *testing.T) {
	type args struct {
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
		poolConfig        *PoolConfig
	}
	tests := []struct {
		name string
		args args
		want *redisClusterInfoCache
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRedisClusterInfoCache(tt.args.connectionTimeout, tt.args.soTimeout, tt.args.password, tt.args.poolConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedisClusterInfoCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterCommand_releaseConnection(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
		ctx               context.Context
		execute           func(redis *Redis) (interface{}, error)
	}
	type args struct {
		redis *Redis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterCommand{
				MaxAttempts:       tt.fields.MaxAttempts,
				ConnectionHandler: tt.fields.ConnectionHandler,
				execute:           tt.fields.execute,
			}
			if err := r.releaseConnection(tt.args.redis); (err != nil) != tt.wantErr {
				t.Errorf("releaseConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisClusterCommand_run(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
		ctx               context.Context
		execute           func(redis *Redis) (interface{}, error)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterCommand{
				MaxAttempts:       tt.fields.MaxAttempts,
				ConnectionHandler: tt.fields.ConnectionHandler,
				execute:           tt.fields.execute,
			}
			got, err := r.run(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterCommand_runBatch(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
		ctx               context.Context
		execute           func(redis *Redis) (interface{}, error)
	}
	type args struct {
		keyCount int
		keys     []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterCommand{
				MaxAttempts:       tt.fields.MaxAttempts,
				ConnectionHandler: tt.fields.ConnectionHandler,
				execute:           tt.fields.execute,
			}
			got, err := r.runBatch(tt.args.keyCount, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("runBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterCommand_runWithAnyNode(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
		ctx               context.Context
		execute           func(redis *Redis) (interface{}, error)
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterCommand{
				MaxAttempts:       tt.fields.MaxAttempts,
				ConnectionHandler: tt.fields.ConnectionHandler,
				execute:           tt.fields.execute,
			}
			got, err := r.runWithAnyNode()
			if (err != nil) != tt.wantErr {
				t.Errorf("runWithAnyNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runWithAnyNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterCommand_runWithRetries(t *testing.T) {
	type fields struct {
		MaxAttempts       int
		ConnectionHandler *redisClusterConnectionHandler
		ctx               context.Context
		execute           func(redis *Redis) (interface{}, error)
	}
	type args struct {
		key           []byte
		attempts      int
		tryRandomNode bool
		asking        bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterCommand{
				MaxAttempts:       tt.fields.MaxAttempts,
				ConnectionHandler: tt.fields.ConnectionHandler,
				execute:           tt.fields.execute,
			}
			got, err := r.runWithRetries(tt.args.key, tt.args.attempts, tt.args.tryRandomNode, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("runWithRetries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runWithRetries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterConnectionHandler_getConnection(t *testing.T) {
	type fields struct {
		cache *redisClusterInfoCache
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Redis
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterConnectionHandler{
				cache: tt.fields.cache,
			}
			got, err := r.getConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterConnectionHandler_getConnectionFromNode(t *testing.T) {
	type fields struct {
		cache *redisClusterInfoCache
	}
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Redis
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterConnectionHandler{
				cache: tt.fields.cache,
			}
			got, err := r.getConnectionFromNode(tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnectionFromNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConnectionFromNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterConnectionHandler_getConnectionFromSlot(t *testing.T) {
	type fields struct {
		cache *redisClusterInfoCache
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Redis
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterConnectionHandler{
				cache: tt.fields.cache,
			}
			got, err := r.getConnectionFromSlot(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnectionFromSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConnectionFromSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterConnectionHandler_getNodes(t *testing.T) {
	type fields struct {
		cache *redisClusterInfoCache
	}
	tests := []struct {
		name   string
		fields fields
		want   sync.Map
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterConnectionHandler{
				cache: tt.fields.cache,
			}
			if got := r.getNodes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterConnectionHandler_renewSlotCache(t *testing.T) {
	type fields struct {
		cache *redisClusterInfoCache
	}
	type args struct {
		redis []*Redis
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &redisClusterConnectionHandler{
				cache: tt.fields.cache,
			}
		})
	}
}

func Test_redisClusterHashTagUtil_extractHashTag(t *testing.T) {
	type args struct {
		key                string
		returnKeyOnAbsence bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterHashTagUtil{}
			if got := r.extractHashTag(tt.args.key, tt.args.returnKeyOnAbsence); got != tt.want {
				t.Errorf("extractHashTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterHashTagUtil_getHashTag(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterHashTagUtil{}
			if got := r.getHashTag(tt.args.key); got != tt.want {
				t.Errorf("getHashTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterHashTagUtil_isClusterCompliantMatchPattern(t *testing.T) {
	type args struct {
		matchPattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterHashTagUtil{}
			if got := r.isClusterCompliantMatchPattern(tt.args.matchPattern); got != tt.want {
				t.Errorf("isClusterCompliantMatchPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_assignSlotToNode(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		slot int
		host string
		port int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
		})
	}
}

func Test_redisClusterInfoCache_assignSlotsToNode(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		lock  bool
		slots []int
		host  string
		port  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
		})
	}
}

func Test_redisClusterInfoCache_discoverClusterNodesAndSlots(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		redis *Redis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if err := r.discoverClusterNodesAndSlots(tt.args.redis); (err != nil) != tt.wantErr {
				t.Errorf("discoverClusterNodesAndSlots() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisClusterInfoCache_discoverClusterSlots(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		redis *Redis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if err := r.discoverClusterSlots(tt.args.redis); (err != nil) != tt.wantErr {
				t.Errorf("discoverClusterSlots() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisClusterInfoCache_generateHostAndPort(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		hostInfos []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  int
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			got, got1 := r.generateHostAndPort(tt.args.hostInfos)
			if got != tt.want {
				t.Errorf("generateHostAndPort() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("generateHostAndPort() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_redisClusterInfoCache_getAssignedSlotArray(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		slotInfo []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.getAssignedSlotArray(tt.args.slotInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAssignedSlotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_getNode(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		nodeKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Pool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.getNode(tt.args.nodeKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_getNodes(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	tests := []struct {
		name   string
		fields fields
		want   sync.Map
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.getNodes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_getShuffledNodesPool(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Pool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.getShuffledNodesPool(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShuffledNodesPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_getSlotPool(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Pool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.getSlotPool(tt.args.slot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSlotPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_renewClusterSlots(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		redis *Redis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if err := r.renewClusterSlots(tt.args.redis); (err != nil) != tt.wantErr {
				t.Errorf("renewClusterSlots() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisClusterInfoCache_reset(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		lock bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
		})
	}
}

func Test_redisClusterInfoCache_setupNodeIfNotExist(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		lock bool
		host string
		port int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Pool
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
			if got := r.setupNodeIfNotExist(tt.args.lock, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupNodeIfNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisClusterInfoCache_shuffle(t *testing.T) {
	type fields struct {
		nodes             sync.Map
		slots             sync.Map
		rwLock            sync.RWMutex
		rLock             sync.Mutex
		wLock             sync.Mutex
		rediscovering     bool
		poolConfig        *PoolConfig
		connectionTimeout time.Duration
		soTimeout         time.Duration
		password          string
	}
	type args struct {
		vals []*Pool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name: "append",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &redisClusterInfoCache{
				nodes:             tt.fields.nodes,
				slots:             tt.fields.slots,
				rwLock:            tt.fields.rwLock,
				rLock:             tt.fields.rLock,
				wLock:             tt.fields.wLock,
				rediscovering:     tt.fields.rediscovering,
				poolConfig:        tt.fields.poolConfig,
				connectionTimeout: tt.fields.connectionTimeout,
				soTimeout:         tt.fields.soTimeout,
				password:          tt.fields.password,
			}
		})
	}
}

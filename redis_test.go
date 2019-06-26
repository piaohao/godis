package godis

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

var option = &Option{
	//Host:     "10.1.1.63",
	//Password: "123456",
	Host:              "localhost",
	Port:              6379,
	Db:                0,
	ConnectionTimeout: 100 * time.Second,
	SoTimeout:         100 * time.Second,
}

// run before every test case ,ensure the redis is empty
func flushAll() {
	redis := NewRedis(option)
	defer redis.Close()
	redis.FlushAll()
}

func initDb() {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
}

func TestNewRedis(t *testing.T) {
	type args struct {
		option *Option
	}
	redis := new(Redis)
	redis.client = newClient(option)
	tests := []struct {
		name string
		args args
		want *Redis
	}{
		{
			name: "new",
			args: args{
				option: option,
			},
			want: redis,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedis(tt.args.option); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Append(t *testing.T) {
	flushAll()
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
				value: "very",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "append",
			args: args{
				key:   "godis",
				value: " good",
			},
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Append(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Append() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Asking(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "asking",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Asking()
			if (err != nil) != tt.wantErr {
				t.Errorf("Asking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Asking() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Bitcount(t *testing.T) {
	initDb()
	type args struct {
		key string
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
				key: "godis",
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Bitcount(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitcount() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_BitcountRange(t *testing.T) {
	initDb()
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "BitcountRange",
			args: args{
				key:   "godis",
				start: 0,
				end:   -1,
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.BitcountRange(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitcountRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BitcountRange() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Bitfield(t *testing.T) {
	initDb()
	type args struct {
		key       string
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "Bitfield",
			args: args{
				key:       "godis",
				arguments: []string{"INCRBY"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Bitfield(tt.args.key, tt.args.arguments...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitfield() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bitfield() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Bitpos(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "\x00\xff\xf0")
	redis.Close()
	type args struct {
		key    string
		value  bool
		params []BitPosParams
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Bitpos",
			args: args{
				key:   "godis",
				value: true,
				params: []BitPosParams{
					{params: [][]byte{IntToByteArray(0)}},
				},
			},
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Bitpos(tt.args.key, tt.args.value, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitpos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitpos() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Decr(t *testing.T) {
	flushAll()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Decr",
			args: args{
				key: "godis",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Decr",
			args: args{
				key: "godis",
			},
			want:    -2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Decr(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decr() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_DecrBy(t *testing.T) {
	flushAll()
	type args struct {
		key       string
		decrement int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "DecrBy",
			args: args{
				key:       "godis",
				decrement: 10,
			},
			want:    -10,
			wantErr: false,
		},
		{
			name: "DecrBy",
			args: args{
				key:       "godis",
				decrement: -10,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.DecrBy(tt.args.key, tt.args.decrement)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecrBy() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Echo(t *testing.T) {
	type args struct {
		string string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "echo",
			args: args{
				string: "godis",
			},
			want:    "godis",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Echo(tt.args.string)
			if (err != nil) != tt.wantErr {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Echo() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Expire(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	type args struct {
		key     string
		seconds int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "expire",
			args: args{
				key:     "godis",
				seconds: 1,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Expire(tt.args.key, tt.args.seconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Expire() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
	time.Sleep(2 * time.Second)
	redis = NewRedis(option)
	reply, _ := redis.Get("godis")
	if reply != "" {
		t.Errorf("want empty string ,but got %s", reply)
	}
	redis.Close()
}

func TestRedis_ExpireAt(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	type args struct {
		key      string
		unixtime int64
	}
	deadline := time.Now().Add(time.Second).Unix()
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "ExpireAt",
			args: args{
				key:      "godis",
				unixtime: deadline,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.ExpireAt(tt.args.key, tt.args.unixtime)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpireAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExpireAt() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
	time.Sleep(2 * time.Second)
	redis = NewRedis(option)
	reply, _ := redis.Get("godis")
	if reply != "" {
		t.Errorf("want empty string ,but got %s", reply)
	}
	redis.Close()
}

func TestRedis_Geoadd(t *testing.T) {
	flushAll()
	type args struct {
		key       string
		longitude float64
		latitude  float64
		member    string
	}
	tests := []struct {
		name    string
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

func TestRedis_GeoaddByMap(t *testing.T) {
	type args struct {
		key                 string
		memberCoordinateMap map[string]GeoCoordinate
	}
	tests := []struct {
		name    string
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

func TestRedis_Geodist(t *testing.T) {
	type args struct {
		key     string
		member1 string
		member2 string
		unit    []GeoUnit
	}
	tests := []struct {
		name    string
		args    args
		want    float64
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

func TestRedis_Geohash(t *testing.T) {
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
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

func TestRedis_Geopos(t *testing.T) {
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*GeoCoordinate
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

func TestRedis_Georadius(t *testing.T) {
	type args struct {
		key       string
		longitude float64
		latitude  float64
		radius    float64
		unit      GeoUnit
		param     []GeoRadiusParam
	}
	tests := []struct {
		name    string
		args    args
		want    []*GeoCoordinate
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

func TestRedis_GeoradiusByMember(t *testing.T) {
	type args struct {
		key    string
		member string
		radius float64
		unit   GeoUnit
		param  []GeoRadiusParam
	}
	tests := []struct {
		name    string
		args    args
		want    []*GeoCoordinate
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

func TestRedis_Get(t *testing.T) {
	flushAll()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get",
			args: args{
				key: "godis",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_GetSet(t *testing.T) {
	flushAll()
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getset",
			args: args{
				key:   "godis",
				value: "good",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "getset",
			args: args{
				key:   "godis",
				value: "good1",
			},
			want:    "good",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.GetSet(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSet() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Getbit(t *testing.T) {
	initDb()
	type args struct {
		key    string
		offset int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "getbit",
			args: args{
				key:    "godis",
				offset: 1,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Getbit(tt.args.key, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getbit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getbit() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Getrange(t *testing.T) {
	initDb()
	type args struct {
		key         string
		startOffset int64
		endOffset   int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getrange",
			args: args{
				key:         "godis",
				startOffset: 0,
				endOffset:   -1,
			},
			want:    "good",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
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

func TestRedis_Hdel(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
	type args struct {
		key    string
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "hdel",
			args: args{
				key:    "godis",
				fields: []string{"a"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "hdel",
			args: args{
				key:    "godis",
				fields: []string{"b"},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Hdel(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hdel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hdel() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Hexists(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "hexists",
			args: args{
				key:   "godis",
				field: "a",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "hexists",
			args: args{
				key:   "godis",
				field: "b",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
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

func TestRedis_Hget(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
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
		{
			name: "hget",
			args: args{
				key:   "godis",
				field: "a",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "hget",
			args: args{
				key:   "godis",
				field: "b",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.Hget(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hget() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_HgetAll(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "hgetall",
			args: args{
				key: "godis",
			},
			want:    map[string]string{"a": "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.HgetAll(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HgetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HgetAll() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_HincrBy(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
	type args struct {
		key   string
		field string
		value int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "hincrby",
			args: args{
				key:   "godis",
				field: "a",
				value: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "hincrby",
			args: args{
				key:   "godis",
				field: "b",
				value: 5,
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.HincrBy(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrBy() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_HincrByFloat(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Hset("godis", "a", "1")
	redis.Close()
	type args struct {
		key   string
		field string
		value float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "HincrByFloat",
			args: args{
				key:   "godis",
				field: "a",
				value: 1.5,
			},
			want:    2.5,
			wantErr: false,
		},
		{
			name: "HincrByFloat",
			args: args{
				key:   "godis",
				field: "b",
				value: 5.0987,
			},
			want:    5.0987,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			got, err := r.HincrByFloat(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrByFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrByFloat() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Hkeys(t *testing.T) {

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

func TestRedis_Hlen(t *testing.T) {

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

func TestRedis_Hmget(t *testing.T) {

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

func TestRedis_Hmset(t *testing.T) {

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

func TestRedis_Hscan(t *testing.T) {

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

func TestRedis_Hset(t *testing.T) {

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

func TestRedis_Hsetnx(t *testing.T) {

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

func TestRedis_Hvals(t *testing.T) {

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

func TestRedis_Incr(t *testing.T) {
	flushAll()
	pool := NewPool(nil, option)
	i := 0
	for ; i < 10000; i++ {
		redis, err := pool.Get()
		if err != nil {
			t.Errorf("err happen,%v", err)
			return
		}
		_, err = redis.Incr("godis")
		if err != nil {
			t.Errorf("err happen,%v", err)
			return
		}
		redis.Close()
	}
	redis, err := pool.Get()
	if err != nil {
		t.Errorf("err happen,%v", err)
		return
	}
	reply, err := redis.Get("godis")
	if err != nil {
		t.Errorf("err happen,%v", err)
		return
	}
	if reply != "10000" {
		t.Errorf("want 10000,but %s", reply)
	}
	redis.Close()
}

func TestRedis_IncrBy(t *testing.T) {
	flushAll()
	pool := NewPool(nil, option)
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	for i := 0; i < 10000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
			redis, err := pool.Get()
			if err != nil {
				t.Errorf("err happen,%v", err)
				return
			}
			_, err = redis.IncrBy("godis", 2)
			if err != nil {
				t.Errorf("err happen,%v", err)
				return
			}
			redis.Close()
			<-ch
		}()
	}
	group.Wait()
	redis, err := pool.Get()
	if err != nil {
		t.Errorf("err happen,%v", err)
		return
	}
	reply, err := redis.Get("godis")
	if err != nil {
		t.Errorf("err happen,%v", err)
		return
	}
	if reply != "20000" {
		t.Errorf("want 20000,but %s", reply)
	}
	redis.Close()
}

func TestRedis_IncrByFloat(t *testing.T) {

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

func TestRedis_Lindex(t *testing.T) {

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

func TestRedis_Linsert(t *testing.T) {

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

func TestRedis_Llen(t *testing.T) {

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

func TestRedis_Lpop(t *testing.T) {

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

func TestRedis_Lpush(t *testing.T) {

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

func TestRedis_Lpushx(t *testing.T) {

	type args struct {
		key    string
		string []string
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
			got, err := r.Lpushx(tt.args.key, tt.args.string...)
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

func TestRedis_Lrange(t *testing.T) {

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

func TestRedis_Lrem(t *testing.T) {

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

func TestRedis_Lset(t *testing.T) {

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

func TestRedis_Ltrim(t *testing.T) {

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

func TestRedis_Move(t *testing.T) {

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

func TestRedis_Multi(t *testing.T) {

	tests := []struct {
		name string

		want    *transaction
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
			got, err := r.Multi()
			if (err != nil) != tt.wantErr {
				t.Errorf("Multi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Persist(t *testing.T) {

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

func TestRedis_Pexpire(t *testing.T) {

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

func TestRedis_PexpireAt(t *testing.T) {

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

func TestRedis_Pfadd(t *testing.T) {

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

func TestRedis_Psetex(t *testing.T) {

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

func TestRedis_Pttl(t *testing.T) {

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

func TestRedis_PubsubChannels(t *testing.T) {

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
			got, err := r.PubsubChannels(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("PubsubChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PubsubChannels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Readonly(t *testing.T) {

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
			got, err := r.Readonly()
			if (err != nil) != tt.wantErr {
				t.Errorf("Readonly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Readonly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Receive(t *testing.T) {

	tests := []struct {
		name string

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
			if err := r.Receive(); (err != nil) != tt.wantErr {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Rpop(t *testing.T) {

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

func TestRedis_Rpoplpush(t *testing.T) {

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

func TestRedis_Rpush(t *testing.T) {

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

func TestRedis_Rpushx(t *testing.T) {

	type args struct {
		key    string
		string []string
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
			got, err := r.Rpushx(tt.args.key, tt.args.string...)
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

func TestRedis_Sadd(t *testing.T) {

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

func TestRedis_Scard(t *testing.T) {

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

func TestRedis_Send(t *testing.T) {

	type args struct {
		command protocolCommand
		args    [][]byte
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
			if err := r.Send(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_SendByStr(t *testing.T) {

	type args struct {
		command string
		args    [][]byte
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
			if err := r.SendByStr(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("SendByStr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Set(t *testing.T) {

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
		{
			name: "set",

			args: args{
				key:   "godis",
				value: "good",
			},
			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Set(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_SetWithParams(t *testing.T) {

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

func TestRedis_SetWithParamsAndTime(t *testing.T) {

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

func TestRedis_Setbit(t *testing.T) {

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

func TestRedis_SetbitWithBool(t *testing.T) {

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

func TestRedis_Setex(t *testing.T) {

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

func TestRedis_Setnx(t *testing.T) {

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

func TestRedis_Setrange(t *testing.T) {

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

func TestRedis_Sismember(t *testing.T) {

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

func TestRedis_Smembers(t *testing.T) {

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

func TestRedis_Spop(t *testing.T) {

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

func TestRedis_SpopBatch(t *testing.T) {

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

func TestRedis_Srandmember(t *testing.T) {

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

func TestRedis_SrandmemberBatch(t *testing.T) {

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

func TestRedis_Srem(t *testing.T) {

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

func TestRedis_Sscan(t *testing.T) {

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

func TestRedis_Strlen(t *testing.T) {

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

func TestRedis_Substr(t *testing.T) {

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

func TestRedis_Ttl(t *testing.T) {

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

func TestRedis_Type(t *testing.T) {

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

func TestRedis_Zadd(t *testing.T) {

	type args struct {
		key     string
		score   float64
		member  string
		mparams []ZAddParams
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
			got, err := r.Zadd(tt.args.key, tt.args.score, tt.args.member, tt.args.mparams...)
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

func TestRedis_ZaddByMap(t *testing.T) {

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

func TestRedis_Zcard(t *testing.T) {

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

func TestRedis_Zcount(t *testing.T) {

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

func TestRedis_Zincrby(t *testing.T) {

	type args struct {
		key       string
		increment float64
		member    string
		params    []ZAddParams
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
			got, err := r.Zincrby(tt.args.key, tt.args.increment, tt.args.member, tt.args.params...)
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

func TestRedis_Zlexcount(t *testing.T) {

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

func TestRedis_Zrange(t *testing.T) {

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
			got, err := r.Zrange(tt.args.key, tt.args.start, tt.args.stop)
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

func TestRedis_ZrangeByLex(t *testing.T) {

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

func TestRedis_ZrangeByLexBatch(t *testing.T) {

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

func TestRedis_ZrangeByScore(t *testing.T) {

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

func TestRedis_ZrangeByScoreBatch(t *testing.T) {

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

func TestRedis_ZrangeByScoreWithScores(t *testing.T) {

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

func TestRedis_ZrangeByScoreWithScoresBatch(t *testing.T) {

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

func TestRedis_ZrangeWithScores(t *testing.T) {

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

func TestRedis_Zrank(t *testing.T) {

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

func TestRedis_Zrem(t *testing.T) {

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
			got, err := r.Zrem(tt.args.key, tt.args.members...)
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

func TestRedis_ZremrangeByLex(t *testing.T) {

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

func TestRedis_ZremrangeByRank(t *testing.T) {

	type args struct {
		key   string
		start int64
		stop  int64
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
			got, err := r.ZremrangeByRank(tt.args.key, tt.args.start, tt.args.stop)
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

func TestRedis_ZremrangeByScore(t *testing.T) {

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

func TestRedis_Zrevrange(t *testing.T) {

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
			got, err := r.Zrevrange(tt.args.key, tt.args.start, tt.args.stop)
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

func TestRedis_ZrevrangeByLex(t *testing.T) {

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

func TestRedis_ZrevrangeByLexBatch(t *testing.T) {

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

func TestRedis_ZrevrangeByScore(t *testing.T) {

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

func TestRedis_ZrevrangeByScoreWithScores(t *testing.T) {

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

func TestRedis_ZrevrangeByScoreWithScoresBatch(t *testing.T) {

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

func TestRedis_ZrevrangeWithScores(t *testing.T) {

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

func TestRedis_Zrevrank(t *testing.T) {

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

func TestRedis_Zscan(t *testing.T) {

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

func TestRedis_Zscore(t *testing.T) {

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

func TestRedis_checkIsInMultiOrPipeline(t *testing.T) {

	tests := []struct {
		name string

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
			if err := r.checkIsInMultiOrPipeline(); (err != nil) != tt.wantErr {
				t.Errorf("checkIsInMultiOrPipeline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

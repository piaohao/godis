package godis

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sync"
	"testing"
	"time"
)

var option = &Option{
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
}

func TestRedis_GeoaddByMap(t *testing.T) {
}

func TestRedis_Geodist(t *testing.T) {
}

func TestRedis_Geohash(t *testing.T) {
}

func TestRedis_Geopos(t *testing.T) {
}

func TestRedis_Georadius(t *testing.T) {
}

func TestRedis_GeoradiusByMember(t *testing.T) {
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
	defer redis.Close()
	redis.Hset("godis", "a", "1")
	ret, err := redis.HincrByFloat("godis", "a", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 2.5, ret)

	ret, err = redis.HincrByFloat("godis", "b", 5.0987)
	assert.Nil(t, err)
	assert.Equal(t, 5.0987, ret)
}

func TestRedis_Hkeys(t *testing.T) {
}

func TestRedis_Hlen(t *testing.T) {
}

func TestRedis_Hmget(t *testing.T) {
}

func TestRedis_Hmset(t *testing.T) {
}

func TestRedis_Hscan(t *testing.T) {
}

func TestRedis_Hset(t *testing.T) {
}

func TestRedis_Hsetnx(t *testing.T) {
}

func TestRedis_Hvals(t *testing.T) {
}

func TestRedis_Incr(t *testing.T) {
	flushAll()
	pool := NewPool(nil, option)
	i := 0
	for ; i < 10000; i++ {
		redis, err := pool.GetResource()
		if err != nil {
			assert.Errorf(t, err, "err happen")
			return
		}
		_, err = redis.Incr("godis")
		assert.Nil(t, err)
		redis.Close()
	}
	redis, err := pool.GetResource()
	if err != nil {
		assert.Errorf(t, err, "err happen")
		return
	}
	reply, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "10000", reply)
	redis.Close()
}

func TestRedis_IncrBy(t *testing.T) {
	flushAll()
	pool := NewPool(nil, option)
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	for i := 0; i < 100000; i++ {
		group.Add(1)
		ch <- true
		go func() {
			defer group.Done()
			redis, err := pool.GetResource()
			if err != nil {
				assert.Errorf(t, err, "err happen")
				return
			}
			_, err = redis.IncrBy("godis", 2)
			assert.Nil(t, err)
			redis.Close()
			<-ch
		}()
	}
	group.Wait()
	redis, err := pool.GetResource()
	if err != nil {
		assert.Errorf(t, err, "err happen")
		return
	}
	reply, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "200000", reply)
	redis.Close()
}

func TestRedis_IncrByFloat(t *testing.T) {
}

func TestRedis_Lindex(t *testing.T) {
}

func TestRedis_Linsert(t *testing.T) {
}

func TestRedis_Llen(t *testing.T) {
}

func TestRedis_Lpop(t *testing.T) {
}

func TestRedis_Lpush(t *testing.T) {
}

func TestRedis_Lpushx(t *testing.T) {
}

func TestRedis_Lrange(t *testing.T) {
}

func TestRedis_Lrem(t *testing.T) {
}

func TestRedis_Lset(t *testing.T) {
}

func TestRedis_Ltrim(t *testing.T) {
}

func TestRedis_Move(t *testing.T) {
}

func TestRedis_Multi(t *testing.T) {
}

func TestRedis_Persist(t *testing.T) {
}

func TestRedis_Pexpire(t *testing.T) {
}

func TestRedis_PexpireAt(t *testing.T) {
}

func TestRedis_Pfadd(t *testing.T) {
}

func TestRedis_Psetex(t *testing.T) {
}

func TestRedis_Pttl(t *testing.T) {
}

func TestRedis_PubsubChannels(t *testing.T) {
}

func TestRedis_Readonly(t *testing.T) {
}

func TestRedis_Receive(t *testing.T) {
}

func TestRedis_Rpop(t *testing.T) {
}

func TestRedis_Rpoplpush(t *testing.T) {
}

func TestRedis_Rpush(t *testing.T) {
}

func TestRedis_Rpushx(t *testing.T) {
}

func TestRedis_Sadd(t *testing.T) {
}

func TestRedis_Scard(t *testing.T) {
}

func TestRedis_Send(t *testing.T) {
}

func TestRedis_SendByStr(t *testing.T) {
}

func TestRedis_Set(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	ret, err := redis.Set("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)
	redis.Close()
}

func TestRedis_SetWithParams(t *testing.T) {
}

func TestRedis_SetWithParamsAndTime(t *testing.T) {
}

func TestRedis_Setbit(t *testing.T) {
}

func TestRedis_SetbitWithBool(t *testing.T) {
}

func TestRedis_Setex(t *testing.T) {
}

func TestRedis_Setnx(t *testing.T) {
}

func TestRedis_Setrange(t *testing.T) {
}

func TestRedis_Sismember(t *testing.T) {
}

func TestRedis_Smembers(t *testing.T) {
}

func TestRedis_Spop(t *testing.T) {
}

func TestRedis_SpopBatch(t *testing.T) {
}

func TestRedis_Srandmember(t *testing.T) {
}

func TestRedis_SrandmemberBatch(t *testing.T) {
}

func TestRedis_Srem(t *testing.T) {
}

func TestRedis_Sscan(t *testing.T) {
}

func TestRedis_Strlen(t *testing.T) {
}

func TestRedis_Substr(t *testing.T) {
}

func TestRedis_Ttl(t *testing.T) {
}

func TestRedis_Type(t *testing.T) {
}

func TestRedis_Zadd(t *testing.T) {
}

func TestRedis_ZaddByMap(t *testing.T) {
}

func TestRedis_Zcard(t *testing.T) {
}

func TestRedis_Zcount(t *testing.T) {
}

func TestRedis_Zincrby(t *testing.T) {
}

func TestRedis_Zlexcount(t *testing.T) {
}

func TestRedis_Zrange(t *testing.T) {
}

func TestRedis_ZrangeByLex(t *testing.T) {
}

func TestRedis_ZrangeByLexBatch(t *testing.T) {
}

func TestRedis_ZrangeByScore(t *testing.T) {
}

func TestRedis_ZrangeByScoreBatch(t *testing.T) {
}

func TestRedis_ZrangeByScoreWithScores(t *testing.T) {
}

func TestRedis_ZrangeByScoreWithScoresBatch(t *testing.T) {
}

func TestRedis_ZrangeWithScores(t *testing.T) {
}

func TestRedis_Zrank(t *testing.T) {
}

func TestRedis_Zrem(t *testing.T) {
}

func TestRedis_ZremrangeByLex(t *testing.T) {
}

func TestRedis_ZremrangeByRank(t *testing.T) {
}

func TestRedis_ZremrangeByScore(t *testing.T) {
}

func TestRedis_Zrevrange(t *testing.T) {
}

func TestRedis_ZrevrangeByLex(t *testing.T) {
}

func TestRedis_ZrevrangeByLexBatch(t *testing.T) {
}

func TestRedis_ZrevrangeByScore(t *testing.T) {
}

func TestRedis_ZrevrangeByScoreWithScores(t *testing.T) {
}

func TestRedis_ZrevrangeByScoreWithScoresBatch(t *testing.T) {
}

func TestRedis_ZrevrangeWithScores(t *testing.T) {
}

func TestRedis_Zrevrank(t *testing.T) {
}

func TestRedis_Zscan(t *testing.T) {
}

func TestRedis_Zscore(t *testing.T) {
}

func TestRedis_checkIsInMultiOrPipeline(t *testing.T) {
}

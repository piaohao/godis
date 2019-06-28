package godis

import (
	"github.com/stretchr/testify/assert"
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
	cluster := NewRedisCluster(clusterOption)
	_, _ = cluster.Del("godis")
	count, err := cluster.Append("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

func TestRedisCluster_Bitcount(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	_, _ = cluster.Set("godis", "good")
	count, err := cluster.Bitcount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(20), count)
}

func TestRedisCluster_BitcountRange(t *testing.T) {
}

func TestRedisCluster_Bitfield(t *testing.T) {
}

func TestRedisCluster_Bitop(t *testing.T) {
}

func TestRedisCluster_Bitpos(t *testing.T) {
}

func TestRedisCluster_Blpop(t *testing.T) {
}

func TestRedisCluster_BlpopTimout(t *testing.T) {
}

func TestRedisCluster_Brpop(t *testing.T) {
}

func TestRedisCluster_BrpopTimout(t *testing.T) {
}

func TestRedisCluster_Brpoplpush(t *testing.T) {
}

func TestRedisCluster_Decr(t *testing.T) {
}

func TestRedisCluster_DecrBy(t *testing.T) {
}

func TestRedisCluster_Del(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	_, _ = cluster.Set("godis", "good")
	count, err := cluster.Del("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)
	str, err := cluster.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", str)
}

func TestRedisCluster_Echo(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	s, err := redis.Echo("godis")
	assert.Nil(t, err)
	assert.Equal(t, "godis", s)
}

func TestRedisCluster_Eval(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Eval(`return redis.call("get",KEYS[1])`, 1, "godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)

	s, err = redis.Eval(`return redis.call("set",KEYS[1],ARGV[1])`, 1, "eval", "godis")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	s, err = redis.Eval(`return redis.call("get",KEYS[1])`, 1, "eval")
	assert.Nil(t, err)
	assert.Equal(t, "godis", s)
}

func TestRedisCluster_Exists(t *testing.T) {
}

func TestRedisCluster_Expire(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Expire("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedisCluster_ExpireAt(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.ExpireAt("godis", time.Now().Add(1*time.Second).Unix())
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedisCluster_Geo(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	c, err := redis.Geoadd("godis", 121, 37, "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.GeoaddByMap("godis", map[string]GeoCoordinate{
		"b": {
			longitude: 122,
			latitude:  37,
		},
		"c": {
			longitude: 123,
			latitude:  37,
		},
		"d": {
			longitude: 124,
			latitude:  37,
		},
		"e": {
			longitude: 125,
			latitude:  37,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	arr, err := redis.Geohash("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"www43rts870"}, arr)

	_, err = redis.Geopos("godis", "b")
	assert.Nil(t, err)

	d, err := redis.Geodist("godis", "a", "b", GeounitKm)
	assert.Nil(t, err)
	assert.Equal(t, 88.8291, d)

	resp, err := redis.Georadius("godis", 121, 37, 500, GeounitKm,
		NewGeoRadiusParam().WithCoord().WithDist())
	assert.Nil(t, err)
	t.Log(resp)

	resp, err = redis.GeoradiusByMember("godis", "a", 500, GeounitKm,
		NewGeoRadiusParam().WithCoord().WithDist())
	assert.Nil(t, err)
	t.Log(resp)
}

func TestRedisCluster_Get(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedisCluster_GetSet(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	s, err := redis.GetSet("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "", s)
	s, err = redis.GetSet("godis", "good1")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedisCluster_Getbit(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Getbit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, s)
}

func TestRedisCluster_Getrange(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Getrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
	s, err = redis.Getrange("godis", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, "go", s)
}

func TestRedisCluster_Hdel(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hdel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	s, err = redis.Hdel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedisCluster_Hexists(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hexists("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, true, s)

	s, err = redis.Hexists("godis", "b")
	assert.Nil(t, err)
	assert.Equal(t, false, s)
}

func TestRedisCluster_Hget(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)
}

func TestRedisCluster_HgetAll(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.HgetAll("godis")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, s)
}

func TestRedisCluster_HincrBy(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.HincrBy("godis", "a", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	s, err = redis.HincrBy("godis", "b", 5)
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)
}

func TestRedisCluster_HincrByFloat(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")
	ret, err := redis.HincrByFloat("godis", "a", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 2.5, ret)

	ret, err = redis.HincrByFloat("godis", "b", 5.0987)
	assert.Nil(t, err)
	assert.Equal(t, 5.0987, ret)
}

func TestRedisCluster_Hkeys(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hkeys("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a"}, s)
}

func TestRedisCluster_Hlen(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedisCluster_Hmget(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hmget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
}

func TestRedisCluster_Hmset(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hmset("godis", map[string]string{"b": "2", "c": "3"})
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, _ := redis.Hlen("godis")
	assert.Equal(t, int64(3), c)
}

func TestRedisCluster_Hscan(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
}

func TestRedisCluster_Hset(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	ret, err := redis.Hset("godis", "a", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), ret)

	s, err := redis.Hlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedisCluster_Hsetnx(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	ret, err := redis.Hsetnx("godis", "a", "2")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ret)

	s, err = redis.Hget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)
}

func TestRedisCluster_Hvals(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	redis.Hset("godis", "a", "1")

	s, err := redis.Hvals("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
}

func TestRedisCluster_Incr(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	cluster.Del("godis")
	for i := 0; i < 10000; i++ {
		cluster.Incr("godis")
	}
	reply, _ := cluster.Get("godis")
	assert.Equal(t, "10000", reply)
}

func TestRedisCluster_IncrBy(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	for i := 0; i < 100000; i++ {
		group.Add(1)
		ch <- true
		go func() {
			defer group.Done()
			_, err := redis.IncrBy("godis", 2)
			assert.Nil(t, err)
			<-ch
		}()
	}
	group.Wait()
	reply, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "200000", reply)
}

func TestRedisCluster_IncrByFloat(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	s, err := redis.IncrByFloat("godis", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, s)

	s, err = redis.IncrByFloat("godis", 1.62)
	assert.Nil(t, err)
	assert.Equal(t, 3.12, s)
}

func TestRedisCluster_Keys(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
}

func TestRedisCluster_Lindex(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	s, err := redis.Lpush("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), s)

	el, err := redis.Lindex("godis", 0)
	assert.Nil(t, err)
	assert.Equal(t, "1", el)

	el, err = redis.Lindex("godis", -1)
	assert.Nil(t, err)
	assert.Equal(t, "3", el)

	el, err = redis.Lindex("godis", 3)
	assert.Nil(t, err)
	assert.Equal(t, "", el)
}

func TestRedisCluster_List(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	s, err := redis.Lpush("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), s)

	s, err = redis.Linsert("godis", ListoptionBefore, "2", "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s)

	s, err = redis.Linsert("godis", ListoptionAfter, "3", "3.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)

	s, err = redis.Linsert("godis", ListoptionBefore, "2", "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(6), s)

	s, err = redis.Linsert("godis", ListoptionBefore, "1.5", "1.4")
	assert.Nil(t, err)
	assert.Equal(t, int64(7), s)

	arr, err := redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "1.4", "1.5", "1.5", "2", "3", "3.5"}, arr)

	llen, err := redis.Llen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(7), llen)

	lpop, err := redis.Lpop("godis")
	assert.Nil(t, err)
	assert.Equal(t, "1", lpop)

	rpop, err := redis.Rpop("godis")
	assert.Nil(t, err)
	assert.Equal(t, "3.5", rpop)

	s, err = redis.Rpush("godis", "4")
	assert.Nil(t, err)
	assert.Equal(t, int64(6), s)

	rpoplpush, err := redis.Rpoplpush("godis", "0.5")
	assert.NotNil(t, err)
	assert.Equal(t, "", rpoplpush)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "1.5", "1.5", "2", "3", "4"}, arr)

	llen, err = redis.Llen("0.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), llen)

	s, err = redis.Lrem("godis", 0, "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "3", "4"}, arr)

	lset, err := redis.Lset("godis", 2, "2.0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", lset)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "2.0", "4"}, arr)

	lset, err = redis.Lset("godis", 4, "2.0")
	assert.NotNil(t, err)

	ltrim, err := redis.Ltrim("godis", 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, "OK", ltrim)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"2", "2.0"}, arr)
}

func TestRedisCluster_Move(t *testing.T) {
}

func TestRedisCluster_Persist(t *testing.T) {
}

func TestRedisCluster_Pexpire(t *testing.T) {
}

func TestRedisCluster_PexpireAt(t *testing.T) {
}

func TestRedisCluster_Pfadd(t *testing.T) {
}

func TestRedisCluster_Pfcount(t *testing.T) {
}

func TestRedisCluster_Pfmerge(t *testing.T) {
}

func TestRedisCluster_Psetex(t *testing.T) {
}

func TestRedisCluster_Psubscribe(t *testing.T) {
}

func TestRedisCluster_Pttl(t *testing.T) {
}

func TestRedisCluster_Publish(t *testing.T) {
}

func TestRedisCluster_RandomKey(t *testing.T) {
}

func TestRedisCluster_Rename(t *testing.T) {
}

func TestRedisCluster_Renamenx(t *testing.T) {
}

func TestRedisCluster_Scan(t *testing.T) {
}

func TestRedisCluster_ScriptLoad(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	sha, err := redis.ScriptLoad("godis1", `return redis.call("get",KEYS[1])`)
	assert.Nil(t, err)

	bools, err := redis.ScriptExists("godis1", sha)
	assert.Nil(t, err)
	assert.Equal(t, []bool{true}, bools)

	s, err := redis.Evalsha(sha, 1, "godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedisCluster_Sdiff(t *testing.T) {
}

func TestRedisCluster_Sdiffstore(t *testing.T) {
}

func TestRedisCluster_Set(t *testing.T) {
}

func TestRedisCluster_SetWithParams(t *testing.T) {
}

func TestRedisCluster_SetWithParamsAndTime(t *testing.T) {
}

func TestRedisCluster_Setbit(t *testing.T) {
}

func TestRedisCluster_SetbitWithBool(t *testing.T) {
}

func TestRedisCluster_Setex(t *testing.T) {
}

func TestRedisCluster_Setnx(t *testing.T) {
}

func TestRedisCluster_Setrange(t *testing.T) {
}

func TestRedisCluster_Smembers(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	c, err := redis.Sadd("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.Scard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	b, err := redis.Sismember("godis", "1")
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	arr, err := redis.Smembers("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, arr)

	s, err := redis.Srandmember("godis")
	assert.Nil(t, err)
	assert.Contains(t, []string{"1", "2", "3"}, s)

	arr, err = redis.SrandmemberBatch("godis", 2)
	assert.Nil(t, err)
	assert.Len(t, arr, 2)

	c, err = redis.Srem("godis", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	spop, err := redis.Spop("godis")
	assert.Nil(t, err)
	assert.Contains(t, []string{"2", "3"}, spop)

	arr, err = redis.SpopBatch("godis", 2)
	assert.Nil(t, err)
	assert.Subset(t, []string{"2", "3"}, arr)
}

func TestRedisCluster_Smove(t *testing.T) {
}

func TestRedisCluster_Sort(t *testing.T) {
}

func TestRedisCluster_SortMulti(t *testing.T) {
}

func TestRedisCluster_Sscan(t *testing.T) {
}

func TestRedisCluster_Strlen(t *testing.T) {
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
}

func TestRedisCluster_Sunion(t *testing.T) {
}

func TestRedisCluster_Sunionstore(t *testing.T) {
}

func TestRedisCluster_Ttl(t *testing.T) {
}

func TestRedisCluster_Type(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Type("godis")
	assert.Nil(t, err)
	assert.Equal(t, "string", s)
}

func TestRedisCluster_Unwatch(t *testing.T) {
}

func TestRedisCluster_Watch(t *testing.T) {
}

func TestRedisCluster_Zadd(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Del("godis")
	zaddParam := NewZAddParams().NX()
	c, err := redis.Zadd("godis", 1, "a", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.ZaddByMap("godis", map[string]float64{"b": 2, "c": 3, "d": 4, "e": 5}, zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	c, err = redis.Zcard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(5), c)

	//zcount include the boundary
	c, err = redis.Zcount("godis", "2", "5")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	f, err := redis.Zincrby("godis", 1.5, "e", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, 6.5, f)

	c, err = redis.Zrem("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	arr, err := redis.Zrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"b", "c", "d", "e"}, arr)

	tuples, err := redis.ZrangeByScoreWithScores("godis", "2", "6.5")
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: []byte("b"), score: 2},
		{element: []byte("c"), score: 3},
		{element: []byte("d"), score: 4},
		{element: []byte("e"), score: 6.5},
	}, tuples)
}

func TestRedisCluster_Zscan(t *testing.T) {
}

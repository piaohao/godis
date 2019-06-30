package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

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

func clearKeys(cluster *RedisCluster) {
	for _, k := range []string{"godis", "godis1", "godis2", "godis3", "godis4", "godis5"} {
		cluster.Del(k)
	}
}

func TestRedisCluster_Append(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	clearKeys(cluster)
	count, err := cluster.Append("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

func TestRedisCluster_Bitcount(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	cluster.Set("godis", "good")
	count, err := cluster.BitCount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(20), count)
}

func TestRedisCluster_BitcountRange(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.BitCountRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s)
}

func TestRedisCluster_Bitfield(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	_, err := redis.BitField("godis", "INCRBY")
	assert.NotNil(t, err)
}

func TestRedisCluster_Bitop(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.Del("bit-1")
	redis.Del("bit-2")
	redis.Del("and-result")

	b, e := redis.SetBit("bit-1", 0, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.SetBit("bit-1", 3, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.SetBit("bit-2", 0, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.SetBit("bit-2", 1, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.SetBit("bit-2", 3, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	i, e := redis.BitOp(BitOpAnd, "c", "bit-1", "bit-2")
	assert.NotNil(t, e)
	assert.Equal(t, int64(0), i)

	b, e = redis.GetBit("and-result", 0)
	assert.Nil(t, e)
	assert.Equal(t, false, b)
}

func TestRedisCluster_Bitpos(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "\x00\xff\xf0")
	s, err := redis.BitPos("godis", true, BitPosParams{params: [][]byte{IntToByteArray(0)}})
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s)
}

func TestRedisCluster_Blpop(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")

	_, e := redis.BLPop("job", "command", "request", "0")
	assert.NotNil(t, e)

}

func TestRedisCluster_BlpopTimout(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	go func() {
		_, e := redis.BLPopTimeout(5, "command", "update system...")
		assert.NotNil(t, e)
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedisCluster_Brpop(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")

	_, e := redis.BRPop("job", "command", "request", "0")
	assert.NotNil(t, e)

}

func TestRedisCluster_BrpopTimout(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	go func() {
		redis.BRPopTimeout(5, "command", "update system...")
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedisCluster_Brpoplpush(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	go func() {
		redis.BRPopLPush("command", "update system...", 5)
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedisCluster_Decr(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.Decr("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)

	s, err = redis.Decr("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-2), s)
}

func TestRedisCluster_DecrBy(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.DecrBy("godis", 10)
	assert.Nil(t, err)
	assert.Equal(t, int64(-10), s)

	s, err = redis.DecrBy("godis", -10)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedisCluster_Del(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	_, _ = cluster.Set("godis", "good")
	count, err := cluster.Del("godis", "godis1", "godis2", "godis3", "godis4", "godis5")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), count)
	count, err = cluster.Del("godis")
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
	clearKeys(redis)
	c, err := redis.GeoAdd("godis", 121, 37, "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.GeoAddByMap("godis", map[string]GeoCoordinate{
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

	arr, err := redis.GeoHash("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"www43rts870"}, arr)

	_, err = redis.GeoPos("godis", "b")
	assert.Nil(t, err)

	d, err := redis.GeoDist("godis", "a", "b", GeoUnitKm)
	assert.Nil(t, err)
	assert.Equal(t, 88.8291, d)

	resp, err := redis.GeoRadius("godis", 121, 37, 500, GeoUnitKm,
		NewGeoRadiusParam().WithCoord().WithDist())
	assert.Nil(t, err)
	t.Log(resp)

	resp, err = redis.GeoRadiusByMember("godis", "a", 500, GeoUnitKm,
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
	clearKeys(redis)
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
	s, err := redis.GetBit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, s)
}

func TestRedisCluster_Getrange(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.GetRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
	s, err = redis.GetRange("godis", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, "go", s)
}

func TestRedisCluster_Hdel(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HDel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	s, err = redis.HDel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedisCluster_Hexists(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HExists("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, true, s)

	s, err = redis.HExists("godis", "b")
	assert.Nil(t, err)
	assert.Equal(t, false, s)
}

func TestRedisCluster_Hget(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)
}

func TestRedisCluster_HgetAll(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HGetAll("godis")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, s)
}

func TestRedisCluster_HincrBy(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HIncrBy("godis", "a", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	s, err = redis.HIncrBy("godis", "b", 5)
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)
}

func TestRedisCluster_HincrByFloat(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")
	ret, err := redis.HIncrByFloat("godis", "a", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 2.5, ret)

	ret, err = redis.HIncrByFloat("godis", "b", 5.0987)
	assert.Nil(t, err)
	assert.Equal(t, 5.0987, ret)
}

func TestRedisCluster_Hkeys(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HKeys("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a"}, s)
}

func TestRedisCluster_Hlen(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedisCluster_Hmget(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HMGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
}

func TestRedisCluster_Hmset(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HMSet("godis", map[string]string{"b": "2", "c": "3"})
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, _ := redis.HLen("godis")
	assert.Equal(t, int64(3), c)
}

func TestRedisCluster_Hscan(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
}

func TestRedisCluster_Hset(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	ret, err := redis.HSet("godis", "a", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), ret)

	s, err := redis.HLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedisCluster_Hsetnx(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	ret, err := redis.HSetNx("godis", "a", "2")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ret)

	s, err = redis.HGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)
}

func TestRedisCluster_Hvals(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.HSet("godis", "a", "1")

	s, err := redis.HVals("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
}

func TestRedisCluster_Incr(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	clearKeys(cluster)
	for i := 0; i < 10000; i++ {
		cluster.Incr("godis")
	}
	reply, _ := cluster.Get("godis")
	assert.Equal(t, "10000", reply)
}

func TestRedisCluster_IncrBy(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
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
	clearKeys(redis)
	s, err := redis.IncrByFloat("godis", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, s)

	s, err = redis.IncrByFloat("godis", 1.62)
	assert.Nil(t, err)
	assert.Equal(t, 3.12, s)
}

func TestRedisCluster_Keys(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
}

func TestRedisCluster_Lindex(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	s, err := redis.LPush("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), s)

	el, err := redis.LIndex("godis", 0)
	assert.Nil(t, err)
	assert.Equal(t, "1", el)

	el, err = redis.LIndex("godis", -1)
	assert.Nil(t, err)
	assert.Equal(t, "3", el)

	el, err = redis.LIndex("godis", 3)
	assert.Nil(t, err)
	assert.Equal(t, "", el)
}

func TestRedisCluster_List(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	s, err := redis.LPush("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), s)

	s, err = redis.LInsert("godis", ListOptionBefore, "2", "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s)

	s, err = redis.LInsert("godis", ListOptionAfter, "3", "3.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)

	s, err = redis.LInsert("godis", ListOptionBefore, "2", "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(6), s)

	s, err = redis.LInsert("godis", ListOptionBefore, "1.5", "1.4")
	assert.Nil(t, err)
	assert.Equal(t, int64(7), s)

	arr, err := redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "1.4", "1.5", "1.5", "2", "3", "3.5"}, arr)

	llen, err := redis.LLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(7), llen)

	lpop, err := redis.LPop("godis")
	assert.Nil(t, err)
	assert.Equal(t, "1", lpop)

	rpop, err := redis.RPop("godis")
	assert.Nil(t, err)
	assert.Equal(t, "3.5", rpop)

	s, err = redis.RPush("godis", "4")
	assert.Nil(t, err)
	assert.Equal(t, int64(6), s)

	rpoplpush, err := redis.RPopLPush("godis", "0.5")
	assert.NotNil(t, err)
	assert.Equal(t, "", rpoplpush)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "1.5", "1.5", "2", "3", "4"}, arr)

	llen, err = redis.LLen("0.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), llen)

	s, err = redis.LRem("godis", 0, "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "3", "4"}, arr)

	lset, err := redis.LSet("godis", 2, "2.0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", lset)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "2.0", "4"}, arr)

	lset, err = redis.LSet("godis", 4, "2.0")
	assert.NotNil(t, err)
	assert.Equal(t, "", lset)

	ltrim, err := redis.LTrim("godis", 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, "OK", ltrim)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"2", "2.0"}, arr)
}

func TestRedisCluster_Persist(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.Expire("godis", 100)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	s, err = redis.Persist("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	c, err := redis.TTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), c)
}

func TestRedisCluster_Pexpire(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.PExpire("godis", 1000)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedisCluster_PexpireAt(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.PExpireAt("godis", time.Now().Add(1*time.Second).UnixNano()/1e6)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedisCluster_Pfadd(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	c, err := redis.PfAdd("godis", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.PfCount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	c, err = redis.PfAdd("godis1", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.PfCount("godis1")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	s, err := redis.PfMerge("godis3", "godis", "godis1")
	assert.NotNil(t, err)
	assert.Equal(t, "", s)

	c, err = redis.PfCount("godis3")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), c)
}

func TestRedisCluster_Psetex(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.PSetEx("godis", 1000, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", get)
}

func TestRedisCluster_Psubscribe(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	pubsub := &RedisPubSub{
		OnMessage: func(channel, message string) {
			t.Logf("receive message ,channel:%s,message:%s", channel, message)
		},
		OnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive subscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnUnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive unsubscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnPMessage: func(pattern string, channel, message string) {
			t.Logf("receive pmessage ,pattern:%s,channel:%s,message:%s", pattern, channel, message)
		},
		OnPSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive psubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPUnSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive punsubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPong: func(channel string) {
			t.Logf("receive pong ,channel:%s", channel)
		},
	}
	go func() {
		redis.PSubscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
}

func TestRedisCluster_Pttl(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.PTTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)
}

func TestRedisCluster_Rename(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, e := redis.Rename("godis", "godis1")
	assert.NotNil(t, e)
	assert.Equal(t, "", s)

	redis.Set("godis", "good")
	c, e := redis.RenameNx("godis", "godis1")
	assert.NotNil(t, e)
	assert.Equal(t, int64(0), c)
}

func TestRedisCluster_Scan(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	for i := 0; i < 1000; i++ {
		redis.Set(fmt.Sprintf("{godis}%d", i), fmt.Sprintf("godis%d", i))
	}

	params := &ScanParams{
		params: map[*keyword][]byte{
			keywordMatch: []byte("{godis}*"),
			keywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.Scan(cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	assert.Equal(t, 1000, total)
}

func TestRedisCluster_ScriptLoad(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	sha, err := redis.ScriptLoad("godis1", `return redis.call("get",KEYS[1])`)
	assert.Nil(t, err)

	bools, err := redis.ScriptExists("godis1", sha)
	assert.Nil(t, err)
	assert.Equal(t, []bool{true}, bools)

	s, err := redis.EvalSha(sha, 1, "godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedisCluster_Sdiff(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.SAdd("godis1", "1", "2", "3")
	redis.SAdd("godis2", "2", "3", "4")

	_, e := redis.SDiff("godis1", "godis2")
	assert.NotNil(t, e)

	_, e = redis.SDiffStore("godis3", "godis1", "godis2")
	assert.NotNil(t, e)

	_, e = redis.SInter("godis1", "godis2")
	assert.NotNil(t, e)

	_, e = redis.SInterStore("godis4", "godis1", "godis2")
	assert.NotNil(t, e)

	_, e = redis.SUnion("godis1", "godis2")
	assert.NotNil(t, e)

	_, e = redis.SUnionStore("godis5", "godis1", "godis2")
	assert.NotNil(t, e)
}

func TestRedisCluster_Set(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	ret, err := redis.Set("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)
}

func TestRedisCluster_SetWithParams(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.SetWithParams("godis", "good", "xx")
	assert.Nil(t, err)
	assert.Equal(t, "", s)

	redis.Set("godis", "good")
	s, err = redis.SetWithParams("godis", "good1", "xx")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good1", get)
}

func TestRedisCluster_SetWithParamsAndTime(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)
	s, err = redis.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.Nil(t, err)
	assert.Equal(t, "", s)
}

func TestRedisCluster_Setbit(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "a")
	c, err := redis.GetBit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, c)

	c, err = redis.SetBit("godis", 6, "1")
	assert.Nil(t, err)
	assert.Equal(t, false, c)

	c, err = redis.SetBitWithBool("godis", 7, false)
	assert.Nil(t, err)
	assert.Equal(t, true, c)

	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "b", get)
}

func TestRedisCluster_Setex(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.SetEx("godis", 1, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)
}

func TestRedisCluster_Setnx(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.SetNx("godis", "good1")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedisCluster_Setrange(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	c, err := redis.SetRange("godis", 5, " ok")
	assert.Nil(t, err)
	assert.Equal(t, int64(8), c)
}

func TestRedisCluster_Smembers(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	c, err := redis.SAdd("godis", "1", "2", "3")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.SCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	b, err := redis.SIsMember("godis", "1")
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	arr, err := redis.SMembers("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, arr)

	s, err := redis.SRandMember("godis")
	assert.Nil(t, err)
	assert.Contains(t, []string{"1", "2", "3"}, s)

	arr, err = redis.SRandMemberBatch("godis", 2)
	assert.Nil(t, err)
	assert.Len(t, arr, 2)

	c, err = redis.SRem("godis", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	spop, err := redis.SPop("godis")
	assert.Nil(t, err)
	assert.Contains(t, []string{"2", "3"}, spop)

	arr, err = redis.SPopBatch("godis", 2)
	assert.Nil(t, err)
	assert.Subset(t, []string{"2", "3"}, arr)
}

func TestRedisCluster_Smove(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.SAdd("godis", "1", "2")
	redis.SAdd("godis1", "3", "4")

	_, e := redis.SMove("godis", "godis1", "2")
	assert.NotNil(t, e)
}

func TestRedisCluster_Sort(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	redis.LPush("godis", "3", "2", "1", "4", "6", "5")
	p := NewSortingParams().Desc()
	arr, e := redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"6", "5", "4", "3", "2", "1"}, arr)

	p = NewSortingParams().Asc()
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, arr)

	_, e = redis.SortStore("godis", "godis1", *p)
	assert.NotNil(t, e)
}

func TestRedisCluster_Sscan(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	for i := 0; i < 1000; i++ {
		redis.SAdd("godis", fmt.Sprintf("%d", i))
	}
	c, err := redis.SCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := &ScanParams{
		params: map[*keyword][]byte{
			keywordMatch: []byte("*"),
			keywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.SScan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	assert.Equal(t, 1000, total)
}

func TestRedisCluster_Strlen(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.StrLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s)
}

func TestRedisCluster_Subscribe(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	clearKeys(cluster)
	pubsub := &RedisPubSub{
		OnMessage: func(channel, message string) {
			t.Logf("receive message ,channel:%s,message:%s", channel, message)
		},
		OnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive subscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnUnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive unsubscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnPMessage: func(pattern string, channel, message string) {
			t.Logf("receive pmessage ,pattern:%s,channel:%s,message:%s", pattern, channel, message)
		},
		OnPSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive psubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPUnSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive punsubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPong: func(channel string) {
			t.Logf("receive pong ,channel:%s", channel)
		},
	}
	go func() {
		cluster.Subscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	cluster.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
}

func TestRedisCluster_Substr(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	redis.Set("godis", "good")
	s, err := redis.SubStr("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedisCluster_Ttl(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	s, err := redis.TTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-2), s)
}

func TestRedisCluster_Type(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	redis.Set("godis", "good")
	s, err := redis.Type("godis")
	assert.Nil(t, err)
	assert.Equal(t, "string", s)
}

func TestRedisCluster_Zadd(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)
	zaddParam := NewZAddParams().NX()
	c, err := redis.ZAdd("godis", 1, "a", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.ZAddByMap("godis", map[string]float64{"b": 2, "c": 3, "d": 4, "e": 5}, zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	c, err = redis.ZCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(5), c)

	//zcount include the boundary
	c, err = redis.ZCount("godis", "2", "5")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	f, err := redis.ZIncrBy("godis", 1.5, "e", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, 6.5, f)

	c, err = redis.ZRem("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	arr, err := redis.ZRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"b", "c", "d", "e"}, arr)

	tuples, err := redis.ZRangeByScoreWithScores("godis", "2", "6.5")
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: []byte("b"), score: 2},
		{element: []byte("c"), score: 3},
		{element: []byte("d"), score: 4},
		{element: []byte("e"), score: 6.5},
	}, tuples)
}

func TestRedisCluster_Zscan(t *testing.T) {
	redis := NewRedisCluster(clusterOption)
	clearKeys(redis)

	for i := 0; i < 1000; i++ {
		redis.ZAdd("godis", float64(i), fmt.Sprintf("a%d", i))
	}
	c, err := redis.ZCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := &ScanParams{
		params: map[*keyword][]byte{
			keywordMatch: []byte("*"),
			keywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.ZScan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	assert.Equal(t, 2000, total)
}

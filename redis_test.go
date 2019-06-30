package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestRedis_Append(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.Append("godis", "very")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	c, err = redis.Append("godis", " good")
	assert.Nil(t, err)
	assert.Equal(t, int64(9), c)
}

func TestRedis_Asking(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.Asking()
	assert.NotNil(t, err)

	redis1 := NewRedis(option1)
	defer redis1.Close()
	s, err := redis1.Asking()
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)
}

func TestRedis_Bitcount(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Bitcount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s)
}

func TestRedis_BitcountRange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.BitcountRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s)
}

func TestRedis_Bitfield(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.Bitfield("godis", "INCRBY")
	assert.NotNil(t, err)
}

func TestRedis_Bitpos(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "\x00\xff\xf0")
	s, err := redis.Bitpos("godis", true, BitPosParams{params: [][]byte{IntToByteArray(0)}})
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s)
}

func TestRedis_Decr(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Decr("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)

	s, err = redis.Decr("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-2), s)
}

func TestRedis_DecrBy(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.DecrBy("godis", 10)
	assert.Nil(t, err)
	assert.Equal(t, int64(-10), s)

	s, err = redis.DecrBy("godis", -10)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedis_Echo(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Echo("godis")
	assert.Nil(t, err)
	assert.Equal(t, "godis", s)
}

func TestRedis_Expire(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Expire("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedis_ExpireAt(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.ExpireAt("godis", time.Now().Add(1*time.Second).Unix())
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedis_Geo(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_Get(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedis_GetSet(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.GetSet("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "", s)
	s, err = redis.GetSet("godis", "good1")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedis_Getbit(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Getbit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, s)
}

func TestRedis_Getrange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Getrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
	s, err = redis.Getrange("godis", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, "go", s)
}

func TestRedis_Hdel(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hdel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	s, err = redis.Hdel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedis_Hexists(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hexists("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, true, s)

	s, err = redis.Hexists("godis", "b")
	assert.Nil(t, err)
	assert.Equal(t, false, s)
}

func TestRedis_Hget(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)
}

func TestRedis_HgetAll(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.HgetAll("godis")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, s)
}

func TestRedis_HincrBy(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.HincrBy("godis", "a", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	s, err = redis.HincrBy("godis", "b", 5)
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)
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
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hkeys("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a"}, s)
}

func TestRedis_Hlen(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedis_Hmget(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hmget("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
}

func TestRedis_Hmset(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hmset("godis", map[string]string{"b": "2", "c": "3"})
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, _ := redis.Hlen("godis")
	assert.Equal(t, int64(3), c)
}

func TestRedis_Hscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.Hset("godis", fmt.Sprintf("a%d", i), fmt.Sprintf("%d", i))
	}
	c, err := redis.Hlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := &ScanParams{
		params: map[*keyword][]byte{
			KeywordMatch: []byte("a*"),
			KeywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.Hscan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	//total contains key and value
	assert.Equal(t, 2000, total)
}

func TestRedis_Hset(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Hset("godis", "a", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), ret)

	s, err := redis.Hlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
}

func TestRedis_Hsetnx(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_Hvals(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Hset("godis", "a", "1")

	s, err := redis.Hvals("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)
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
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.IncrByFloat("godis", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, s)

	s, err = redis.IncrByFloat("godis", 1.62)
	assert.Nil(t, err)
	assert.Equal(t, 3.12, s)
}

func TestRedis_Lindex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_List(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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
	assert.Nil(t, err)
	assert.Equal(t, "4", rpoplpush)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "1.5", "1.5", "2", "3"}, arr)

	llen, err = redis.Llen("0.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), llen)

	s, err = redis.Lrem("godis", 0, "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "3"}, arr)

	lset, err := redis.Lset("godis", 2, "2.0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", lset)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "2.0"}, arr)

	lset, err = redis.Lset("godis", 4, "2.0")
	assert.NotNil(t, err)

	ltrim, err := redis.Ltrim("godis", 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, "OK", ltrim)

	arr, err = redis.Lrange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"2", "2.0"}, arr)
}

func TestRedis_Move(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Move("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), ret)

	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)

	redis.Select(1)

	get, err = redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", get)
}

func TestRedis_Persist(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Expire("godis", 100)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	s, err = redis.Persist("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	c, err := redis.Ttl("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), c)
}

func TestRedis_Pexpire(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Pexpire("godis", 1000)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedis_PexpireAt(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.PexpireAt("godis", time.Now().Add(1*time.Second).UnixNano()/1e6)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedis_Pfadd(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.Pfadd("godis", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.Pfcount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	c, err = redis.Pfadd("godis1", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.Pfcount("godis1")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	s, err := redis.Pfmerge("godis3", "godis", "godis1")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, err = redis.Pfcount("godis3")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)
}

func TestRedis_Psetex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Psetex("godis", 1000, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", get)
}

func TestRedis_Pttl(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Pttl("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)
}

func TestRedis_PubsubChannels(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.PubsubChannels("godis")
	assert.Nil(t, err)

}

func TestRedis_Readonly(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.Readonly()
	assert.NotNil(t, err)
}

func TestRedis_Send(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	err := redis.Send(CmdGet, []byte("godis"))
	assert.Nil(t, err)
	s, err := ToStringReply(redis.Receive())
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedis_SendByStr(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	err := redis.SendByStr("get", []byte("godis"))
	assert.Nil(t, err)
	s, err := ToStringReply(redis.Receive())
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
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
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_SetWithParamsAndTime(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)
	s, err = redis.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.Nil(t, err)
	assert.Equal(t, "", s)
}

func TestRedis_Setbit(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "a")
	c, err := redis.Getbit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, c)

	c, err = redis.Setbit("godis", 6, "1")
	assert.Nil(t, err)
	assert.Equal(t, false, c)

	c, err = redis.SetbitWithBool("godis", 7, false)
	assert.Nil(t, err)
	assert.Equal(t, true, c)

	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "b", get)
}

func TestRedis_Setex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Setex("godis", 1, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)
}

func TestRedis_Setnx(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Setnx("godis", "good1")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)
}

func TestRedis_Setrange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.Setrange("godis", 5, " ok")
	assert.Nil(t, err)
	assert.Equal(t, int64(8), c)
}

func TestRedis_Smembers(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_Sscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.Sadd("godis", fmt.Sprintf("%d", i))
	}
	c, err := redis.Scard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := &ScanParams{
		params: map[*keyword][]byte{
			KeywordMatch: []byte("*"),
			KeywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.Sscan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	assert.Equal(t, 1000, total)
}

func TestRedis_Strlen(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Strlen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s)
}

func TestRedis_Substr(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Substr("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
}

func TestRedis_Ttl(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Ttl("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)
}

func TestRedis_Type(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Type("godis")
	assert.Nil(t, err)
	assert.Equal(t, "string", s)
}

func TestRedis_Zadd(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

func TestRedis_Zscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.Zadd("godis", float64(i), fmt.Sprintf("a%d", i))
	}
	c, err := redis.Zcard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := &ScanParams{
		params: map[*keyword][]byte{
			KeywordMatch: []byte("*"),
			KeywordCount: IntToByteArray(10),
		},
	}
	cursor := "0"
	total := 0
	for {
		result, err := redis.Zscan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	assert.Equal(t, 2000, total)
}

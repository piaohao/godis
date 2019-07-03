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

	m, _ := redis.Multi()
	_, err = redis.Append("godis", "good")
	assert.NotNil(t, err)
	m.Discard()
	redis.client.connection.host = "localhost1"
	redis.Close()
	_, err = redis.Append("godis", "good")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Asking()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Asking()
	assert.NotNil(t, err)
}

func TestRedis_Bitcount(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.BitCount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BitCount("godis")
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BitCount("godis")
	assert.NotNil(t, err)
}

func TestRedis_BitcountRange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.BitCountRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BitCountRange("godis", 0, -1)
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BitCountRange("godis", 0, -1)
	assert.NotNil(t, err)
}

func TestRedis_Bitfield(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.BitField("godis", "INCRBY")
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BitField("godis", "INCRBY")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BitField("godis", "INCRBY")
	assert.NotNil(t, err)
}

func TestRedis_Bitpos(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "\x00\xff\xf0")
	s, err := redis.BitPos("godis", true, &BitPosParams{params: [][]byte{IntToByteArr(0)}})
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BitPos("godis", true, &BitPosParams{params: [][]byte{IntToByteArr(0)}})
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BitPos("godis", true, &BitPosParams{params: [][]byte{IntToByteArr(0)}})
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Decr("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Decr("godis")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.DecrBy("godis", 10)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.DecrBy("godis", 10)
	assert.NotNil(t, err)
}

func TestRedis_Echo(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Echo("godis")
	assert.Nil(t, err)
	assert.Equal(t, "godis", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Echo("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Echo("godis")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Expire("godis", 1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Expire("godis", 1)
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.ExpireAt("godis", time.Now().Add(1*time.Second).Unix())
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ExpireAt("godis", time.Now().Add(1*time.Second).Unix())
	assert.NotNil(t, err)
}

func TestRedis_Geo(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.GeoAdd("godis", 121, 37, "a")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.GeoAdd("godis", 121, 37, "a")
	assert.NotNil(t, err)
}

func TestRedis_Geo2(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	resp, err := redis.GeoRadius("godis", 121, 37, 500, GeoUnitKm,
		NewGeoRadiusParam().WithCoord().WithDist().SortAscending())
	assert.Nil(t, err)
	arr := make([]string, 0)
	for _, re := range resp {
		arr = append(arr, string(re.member))
	}
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, arr)

	resp, err = redis.GeoRadius("godis", 121, 37, 500, GeoUnitKm,
		NewGeoRadiusParam().WithCoord().WithDist().SortDescending().Count(3))
	assert.Nil(t, err)
	arr = make([]string, 0)
	for _, re := range resp {
		arr = append(arr, string(re.member))
	}
	assert.Equal(t, []string{"e", "d", "c"}, arr)
}

func TestRedis_Get(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Get("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Get("godis")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.GetSet("godis", "good")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.GetSet("godis", "good")
	assert.NotNil(t, err)
}

func TestRedis_Getbit(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.GetBit("godis", 1)
	assert.Nil(t, err)
	assert.Equal(t, true, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.GetBit("godis", 1)
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.GetBit("godis", 1)
	assert.NotNil(t, err)
}

func TestRedis_Getrange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.GetRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)
	s, err = redis.GetRange("godis", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, "go", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.GetRange("godis", 0, -1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.GetRange("godis", 0, -1)
	assert.NotNil(t, err)
}

func TestRedis_Hdel(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HDel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	s, err = redis.HDel("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HDel("godis", "a")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HDel("godis", "a")
	assert.NotNil(t, err)
}

func TestRedis_Hexists(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HExists("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, true, s)

	s, err = redis.HExists("godis", "b")
	assert.Nil(t, err)
	assert.Equal(t, false, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HExists("godis", "a")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HExists("godis", "a")
	assert.NotNil(t, err)
}

func TestRedis_Hget(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, "1", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HGet("godis", "a")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HGet("godis", "a")
	assert.NotNil(t, err)
}

func TestRedis_HgetAll(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HGetAll("godis")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HGetAll("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HGetAll("godis")
	assert.NotNil(t, err)
}

func TestRedis_HincrBy(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HIncrBy("godis", "a", 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	s, err = redis.HIncrBy("godis", "b", 5)
	assert.Nil(t, err)
	assert.Equal(t, int64(5), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HIncrBy("godis", "a", 1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HIncrBy("godis", "a", 1)
	assert.NotNil(t, err)
}

func TestRedis_HincrByFloat(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")
	ret, err := redis.HIncrByFloat("godis", "a", 1.5)
	assert.Nil(t, err)
	assert.Equal(t, 2.5, ret)

	ret, err = redis.HIncrByFloat("godis", "b", 5.0987)
	assert.Nil(t, err)
	assert.Equal(t, 5.0987, ret)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HIncrByFloat("godis", "a", 1.5)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HIncrByFloat("godis", "a", 1.5)
	assert.NotNil(t, err)
}

func TestRedis_Hkeys(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HKeys("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"a"}, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HKeys("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HKeys("godis")
	assert.NotNil(t, err)
}

func TestRedis_Hlen(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HLen("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HLen("godis")
	assert.NotNil(t, err)
}

func TestRedis_Hmget(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HMGet("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HMGet("godis", "a")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HMGet("godis", "a")
	assert.NotNil(t, err)
}

func TestRedis_Hmset(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HMSet("godis", map[string]string{"b": "2", "c": "3"})
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, _ := redis.HLen("godis")
	assert.Equal(t, int64(3), c)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HMSet("godis", map[string]string{"b": "2", "c": "3"})
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HMSet("godis", map[string]string{"b": "2", "c": "3"})
	assert.NotNil(t, err)
}

func TestRedis_Hscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.HSet("godis", fmt.Sprintf("a%d", i), fmt.Sprintf("%d", i))
	}
	c, err := redis.HLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := NewScanParams().Match("a*").Count(10)
	cursor := "0"
	total := 0
	for {
		result, err := redis.HScan("godis", cursor, params)
		assert.Nil(t, err)
		total += len(result.Results)
		cursor = result.Cursor
		if result.Cursor == "0" {
			break
		}
	}
	//total contains key and value
	assert.Equal(t, 2000, total)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HScan("godis", cursor, params)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HScan("godis", cursor, params)
	assert.NotNil(t, err)
}

func TestRedis_Hset(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.HSet("godis", "a", "1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), ret)

	s, err := redis.HLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HSet("godis", "a", "1")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HSet("godis", "a", "1")
	assert.NotNil(t, err)
}

func TestRedis_Hsetnx(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HSetNx("godis", "a", "2")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HSetNx("godis", "a", "2")
	assert.NotNil(t, err)
}

func TestRedis_Hvals(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.HSet("godis", "a", "1")

	s, err := redis.HVals("godis")
	assert.Nil(t, err)
	assert.Equal(t, []string{"1"}, s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.HVals("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.HVals("godis")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Incr("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Incr("godis")
	assert.NotNil(t, err)
}

func TestRedis_IncrBy(t *testing.T) {
	flushAll()
	pool := NewPool(nil, option)
	var group sync.WaitGroup
	ch := make(chan bool, 2)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.IncrBy("godis", 2)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.IncrBy("godis", 2)
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.IncrByFloat("godis", 1.62)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.IncrByFloat("godis", 1.62)
	assert.NotNil(t, err)
}

func TestRedis_Lindex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.LIndex("godis", 0)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.LIndex("godis", 0)
	assert.NotNil(t, err)
}

func TestRedis_List0(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	arr := make([]string, 0)
	for i := 0; i < 100000; i++ {
		arr = append(arr, "1")
	}
	l, e := redis.LPush("godis", arr...)
	assert.Nil(t, e)
	assert.Equal(t, int64(100000), l)
}

func TestRedis_List(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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
	assert.Nil(t, err)
	assert.Equal(t, "4", rpoplpush)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "1.5", "1.5", "2", "3"}, arr)

	llen, err = redis.LLen("0.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), llen)

	s, err = redis.LRem("godis", 0, "1.5")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), s)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "3"}, arr)

	lset, err := redis.LSet("godis", 2, "2.0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", lset)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.4", "2", "2.0"}, arr)

	_, err = redis.LSet("godis", 4, "2.0")
	assert.NotNil(t, err)

	ltrim, err := redis.LTrim("godis", 1, 2)
	assert.Nil(t, err)
	assert.Equal(t, "OK", ltrim)

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"2", "2.0"}, arr)

	redis.LPushX("godis", "1")
	redis.RPushX("godis", "3")

	arr, err = redis.LRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "2.0", "3"}, arr)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.LPush("godis", "1", "2", "3")
	assert.NotNil(t, err)
	_, err = redisBroken.LInsert("godis", ListOptionBefore, "2", "1.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.LLen("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.LPop("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.RPop("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.RPush("godis", "4")
	assert.NotNil(t, err)
	_, err = redisBroken.RPopLPush("godis", "0.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LRem("godis", 0, "1.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LSet("godis", 2, "2.0")
	assert.NotNil(t, err)
	_, err = redisBroken.LTrim("godis", 1, 2)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.LPush("godis", "1", "2", "3")
	assert.NotNil(t, err)
	_, err = redisBroken.LInsert("godis", ListOptionBefore, "2", "1.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.LLen("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.LPop("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.RPop("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.RPush("godis", "4")
	assert.NotNil(t, err)
	_, err = redisBroken.RPopLPush("godis", "0.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LRem("godis", 0, "1.5")
	assert.NotNil(t, err)
	_, err = redisBroken.LSet("godis", 2, "2.0")
	assert.NotNil(t, err)
	_, err = redisBroken.LTrim("godis", 1, 2)
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Move("godis", 1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Move("godis", 1)
	assert.NotNil(t, err)
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
	c, err := redis.TTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), c)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Persist("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Persist("godis")
	assert.NotNil(t, err)
}

func TestRedis_Pexpire(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.PExpire("godis", 1000)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PExpire("godis", 1000)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PExpire("godis", 1000)
	assert.NotNil(t, err)
}

func TestRedis_PexpireAt(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.PExpireAt("godis", time.Now().Add(1*time.Second).UnixNano()/1e6)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), s)
	time.Sleep(2 * time.Second)
	ret, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", ret)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PExpireAt("godis", time.Now().Add(1*time.Second).UnixNano()/1e6)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PExpireAt("godis", time.Now().Add(1*time.Second).UnixNano()/1e6)
	assert.NotNil(t, err)
}

func TestRedis_Pfadd(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	c, err = redis.PfCount("godis3")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PfAdd("godis", "a", "b", "c", "d")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PfAdd("godis", "a", "b", "c", "d")
	assert.NotNil(t, err)
}

func TestRedis_Psetex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.PSetEx("godis", 1000, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", get)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PSetEx("godis", 1000, "good")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PSetEx("godis", 1000, "good")
	assert.NotNil(t, err)
}

func TestRedis_Pttl(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.PTTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PTTL("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PTTL("godis")
	assert.NotNil(t, err)
}

func TestRedis_PubsubChannels(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.PubSubChannels("godis")
	assert.Nil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.PubSubChannels("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.PubSubChannels("godis")
	assert.NotNil(t, err)

}

func TestRedis_Readonly(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.Readonly()
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Readonly()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Readonly()
	assert.NotNil(t, err)
}

func TestRedis_Send(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	err := redis.Send(cmdGet, []byte("godis"))
	assert.Nil(t, err)
	s, err := ToStrReply(redis.Receive())
	assert.Nil(t, err)
	assert.Equal(t, "good", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	err = redisBroken.Send(cmdGet, []byte("godis"))
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	err = redisBroken.Send(cmdGet, []byte("godis"))
	assert.NotNil(t, err)
}

func TestRedis_SendByStr(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	err := redis.SendByStr("get", []byte("godis"))
	assert.Nil(t, err)
	s, err := ToStrReply(redis.Receive())
	assert.Nil(t, err)
	assert.Equal(t, "good", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	err = redisBroken.SendByStr("get", []byte("godis"))
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	err = redisBroken.SendByStr("get", []byte("godis"))
	assert.NotNil(t, err)
}

func TestRedis_Set(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Set("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)

	str := ""
	for i := 0; i < 20000; i++ {
		str += "1"
	}
	redis.Set("godis", str)
	redis.Get("godis")

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Set("godis", "good")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Set("godis", "good")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetWithParams("godis", "good", "xx")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetWithParams("godis", "good", "xx")
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetWithParamsAndTime("godis", "good", "nx", "px", 1500)
	assert.NotNil(t, err)
}

func TestRedis_Setbit(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetBit("godis", 6, "1")
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetBit("godis", 6, "1")
	assert.NotNil(t, err)
}

func TestRedis_Setex(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SetEx("godis", 1, "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	time.Sleep(2 * time.Second)
	get, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetEx("godis", 1, "good")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetEx("godis", 1, "good")
	assert.NotNil(t, err)
}

func TestRedis_Setnx(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SetNx("godis", "good1")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetNx("godis", "good1")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetNx("godis", "good1")
	assert.NotNil(t, err)
}

func TestRedis_Setrange(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.SetRange("godis", 5, " ok")
	assert.Nil(t, err)
	assert.Equal(t, int64(8), c)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SetRange("godis", 5, " ok")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SetRange("godis", 5, " ok")
	assert.NotNil(t, err)
}

func TestRedis_Smembers(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SAdd("godis", "1", "2", "3")
	assert.NotNil(t, err)
	_, err = redisBroken.SCard("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SIsMember("godis", "1")
	assert.NotNil(t, err)
	_, err = redisBroken.SMembers("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SRandMember("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SRandMemberBatch("godis", 2)
	assert.NotNil(t, err)
	_, err = redisBroken.SRem("godis", "1")
	assert.NotNil(t, err)
	_, err = redisBroken.SPopBatch("godis", 2)
	assert.NotNil(t, err)
	_, err = redisBroken.SPop("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SAdd("godis", "1", "2", "3")
	assert.NotNil(t, err)
	_, err = redisBroken.SCard("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SIsMember("godis", "1")
	assert.NotNil(t, err)
	_, err = redisBroken.SMembers("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SRandMember("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.SRandMemberBatch("godis", 2)
	assert.NotNil(t, err)
	_, err = redisBroken.SRem("godis", "1")
	assert.NotNil(t, err)
	_, err = redisBroken.SPopBatch("godis", 2)
	assert.NotNil(t, err)
	_, err = redisBroken.SPop("godis")
	assert.NotNil(t, err)
}

func TestRedis_Sscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.SAdd("godis", fmt.Sprintf("%d", i))
	}
	c, err := redis.SCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := NewScanParams().Match("*").Count(10)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SScan("godis", cursor, params)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SScan("godis", cursor, params)
	assert.NotNil(t, err)
}

func TestRedis_Strlen(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.StrLen("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.StrLen("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.StrLen("godis")
	assert.NotNil(t, err)
}

func TestRedis_Substr(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SubStr("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, "good", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SubStr("godis", 0, -1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SubStr("godis", 0, -1)
	assert.NotNil(t, err)
}

func TestRedis_Ttl(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.TTL("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.TTL("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.TTL("godis")
	assert.NotNil(t, err)
}

func TestRedis_Type(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.Type("godis")
	assert.Nil(t, err)
	assert.Equal(t, "string", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Type("godis")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Type("godis")
	assert.NotNil(t, err)
}

func TestRedis_Zadd(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	f0, err := redis.ZScore("godis", "a")
	assert.Nil(t, err)
	assert.Equal(t, float64(1), f0)

	//zcount include the boundary
	c, err = redis.ZCount("godis", 2, 5)
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

	arr, err = redis.ZRangeByScore("godis", 2, 6.5)
	assert.Nil(t, err)
	assert.Equal(t, []string{"b", "c", "d", "e"}, arr)

	arr, err = redis.ZRevRangeByScore("godis", 6.5, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"e", "d", "c", "b"}, arr)

	arr, err = redis.ZRevRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"e", "d", "c", "b"}, arr)

	tuples, err := redis.ZRangeByScoreWithScores("godis", 2, 6.5)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "b", score: 2},
		{element: "c", score: 3},
		{element: "d", score: 4},
		{element: "e", score: 6.5},
	}, tuples)

	tuples, err = redis.ZRangeWithScores("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "b", score: 2},
		{element: "c", score: 3},
		{element: "d", score: 4},
		{element: "e", score: 6.5},
	}, tuples)

	tuples, err = redis.ZRevRangeByScoreWithScores("godis", 6.5, 2)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "e", score: 6.5},
		{element: "d", score: 4},
		{element: "c", score: 3},
		{element: "b", score: 2},
	}, tuples)

	tuples, err = redis.ZRevRangeWithScores("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "e", score: 6.5},
		{element: "d", score: 4},
		{element: "c", score: 3},
		{element: "b", score: 2},
	}, tuples)

	arr, err = redis.ZRangeByScoreBatch("godis", 2, 6.5, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"b", "c"}, arr)

	tuples, err = redis.ZRangeByScoreWithScoresBatch("godis", 2, 6.5, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "b", score: 2},
		{element: "c", score: 3},
	}, tuples)

	tuples, err = redis.ZRevRangeByScoreWithScoresBatch("godis", 6.5, 2, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []Tuple{
		{element: "e", score: 6.5},
		{element: "d", score: 4},
	}, tuples)

	c, err = redis.ZRemRangeByScore("godis", 2, 2.5)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.ZLexCount("godis", "-", "+")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	arr, err = redis.ZRangeByLex("godis", "-", "+")
	assert.Nil(t, err)
	assert.Equal(t, []string{"c", "d", "e"}, arr)

	arr, err = redis.ZRangeByLexBatch("godis", "-", "+", 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"c", "d"}, arr)

	arr, err = redis.ZRevRangeByLex("godis", "+", "-")
	assert.Nil(t, err)
	assert.Equal(t, []string{"e", "d", "c"}, arr)

	arr, err = redis.ZRevRangeByLexBatch("godis", "+", "-", 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"e", "d"}, arr)

	c, err = redis.ZRemRangeByLex("godis", "[c", "[d")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.ZCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c) //{"e":6.5}

	redis.ZAddByMap("godis", map[string]float64{"b": 2, "c": 3, "d": 4}, zaddParam)

	c, err = redis.ZRank("godis", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.ZRevRank("godis", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.ZRank("godis", "f")
	assert.Nil(t, err)
	assert.Equal(t, int64(-1), c) //f is not in godis

	c, err = redis.ZRemRangeByRank("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), c) //f is not in godis

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.ZAdd("godis", 1, "a", zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZAddByMap("godis", map[string]float64{"b": 2, "c": 3, "d": 4, "e": 5}, zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZCard("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.ZScore("godis", "a")
	assert.NotNil(t, err)
	_, err = redisBroken.ZCount("godis", 2, 5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZIncrBy("godis", 1.5, "e", zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRem("godis", "a")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScore("godis", 2, 6.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScore("godis", 6.5, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreWithScores("godis", 2, 6.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeWithScores("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScoreWithScores("godis", 6.5, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeWithScores("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreBatch("godis", 2, 6.5, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreWithScoresBatch("godis", 2, 6.5, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScoreWithScoresBatch("godis", 6.5, 2, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByScore("godis", 2, 2.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZLexCount("godis", "-", "+")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByLex("godis", "-", "+")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByLexBatch("godis", "-", "+", 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByLex("godis", "+", "-")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByLexBatch("godis", "+", "-", 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByLex("godis", "[c", "[d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRank("godis", "d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRank("godis", "d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByRank("godis", 0, -1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ZAdd("godis", 1, "a", zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZAddByMap("godis", map[string]float64{"b": 2, "c": 3, "d": 4, "e": 5}, zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZCard("godis")
	assert.NotNil(t, err)
	_, err = redisBroken.ZScore("godis", "a")
	assert.NotNil(t, err)
	_, err = redisBroken.ZCount("godis", 2, 5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZIncrBy("godis", 1.5, "e", zaddParam)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRem("godis", "a")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScore("godis", 2, 6.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScore("godis", 6.5, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRange("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreWithScores("godis", 2, 6.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeWithScores("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScoreWithScores("godis", 6.5, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeWithScores("godis", 0, -1)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreBatch("godis", 2, 6.5, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByScoreWithScoresBatch("godis", 2, 6.5, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByScoreWithScoresBatch("godis", 6.5, 2, 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByScore("godis", 2, 2.5)
	assert.NotNil(t, err)
	_, err = redisBroken.ZLexCount("godis", "-", "+")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByLex("godis", "-", "+")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRangeByLexBatch("godis", "-", "+", 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByLex("godis", "+", "-")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRangeByLexBatch("godis", "+", "-", 0, 2)
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByLex("godis", "[c", "[d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRank("godis", "d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRevRank("godis", "d")
	assert.NotNil(t, err)
	_, err = redisBroken.ZRemRangeByRank("godis", 0, -1)
	assert.NotNil(t, err)
}

func TestRedis_Zadd2(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	zaddParam := NewZAddParams().NX()
	c, err := redis.ZAdd("godis", 1, "a", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c) //{"a":1}

	zaddParam = NewZAddParams().XX()
	c, err = redis.ZAdd("godis", 2, "a", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), c) //{"a":2}

	c, err = redis.ZAdd("godis", 3, "b", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), c) //{"a":2}

	zaddParam = NewZAddParams().CH()
	c, err = redis.ZAdd("godis", 3, "c", zaddParam)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	arr, err := redis.ZRange("godis", 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "c"}, arr)
}

func TestRedis_Zscan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.ZAdd("godis", float64(i), fmt.Sprintf("a%d", i))
	}
	c, err := redis.ZCard("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), c)

	params := NewScanParams().Match("*").Count(10)
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

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.ZScan("godis", cursor, params)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ZScan("godis", cursor, params)
	assert.NotNil(t, err)
}

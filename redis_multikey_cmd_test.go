package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis_Keys(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis1", "good")
	arr, e := redis.Keys("*")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"godis", "godis1"}, arr)

	b, e := redis.Exists("godis")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), b)
}

func TestRedis_Del(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.Del("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)
}

func TestRedis_Blpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")

	arr, e := redis.Blpop("job", "command", "request", "0")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
}

func TestRedis_BlpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.BlpopTimout(5, "command", "update system...")
		assert.Nil(t, e)
		assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
	}()
	time.Sleep(1 * time.Second)
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedis_Brpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")

	arr, e := redis.Brpop("job", "command", "request", "0")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
}

func TestRedis_BrpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.BrpopTimout(5, "command", "update system...")
		assert.Nil(t, e)
		assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
	}()
	time.Sleep(1 * time.Second)
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedis_Mset(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()

	s, e := redis.Mset("godis1", "good", "godis2", "good")
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)

	c, e := redis.Msetnx("godis1", "good1")
	assert.Nil(t, e)
	assert.Equal(t, int64(0), c)

	arr, e := redis.Mget("godis", "godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"good", "good", "good"}, arr)
}

func TestRedis_Rename(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, e := redis.Rename("godis", "godis1")
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)

	redis.Set("godis", "good")
	c, e := redis.Renamenx("godis", "godis1")
	assert.Nil(t, e)
	assert.Equal(t, int64(0), c)
}

func TestRedis_Brpoplpush(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.Brpoplpush("command", "update system...", 5)
		assert.Nil(t, e)
		assert.Equal(t, "update system...", arr)
	}()
	time.Sleep(1 * time.Second)
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")
	time.Sleep(1 * time.Second)
}

func TestRedis_Sdiff(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Sadd("godis1", "1", "2", "3")
	redis.Sadd("godis2", "2", "3", "4")

	arr, e := redis.Sdiff("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1"}, arr)

	c, e := redis.Sdiffstore("godis3", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), c)

	arr, e = redis.Smembers("godis3")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1"}, arr)

	arr, e = redis.Sinter("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"2", "3"}, arr)

	c, e = redis.Sinterstore("godis4", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(2), c)

	arr, e = redis.Sunion("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1", "2", "3", "4"}, arr)

	c, e = redis.Sunionstore("godis5", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(4), c)
}

func TestRedis_Smove(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Sadd("godis", "1", "2")
	redis.Sadd("godis1", "3", "4")

	s, e := redis.Smove("godis", "godis1", "2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), s)

	arr, _ := redis.Smembers("godis")
	assert.ElementsMatch(t, []string{"1"}, arr)

	arr, _ = redis.Smembers("godis1")
	assert.ElementsMatch(t, []string{"2", "3", "4"}, arr)

}

func TestRedis_Sort(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Lpush("godis", "3", "2", "1", "4", "6", "5")
	p := NewSortingParams().Desc()
	arr, e := redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"6", "5", "4", "3", "2", "1"}, arr)

	p = NewSortingParams().Asc()
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, arr)

	c, e := redis.SortStore("godis", "godis1", *p)
	assert.Nil(t, e)
	assert.Equal(t, int64(6), c)
}

func TestRedis_Watch(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	s, e := redis.Watch("godis")
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)

	s, e = redis.Unwatch()
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)
}

func TestRedis_Zinterstore(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.ZaddByMap("godis1", map[string]float64{"a": 1, "b": 2, "c": 3})
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.ZaddByMap("godis2", map[string]float64{"a": 1, "b": 2})
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.Zinterstore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.ZinterstoreWithParams("godis3", ZparamsSum, "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.Zunionstore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.ZunionstoreWithParams("godis3", ZparamsMax, "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)
}

func TestRedis_Subscribe(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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
	go func() {
		r := NewRedis(option)
		defer r.Close()
		r.Subscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
}

func TestRedis_Psubscribe(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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
	go func() {
		r := NewRedis(option)
		defer r.Close()
		r.Psubscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
}

func TestRedis_RandomKey(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, e := redis.RandomKey()
	assert.Nil(t, e)
	assert.Equal(t, "godis", s)
}

func TestRedis_Bitop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	b, e := redis.Setbit("bit-1", 0, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.Setbit("bit-1", 3, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.Setbit("bit-2", 0, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.Setbit("bit-2", 1, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = redis.Setbit("bit-2", 3, "1")
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	i, e := redis.Bitop(BitopAnd, "and-result", "bit-1", "bit-2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), i)

	b, e = redis.Getbit("and-result", 0)
	assert.Nil(t, e)
	assert.Equal(t, true, b)
}

func TestRedis_Scan(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	for i := 0; i < 1000; i++ {
		redis.Set(fmt.Sprintf("godis%d", i), fmt.Sprintf("godis%d", i))
	}
	c, err := redis.Keys("godis*")
	assert.Nil(t, err)
	assert.Len(t, c, 1000)

	params := &ScanParams{
		params: map[*keyword][]byte{
			KeywordMatch: []byte("godis*"),
			KeywordCount: IntToByteArray(10),
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

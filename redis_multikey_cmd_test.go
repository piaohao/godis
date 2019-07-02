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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Keys("*")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Keys("*")
	assert.NotNil(t, e)
}

func TestRedis_Del(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	c, e := redis.Del("godis")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), c)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Del("godis")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Del("godis")
	assert.NotNil(t, e)
}

func TestRedis_Blpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")

	arr, e := redis.BLPop("job", "command", "request", "0")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, arr)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.BLPop("job", "command", "request", "0")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.BLPop("job", "command", "request", "0")
	assert.NotNil(t, e)
}

func TestRedis_BlpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.BLPopTimeout(5, "command", "update system...")
		assert.Nil(t, e)
		assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e := redisBroken.BLPopTimeout(5, "command", "update system...")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.BLPopTimeout(5, "command", "update system...")
	assert.NotNil(t, e)
}

func TestRedis_Brpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")

	arr, e := redis.BRPop("job", "command", "request", "0")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, arr)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.BRPop("job", "command", "request", "0")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.BRPop("job", "command", "request", "0")
	assert.NotNil(t, e)
}

func TestRedis_BrpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.BRPopTimeout(5, "command", "update system...")
		assert.Nil(t, e)
		assert.ElementsMatch(t, []string{"command", "update system..."}, arr)
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e := redisBroken.BRPopTimeout(5, "command", "update system...")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.BRPopTimeout(5, "command", "update system...")
	assert.NotNil(t, e)
}

func TestRedis_Mset(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()

	s, e := redis.MSet("godis1", "good", "godis2", "good")
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)

	c, e := redis.MSetNx("godis1", "good1")
	assert.Nil(t, e)
	assert.Equal(t, int64(0), c)

	arr, e := redis.MGet("godis", "godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"good", "good", "good"}, arr)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.MSet("godis1", "good", "godis2", "good")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.MSet("godis1", "good", "godis2", "good")
	assert.NotNil(t, e)
}

func TestRedis_Rename(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, e := redis.Rename("godis", "godis1")
	assert.Nil(t, e)
	assert.Equal(t, "OK", s)

	redis.Set("godis", "good")
	c, e := redis.RenameNx("godis", "godis1")
	assert.Nil(t, e)
	assert.Equal(t, int64(0), c)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Rename("godis", "godis1")
	assert.NotNil(t, e)
	_, e = redisBroken.RenameNx("godis", "godis1")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Rename("godis", "godis1")
	assert.NotNil(t, e)
	_, e = redisBroken.RenameNx("godis", "godis1")
	assert.NotNil(t, e)
}

func TestRedis_Brpoplpush(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		redis := NewRedis(option)
		defer redis.Close()
		arr, e := redis.BRPopLPush("command", "update system...", 5)
		assert.Nil(t, e)
		assert.Equal(t, "update system...", arr)
	}()
	time.Sleep(1 * time.Second)
	redis.LPush("command", "update system...")
	redis.LPush("request", "visit page")
	time.Sleep(1 * time.Second)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e := redisBroken.BRPopLPush("command", "update system...", 5)
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.BRPopLPush("command", "update system...", 5)
	assert.NotNil(t, e)
}

func TestRedis_Sdiff(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.SAdd("godis1", "1", "2", "3")
	redis.SAdd("godis2", "2", "3", "4")

	arr, e := redis.SDiff("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1"}, arr)

	c, e := redis.SDiffStore("godis3", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), c)

	arr, e = redis.SMembers("godis3")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1"}, arr)

	arr, e = redis.SInter("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"2", "3"}, arr)

	c, e = redis.SInterStore("godis4", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(2), c)

	arr, e = redis.SUnion("godis1", "godis2")
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1", "2", "3", "4"}, arr)

	c, e = redis.SUnionStore("godis5", "godis1", "godis2")
	assert.Nil(t, e)
	assert.Equal(t, int64(4), c)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.SDiff("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SDiffStore("godis3", "godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SMembers("godis3")
	assert.NotNil(t, e)
	_, e = redisBroken.SInter("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SInterStore("godis4", "godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SUnion("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SUnionStore("godis5", "godis1", "godis2")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.SDiff("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SDiffStore("godis3", "godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SMembers("godis3")
	assert.NotNil(t, e)
	_, e = redisBroken.SInter("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SInterStore("godis4", "godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SUnion("godis1", "godis2")
	assert.NotNil(t, e)
	_, e = redisBroken.SUnionStore("godis5", "godis1", "godis2")
	assert.NotNil(t, e)
}

func TestRedis_Smove(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.SAdd("godis", "1", "2")
	redis.SAdd("godis1", "3", "4")

	s, e := redis.SMove("godis", "godis1", "2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), s)

	arr, _ := redis.SMembers("godis")
	assert.ElementsMatch(t, []string{"1"}, arr)

	arr, _ = redis.SMembers("godis1")
	assert.ElementsMatch(t, []string{"2", "3", "4"}, arr)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.SMove("godis", "godis1", "2")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.SMove("godis", "godis1", "2")
	assert.NotNil(t, e)

}

func TestRedis_Sort(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.LPush("godis", "3", "2", "1", "4", "6", "5")
	p := NewSortingParams().Desc()
	arr, e := redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"6", "5", "4", "3", "2", "1"}, arr)

	p = NewSortingParams().Desc().Limit(0, 1)
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"6"}, arr)

	p = NewSortingParams().Asc()
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, arr)

	p = NewSortingParams().Alpha()
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, arr)

	p = NewSortingParams().By("*").Get("*")
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"", "", "", "", "", ""}, arr)

	c, e := redis.SortStore("godis", "godis1", *p)
	assert.Nil(t, e)
	assert.Equal(t, int64(6), c)

	p = NewSortingParams().NoSort()
	arr, e = redis.Sort("godis", *p)
	assert.Nil(t, e)
	assert.Equal(t, []string{"3", "2", "1", "4", "6", "5"}, arr)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Sort("godis", *p)
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Sort("godis", *p)
	assert.NotNil(t, e)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Watch("godis")
	assert.NotNil(t, e)
	_, e = redisBroken.Unwatch()
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Watch("godis")
	assert.NotNil(t, e)
	_, e = redisBroken.Unwatch()
	assert.NotNil(t, e)
}

func TestRedis_Zinterstore(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	c, err := redis.ZAddByMap("godis1", map[string]float64{"a": 1, "b": 2, "c": 3})
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.ZAddByMap("godis2", map[string]float64{"a": 1, "b": 2})
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.ZInterStore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	param := newZParams().Aggregate(AggregateSum)
	c, err = redis.ZInterStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	c, err = redis.ZUnionStore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	param = newZParams().Aggregate(AggregateMax)
	c, err = redis.ZUnionStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.ZInterStore("godis3", "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZInterStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZUnionStore("godis3", "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZUnionStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ZInterStore("godis3", "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZInterStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZUnionStore("godis3", "godis1", "godis2")
	assert.NotNil(t, err)
	_, err = redisBroken.ZUnionStoreWithParams("godis3", *param, "godis1", "godis2")
	assert.NotNil(t, err)
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
		r := NewRedis(option)
		defer r.Close()
		r.Subscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)

	pubsub2 := &RedisPubSub{
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
	redis1 := NewRedis(option)
	defer redis1.Close()
	pubsub2.redis = redis1
	go func() {
		pubsub2.Subscribe("godis1")
		pubsub2.process(redis1)
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis1", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	pubsub2.UnSubscribe("godis1")
	time.Sleep(500 * time.Millisecond)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e := redisBroken.Publish("godis1", "publish a message to godis channel")
	assert.NotNil(t, e)
	pubsub2.UnSubscribe("godis1")
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Publish("godis1", "publish a message to godis channel")
	assert.NotNil(t, e)
	pubsub2.UnSubscribe("godis1")
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
		r := NewRedis(option)
		defer r.Close()
		r.PSubscribe(pubsub, "godis")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)

	pubsub2 := &RedisPubSub{
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
	redis1 := NewRedis(option)
	defer redis1.Close()
	pubsub2.redis = redis1
	pubsub2.PSubscribe("godis1")
	pubsub2.PUnSubscribe("godis1")
	time.Sleep(500 * time.Millisecond)

	pubsub3 := &RedisPubSub{
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
	redis2 := NewRedis(option)
	defer redis2.Close()
	pubsub3.redis = redis2
	go func() {
		//pubsub2.PSubscribe("godis1")
		pubsub3.proceedWithPatterns(redis2, "godis2")
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis2", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	e := pubsub3.PUnSubscribe("godis2")
	assert.Nil(t, e)
	time.Sleep(500 * time.Millisecond)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.Publish("godis1", "publish a message to godis channel")
	assert.NotNil(t, e)
	pubsub2.PUnSubscribe("godis1")
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Publish("godis1", "publish a message to godis channel")
	assert.NotNil(t, e)
	pubsub2.PUnSubscribe("godis1")
}

func TestRedis_RandomKey(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	s, e := redis.RandomKey()
	assert.Nil(t, e)
	assert.Equal(t, "godis", s)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.RandomKey()
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.RandomKey()
	assert.NotNil(t, e)
}

func TestRedis_Bitop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
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

	i, e := redis.BitOp(BitOpAnd, "and-result", "bit-1", "bit-2")
	assert.Nil(t, e)
	assert.Equal(t, int64(1), i)

	b, e = redis.GetBit("and-result", 0)
	assert.Nil(t, e)
	assert.Equal(t, true, b)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e = redisBroken.SetBit("bit-1", 0, "1")
	assert.Nil(t, e)
	_, e = redisBroken.BitOp(BitOpAnd, "and-result", "bit-1", "bit-2")
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.SetBit("bit-1", 0, "1")
	assert.NotNil(t, e)
	_, e = redisBroken.BitOp(BitOpAnd, "and-result", "bit-1", "bit-2")
	assert.NotNil(t, e)
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
			keywordMatch: []byte("godis*"),
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, e := redisBroken.Scan(cursor, params)
	assert.NotNil(t, e)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, e = redisBroken.Scan(cursor, params)
	assert.NotNil(t, e)
}

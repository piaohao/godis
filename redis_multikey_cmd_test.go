package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis_Keys(t *testing.T) {
}

func TestRedis_Exists(t *testing.T) {
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
}

func TestRedis_BlpopTimout(t *testing.T) {
}

func TestRedis_Brpop(t *testing.T) {
}

func TestRedis_BrpopTimout(t *testing.T) {
}

func TestRedis_Mget(t *testing.T) {
}

func TestRedis_Mset(t *testing.T) {
}

func TestRedis_Msetnx(t *testing.T) {
}

func TestRedis_Rename(t *testing.T) {
}

func TestRedis_Renamenx(t *testing.T) {
}

func TestRedis_Brpoplpush(t *testing.T) {
}

func TestRedis_Sdiff(t *testing.T) {
}

func TestRedis_Sdiffstore(t *testing.T) {
}

func TestRedis_Sinter(t *testing.T) {
}

func TestRedis_Sinterstore(t *testing.T) {
}

func TestRedis_Smove(t *testing.T) {
}

func TestRedis_Sort(t *testing.T) {
}

func TestRedis_SortMulti(t *testing.T) {
}

func TestRedis_Sunion(t *testing.T) {
}

func TestRedis_Sunionstore(t *testing.T) {
}

func TestRedis_Watch(t *testing.T) {
}

func TestRedis_Unwatch(t *testing.T) {
}

func TestRedis_Zinterstore(t *testing.T) {
}

func TestRedis_ZinterstoreWithParams(t *testing.T) {
}

func TestRedis_Zunionstore(t *testing.T) {
}

func TestRedis_ZunionstoreWithParams(t *testing.T) {
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
		err := r.Subscribe(pubsub, "godis")
		assert.Nil(t, err)
	}()
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
	redis.Publish("godis", "publish a message to godis channel")
	//sleep mills, ensure message can publish to subscribers
	time.Sleep(500 * time.Millisecond)
}

func TestRedis_Psubscribe(t *testing.T) {
}

func TestRedis_RandomKey(t *testing.T) {
}

func TestRedis_Bitop(t *testing.T) {
}

func TestRedis_Scan(t *testing.T) {
}

func TestRedis_Pfmerge(t *testing.T) {
}

package godis_test

import (
	"github.com/gogf/gf/g/test/gtest"
	"github.com/piaohao/godis"
	"testing"
	"time"
)

var option = godis.Option{
	Host:     "10.1.1.63",
	Password: "123456",
	//Host: "localhost",
	Port: 6379,
	Db:   0,
}

func Test_GetSet(t *testing.T) {
	gtest.Case(t, func() {
		redis := godis.NewRedis(option)
		err := redis.Connect()
		defer redis.Close()
		gtest.Assert(err, nil)
		ret, err := redis.Set("godis", "1")
		gtest.Assert(err, nil)
		t.Log(ret)

		arr, err := redis.Get("godis")
		gtest.Assert(err, nil)
		t.Log(string(arr))

		count, err := redis.Del("godis")
		gtest.Assert(err, nil)
		t.Log(count)

		count, err = redis.Del("godis", "godis2")
		gtest.Assert(err, nil)
		t.Log(count)

		arr, err = redis.Get("godis")
		gtest.Assert(err, nil)
		t.Log(string(arr))

		info, err := redis.Info()
		gtest.Assert(err, nil)
		t.Log(info)
	})
}

func Test_Pool(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(option)
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, err := pool.GetResource()
		defer redis.Close()
		gtest.Assert(err, nil)
		arr, err := redis.Info()
		gtest.Assert(err, nil)
		t.Log(string(arr))

		keys, err := redis.Keys("*")
		gtest.Assert(err, nil)
		t.Log(keys)
	})
}

func Test_PubSub(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(option)
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		{
			redis, _ := pool.GetResource()
			reply, err := redis.Exists("godis")
			gtest.Assert(err, nil)
			gtest.Assert(reply, 0)
			redis.Close()
		}

		{
			redis, err := pool.GetResource()
			gtest.Assert(err, nil)
			_, err = redis.Publish("godis", "godis pubsub")
			gtest.Assert(err, nil)
			redis.Close()
		}
		{
			redis, err := pool.GetResource()
			gtest.Assert(err, nil)
			reply, err := redis.Exists("godis")
			gtest.Assert(err, nil)
			gtest.Assert(reply, 0)
			redis.Close()
		}
		go func() {
			redis, err := pool.GetResource()
			gtest.Assert(err, nil)
			pubsub := &godis.RedisPubSub{
				Redis: redis,
				OnMessage: func(channel, message string) {
					t.Log(channel, message)
				},
				OnSubscribe: func(channel string, subscribedChannels int) {
					t.Log(channel, subscribedChannels)
				},
				OnPong: func(channel string) {
					t.Log("receive pong")
				},
			}
			newErr := redis.Subscribe(pubsub, "godis")
			gtest.Assert(newErr, nil)
		}()
		time.Sleep(1 * time.Second)
		{
			redis, err := pool.GetResource()
			gtest.Assert(err, nil)
			_, err = redis.Publish("godis", "godis pubsub")
			gtest.Assert(err, nil)
			redis.Close()
		}
		time.Sleep(1 * time.Second)
	})
}

func Test_PubSub2(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(option)
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, _ := pool.GetResource()
		defer redis.Close()
		t.Log(redis.PubsubChannels("godis"))
	})
}

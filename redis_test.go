package godis_test

import (
	"github.com/gogf/gf/g/test/gtest"
	"github.com/piaohao/godis"
	"testing"
	"time"
)

func Test_GetSet(t *testing.T) {
	gtest.Case(t, func() {
		redis := godis.NewRedis(godis.ShardInfo{
			Host: "172.17.0.2",
			Port: 6379,
			Db:   0,
		})
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
		factory := godis.NewFactory(godis.ShardInfo{
			Host: "172.17.0.2",
			Port: 6379,
			Db:   0,
		})
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
		factory := godis.NewFactory(godis.ShardInfo{
			Host:     "10.1.1.63",
			Port:     6379,
			Db:       0,
			Password: "123456",
		})
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		{
			redis, _ := pool.GetResource()
			reply, err := redis.Exists("gf")
			gtest.Assert(err, nil)
			gtest.Assert(reply, 0)
			redis.Close()
		}

		{
			redis, err := pool.GetResource()
			_, err = redis.Publish("gf", "godis pubsub")
			gtest.Assert(err, nil)
			redis.Close()
		}
		{
			redis, err := pool.GetResource()
			gtest.Assert(err, nil)
			reply, err := redis.Exists("gf")
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
					t.Log("recieve pong")
				},
			}
			newErr := redis.Subscribe(pubsub, "gf")
			gtest.Assert(newErr, nil)
		}()
		time.Sleep(1 * time.Second)
		{
			redis, err := pool.GetResource()
			_, err = redis.Publish("gf", "godis pubsub")
			gtest.Assert(err, nil)
			redis.Close()
		}
		time.Sleep(1 * time.Second)
	})
}

func Test_PubSub2(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(godis.ShardInfo{
			Host:     "10.1.1.63",
			Port:     6379,
			Db:       0,
			Password: "123456",
		})
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, _ := pool.GetResource()
		defer redis.Close()
		t.Log(redis.PubsubChannels("gf"))
	})
}

func Test_PubSub3(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(godis.ShardInfo{
			Host:     "10.1.1.63",
			Port:     6379,
			Db:       0,
			Password: "123456",
		})
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, _ := pool.GetResource()
		defer redis.Close()
		pubsub := &godis.RedisPubSub{
			Redis: redis,
			OnMessage: func(channel, message string) {
				t.Log(channel, message)
			},
			OnSubscribe: func(channel string, subscribedChannels int) {
				t.Log(channel, subscribedChannels)
			},
			OnPong: func(channel string) {
				t.Log("recieve pong")
			},
		}
		newErr := redis.Subscribe(pubsub, "gf1")
		gtest.Assert(newErr, nil)
	})
}

func Test_Cluster(t *testing.T) {
	gtest.Case(t, func() {
		//cluster := godis.NewRedisCluster([]string{"192.168.6.149:8001", "192.168.6.149:8002", "192.168.6.149:8003", "192.168.6.149:8004", "192.168.6.149:8005", "192.168.6.149:8006"},
		//	0, 0, 1, "", godis.PoolConfig{})

		cluster := godis.NewRedisCluster([]string{"192.168.1.6:8001", "192.168.1.6:8002", "192.168.1.6:8003", "192.168.1.6:8004", "192.168.1.6:8005", "192.168.1.6:8006"},
			0, 0, 1, "", godis.PoolConfig{})
		_, err := cluster.Set("cluster", "godis cluster")
		gtest.Assert(err, nil)
		reply, err := cluster.Get("cluster")
		gtest.Assert(err, nil)
		gtest.Assert(reply, "godis cluster")

		int64Reply, err := cluster.Exists("cluster", "cluster1")
		gtest.AssertNE(err, nil)
		int64Reply, err = cluster.Exists("cluster", "cluster")
		gtest.AssertEQ(err, nil)
		gtest.Assert(int64Reply, 2)
	})
}

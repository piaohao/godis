package godis_test

import (
	"fmt"
	"github.com/gogf/gf/g/test/gtest"
	"github.com/piaohao/godis"
	"testing"
	"time"
)

func _TestCluster_Basic(t *testing.T) {
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

func _TestCluster_PubSub(t *testing.T) {
	gtest.Case(t, func() {
		cluster := godis.NewRedisCluster([]string{"192.168.6.149:8001", "192.168.6.149:8002", "192.168.6.149:8003", "192.168.6.149:8004", "192.168.6.149:8005", "192.168.6.149:8006"},
			0, 0, 1, "", godis.PoolConfig{})

		//cluster := godis.NewRedisCluster([]string{"192.168.1.6:8001", "192.168.1.6:8002", "192.168.1.6:8003", "192.168.1.6:8004", "192.168.1.6:8005", "192.168.1.6:8006"},
		//	0, 0, 1, "", godis.PoolConfig{})

		go func() {
			err := cluster.Subscribe(&godis.RedisPubSub{
				OnMessage: func(channel, message string) {
					t.Log(fmt.Sprintf("channel:%s,receive:%s\n", channel, message))
					fmt.Println(fmt.Sprintf("channel:%s,receive:%s\n", channel, message))
				},
				OnPMessage: nil,
				OnSubscribe: func(channel string, subscribedChannels int) {
					t.Log(fmt.Sprintf("channel:%s connect one, total subscribers:%d\n", channel, subscribedChannels))
				},
				OnUnsubscribe: func(channel string, subscribedChannels int) {
					t.Log(fmt.Sprintf("channel:%s disconnect one, total subscribers:%d\n", channel, subscribedChannels))
				},
				OnPUnsubscribe: nil,
				OnPSubscribe:   nil,
				OnPong:         nil,
			}, "channel-godis")
			gtest.Assert(err, nil)
		}()
		time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			c, err := cluster.Publish("channel-godis", fmt.Sprintf("publish message:%d\n", i))
			gtest.Assert(err, nil)
			t.Log(c)
			time.Sleep(20 * time.Millisecond)
		}
		time.Sleep(100 * time.Second)

	})
}

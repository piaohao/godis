package godis

import (
	"github.com/stretchr/testify/assert"
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
}

func TestRedisCluster_Eval(t *testing.T) {
}

func TestRedisCluster_Evalsha(t *testing.T) {
}

func TestRedisCluster_Exists(t *testing.T) {
}

func TestRedisCluster_Expire(t *testing.T) {
}

func TestRedisCluster_ExpireAt(t *testing.T) {
}

func TestRedisCluster_Geoadd(t *testing.T) {
}

func TestRedisCluster_GeoaddByMap(t *testing.T) {
}

func TestRedisCluster_Geodist(t *testing.T) {
}

func TestRedisCluster_Geohash(t *testing.T) {
}

func TestRedisCluster_Geopos(t *testing.T) {
}

func TestRedisCluster_Georadius(t *testing.T) {
}

func TestRedisCluster_GeoradiusByMember(t *testing.T) {
}

func TestRedisCluster_Get(t *testing.T) {
}

func TestRedisCluster_GetSet(t *testing.T) {
}

func TestRedisCluster_Getbit(t *testing.T) {
}

func TestRedisCluster_Getrange(t *testing.T) {
}

func TestRedisCluster_Hdel(t *testing.T) {
}

func TestRedisCluster_Hexists(t *testing.T) {
}

func TestRedisCluster_Hget(t *testing.T) {
}

func TestRedisCluster_HgetAll(t *testing.T) {
}

func TestRedisCluster_HincrBy(t *testing.T) {
}

func TestRedisCluster_HincrByFloat(t *testing.T) {
}

func TestRedisCluster_Hkeys(t *testing.T) {
}

func TestRedisCluster_Hlen(t *testing.T) {
}

func TestRedisCluster_Hmget(t *testing.T) {
}

func TestRedisCluster_Hmset(t *testing.T) {
}

func TestRedisCluster_Hscan(t *testing.T) {
}

func TestRedisCluster_Hset(t *testing.T) {
}

func TestRedisCluster_Hsetnx(t *testing.T) {
}

func TestRedisCluster_Hvals(t *testing.T) {
}

func TestRedisCluster_Incr(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	cluster.Del("godis")
	for i := 0; i < 10000; i++ {
		cluster.Incr("godis")
	}
	reply, _ := cluster.Get("godis")
	if reply != "10000" {
		t.Errorf("want 10000,but %s", reply)
	}
}

func TestRedisCluster_IncrBy(t *testing.T) {
}

func TestRedisCluster_IncrByFloat(t *testing.T) {
}

func TestRedisCluster_Keys(t *testing.T) {
}

func TestRedisCluster_Lindex(t *testing.T) {
}

func TestRedisCluster_Linsert(t *testing.T) {
}

func TestRedisCluster_Llen(t *testing.T) {
}

func TestRedisCluster_Lpop(t *testing.T) {
}

func TestRedisCluster_Lpush(t *testing.T) {
}

func TestRedisCluster_Lpushx(t *testing.T) {
}

func TestRedisCluster_Lrange(t *testing.T) {
}

func TestRedisCluster_Lrem(t *testing.T) {
}

func TestRedisCluster_Lset(t *testing.T) {
}

func TestRedisCluster_Ltrim(t *testing.T) {
}

func TestRedisCluster_Mget(t *testing.T) {
}

func TestRedisCluster_Move(t *testing.T) {
}

func TestRedisCluster_Mset(t *testing.T) {
}

func TestRedisCluster_Msetnx(t *testing.T) {
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

func TestRedisCluster_Rpop(t *testing.T) {
}

func TestRedisCluster_Rpoplpush(t *testing.T) {
}

func TestRedisCluster_Rpush(t *testing.T) {
}

func TestRedisCluster_Rpushx(t *testing.T) {
}

func TestRedisCluster_Sadd(t *testing.T) {
}

func TestRedisCluster_Scan(t *testing.T) {
}

func TestRedisCluster_Scard(t *testing.T) {
}

func TestRedisCluster_ScriptExists(t *testing.T) {
}

func TestRedisCluster_ScriptLoad(t *testing.T) {
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

func TestRedisCluster_Sinter(t *testing.T) {
}

func TestRedisCluster_Sinterstore(t *testing.T) {
}

func TestRedisCluster_Sismember(t *testing.T) {
}

func TestRedisCluster_Smembers(t *testing.T) {
}

func TestRedisCluster_Smove(t *testing.T) {
}

func TestRedisCluster_Sort(t *testing.T) {
}

func TestRedisCluster_SortMulti(t *testing.T) {
}

func TestRedisCluster_Spop(t *testing.T) {
}

func TestRedisCluster_SpopBatch(t *testing.T) {
}

func TestRedisCluster_Srandmember(t *testing.T) {
}

func TestRedisCluster_SrandmemberBatch(t *testing.T) {
}

func TestRedisCluster_Srem(t *testing.T) {
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
}

func TestRedisCluster_Unwatch(t *testing.T) {
}

func TestRedisCluster_Watch(t *testing.T) {
}

func TestRedisCluster_Zadd(t *testing.T) {
}

func TestRedisCluster_ZaddByMap(t *testing.T) {
}

func TestRedisCluster_Zcard(t *testing.T) {
}

func TestRedisCluster_Zcount(t *testing.T) {
}

func TestRedisCluster_Zincrby(t *testing.T) {
}

func TestRedisCluster_Zinterstore(t *testing.T) {
}

func TestRedisCluster_ZinterstoreWithParams(t *testing.T) {
}

func TestRedisCluster_Zlexcount(t *testing.T) {
}

func TestRedisCluster_Zrange(t *testing.T) {
}

func TestRedisCluster_ZrangeByLex(t *testing.T) {
}

func TestRedisCluster_ZrangeByLexBatch(t *testing.T) {
}

func TestRedisCluster_ZrangeByScore(t *testing.T) {
}

func TestRedisCluster_ZrangeByScoreBatch(t *testing.T) {
}

func TestRedisCluster_ZrangeByScoreWithScores(t *testing.T) {
}

func TestRedisCluster_ZrangeByScoreWithScoresBatch(t *testing.T) {
}

func TestRedisCluster_ZrangeWithScores(t *testing.T) {
}

func TestRedisCluster_Zrank(t *testing.T) {
}

func TestRedisCluster_Zrem(t *testing.T) {
}

func TestRedisCluster_ZremrangeByLex(t *testing.T) {
}

func TestRedisCluster_ZremrangeByRank(t *testing.T) {
}

func TestRedisCluster_ZremrangeByScore(t *testing.T) {
}

func TestRedisCluster_Zrevrange(t *testing.T) {
}

func TestRedisCluster_ZrevrangeByLex(t *testing.T) {
}

func TestRedisCluster_ZrevrangeByLexBatch(t *testing.T) {
}

func TestRedisCluster_ZrevrangeByScore(t *testing.T) {
}

func TestRedisCluster_ZrevrangeByScoreWithScores(t *testing.T) {
}

func TestRedisCluster_ZrevrangeByScoreWithScoresBatch(t *testing.T) {
}

func TestRedisCluster_ZrevrangeWithScores(t *testing.T) {
}

func TestRedisCluster_Zrevrank(t *testing.T) {
}

func TestRedisCluster_Zscan(t *testing.T) {
}

func TestRedisCluster_Zscore(t *testing.T) {
}

func TestRedisCluster_Zunionstore(t *testing.T) {
}

func TestRedisCluster_ZunionstoreWithParams(t *testing.T) {
}

func Test_newRedisClusterCommand(t *testing.T) {
}

func Test_newRedisClusterConnectionHandler(t *testing.T) {
}

func Test_newRedisClusterHashTagUtil(t *testing.T) {
}

func Test_newRedisClusterInfoCache(t *testing.T) {
}

func Test_redisClusterCommand_releaseConnection(t *testing.T) {
}

func Test_redisClusterCommand_run(t *testing.T) {
}

func Test_redisClusterCommand_runBatch(t *testing.T) {
}

func Test_redisClusterCommand_runWithAnyNode(t *testing.T) {
}

func Test_redisClusterCommand_runWithRetries(t *testing.T) {
}

func Test_redisClusterConnectionHandler_getConnection(t *testing.T) {
}

func Test_redisClusterConnectionHandler_getConnectionFromNode(t *testing.T) {
}

func Test_redisClusterConnectionHandler_getConnectionFromSlot(t *testing.T) {
}

func Test_redisClusterConnectionHandler_getNodes(t *testing.T) {
}

func Test_redisClusterConnectionHandler_renewSlotCache(t *testing.T) {
}

func Test_redisClusterHashTagUtil_extractHashTag(t *testing.T) {
}

func Test_redisClusterHashTagUtil_getHashTag(t *testing.T) {
}

func Test_redisClusterHashTagUtil_isClusterCompliantMatchPattern(t *testing.T) {
}

func Test_redisClusterInfoCache_assignSlotToNode(t *testing.T) {
}

func Test_redisClusterInfoCache_assignSlotsToNode(t *testing.T) {
}

func Test_redisClusterInfoCache_discoverClusterNodesAndSlots(t *testing.T) {
}

func Test_redisClusterInfoCache_discoverClusterSlots(t *testing.T) {
}

func Test_redisClusterInfoCache_generateHostAndPort(t *testing.T) {
}

func Test_redisClusterInfoCache_getAssignedSlotArray(t *testing.T) {
}

func Test_redisClusterInfoCache_getNode(t *testing.T) {
}

func Test_redisClusterInfoCache_getNodes(t *testing.T) {
}

func Test_redisClusterInfoCache_getShuffledNodesPool(t *testing.T) {
}

func Test_redisClusterInfoCache_getSlotPool(t *testing.T) {
}

func Test_redisClusterInfoCache_renewClusterSlots(t *testing.T) {
}

func Test_redisClusterInfoCache_reset(t *testing.T) {
}

func Test_redisClusterInfoCache_setupNodeIfNotExist(t *testing.T) {
}

func Test_redisClusterInfoCache_shuffle(t *testing.T) {
}

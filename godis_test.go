package godis_test

import (
	"github.com/gogf/gf/g/test/gtest"
	"github.com/piaohao/godis"
	"strings"
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
		pool := godis.NewPool(godis.PoolConfig{}, *factory)
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
		pool := godis.NewPool(godis.PoolConfig{}, *factory)
		{
			redis, _ := pool.GetResource()
			reply, err := redis.Exists("gf")
			gtest.Assert(err, nil)
			gtest.Assert(reply, 0)
			redis.Close()
		}

		{
			redis, err := pool.GetResource()
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
			err = pubsub.Subscribe("gf")
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
		time.Sleep(500 * time.Second)
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
		pool := godis.NewPool(godis.PoolConfig{}, *factory)
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
		pool := godis.NewPool(godis.PoolConfig{}, *factory)
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

func Test_Cmd(t *testing.T) {
	str := `PING, SET, GET, QUIT, EXISTS, DEL, UNLINK, TYPE, FLUSHDB, KEYS, RANDOMKEY, RENAME, RENAMENX,
	RENAMEX, DBSIZE, EXPIRE, EXPIREAT, TTL, SELECT, MOVE, FLUSHALL, GETSET, MGET, SETNX, SETEX,
	MSET, MSETNX, DECRBY, DECR, INCRBY, INCR, APPEND, SUBSTR, HSET, HGET, HSETNX, HMSET, HMGET,
	HINCRBY, HEXISTS, HDEL, HLEN, HKEYS, HVALS, HGETALL, RPUSH, LPUSH, LLEN, LRANGE, LTRIM, LINDEX,
	LSET, LREM, LPOP, RPOP, RPOPLPUSH, SADD, SMEMBERS, SREM, SPOP, SMOVE, SCARD, SISMEMBER, SINTER,
	SINTERSTORE, SUNION, SUNIONSTORE, SDIFF, SDIFFSTORE, SRANDMEMBER, ZADD, ZRANGE, ZREM, ZINCRBY,
	ZRANK, ZREVRANK, ZREVRANGE, ZCARD, ZSCORE, MULTI, DISCARD, EXEC, WATCH, UNWATCH, SORT, BLPOP,
	BRPOP, AUTH, SUBSCRIBE, PUBLISH, UNSUBSCRIBE, PSUBSCRIBE, PUNSUBSCRIBE, PUBSUB, ZCOUNT,
	ZRANGEBYSCORE, ZREVRANGEBYSCORE, ZREMRANGEBYRANK, ZREMRANGEBYSCORE, ZUNIONSTORE, ZINTERSTORE,
	ZLEXCOUNT, ZRANGEBYLEX, ZREVRANGEBYLEX, ZREMRANGEBYLEX, SAVE, BGSAVE, BGREWRITEAOF, LASTSAVE,
	SHUTDOWN, INFO, MONITOR, SLAVEOF, CONFIG, STRLEN, SYNC, LPUSHX, PERSIST, RPUSHX, ECHO, LINSERT,
	DEBUG, BRPOPLPUSH, SETBIT, GETBIT, BITPOS, SETRANGE, GETRANGE, EVAL, EVALSHA, SCRIPT, SLOWLOG,
	OBJECT, BITCOUNT, BITOP, SENTINEL, DUMP, RESTORE, PEXPIRE, PEXPIREAT, PTTL, INCRBYFLOAT,
	PSETEX, CLIENT, TIME, MIGRATE, HINCRBYFLOAT, SCAN, HSCAN, SSCAN, ZSCAN, WAIT, CLUSTER, ASKING,
	PFADD, PFCOUNT, PFMERGE, READONLY, GEOADD, GEODIST, GEOHASH, GEOPOS, GEORADIUS, GEORADIUS_RO,
	GEORADIUSBYMEMBER, GEORADIUSBYMEMBER_RO, MODULE, BITFIELD, HSTRLEN, TOUCH, SWAPDB, MEMORY,
	XADD, XLEN, XDEL, XTRIM, XRANGE, XREVRANGE, XREAD, XACK, XGROUP, XREADGROUP, XPENDING, XCLAIM`
	for _, item := range strings.Split(str, ",") {
		println(strings.TrimSpace(item) + "=newProtocolCommand(\"" + strings.TrimSpace(item) + "\")")
	}
	str = `AGGREGATE, ALPHA, ASC, BY, DESC, GET, LIMIT, MESSAGE, NO, NOSORT, PMESSAGE, PSUBSCRIBE,
    PUNSUBSCRIBE, OK, ONE, QUEUED, SET, STORE, SUBSCRIBE, UNSUBSCRIBE, WEIGHTS, WITHSCORES,
    RESETSTAT, REWRITE, RESET, FLUSH, EXISTS, LOAD, KILL, LEN, REFCOUNT, ENCODING, IDLETIME,
    GETNAME, SETNAME, LIST, MATCH, COUNT, PING, PONG, UNLOAD, REPLACE, KEYS, PAUSE, DOCTOR, 
    BLOCK, NOACK, STREAMS, KEY, CREATE, MKSTREAM, SETID, DESTROY, DELCONSUMER, MAXLEN, GROUP, 
    IDLE, TIME, RETRYCOUNT, FORCE`
	println()
	println()
	for _, item := range strings.Split(str, ",") {
		println(strings.TrimSpace(item) + "=newKeyword(\"" + strings.TrimSpace(item) + "\")")
	}
}

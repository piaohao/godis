package godis

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func Test_multiKeyPipelineBase_Bgrewriteaof(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	p := redis.Pipelined()
	r, err := p.Bgrewriteaof()
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(r.Get())
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)

}

func Test_multiKeyPipelineBase_Bgsave(t *testing.T) {
	//sleep 2 second to wait previous Bgrewriteaof  stop,otherwise this case
	time.Sleep(2 * time.Second)
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	p := redis.Pipelined()
	r, err := p.Bgsave()
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(r.Get())
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}

func Test_multiKeyPipelineBase_Bitop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
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

	i, e := p.Bitop(BitopAnd, "and-result", "bit-1", "bit-2")
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToInt64Reply(i.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(1), resp)

	b, e = redis.Getbit("and-result", 0)
	assert.Nil(t, e)
	assert.Equal(t, true, b)
}

func Test_multiKeyPipelineBase_Blpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")

	arr, e := p.Blpop("job", "command", "request", "0")
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToStringArrayReply(arr.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, resp)

}

func Test_multiKeyPipelineBase_BlpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		p := redis.Pipelined()
		_, e := p.BlpopTimout(5, "command", "update system...")
		assert.Nil(t, e)
	}()
	time.Sleep(500 * time.Millisecond)
	redis2 := NewRedis(option)
	redis2.Lpush("command", "update system...")
	redis2.Lpush("request", "visit page")
	time.Sleep(500 * time.Millisecond)
}

func Test_multiKeyPipelineBase_Brpop(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	redis.Lpush("command", "update system...")
	redis.Lpush("request", "visit page")

	arr, e := p.Brpop("job", "command", "request", "0")
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToStringArrayReply(arr.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"command", "update system..."}, resp)

}

func Test_multiKeyPipelineBase_BrpopTimout(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		p := redis.Pipelined()
		_, e := p.BrpopTimout(5, "command", "update system...")
		assert.Nil(t, e)
	}()
	time.Sleep(1 * time.Second)
	redis2 := NewRedis(option)
	redis2.Lpush("command", "update system...")
	redis2.Lpush("request", "visit page")
	time.Sleep(1 * time.Second)

}

func Test_multiKeyPipelineBase_Brpoplpush(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	go func() {
		p := redis.Pipelined()
		_, e := p.Brpoplpush("command", "update system...", 5)
		assert.Nil(t, e)
	}()
	time.Sleep(1 * time.Second)
	redis2 := NewRedis(option)
	redis2.Lpush("command", "update system...")
	redis2.Lpush("request", "visit page")
	time.Sleep(1 * time.Second)

}

func Test_multiKeyPipelineBase_ClusterAddSlots(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	p := redis.Pipelined()
	slots, err := p.ClusterAddSlots(10000)
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(slots.Get())
	assert.NotNil(t, err)
	assert.Equal(t, "", resp)
}

func Test_multiKeyPipelineBase_ClusterDelSlots(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterGetKeysInSlot(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterInfo(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	p := redis.Pipelined()
	s, err := p.ClusterInfo()
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(s.Get())
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}

func Test_multiKeyPipelineBase_ClusterMeet(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterNodes(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	p := redis.Pipelined()
	s, err := p.ClusterNodes()
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(s.Get())
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)

	nodeId := resp[:strings.Index(resp, " ")]
	redis1 := NewRedis(option1)
	defer redis1.Close()
	redis1.ClusterSlaves(nodeId)
	//assert.Nil(t, err)
	//assert.NotEmpty(t, slaves)
}

func Test_multiKeyPipelineBase_ClusterSetSlotImporting(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotMigrating(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotNode(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigGet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	reply, err := p.ConfigGet("timeout")
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringArrayReply(reply.Get())
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "0"}, resp)
}

func Test_multiKeyPipelineBase_ConfigResetStat(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	_, err := p.ConfigResetStat()
	assert.Nil(t, err)
	p.Sync()
}

func Test_multiKeyPipelineBase_ConfigSet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	reply, err := p.ConfigSet("timeout", "30")
	assert.Nil(t, err)
	//assert.Equal(t, "OK", reply)
	reply1, err := p.ConfigGet("timeout")
	assert.Nil(t, err)
	//assert.Equal(t, []string{"timeout", "30"}, reply1)
	reply2, err := p.ConfigSet("timeout", "0")
	assert.Nil(t, err)
	//assert.Equal(t, "OK", reply)
	reply3, err := p.ConfigGet("timeout")
	assert.Nil(t, err)
	//assert.Equal(t, []string{"timeout", "0"}, reply1)
	p.Sync()
	resp1, err := ToStringReply(reply.Get())
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp1)
	resp2, err := ToStringArrayReply(reply1.Get())
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "30"}, resp2)
	resp3, err := ToStringReply(reply2.Get())
	assert.Nil(t, err)
	assert.Equal(t, "\x00\x00", resp3)
	resp4, err := ToStringArrayReply(reply3.Get())
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "0"}, resp4)
}

func Test_multiKeyPipelineBase_DbSize(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	redis.Set("godis1", "good")
	p := redis.Pipelined()
	ret, err := p.DbSize()
	assert.Nil(t, err)
	p.Sync()
	c, err := ToInt64Reply(ret.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)
}

func Test_multiKeyPipelineBase_Del(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Del("godis")
	assert.Nil(t, err)
	obj, err := ToInt64Reply(del.Get())
	assert.NotNil(t, err)
	p.Sync()
	obj, err = ToInt64Reply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(0), obj)
	redis.Close()
}

func Test_multiKeyPipelineBase_Eval(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	s1, err := p.Eval(`return redis.call("get",KEYS[1])`, 1, "godis")
	assert.Nil(t, err)

	s2, err := p.Eval(`return redis.call("set",KEYS[1],ARGV[1])`, 1, "eval", "godis")
	assert.Nil(t, err)

	s3, err := p.Eval(`return redis.call("get",KEYS[1])`, 1, "eval")
	assert.Nil(t, err)

	p.Sync()
	resp1, _ := ToStringReply(s1.Get())
	assert.Equal(t, "good", resp1)
	resp2, _ := ToStringReply(s2.Get())
	assert.Equal(t, "\x00\x00", resp2)
	resp3, _ := ToStringReply(s3.Get())
	assert.Equal(t, "godis", resp3)

}

func Test_multiKeyPipelineBase_Evalsha(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()

	s, err := p.Evalsha("111", 1, "godis")
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(s.Get())
	assert.NotNil(t, err)
	assert.Equal(t, "", resp)
}

func Test_multiKeyPipelineBase_Exists(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	defer redis.Close()

	p := redis.Pipelined()
	del, err := p.Exists("godis")
	assert.Nil(t, err)
	del2, err := p.Exists("godis2")
	assert.Nil(t, err)
	p.Sync()
	obj, err := ToInt64Reply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(1), obj)

	obj, err = ToInt64Reply(del2.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(0), obj)
}

func Test_multiKeyPipelineBase_FlushAll(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	redis = NewRedis(option)
	p := redis.Pipelined()
	del, err := p.FlushAll()
	assert.Nil(t, err)
	obj, err := ToStringReply(del.Get())
	assert.NotNil(t, err)
	p.Sync()
	obj, err = ToStringReply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, "OK", obj)
	redis.Close()

	redis = NewRedis(option)
	ret, _ := redis.Get("godis")
	assert.Equal(t, "", ret)
	redis.Close()
}

func Test_multiKeyPipelineBase_FlushDB(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	redis = NewRedis(option)
	p := redis.Pipelined()
	del, err := p.FlushDB()
	assert.Nil(t, err)
	obj, err := ToStringReply(del.Get())
	assert.NotNil(t, err)
	p.Sync()
	obj, err = ToStringReply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, "OK", obj)
	redis.Close()

	redis = NewRedis(option)
	ret, _ := redis.Get("godis")
	assert.Equal(t, "", ret)
	redis.Close()
}

func Test_multiKeyPipelineBase_Info(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Info()
	assert.Nil(t, err)
	_, err = ToStringReply(del.Get())
	assert.NotNil(t, err)
	p.Sync()
	_, err = ToStringReply(del.Get())
	assert.Nil(t, err)
	redis.Close()
}

func Test_multiKeyPipelineBase_Keys(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	redis = NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Keys("*")
	assert.Nil(t, err)
	obj, err := ToStringArrayReply(del.Get())
	assert.NotNil(t, err)
	p.Sync()
	obj, err = ToStringArrayReply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, []string{"godis"}, obj)
	redis.Close()
}

func Test_multiKeyPipelineBase_Lastsave(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	_, err := p.Lastsave()
	assert.Nil(t, err)
}

func Test_multiKeyPipelineBase_Mget(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()

	s, e := p.Mset("godis1", "good", "godis2", "good")
	assert.Nil(t, e)

	c, e := p.Msetnx("godis1", "good1")
	assert.Nil(t, e)

	arr, e := p.Mget("godis", "godis1", "godis2")
	assert.Nil(t, e)

	p.Sync()
	resp1, e := ToStringReply(s.Get())
	assert.Nil(t, e)
	assert.Equal(t, "OK", resp1)
	resp2, e := ToInt64Reply(c.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(0), resp2)

	resp3, e := ToStringArrayReply(arr.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"good", "good", "good"}, resp3)
}

func Test_multiKeyPipelineBase_Pfcount(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	c, err := redis.Pfadd("godis", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	c, err = redis.Pfadd("godis1", "a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)

	reply1, err := p.Pfcount("godis")
	assert.Nil(t, err)

	reply2, err := p.Pfcount("godis1")
	assert.Nil(t, err)

	reply3, err := p.Pfmerge("godis3", "godis", "godis1")
	assert.Nil(t, err)

	reply4, err := p.Pfcount("godis3")
	assert.Nil(t, err)

	p.Sync()
	resp1, err := ToInt64Reply(reply1.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(4), resp1)
	resp2, err := ToInt64Reply(reply2.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(4), resp2)
	resp3, err := ToStringReply(reply3.Get())
	assert.Nil(t, err)
	assert.Equal(t, "\x00\x00", resp3)
	resp4, err := ToInt64Reply(reply4.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(4), resp4)
}

func Test_multiKeyPipelineBase_Ping(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Ping()
	assert.Nil(t, err)
	p.Sync()
	obj, err := ToStringReply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, "PONG", obj)
	redis.Close()
}

func Test_multiKeyPipelineBase_Publish(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Publish("godis", "good")
	assert.Nil(t, err)
	p.Sync()
	obj, err := ToInt64Reply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(0), obj)
	redis.Close()
}

func Test_multiKeyPipelineBase_RandomKey(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	s, e := p.RandomKey()
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToStringReply(s.Get())
	assert.Nil(t, e)
	assert.Equal(t, "godis", resp)
}

func Test_multiKeyPipelineBase_Rename(t *testing.T) {
	initDb()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	s, e := p.Rename("godis", "godis1")
	assert.Nil(t, e)
	c, e := p.Renamenx("godis1", "godis2")
	assert.Nil(t, e)
	p.Sync()
	resp1, e := ToStringReply(s.Get())
	assert.Nil(t, e)
	assert.Equal(t, "OK", resp1)
	resp2, e := ToInt64Reply(c.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(1), resp2)
}

func Test_multiKeyPipelineBase_Rpoplpush(t *testing.T) {
}

func Test_multiKeyPipelineBase_Save(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	p := redis.Pipelined()
	ret, err := p.Save()
	assert.Nil(t, err)
	p.Sync()
	resp, err := ToStringReply(ret.Get())
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp)
}

func Test_multiKeyPipelineBase_Sdiff(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Sadd("godis1", "1", "2", "3")
	redis.Sadd("godis2", "2", "3", "4")

	p := redis.Pipelined()
	reply1, e := p.Sdiff("godis1", "godis2")
	assert.Nil(t, e)
	reply2, e := p.Sdiffstore("godis3", "godis1", "godis2")
	assert.Nil(t, e)
	reply3, e := p.Sinter("godis1", "godis2")
	assert.Nil(t, e)
	reply4, e := p.Sinterstore("godis4", "godis1", "godis2")
	assert.Nil(t, e)
	reply5, e := p.Sunion("godis1", "godis2")
	assert.Nil(t, e)
	reply6, e := p.Sunionstore("godis5", "godis1", "godis2")
	assert.Nil(t, e)

	p.Sync()
	resp1, e := ToStringArrayReply(reply1.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1"}, resp1)
	resp2, e := ToInt64Reply(reply2.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(1), resp2)
	resp3, e := ToStringArrayReply(reply3.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"2", "3"}, resp3)
	resp4, e := ToInt64Reply(reply4.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(2), resp4)
	resp5, e := ToStringArrayReply(reply5.Get())
	assert.Nil(t, e)
	assert.ElementsMatch(t, []string{"1", "2", "3", "4"}, resp5)
	resp6, e := ToInt64Reply(reply6.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(4), resp6)
}

func Test_multiKeyPipelineBase_Select(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	reply1, err := p.Select(15)
	assert.Nil(t, err)
	reply2, err := p.Select(16)
	assert.Nil(t, err)
	p.Sync()
	resp1, err := ToStringReply(reply1.Get())
	assert.Nil(t, err)
	assert.Equal(t, "OK", resp1)
	resp2, err := ToStringReply(reply2.Get())
	assert.NotNil(t, err)
	assert.Equal(t, "", resp2)
}

func Test_multiKeyPipelineBase_Shutdown(t *testing.T) {
}

func Test_multiKeyPipelineBase_Smove(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Sadd("godis", "1", "2")
	redis.Sadd("godis1", "3", "4")

	p := redis.Pipelined()
	s, e := p.Smove("godis", "godis1", "2")
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToInt64Reply(s.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(1), resp)

	arr, _ := redis.Smembers("godis")
	assert.ElementsMatch(t, []string{"1"}, arr)

	arr, _ = redis.Smembers("godis1")
	assert.ElementsMatch(t, []string{"2", "3", "4"}, arr)
}

func Test_multiKeyPipelineBase_SortMulti(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Lpush("godis", "3", "2", "1", "4", "6", "5")
	param := NewSortingParams().Desc()

	p := redis.Pipelined()
	//arr, e := p.Sort("godis", *param)
	//assert.Nil(t, e)
	//assert.Equal(t, []string{"6", "5", "4", "3", "2", "1"}, arr)
	//
	//param = NewSortingParams().Asc()
	//arr, e = p.Sort("godis", *param)
	//assert.Nil(t, e)
	//assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, arr)

	c, e := p.SortStore("godis", "godis1", *param)
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToInt64Reply(c.Get())
	assert.Nil(t, e)
	assert.Equal(t, int64(6), resp)
}

func Test_multiKeyPipelineBase_Time(t *testing.T) {
}

func Test_multiKeyPipelineBase_Watch(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	s, e := p.Watch("godis")
	assert.Nil(t, e)
	p.Sync()
	resp, e := ToStringReply(s.Get())
	assert.Nil(t, e)
	assert.Equal(t, "OK", resp)
}

func Test_multiKeyPipelineBase_Zinterstore(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	p := redis.Pipelined()
	c, err := redis.ZaddByMap("godis1", map[string]float64{"a": 1, "b": 2, "c": 3})
	assert.Nil(t, err)
	assert.Equal(t, int64(3), c)

	c, err = redis.ZaddByMap("godis2", map[string]float64{"a": 1, "b": 2})
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)

	r1, err := p.Zinterstore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	r2, err := p.ZinterstoreWithParams("godis3", ZparamsSum, "godis1", "godis2")
	assert.Nil(t, err)
	r3, err := p.Zunionstore("godis3", "godis1", "godis2")
	assert.Nil(t, err)
	r4, err := p.ZunionstoreWithParams("godis3", ZparamsMax, "godis1", "godis2")
	assert.Nil(t, err)

	p.Sync()
	resp1, err := ToInt64Reply(r1.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(2), resp1)
	resp2, err := ToInt64Reply(r2.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(2), resp2)
	resp3, err := ToInt64Reply(r3.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(3), resp3)
	resp4, err := ToInt64Reply(r4.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(3), resp4)
}

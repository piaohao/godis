package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiKeyPipelineBase_Bgrewriteaof(t *testing.T) {
}

func Test_multiKeyPipelineBase_Bgsave(t *testing.T) {
}

func Test_multiKeyPipelineBase_Bitop(t *testing.T) {
}

func Test_multiKeyPipelineBase_Blpop(t *testing.T) {
}

func Test_multiKeyPipelineBase_BlpopTimout(t *testing.T) {
}

func Test_multiKeyPipelineBase_Brpop(t *testing.T) {
}

func Test_multiKeyPipelineBase_BrpopTimout(t *testing.T) {
}

func Test_multiKeyPipelineBase_Brpoplpush(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterAddSlots(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterDelSlots(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterGetKeysInSlot(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterInfo(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterMeet(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterNodes(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotImporting(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotMigrating(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotNode(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigGet(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigResetStat(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigSet(t *testing.T) {
}

func Test_multiKeyPipelineBase_DbSize(t *testing.T) {
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
}

func Test_multiKeyPipelineBase_Evalsha(t *testing.T) {
}

func Test_multiKeyPipelineBase_Exists(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()

	redis = NewRedis(option)
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
	redis.Close()
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
}

func Test_multiKeyPipelineBase_Mget(t *testing.T) {
}

func Test_multiKeyPipelineBase_Mset(t *testing.T) {
}

func Test_multiKeyPipelineBase_Msetnx(t *testing.T) {
}

func Test_multiKeyPipelineBase_Pfcount(t *testing.T) {
}

func Test_multiKeyPipelineBase_Pfmerge(t *testing.T) {
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
}

func Test_multiKeyPipelineBase_Rename(t *testing.T) {
}

func Test_multiKeyPipelineBase_Renamenx(t *testing.T) {
}

func Test_multiKeyPipelineBase_Rpoplpush(t *testing.T) {
}

func Test_multiKeyPipelineBase_Save(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sdiff(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sdiffstore(t *testing.T) {}

func Test_multiKeyPipelineBase_Select(t *testing.T) {
}

func Test_multiKeyPipelineBase_Shutdown(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sinter(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sinterstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_Smove(t *testing.T) {

}

func Test_multiKeyPipelineBase_SortMulti(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sunion(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sunionstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_Time(t *testing.T) {
}

func Test_multiKeyPipelineBase_Watch(t *testing.T) {
}

func Test_multiKeyPipelineBase_Zinterstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_ZinterstoreWithParams(t *testing.T) {
}

func Test_multiKeyPipelineBase_Zunionstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_ZunionstoreWithParams(t *testing.T) {
}

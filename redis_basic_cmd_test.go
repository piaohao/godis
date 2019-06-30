package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis_Auth(t *testing.T) {
	redis := NewRedis(option)
	ok, err := redis.Auth("")
	redis.Close()
	assert.NotNil(t, err)
	assert.Equal(t, "", ok)
	redis = NewRedis(option)
	ok, err = redis.Auth("123456")
	redis.Close()
	assert.NotNil(t, err)
	assert.Equal(t, "", ok)
}

func TestRedis_Ping(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Ping()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", ret)
}

func TestRedis_Quit(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Quit()
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)
}

func TestRedis_FlushDB(t *testing.T) {
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	reply, err := redis.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "good", reply)
	redis = NewRedis(option)
	redis.Select(2)
	redis.Set("godis", "good")
	reply, err = redis.FlushDB()
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply, err = redis.Get("godis")
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "", reply)
	redis = NewRedis(option)
	redis.Select(0)
	reply, err = redis.Get("godis")
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "good", reply)
}

func TestRedis_FlushAll(t *testing.T) {
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	reply, err := redis.Get("godis")
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "good", reply)
	redis = NewRedis(option)
	redis.Select(2)
	redis.Set("godis", "good")
	reply, err = redis.FlushAll()
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply, err = redis.Get("godis")
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "", reply)
	redis = NewRedis(option)
	redis.Select(0)
	reply, err = redis.Get("godis")
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "", reply)
}

func TestRedis_DbSize(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	redis.Set("godis1", "good")
	redis.Close()
	redis = NewRedis(option)
	ret, err := redis.DbSize()
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, int64(2), ret)
}

func TestRedis_Select(t *testing.T) {
	redis := NewRedis(option)
	ret, err := redis.Select(15)
	redis = NewRedis(option)
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)
	redis = NewRedis(option)
	ret, err = redis.Select(16)
	redis.Close()
	assert.NotNil(t, err)
	assert.Equal(t, "", ret)
}

func TestRedis_Save(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	ret, err := redis.Save()
	redis.Close()
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)
}

func TestRedis_Bgsave(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	_, err := redis.BgSave()
	redis.Close()
	assert.Nil(t, err)
}

func TestRedis_Bgrewriteaof(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	_, err := redis.BgRewriteAof()
	redis.Close()
	assert.Nil(t, err)
}

func TestRedis_Lastsave(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.LastSave()
	assert.Nil(t, err)
}

// ignore this case,cause it will shutdown redis
func _TestRedisShutdown(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.Shutdown()
	assert.NotNil(t, err)
}

func TestRedis_Info(t *testing.T) {
	redis := NewRedis(option)
	_, err := redis.Info()
	redis.Close()
	assert.Nil(t, err)
	redis = NewRedis(option)
	_, err = redis.Info("stats")
	redis.Close()
	assert.Nil(t, err)
	redis = NewRedis(option)
	_, err = redis.Info("clients", "memory")
	redis.Close()
	assert.NotNil(t, err)
}

func TestRedis_Slaveof(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SlaveOf("localhost", 6379)
	assert.Nil(t, err)
}

func TestRedis_SlaveofNoOne(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SlaveOfNoOne()
	assert.Nil(t, err)
}

func TestRedis_Debug(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	redis = NewRedis(option)
	_, err := redis.Debug(*NewDebugParamsObject("godis"))
	redis.Close()
	assert.Nil(t, err)
}

func TestRedis_ConfigResetStat(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.ConfigResetStat()
	assert.Nil(t, err)
}

func TestRedis_WaitReplicas(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.WaitReplicas(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ret)
}

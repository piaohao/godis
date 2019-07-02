package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Auth("123456")
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Auth("123456")
	assert.NotNil(t, err)
}

func TestRedis_Ping(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Ping()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", ret)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Ping()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Ping()
	assert.NotNil(t, err)
}

func TestRedis_Quit(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.Quit()
	assert.Nil(t, err)
	assert.Equal(t, "OK", ret)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Quit()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Quit()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.FlushDB()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.FlushDB()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.FlushAll()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.FlushAll()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.DbSize()
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.DbSize()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Select(15)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Select(15)
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Save()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Save()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BgSave()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BgSave()
	assert.NotNil(t, err)
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.BgRewriteAof()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.BgRewriteAof()
	assert.NotNil(t, err)
}

func TestRedis_Lastsave(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.LastSave()
	assert.Nil(t, err)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.LastSave()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.LastSave()
	assert.NotNil(t, err)
}

// ignore this case,cause it will shutdown redis
func TestRedis_Shutdown(t *testing.T) {
	redis := NewRedis(&Option{
		Host: "localhost",
		Port: 8888,
	})
	defer redis.Close()
	_, err := redis.Shutdown()
	assert.NotNil(t, err)

	time.Sleep(time.Second)
	redis1 := NewRedis(option)
	defer redis1.Close()
	s, err := redis1.Set("godis", "good")
	assert.Nil(t, err)
	assert.Equal(t, "OK", s)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Shutdown()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Shutdown()
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

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Info()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Info()
	assert.NotNil(t, err)
}

func TestRedis_Slaveof(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SlaveOf("localhost", 6379)
	assert.Nil(t, err)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SlaveOf("localhost", 6379)
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SlaveOf("localhost", 6379)
	assert.NotNil(t, err)
}

func TestRedis_SlaveofNoOne(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SlaveOfNoOne()
	assert.Nil(t, err)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.SlaveOfNoOne()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SlaveOfNoOne()
	assert.NotNil(t, err)
}

func TestRedis_Debug(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	_, err := redis.Debug(*NewDebugParamsObject("godis"))
	assert.Nil(t, err)

	_, err = redis.Debug(*NewDebugParamsSegfault())
	assert.NotNil(t, err) //EOF error

	get, err := redis.Get("godis")
	assert.NotNil(t, err)
	assert.Equal(t, "", get)

	time.Sleep(1 * time.Second)
	redis1 := NewRedis(option)
	defer redis1.Close()
	get, err = redis1.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)

	_, err = redis1.Debug(*NewDebugParamsReload())
	assert.Nil(t, err) //EOF error

	get, err = redis1.Get("godis")
	assert.Nil(t, err)
	assert.Equal(t, "", get)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.Debug(*NewDebugParamsObject("godis"))
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.Debug(*NewDebugParamsObject("godis"))
	assert.NotNil(t, err)
}

func TestRedis_ConfigResetStat(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.ConfigResetStat()
	assert.Nil(t, err)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.ConfigResetStat()
	assert.Nil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ConfigResetStat()
	assert.NotNil(t, err)
}

func TestRedis_WaitReplicas(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	ret, err := redis.WaitReplicas(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ret)

	redisBroken := NewRedis(option1)
	defer redisBroken.Close()
	m, _ := redisBroken.Multi()
	_, err = redisBroken.WaitReplicas(1, 1)
	assert.NotNil(t, err)
	m.Discard()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.WaitReplicas(1, 1)
	assert.NotNil(t, err)
}

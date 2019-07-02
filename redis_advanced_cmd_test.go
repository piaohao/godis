package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis_ConfigGet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	reply, err := redis.ConfigGet("timeout")
	assert.Nil(t, err, "err is nil")
	assert.Equal(t, []string{"timeout", "0"}, reply)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ConfigGet("timeout")
	assert.NotNil(t, err)
}

func TestRedis_ConfigSet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	reply, err := redis.ConfigSet("timeout", "30")
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply1, err := redis.ConfigGet("timeout")
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "30"}, reply1)
	reply, err = redis.ConfigSet("timeout", "0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply1, err = redis.ConfigGet("timeout")
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "0"}, reply1)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ConfigSet("timeout", "30")
	assert.NotNil(t, err)
}

func TestRedis_SlowlogGet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	flushAll()
	arr, err := redis.SlowLogGet()
	assert.Nil(t, err)
	t.Log(arr)
	//assert.NotEmpty(t, arr)

	redis.SlowLogGet(1)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SlowLogGet()
	assert.NotNil(t, err)
}

func TestRedis_SlowlogLen(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	l, err := redis.SlowLogLen()
	assert.Nil(t, err)
	assert.True(t, l >= 0)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SlowLogLen()
	assert.NotNil(t, err)
}

func TestRedis_SlowlogReset(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	str, err := redis.SlowLogReset()
	assert.Nil(t, err)
	assert.Equal(t, "OK", str)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SlowLogReset()
	assert.NotNil(t, err)
}

func TestRedis_ObjectEncoding(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	encode, err := redis.ObjectEncoding("godis")
	assert.Nil(t, err)
	assert.Equal(t, "embstr", encode)
	redis.Set("godis", "12")
	encode, err = redis.ObjectEncoding("godis")
	assert.Nil(t, err)
	assert.Equal(t, "int", encode)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ObjectEncoding("godis")
	assert.NotNil(t, err)
}

func TestRedis_ObjectIdletime(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	time.Sleep(1000 * time.Millisecond)
	idle, err := redis.ObjectIdleTime("godis")
	assert.Nil(t, err)
	assert.True(t, idle > 0)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ObjectIdleTime("godis")
	assert.NotNil(t, err)
}

func TestRedis_ObjectRefcount(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	count, err := redis.ObjectRefCount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.ObjectRefCount("godis")
	assert.NotNil(t, err)
}

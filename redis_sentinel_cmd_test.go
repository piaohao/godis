package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis_SentinelFailover(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SentinelFailOver("a")
	assert.NotNil(t, err)
	assert.Equal(t, "", s)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelFailOver("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelGetMasterAddrByName(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelGetMasterAddrByName("a")
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelGetMasterAddrByName("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelMasters(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelMasters()
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelMasters()
	assert.NotNil(t, err)
}

func TestRedis_SentinelMonitor(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelMonitor("a", "", 0, 0)
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelMonitor("a", "", 0, 0)
	assert.NotNil(t, err)
}

func TestRedis_SentinelRemove(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelRemove("a")
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelRemove("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelReset(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelReset("a")
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelReset("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelSet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelSet("a", nil)
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelSet("a", nil)
	assert.NotNil(t, err)
}

func TestRedis_SentinelSlaves(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelSlaves("a")
	assert.NotNil(t, err)

	redisBroken := NewRedis(option)
	defer redisBroken.Close()
	redisBroken.client.connection.host = "localhost1"
	redisBroken.Close()
	_, err = redisBroken.SentinelSlaves("a")
	assert.NotNil(t, err)
}

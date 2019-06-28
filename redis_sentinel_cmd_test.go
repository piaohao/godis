package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis_SentinelFailover(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	s, err := redis.SentinelFailover("a")
	assert.NotNil(t, err)
	assert.Equal(t, "", s)
}

func TestRedis_SentinelGetMasterAddrByName(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelGetMasterAddrByName("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelMasters(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelMasters()
	assert.NotNil(t, err)
}

func TestRedis_SentinelMonitor(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelMonitor("a", "", 0, 0)
	assert.NotNil(t, err)
}

func TestRedis_SentinelRemove(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelRemove("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelReset(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelReset("a")
	assert.NotNil(t, err)
}

func TestRedis_SentinelSet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelSet("a", nil)
	assert.NotNil(t, err)
}

func TestRedis_SentinelSlaves(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SentinelSlaves("a")
	assert.NotNil(t, err)
}

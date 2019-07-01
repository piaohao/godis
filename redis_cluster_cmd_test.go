package godis

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var option1 = &Option{
	Host: "localhost",
	Port: 7000,
}

func TestRedis_ClusterAddSlots(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	slots, err := redis.ClusterAddSlots(10000)
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterCountKeysInSlot(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	slots, err := redis.ClusterCountKeysInSlot(10000)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), slots)
}

func TestRedis_ClusterDelSlots(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterDelSlots(10000)
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterFailover(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	slots, err := redis.ClusterFailOver()
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterFlushSlots(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterFlushSlots()
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterForget(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	_, err := redis.ClusterForget("1")
	assert.NotNil(t, err)
}

func TestRedis_ClusterGetKeysInSlot(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterGetKeysInSlot(1, 1)
	assert.NotNil(t, err)
	assert.Empty(t, slots)
}

func TestRedis_ClusterInfo(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	s, err := redis.ClusterInfo()
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
}

func TestRedis_ClusterKeySlot(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterKeySlot("godis")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), slots)
}

func TestRedis_ClusterMeet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterMeet("localhost", 8000)
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterNodes(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	s, err := redis.ClusterNodes()
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
	//t.Log(s)

	nodeID := s[:strings.Index(s, " ")]
	redis.ClusterSlaves(nodeID)
	//assert.Nil(t, err)
	//assert.NotEmpty(t, slaves)
}

func TestRedis_ClusterReplicate(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterReplicate("godis")
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterReset(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterReset(*ResetSoft)
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSaveConfig(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterSaveConfig()
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSetSlotImporting(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterSetSlotImporting(1, "godis")
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSetSlotMigrating(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterSetSlotMigrating(1, "godis")
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSetSlotNode(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterSetSlotNode(1, "godis")
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSetSlotStable(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	slots, err := redis.ClusterSetSlotStable(1)
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterSlots(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	s, err := redis.ClusterSlots()
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
}

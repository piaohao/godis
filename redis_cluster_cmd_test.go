package godis

import (
	"github.com/stretchr/testify/assert"
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
	//redis := NewRedis(option1)
	//defer redis.Close()
	//slots, err := redis.ClusterDelSlots(10000)
	//assert.Nil(t, err)
	//assert.Equal(t, "OK", slots)
}

func TestRedis_ClusterFailover(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	slots, err := redis.ClusterFailover()
	assert.NotNil(t, err)
	assert.Equal(t, "", slots)
}

func TestRedis_ClusterFlushSlots(t *testing.T) {
	//redis := NewRedis(option1)
	//defer redis.Close()
	//slots, err := redis.ClusterFlushSlots()
	//assert.Nil(t, err)
	//assert.Equal(t, "OK", slots)
}

func TestRedis_ClusterForget(t *testing.T) {
	redis := NewRedis(option1)
	defer redis.Close()
	_, err := redis.ClusterForget("1")
	assert.NotNil(t, err)
}

func TestRedis_ClusterGetKeysInSlot(t *testing.T) {
}

func TestRedis_ClusterInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ClusterInfo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(&Option{
				Host: "localhost",
				Port: 7000,
			})
			got, err := r.ClusterInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestRedis_ClusterKeySlot(t *testing.T) {
}

func TestRedis_ClusterMeet(t *testing.T) {
}

func TestRedis_ClusterNodes(t *testing.T) {
}

func TestRedis_ClusterReplicate(t *testing.T) {
}

func TestRedis_ClusterReset(t *testing.T) {
}

func TestRedis_ClusterSaveConfig(t *testing.T) {
}

func TestRedis_ClusterSetSlotImporting(t *testing.T) {
}

func TestRedis_ClusterSetSlotMigrating(t *testing.T) {
}

func TestRedis_ClusterSetSlotNode(t *testing.T) {
}

func TestRedis_ClusterSetSlotStable(t *testing.T) {
}

func TestRedis_ClusterSlaves(t *testing.T) {
}

func TestRedis_ClusterSlots(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ClusterSlots",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(&Option{
				Host: "localhost",
				Port: 7000,
			})
			got, err := r.ClusterSlots()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

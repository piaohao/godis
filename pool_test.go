package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPool_Basic(t *testing.T) {
	pool := NewPool(&PoolConfig{
		MaxTotal:             4,
		MaxIdle:              2,
		MinIdle:              2,
		MinEvictableIdleTime: 10,
		TestOnBorrow:         true,
	}, &Option{
		Host:              "localhost",
		Port:              6379,
		ConnectionTimeout: 2 * time.Second,
		SoTimeout:         2 * time.Second,
		Password:          "",
		Db:                0,
	})
	redis, e := pool.GetResource()
	assert.Nil(t, e)

	s, e := redis.Echo("godis")
	assert.Nil(t, e)
	assert.Equal(t, "godis", s)

	e = pool.returnResourceObject(redis)
	assert.Nil(t, e)
	s, e = redis.Echo("godis")
	assert.Nil(t, e)
	assert.Equal(t, "godis", s)

	//redis1, e := pool.GetResource()
	//assert.Nil(t, e)
	//e = pool.returnBrokenResourceObject(redis1)
	//assert.Nil(t, e)
	//s, e = redis1.Echo("godis")
	//assert.NotNil(t, e)
	//assert.Equal(t, "", s)

	redis2, e := pool.GetResource()
	assert.Nil(t, e)
	pool.Destroy()
	s, e = redis2.Echo("godis")
	assert.Nil(t, e)
	assert.Equal(t, "godis", s)

	_, e = pool.GetResource()
	assert.NotNil(t, e)
}

func TestPool_Basic2(t *testing.T) {
	pool := NewPool(&PoolConfig{
		MaxTotal:             4,
		MaxIdle:              2,
		MinIdle:              2,
		MinEvictableIdleTime: 10,
		TestOnBorrow:         true,
	}, &Option{
		Host:              "localhost",
		Port:              6379,
		ConnectionTimeout: 2 * time.Second,
		SoTimeout:         2 * time.Second,
		Password:          "123456",
		Db:                0,
	})
	_, e := pool.GetResource()
	assert.NotNil(t, e) //auth error
}

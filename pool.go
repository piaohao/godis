package godis

import (
	pool "github.com/jolestar/go-commons-pool"
	"time"
)

type Pool struct {
	internalPool *pool.ObjectPool
}

type PoolConfig struct {
	MaxTotal             *int
	MaxIdle              *int
	MinIdle              *int
	MinEvictableIdleTime *time.Duration
	TestOnBorrow         *bool
}

func NewPool(config PoolConfig, factory Factory) *Pool {
	poolConfig := pool.NewDefaultPoolConfig()
	if config.MaxTotal != nil {
		poolConfig.MaxTotal = *config.MaxTotal
	}
	if config.MaxTotal != nil {
		poolConfig.MaxIdle = *config.MaxIdle
	}
	if config.MaxTotal != nil {
		poolConfig.MinIdle = *config.MinIdle
	}
	if config.MaxTotal != nil {
		poolConfig.MinEvictableIdleTime = *config.MinEvictableIdleTime
	}
	if config.MaxTotal != nil {
		poolConfig.TestOnBorrow = *config.TestOnBorrow
	}
	return &Pool{
		internalPool: pool.NewObjectPool(nil, factory, poolConfig),
	}
}

func (p *Pool) GetResource() (*Redis, error) {
	obj, err := p.internalPool.BorrowObject(nil)
	if err != nil {
		return nil, err
	}
	return obj.(*Redis), nil
}

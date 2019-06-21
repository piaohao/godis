package godis

import (
	"context"
	"github.com/jolestar/go-commons-pool"
	"time"
)

//Pool
type Pool struct {
	internalPool *pool.ObjectPool
	ctx          context.Context
}

//PoolConfig
type PoolConfig struct {
	MaxTotal             int
	MaxIdle              int
	MinIdle              int
	MinEvictableIdleTime time.Duration
	TestOnBorrow         bool
}

//NewPool
func NewPool(config PoolConfig, factory *Factory) *Pool {
	poolConfig := pool.NewDefaultPoolConfig()
	if config.MaxTotal != 0 {
		poolConfig.MaxTotal = config.MaxTotal
	}
	if config.MaxIdle != 0 {
		poolConfig.MaxIdle = config.MaxIdle
	}
	if config.MaxIdle != 0 {
		poolConfig.MinIdle = config.MinIdle
	}
	if config.MinEvictableIdleTime != 0 {
		poolConfig.MinEvictableIdleTime = config.MinEvictableIdleTime
	}
	if config.TestOnBorrow != false {
		poolConfig.TestOnBorrow = config.TestOnBorrow
	}
	ctx := context.Background()
	return &Pool{
		ctx:          ctx,
		internalPool: pool.NewObjectPool(ctx, factory, poolConfig),
	}
}

//GetResource
func (p *Pool) GetResource() (*Redis, error) {
	obj, err := p.internalPool.BorrowObject(p.ctx)
	if err != nil {
		return nil, err
	}
	return obj.(*Redis), nil
}

//Destroy
func (p *Pool) Destroy() {
	p.internalPool.Close(p.ctx)
}

//Factory
type Factory struct {
	option Option
}

//NewFactory
func NewFactory(option Option) *Factory {
	return &Factory{option: option}
}

//MakeObject
func (f Factory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	redis := NewRedis(f.option)
	err := redis.Connect()
	if err != nil {
		return nil, err
	}
	return pool.NewPooledObject(redis), nil
}

//DestroyObject
func (f Factory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	_, err := redis.Quit()
	if err != nil {
		return err
	}
	return nil
}

//ValidateObject
func (f Factory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	redis := object.Object.(*Redis)
	if redis.client.host() != f.option.Host {
		return false
	}
	if redis.client.port() != f.option.Port {
		return false
	}
	reply, err := redis.Ping()
	if err != nil {
		return false
	}
	return reply == "PONG"
}

//ActivateObject
func (f Factory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	if redis.client.Db == f.option.Db {
		return nil
	}
	_, err := redis.Select(f.option.Db)
	if err != nil {
		return err
	}
	return nil
}

//PassivateObject
func (f Factory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// TODO maybe should select db 0? Not sure right now.
	return nil
}

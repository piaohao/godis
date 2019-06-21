package godis

import (
	"context"
	pool "github.com/jolestar/go-commons-pool"
	"time"
)

//Pool
type Pool struct {
	internalPool *pool.ObjectPool
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
	poolConfig.MaxTotal = config.MaxTotal
	poolConfig.MaxIdle = config.MaxIdle
	poolConfig.MinIdle = config.MinIdle
	poolConfig.MinEvictableIdleTime = config.MinEvictableIdleTime
	poolConfig.TestOnBorrow = config.TestOnBorrow
	return &Pool{
		internalPool: pool.NewObjectPool(nil, factory, poolConfig),
	}
}

//GetResource
func (p *Pool) GetResource() (*Redis, error) {
	obj, err := p.internalPool.BorrowObject(nil)
	if err != nil {
		return nil, err
	}
	return obj.(*Redis), nil
}

//Destroy
func (p *Pool) Destroy() {
	p.internalPool.Close(nil)
}

//Factory
type Factory struct {
	option Option
}

//NewFactory
func NewFactory(shardInfo Option) *Factory {
	return &Factory{option: shardInfo}
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

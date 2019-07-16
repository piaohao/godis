package godis

import (
	"context"
	"errors"
	"github.com/jolestar/go-commons-pool"
	"time"
)

var (
	//ErrClosed when pool is closed,continue operate pool will return this error
	ErrClosed = errors.New("pool is closed")
)

//Pool redis pool
type Pool struct {
	internalPool *pool.ObjectPool
	ctx          context.Context
}

//PoolConfig redis pool config, see go-commons-pool ObjectPoolConfig
type PoolConfig struct {
	MaxTotal int //The cap on the number of objects that can be allocated
	MaxIdle  int //The cap on the number of "idle" instances in the pool
	MinIdle  int //The minimum number of idle objects to maintain in the pool

	LIFO               bool //Whether the pool has LIFO (last in, first out) behaviour
	TestOnBorrow       bool //Whether objects borrowed from the pool will be validated before being returned from the ObjectPool.BorrowObject() method
	TestWhileIdle      bool //Whether objects sitting idle in the pool will be validated by the idle object evictor (if any - see TimeBetweenEvictionRuns )
	TestOnReturn       bool //Whether objects borrowed from the pool will be validated when they are returned to the pool via the ObjectPool.ReturnObject() method
	TestOnCreate       bool //Whether objects created for the pool will be validated before being returned from the ObjectPool.BorrowObject() method.
	BlockWhenExhausted bool //Whether to block when the ObjectPool.BorrowObject() method is invoked when the pool is exhausted

	MinEvictableIdleTime     time.Duration //The minimum amount of time an object may sit idle in the pool
	SoftMinEvictableIdleTime time.Duration //if MinEvictableIdleTime is positive, then SoftMinEvictableIdleTime is ignored
	TimeBetweenEvictionRuns  time.Duration //The amount of time sleep between runs of the idle object evictor goroutine.
	EvictionPolicyName       string        //The name of the EvictionPolicy implementation
	NumTestsPerEvictionRun   int           //The maximum number of objects to examine during each run
}

//NewPool create new pool
func NewPool(config *PoolConfig, option *Option) *Pool {
	poolConfig := pool.NewDefaultPoolConfig()
	if config != nil && config.MaxTotal != 0 {
		poolConfig.MaxTotal = config.MaxTotal
	}
	if config != nil && config.MaxIdle != 0 {
		poolConfig.MaxIdle = config.MaxIdle
	}
	if config != nil && config.MinIdle != 0 {
		poolConfig.MinIdle = config.MinIdle
	}
	if config != nil && config.MinEvictableIdleTime != 0 {
		poolConfig.MinEvictableIdleTime = config.MinEvictableIdleTime
	}
	if config != nil && config.TestOnBorrow != false {
		poolConfig.TestOnBorrow = config.TestOnBorrow
	}
	ctx := context.Background()
	return &Pool{
		ctx:          ctx,
		internalPool: pool.NewObjectPool(ctx, newFactory(option), poolConfig),
	}
}

//GetResource get redis instance from pool
func (p *Pool) GetResource() (*Redis, error) {
	obj, err := p.internalPool.BorrowObject(p.ctx)
	if err != nil {
		return nil, newConnectError(err.Error())
	}
	redis := obj.(*Redis)
	redis.client.socket.SetDeadline(time.Now().Add(redis.client.soTimeout))
	redis.setDataSource(p)
	return redis, nil
}

func (p *Pool) returnBrokenResourceObject(resource *Redis) error {
	if resource != nil {
		return p.internalPool.InvalidateObject(p.ctx, resource)
	}
	return nil
}

func (p *Pool) returnResourceObject(resource *Redis) error {
	if resource == nil {
		return nil
	}
	return p.internalPool.ReturnObject(p.ctx, resource)
}

//Destroy destroy pool
func (p *Pool) Destroy() {
	p.internalPool.Close(p.ctx)
}

//Factory redis pool factory
type factory struct {
	option *Option
}

//NewFactory create new redis pool factory
func newFactory(option *Option) *factory {
	return &factory{option: option}
}

//MakeObject make new object from pool
func (f factory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	redis := NewRedis(f.option)
	defer func() {
		if e := recover(); e != nil {
			redis.Close()
		}
	}()
	err := redis.Connect()
	if err != nil {
		return nil, err
	}
	return pool.NewPooledObject(redis), nil
}

//DestroyObject destroy object of pool
func (f factory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	_, err := redis.Quit()
	if err != nil {
		return err
	}
	return nil
}

//ValidateObject validate object is available
func (f factory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
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

//ActivateObject active object
func (f factory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
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

//PassivateObject passivate object
func (f factory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	//todo how to passivate redis object
	return nil
}

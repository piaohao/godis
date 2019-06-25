package godis

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrClosed = errors.New("pool is closed")
)

//PoolConfig
type PoolConfig struct {
	MaxTotal             int
	MaxIdle              int
	MinIdle              int
	MinEvictableIdleTime time.Duration
	TestOnBorrow         bool
}

type pool interface {
	Get() (*Redis, error)
	Put(redis *Redis) error
	Destroy() error
	Close(*Redis) error
}

type Pool struct {
	lock        sync.Mutex
	redisPool   chan *Redis
	maxTotal    int
	minIdle     int
	maxLifetime time.Duration
	create      func() (*Redis, error)
	activeCount int
}

//NewPool create new pool
func NewPool(config *PoolConfig, option *Option) *Pool {
	create := func() (*Redis, error) {
		redis := NewRedis(option)
		defer func() {
			if e := recover(); e != nil {
				redis.client.close()
			}
		}()
		err := redis.Connect()
		if err != nil {
			return nil, err
		}
		if option.Password != "" {
			_, err := redis.Auth(option.Password)
			if err != nil {
				return nil, err
			}
		}
		if option.Db != 0 {
			_, err := redis.Select(option.Db)
			if err != nil {
				return nil, err
			}
		}
		redis.activeTime = time.Now()
		return redis, nil
	}
	pool := &Pool{
		maxTotal:    10,
		minIdle:     3,
		maxLifetime: 30 * time.Second,
		create:      create,
	}
	if config != nil && config.MaxTotal != 0 {
		pool.maxTotal = config.MaxTotal
	}
	if config != nil && config.MinIdle != 0 {
		pool.minIdle = config.MinIdle
	}
	pool.redisPool = make(chan *Redis, pool.maxTotal)
	for i := 0; i < pool.minIdle; i++ {
		redis, err := create()
		if err != nil {
			continue
		}
		pool.redisPool <- redis
	}
	return pool
}

func (p *Pool) getPool() chan *Redis {
	p.lock.Lock()
	pool := p.redisPool
	p.lock.Unlock()
	return pool
}

func (p *Pool) Get() (*Redis, error) {
	pool := p.getPool()
	if pool == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case redis := <-p.getPool():
			if redis == nil {
				return nil, ErrClosed
			}
			if p.maxLifetime > 0 && redis.activeTime.Add(p.maxLifetime).Before(time.Now()) {
				p.Close(redis)
				continue
			}
			redis.setDataSource(p)
			return redis, nil
		default:
			p.lock.Lock()
			redis, err := p.create()
			p.lock.Unlock()
			if err != nil {
				return nil, err
			}

			return redis, nil
		}
	}
}

func (p *Pool) Put(redis *Redis) error {
	if redis == nil {
		return errors.New("redis is nil")
	}
	p.lock.Lock()
	if p.redisPool == nil {
		p.lock.Unlock()
		return p.Close(redis)
	}
	//newRedis, err := redis
	//if err != nil {
	//	p.lock.Unlock()
	//	return nil
	//}
	redis.activeTime = time.Now()
	select {
	case p.redisPool <- redis:
		p.lock.Unlock()
		return nil
	default:
		p.lock.Unlock()
		return p.Close(redis)
	}
}

func (p *Pool) Close(redis *Redis) error {
	if redis == nil {
		return errors.New("redis is nil")
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	return redis.client.close()
}

func (p *Pool) Destroy() error {
	p.lock.Lock()
	redisPool := p.redisPool
	p.redisPool = nil
	p.lock.Unlock()

	if redisPool == nil {
		return nil
	}
	close(redisPool)
	for redis := range redisPool {
		redis.client.close()
	}
	return nil
}

/*//Pool
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

//NewPool create new pool
func NewPool(config *PoolConfig, factory *Factory) *Pool {
	poolConfig := pool.NewDefaultPoolConfig()
	if config != nil && config.MaxTotal != 0 {
		poolConfig.MaxTotal = config.MaxTotal
	}
	if config != nil && config.MaxIdle != 0 {
		poolConfig.MaxIdle = config.MaxIdle
	}
	if config != nil && config.MaxIdle != 0 {
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
		internalPool: pool.NewObjectPool(ctx, factory, poolConfig),
	}
}

//GetResource get redis instance from pool
func (p *Pool) GetResource() (*Redis, error) {
	obj, err := p.internalPool.BorrowObject(p.ctx)
	if err != nil {
		return nil, NewConnectError(err.Error())
	}
	redis := obj.(*Redis)
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
type Factory struct {
	option *Option
}

//NewFactory create new redis pool factory
func NewFactory(option *Option) *Factory {
	return &Factory{option: option}
}

//MakeObject make new object from pool
func (f Factory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
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
	if f.option.Password != "" {
		_, err := redis.Auth(f.option.Password)
		if err != nil {
			return nil, err
		}
	}

	return pool.NewPooledObject(redis), nil
}

//DestroyObject destroy object of pool
func (f Factory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	_, err := redis.Quit()
	if err != nil {
		return err
	}
	return nil
}

//ValidateObject validate object is available
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

//ActivateObject active object
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

//PassivateObject ...
func (f Factory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// TODO maybe should select db 0? Not sure right now.
	return nil
}*/

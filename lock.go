package godis

import (
	"strconv"
	"sync/atomic"
	"time"
)

type Locker interface {
	TryLock(key string) (bool, error)
	UnLock() error
}

type locker struct {
	timeout time.Duration
	//retryCount int
	//retryDelay time.Duration

	ch    chan bool
	state int32

	key  string
	pool *Pool
}

func NewLocker(option *Option, lockOption *LockOption) *locker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	if lockOption.RetryCount <= 0 {
		lockOption.RetryCount = 0
	}
	if lockOption.RetryDelay.Nanoseconds() == 0 {
		lockOption.RetryDelay = 100 * time.Millisecond
	}
	pool := NewPool(&PoolConfig{MaxTotal: 500}, option)
	return &locker{
		timeout: lockOption.Timeout,
		//retryCount: lockOption.RetryCount,
		//retryDelay: lockOption.RetryDelay,
		ch:   make(chan bool, 1),
		pool: pool,
	}
}

type LockOption struct {
	Timeout    time.Duration
	RetryCount int
	RetryDelay time.Duration
}

var m int32 = 0
var n int32 = 0

func (l *locker) TryLock(key string) (bool, error) {
	l.key = key
	deadline := time.Now().Add(l.timeout)
	value := deadline.UnixNano()
	for {
		redis, err := l.pool.Get()
		//println(1)
		if err != nil {
			return false, err
		}
		status, err := redis.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
		//println(2)
		redis.Close()
		//get lock success
		if err == nil {
			//println(3)

			//println(4)
			if status == KeywordOk.Name {
				//fmt.Printf("m:%d\n", atomic.AddInt32(&m, 1))
				return true, nil
			}
		}

		//fmt.Printf("l.state:%d\n", atomic.AddInt32(&l.state, 1))
		atomic.AddInt32(&l.state, 1)
		select {
		case <-l.ch:
			atomic.AddInt32(&l.state, -1)
			continue
		case <-time.After(l.timeout):
			//fmt.Printf("n:%d\n", atomic.AddInt32(&n, 1))
			atomic.AddInt32(&l.state, -1)
			return false, nil
		}
	}
}

var i int32 = 0
var j int32 = 0

func (l *locker) UnLock() error {
	redis, err := l.pool.Get()
	if err != nil {
		return err
	}
	//fmt.Printf("i:%d\n", atomic.AddInt32(&i, 1))
	c, err := redis.Del(l.key)
	//fmt.Printf("j:%d\n", atomic.AddInt32(&j, 1))
	//fmt.Printf("state:%d\n", atomic.LoadInt32(&l.state))
	redis.Close()
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}
	if atomic.LoadInt32(&l.state) > 0 {
		//fmt.Printf("j:%d\n", atomic.AddInt32(&j, 1))
		l.ch <- true
	}
	return nil
}

type clusterLocker struct {
	timeout time.Duration
	//retryCount int
	//retryDelay time.Duration

	ch    chan bool
	state int32

	key          string
	redisCluster *RedisCluster
}

func NewClusterLocker(option *ClusterOption, lockOption *LockOption) *clusterLocker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	if lockOption.RetryCount <= 0 {
		lockOption.RetryCount = 0
	}
	if lockOption.RetryDelay.Nanoseconds() == 0 {
		lockOption.RetryDelay = 500 * time.Millisecond
	}
	return &clusterLocker{
		timeout: lockOption.Timeout,
		//retryCount: lockOption.RetryCount,
		//retryDelay: lockOption.RetryDelay,
		ch:           make(chan bool, 1),
		redisCluster: NewRedisCluster(option),
	}
}

func (l *clusterLocker) TryLock(key string) (bool, error) {
	l.key = key
	deadline := time.Now().Add(l.timeout)
	value := deadline.UnixNano()
	for {
		status, err := l.redisCluster.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
		//get lock success
		if err == nil && status == KeywordOk.Name {
			return true, nil
		}
		atomic.AddInt32(&l.state, 1)
		select {
		case <-l.ch:
			atomic.AddInt32(&l.state, -1)
			continue
		case <-time.After(l.timeout):
			atomic.AddInt32(&l.state, -1)
			return false, nil
		}
	}
}

func (l *clusterLocker) UnLock() error {
	c, err := l.redisCluster.Del(l.key)
	if c == 0 {
		return nil
	}
	if atomic.LoadInt32(&l.state) > 0 {
		l.ch <- true
	}
	return err
}

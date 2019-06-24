package godis

import (
	"fmt"
	"strconv"
	"sync"
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

	ch   chan bool
	lock sync.Mutex

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
		lockOption.RetryDelay = 200 * time.Millisecond
	}
	pool := NewPool(nil, NewFactory(option))
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

func (l *locker) TryLock(key string) (bool, error) {
	l.key = key
	deadline := time.Now().Add(l.timeout)
	value := deadline.UnixNano()
	for {
		redis, err := l.pool.GetResource()
		if err != nil {
			return false, err
		}
		status, err := redis.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
		//get lock success
		redis.Close()
		if err == nil && status == KeywordOk.Name {
			return true, nil
		}
		var ch chan bool
		l.lock.Lock()
		ch = l.ch
		l.lock.Unlock()
		elapsed := time.Until(deadline)
		if elapsed <= 0 {
			return false, nil
		}
		select {
		case <-ch:
			continue
		case <-time.After(elapsed):
			return false, nil
		}
	}
}

func (l *locker) UnLock() error {
	redis, err := l.pool.GetResource()
	if err != nil {
		return err
	}
	_, err = redis.Del(l.key)
	redis.Close()
	if err != nil {
		return err
	}
	newCh := make(chan bool, 1)
	l.lock.Lock()
	ch := l.ch
	l.ch = newCh
	l.lock.Unlock()
	close(ch)
	return nil
}

type clusterLocker struct {
	timeout time.Duration
	//retryCount int
	//retryDelay time.Duration

	ch     chan bool
	writer int32

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
		if err == nil {
			if status == KeywordOk.Name {
				return true, nil
			}
		} else {
			println(err)
		}
		elapsed := time.Until(deadline)
		if elapsed <= 0 {
			return false, nil
		}
		println(atomic.AddInt32(&l.writer, 1))
		select {
		case <-l.ch:
			atomic.AddInt32(&l.writer, -1)
			continue
		case <-time.After(l.timeout):
			return false, nil
		}
	}
}

func (l *clusterLocker) UnLock() error {
	c, err := l.redisCluster.Del(l.key)
	if c == 0 {
		return nil
	}
	fmt.Printf("delete success,writer:%d\n", atomic.LoadInt32(&l.writer))
	if atomic.LoadInt32(&l.writer) > 0 {
		l.ch <- true
	}
	return err
}

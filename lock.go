package godis

import (
	"errors"
	"strconv"
	"time"
)

var LockTimeoutErr = errors.New("get lock timeout")

type lock struct {
	Name string
}

type locker struct {
	timeout time.Duration
	ch      chan bool
	pool    *Pool
}

// create new locker
func NewLocker(option *Option, lockOption *LockOption) *locker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	pool := NewPool(&PoolConfig{MaxTotal: 500}, option)
	return &locker{
		timeout: lockOption.Timeout,
		ch:      make(chan bool, 1),
		pool:    pool,
	}
}

// locker options
type LockOption struct {
	Timeout time.Duration //lock wait timeout
}

// acquire a lock,when it returns a non nil locker,get lock success,
// otherwise, it returns an error,get lock failed
func (l *locker) TryLock(key string) (*lock, error) {
	deadline := time.Now().Add(l.timeout)
	value := strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		redis, err := l.pool.GetResource()
		if err != nil {
			return nil, err
		}
		if time.Now().After(deadline) {
			return nil, LockTimeoutErr
		}
		status, err := redis.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
		redis.Close()
		if err == nil {
			if status == KeywordOk.Name {
				return &lock{Name: key}, nil
			}
		}
		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return nil, LockTimeoutErr
		}
	}
}

// when your business end,then release the locker
func (l *locker) UnLock(lock *lock) error {
	redis, err := l.pool.GetResource()
	if err != nil {
		return err
	}
	defer redis.Close()
	l.ch <- true
	c, err := redis.Del(lock.Name)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}
	return nil
}

type clusterLocker struct {
	timeout      time.Duration
	ch           chan int
	redisCluster *RedisCluster
}

// create new cluster locker
func NewClusterLocker(option *ClusterOption, lockOption *LockOption) *clusterLocker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	return &clusterLocker{
		timeout:      lockOption.Timeout,
		ch:           make(chan int, 1),
		redisCluster: NewRedisCluster(option),
	}
}

// acquire a lock,when it returns a non nil locker,get lock success,
// otherwise, it returns an error,get lock failed
func (l *clusterLocker) TryLock(key string) (*lock, error) {
	deadline := time.Now().Add(l.timeout)
	value := strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		if time.Now().After(deadline) {
			return nil, LockTimeoutErr
		}
		if len(l.ch) == 0 {
			status, err := l.redisCluster.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
			//get lock success
			if err == nil && status == KeywordOk.Name {
				return &lock{Name: key}, nil
			}
		}
		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return nil, LockTimeoutErr
		}
	}
}

// when your business end,then release the locker
func (l *clusterLocker) UnLock(lock *lock) error {
	l.ch <- 1
	c, err := l.redisCluster.Del(lock.Name)
	if c == 0 {
		return nil
	}
	return err
}

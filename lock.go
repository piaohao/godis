package godis

import (
	"errors"
	"strconv"
	"time"
)

//ErrLockTimeOut when get lock exceed the timeout,then return error
var ErrLockTimeOut = errors.New("get lock timeout")

//Lock different keys with different lock
type Lock struct {
	name string
}

//Locker the lock client
type Locker struct {
	timeout time.Duration
	ch      chan bool
	pool    *Pool
}

//NewLocker create new locker
func NewLocker(option *Option, lockOption *LockOption) *Locker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	pool := NewPool(&PoolConfig{MaxTotal: 500}, option)
	return &Locker{
		timeout: lockOption.Timeout,
		ch:      make(chan bool, 1),
		pool:    pool,
	}
}

//LockOption locker options
type LockOption struct {
	Timeout time.Duration //lock wait timeout
}

//TryLock acquire a lock,when it returns a non nil locker,get lock success,
// otherwise, it returns an error,get lock failed
func (l *Locker) TryLock(key string) (*Lock, error) {
	deadline := time.Now().Add(l.timeout)
	value := strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		redis, err := l.pool.GetResource()
		if err != nil {
			return nil, err
		}
		if time.Now().After(deadline) {
			return nil, ErrLockTimeOut
		}
		status, err := redis.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
		redis.Close()
		if err == nil && status == keywordOk.name {
			if len(l.ch) > 0 {
				<-l.ch
			}
			return &Lock{name: key}, nil
		}
		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return nil, ErrLockTimeOut
		}
	}
}

//UnLock when your business end,then release the locker
func (l *Locker) UnLock(lock *Lock) error {
	redis, err := l.pool.GetResource()
	if err != nil {
		return err
	}
	defer redis.Close()
	if len(l.ch) == 0 {
		l.ch <- true
	}
	c, err := redis.Del(lock.name)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}
	return nil
}

//ClusterLocker cluster lock client
type ClusterLocker struct {
	timeout      time.Duration
	ch           chan bool
	redisCluster *RedisCluster
}

//NewClusterLocker create new cluster locker
func NewClusterLocker(option *ClusterOption, lockOption *LockOption) *ClusterLocker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	return &ClusterLocker{
		timeout:      lockOption.Timeout,
		ch:           make(chan bool, 1),
		redisCluster: NewRedisCluster(option),
	}
}

//TryLock acquire a lock,when it returns a non nil locker,get lock success,
// otherwise, it returns an error,get lock failed
func (l *ClusterLocker) TryLock(key string) (*Lock, error) {
	deadline := time.Now().Add(l.timeout)
	value := strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		if time.Now().After(deadline) {
			return nil, ErrLockTimeOut
		}
		if len(l.ch) == 0 {
			status, err := l.redisCluster.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
			//get lock success
			if err == nil && status == keywordOk.name {
				if len(l.ch) > 0 {
					<-l.ch
				}
				return &Lock{name: key}, nil
			}
		}
		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return nil, ErrLockTimeOut
		}
	}
}

//UnLock when your business end,then release the locker
func (l *ClusterLocker) UnLock(lock *Lock) error {
	if len(l.ch) == 0 {
		l.ch <- true
	}
	c, err := l.redisCluster.Del(lock.name)
	if c == 0 {
		return nil
	}
	return err
}

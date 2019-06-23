package godis

import (
	"fmt"
	"strconv"
	"time"
)

type Locker interface {
	TryLock(key string) (bool, error)
	UnLock() error
}

type locker struct {
	redis      *Redis
	timeout    time.Duration
	retryCount int
	retryDelay time.Duration

	key string
}

func newLocker(redis *Redis, option *LockOption) *locker {
	if option == nil {
		option = &LockOption{}
	}
	if option.Timeout.Nanoseconds() == 0 {
		option.Timeout = 5 * time.Second
	}
	if option.RetryCount <= 0 {
		option.RetryCount = 0
	}
	if option.RetryDelay.Nanoseconds() == 0 {
		option.RetryDelay = 200 * time.Millisecond
	}
	return &locker{
		redis:      redis,
		timeout:    option.Timeout,
		retryCount: option.RetryCount,
		retryDelay: option.RetryDelay,
	}
}

type LockOption struct {
	Timeout    time.Duration
	RetryCount int
	RetryDelay time.Duration
}

func (l *locker) TryLock(key string) (bool, error) {
	l.key = key
	value := time.Now().Add(l.timeout).UnixNano()
	status, err := l.redis.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
	//get lock success
	if err == nil && status == KeywordOk.Name {
		return true, nil
	}
	if l.retryCount <= 0 {
		ch := make(chan bool, 1)
		go func() {
			for {
				time.Sleep(l.retryDelay)
				value := time.Now().Add(l.timeout).UnixNano()
				status, err := l.redis.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
				//get lock success
				if err == nil && status == KeywordOk.Name {
					ch <- true
				}
			}
		}()
		select {
		case <-ch:
			return true, nil
		case <-time.After(l.timeout):
			return false, NewRedisError(fmt.Sprintf("get lock failed after %s", l.timeout))
		}
	} else {
		ch := make(chan bool, 1)
		go func() {
			//get lock failed,retry
			for i := 0; i < l.retryCount; i++ {
				time.Sleep(l.retryDelay)
				value := time.Now().Add(l.timeout).UnixNano()
				status, err := l.redis.SetWithParamsAndTime(key, strconv.FormatInt(value, 10), "nx", "px", l.timeout.Nanoseconds()/1e6)
				//get lock success
				if err == nil && status == KeywordOk.Name {
					ch <- true
				}
			}
		}()
		select {
		case <-ch:
			return true, nil
		case <-time.After(l.timeout):
			return false, NewRedisError(fmt.Sprintf("get lock failed after %s", l.timeout))
		}
	}
}

func (l *locker) UnLock() error {
	_, err := l.redis.Del(l.key)
	if err != nil {
		return err
	}
	return nil
}

type clusterLocker struct {
	redisCluster *RedisCluster
	timeout      time.Duration
	retryCount   int
	retryDelay   time.Duration

	key string
}

func newClusterLocker(redisCluster *RedisCluster, option *LockOption) *clusterLocker {
	if option == nil {
		option = &LockOption{}
	}
	if option.Timeout.Nanoseconds() == 0 {
		option.Timeout = 5 * time.Second
	}
	if option.RetryCount <= 0 {
		option.RetryCount = 0
	}
	if option.RetryDelay.Nanoseconds() == 0 {
		option.RetryDelay = 500 * time.Millisecond
	}
	return &clusterLocker{
		redisCluster: redisCluster,
		timeout:      option.Timeout,
		retryCount:   option.RetryCount,
		retryDelay:   option.RetryDelay,
	}
}

func (l *clusterLocker) TryLock(key string) (bool, error) {
	l.key = key
	command := newRedisClusterCommand(l.redisCluster.MaxAttempts, l.redisCluster.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		locker := newLocker(redis, &LockOption{
			Timeout:    l.timeout,
			RetryCount: l.retryCount,
			RetryDelay: l.retryDelay,
		})
		return locker.TryLock(key)
	}
	reply, err := command.run(key)
	if err != nil {
		return false, err
	}
	return reply.(bool), nil
}

func (l *clusterLocker) UnLock() error {
	command := newRedisClusterCommand(l.redisCluster.MaxAttempts, l.redisCluster.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Del(l.key)
	}
	_, err := command.run(l.key)
	return err
}

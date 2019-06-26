package godis

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var LockTimeoutErr = errors.New("get lock timeout")

func GoID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id, nil
}

type Locker interface {
	TryLock(key string) (bool, error)
	UnLock() error
}

type locker struct {
	timeout time.Duration

	ch    chan bool
	state int32

	key  string
	pool *Pool

	vMap map[int]string
}

func NewLocker(option *Option, lockOption *LockOption) *locker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	//if lockOption.RetryCount <= 0 {
	//	lockOption.RetryCount = 0
	//}
	//if lockOption.RetryDelay.Nanoseconds() == 0 {
	//	lockOption.RetryDelay = 100 * time.Millisecond
	//}
	pool := NewPool(&PoolConfig{MaxTotal: 500}, option)
	return &locker{
		timeout: lockOption.Timeout,
		ch:      make(chan bool, 1),
		pool:    pool,
		vMap:    make(map[int]string),
	}
}

type LockOption struct {
	Timeout time.Duration
	//RetryCount int
	//RetryDelay time.Duration
}

func (l *locker) TryLock(key string) (bool, error) {
	l.key = key
	deadline := time.Now().Add(l.timeout)
	id, err := GoID()
	if err != nil {
		return false, err
	}
	value := strconv.FormatInt(int64(id), 10) + "-" + strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		redis, err := l.pool.Get()
		if err != nil {
			return false, err
		}
		if time.Now().After(deadline) {
			return false, LockTimeoutErr
		}
		status, err := redis.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
		redis.Close()
		if err == nil {
			if status == KeywordOk.Name {
				l.vMap[id] = value
				return true, nil
			}
		}

		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return false, LockTimeoutErr
		}
	}
}

func (l *locker) UnLock() error {
	redis, err := l.pool.Get()
	if err != nil {
		return err
	}
	defer redis.Close()
	v, err := redis.Get(l.key)
	if err != nil {
		return err
	}
	arr := strings.Split(v, "-")
	if len(arr) < 1 {
		return nil
	}
	goid, _ := strconv.Atoi(arr[0])
	if l.vMap[goid] != v {
		return nil
	}
	l.ch <- true
	c, err := redis.Del(l.key)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}
	return nil
}

type clusterLocker struct {
	timeout time.Duration

	ch    chan int
	state int32

	key          string
	redisCluster *RedisCluster

	vMap map[int]string
}

func NewClusterLocker(option *ClusterOption, lockOption *LockOption) *clusterLocker {
	if lockOption == nil {
		lockOption = &LockOption{}
	}
	if lockOption.Timeout.Nanoseconds() == 0 {
		lockOption.Timeout = 5 * time.Second
	}
	//if lockOption.RetryCount <= 0 {
	//	lockOption.RetryCount = 0
	//}
	//if lockOption.RetryDelay.Nanoseconds() == 0 {
	//	lockOption.RetryDelay = 500 * time.Millisecond
	//}
	return &clusterLocker{
		timeout:      lockOption.Timeout,
		ch:           make(chan int, 1),
		redisCluster: NewRedisCluster(option),
		vMap:         make(map[int]string),
	}
}

var inCh int32 = 0
var outCh int32 = 0

func (l *clusterLocker) TryLock(key string) (bool, error) {
	l.key = key
	deadline := time.Now().Add(l.timeout)
	id, err := GoID()
	if err != nil {
		return false, err
	}
	value := strconv.FormatInt(int64(id), 10) + "-" + strconv.FormatInt(deadline.UnixNano(), 10)
	for {
		if time.Now().After(deadline) {
			return false, LockTimeoutErr
		}
		if len(l.ch) == 0 {
			status, err := l.redisCluster.SetWithParamsAndTime(key, value, "nx", "px", l.timeout.Nanoseconds()/1e6)
			//get lock success
			if err == nil && status == KeywordOk.Name {
				l.vMap[id] = value
				return true, nil
			}
		}
		select {
		case <-l.ch:
			continue
		case <-time.After(l.timeout):
			return false, LockTimeoutErr
		}
	}
}

func (l *clusterLocker) UnLock() error {
	v, err := l.redisCluster.Get(l.key)
	if err != nil {
		return err
	}
	arr := strings.Split(v, "-")
	if len(arr) < 1 {
		return nil
	}
	goid, _ := strconv.Atoi(arr[0])
	if l.vMap[goid] != v {
		return nil
	}
	l.ch <- 1
	c, err := l.redisCluster.Del(l.key)
	if c == 0 {
		return nil
	}
	return err
}

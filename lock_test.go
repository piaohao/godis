package godis

import (
	"sync"
	"testing"
	"time"
)

func TestRedisCluster_Lock(t *testing.T) {
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 4)
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
			cluster := newCluster()
			locker := newClusterLocker(cluster, &LockOption{
				Timeout: 5 * time.Second,
			})
			//start := time.Now()
			ok, err := locker.TryLock("lock")
			//t.Logf("cost time:%s", time.Now().Sub(start))
			if err == nil && ok {
				count++
			}
			//start = time.Now()
			locker.UnLock()
			//t.Logf("cost time:%s", time.Now().Sub(start))
			<-ch
		}()
	}
	group.Wait()
	t.Log(count)
}

//ignore this case,cause race data
func _TestRedis_NoLock(t *testing.T) {
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
			count++
			<-ch
		}()
	}
	group.Wait()
	t.Log(count)
}

func TestRedis_Lock(t *testing.T) {
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 4)
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
			redis := NewRedis(option)
			locker := newLocker(redis, nil)
			ok, err := locker.TryLock("lock")
			if err == nil && ok {
				count++
			}
			locker.UnLock()
			redis.Close()
			<-ch
		}()
	}
	group.Wait()
	t.Log(count)
	if count != 1000 {
		t.Errorf("want 1000,but %d", count)
	}
}

package godis

import (
	"sync"
	"testing"
)

func TestRedisCluster_Lock(t *testing.T) {
	count := 0
	var group sync.WaitGroup
	locker := NewClusterLocker(clusterOption, nil)
	ch := make(chan bool, 2)
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
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
	if count != 100 {
		t.Errorf("want 100,but %d", count)
	}
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
	locker := NewLocker(option, nil)
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 4)
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			ch <- true
			ok, err := locker.TryLock("lock")
			if err == nil && ok {
				count++
			}
			locker.UnLock()
			<-ch
		}()
	}
	group.Wait()
	t.Log(count)
	if count != 100 {
		t.Errorf("want 100,but %d", count)
	}
}

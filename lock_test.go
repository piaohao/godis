package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRedisCluster_Lock(t *testing.T) {
	//os.Remove("count.txt")
	count := 0
	var group sync.WaitGroup
	locker := NewClusterLocker(clusterOption, nil)
	ch := make(chan bool, 4)
	total := 10000
	var timeoutCount int32 = 0
	for i := 0; i < total; i++ {
		group.Add(1)
		ch <- true
		//println(i)
		go func() {
			defer group.Done()
			lock, err := locker.TryLock("lock")
			if err == nil {
				if lock != nil {
					count++
					/*file, _ := os.OpenFile(
						"count.txt",
						os.O_RDWR|os.O_CREATE,
						0664,
					)
					arr, err := ioutil.ReadFile("count.txt")
					if err != nil {
						fmt.Printf("%v\n", err)
					}
					oldNum := 0
					if string(arr) != "" {
						oldNum, _ = strconv.Atoi(string(arr))
					}
					file.WriteString(strconv.Itoa(oldNum + 1))
					file.Close()
					if chCount := len(locker.ch); chCount > 0 {
						fmt.Printf("locker ch:%d\n", chCount)
					}*/
					locker.UnLock(lock)
				} else {
					atomic.AddInt32(&timeoutCount, 1)
				}
			} else {
				atomic.AddInt32(&timeoutCount, 1)
			}
			<-ch
		}()
	}
	group.Wait()
	realCount := count + int(timeoutCount)
	assert.Equal(t, total, realCount, "want %d,but %d", total, realCount)
}

func TestRedis_Lock(t *testing.T) {
	locker := NewLocker(option, &LockOption{Timeout: 1 * time.Second})
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	total := 10000
	var timeoutCount int32 = 0
	for i := 0; i < total; i++ {
		group.Add(1)
		ch <- true
		go func() {
			defer group.Done()
			lock, err := locker.TryLock("lock")
			if err == nil {
				if lock != nil {
					count++
					locker.UnLock(lock)
				} else {
					atomic.AddInt32(&timeoutCount, 1)
				}
			} else {
				fmt.Printf("%v\n", err)
				atomic.AddInt32(&timeoutCount, 1)
			}
			<-ch
		}()
	}
	group.Wait()
	realCount := count + int(timeoutCount)
	assert.Equal(t, total, realCount, "want %d,but %d", total, realCount)
}

func TestRedis_Lock_Exception(t *testing.T) {
	locker := NewLocker(option, &LockOption{Timeout: 1 * time.Second})
	count := 0
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	total := 20
	var timeoutCount int32 = 0
	for i := 0; i < total; i++ {
		group.Add(1)
		ch <- true
		go func() {
			defer group.Done()
			lock, err := locker.TryLock("lock")
			if err == nil {
				if lock != nil {
					count++
					//simulate business exception
					rand.NewSource(time.Now().UnixNano())
					if rand.Intn(10) > 4 {
						<-ch
						return
					}
					locker.UnLock(lock)
				} else {
					atomic.AddInt32(&timeoutCount, 1)
				}
			} else {
				fmt.Printf("%v\n", err)
				atomic.AddInt32(&timeoutCount, 1)
			}
			<-ch
		}()
	}
	group.Wait()
	realCount := count + int(timeoutCount)
	assert.Equal(t, total, realCount, "want %d,but %d", total, realCount)
}

func _BenchmarkRedis_Lock(b *testing.B) {
	locker := NewLocker(option, &LockOption{Timeout: 3 * time.Second})
	count := 0
	for i := 0; i < 100; i++ {
		lock, err := locker.TryLock("lock")
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if lock != nil {
			count++
			fmt.Printf("%d\n", count)
			locker.UnLock(lock)
		}
	}
	b.Log(count)
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

func TestCluster_Set(t *testing.T) {
	cluster := NewRedisCluster(clusterOption)
	var count int32 = 0
	var group sync.WaitGroup
	ch := make(chan bool, 8)
	for i := 0; i < 100000; i++ {
		group.Add(1)
		ch <- true
		go func() {
			defer group.Done()
			atomic.AddInt32(&count, 1)
			reply, err := cluster.SetWithParamsAndTime("lock", "1", "nx", "px", 5*time.Second.Nanoseconds()/1e6)
			if err != nil {
				fmt.Printf("err:%v\n", err)
			} else {
				if reply == "OK" {
					fmt.Printf("reply:%s\n", reply)
				}
			}
			<-ch
		}()
	}
	group.Wait()
	t.Log(count)
}

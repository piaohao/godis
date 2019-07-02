# godis
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/piaohao/godis)
[![Build Status](https://travis-ci.com/piaohao/godis.svg?branch=dev.master)](https://travis-ci.com/piaohao/godis) 
[![Go Report](https://goreportcard.com/badge/github.com/piaohao/godis?123)](https://goreportcard.com/report/github.com/piaohao/godis) 
[![codecov](https://codecov.io/gh/piaohao/godis/branch/master/graph/badge.svg)](https://codecov.io/gh/piaohao/godis)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/piaohao/godis)

godis是一个golang实现的redis客户端,参考jedis实现.  
godis实现了几乎所有的redis命令,包括单机命令,集群命令,管道命令和事物命令等.  
如果你用过jedis,你就能非常容易地上手godis,因为godis的方法命名几乎全部来自jedis.  
值得一提的是,godis实现了单机和集群模式下的分布式锁,godis的锁比redisson快很多,在i7,8核32g的电脑测试,10万次for循环,8个线程,业务逻辑是简单的count++,redisson需要18-20秒,而godis只需要7秒左右.  
godis已经完成了大多数命令的测试用例,比较稳定.  
非常高兴你能提出任何建议,我会积极地迭代这个项目.  

* [github地址: https://github.com/piaohao/godis](https://github.com/piaohao/godis)
* [gitee地址: https://gitee.com/piaohao/godis](https://gitee.com/piaohao/godis) 

# 特色
* cluster集群
* pipeline管道
* transaction事物
* distributed lock分布式锁
* 其他功能在持续开发中

# 安装
```
go get -u github.com/piaohao/godis
```
或者使用 `go.mod`:
```
require github.com/piaohao/godis latest
```
# 文档

* [ApiDoc](https://godoc.org/github.com/piaohao/godis)

# 快速开始
1. 基本例子

    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
    )
    
    func main() {
        redis := godis.NewRedis(&godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        defer redis.Close()
        redis.Set("godis", "1")
        arr, _ := redis.Get("godis")
        println(arr)
    }
    ```
1. 使用连接池
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
    )
    
    func main() {
        option:=&godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        }
        pool := godis.NewPool(&godis.PoolConfig{}, option)
        redis, _ := pool.GetResource()
        defer redis.Close()
        redis.Set("godis", "1")
        arr, _ := redis.Get("godis")
        println(arr)
    }
    ```
1. 发布订阅
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        option:=&godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        }
        pool := godis.NewPool(&godis.PoolConfig{}, option)
        go func() {
            redis, _ := pool.GetResource()
            defer redis.Close()
            pubsub := &godis.RedisPubSub{
                OnMessage: func(channel, message string) {
                    println(channel, message)
                },
                OnSubscribe: func(channel string, subscribedChannels int) {
                    println(channel, subscribedChannels)
                },
                OnPong: func(channel string) {
                    println("recieve pong")
                },
            }
            redis.Subscribe(pubsub, "godis")
        }()
        time.Sleep(1 * time.Second)
        {
            redis, _ := pool.GetResource()
            defer redis.Close()
            redis.Publish("godis", "godis pubsub")
            redis.Close()
        }
        time.Sleep(1 * time.Second)
    }
    ```
1. cluster集群
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        cluster := godis.NewRedisCluster(&godis.ClusterOption{
            Nodes:             []string{"localhost:7000", "localhost:7001", "localhost:7002", "localhost:7003", "localhost:7004", "localhost:7005"},
            ConnectionTimeout: 0,
            SoTimeout:         0,
            MaxAttempts:       0,
            Password:          "",
            PoolConfig:        &godis.PoolConfig{},
        })
        cluster.Set("cluster", "godis cluster")
        reply, _ := cluster.Get("cluster")
        println(reply)
    }
    ```
1. pipeline管道
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        option:=&godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        }
        pool := godis.NewPool(&godis.PoolConfig{}, option)
        redis, _ := pool.GetResource()
        defer redis.Close()
        p := redis.Pipelined()
        infoResp, _ := p.Info()
        timeResp, _ := p.Time()
        p.Sync()
        timeList, _ := timeResp.Get()
        println(timeList)
        info, _ := infoResp.Get()
        println(info)
    }
    ```
1. transaction事物
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        option:=&godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        }
        pool := godis.NewPool(nil, option)
        redis, _ := pool.GetResource()
        defer redis.Close()
        p, _ := redis.Multi()
        infoResp, _ := p.Info()
        timeResp, _ := p.Time()
        p.Exec()
        timeList, _ := timeResp.Get()
        println(timeList)
        info, _ := infoResp.Get()
        println(info)
    }
    ``` 
1. distribute lock分布式锁
    * single redis   
         ```go
            package main
            
            import (
                "github.com/piaohao/godis"
                "time"
            )
            
            func main() {
                locker := godis.NewLocker(&godis.Option{
                      Host: "localhost",
                      Port: 6379,
                      Db:   0,
                  }, &godis.LockOption{
                      Timeout: 5*time.Second,
                  })
                lock, err := locker.TryLock("lock")
                if err == nil && lock!=nil {
                    //do something
                    locker.UnLock(lock)
                }
                
            }
        ``` 
    * redis cluster   
         ```go
            package main
            
            import (
                "github.com/piaohao/godis"
                "time"
            )
            
            func main() {
                locker := godis.NewClusterLocker(&godis.ClusterOption{
                	Nodes:             []string{"localhost:7000", "localhost:7001", "localhost:7002", "localhost:7003", "localhost:7004", "localhost:7005"},
                    ConnectionTimeout: 0,
                    SoTimeout:         0,
                    MaxAttempts:       0,
                    Password:          "",
                    PoolConfig:        &godis.PoolConfig{},
                },&godis.LockOption{
                    Timeout: 5*time.Second,
                })
                lock, err := locker.TryLock("lock")
                if err == nil && lock!=nil {
                    //do something
                    locker.UnLock(lock)
                }
            }
        ```   
# 证书

`godis` 使用的是 [MIT License](LICENSE), 永远100%免费和开源.      

# 鸣谢
* [jedis,java非常出名的redis客户端](https://github.com/xetorthio/jedis)
* [gf,功能非常强大的go web框架](https://github.com/gogf/gf)
* [go-commons-pool,参照apache common-pool实现的go连接池](https://github.com/jolestar/go-commons-pool)

# 联系

piao.hao@qq.com
     
# godis
<div align=center>

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/piaohao/godis)
[![Build Status](https://travis-ci.com/piaohao/godis.svg?branch=dev.master)](https://travis-ci.com/piaohao/godis) 
[![Go Report](https://goreportcard.com/badge/github.com/piaohao/godis?123)](https://goreportcard.com/report/github.com/piaohao/godis) 
[![codecov](https://codecov.io/gh/piaohao/godis/branch/master/graph/badge.svg)](https://codecov.io/gh/piaohao/godis)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/piaohao/godis)


</div>

redis client implement by golang, refer to jedis.

# Features
* cluster
* pipeline
* transaction
* distributed lock
* other feature under development

# Installation
```
go get -u github.com/piaohao/godis
```
or use `go.mod`:
```
require github.com/piaohao/godis latest
```
# Documentation

* [ApiDoc](https://godoc.org/github.com/piaohao/godis)

# Quick Start
1. basic example

    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
    )
    
    func main() {
        redis := godis.NewRedis(godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        defer redis.Close()
        redis.Set("godis", "1")
        arr, _ := redis.Get("godis")
        println(string(arr))
    }
    ```
1. use pool
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
    )
    
    func main() {
        factory := godis.NewFactory(godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        pool := godis.NewPool(godis.PoolConfig{}, factory)
        redis, _ := pool.GetResource()
        defer redis.Close()
        redis.Set("godis", "1")
        arr, _ := redis.Get("godis")
        println(string(arr))
    }
    ```
1. pubsub
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        factory := godis.NewFactory(godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        pool := godis.NewPool(godis.PoolConfig{}, factory)
        go func() {
            redis, _ := pool.GetResource()
            defer redis.Close()
            pubsub := &godis.RedisPubSub{
                Redis: redis,
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
1. cluster
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        cluster := godis.NewRedisCluster([]string{"192.168.1.6:8001", "192.168.1.6:8002", "192.168.1.6:8003", "192.168.1.6:8004", "192.168.1.6:8005", "192.168.1.6:8006"},
        	0, 0, 1, "", godis.PoolConfig{})
        cluster.Set("cluster", "godis cluster")
        reply, _ := cluster.Get("cluster")
        println(reply)
    }
    ```
1. pipeline
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        factory := godis.NewFactory(godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        pool := godis.NewPool(godis.PoolConfig{}, factory)
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
1. transaction
    ```go
    package main
    
    import (
        "github.com/piaohao/godis"
        "time"
    )
    
    func main() {
        factory := godis.NewFactory(godis.Option{
            Host: "localhost",
            Port: 6379,
            Db:   0,
        })
        pool := godis.NewPool(godis.PoolConfig{}, factory)
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
# License

`godis` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.      

# Thanks
* [jedis](https://github.com/xetorthio/jedis)
* [gf](https://github.com/gogf/gf)

# Contact

piao.hao@qq.com
     
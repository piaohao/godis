package godis_test

import (
	"github.com/gogf/gf/g/test/gtest"
	"github.com/piaohao/godis"
	"testing"
)

func TestPipeline_Basic(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(godis.ShardInfo{
			Host: "localhost",
			Port: 6379,
			Db:   0,
		})
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, _ := pool.GetResource()
		defer redis.Close()
		p := redis.Pipelined()
		infoResp, err := p.Info()
		gtest.Assert(err, nil)
		_, err = infoResp.Get()
		gtest.AssertNE(err, nil)
		timeResp, err := p.Time()
		gtest.Assert(err, nil)
		err = p.Sync()
		gtest.Assert(err, nil)
		timeList, err := timeResp.Get()
		gtest.Assert(err, nil)
		t.Log(timeList)
		info, err := infoResp.Get()
		gtest.Assert(err, nil)
		t.Log(info)
	})
}

func TestTransaction_Basic(t *testing.T) {
	gtest.Case(t, func() {
		factory := godis.NewFactory(godis.ShardInfo{
			Host: "localhost",
			Port: 6379,
			Db:   0,
		})
		pool := godis.NewPool(godis.PoolConfig{}, factory)
		redis, _ := pool.GetResource()
		defer redis.Close()
		p, err := redis.Multi()
		gtest.Assert(err, nil)
		infoResp, err := p.Info()
		gtest.Assert(err, nil)
		_, err = infoResp.Get()
		gtest.AssertNE(err, nil)
		timeResp, err := p.Time()
		gtest.Assert(err, nil)
		_, err = p.Exec()
		gtest.Assert(err, nil)
		timeList, err := timeResp.Get()
		gtest.Assert(err, nil)
		t.Log(timeList)
		info, err := infoResp.Get()
		gtest.Assert(err, nil)
		t.Log(info)
	})
}

package godis

import (
	"context"
	"github.com/jolestar/go-commons-pool"
)

type Factory struct {
	shardInfo ShardInfo
}

func NewFactory(shardInfo ShardInfo) *Factory {
	return &Factory{shardInfo: shardInfo}
}

func (f Factory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	redis := NewRedis(f.shardInfo)
	err := redis.Connect()
	if err != nil {
		return nil, err
	}
	return pool.NewPooledObject(redis), nil
}

func (f Factory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	_, err := redis.Quit()
	if err != nil {
		return err
	}
	return nil
}

func (f Factory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	redis := object.Object.(*Redis)
	if redis.Client.Host() != f.shardInfo.Host {
		return false
	}
	if redis.Client.Port() != f.shardInfo.Port {
		return false
	}
	reply, err := redis.Ping()
	if err != nil {
		return false
	}
	return string(reply) == "PONG"
}

func (f Factory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	redis := object.Object.(*Redis)
	if redis.Client.Db == f.shardInfo.Db {
		return nil
	}
	_, err := redis.Select(f.shardInfo.Db)
	if err != nil {
		return err
	}
	return nil
}

func (f Factory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// TODO maybe should select db 0? Not sure right now.
	return nil
}

package godis

import (
	"testing"
	"time"
)

func TestRedis_Keys(t *testing.T) {
}

func TestRedis_Exists(t *testing.T) {
}

func TestRedis_Del(t *testing.T) {
	TestRedis_Set(t)

	type args struct {
		key []string
	}
	tests := []struct {
		name string

		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "del",

			args: args{
				key: []string{"godis"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(option)
			r.Connect()
			got, err := r.Del(tt.args.key...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Del() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Blpop(t *testing.T) {
}

func TestRedis_BlpopTimout(t *testing.T) {
}

func TestRedis_Brpop(t *testing.T) {
}

func TestRedis_BrpopTimout(t *testing.T) {
}

func TestRedis_Mget(t *testing.T) {
}

func TestRedis_Mset(t *testing.T) {
}

func TestRedis_Msetnx(t *testing.T) {
}

func TestRedis_Rename(t *testing.T) {
}

func TestRedis_Renamenx(t *testing.T) {
}

func TestRedis_Brpoplpush(t *testing.T) {
}

func TestRedis_Sdiff(t *testing.T) {
}

func TestRedis_Sdiffstore(t *testing.T) {
}

func TestRedis_Sinter(t *testing.T) {
}

func TestRedis_Sinterstore(t *testing.T) {
}

func TestRedis_Smove(t *testing.T) {
}

func TestRedis_Sort(t *testing.T) {
}

func TestRedis_SortMulti(t *testing.T) {
}

func TestRedis_Sunion(t *testing.T) {
}

func TestRedis_Sunionstore(t *testing.T) {
}

func TestRedis_Watch(t *testing.T) {
}

func TestRedis_Unwatch(t *testing.T) {
}

func TestRedis_Zinterstore(t *testing.T) {
}

func TestRedis_ZinterstoreWithParams(t *testing.T) {
}

func TestRedis_Zunionstore(t *testing.T) {
}

func TestRedis_ZunionstoreWithParams(t *testing.T) {
}

func TestRedis_Publish(t *testing.T) {
}

func TestRedis_Subscribe(t *testing.T) {
	flushAll()

	type args struct {
		redisPubSub *RedisPubSub
		channels    []string
	}
	pubsub := &RedisPubSub{
		OnMessage: func(channel, message string) {
			t.Logf("receive message ,channel:%s,message:%s", channel, message)
		},
		OnSubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive subscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnUnsubscribe: func(channel string, subscribedChannels int) {
			t.Logf("receive unsubscribe command ,channel:%s,subscribedChannels:%d", channel, subscribedChannels)
		},
		OnPMessage: func(pattern string, channel, message string) {
			t.Logf("receive pmessage ,pattern:%s,channel:%s,message:%s", pattern, channel, message)
		},
		OnPSubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive psubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPUnsubscribe: func(pattern string, subscribedChannels int) {
			t.Logf("receive punsubscribe command ,pattern:%s,subscribedChannels:%d", pattern, subscribedChannels)
		},
		OnPong: func(channel string) {
			t.Logf("receive pong ,channel:%s", channel)
		},
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "pubsub",

			args: args{
				redisPubSub: pubsub,
				channels:    []string{"godis"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				r := NewRedis(option)
				r.Connect()
				if err := r.Subscribe(tt.args.redisPubSub, tt.args.channels...); (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
			redis := NewRedis(option)
			redis.Publish("godis", "publish a message to godis channel")
			redis.Close()
			//sleep mills, ensure message can publish to subscribers
			time.Sleep(500 * time.Millisecond)
		})
	}
}

func TestRedis_Psubscribe(t *testing.T) {
}

func TestRedis_RandomKey(t *testing.T) {
}

func TestRedis_Bitop(t *testing.T) {
}

func TestRedis_Scan(t *testing.T) {
}

func TestRedis_Pfmerge(t *testing.T) {
}

func TestRedis_Pfcount(t *testing.T) {
}

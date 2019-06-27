package godis

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_multiKeyPipelineBase_Bgrewriteaof(t *testing.T) {
}

func Test_multiKeyPipelineBase_Bgsave(t *testing.T) {
}

func Test_multiKeyPipelineBase_Bitop(t *testing.T) {
}

func Test_multiKeyPipelineBase_Blpop(t *testing.T) {
}

func Test_multiKeyPipelineBase_BlpopTimout(t *testing.T) {
}

func Test_multiKeyPipelineBase_Brpop(t *testing.T) {
}

func Test_multiKeyPipelineBase_BrpopTimout(t *testing.T) {
}

func Test_multiKeyPipelineBase_Brpoplpush(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterAddSlots(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterDelSlots(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterGetKeysInSlot(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterInfo(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterMeet(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterNodes(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotImporting(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotMigrating(t *testing.T) {
}

func Test_multiKeyPipelineBase_ClusterSetSlotNode(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigGet(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigResetStat(t *testing.T) {
}

func Test_multiKeyPipelineBase_ConfigSet(t *testing.T) {
}

func Test_multiKeyPipelineBase_DbSize(t *testing.T) {
}

func Test_multiKeyPipelineBase_Del(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	p := redis.Pipelined()
	del, err := p.Del("godis")
	assert.Nil(t, err)
	obj, err := ToInt64Reply(del.Get())
	assert.NotNil(t, err)
	//assert.Equal(t, int64(0), obj)
	p.Sync()
	obj, err = ToInt64Reply(del.Get())
	assert.Nil(t, err)
	assert.Equal(t, int64(0), obj)
}

func Test_multiKeyPipelineBase_Eval(t *testing.T) {
}

func Test_multiKeyPipelineBase_Evalsha(t *testing.T) {
}

func Test_multiKeyPipelineBase_Exists(t *testing.T) {
}

func Test_multiKeyPipelineBase_FlushAll(t *testing.T) {
}

func Test_multiKeyPipelineBase_FlushDB(t *testing.T) {
}

func Test_multiKeyPipelineBase_Info(t *testing.T) {
	tests := []struct {
		name    string
		want    *response
		wantErr bool
	}{
		{
			name: "keys",
			want: &response{
				builder: StringBuilder,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Info()
			if (err != nil) != tt.wantErr {
				t.Errorf("Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info() got = %v, want %v", got, tt.want)
				return
			}
			reply, _ := got.Get()
			if reply == "" {
				t.Errorf("Info() got empty string , want not empty string")
			}
		})
	}
}

func Test_multiKeyPipelineBase_Keys(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	redis.Set("godis", "good")
	redis.Close()
	type args struct {
		pattern string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		{
			name: "keys",
			args: args{
				pattern: "*",
			},
			want: &response{
				builder: StringArrayBuilder,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Keys(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
				return
			}
			p.Sync()
			reply, _ := got.Get()
			if !reflect.DeepEqual(reply, []string{"godis"}) {
				t.Errorf("Keys() got = %v, want %v", got, []string{"godis"})
			}
			redis.Close()
		})
	}
}

func Test_multiKeyPipelineBase_Lastsave(t *testing.T) {
}

func Test_multiKeyPipelineBase_Mget(t *testing.T) {
}

func Test_multiKeyPipelineBase_Mset(t *testing.T) {
}

func Test_multiKeyPipelineBase_Msetnx(t *testing.T) {
}

func Test_multiKeyPipelineBase_Pfcount(t *testing.T) {
}

func Test_multiKeyPipelineBase_Pfmerge(t *testing.T) {
}

func Test_multiKeyPipelineBase_Ping(t *testing.T) {
}

func Test_multiKeyPipelineBase_Publish(t *testing.T) {
}

func Test_multiKeyPipelineBase_RandomKey(t *testing.T) {
}

func Test_multiKeyPipelineBase_Rename(t *testing.T) {
}

func Test_multiKeyPipelineBase_Renamenx(t *testing.T) {
}

func Test_multiKeyPipelineBase_Rpoplpush(t *testing.T) {
}

func Test_multiKeyPipelineBase_Save(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sdiff(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sdiffstore(t *testing.T) {}

func Test_multiKeyPipelineBase_Select(t *testing.T) {
}

func Test_multiKeyPipelineBase_Shutdown(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sinter(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sinterstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_Smove(t *testing.T) {

}

func Test_multiKeyPipelineBase_SortMulti(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sunion(t *testing.T) {
}

func Test_multiKeyPipelineBase_Sunionstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_Time(t *testing.T) {
}

func Test_multiKeyPipelineBase_Watch(t *testing.T) {
}

func Test_multiKeyPipelineBase_Zinterstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_ZinterstoreWithParams(t *testing.T) {
}

func Test_multiKeyPipelineBase_Zunionstore(t *testing.T) {
}

func Test_multiKeyPipelineBase_ZunionstoreWithParams(t *testing.T) {
}

func Test_newMultiKeyPipelineBase(t *testing.T) {
}

func Test_newPipeline(t *testing.T) {
}

func Test_newQueable(t *testing.T) {
}

func Test_newResponse(t *testing.T) {
}

func Test_newTransaction(t *testing.T) {
}

func Test_pipeline_Sync(t *testing.T) {
}

func Test_queable_clean(t *testing.T) {
}

func Test_queable_generateResponse(t *testing.T) {
}

func Test_queable_getPipelinedResponseLength(t *testing.T) {
}

func Test_queable_getResponse(t *testing.T) {
}

func Test_queable_hasPipelinedResponse(t *testing.T) {
}

func Test_response_Get(t *testing.T) {
}

func Test_response_build(t *testing.T) {
}

func Test_response_setDependency(t *testing.T) {
}

func Test_response_set_(t *testing.T) {
}

func Test_transaction_Clear(t1 *testing.T) {
}

func Test_transaction_Discard(t1 *testing.T) {
}

func Test_transaction_Exec(t1 *testing.T) {
}

func Test_transaction_ExecGetResponse(t1 *testing.T) {
}

func Test_transaction_clean(t1 *testing.T) {
}

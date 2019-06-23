package godis

import (
	"reflect"
	"testing"
)

func Test_multiKeyPipelineBase_Bgrewriteaof(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Bgrewriteaof()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgrewriteaof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bgrewriteaof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Bgsave(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Bgsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bgsave() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Bitop(t *testing.T) {

	type args struct {
		op      BitOP
		destKey string
		srcKeys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Bitop(tt.args.op, tt.args.destKey, tt.args.srcKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bitop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Blpop(t *testing.T) {

	type args struct {
		args []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Blpop(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Blpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_BlpopTimout(t *testing.T) {

	type args struct {
		timeout int
		keys    []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.BlpopTimout(tt.args.timeout, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlpopTimout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlpopTimout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Brpop(t *testing.T) {

	type args struct {
		args []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Brpop(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Brpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Brpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_BrpopTimout(t *testing.T) {

	type args struct {
		timeout int
		keys    []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.BrpopTimout(tt.args.timeout, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrpopTimout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BrpopTimout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Brpoplpush(t *testing.T) {

	type args struct {
		source      string
		destination string
		timeout     int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Brpoplpush(tt.args.source, tt.args.destination, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Brpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Brpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterAddSlots(t *testing.T) {

	type args struct {
		slots []int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterAddSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterAddSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterAddSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterDelSlots(t *testing.T) {

	type args struct {
		slots []int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterDelSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterDelSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterDelSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterGetKeysInSlot(t *testing.T) {

	type args struct {
		slot  int
		count int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterGetKeysInSlot(tt.args.slot, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterGetKeysInSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterGetKeysInSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterInfo(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterMeet(t *testing.T) {

	type args struct {
		ip   string
		port int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterMeet(tt.args.ip, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterMeet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterMeet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterNodes(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterNodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterSetSlotImporting(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterSetSlotImporting(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotImporting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSetSlotImporting() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterSetSlotMigrating(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterSetSlotMigrating(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotMigrating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSetSlotMigrating() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ClusterSetSlotNode(t *testing.T) {

	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ClusterSetSlotNode(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSetSlotNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ConfigGet(t *testing.T) {

	type args struct {
		pattern string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ConfigGet(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ConfigResetStat(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ConfigResetStat()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigResetStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigResetStat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ConfigSet(t *testing.T) {

	type args struct {
		parameter string
		value     string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ConfigSet(tt.args.parameter, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_DbSize(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.DbSize()
			if (err != nil) != tt.wantErr {
				t.Errorf("DbSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Del(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Del(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Del() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Eval(t *testing.T) {

	type args struct {
		script   string
		keyCount int
		params   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Eval(tt.args.script, tt.args.keyCount, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Evalsha(t *testing.T) {

	type args struct {
		sha1     string
		keyCount int
		params   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Evalsha(tt.args.sha1, tt.args.keyCount, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evalsha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evalsha() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Exists(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Exists(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_FlushAll(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.FlushAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlushAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_FlushDB(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.FlushDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlushDB() got = %v, want %v", got, tt.want)
			}
		})
	}
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

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Lastsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lastsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lastsave() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Mget(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Mget(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Mset(t *testing.T) {
	type args struct {
		keysvalues []string
	}
	tests := []struct {
		name    string
		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Mset(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Msetnx(t *testing.T) {

	type args struct {
		keysvalues []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Msetnx(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Msetnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Msetnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Pfcount(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Pfcount(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pfcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Pfmerge(t *testing.T) {

	type args struct {
		destkey    string
		sourcekeys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Pfmerge(tt.args.destkey, tt.args.sourcekeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfmerge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pfmerge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Ping(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Publish(t *testing.T) {

	type args struct {
		channel string
		message string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Publish(tt.args.channel, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Publish() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_RandomKey(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.RandomKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Rename(t *testing.T) {

	type args struct {
		oldkey string
		newkey string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Rename(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rename() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Renamenx(t *testing.T) {

	type args struct {
		oldkey string
		newkey string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Renamenx(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renamenx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Renamenx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Rpoplpush(t *testing.T) {

	type args struct {
		srckey string
		dstkey string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Rpoplpush(tt.args.srckey, tt.args.dstkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Save(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Save()
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sdiff(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sdiff(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sdiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sdiff() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sdiffstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sdiffstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sdiffstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sdiffstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Select(t *testing.T) {

	type args struct {
		index int
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Select(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Shutdown(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Shutdown()
			if (err != nil) != tt.wantErr {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shutdown() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sinter(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sinter(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sinter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sinter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sinterstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sinterstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Smove(t *testing.T) {

	type args struct {
		srckey string
		dstkey string
		member string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Smove(tt.args.srckey, tt.args.dstkey, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Smove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Smove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_SortMulti(t *testing.T) {

	type args struct {
		key               string
		dstkey            string
		sortingParameters []SortingParams
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.SortMulti(tt.args.key, tt.args.dstkey, tt.args.sortingParameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortMulti() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMulti() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sunion(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sunion(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sunion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sunion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Sunionstore(t *testing.T) {

	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Sunionstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Time(t *testing.T) {

	tests := []struct {
		name string

		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Watch(t *testing.T) {

	type args struct {
		keys []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Watch(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Watch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Zinterstore(t *testing.T) {

	type args struct {
		dstkey string
		sets   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Zinterstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ZinterstoreWithParams(t *testing.T) {

	type args struct {
		dstkey string
		params ZParams
		sets   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ZinterstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZinterstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZinterstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_Zunionstore(t *testing.T) {

	type args struct {
		dstkey string
		sets   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.Zunionstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiKeyPipelineBase_ZunionstoreWithParams(t *testing.T) {

	type args struct {
		dstkey string
		params ZParams
		sets   []string
	}
	tests := []struct {
		name string

		args    args
		want    *response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redis := NewRedis(option)
			p := redis.Pipelined()
			got, err := p.ZunionstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZunionstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZunionstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newMultiKeyPipelineBase(t *testing.T) {
	type args struct {
		client *client
	}
	tests := []struct {
		name string
		args args
		want *multiKeyPipelineBase
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMultiKeyPipelineBase(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMultiKeyPipelineBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPipeline(t *testing.T) {
	type args struct {
		c *client
	}
	tests := []struct {
		name string
		args args
		want *pipeline
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPipeline(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newQueable(t *testing.T) {
	tests := []struct {
		name string
		want *queable
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newResponse(t *testing.T) {
	tests := []struct {
		name string
		want *response
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newTransaction(t *testing.T) {
	type args struct {
		c *client
	}
	tests := []struct {
		name string
		args args
		want *transaction
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransaction(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pipeline_Sync(t *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pipeline{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
			}
			if err := p.Sync(); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_queable_clean(t *testing.T) {
	type fields struct {
		pipelinedResponses []*response
	}
	tests := []struct {
		name   string
		fields fields
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &queable{
				pipelinedResponses: tt.fields.pipelinedResponses,
			}
		})
	}
}

func Test_queable_generateResponse(t *testing.T) {
	type fields struct {
		pipelinedResponses []*response
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *response
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queable{
				pipelinedResponses: tt.fields.pipelinedResponses,
			}
			if got := q.generateResponse(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queable_getPipelinedResponseLength(t *testing.T) {
	type fields struct {
		pipelinedResponses []*response
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queable{
				pipelinedResponses: tt.fields.pipelinedResponses,
			}
			if got := q.getPipelinedResponseLength(); got != tt.want {
				t.Errorf("getPipelinedResponseLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queable_getResponse(t *testing.T) {
	type fields struct {
		pipelinedResponses []*response
	}
	type args struct {
		builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *response
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queable{
				pipelinedResponses: tt.fields.pipelinedResponses,
			}
			if got := q.getResponse(tt.args.builder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queable_hasPipelinedResponse(t *testing.T) {
	type fields struct {
		pipelinedResponses []*response
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queable{
				pipelinedResponses: tt.fields.pipelinedResponses,
			}
			if got := q.hasPipelinedResponse(); got != tt.want {
				t.Errorf("hasPipelinedResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Get(t *testing.T) {
	type fields struct {
		response   interface{}
		building   bool
		built      bool
		set        bool
		builder    Builder
		data       interface{}
		dependency *response
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				response:   tt.fields.response,
				building:   tt.fields.building,
				built:      tt.fields.built,
				set:        tt.fields.set,
				builder:    tt.fields.builder,
				data:       tt.fields.data,
				dependency: tt.fields.dependency,
			}
			got, err := r.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_build(t *testing.T) {
	type fields struct {
		response   interface{}
		building   bool
		built      bool
		set        bool
		builder    Builder
		data       interface{}
		dependency *response
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				response:   tt.fields.response,
				building:   tt.fields.building,
				built:      tt.fields.built,
				set:        tt.fields.set,
				builder:    tt.fields.builder,
				data:       tt.fields.data,
				dependency: tt.fields.dependency,
			}
			if err := r.build(); (err != nil) != tt.wantErr {
				t.Errorf("build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_response_setDependency(t *testing.T) {
	type fields struct {
		response   interface{}
		building   bool
		built      bool
		set        bool
		builder    Builder
		data       interface{}
		dependency *response
	}
	type args struct {
		dependency *response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &response{
				response:   tt.fields.response,
				building:   tt.fields.building,
				built:      tt.fields.built,
				set:        tt.fields.set,
				builder:    tt.fields.builder,
				data:       tt.fields.data,
				dependency: tt.fields.dependency,
			}
		})
	}
}

func Test_response_set_(t *testing.T) {
	type fields struct {
		response   interface{}
		building   bool
		built      bool
		set        bool
		builder    Builder
		data       interface{}
		dependency *response
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &response{
				response:   tt.fields.response,
				building:   tt.fields.building,
				built:      tt.fields.built,
				set:        tt.fields.set,
				builder:    tt.fields.builder,
				data:       tt.fields.data,
				dependency: tt.fields.dependency,
			}
		})
	}
}

func Test_transaction_Clear(t1 *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
		inTransaction        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transaction{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
				inTransaction:        tt.fields.inTransaction,
			}
			got, err := t.Clear()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Clear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Clear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_Discard(t1 *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
		inTransaction        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transaction{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
				inTransaction:        tt.fields.inTransaction,
			}
			got, err := t.Discard()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Discard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Discard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_Exec(t1 *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
		inTransaction        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []interface{}
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transaction{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
				inTransaction:        tt.fields.inTransaction,
			}
			got, err := t.Exec()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Exec() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_ExecGetResponse(t1 *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
		inTransaction        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*response
		wantErr bool
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transaction{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
				inTransaction:        tt.fields.inTransaction,
			}
			got, err := t.ExecGetResponse()
			if (err != nil) != tt.wantErr {
				t1.Errorf("ExecGetResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ExecGetResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transaction_clean(t1 *testing.T) {
	type fields struct {
		multiKeyPipelineBase *multiKeyPipelineBase
		inTransaction        bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		/*{
			name:"keys",
			args:args{
				pattern:"*",
			},
			want:&response{

			},
			wantErr:false,
		},*/
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			_ = &transaction{
				multiKeyPipelineBase: tt.fields.multiKeyPipelineBase,
				inTransaction:        tt.fields.inTransaction,
			}
		})
	}
}

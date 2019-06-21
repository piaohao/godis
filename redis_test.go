package godis

import (
	"reflect"
	"testing"
)

var option = Option{
	//Host:     "10.1.1.63",
	//Password: "123456",
	Host: "localhost",
	Port: 6379,
	Db:   0,
}

func TestNewRedis(t *testing.T) {
	type args struct {
		option Option
	}
	tests := []struct {
		name string
		args args
		want *Redis
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedis(tt.args.option); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Append(t *testing.T) {
	TestRedis_Del(t)
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "godis",
				value: "very",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "godis",
				value: " good",
			},
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			r.Connect()
			got, err := r.Append(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Append() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Asking(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "asking",
			fields: fields{
				client: newClient(option),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			r.Connect()
			got, err := r.Asking()
			if (err != nil) != tt.wantErr {
				t.Errorf("Asking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Asking() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_Auth(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "auth1",
			fields: fields{
				client: newClient(option),
			},
			args:    args{password: "123456"},
			want:    "OK",
			wantErr: false,
		},*/
		{
			name: "auth2",
			fields: fields{
				client: newClient(option),
			},
			args:    args{password: "1234567"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Auth(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bgrewriteaof(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Bgrewriteaof()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgrewriteaof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bgrewriteaof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bgsave(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Bgsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bgsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bgsave() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bitcount(t *testing.T) {
	TestRedis_Set(t)
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key: "godis",
			},
			want:    20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			r.Connect()
			got, err := r.Bitcount(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitcount() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_BitcountRange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.BitcountRange(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitcountRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BitcountRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bitfield(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		arguments []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Bitfield(tt.args.key, tt.args.arguments...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitfield() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bitfield() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bitop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		op      BitOP
		destKey string
		srcKeys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Bitop(tt.args.op, tt.args.destKey, tt.args.srcKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Bitpos(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		value  bool
		params []BitPosParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Bitpos(tt.args.key, tt.args.value, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bitpos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bitpos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Blpop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Blpop(tt.args.args...)
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

func TestRedis_BlpopTimout(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		timeout int
		keys    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.BlpopTimout(tt.args.timeout, tt.args.keys...)
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

func TestRedis_Brpop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Brpop(tt.args.args...)
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

func TestRedis_BrpopTimout(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		timeout int
		keys    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.BrpopTimout(tt.args.timeout, tt.args.keys...)
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

func TestRedis_Brpoplpush(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		source      string
		destination string
		timeout     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Brpoplpush(tt.args.source, tt.args.destination, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Brpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Brpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Close(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_ClusterAddSlots(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slots []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterAddSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterAddSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterAddSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterCountKeysInSlot(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterCountKeysInSlot(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterCountKeysInSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterCountKeysInSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterDelSlots(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slots []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterDelSlots(tt.args.slots...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterDelSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterDelSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterFailover(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterFailover()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterFailover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterFailover() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterFlushSlots(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterFlushSlots()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterFlushSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterFlushSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterForget(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterForget(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterForget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterForget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterGetKeysInSlot(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot  int
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterGetKeysInSlot(tt.args.slot, tt.args.count)
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

func TestRedis_ClusterInfo(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterKeySlot(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterKeySlot(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterKeySlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterKeySlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterMeet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		ip   string
		port int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterMeet(tt.args.ip, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterMeet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterMeet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterNodes(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterNodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterReplicate(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterReplicate(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterReplicate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterReplicate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterReset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		resetType Reset
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterReset(tt.args.resetType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSaveConfig(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSaveConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSaveConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSaveConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotImporting(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSetSlotImporting(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotImporting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotImporting() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotMigrating(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSetSlotMigrating(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotMigrating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotMigrating() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotNode(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot   int
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSetSlotNode(tt.args.slot, tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetSlotStable(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSetSlotStable(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSetSlotStable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClusterSetSlotStable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSlaves(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		nodeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSlaves(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSlaves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSlaves() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSlots(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    []interface{}
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ClusterSlots()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterSlots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterSlots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ConfigGet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ConfigGet(tt.args.pattern)
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

func TestRedis_ConfigResetStat(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ConfigResetStat()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigResetStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfigResetStat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ConfigSet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		parameter string
		value     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ConfigSet(tt.args.parameter, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfigSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Connect(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_DbSize(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.DbSize()
			if (err != nil) != tt.wantErr {
				t.Errorf("DbSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DbSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Debug(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		params DebugParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Debug(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Debug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Debug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Decr(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Decr(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_DecrBy(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		decrement int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.DecrBy(tt.args.key, tt.args.decrement)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Del(t *testing.T) {
	TestRedis_Set(t)
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "del",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key: []string{"godis"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
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

func TestRedis_Echo(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		string string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Echo(tt.args.string)
			if (err != nil) != tt.wantErr {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Echo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Eval(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		script   string
		keyCount int
		params   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Eval(tt.args.script, tt.args.keyCount, tt.args.params...)
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

func TestRedis_Evalsha(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		sha1     string
		keyCount int
		params   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Evalsha(tt.args.sha1, tt.args.keyCount, tt.args.params...)
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

func TestRedis_Exists(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Exists(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Expire(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		seconds int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Expire(tt.args.key, tt.args.seconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Expire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ExpireAt(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key      string
		unixtime int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ExpireAt(tt.args.key, tt.args.unixtime)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpireAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExpireAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_FlushAll(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.FlushAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlushAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_FlushDB(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.FlushDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlushDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlushDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Geoadd(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		longitude float64
		latitude  float64
		member    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Geoadd(tt.args.key, tt.args.longitude, tt.args.latitude, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geoadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Geoadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_GeoaddByMap(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key                 string
		memberCoordinateMap map[string]GeoCoordinate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.GeoaddByMap(tt.args.key, tt.args.memberCoordinateMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoaddByMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoaddByMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Geodist(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		member1 string
		member2 string
		unit    []GeoUnit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Geodist(tt.args.key, tt.args.member1, tt.args.member2, tt.args.unit...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geodist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Geodist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Geohash(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Geohash(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geohash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Geohash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Geopos(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Geopos(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Geopos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Geopos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Georadius(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		longitude float64
		latitude  float64
		radius    float64
		unit      GeoUnit
		param     []GeoRadiusParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Georadius(tt.args.key, tt.args.longitude, tt.args.latitude, tt.args.radius, tt.args.unit, tt.args.param...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Georadius() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Georadius() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_GeoradiusByMember(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		member string
		radius float64
		unit   GeoUnit
		param  []GeoRadiusParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*GeoCoordinate
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.GeoradiusByMember(tt.args.key, tt.args.member, tt.args.radius, tt.args.unit, tt.args.param...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoradiusByMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoradiusByMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Get(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_GetDB(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if got := r.GetDB(); got != tt.want {
				t.Errorf("GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_GetSet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.GetSet(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Getbit(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		offset int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Getbit(tt.args.key, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getbit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getbit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Getrange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key         string
		startOffset int64
		endOffset   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Getrange(tt.args.key, tt.args.startOffset, tt.args.endOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hdel(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hdel(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hdel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hdel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hexists(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hexists(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hexists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hexists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hget(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hget(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_HgetAll(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.HgetAll(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HgetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HgetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_HincrBy(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.HincrBy(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_HincrByFloat(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.HincrByFloat(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HincrByFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HincrByFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hkeys(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hkeys(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hkeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hkeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hlen(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hlen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hlen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hlen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hmget(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hmget(tt.args.key, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hmget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hmget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hmset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key  string
		hash map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hmset(tt.args.key, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hmset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hmset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hscan(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		cursor string
		params []ScanParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ScanResult
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hset(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hsetnx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		field string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hsetnx(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hsetnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hsetnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Hvals(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Hvals(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hvals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hvals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Incr(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Incr(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Incr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Incr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_IncrBy(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		increment int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.IncrBy(tt.args.key, tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncrBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncrBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_IncrByFloat(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		increment float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.IncrByFloat(tt.args.key, tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncrByFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncrByFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Info(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		section []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Info(tt.args.section...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Keys(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Keys(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lastsave(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lastsave()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lastsave() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lastsave() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lindex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		index int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lindex(tt.args.key, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lindex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lindex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Linsert(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		where ListOption
		pivot string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Linsert(tt.args.key, tt.args.where, tt.args.pivot, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Linsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Linsert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Llen(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Llen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Llen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Llen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lpop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lpop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lpush(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		strings []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lpush(tt.args.key, tt.args.strings...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lpushx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		string []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lpushx(tt.args.key, tt.args.string...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lpushx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lpushx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lrange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		stop  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lrange(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lrem(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		count int64
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lrem(tt.args.key, tt.args.count, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lrem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lrem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Lset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		index int64
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Lset(tt.args.key, tt.args.index, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Ltrim(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		stop  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Ltrim(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ltrim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ltrim() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Mget(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Mget(tt.args.keys...)
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

func TestRedis_Move(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		dbIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Move(tt.args.key, tt.args.dbIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Mset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keysvalues []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Mset(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Msetnx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keysvalues []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Msetnx(tt.args.keysvalues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Msetnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Msetnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Multi(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    *transaction
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Multi()
			if (err != nil) != tt.wantErr {
				t.Errorf("Multi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectEncoding(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ObjectEncoding(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectEncoding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectIdletime(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ObjectIdletime(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectIdletime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectIdletime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ObjectRefcount(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ObjectRefcount(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectRefcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ObjectRefcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Persist(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Persist(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Persist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Persist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pexpire(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key          string
		milliseconds int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Pexpire(tt.args.key, tt.args.milliseconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pexpire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pexpire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_PexpireAt(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key                   string
		millisecondsTimestamp int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.PexpireAt(tt.args.key, tt.args.millisecondsTimestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("PexpireAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PexpireAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pfadd(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key      string
		elements []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Pfadd(tt.args.key, tt.args.elements...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pfcount(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Pfcount(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pfmerge(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		destkey    string
		sourcekeys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Pfmerge(tt.args.destkey, tt.args.sourcekeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pfmerge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pfmerge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Ping(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pipelined(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   *pipeline
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if got := r.Pipelined(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pipelined() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Psetex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key          string
		milliseconds int64
		value        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Psetex(tt.args.key, tt.args.milliseconds, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Psetex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Psetex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Psubscribe(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		redisPubSub *RedisPubSub
		patterns    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Psubscribe(tt.args.redisPubSub, tt.args.patterns...); (err != nil) != tt.wantErr {
				t.Errorf("Psubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Pttl(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Pttl(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pttl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pttl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Publish(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		channel string
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Publish(tt.args.channel, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Publish() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_PubsubChannels(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.PubsubChannels(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("PubsubChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PubsubChannels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Quit(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Quit()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_RandomKey(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.RandomKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandomKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Readonly(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Readonly()
			if (err != nil) != tt.wantErr {
				t.Errorf("Readonly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Readonly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Receive(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Receive(); (err != nil) != tt.wantErr {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Rename(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		oldkey string
		newkey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Rename(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rename() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Renamenx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		oldkey string
		newkey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Renamenx(tt.args.oldkey, tt.args.newkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renamenx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Renamenx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Rpop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Rpop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Rpoplpush(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		srckey string
		dstkey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Rpoplpush(tt.args.srckey, tt.args.dstkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpoplpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpoplpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Rpush(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		strings []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Rpush(tt.args.key, tt.args.strings...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Rpushx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		string []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Rpushx(tt.args.key, tt.args.string...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpushx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpushx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sadd(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sadd(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Save(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Save()
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Scan(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		cursor string
		params []ScanParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ScanResult
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Scan(tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Scard(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Scard(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Scard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ScriptExists(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		sha1 []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ScriptExists(tt.args.sha1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ScriptLoad(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		script string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ScriptLoad(tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ScriptLoad() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sdiff(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sdiff(tt.args.keys...)
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

func TestRedis_Sdiffstore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sdiffstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sdiffstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sdiffstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Select(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Select(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Send(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		command protocolCommand
		args    [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Send(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_SendByStr(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		command string
		args    [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.SendByStr(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("SendByStr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_SentinelFailover(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelFailover(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelFailover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelFailover() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelGetMasterAddrByName(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelGetMasterAddrByName(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelGetMasterAddrByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelGetMasterAddrByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelMasters(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    []map[string]string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelMasters()
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelMasters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelMasters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelMonitor(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName string
		ip         string
		port       int
		quorum     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelMonitor(tt.args.masterName, tt.args.ip, tt.args.port, tt.args.quorum)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelMonitor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelRemove(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelRemove(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelRemove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelRemove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelReset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelReset(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelSet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName   string
		parameterMap map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelSet(tt.args.masterName, tt.args.parameterMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentinelSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SentinelSlaves(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		masterName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SentinelSlaves(tt.args.masterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelSlaves() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SentinelSlaves() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Set(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "set",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "godis",
				value: "good",
			},
			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			r.Connect()
			got, err := r.Set(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
			r.Close()
		})
	}
}

func TestRedis_SetWithParams(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
		nxxx  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SetWithParams(tt.args.key, tt.args.value, tt.args.nxxx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SetWithParamsAndTime(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
		nxxx  string
		expx  string
		time  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SetWithParamsAndTime(tt.args.key, tt.args.value, tt.args.nxxx, tt.args.expx, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetWithParamsAndTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetWithParamsAndTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Setbit(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		offset int64
		value  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Setbit(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setbit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setbit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SetbitWithBool(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		offset int64
		value  bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SetbitWithBool(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetbitWithBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetbitWithBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Setex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		seconds int
		value   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Setex(tt.args.key, tt.args.seconds, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Setnx(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Setnx(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setnx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setnx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Setrange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		offset int64
		value  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Setrange(tt.args.key, tt.args.offset, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Setrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Shutdown(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Shutdown()
			if (err != nil) != tt.wantErr {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shutdown() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sinter(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sinter(tt.args.keys...)
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

func TestRedis_Sinterstore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sinterstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sismember(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sismember(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sismember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sismember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Slaveof(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Slaveof(tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("Slaveof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Slaveof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlaveofNoOne(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SlaveofNoOne()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlaveofNoOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlaveofNoOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogGet(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		entries []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Slowlog
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SlowlogGet(tt.args.entries...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlowlogGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogLen(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SlowlogLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlowlogLen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SlowlogReset(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SlowlogReset()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlowlogReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlowlogReset() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Smembers(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Smembers(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Smembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Smembers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Smove(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		srckey string
		dstkey string
		member string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Smove(tt.args.srckey, tt.args.dstkey, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Smove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Smove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sort(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key               string
		sortingParameters []SortingParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sort(tt.args.key, tt.args.sortingParameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SortMulti(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key               string
		dstkey            string
		sortingParameters []SortingParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SortMulti(tt.args.key, tt.args.dstkey, tt.args.sortingParameters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortMulti() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SortMulti() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Spop(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Spop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Spop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SpopBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		count int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SpopBatch(tt.args.key, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpopBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpopBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Srandmember(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Srandmember(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Srandmember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Srandmember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SrandmemberBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.SrandmemberBatch(tt.args.key, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SrandmemberBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SrandmemberBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Srem(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Srem(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Srem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Srem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sscan(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		cursor string
		params []ScanParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ScanResult
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Strlen(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Strlen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strlen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Strlen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Subscribe(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		redisPubSub *RedisPubSub
		channels    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.Subscribe(tt.args.redisPubSub, tt.args.channels...); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_Substr(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int
		end   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Substr(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Substr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Substr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Sunion(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sunion(tt.args.keys...)
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

func TestRedis_Sunionstore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		keys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Sunionstore(tt.args.dstkey, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Ttl(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Ttl(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ttl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ttl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Type(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Type(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Type() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Type() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Unwatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Unwatch()
			if (err != nil) != tt.wantErr {
				t.Errorf("Unwatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unwatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_WaitReplicas(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		replicas int
		timeout  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.WaitReplicas(tt.args.replicas, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaitReplicas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WaitReplicas() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Watch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Watch(tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Watch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zadd(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		score   float64
		member  string
		mparams []ZAddParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zadd(tt.args.key, tt.args.score, tt.args.member, tt.args.mparams...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zadd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zadd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZaddByMap(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key          string
		scoreMembers map[string]float64
		params       []ZAddParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZaddByMap(tt.args.key, tt.args.scoreMembers, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZaddByMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZaddByMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zcard(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zcard(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zcard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zcard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zcount(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zcount(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zincrby(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key       string
		increment float64
		member    string
		params    []ZAddParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zincrby(tt.args.key, tt.args.increment, tt.args.member, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zincrby() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zincrby() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zinterstore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		sets   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zinterstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zinterstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zinterstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZinterstoreWithParams(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		params ZParams
		sets   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZinterstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZinterstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZinterstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zlexcount(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zlexcount(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zlexcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zlexcount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zrange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		stop  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zrange(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByLex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByLex(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByLexBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByLexBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByLexBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByLexBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByScore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByScore(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByScoreBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByScoreBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByScoreWithScores(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByScoreWithScores(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeByScoreWithScoresBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		min    string
		max    string
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeByScoreWithScoresBatch(tt.args.key, tt.args.min, tt.args.max, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeByScoreWithScoresBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeByScoreWithScoresBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrangeWithScores(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrangeWithScores(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrangeWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrangeWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zrank(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zrank(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zrem(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key     string
		members []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zrem(tt.args.key, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZremrangeByLex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		min string
		max string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZremrangeByLex(tt.args.key, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZremrangeByRank(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		stop  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZremrangeByRank(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByRank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZremrangeByScore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZremrangeByScore(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZremrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZremrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zrevrange(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		stop  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zrevrange(tt.args.key, tt.args.start, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrevrange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zrevrange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeByLex(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		max string
		min string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeByLex(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByLex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByLex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeByLexBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		max    string
		min    string
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeByLexBatch(tt.args.key, tt.args.max, tt.args.min, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByLexBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByLexBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeByScore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		max string
		min string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeByScore(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeByScoreWithScores(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key string
		max string
		min string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeByScoreWithScores(tt.args.key, tt.args.max, tt.args.min)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScoreWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScoreWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeByScoreWithScoresBatch(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		max    string
		min    string
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeByScoreWithScoresBatch(tt.args.key, tt.args.max, tt.args.min, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeByScoreWithScoresBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeByScoreWithScoresBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZrevrangeWithScores(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key   string
		start int64
		end   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Tuple
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZrevrangeWithScores(tt.args.key, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZrevrangeWithScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZrevrangeWithScores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zrevrank(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zrevrank(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zrevrank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zrevrank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zscan(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		cursor string
		params []ScanParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ScanResult
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zscan(tt.args.key, tt.args.cursor, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zscan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zscore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		key    string
		member string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zscore(tt.args.key, tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zscore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zscore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_Zunionstore(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		sets   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.Zunionstore(tt.args.dstkey, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zunionstore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Zunionstore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ZunionstoreWithParams(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	type args struct {
		dstkey string
		params ZParams
		sets   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			got, err := r.ZunionstoreWithParams(tt.args.dstkey, tt.args.params, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZunionstoreWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZunionstoreWithParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_checkIsInMultiOrPipeline(t *testing.T) {
	type fields struct {
		client      *client
		pipeline    *pipeline
		transaction *transaction
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		/*{
			name: "append",
			fields: fields{
				client: newClient(option),
			},
			args: args{
				key:   "a",
				value: "b",
			},
			want:    1,
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client:      tt.fields.client,
				pipeline:    tt.fields.pipeline,
				transaction: tt.fields.transaction,
			}
			if err := r.checkIsInMultiOrPipeline(); (err != nil) != tt.wantErr {
				t.Errorf("checkIsInMultiOrPipeline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

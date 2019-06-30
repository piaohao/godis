package godis

import (
	"reflect"
	"testing"
)

func Test_crc16_getByteSlot(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "crc16",
			args: args{key: []byte("godis")},
			want: 3033,
		},
		{
			name: "crc16",
			args: args{key: []byte("test{godis}")},
			want: 3033,
		},
		{
			name: "crc16",
			args: args{key: []byte("test2{godis}")},
			want: 3033,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crc16{
				tagUtil: newRedisClusterHashTagUtil(),
			}
			if got := c.getByteSlot(tt.args.key); got != tt.want {
				t.Errorf("getByteSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crc16_getBytesCRC16(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "crc16",
			args: args{bytes: []byte("godis")},
			want: 19417,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crc16{
				tagUtil: newRedisClusterHashTagUtil(),
			}
			if got := c.getBytesCRC16(tt.args.bytes); got != tt.want {
				t.Errorf("getBytesCRC16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crc16_getCRC16(t *testing.T) {
	type args struct {
		bytes []byte
		s     int
		e     int
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "crc16",
			args: args{bytes: []byte("godis"), s: 0, e: 5},
			want: 19417,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crc16{
				tagUtil: newRedisClusterHashTagUtil(),
			}
			if got := c.getCRC16(tt.args.bytes, tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("getCRC16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crc16_getStringCRC16(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "crc16",
			args: args{key: "godis"},
			want: 19417,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crc16{
				tagUtil: newRedisClusterHashTagUtil(),
			}
			if got := c.getStringCRC16(tt.args.key); got != tt.want {
				t.Errorf("getStringCRC16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crc16_getStringSlot(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "crc16",
			args: args{key: "godis"},
			want: 3033,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crc16{
				tagUtil: newRedisClusterHashTagUtil(),
			}
			if got := c.getStringSlot(tt.args.key); got != tt.want {
				t.Errorf("getStringSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newCRC16(t *testing.T) {
	tests := []struct {
		name string
		want *crc16
	}{
		{
			name: "crc16",
			want: newCRC16(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCRC16(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCRC16() = %v, want %v", got, tt.want)
			}
		})
	}
}

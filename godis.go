package godis

type RedisOptions struct {
	Host              string
	Port              int
	ConnectionTimeout int
	SoTimeout         int
	Password          string
	Db                int
	IsInMulti         bool
	IsInWatch         bool
	Ssl               bool
}

type BinaryRedis struct {
	Client *BinaryClient
}

func NewBinaryRedis(options RedisOptions) (*BinaryRedis, error) {
	client := NewBinaryClient(options)
	r := &BinaryRedis{Client: client}
	err := client.Connect()
	return r, err
}

func (r *BinaryRedis) Set(key, value []byte) (string, error) {
	err := r.Client.Set(key, value)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *BinaryRedis) Get(key []byte) ([]byte, error) {
	err := r.Client.Get(key)
	if err != nil {
		return nil, err
	}
	return r.Client.getBinaryBulkReply()
}

func (r *BinaryRedis) Del(key []byte) (int64, error) {
	err := r.Client.Del(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *BinaryRedis) DelBatch(keys ...[]byte) (int64, error) {
	err := r.Client.DelBatch(keys...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

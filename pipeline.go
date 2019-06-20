package godis

import "errors"

type RedisPipeline interface {
	Set(key, value string) (*Response, error)
	SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (*Response, error)
	SetWithParams(key, value, nxxx string) (*Response, error)
	Get(key string) (*Response, error)
	//Exists(key string) (*Response, error)
	Persist(key string) (*Response, error)
	Type(key string) (*Response, error)
	Expire(key string, seconds int) (*Response, error)
	Pexpire(key string, milliseconds int64) (*Response, error)
	ExpireAt(key string, unixtime int64) (*Response, error)
	PexpireAt(key string, millisecondsTimestamp int64) (*Response, error)
	Ttl(key string) (*Response, error)
	Pttl(key string) (*Response, error)
	SetbitWithBool(key string, offset int64, value bool) (*Response, error)
	Setbit(key string, offset int64, value string) (*Response, error)
	Getbit(key string, offset int64) (*Response, error)
	Setrange(key string, offset int64, value string) (*Response, error)
	Getrange(key string, startOffset, endOffset int64) (*Response, error)
	GetSet(key, value string) (*Response, error)

	Setnx(key, value string) (*Response, error)
	Setex(key string, seconds int, value string) (*Response, error)
	Psetex(key string, milliseconds int64, value string) (*Response, error)
	DecrBy(key string, decrement int64) (*Response, error)
	Decr(key string) (*Response, error)
	IncrBy(key string, increment int64) (*Response, error)
	IncrByFloat(key string, increment float64) (*Response, error)
	Incr(key string) (*Response, error)
	Append(key, value string) (*Response, error)
	Substr(key string, start, end int) (*Response, error)

	Hset(key, field string, value string) (*Response, error)
	Hget(key, field string) (*Response, error)
	Hsetnx(key, field, value string) (*Response, error)
	Hmset(key string, hash map[string]string) (*Response, error)
	Hmget(key string, fields ...string) (*Response, error)
	HincrBy(key, field string, value int64) (*Response, error)
	HincrByFloat(key, field string, value float64) (*Response, error)
	Hexists(key, field string) (*Response, error)
	Hdel(key string, fields ...string) (*Response, error)
	Hlen(key string) (*Response, error)
	Hkeys(key string) (*Response, error)
	Hvals(key string) (*Response, error)
	HgetAll(key string) (*Response, error)

	Rpush(key string, strings ...string) (*Response, error)
	Lpush(key string, strings ...string) (*Response, error)
	Llen(key string) (*Response, error)
	Lrange(key string, start, stop int64) (*Response, error)
	Ltrim(key string, start, stop int64) (*Response, error)
	Lindex(key string, index int64) (*Response, error)
	Lset(key string, index int64, value string) (*Response, error)
	Lrem(key string, count int64, value string) (*Response, error)
	Lpop(key string) (*Response, error)
	Rpop(key string) (*Response, error)
	Sadd(key string, members ...string) (*Response, error)
	Smembers(key string) (*Response, error)
	Srem(key string, members ...string) (*Response, error)
	Spop(key string) (*Response, error)
	SpopBatch(key string, count int64) (*Response, error)

	Scard(key string) (*Response, error)
	Sismember(key string, member string) (*Response, error)
	Srandmember(key string) (*Response, error)
	SrandmemberBatch(key string, count int) (*Response, error)
	Strlen(key string) (*Response, error)

	//Zadd(key string, score float64, member string) (*Response, error)
	Zadd(key string, score float64, member string, params ...ZAddParams) (*Response, error)
	//Zadd(key string, scoreMembers map[string]float64) (*Response, error)
	ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (*Response, error)
	Zrange(key string, start, end int64) (*Response, error)
	Zrem(key string, member ...string) (*Response, error)
	//Zincrby(key string, score float64, member string) (*Response, error)
	Zincrby(key string, score float64, member string, params ...ZAddParams) (*Response, error)
	Zrank(key, member string) (*Response, error)
	Zrevrank(key, member string) (*Response, error)
	Zrevrange(key string, start, end int64) (*Response, error)
	ZrangeWithScores(key string, start, end int64) (*Response, error)
	ZrevrangeWithScores(key string, start, end int64) (*Response, error)
	Zcard(key string) (*Response, error)
	Zscore(key, member string) (*Response, error)
	//Sort(key string) (*Response, error)
	Sort(key string, sortingParameters ...SortingParams) (*Response, error)
	//Zcount(key string, min float64, max float64) (*Response, error)
	Zcount(key string, min string, max string) (*Response, error)
	//ZrangeByScore(key string, min float64, max float64) (*Response, error)
	ZrangeByScore(key string, min string, max string) (*Response, error)
	//ZrevrangeByScore(key string, max float64, min float64) (*Response, error)
	//ZrangeByScoreBatch(key string, min float64, max float64, offset int, count int) (*Response, error)
	ZrevrangeByScore(key string, max string, min string) (*Response, error)
	ZrangeByScoreBatch(key string, min string, max string, offset int, count int) (*Response, error)
	//ZrevrangeByScore(key string, max float64, min float64, offset int, count int) (*Response, error)
	//ZrangeByScoreWithScores(key string, min float64, max float64) (*Response, error)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64) (*Response, error)
	//ZrangeByScoreWithScoresBatch(key string, min float64, max float64, offset int, count int) (*Response, error)
	//Zcore(key, max, min string, offset, count int) (*Response, error)
	ZrangeByScoreWithScores(key, min, max string) (*Response, error)
	ZrevrangeByScoreWithScores(key, max, min string) (*Response, error)
	ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) (*Response, error)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64, offset int, count int) (*Response, error)
	ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) (*Response, error)
	ZremrangeByRank(key string, start, end int64) (*Response, error)
	//ZremrangeByScore(key string, start float64, end float64) (*Response, error)
	ZremrangeByScore(key, start, end string) (*Response, error)
	Zlexcount(key, min, max string) (*Response, error)
	ZrangeByLex(key, min, max string) (*Response, error)
	ZrangeByLexBatch(key, min, max string, offset, count int) (*Response, error)
	ZrevrangeByLex(key, max, min string) (*Response, error)
	ZrevrangeByLexBatch(key, max, min string, offset, count int) (*Response, error)
	ZremrangeByLex(key, min, max string) (*Response, error)
	Linsert(key string, where ListOption, pivot, value string) (*Response, error)
	Lpushx(key string, String ...string) (*Response, error)
	Rpushx(key string, String ...string) (*Response, error)
	//Brpop(timeout int, key string) (*Response, error)
	//Del(key string) (*Response, error)
	Echo(str string) (*Response, error)
	Move(key string, dbIndex int) (*Response, error)
	Bitcount(key string) (*Response, error)
	BitcountRange(key string, start int64, end int64) (*Response, error)
	//Bitpos(key string, value bool) (*Response, error)
	Bitpos(key string, value bool, params ...BitPosParams) (*Response, error)
	//Hscan(key string, cursor string) (ScanResult, error)
	Hscan(key, cursor string, params ...ScanParams) (*Response, error)
	//Sscan(key string, cursor string) (ScanResult, error)
	Sscan(key, cursor string, params ...ScanParams) (*Response, error)
	//Zscan(key string, cursor string) (ScanResult, error)
	Zscan(key, cursor string, params ...ScanParams) (*Response, error)
	Pfadd(key string, elements ...string) (*Response, error)
	//Pfcount(key string) (*Response, error)

	// Geo Commands
	Geoadd(key string, longitude, latitude float64, member string) (*Response, error)
	GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (*Response, error)
	//Geodist(key string, member1, member2 string) (*Response, error)
	Geodist(key string, member1, member2 string, unit ...GeoUnit) (*Response, error)
	Geohash(key string, members ...string) (*Response, error)
	Geopos(key string, members ...string) (*Response, error)
	//Georadius(key string, longitude float64, latitude float64, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) (*Response, error)
	//GeoradiusByMember(key string, member string, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) (*Response, error)
	Bitfield(key string, arguments ...string) (*Response, error)
}

type BasicRedisPipeline interface {
	Bgrewriteaof() (*Response, error)
	Bgsave() (*Response, error)
	ConfigGet(pattern string) (*Response, error)
	ConfigSet(parameter, value string) (*Response, error)
	ConfigResetStat() (*Response, error)
	Save() (*Response, error)
	Lastsave() (*Response, error)
	FlushDB() (*Response, error)
	FlushAll() (*Response, error)
	Info() (*Response, error)
	Time() (*Response, error)
	DbSize() (*Response, error)
	Shutdown() (*Response, error)
	Ping() (*Response, error)
	Select(index int) (*Response, error)
}

type MultiKeyCommandsPipeline interface {
	Del(keys ...string) (*Response, error)
	Exists(keys ...string) (*Response, error)
	BlpopTimout(timeout int, keys ...string) (*Response, error)
	BrpopTimout(timeout int, keys ...string) (*Response, error)
	Blpop(args ...string) (*Response, error)
	Brpop(args ...string) (*Response, error)
	Keys(pattern string) (*Response, error)
	Mget(keys ...string) (*Response, error)
	Mset(keysvalues ...string) (*Response, error)
	Msetnx(keysvalues ...string) (*Response, error)
	Rename(oldkey, newkey string) (*Response, error)
	Renamenx(oldkey, newkey string) (*Response, error)
	Rpoplpush(srckey, dstkey string) (*Response, error)
	Sdiff(keys ...string) (*Response, error)

	Sdiffstore(dstkey string, keys ...string) (*Response, error)
	Sinter(keys ...string) (*Response, error)
	Sinterstore(dstkey string, keys ...string) (*Response, error)
	Smove(srckey, dstkey, member string) (*Response, error)
	SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (*Response, error)
	//Sort(key, dstkey string) (*Response, error)
	Sunion(keys ...string) (*Response, error)
	Sunionstore(dstkey string, keys ...string) (*Response, error)
	Watch(keys ...string) (*Response, error)
	Unwatch() (*Response, error)
	Zinterstore(dstkey string, sets ...string) (*Response, error)
	ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (*Response, error)
	Zunionstore(dstkey string, sets ...string) (*Response, error)
	ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (*Response, error)
	Brpoplpush(source, destination string, timeout int) (*Response, error)
	Publish(channel, message string) (*Response, error)
	RandomKey() (*Response, error)
	Bitop(op BitOP, destKey string, srcKeys ...string) (*Response, error)
	//Scan(cursor string) (ScanResult, error)
	Scan(cursor string, params ...ScanParams) (*Response, error)
	Pfmerge(destkey string, sourcekeys ...string) (*Response, error)
	Pfcount(keys ...string) (*Response, error)
}

type ClusterPipeline interface {
	ClusterNodes() (*Response, error)
	ClusterMeet(ip string, port int) (*Response, error)
	ClusterAddSlots(slots ...int) (*Response, error)
	ClusterDelSlots(slots ...int) (*Response, error)
	ClusterInfo() (*Response, error)
	ClusterGetKeysInSlot(slot int, count int) (*Response, error)
	ClusterSetSlotNode(slot int, nodeId string) (*Response, error)
	ClusterSetSlotMigrating(slot int, nodeId string) (*Response, error)
	ClusterSetSlotImporting(slot int, nodeId string) (*Response, error)
}

type ScriptingCommandsPipeline interface {
	Eval(script string, keyCount int, params ...string) (*Response, error)
	Evalsha(sha1 string, keyCount int, params ...string) (*Response, error)
}

type Response struct {
	response interface{}

	building bool
	built    bool
	set      bool

	builder    Builder
	data       interface{}
	dependency *Response
}

func NewResponse() *Response {
	return &Response{
		building: false,
		built:    false,
		set:      false,
	}
}

func (r *Response) Set(data interface{}) {
	r.data = data
	r.set = true
}

func (r *Response) Get() (interface{}, error) {
	if r.dependency != nil && r.dependency.set && !r.dependency.built {
		err := r.dependency.build()
		if err != nil {
			return nil, err
		}
	}
	if !r.set {
		return nil, errors.New("please close pipeline or multi block before calling this method")
	}
	if !r.built {
		err := r.build()
		if err != nil {
			return nil, err
		}
	}
	return r.response, nil
}

func (r *Response) setDependency(dependency *Response) {
	r.dependency = dependency
}

func (r *Response) build() error {
	if r.building {
		return nil
	}
	r.building = true
	defer func() {
		r.building = false
		r.built = true
	}()
	if r.data != nil {
		r.response = r.builder.build(r.data)
	}
	r.data = nil
	return nil
}

type Transaction struct {
	*MultiKeyPipelineBase
	inTransaction bool
}

func NewTransaction(client *Client) *Transaction {
	return &Transaction{MultiKeyPipelineBase: NewMultiKeyPipelineBase(client)}
}

func (t *Transaction) Clear() (string, error) {
	if t.inTransaction {
		return t.Discard()
	}
	return "", nil
}

func (t *Transaction) Exec() ([]interface{}, error) {
	err := t.client.Exec()
	if err != nil {
		return nil, err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return nil, err
	}
	t.inTransaction = false
	reply, err := t.client.getObjectMultiBulkReply()
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

func (t *Transaction) ExecGetResponse() ([]*Response, error) {
	err := t.client.Exec()
	if err != nil {
		return nil, err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return nil, err
	}
	t.inTransaction = false
	reply, err := t.client.getObjectMultiBulkReply()
	if err != nil {
		return nil, err
	}
	result := make([]*Response, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

func (t *Transaction) Discard() (string, error) {
	err := t.client.Discard()
	if err != nil {
		return "", err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return "", err
	}
	t.inTransaction = false
	t.clean()
	return t.client.getStatusCodeReply()
}

func (t *Transaction) clean() {
	t.pipelinedResponses = make([]*Response, 0)
}

type Pipeline struct {
	*MultiKeyPipelineBase
	//client             *Client
	//pipelinedResponses []*Response
}

func NewPipeline(client *Client) *Pipeline {
	return &Pipeline{MultiKeyPipelineBase: NewMultiKeyPipelineBase(client)}
}

//func (p *Pipeline) generateResponse(data interface{}) *Response {
//	size := len(p.pipelinedResponses)
//	if size == 0 {
//		return nil
//	}
//	r := p.pipelinedResponses[0]
//	r.Set(data)
//	if size == 1 {
//		p.pipelinedResponses = make([]*Response, 0)
//	} else {
//		p.pipelinedResponses = p.pipelinedResponses[1:]
//	}
//	return r
//}

//func (p *Pipeline) getResponse(builder Builder) *Response {
//	response := NewResponse()
//	response.builder = builder
//	p.pipelinedResponses = append(p.pipelinedResponses, response)
//	return response
//}

func (p *Pipeline) Sync() error {
	if len(p.pipelinedResponses) == 0 {
		return nil
	}
	all, err := p.client.Connection.getAll()
	if err != nil {
		return err
	}
	for _, a := range all.([]interface{}) {
		p.generateResponse(a)
	}
	return nil
}

type Queable struct {
	pipelinedResponses []*Response
}

func NewQueable() *Queable {
	return &Queable{pipelinedResponses: make([]*Response, 0)}
}

func (q *Queable) clean() {
	q.pipelinedResponses = make([]*Response, 0)
}

func (q *Queable) generateResponse(data interface{}) *Response {
	size := len(q.pipelinedResponses)
	if size == 0 {
		return nil
	}
	r := q.pipelinedResponses[0]
	r.Set(data)
	if size == 1 {
		q.pipelinedResponses = make([]*Response, 0)
	} else {
		q.pipelinedResponses = q.pipelinedResponses[1:]
	}
	return r
}

func (q *Queable) getResponse(builder Builder) *Response {
	response := NewResponse()
	response.builder = builder
	q.pipelinedResponses = append(q.pipelinedResponses, response)
	return response
}

func (q *Queable) hasPipelinedResponse() bool {
	return q.getPipelinedResponseLength() > 0
}

func (q *Queable) getPipelinedResponseLength() int {
	return len(q.pipelinedResponses)
}

type MultiKeyPipelineBase struct {
	*Queable
	client *Client
}

func NewMultiKeyPipelineBase(client *Client) *MultiKeyPipelineBase {
	return &MultiKeyPipelineBase{Queable: NewQueable(), client: client}
}

//<editor-fold desc="basicpipeline">
func (p *MultiKeyPipelineBase) Bgrewriteaof() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Bgsave() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) ConfigGet(pattern string) (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) ConfigSet(parameter, value string) (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) ConfigResetStat() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Save() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Lastsave() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) FlushDB() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) FlushAll() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Info() (*Response, error) {
	err := p.client.Info()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *MultiKeyPipelineBase) Time() (*Response, error) {
	err := p.client.Time()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *MultiKeyPipelineBase) DbSize() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Shutdown() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Ping() (*Response, error) {
	panic("implement me")
}

func (p *MultiKeyPipelineBase) Select(index int) (*Response, error) {
	panic("implement me")
}

//</editor-fold>

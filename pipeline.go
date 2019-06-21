package godis

import "errors"

//RedisPipeline
type RedisPipeline interface {
	Set(key, value string) (*response, error)
	SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (*response, error)
	SetWithParams(key, value, nxxx string) (*response, error)
	Get(key string) (*response, error)
	Persist(key string) (*response, error)
	Type(key string) (*response, error)
	Expire(key string, seconds int) (*response, error)
	Pexpire(key string, milliseconds int64) (*response, error)
	ExpireAt(key string, unixtime int64) (*response, error)
	PexpireAt(key string, millisecondsTimestamp int64) (*response, error)
	Ttl(key string) (*response, error)
	Pttl(key string) (*response, error)
	SetbitWithBool(key string, offset int64, value bool) (*response, error)
	Setbit(key string, offset int64, value string) (*response, error)
	Getbit(key string, offset int64) (*response, error)
	Setrange(key string, offset int64, value string) (*response, error)
	Getrange(key string, startOffset, endOffset int64) (*response, error)
	GetSet(key, value string) (*response, error)

	Setnx(key, value string) (*response, error)
	Setex(key string, seconds int, value string) (*response, error)
	Psetex(key string, milliseconds int64, value string) (*response, error)
	DecrBy(key string, decrement int64) (*response, error)
	Decr(key string) (*response, error)
	IncrBy(key string, increment int64) (*response, error)
	IncrByFloat(key string, increment float64) (*response, error)
	Incr(key string) (*response, error)
	Append(key, value string) (*response, error)
	Substr(key string, start, end int) (*response, error)

	Hset(key, field string, value string) (*response, error)
	Hget(key, field string) (*response, error)
	Hsetnx(key, field, value string) (*response, error)
	Hmset(key string, hash map[string]string) (*response, error)
	Hmget(key string, fields ...string) (*response, error)
	HincrBy(key, field string, value int64) (*response, error)
	HincrByFloat(key, field string, value float64) (*response, error)
	Hexists(key, field string) (*response, error)
	Hdel(key string, fields ...string) (*response, error)
	Hlen(key string) (*response, error)
	Hkeys(key string) (*response, error)
	Hvals(key string) (*response, error)
	HgetAll(key string) (*response, error)

	Rpush(key string, strings ...string) (*response, error)
	Lpush(key string, strings ...string) (*response, error)
	Llen(key string) (*response, error)
	Lrange(key string, start, stop int64) (*response, error)
	Ltrim(key string, start, stop int64) (*response, error)
	Lindex(key string, index int64) (*response, error)
	Lset(key string, index int64, value string) (*response, error)
	Lrem(key string, count int64, value string) (*response, error)
	Lpop(key string) (*response, error)
	Rpop(key string) (*response, error)
	Sadd(key string, members ...string) (*response, error)
	Smembers(key string) (*response, error)
	Srem(key string, members ...string) (*response, error)
	Spop(key string) (*response, error)
	SpopBatch(key string, count int64) (*response, error)

	Scard(key string) (*response, error)
	Sismember(key string, member string) (*response, error)
	Srandmember(key string) (*response, error)
	SrandmemberBatch(key string, count int) (*response, error)
	Strlen(key string) (*response, error)

	Zadd(key string, score float64, member string, params ...ZAddParams) (*response, error)
	ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (*response, error)
	Zrange(key string, start, end int64) (*response, error)
	Zrem(key string, member ...string) (*response, error)
	Zincrby(key string, score float64, member string, params ...ZAddParams) (*response, error)
	Zrank(key, member string) (*response, error)
	Zrevrank(key, member string) (*response, error)
	Zrevrange(key string, start, end int64) (*response, error)
	ZrangeWithScores(key string, start, end int64) (*response, error)
	ZrevrangeWithScores(key string, start, end int64) (*response, error)
	Zcard(key string) (*response, error)
	Zscore(key, member string) (*response, error)
	Sort(key string, sortingParameters ...SortingParams) (*response, error)
	Zcount(key string, min string, max string) (*response, error)
	ZrangeByScore(key string, min string, max string) (*response, error)
	ZrevrangeByScore(key string, max string, min string) (*response, error)
	ZrangeByScoreBatch(key string, min string, max string, offset int, count int) (*response, error)
	ZrangeByScoreWithScores(key, min, max string) (*response, error)
	ZrevrangeByScoreWithScores(key, max, min string) (*response, error)
	ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) (*response, error)
	ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) (*response, error)
	ZremrangeByRank(key string, start, end int64) (*response, error)
	ZremrangeByScore(key, start, end string) (*response, error)
	Zlexcount(key, min, max string) (*response, error)
	ZrangeByLex(key, min, max string) (*response, error)
	ZrangeByLexBatch(key, min, max string, offset, count int) (*response, error)
	ZrevrangeByLex(key, max, min string) (*response, error)
	ZrevrangeByLexBatch(key, max, min string, offset, count int) (*response, error)
	ZremrangeByLex(key, min, max string) (*response, error)
	Linsert(key string, where ListOption, pivot, value string) (*response, error)
	Lpushx(key string, String ...string) (*response, error)
	Rpushx(key string, String ...string) (*response, error)
	Echo(str string) (*response, error)
	Move(key string, dbIndex int) (*response, error)
	Bitcount(key string) (*response, error)
	BitcountRange(key string, start int64, end int64) (*response, error)
	Bitpos(key string, value bool, params ...BitPosParams) (*response, error)
	Hscan(key, cursor string, params ...ScanParams) (*response, error)
	Sscan(key, cursor string, params ...ScanParams) (*response, error)
	Zscan(key, cursor string, params ...ScanParams) (*response, error)
	Pfadd(key string, elements ...string) (*response, error)

	Geoadd(key string, longitude, latitude float64, member string) (*response, error)
	GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (*response, error)
	Geodist(key string, member1, member2 string, unit ...GeoUnit) (*response, error)
	Geohash(key string, members ...string) (*response, error)
	Geopos(key string, members ...string) (*response, error)
	Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) (*response, error)
	GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) (*response, error)
	Bitfield(key string, arguments ...string) (*response, error)
}

//BasicRedisPipeline
type BasicRedisPipeline interface {
	Bgrewriteaof() (*response, error)
	Bgsave() (*response, error)
	ConfigGet(pattern string) (*response, error)
	ConfigSet(parameter, value string) (*response, error)
	ConfigResetStat() (*response, error)
	Save() (*response, error)
	Lastsave() (*response, error)
	FlushDB() (*response, error)
	FlushAll() (*response, error)
	Info() (*response, error)
	Time() (*response, error)
	DbSize() (*response, error)
	Shutdown() (*response, error)
	Ping() (*response, error)
	Select(index int) (*response, error)
}

//MultiKeyCommandsPipeline
type MultiKeyCommandsPipeline interface {
	Del(keys ...string) (*response, error)
	Exists(keys ...string) (*response, error)
	BlpopTimout(timeout int, keys ...string) (*response, error)
	BrpopTimout(timeout int, keys ...string) (*response, error)
	Blpop(args ...string) (*response, error)
	Brpop(args ...string) (*response, error)
	Keys(pattern string) (*response, error)
	Mget(keys ...string) (*response, error)
	Mset(keysvalues ...string) (*response, error)
	Msetnx(keysvalues ...string) (*response, error)
	Rename(oldkey, newkey string) (*response, error)
	Renamenx(oldkey, newkey string) (*response, error)
	Rpoplpush(srckey, dstkey string) (*response, error)
	Sdiff(keys ...string) (*response, error)

	Sdiffstore(dstkey string, keys ...string) (*response, error)
	Sinter(keys ...string) (*response, error)
	Sinterstore(dstkey string, keys ...string) (*response, error)
	Smove(srckey, dstkey, member string) (*response, error)
	SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (*response, error)
	Sunion(keys ...string) (*response, error)
	Sunionstore(dstkey string, keys ...string) (*response, error)
	Watch(keys ...string) (*response, error)
	Zinterstore(dstkey string, sets ...string) (*response, error)
	ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error)
	Zunionstore(dstkey string, sets ...string) (*response, error)
	ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error)
	Brpoplpush(source, destination string, timeout int) (*response, error)
	Publish(channel, message string) (*response, error)
	RandomKey() (*response, error)
	Bitop(op BitOP, destKey string, srcKeys ...string) (*response, error)
	Pfmerge(destkey string, sourcekeys ...string) (*response, error)
	Pfcount(keys ...string) (*response, error)
}

//ClusterPipeline
type ClusterPipeline interface {
	ClusterNodes() (*response, error)
	ClusterMeet(ip string, port int) (*response, error)
	ClusterAddSlots(slots ...int) (*response, error)
	ClusterDelSlots(slots ...int) (*response, error)
	ClusterInfo() (*response, error)
	ClusterGetKeysInSlot(slot int, count int) (*response, error)
	ClusterSetSlotNode(slot int, nodeId string) (*response, error)
	ClusterSetSlotMigrating(slot int, nodeId string) (*response, error)
	ClusterSetSlotImporting(slot int, nodeId string) (*response, error)
}

//ScriptingCommandsPipeline
type ScriptingCommandsPipeline interface {
	Eval(script string, keyCount int, params ...string) (*response, error)
	Evalsha(sha1 string, keyCount int, params ...string) (*response, error)
}

type response struct {
	response interface{}

	building bool
	built    bool
	set      bool

	builder    Builder
	data       interface{}
	dependency *response
}

func newResponse() *response {
	return &response{
		building: false,
		built:    false,
		set:      false,
	}
}

func (r *response) Set(data interface{}) {
	r.data = data
	r.set = true
}

func (r *response) Get() (interface{}, error) {
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

func (r *response) setDependency(dependency *response) {
	r.dependency = dependency
}

func (r *response) build() error {
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

type transaction struct {
	*multiKeyPipelineBase
	inTransaction bool
}

func newTransaction(client *Client) *transaction {
	base := newMultiKeyPipelineBase(client)
	base.getClient = func(key string) *Client {
		return client
	}
	return &transaction{multiKeyPipelineBase: base}
}

func (t *transaction) Clear() (string, error) {
	if t.inTransaction {
		return t.Discard()
	}
	return "", nil
}

func (t *transaction) Exec() ([]interface{}, error) {
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

func (t *transaction) ExecGetResponse() ([]*response, error) {
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
	result := make([]*response, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

func (t *transaction) Discard() (string, error) {
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

func (t *transaction) clean() {
	t.pipelinedResponses = make([]*response, 0)
}

type pipeline struct {
	*multiKeyPipelineBase
}

func newPipeline(client *Client) *pipeline {
	base := newMultiKeyPipelineBase(client)
	base.getClient = func(key string) *Client {
		return client
	}
	return &pipeline{multiKeyPipelineBase: base}
}

func (p *pipeline) Sync() error {
	if len(p.pipelinedResponses) == 0 {
		return nil
	}
	all, err := p.client.connection.getAll()
	if err != nil {
		return err
	}
	for _, a := range all.([]interface{}) {
		p.generateResponse(a)
	}
	return nil
}

type queable struct {
	pipelinedResponses []*response
}

func newQueable() *queable {
	return &queable{pipelinedResponses: make([]*response, 0)}
}

func (q *queable) clean() {
	q.pipelinedResponses = make([]*response, 0)
}

func (q *queable) generateResponse(data interface{}) *response {
	size := len(q.pipelinedResponses)
	if size == 0 {
		return nil
	}
	r := q.pipelinedResponses[0]
	r.Set(data)
	if size == 1 {
		q.pipelinedResponses = make([]*response, 0)
	} else {
		q.pipelinedResponses = q.pipelinedResponses[1:]
	}
	return r
}

func (q *queable) getResponse(builder Builder) *response {
	response := newResponse()
	response.builder = builder
	q.pipelinedResponses = append(q.pipelinedResponses, response)
	return response
}

func (q *queable) hasPipelinedResponse() bool {
	return q.getPipelinedResponseLength() > 0
}

func (q *queable) getPipelinedResponseLength() int {
	return len(q.pipelinedResponses)
}

type multiKeyPipelineBase struct {
	*queable
	client *Client

	getClient func(key string) *Client
}

func newMultiKeyPipelineBase(client *Client) *multiKeyPipelineBase {
	return &multiKeyPipelineBase{queable: newQueable(), client: client}
}

//<editor-fold desc="basicpipeline">

//Bgrewriteaof
func (p *multiKeyPipelineBase) Bgrewriteaof() (*response, error) {
	err := p.client.Bgrewriteaof()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Bgsave() (*response, error) {
	err := p.client.Bgsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ConfigGet(pattern string) (*response, error) {
	err := p.client.ConfigGet(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) ConfigSet(parameter, value string) (*response, error) {
	err := p.client.ConfigSet(parameter, value)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ConfigResetStat() (*response, error) {
	err := p.client.ConfigResetStat()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Save() (*response, error) {
	err := p.client.Save()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Lastsave() (*response, error) {
	err := p.client.Lastsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) FlushDB() (*response, error) {
	err := p.client.FlushDB()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) FlushAll() (*response, error) {
	err := p.client.FlushAll()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Info() (*response, error) {
	err := p.client.Info()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Time() (*response, error) {
	err := p.client.Time()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) DbSize() (*response, error) {
	err := p.client.DbSize()
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Shutdown() (*response, error) {
	err := p.client.Shutdown()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Ping() (*response, error) {
	err := p.client.Ping()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Select(index int) (*response, error) {
	err := p.client.Select(index)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

//</editor-fold>

//<editor-fold desc="multikeypipeline">

//Del
func (p *multiKeyPipelineBase) Del(keys ...string) (*response, error) {
	err := p.client.Del(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Exists(keys ...string) (*response, error) {
	err := p.client.Exists(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) BlpopTimout(timeout int, keys ...string) (*response, error) {
	err := p.client.BlpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) BrpopTimout(timeout int, keys ...string) (*response, error) {
	err := p.client.BrpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Blpop(args ...string) (*response, error) {
	err := p.client.Blpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Brpop(args ...string) (*response, error) {
	err := p.client.Brpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Keys(pattern string) (*response, error) {
	err := p.client.Keys(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Mget(keys ...string) (*response, error) {
	err := p.client.Mget(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Mset(keysvalues ...string) (*response, error) {
	err := p.client.Mset(keysvalues...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Msetnx(keysvalues ...string) (*response, error) {
	err := p.client.Msetnx(keysvalues...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Rename(oldkey, newkey string) (*response, error) {
	err := p.client.Rename(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Renamenx(oldkey, newkey string) (*response, error) {
	err := p.client.Renamenx(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Rpoplpush(srckey, dstkey string) (*response, error) {
	err := p.client.RpopLpush(srckey, dstkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sdiff(keys ...string) (*response, error) {
	err := p.client.Sdiff(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sdiffstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.Sdiffstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sinter(keys ...string) (*response, error) {
	err := p.client.Sinter(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sinterstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.Sinterstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Smove(srckey, dstkey, member string) (*response, error) {
	err := p.client.Smove(srckey, dstkey, member)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (*response, error) {
	err := p.client.SortMulti(key, dstkey, sortingParameters...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sunion(keys ...string) (*response, error) {
	err := p.client.Sunion(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Sunionstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.Sunionstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Watch(keys ...string) (*response, error) {
	err := p.client.Watch(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Zinterstore(dstkey string, sets ...string) (*response, error) {
	err := p.client.Zinterstore(dstkey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error) {
	err := p.client.ZinterstoreWithParams(dstkey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Zunionstore(dstkey string, sets ...string) (*response, error) {
	err := p.client.Zunionstore(dstkey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error) {
	err := p.client.ZunionstoreWithParams(dstkey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Brpoplpush(source, destination string, timeout int) (*response, error) {
	err := p.client.Brpoplpush(source, destination, timeout)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) Publish(channel, message string) (*response, error) {
	err := p.client.Publish(channel, message)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) RandomKey() (*response, error) {
	err := p.client.RandomKey()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Bitop(op BitOP, destKey string, srcKeys ...string) (*response, error) {
	err := p.client.Bitop(op, destKey, srcKeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

func (p *multiKeyPipelineBase) Pfmerge(destkey string, sourcekeys ...string) (*response, error) {
	err := p.client.Pfmerge(destkey, sourcekeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Pfcount(keys ...string) (*response, error) {
	err := p.client.Pfcount(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(INT64_BUILDER), nil
}

//</editor-fold>

//<editor-fold desc="cluster pipeline">

//ClusterNodes
func (p *multiKeyPipelineBase) ClusterNodes() (*response, error) {
	err := p.client.ClusterNodes()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterMeet(ip string, port int) (*response, error) {
	err := p.client.ClusterMeet(ip, port)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterAddSlots(slots ...int) (*response, error) {
	err := p.client.ClusterAddSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterDelSlots(slots ...int) (*response, error) {
	err := p.client.ClusterDelSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterInfo() (*response, error) {
	err := p.client.ClusterInfo()
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterGetKeysInSlot(slot int, count int) (*response, error) {
	err := p.client.ClusterGetKeysInSlot(slot, count)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_ARRAY_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterSetSlotNode(slot int, nodeId string) (*response, error) {
	err := p.client.ClusterSetSlotNode(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterSetSlotMigrating(slot int, nodeId string) (*response, error) {
	err := p.client.ClusterSetSlotMigrating(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) ClusterSetSlotImporting(slot int, nodeId string) (*response, error) {
	err := p.client.ClusterSetSlotImporting(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

//</editor-fold>

//<editor-fold desc="scripting pipeline">

//Eval
func (p *multiKeyPipelineBase) Eval(script string, keyCount int, params ...string) (*response, error) {
	err := p.getClient(script).Eval(script, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

func (p *multiKeyPipelineBase) Evalsha(sha1 string, keyCount int, params ...string) (*response, error) {
	err := p.getClient(sha1).Evalsha(sha1, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(STRING_BUILDER), nil
}

//</editor-fold>

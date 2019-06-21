package godis

import (
	"errors"
	"fmt"
	"strings"
)

//RedisCommands
type RedisCommands interface {
	Set(key, value string) (string, error)
	SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (string, error)
	SetWithParams(key, value, nxxx string) (string, error)
	Get(key string) (string, error)
	//Exists(key string) ([]string, error)
	Persist(key string) (int64, error)
	Type(key string) (string, error)
	Expire(key string, seconds int) (int64, error)
	Pexpire(key string, milliseconds int64) (int64, error)
	ExpireAt(key string, unixtime int64) (int64, error)
	PexpireAt(key string, millisecondsTimestamp int64) (int64, error)
	Ttl(key string) (int64, error)
	Pttl(key string) (int64, error)
	SetbitWithBool(key string, offset int64, value bool) (bool, error)
	Setbit(key string, offset int64, value string) (bool, error)
	Getbit(key string, offset int64) (bool, error)
	Setrange(key string, offset int64, value string) (int64, error)
	Getrange(key string, startOffset, endOffset int64) (string, error)
	GetSet(key, value string) (string, error)

	Setnx(key, value string) (int64, error)
	Setex(key string, seconds int, value string) (string, error)
	Psetex(key string, milliseconds int64, value string) (string, error)
	DecrBy(key string, decrement int64) (int64, error)
	Decr(key string) (int64, error)
	IncrBy(key string, increment int64) (int64, error)
	IncrByFloat(key string, increment float64) (float64, error)
	Incr(key string) (int64, error)
	Append(key, value string) (int64, error)
	Substr(key string, start, end int) (string, error)

	Hset(key, field string, value string) (int64, error)
	Hget(key, field string) (string, error)
	Hsetnx(key, field, value string) (int64, error)
	Hmset(key string, hash map[string]string) (string, error)
	Hmget(key string, fields ...string) ([]string, error)
	HincrBy(key, field string, value int64) (int64, error)
	HincrByFloat(key, field string, value float64) (float64, error)
	Hexists(key, field string) (bool, error)
	Hdel(key string, fields ...string) (int64, error)
	Hlen(key string) (int64, error)
	Hkeys(key string) ([]string, error)
	Hvals(key string) ([]string, error)
	HgetAll(key string) (map[string]string, error)

	Rpush(key string, strings ...string) (int64, error)
	Lpush(key string, strings ...string) (int64, error)
	Llen(key string) (int64, error)
	Lrange(key string, start, stop int64) ([]string, error)
	Ltrim(key string, start, stop int64) (string, error)
	Lindex(key string, index int64) (string, error)
	Lset(key string, index int64, value string) (string, error)
	Lrem(key string, count int64, value string) (int64, error)
	Lpop(key string) (string, error)
	Rpop(key string) (string, error)
	Sadd(key string, members ...string) (int64, error)
	Smembers(key string) ([]string, error)
	Srem(key string, members ...string) (int64, error)
	Spop(key string) (string, error)
	SpopBatch(key string, count int64) ([]string, error)

	Scard(key string) (int64, error)
	Sismember(key string, member string) (bool, error)
	Srandmember(key string) (string, error)
	SrandmemberBatch(key string, count int) ([]string, error)
	Strlen(key string) (int64, error)

	//Zadd(key string, score float64, member string) (int64, error)
	Zadd(key string, score float64, member string, params ...ZAddParams) (int64, error)
	//Zadd(key string, scoreMembers map[string]float64) (int64, error)
	ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (int64, error)
	Zrange(key string, start, end int64) ([]string, error)
	Zrem(key string, member ...string) (int64, error)
	//Zincrby(key string, score float64, member string) (float64, error)
	Zincrby(key string, score float64, member string, params ...ZAddParams) (float64, error)
	Zrank(key, member string) (int64, error)
	Zrevrank(key, member string) (int64, error)
	Zrevrange(key string, start, end int64) ([]string, error)
	ZrangeWithScores(key string, start, end int64) ([]Tuple, error)
	ZrevrangeWithScores(key string, start, end int64) ([]Tuple, error)
	Zcard(key string) (int64, error)
	Zscore(key, member string) (float64, error)
	//Sort(key string) ([]string, error)
	Sort(key string, sortingParameters ...SortingParams) ([]string, error)
	//Zcount(key string, min float64, max float64) (int64, error)
	Zcount(key string, min string, max string) (int64, error)
	//ZrangeByScore(key string, min float64, max float64) ([]string, error)
	ZrangeByScore(key string, min string, max string) ([]string, error)
	//ZrevrangeByScore(key string, max float64, min float64) ([]string, error)
	//ZrangeByScoreBatch(key string, min float64, max float64, offset int, count int) ([]string, error)
	ZrevrangeByScore(key string, max string, min string) ([]string, error)
	ZrangeByScoreBatch(key string, min string, max string, offset int, count int) ([]string, error)
	//ZrevrangeByScore(key string, max float64, min float64, offset int, count int) ([]string, error)
	//ZrangeByScoreWithScores(key string, min float64, max float64) ([]Tuple, error)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64) ([]Tuple, error)
	//ZrangeByScoreWithScoresBatch(key string, min float64, max float64, offset int, count int) ([]Tuple, error)
	//Zcore(key, max, min string, offset, count int) ([]string, error)
	ZrangeByScoreWithScores(key, min, max string) ([]Tuple, error)
	ZrevrangeByScoreWithScores(key, max, min string) ([]Tuple, error)
	ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) ([]Tuple, error)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64, offset int, count int) ([]Tuple, error)
	ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) ([]Tuple, error)
	ZremrangeByRank(key string, start, end int64) (int64, error)
	//ZremrangeByScore(key string, start float64, end float64) (int64, error)
	ZremrangeByScore(key, start, end string) (int64, error)
	Zlexcount(key, min, max string) (int64, error)
	ZrangeByLex(key, min, max string) ([]string, error)
	ZrangeByLexBatch(key, min, max string, offset, count int) ([]string, error)
	ZrevrangeByLex(key, max, min string) ([]string, error)
	ZrevrangeByLexBatch(key, max, min string, offset, count int) ([]string, error)
	ZremrangeByLex(key, min, max string) (int64, error)
	Linsert(key string, where ListOption, pivot, value string) (int64, error)
	Lpushx(key string, String ...string) (int64, error)
	Rpushx(key string, String ...string) (int64, error)
	//Brpop(timeout int, key string) ([]string, error)
	//Del(key string) (int64, error)
	Echo(str string) (string, error)
	Move(key string, dbIndex int) (int64, error)
	Bitcount(key string) (int64, error)
	BitcountRange(key string, start int64, end int64) (int64, error)
	//Bitpos(key string, value bool) (int64, error)
	Bitpos(key string, value bool, params ...BitPosParams) (int64, error)
	//Hscan(key string, cursor string) (ScanResult, error)
	Hscan(key, cursor string, params ...ScanParams) (*ScanResult, error)
	//Sscan(key string, cursor string) (ScanResult, error)
	Sscan(key, cursor string, params ...ScanParams) (*ScanResult, error)
	//Zscan(key string, cursor string) (ScanResult, error)
	Zscan(key, cursor string, params ...ScanParams) (*ScanResult, error)
	Pfadd(key string, elements ...string) (int64, error)
	//Pfcount(key string) (int64, error)

	// Geo Commands
	Geoadd(key string, longitude, latitude float64, member string) (int64, error)
	GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error)
	//Geodist(key string, member1, member2 string) (float64, error)
	Geodist(key string, member1, member2 string, unit ...GeoUnit) (float64, error)
	Geohash(key string, members ...string) ([]string, error)
	Geopos(key string, members ...string) ([]*GeoCoordinate, error)
	//Georadius(key string, longitude float64, latitude float64, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error)
	//GeoradiusByMember(key string, member string, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error)
	Bitfield(key string, arguments ...string) ([]int64, error)
}

//ZAddParams
type ZAddParams struct {
	XX     bool
	NX     bool
	CH     bool
	params map[string]string
}

//GetByteParams
func (p ZAddParams) GetByteParams(key []byte, args ...[]byte) [][]byte {
	arr := make([][]byte, 0)
	arr = append(arr, key)
	if p.Contains("XX") {
		arr = append(arr, []byte("XX"))
	}
	if p.Contains("NX") {
		arr = append(arr, []byte("NX"))
	}
	if p.Contains("CH") {
		arr = append(arr, []byte("CH"))
	}
	for _, a := range args {
		arr = append(arr, a)
	}
	return arr
}

//Contains
func (p ZAddParams) Contains(key string) bool {
	_, ok := p.params[key]
	return ok
}

//BitPosParams
type BitPosParams struct {
	params [][]byte
}

//SortingParams
type SortingParams struct {
	params [][]byte
}

//ScanParams
type ScanParams struct {
	params map[keyword][]byte
}

//NewScanParams
func NewScanParams() *ScanParams {
	return &ScanParams{params: make(map[keyword][]byte)}
}

//GetParams
func (s ScanParams) GetParams() [][]byte {
	arr := make([][]byte, 0)
	for k, v := range s.params {
		arr = append(arr, k.GetRaw())
		arr = append(arr, []byte(v))
	}
	return arr
}

//Match
func (s ScanParams) Match() string {
	if v, ok := s.params[KEYWORD_MATCH]; !ok {
		return ""
	} else {
		return string(v)
	}
}

//Count
func (s ScanParams) Count() int {
	if v, ok := s.params[KEYWORD_COUNT]; !ok {
		return 0
	} else {
		return int(ByteArrayToInt64(v))
	}
}

//ListOption
type ListOption struct {
	Name string
}

//GetRaw
func (l ListOption) GetRaw() []byte {
	return []byte(l.Name)
}

//NewListOption
func NewListOption(name string) ListOption {
	return ListOption{name}
}

var (
	ListOption_BEFORE = NewListOption("BEFORE")
	ListOption_AFTER  = NewListOption("AFTER")
)

//GeoUnit
type GeoUnit struct {
	Name string
}

//GetRaw
func (g GeoUnit) GetRaw() []byte {
	return []byte(g.Name)
}

//NewGeoUnit
func NewGeoUnit(name string) GeoUnit {
	return GeoUnit{name}
}

var (
	GEOUNIT_MI = NewGeoUnit("MI")
	GEOUNIT_M  = NewGeoUnit("M")
	GEOUNIT_KM = NewGeoUnit("KM")
	GEOUNIT_FT = NewGeoUnit("FT")
)

//GeoRadiusParam
type GeoRadiusParam struct {
	params map[string]interface{}
}

//GetParams
func (g GeoRadiusParam) GetParams(args [][]byte) [][]byte {
	arr := make([][]byte, 0)
	for _, a := range args {
		arr = append(arr, a)
	}

	if g.Contains("WITHCOORD") {
		arr = append(arr, []byte("WITHCOORD"))
	}
	if g.Contains("WITHDIST") {
		arr = append(arr, []byte("WITHDIST"))
	}

	if g.Contains("COUNT") {
		arr = append(arr, []byte("COUNT"))
		arr = append(arr, IntToByteArray(g.params["COUNT"].(int)))
	}

	if g.Contains("ASC") {
		arr = append(arr, []byte("ASC"))
	} else if g.Contains("DESC") {
		arr = append(arr, []byte("DESC"))
	}

	return arr
}

//Contains
func (g GeoRadiusParam) Contains(key string) bool {
	_, ok := g.params[key]
	return ok
}

//Tuple
type Tuple struct {
	element []byte
	score   float64
}

//GeoRadiusResponse
type GeoRadiusResponse struct {
	member     []byte
	distance   float64
	coordinate GeoCoordinate
}

//GeoCoordinate
type GeoCoordinate struct {
	longitude float64
	latitude  float64
}

//MultiKeyCommands
type MultiKeyCommands interface {
	Del(keys ...string) (int64, error)
	Exists(keys ...string) (int64, error)
	BlpopTimout(timeout int, keys ...string) ([]string, error)
	BrpopTimout(timeout int, keys ...string) ([]string, error)
	Blpop(args ...string) ([]string, error)
	Brpop(args ...string) ([]string, error)
	Keys(pattern string) ([]string, error)
	Mget(keys ...string) ([]string, error)
	Mset(keysvalues ...string) (string, error)
	Msetnx(keysvalues ...string) (int64, error)
	Rename(oldkey, newkey string) (string, error)
	Renamenx(oldkey, newkey string) (int64, error)
	Rpoplpush(srckey, dstkey string) (string, error)
	Sdiff(keys ...string) ([]string, error)

	Sdiffstore(dstkey string, keys ...string) (int64, error)
	Sinter(keys ...string) ([]string, error)
	Sinterstore(dstkey string, keys ...string) (int64, error)
	Smove(srckey, dstkey, member string) (int64, error)
	SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (int64, error)
	//Sort(key, dstkey string) (int64, error)
	Sunion(keys ...string) ([]string, error)
	Sunionstore(dstkey string, keys ...string) (int64, error)
	Watch(keys ...string) (string, error)
	Unwatch() (string, error)
	Zinterstore(dstkey string, sets ...string) (int64, error)
	ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error)
	Zunionstore(dstkey string, sets ...string) (int64, error)
	ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error)
	Brpoplpush(source, destination string, timeout int) (string, error)
	Publish(channel, message string) (int64, error)
	Subscribe(redisPubSub *RedisPubSub, channels ...string) error
	Psubscribe(redisPubSub *RedisPubSub, patterns ...string) error
	RandomKey() (string, error)
	Bitop(op BitOP, destKey string, srcKeys ...string) (int64, error)
	//Scan(cursor string) (ScanResult, error)
	Scan(cursor string, params ...ScanParams) (*ScanResult, error)
	Pfmerge(destkey string, sourcekeys ...string) (string, error)
	Pfcount(keys ...string) (int64, error)
}

//ScanResult
type ScanResult struct {
	Cursor  string
	Results []string
}

//ZParams
type ZParams struct {
	Name   string
	params [][]byte
}

//GetRaw
func (g ZParams) GetRaw() []byte {
	return []byte(g.Name)
}

//GetParams
func (g ZParams) GetParams() [][]byte {
	return g.params
}

//NewZParams
func NewZParams(name string) ZParams {
	return ZParams{Name: name}
}

var (
	ZParams_SUM = NewZParams("SUM")
	ZParams_MIN = NewZParams("MIN")
	ZParams_MAX = NewZParams("MAX")
)

//RedisPubSub
type RedisPubSub struct {
	subscribedChannels int
	Redis              *Redis
	OnMessage          func(channel, message string)
	OnPMessage         func(pattern string, channel, message string)
	OnSubscribe        func(channel string, subscribedChannels int)
	OnUnsubscribe      func(channel string, subscribedChannels int)
	OnPUnsubscribe     func(pattern string, subscribedChannels int)
	OnPSubscribe       func(pattern string, subscribedChannels int)
	OnPong             func(channel string)
}

//Subscribe
func (r *RedisPubSub) Subscribe(channels ...string) error {
	if r.Redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.Redis.client.Subscribe(channels...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Unsubscribe
func (r *RedisPubSub) Unsubscribe(channels ...string) error {
	if r.Redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.Redis.client.Unsubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Psubscribe
func (r *RedisPubSub) Psubscribe(channels ...string) error {
	if r.Redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.Redis.client.Psubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Punsubscribe
func (r *RedisPubSub) Punsubscribe(channels ...string) error {
	if r.Redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.Redis.client.Punsubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisPubSub) proceed(redis *Redis, channels ...string) error {
	r.Redis = redis
	err := r.Redis.client.Subscribe(channels...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return r.process(redis)
}

func (r *RedisPubSub) isSubscribed() bool {
	return r.subscribedChannels > 0
}

func (r *RedisPubSub) proceedWithPatterns(redis *Redis, patterns ...string) error {
	r.Redis = redis
	err := r.Redis.client.Psubscribe(patterns...)
	if err != nil {
		return err
	}
	err = r.Redis.client.flush()
	if err != nil {
		return err
	}
	return r.process(redis)
}

func (r *RedisPubSub) process(redis *Redis) error {
	for {
		reply, err := redis.client.connection.getRawObjectMultiBulkReply()
		if err != nil {
			return err
		}
		firstObj := reply[0]
		//fmt.Printf("%T\n", firstObj)
		//switch t := firstObj.(type) {
		//case []uint8:
		//	break
		//default:
		//	return errors.New(fmt.Sprintf("Unknown message type: %v", t))
		//}
		resp := firstObj.([]byte)
		respUpper := strings.ToUpper(string(resp))
		if string(KEYWORD_SUBSCRIBE.GetRaw()) == respUpper {
			r.subscribedChannels = int(reply[2].(int64))
			bchannel := reply[1].([]byte)
			strchannel := ""
			if bchannel != nil {
				strchannel = string(bchannel)
			}
			r.OnSubscribe(strchannel, int(r.subscribedChannels))
		} else if string(KEYWORD_UNSUBSCRIBE.GetRaw()) == respUpper {
			r.subscribedChannels = int(reply[2].(int64))
			bchannel := reply[1].([]byte)
			strchannel := ""
			if bchannel != nil {
				strchannel = string(bchannel)
			}
			r.OnUnsubscribe(strchannel, int(r.subscribedChannels))
		} else if string(KEYWORD_MESSAGE.GetRaw()) == respUpper {
			bchannel := reply[1].([]byte)
			bmesg := reply[2].([]byte)
			strchannel := ""
			if bchannel != nil {
				strchannel = string(bchannel)
			}
			strmesg := ""
			if bchannel != nil {
				strmesg = string(bmesg)
			}
			r.OnMessage(strchannel, strmesg)
		} else if string(KEYWORD_PMESSAGE.GetRaw()) == respUpper {
			bpattern := reply[1].([]byte)
			bchannel := reply[2].([]byte)
			bmesg := reply[31].([]byte)
			strpattern := ""
			if bpattern != nil {
				strpattern = string(bpattern)
			}
			strchannel := ""
			if bchannel != nil {
				strchannel = string(bchannel)
			}
			strmesg := ""
			if bchannel != nil {
				strmesg = string(bmesg)
			}
			r.OnPMessage(strpattern, strchannel, strmesg)
		} else if string(KEYWORD_PSUBSCRIBE.GetRaw()) == respUpper {
			r.subscribedChannels = int(reply[2].(int64))
			bpattern := reply[1].([]byte)
			strpattern := ""
			if bpattern != nil {
				strpattern = string(bpattern)
			}
			r.OnPSubscribe(strpattern, int(r.subscribedChannels))
		} else if string(CMD_PUNSUBSCRIBE.GetRaw()) == respUpper {
			r.subscribedChannels = int(reply[2].(int64))
			bpattern := reply[1].([]byte)
			strpattern := ""
			if bpattern != nil {
				strpattern = string(bpattern)
			}
			r.OnPUnsubscribe(strpattern, int(r.subscribedChannels))
		} else if string(KEYWORD_PONG.GetRaw()) == respUpper {
			bpattern := reply[1].([]byte)
			strpattern := ""
			if bpattern != nil {
				strpattern = string(bpattern)
			}
			r.OnPong(strpattern)
		} else {
			return errors.New(fmt.Sprintf("Unknown message type: %v", firstObj))
		}

		if !r.isSubscribed() {
			break
		}
	}
	/*
	 * Reset pipeline count because subscribe() calls would have increased it but nothing
	 * decremented it.
	 */
	redis.client.resetPipelinedCount()
	/* Invalidate instance since this thread is no longer listening */
	r.Redis.client = nil
	return nil
}

//BitOP
type BitOP struct {
	Name string
}

//GetRaw
func (g BitOP) GetRaw() []byte {
	return []byte(g.Name)
}

//NewBitOP
func NewBitOP(name string) BitOP {
	return BitOP{name}
}

var (
	BitOP_AND = NewBitOP("AND")
	BitOP_OR  = NewBitOP("OR")
	BitOP_XOR = NewBitOP("XOR")
	BitOP_NOT = NewBitOP("NOT")
)

//AdvancedRedisCommands
type AdvancedRedisCommands interface {
	ConfigGet(pattern string) ([]string, error)
	ConfigSet(parameter string, value string) (string, error)
	SlowlogReset() (string, error)
	SlowlogLen() (int64, error)
	//SlowlogGet() ([]Slowlog, error)
	SlowlogGet(entries ...int64) ([]Slowlog, error)
	ObjectRefcount(str string) (int64, error)
	ObjectEncoding(str string) (string, error)
	ObjectIdletime(str string) (int64, error)
}

//Slowlog
type Slowlog struct {
	id            int64
	timeStamp     int64
	executionTime int64
	args          []string
}

//ScriptingCommands
type ScriptingCommands interface {
	Eval(script string, keyCount int, params ...string) (interface{}, error)
	Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error)
	ScriptExists(sha1 ...string) ([]bool, error)
	ScriptLoad(script string) (string, error)
}

//BasicCommands
type BasicCommands interface {
	Ping() (string, error)
	Quit() (string, error)
	FlushDB() (string, error)
	DbSize() (int64, error)
	Select(index int) (string, error)
	FlushAll() (string, error)
	Auth(password string) (string, error)
	Save() (string, error)
	Bgsave() (string, error)
	Bgrewriteaof() (string, error)
	Lastsave() (int64, error)
	Shutdown() (string, error)
	//Info() (string, error)
	Info(section ...string) (string, error)
	Slaveof(host string, port int) (string, error)
	SlaveofNoOne() (string, error)
	GetDB() int
	Debug(params DebugParams) (string, error)
	ConfigResetStat() (string, error)
	WaitReplicas(replicas int, timeout int64) (int64, error)
}

//DebugParams
type DebugParams struct {
	command []string
}

//ClusterCommands
type ClusterCommands interface {
	ClusterNodes() (string, error)
	ClusterMeet(ip string, port int) (string, error)
	ClusterAddSlots(slots ...int) (string, error)
	ClusterDelSlots(slots ...int) (string, error)
	ClusterInfo() (string, error)
	ClusterGetKeysInSlot(slot int, count int) ([]string, error)
	ClusterSetSlotNode(slot int, nodeId string) (string, error)
	ClusterSetSlotMigrating(slot int, nodeId string) (string, error)
	ClusterSetSlotImporting(slot int, nodeId string) (string, error)
	ClusterSetSlotStable(slot int) (string, error)
	ClusterForget(nodeId string) (string, error)
	ClusterFlushSlots() (string, error)
	ClusterKeySlot(key string) (int64, error)
	ClusterCountKeysInSlot(slot int) (int64, error)
	ClusterSaveConfig() (string, error)
	ClusterReplicate(nodeId string) (string, error)
	ClusterSlaves(nodeId string) ([]string, error)
	ClusterFailover() (string, error)
	ClusterSlots() ([]interface{}, error)
	ClusterReset(resetType Reset) (string, error)
	Readonly() (string, error)
}

//Reset
type Reset struct {
	Name string
}

//GetRaw
func (g Reset) GetRaw() []byte {
	return []byte(g.Name)
}

//NewReset
func NewReset(name string) Reset {
	return Reset{name}
}

var (
	Reset_SOFT = NewReset("SOFT")
	Reset_HARD = NewReset("HARD")
)

//SentinelCommands
type SentinelCommands interface {
	SentinelMasters() ([]map[string]string, error)
	SentinelGetMasterAddrByName(masterName string) ([]string, error)
	SentinelReset(pattern string) (int64, error)
	SentinelSlaves(masterName string) ([]map[string]string, error)
	SentinelFailover(masterName string) (string, error)
	SentinelMonitor(masterName, ip string, port, quorum int) (string, error)
	SentinelRemove(masterName string) (string, error)
	SentinelSet(masterName string, parameterMap map[string]string) (string, error)
}

//ClusterScriptingCommands
type ClusterScriptingCommands interface {
	Eval(script string, keyCount int, params ...string) (interface{}, error)
	Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error)
	ScriptExists(key string, sha1 ...string) ([]bool, error)
	ScriptLoad(key, script string) (string, error)
}

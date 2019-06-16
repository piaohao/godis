package godis

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
	SetbitWithBool(key string, offset int64, value bool)
	Setbit(key string, offset int64, value string)
	Getbit(key string, offset int64)
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
	Zcore(key string, max string, min string, offset int, count int) ([]string, error)
	ZrangeByScoreWithScores(key string, min string, max string) ([]Tuple, error)
	ZrevrangeByScoreWithScores(key string, max string, min string) ([]Tuple, error)
	ZrangeByScoreWithScoresBatch(key string, min string, max string, offset int, count int) ([]Tuple, error)
	//ZrevrangeByScoreWithScores(key string, max float64, min float64, offset int, count int) ([]Tuple, error)
	ZrevrangeByScoreWithScoresBatch(key string, max string, min string, offset int, count int) ([]Tuple, error)
	ZremrangeByRank(key string, start int64, end int64) (int64, error)
	//ZremrangeByScore(key string, start float64, end float64) (int64, error)
	ZremrangeByScore(key string, start string, end string) (int64, error)
	Zlexcount(key string, min string, max string) (int64, error)
	ZrangeByLex(key string, min string, max string) ([]string, error)
	ZrangeByLexBatch(key string, min string, max string, offset int, count int) ([]string, error)
	ZrevrangeByLex(key string, max string, min string) ([]string, error)
	ZrevrangeByLexBatch(key string, max string, min string, offset int, count int) ([]string, error)
	ZremrangeByLex(key string, min string, max string) (int64, error)
	Linsert(key string, where ListOption, pivot, value string) (int64, error)
	Lpushx(key string, String ...string) (int64, error)
	Rpushx(key string, String ...string) (int64, error)
	//Brpop(timeout int, key string) ([]string, error)
	//Del(key string) (int64, error)
	Echo(String string) (string, error)
	Move(key string, dbIndex int) (int64, error)
	Bitcount(key string) (int64, error)
	BitcountRange(key string, start int64, end int64) (int64, error)
	//Bitpos(key string, value bool) (int64, error)
	Bitpos(key string, value bool, params ...BitPosParams) (int64, error)
	//Hscan(key string, cursor string) (ScanResult, error)
	Hscan(key string, cursor string, params ...ScanParams) (ScanResult, error)
	//Sscan(key string, cursor string) (ScanResult, error)
	Sscan(key string, cursor string, params ...ScanParams) (ScanResult, error)
	//Zscan(key string, cursor string) (ScanResult, error)
	Zscan(key string, cursor string, params ...ScanParams) (ScanResult, error)
	Pfadd(key string, elements ...string) (int64, error)
	//Pfcount(key string) (int64, error)

	// Geo Commands
	Geoadd(key string, longitude float64, latitude float64, member string) (int64, error)
	GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error)
	//Geodist(key string, member1, member2 string) (float64, error)
	Geodist(key string, member1, member2 string, unit ...GeoUnit) (float64, error)
	Geohash(key string, members ...string) ([]string, error)
	Geopos(key string, members ...string) ([]GeoCoordinate, error)
	//Georadius(key string, longitude float64, latitude float64, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	Georadius(key string, longitude float64, latitude float64, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]GeoCoordinate, error)
	//GeoradiusByMember(key string, member string, radius float64, unit GeoUnit) ([]GeoCoordinate, error)
	GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]GeoCoordinate, error)
	Bitfield(key string, arguments ...string) ([]int64, error)
}

type ZAddParams struct {
	XX bool
	NX bool
	CH bool
}

type BitPosParams struct {
	params [][]byte
}

type SortingParams struct {
	params [][]byte
}

type ScanParams struct {
	params map[keyword][]byte
}

type ListOption struct {
	Name string
}

func (l ListOption) GetRaw() []byte {
	return []byte(l.Name)
}

func newListOption(name string) ListOption {
	return ListOption{name}
}

var (
	ListOption_BEFORE = newListOption("BEFORE")
	ListOption_AFTER  = newListOption("AFTER")
)

type GeoUnit struct {
	Name string
}

func (g GeoUnit) GetRaw() []byte {
	return []byte(g.Name)
}

func newGeoUnit(name string) GeoUnit {
	return GeoUnit{name}
}

var (
	GEOUNIT_MI = newGeoUnit("MI")
	GEOUNIT_M  = newGeoUnit("M")
	GEOUNIT_KM = newGeoUnit("KM")
	GEOUNIT_FT = newGeoUnit("FT")
)

type GeoRadiusParam struct {
	params map[string][]byte
}

type Tuple struct {
	element []byte
	score   float64
}

type GeoRadiusResponse struct {
	member     []byte
	distance   float64
	coordinate GeoCoordinate
}

type GeoCoordinate struct {
	longitude float64
	latitude  float64
}

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
	Subscribe(jedisPubSub JedisPubSub, channels ...string) error
	Psubscribe(jedisPubSub JedisPubSub, patterns ...string) error
	RandomKey() (string, error)
	Bitop(op BitOP, destKey string, srcKeys ...string) (int64, error)
	//Scan(cursor string) (ScanResult, error)
	Scan(cursor string, params ...ScanParams) (ScanResult, error)
	Pfmerge(destkey string, sourcekeys ...string) (string, error)
	Pfcount(keys ...string) (int64, error)
}

type ScanResult struct {
	Cursor  []byte
	Results []string
}

type ZParams struct {
	Name string
}

func (g ZParams) GetRaw() []byte {
	return []byte(g.Name)
}

func newZParams(name string) ZParams {
	return ZParams{name}
}

var (
	ZParams_SUM = newZParams("SUM")
	ZParams_MIN = newZParams("MIN")
	ZParams_MAX = newZParams("MAX")
)

type JedisPubSub struct {
}

type BitOP struct {
	Name string
}

func (g BitOP) GetRaw() []byte {
	return []byte(g.Name)
}

func newBitOP(name string) BitOP {
	return BitOP{name}
}

var (
	BitOP_AND = newBitOP("AND")
	BitOP_OR  = newBitOP("OR")
	BitOP_XOR = newBitOP("XOR")
	BitOP_NOT = newBitOP("NOT")
)

type AdvancedJedisCommands interface {
	ConfigGet(pattern string) ([]string, error)
	ConfigSet(parameter string, value string) (string, error)
	SlowlogReset() (string, error)
	SlowlogLen() (int64, error)
	//SlowlogGet() ([]Slowlog, error)
	SlowlogGet(entries ...int64) ([]Slowlog, error)
	ObjectRefcount(String string) (int64, error)
	ObjectEncoding(String string) (string, error)
	ObjectIdletime(String string) (int64, error)
}

type Slowlog struct {
	id            int64
	timeStamp     int64
	executionTime int64
	args          []string
}

const COMMA = ","

type ScriptingCommands interface {
	Eval(script string, keyCount int, params ...string) (interface{}, error)
	//Eval(script string, keys, args []string) (interface{}, error)
	//Eval(script string) (interface{}, error)
	//Evalsha(script string) (interface{}, error)
	//Evalsha(sha1 string, keys, args []string) (interface{}, error)
	Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error)
	//ScriptExists(sha1 string) (bool, error)
	ScriptExists(sha1 ...string) ([]bool, error)
	ScriptLoad(script string) (string, error)
}

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
	GetDB() (int64, error)
	Debug(params DebugParams) (string, error)
	ConfigResetStat() (string, error)
	WaitReplicas(replicas int, timeout int64) (int64, error)
}

type DebugParams struct {
	command []string
}

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

type Reset struct {
	Name string
}

func (g Reset) GetRaw() []byte {
	return []byte(g.Name)
}

func newReset(name string) Reset {
	return Reset{name}
}

var (
	Reset_SOFT = newReset("SOFT")
	Reset_HARD = newReset("HARD")
)

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

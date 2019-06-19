package godis

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	MASTER_NODE_INDEX = 2
)

type RedisClusterInfoCache struct {
	nodes map[string]*Pool
	slots map[int]*Pool

	rwLock        sync.RWMutex
	rLock         sync.Mutex
	wLock         sync.Mutex
	rediscovering bool
	poolConfig    PoolConfig

	connectionTimeout int
	soTimeout         int
	password          string
}

func NewRedisClusterInfoCache(connectionTimeout, soTimeout int, password string, poolConfig PoolConfig) *RedisClusterInfoCache {
	return &RedisClusterInfoCache{
		poolConfig:        poolConfig,
		connectionTimeout: connectionTimeout,
		soTimeout:         soTimeout,
		password:          password,
	}
}

func (r *RedisClusterInfoCache) discoverClusterNodesAndSlots(redis *Redis) error {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	r.reset(false)
	slots, err := redis.ClusterSlots()
	if err != nil {
		return err
	}
	for _, s := range slots {
		slotInfo := s.([]interface{})
		size := len(slotInfo)
		if size <= MASTER_NODE_INDEX {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		for i := MASTER_NODE_INDEX; i < size; i++ {
			hostInfos := slotInfo[i].([]interface{})
			if len(hostInfos) <= 0 {
				continue
			}
			host, port := r.generateHostAndPort(hostInfos)
			r.setupNodeIfNotExist(false, host, port)
			if i == MASTER_NODE_INDEX {
				r.assignSlotsToNode(false, slotNums, host, port)
			}
		}
	}
	return nil
}

func (r *RedisClusterInfoCache) renewClusterSlots(redis *Redis) error {
	if r.rediscovering {
		return nil
	}
	r.wLock.Lock()
	defer func() {
		r.wLock.Unlock()
		r.rediscovering = false
	}()
	if redis != nil {
		return r.discoverClusterSlots(redis)
	}
	for _, jp := range r.getShuffledNodesPool() {
		newRedis, err := jp.GetResource()
		if err != nil {
			continue
		}
		err = r.discoverClusterSlots(newRedis)
		if err != nil {
			continue
		}
		err = newRedis.Close()
		return err
	}
	return nil
}

func (r *RedisClusterInfoCache) discoverClusterSlots(redis *Redis) error {
	slots, err := redis.ClusterSlots()
	if err != nil {
		return err
	}
	r.slots = make(map[int]*Pool)
	for _, s := range slots {
		slotInfo := s.([]interface{})
		size := len(slotInfo)
		if size <= MASTER_NODE_INDEX {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		hostInfos := slotInfo[MASTER_NODE_INDEX].([]interface{})
		if len(hostInfos) == 0 {
			continue
		}
		host, port := r.generateHostAndPort(hostInfos)
		r.assignSlotsToNode(true, slotNums, host, port)
	}
	return nil
}

func (r *RedisClusterInfoCache) reset(lock bool) {
	if lock {
		r.wLock.Lock()
	}
	defer func() {
		if lock {
			r.wLock.Unlock()
		}
	}()
	for _, v := range r.nodes {
		if v != nil {
			v.Destroy()
		}
	}
	r.nodes = make(map[string]*Pool)
	r.slots = make(map[int]*Pool)
}

func (r *RedisClusterInfoCache) getAssignedSlotArray(slotInfo []interface{}) []int {
	slotNums := make([]int, 0)
	for slot := slotInfo[0].(int64); slot <= slotInfo[1].(int64); slot++ {
		slotNums = append(slotNums, int(slot))
	}
	return slotNums
}

func (r *RedisClusterInfoCache) generateHostAndPort(hostInfos []interface{}) (string, int) {
	return string(hostInfos[0].([]byte)), int(hostInfos[1].(int64))
}

func (r *RedisClusterInfoCache) setupNodeIfNotExist(lock bool, host string, port int) *Pool {
	if lock {
		r.wLock.Lock()
	}
	defer func() {
		if lock {
			r.wLock.Unlock()
		}
	}()
	nodeKey := host + ":" + strconv.Itoa(port)
	existingPool, ok := r.nodes[nodeKey]
	if ok && existingPool != nil {
		return existingPool
	}
	nodePool := NewPool(r.poolConfig, NewFactory(ShardInfo{
		Host:              host,
		Port:              port,
		ConnectionTimeout: r.connectionTimeout,
		SoTimeout:         r.soTimeout,
		Password:          r.password,
	}))
	return nodePool
}

func (r *RedisClusterInfoCache) assignSlotToNode(slot int, host string, port int) {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	targetPool := r.setupNodeIfNotExist(false, host, port)
	r.slots[slot] = targetPool
}

func (r *RedisClusterInfoCache) assignSlotsToNode(lock bool, slots []int, host string, port int) {
	if lock {
		r.wLock.Lock()
	}
	defer func() {
		if lock {
			r.wLock.Unlock()
		}
	}()
	targetPool := r.setupNodeIfNotExist(false, host, port)
	for _, slot := range slots {
		r.slots[slot] = targetPool
	}
}

func (r *RedisClusterInfoCache) getShuffledNodesPool() []*Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	pools := make([]*Pool, 0)
	for _, v := range r.nodes {
		pools = append(pools, v)
	}
	r.shuffle(pools)
	return pools
}

func (r *RedisClusterInfoCache) shuffle(vals []*Pool) {
	ra := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := ra.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func (r *RedisClusterInfoCache) getNode(nodeKey string) *Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.nodes[nodeKey]
}

func (r *RedisClusterInfoCache) getNodes() map[string]*Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.nodes
}

func (r *RedisClusterInfoCache) getSlotPool(slot int) *Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.slots[slot]
}

type RedisClusterConnectionHandler struct {
	cache *RedisClusterInfoCache
}

func NewRedisClusterConnectionHandler(nodes []string, connectionTimeout, soTimeout int, password string, poolConfig PoolConfig) *RedisClusterConnectionHandler {
	cache := NewRedisClusterInfoCache(connectionTimeout, soTimeout, password, poolConfig)
	for _, node := range nodes {
		arr := strings.Split(node, ":")
		port, err := strconv.Atoi(arr[1])
		if err != nil {
			continue
		}
		redis := NewRedis(ShardInfo{
			Host: arr[0],
			Port: port,
		})
		if password != "" {
			_, err := redis.Auth(password)
			if err != nil {
				continue
			}
		}
		err = cache.discoverClusterNodesAndSlots(redis)
		if err != nil {
			continue
		}
		redis.Close()
		break
	}

	return &RedisClusterConnectionHandler{cache: cache}
}

func (r *RedisClusterConnectionHandler) getConnection() (*Redis, error) {
	pools := r.cache.getShuffledNodesPool()
	for _, pool := range pools {
		redis, err := pool.GetResource()
		if err != nil {
			continue
		}
		result, err := redis.Ping()
		if err != nil {
			continue
		}
		redis.Close()
		if strings.ToUpper(result) == KEYWORD_PONG.Name {
			return redis, nil
		}
	}
	return nil, errors.New("no reachable node in cluster")
}

func (r *RedisClusterConnectionHandler) getConnectionFromSlot(slot int) (*Redis, error) {
	connectionPool := r.cache.getSlotPool(slot)
	if connectionPool != nil {
		return connectionPool.GetResource()
	}
	r.renewSlotCache()
	connectionPool = r.cache.getSlotPool(slot)
	if connectionPool != nil {
		return connectionPool.GetResource()
	}
	return r.getConnection()
}

func (r *RedisClusterConnectionHandler) getConnectionFromNode(host string, port int) (*Redis, error) {
	return r.cache.setupNodeIfNotExist(true, host, port).GetResource()
}

func (r *RedisClusterConnectionHandler) getNodes() map[string]*Pool {
	return r.cache.getNodes()
}

func (r *RedisClusterConnectionHandler) renewSlotCache(redis ...*Redis) {
	if len(redis) == 0 {
		r.cache.renewClusterSlots(nil)
		return
	}
	for _, re := range redis {
		r.cache.renewClusterSlots(re)
	}
}

type RedisClusterHashTagUtil struct {
}

func NewRedisClusterHashTagUtil() *RedisClusterHashTagUtil {
	return &RedisClusterHashTagUtil{}
}

func (r *RedisClusterHashTagUtil) getHashTag(key string) string {
	return r.extractHashTag(key, true)
}

func (r *RedisClusterHashTagUtil) isClusterCompliantMatchPattern(matchPattern string) bool {
	tag := r.extractHashTag(matchPattern, false)
	return tag != ""
}

func (r *RedisClusterHashTagUtil) extractHashTag(key string, returnKeyOnAbsence bool) string {
	s := strings.Index(key, "{")
	if s > -1 {
		e := strings.Index(key[s+1:], "}")
		if e > -1 && e != s+1 {
			return key[s+1 : e]
		}
	}
	if returnKeyOnAbsence {
		return key
	} else {
		return ""
	}
}

type RedisClusterCommand struct {
	MaxAttempts       int
	ConnectionHandler *RedisClusterConnectionHandler

	ctx context.Context

	execute func(redis *Redis) (interface{}, error)
}

func NewRedisClusterCommand(maxAttempts int, connectionHandler *RedisClusterConnectionHandler) *RedisClusterCommand {
	return &RedisClusterCommand{MaxAttempts: maxAttempts, ConnectionHandler: connectionHandler}
}

func (r *RedisClusterCommand) run(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("no way to dispatch this command to Redis Cluster")
	}
	return r.runWithRetries([]byte(key), r.MaxAttempts, false, false)
}

func (r *RedisClusterCommand) runBatch(keyCount int, keys ...string) (interface{}, error) {
	if len(keys) == 0 {
		return nil, errors.New("no way to dispatch this command to Redis Cluster")
	}
	if len(keys) > 1 {
		crc16 := NewCRC16()
		slot := crc16.getStringSlot(keys[0])
		for i := 1; i < keyCount; i++ {
			nextSlot := crc16.getStringSlot(keys[i])
			if nextSlot != slot {
				return nil, errors.New("no way to dispatch this command to Redis Cluster,because keys have different slots")
			}
		}
	}
	return r.runWithRetries([]byte(keys[0]), r.MaxAttempts, false, false)
}

func (r *RedisClusterCommand) runWithAnyNode() (interface{}, error) {
	connection := new(Redis)
	defer connection.Close()
	connection, err := r.ConnectionHandler.getConnection()
	if err != nil {
		return nil, err
	}
	return r.execute(connection)
}

func (r *RedisClusterCommand) releaseConnection(redis *Redis) error {
	if redis != nil {
		return redis.Close()
	}
	return nil
}

func (r *RedisClusterCommand) runWithRetries(key []byte, attempts int, tryRandomNode, asking bool) (interface{}, error) {
	if attempts <= 0 {
		return nil, errors.New("too many Cluster redirections")
	}
	connection := new(Redis)
	var err error
	if asking {
		connection = r.ctx.Value("").(*Redis)
		_, err = connection.Asking()
		if err != nil {
			return nil, err
		}
		asking = false
	} else {
		if tryRandomNode {
			connection, err = r.ConnectionHandler.getConnection()
			if err != nil {
				return nil, err
			}
		} else {
			connection, err = r.ConnectionHandler.getConnectionFromSlot(int(NewCRC16().getByteSlot(key)))
			if err != nil {
				return nil, err
			}
		}
	}
	//todo 根据各种error，进行重试或者重新分配slot的逻辑
	return r.execute(connection)
}

type RedisCluster struct {
	MaxAttempts       int
	ConnectionHandler *RedisClusterConnectionHandler
}

func NewRedisCluster(nodes []string, connectionTimeout, soTimeout, maxAttempts int, password string, poolConfig PoolConfig) *RedisCluster {
	return &RedisCluster{
		MaxAttempts:       maxAttempts,
		ConnectionHandler: NewRedisClusterConnectionHandler(nodes, connectionTimeout, soTimeout, password, poolConfig),
	}
}

//<editor-fold desc="rediscommands">
func (r *RedisCluster) Set(key, value string) (string, error) {
	command := NewRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Set(key, value)
	}
	return ToStringReply(command.run(key))
}

func (r *RedisCluster) SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) SetWithParams(key, value, nxxx string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Get(key string) (string, error) {
	command := NewRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Get(key)
	}
	return ToStringReply(command.run(key))
}

func (r *RedisCluster) Persist(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Type(key string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Expire(key string, seconds int) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Pexpire(key string, milliseconds int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ExpireAt(key string, unixtime int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) PexpireAt(key string, millisecondsTimestamp int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Ttl(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Pttl(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) SetbitWithBool(key string, offset int64, value bool) (bool, error) {
	panic("implement me")
}

func (r *RedisCluster) Setbit(key string, offset int64, value string) (bool, error) {
	panic("implement me")
}

func (r *RedisCluster) Getbit(key string, offset int64) (bool, error) {
	panic("implement me")
}

func (r *RedisCluster) Setrange(key string, offset int64, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Getrange(key string, startOffset, endOffset int64) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) GetSet(key, value string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Setnx(key, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Setex(key string, seconds int, value string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Psetex(key string, milliseconds int64, value string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) DecrBy(key string, decrement int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Decr(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) IncrBy(key string, increment int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) IncrByFloat(key string, increment float64) (float64, error) {
	panic("implement me")
}

func (r *RedisCluster) Incr(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Append(key, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Substr(key string, start, end int) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Hset(key, field string, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hget(key, field string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Hsetnx(key, field, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hmset(key string, hash map[string]string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Hmget(key string, fields ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) HincrBy(key, field string, value int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) HincrByFloat(key, field string, value float64) (float64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hexists(key, field string) (bool, error) {
	panic("implement me")
}

func (r *RedisCluster) Hdel(key string, fields ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hlen(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hkeys(key string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Hvals(key string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) HgetAll(key string) (map[string]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Rpush(key string, strings ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Lpush(key string, strings ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Llen(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Lrange(key string, start, stop int64) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Ltrim(key string, start, stop int64) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Lindex(key string, index int64) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Lset(key string, index int64, value string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Lrem(key string, count int64, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Lpop(key string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Rpop(key string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Sadd(key string, members ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Smembers(key string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Srem(key string, members ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Spop(key string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) SpopBatch(key string, count int64) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Scard(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Sismember(key string, member string) (bool, error) {
	panic("implement me")
}

func (r *RedisCluster) Srandmember(key string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) SrandmemberBatch(key string, count int) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Strlen(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zadd(key string, score float64, member string, params ...ZAddParams) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zrange(key string, start, end int64) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Zrem(key string, member ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zincrby(key string, score float64, member string, params ...ZAddParams) (float64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zrank(key, member string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zrevrank(key, member string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zrevrange(key string, start, end int64) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) Zcard(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zscore(key, member string) (float64, error) {
	panic("implement me")
}

func (r *RedisCluster) Sort(key string, sortingParameters ...SortingParams) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Zcount(key string, min string, max string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByScore(key string, min string, max string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeByScore(key string, max string, min string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByScoreBatch(key string, min string, max string, offset int, count int) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByScoreWithScores(key, min, max string) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeByScoreWithScores(key, max, min string) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) ([]Tuple, error) {
	panic("implement me")
}

func (r *RedisCluster) ZremrangeByRank(key string, start, end int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZremrangeByScore(key, start, end string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zlexcount(key, min, max string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByLex(key, min, max string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrangeByLexBatch(key, min, max string, offset, count int) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeByLex(key, max, min string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZrevrangeByLexBatch(key, max, min string, offset, count int) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) ZremrangeByLex(key, min, max string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Linsert(key string, where ListOption, pivot, value string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Lpushx(key string, String ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Rpushx(key string, String ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Echo(str string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Move(key string, dbIndex int) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Bitcount(key string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) BitcountRange(key string, start int64, end int64) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Bitpos(key string, value bool, params ...BitPosParams) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Hscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	panic("implement me")
}

func (r *RedisCluster) Sscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	panic("implement me")
}

func (r *RedisCluster) Zscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	panic("implement me")
}

func (r *RedisCluster) Pfadd(key string, elements ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Geoadd(key string, longitude, latitude float64, member string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Geodist(key string, member1, member2 string, unit ...GeoUnit) (float64, error) {
	panic("implement me")
}

func (r *RedisCluster) Geohash(key string, members ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Geopos(key string, members ...string) ([]*GeoCoordinate, error) {
	panic("implement me")
}

func (r *RedisCluster) Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error) {
	panic("implement me")
}

func (r *RedisCluster) GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error) {
	panic("implement me")
}

func (r *RedisCluster) Bitfield(key string, arguments ...string) ([]int64, error) {
	panic("implement me")
}

//</editor-fold>

//<editor-fold desc="multikeycommands">
func (r *RedisCluster) Del(keys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Exists(keys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) BlpopTimout(timeout int, keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) BrpopTimout(timeout int, keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Blpop(args ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Brpop(args ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Keys(pattern string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Mget(keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Mset(keysvalues ...string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Msetnx(keysvalues ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Rename(oldkey, newkey string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Renamenx(oldkey, newkey string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Rpoplpush(srckey, dstkey string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Sdiff(keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Sdiffstore(dstkey string, keys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Sinter(keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Sinterstore(dstkey string, keys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Smove(srckey, dstkey, member string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Sunion(keys ...string) ([]string, error) {
	panic("implement me")
}

func (r *RedisCluster) Sunionstore(dstkey string, keys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Watch(keys ...string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Unwatch() (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Zinterstore(dstkey string, sets ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Zunionstore(dstkey string, sets ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Brpoplpush(source, destination string, timeout int) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Publish(channel, message string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Subscribe(redisPubSub *RedisPubSub, channels ...string) error {
	panic("implement me")
}

func (r *RedisCluster) Psubscribe(redisPubSub *RedisPubSub, patterns ...string) error {
	panic("implement me")
}

func (r *RedisCluster) RandomKey() (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Bitop(op BitOP, destKey string, srcKeys ...string) (int64, error) {
	panic("implement me")
}

func (r *RedisCluster) Scan(cursor string, params ...ScanParams) (*ScanResult, error) {
	panic("implement me")
}

func (r *RedisCluster) Pfmerge(destkey string, sourcekeys ...string) (string, error) {
	panic("implement me")
}

func (r *RedisCluster) Pfcount(keys ...string) (int64, error) {
	panic("implement me")
}

//</editor-fold>

//<editor-fold desc="scriptcommands">
func (r *RedisCluster) Eval(script string, keyCount int, params ...string) (interface{}, error) {
	panic("implement me")
}

func (r *RedisCluster) Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error) {
	panic("implement me")
}

func (r *RedisCluster) ScriptExists(sha1 ...string) ([]bool, error) {
	panic("implement me")
}

func (r *RedisCluster) ScriptLoad(script string) (string, error) {
	panic("implement me")
}

//</editor-fold>

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
	MasterNodeIndex = 2
)

type redisClusterInfoCache struct {
	nodes map[string]*Pool
	slots map[int]*Pool

	rwLock        sync.RWMutex
	rLock         sync.Mutex
	wLock         sync.Mutex
	rediscovering bool
	poolConfig    *PoolConfig

	connectionTimeout time.Duration
	soTimeout         time.Duration
	password          string
}

func newRedisClusterInfoCache(connectionTimeout, soTimeout time.Duration, password string, poolConfig *PoolConfig) *redisClusterInfoCache {
	return &redisClusterInfoCache{
		poolConfig:        poolConfig,
		connectionTimeout: connectionTimeout,
		soTimeout:         soTimeout,
		password:          password,
	}
}

func (r *redisClusterInfoCache) discoverClusterNodesAndSlots(redis *Redis) error {
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
		if size <= MasterNodeIndex {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		for i := MasterNodeIndex; i < size; i++ {
			hostInfos := slotInfo[i].([]interface{})
			if len(hostInfos) <= 0 {
				continue
			}
			host, port := r.generateHostAndPort(hostInfos)
			r.setupNodeIfNotExist(false, host, port)
			if i == MasterNodeIndex {
				r.assignSlotsToNode(false, slotNums, host, port)
			}
		}
	}
	return nil
}

func (r *redisClusterInfoCache) renewClusterSlots(redis *Redis) error {
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
		newRedis, err := jp.Get()
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

func (r *redisClusterInfoCache) discoverClusterSlots(redis *Redis) error {
	slots, err := redis.ClusterSlots()
	if err != nil {
		return err
	}
	r.slots = make(map[int]*Pool)
	for _, s := range slots {
		slotInfo := s.([]interface{})
		size := len(slotInfo)
		if size <= MasterNodeIndex {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		hostInfos := slotInfo[MasterNodeIndex].([]interface{})
		if len(hostInfos) == 0 {
			continue
		}
		host, port := r.generateHostAndPort(hostInfos)
		r.assignSlotsToNode(true, slotNums, host, port)
	}
	return nil
}

func (r *redisClusterInfoCache) reset(lock bool) {
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

func (r *redisClusterInfoCache) getAssignedSlotArray(slotInfo []interface{}) []int {
	slotNums := make([]int, 0)
	for slot := slotInfo[0].(int64); slot <= slotInfo[1].(int64); slot++ {
		slotNums = append(slotNums, int(slot))
	}
	return slotNums
}

func (r *redisClusterInfoCache) generateHostAndPort(hostInfos []interface{}) (string, int) {
	return string(hostInfos[0].([]byte)), int(hostInfos[1].(int64))
}

func (r *redisClusterInfoCache) setupNodeIfNotExist(lock bool, host string, port int) *Pool {
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
	nodePool := NewPool(r.poolConfig, &Option{
		Host:              host,
		Port:              port,
		ConnectionTimeout: r.connectionTimeout,
		SoTimeout:         r.soTimeout,
		Password:          r.password,
	})
	/*nodePool := NewPool(r.poolConfig, NewFactory(&Option{
		Host:              host,
		Port:              port,
		ConnectionTimeout: r.connectionTimeout,
		SoTimeout:         r.soTimeout,
		Password:          r.password,
	}))*/
	r.nodes[nodeKey] = nodePool
	return nodePool
}

func (r *redisClusterInfoCache) assignSlotToNode(slot int, host string, port int) {
	r.wLock.Lock()
	defer r.wLock.Unlock()
	targetPool := r.setupNodeIfNotExist(false, host, port)
	r.slots[slot] = targetPool
}

func (r *redisClusterInfoCache) assignSlotsToNode(lock bool, slots []int, host string, port int) {
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

func (r *redisClusterInfoCache) getShuffledNodesPool() []*Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	pools := make([]*Pool, 0)
	for _, v := range r.nodes {
		pools = append(pools, v)
	}
	r.shuffle(pools)
	return pools
}

func (r *redisClusterInfoCache) shuffle(vals []*Pool) {
	ra := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := ra.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func (r *redisClusterInfoCache) getNode(nodeKey string) *Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.nodes[nodeKey]
}

func (r *redisClusterInfoCache) getNodes() map[string]*Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.nodes
}

func (r *redisClusterInfoCache) getSlotPool(slot int) *Pool {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	return r.slots[slot]
}

type redisClusterConnectionHandler struct {
	cache *redisClusterInfoCache
}

func newRedisClusterConnectionHandler(nodes []string, connectionTimeout, soTimeout time.Duration, password string, poolConfig *PoolConfig) *redisClusterConnectionHandler {
	cache := newRedisClusterInfoCache(connectionTimeout, soTimeout, password, poolConfig)
	for _, node := range nodes {
		arr := strings.Split(node, ":")
		port, err := strconv.Atoi(arr[1])
		if err != nil {
			continue
		}
		redis := NewRedis(&Option{
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
		_ = redis.Close()
		break
	}

	return &redisClusterConnectionHandler{cache: cache}
}

func (r *redisClusterConnectionHandler) getConnection() (*Redis, error) {
	pools := r.cache.getShuffledNodesPool()
	for _, pool := range pools {
		redis, err := pool.Get()
		if err != nil {
			continue
		}
		result, err := redis.Ping()
		if err != nil {
			continue
		}
		if strings.ToUpper(result) == KeywordPong.Name {
			return redis, nil
		}
	}
	return nil, NewNoReachableClusterNodeError("no reachable node in cluster")
}

func (r *redisClusterConnectionHandler) getConnectionFromSlot(slot int) (*Redis, error) {
	connectionPool := r.cache.getSlotPool(slot)
	if connectionPool != nil {
		return connectionPool.Get()
	}
	r.renewSlotCache()
	connectionPool = r.cache.getSlotPool(slot)
	if connectionPool != nil {
		return connectionPool.Get()
	}
	return r.getConnection()
}

func (r *redisClusterConnectionHandler) getConnectionFromNode(host string, port int) (*Redis, error) {
	return r.cache.setupNodeIfNotExist(true, host, port).Get()
}

func (r *redisClusterConnectionHandler) getNodes() map[string]*Pool {
	return r.cache.getNodes()
}

func (r *redisClusterConnectionHandler) renewSlotCache(redis ...*Redis) {
	if len(redis) == 0 {
		_ = r.cache.renewClusterSlots(nil)
		return
	}
	for _, re := range redis {
		_ = r.cache.renewClusterSlots(re)
	}
}

type redisClusterHashTagUtil struct {
}

func newRedisClusterHashTagUtil() *redisClusterHashTagUtil {
	return &redisClusterHashTagUtil{}
}

func (r *redisClusterHashTagUtil) getHashTag(key string) string {
	return r.extractHashTag(key, true)
}

func (r *redisClusterHashTagUtil) isClusterCompliantMatchPattern(matchPattern string) bool {
	tag := r.extractHashTag(matchPattern, false)
	return tag != ""
}

func (r *redisClusterHashTagUtil) extractHashTag(key string, returnKeyOnAbsence bool) string {
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

type redisClusterCommand struct {
	MaxAttempts       int
	ConnectionHandler *redisClusterConnectionHandler

	ctx context.Context

	execute func(redis *Redis) (interface{}, error)
}

func newRedisClusterCommand(maxAttempts int, connectionHandler *redisClusterConnectionHandler) *redisClusterCommand {
	return &redisClusterCommand{MaxAttempts: maxAttempts, ConnectionHandler: connectionHandler, ctx: context.Background()}
}

func (r *redisClusterCommand) run(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("no way to dispatch this command to Redis cluster")
	}
	return r.runWithRetries([]byte(key), r.MaxAttempts, false, false)
}

func (r *redisClusterCommand) runBatch(keyCount int, keys ...string) (interface{}, error) {
	if len(keys) == 0 {
		return nil, errors.New("no way to dispatch this command to Redis cluster")
	}
	if len(keys) > 1 {
		crc16 := newCRC16()
		slot := crc16.getStringSlot(keys[0])
		for i := 1; i < keyCount; i++ {
			nextSlot := crc16.getStringSlot(keys[i])
			if nextSlot != slot {
				return nil, errors.New("no way to dispatch this command to Redis cluster,because keys have different slots")
			}
		}
	}
	return r.runWithRetries([]byte(keys[0]), r.MaxAttempts, false, false)
}

func (r *redisClusterCommand) runWithAnyNode() (interface{}, error) {
	connection, err := r.ConnectionHandler.getConnection()
	if err != nil {
		return nil, err
	}
	result, err := r.execute(connection)
	if err != nil {
		return nil, err
	}
	_ = r.releaseConnection(connection)
	return result, nil
}

func (r *redisClusterCommand) releaseConnection(redis *Redis) error {
	if redis != nil {
		return redis.Close()
	}
	return nil
}

func (r *redisClusterCommand) runWithRetries(key []byte, attempts int, tryRandomNode, asking bool) (interface{}, error) {
	if attempts <= 0 {
		return nil, errors.New("too many cluster redirections")
	}
	connection := new(Redis)
	var err error
	if asking {
		connection = r.ctx.Value("redis").(*Redis)
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
			connection, err = r.ConnectionHandler.getConnectionFromSlot(int(newCRC16().getByteSlot(key)))
			if err != nil {
				return nil, err
			}
		}
	}
	result, err := r.execute(connection)
	defer r.releaseConnection(connection)
	if err == nil {
		return result, nil
	}
	// 根据各种error，进行重试或者重新分配slot的逻辑
	// 判断 NoReachableClusterNodeException，直接返回错误
	// 判断 ConnectionException，重试，当attempt<=1时，重新分配slot
	// 判断 RedirectionException，如果是MovedDataException，则重新分配slot，如果是AskDataException，则设置ctx，如果是其他错误，直接返回错误，继续重试
	switch err.(type) {
	case *NoReachableClusterNodeError:
		return nil, err
	case *ConnectError:
		_ = r.releaseConnection(connection)
		connection = nil
		if attempts <= 1 {
			r.ConnectionHandler.renewSlotCache()
			return nil, err
		}
		return r.runWithRetries(key, attempts-1, tryRandomNode, asking)
	case *MovedDataError:
		r.ConnectionHandler.renewSlotCache(connection)
		r.releaseConnection(connection)
		connection = nil
		return r.runWithRetries(key, attempts-1, false, asking)
	case *AskDataError:
		r.releaseConnection(connection)
		connection = nil
		asking = true
		dataError := err.(*AskDataError)
		redis, err := r.ConnectionHandler.getConnectionFromNode(dataError.Host, dataError.Port)
		if err != nil {
			return nil, err
		}
		context.WithValue(r.ctx, "redis", redis)
		return r.runWithRetries(key, attempts-1, false, asking)
	}
	return nil, err
}

type ClusterOption struct {
	Nodes             []string
	ConnectionTimeout time.Duration
	SoTimeout         time.Duration
	MaxAttempts       int
	Password          string
	PoolConfig        *PoolConfig
}

//RedisCluster redis cluster tool
type RedisCluster struct {
	MaxAttempts       int
	ConnectionHandler *redisClusterConnectionHandler
}

//NewRedisCluster constructor
func NewRedisCluster(option *ClusterOption) *RedisCluster {
	if option.MaxAttempts <= 0 {
		option.MaxAttempts = 1
	}
	conTimeout := option.ConnectionTimeout
	if option.ConnectionTimeout == 0 {
		conTimeout = 5 * time.Second
	}
	soTimeout := option.SoTimeout
	if option.SoTimeout == 0 {
		soTimeout = 5 * time.Second
	}
	return &RedisCluster{
		MaxAttempts:       option.MaxAttempts,
		ConnectionHandler: newRedisClusterConnectionHandler(option.Nodes, conTimeout, soTimeout, option.Password, option.PoolConfig),
	}
}

//<editor-fold desc="rediscommands">

//Set ...
func (r *RedisCluster) Set(key, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Set(key, value)
	}
	return ToStringReply(command.run(key))
}

//SetWithParamsAndTime ...
func (r *RedisCluster) SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		defer redis.Close()
		return redis.SetWithParamsAndTime(key, value, nxxx, expx, time)
	}
	return ToStringReply(command.run(key))
}

//SetWithParams ...
func (r *RedisCluster) SetWithParams(key, value, nxxx string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetWithParams(key, value, nxxx)
	}
	return ToStringReply(command.run(key))
}

//Get ...
func (r *RedisCluster) Get(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Get(key)
	}
	return ToStringReply(command.run(key))
}

//Persist ...
func (r *RedisCluster) Persist(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Persist(key)
	}
	return ToInt64Reply(command.run(key))
}

//Type ...
func (r *RedisCluster) Type(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Type(key)
	}
	return ToStringReply(command.run(key))
}

//Expire ...
func (r *RedisCluster) Expire(key string, seconds int) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Expire(key, seconds)
	}
	return ToInt64Reply(command.run(key))
}

//Pexpire ...
func (r *RedisCluster) Pexpire(key string, milliseconds int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Pexpire(key, milliseconds)
	}
	return ToInt64Reply(command.run(key))
}

//ExpireAt ...
func (r *RedisCluster) ExpireAt(key string, unixtime int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ExpireAt(key, unixtime)
	}
	return ToInt64Reply(command.run(key))
}

//PexpireAt ...
func (r *RedisCluster) PexpireAt(key string, millisecondsTimestamp int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PexpireAt(key, millisecondsTimestamp)
	}
	return ToInt64Reply(command.run(key))
}

//Ttl ...
func (r *RedisCluster) Ttl(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Ttl(key)
	}
	return ToInt64Reply(command.run(key))
}

//Pttl ...
func (r *RedisCluster) Pttl(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Pttl(key)
	}
	return ToInt64Reply(command.run(key))
}

//SetbitWithBool ...
func (r *RedisCluster) SetbitWithBool(key string, offset int64, value bool) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetbitWithBool(key, offset, value)
	}
	return ToBoolReply(command.run(key))
}

//Setbit ...
func (r *RedisCluster) Setbit(key string, offset int64, value string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Setbit(key, offset, value)
	}
	return ToBoolReply(command.run(key))
}

//Getbit ...
func (r *RedisCluster) Getbit(key string, offset int64) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Getbit(key, offset)
	}
	return ToBoolReply(command.run(key))
}

//Setrange ...
func (r *RedisCluster) Setrange(key string, offset int64, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Setrange(key, offset, value)
	}
	return ToInt64Reply(command.run(key))
}

//Getrange ...
func (r *RedisCluster) Getrange(key string, startOffset, endOffset int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Getrange(key, startOffset, endOffset)
	}
	return ToStringReply(command.run(key))
}

//GetSet ...
func (r *RedisCluster) GetSet(key, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GetSet(key, value)
	}
	return ToStringReply(command.run(key))
}

//Setnx ...
func (r *RedisCluster) Setnx(key, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Setnx(key, value)
	}
	return ToInt64Reply(command.run(key))
}

//Setex ...
func (r *RedisCluster) Setex(key string, seconds int, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Setex(key, seconds, value)
	}
	return ToStringReply(command.run(key))
}

//Psetex ...
func (r *RedisCluster) Psetex(key string, milliseconds int64, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Psetex(key, milliseconds, value)
	}
	return ToStringReply(command.run(key))
}

//DecrBy ...
func (r *RedisCluster) DecrBy(key string, decrement int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.DecrBy(key, decrement)
	}
	return ToInt64Reply(command.run(key))
}

//Decr ...
func (r *RedisCluster) Decr(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Decr(key)
	}
	return ToInt64Reply(command.run(key))
}

//IncrBy ...
func (r *RedisCluster) IncrBy(key string, increment int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.IncrBy(key, increment)
	}
	return ToInt64Reply(command.run(key))
}

//IncrByFloat ...
func (r *RedisCluster) IncrByFloat(key string, increment float64) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.IncrByFloat(key, increment)
	}
	return ToFloat64Reply(command.run(key))
}

//Incr ...
func (r *RedisCluster) Incr(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Incr(key)
	}
	return ToInt64Reply(command.run(key))
}

//Append ...
func (r *RedisCluster) Append(key, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Append(key, value)
	}
	return ToInt64Reply(command.run(key))
}

//Substr ...
func (r *RedisCluster) Substr(key string, start, end int) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Substr(key, start, end)
	}
	return ToStringReply(command.run(key))
}

//Hset ...
func (r *RedisCluster) Hset(key, field string, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hset(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//Hget ...
func (r *RedisCluster) Hget(key, field string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hget(key, field)
	}
	return ToStringReply(command.run(key))
}

//Hsetnx ...
func (r *RedisCluster) Hsetnx(key, field, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hsetnx(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//Hmset ...
func (r *RedisCluster) Hmset(key string, hash map[string]string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hmset(key, hash)
	}
	return ToStringReply(command.run(key))
}

//Hmget ...
func (r *RedisCluster) Hmget(key string, fields ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hmget(key, fields...)
	}
	return ToStringArrayReply(command.run(key))
}

//HincrBy ...
func (r *RedisCluster) HincrBy(key, field string, value int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HincrBy(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//HincrByFloat ...
func (r *RedisCluster) HincrByFloat(key, field string, value float64) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HincrByFloat(key, field, value)
	}
	return ToFloat64Reply(command.run(key))
}

//Hexists ...
func (r *RedisCluster) Hexists(key, field string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hexists(key, field)
	}
	return ToBoolReply(command.run(key))
}

//Hdel ...
func (r *RedisCluster) Hdel(key string, fields ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hdel(key, fields...)
	}
	return ToInt64Reply(command.run(key))
}

//Hlen ...
func (r *RedisCluster) Hlen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hlen(key)
	}
	return ToInt64Reply(command.run(key))
}

//Hkeys ...
func (r *RedisCluster) Hkeys(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hkeys(key)
	}
	return ToStringArrayReply(command.run(key))
}

//Hvals ...
func (r *RedisCluster) Hvals(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hvals(key)
	}
	return ToStringArrayReply(command.run(key))
}

//HgetAll ...
func (r *RedisCluster) HgetAll(key string) (map[string]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HgetAll(key)
	}
	return ToMapReply(command.run(key))
}

//Rpush ...
func (r *RedisCluster) Rpush(key string, strings ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rpush(key, strings...)
	}
	return ToInt64Reply(command.run(key))
}

//Lpush ...
func (r *RedisCluster) Lpush(key string, strings ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lpush(key, strings...)
	}
	return ToInt64Reply(command.run(key))
}

//Llen ...
func (r *RedisCluster) Llen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Llen(key)
	}
	return ToInt64Reply(command.run(key))
}

//Lrange ...
func (r *RedisCluster) Lrange(key string, start, stop int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lrange(key, start, stop)
	}
	return ToStringArrayReply(command.run(key))
}

//Ltrim ...
func (r *RedisCluster) Ltrim(key string, start, stop int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Ltrim(key, start, stop)
	}
	return ToStringReply(command.run(key))
}

//Lindex ...
func (r *RedisCluster) Lindex(key string, index int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lindex(key, index)
	}
	return ToStringReply(command.run(key))
}

//Lset ...
func (r *RedisCluster) Lset(key string, index int64, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lset(key, index, value)
	}
	return ToStringReply(command.run(key))
}

//Lrem ...
func (r *RedisCluster) Lrem(key string, count int64, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lrem(key, count, value)
	}
	return ToInt64Reply(command.run(key))
}

//Lpop ...
func (r *RedisCluster) Lpop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lpop(key)
	}
	return ToStringReply(command.run(key))
}

//Rpop ...
func (r *RedisCluster) Rpop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rpop(key)
	}
	return ToStringReply(command.run(key))
}

//Sadd ...
func (r *RedisCluster) Sadd(key string, members ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sadd(key, members...)
	}
	return ToInt64Reply(command.run(key))
}

//Smembers ...
func (r *RedisCluster) Smembers(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Smembers(key)
	}
	return ToStringArrayReply(command.run(key))
}

//Srem ...
func (r *RedisCluster) Srem(key string, members ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Srem(key, members...)
	}
	return ToInt64Reply(command.run(key))
}

//Spop ...
func (r *RedisCluster) Spop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Spop(key)
	}
	return ToStringReply(command.run(key))
}

//SpopBatch  wait complete
func (r *RedisCluster) SpopBatch(key string, count int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SpopBatch(key, count)
	}
	return ToStringArrayReply(command.run(key))
}

//Scard  wait complete
func (r *RedisCluster) Scard(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Scard(key)
	}
	return ToInt64Reply(command.run(key))
}

//Sismember  wait complete
func (r *RedisCluster) Sismember(key string, member string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sismember(key, member)
	}
	return ToBoolReply(command.run(key))
}

//Srandmember  wait complete
func (r *RedisCluster) Srandmember(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Srandmember(key)
	}
	return ToStringReply(command.run(key))
}

//SrandmemberBatch  wait complete
func (r *RedisCluster) SrandmemberBatch(key string, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SrandmemberBatch(key, count)
	}
	return ToStringArrayReply(command.run(key))
}

//Strlen  wait complete
func (r *RedisCluster) Strlen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Strlen(key)
	}
	return ToInt64Reply(command.run(key))
}

//Zadd  wait complete
func (r *RedisCluster) Zadd(key string, score float64, member string, params ...ZAddParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zadd(key, score, member, params...)
	}
	return ToInt64Reply(command.run(key))
}

//ZaddByMap  wait complete
func (r *RedisCluster) ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZaddByMap(key, scoreMembers, params...)
	}
	return ToInt64Reply(command.run(key))
}

//Zrange  wait complete
func (r *RedisCluster) Zrange(key string, start, end int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zrange(key, start, end)
	}
	return ToStringArrayReply(command.run(key))
}

//Zrem  wait complete
func (r *RedisCluster) Zrem(key string, member ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zrem(key, member...)
	}
	return ToInt64Reply(command.run(key))
}

//Zincrby  wait complete
func (r *RedisCluster) Zincrby(key string, score float64, member string, params ...ZAddParams) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zincrby(key, score, member, params...)
	}
	return ToFloat64Reply(command.run(key))
}

//Zrank  wait complete
func (r *RedisCluster) Zrank(key, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zrank(key, member)
	}
	return ToInt64Reply(command.run(key))
}

//Zrevrank  wait complete
func (r *RedisCluster) Zrevrank(key, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zrevrank(key, member)
	}
	return ToInt64Reply(command.run(key))
}

//Zrevrange  wait complete
func (r *RedisCluster) Zrevrange(key string, start, end int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zrevrange(key, start, end)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrangeWithScores  wait complete
func (r *RedisCluster) ZrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeWithScores(key, start, end)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZrevrangeWithScores  wait complete
func (r *RedisCluster) ZrevrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeWithScores(key, start, end)
	}
	return ToTupleArrayReply(command.run(key))
}

//Zcard  wait complete
func (r *RedisCluster) Zcard(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zcard(key)
	}
	return ToInt64Reply(command.run(key))
}

//Zscore  wait complete
func (r *RedisCluster) Zscore(key, member string) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zscore(key, member)
	}
	return ToFloat64Reply(command.run(key))
}

//Sort  wait complete
func (r *RedisCluster) Sort(key string, sortingParameters ...SortingParams) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sort(key, sortingParameters...)
	}
	return ToStringArrayReply(command.run(key))
}

//Zcount  wait complete
func (r *RedisCluster) Zcount(key string, min string, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zcount(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//ZrangeByScore  wait complete
func (r *RedisCluster) ZrangeByScore(key string, min string, max string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByScore(key, min, max)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrevrangeByScore  wait complete
func (r *RedisCluster) ZrevrangeByScore(key string, max string, min string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeByScore(key, max, min)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrangeByScoreBatch  wait complete
func (r *RedisCluster) ZrangeByScoreBatch(key string, min string, max string, offset int, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByScoreBatch(key, min, max, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrangeByScoreWithScores  wait complete
func (r *RedisCluster) ZrangeByScoreWithScores(key, min, max string) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByScoreWithScores(key, min, max)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZrevrangeByScoreWithScores  wait complete
func (r *RedisCluster) ZrevrangeByScoreWithScores(key, max, min string) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeByScoreWithScores(key, max, min)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZrangeByScoreWithScoresBatch  wait complete
func (r *RedisCluster) ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByScoreWithScoresBatch(key, min, max, offset, count)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZrevrangeByScoreWithScoresBatch  wait complete
func (r *RedisCluster) ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeByScoreWithScoresBatch(key, max, min, offset, count)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZremrangeByRank  wait complete
func (r *RedisCluster) ZremrangeByRank(key string, start, end int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZremrangeByRank(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//ZremrangeByScore  wait complete
func (r *RedisCluster) ZremrangeByScore(key, start, end string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZremrangeByScore(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//Zlexcount  wait complete
func (r *RedisCluster) Zlexcount(key, min, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zlexcount(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//ZrangeByLex  wait complete
func (r *RedisCluster) ZrangeByLex(key, min, max string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByLex(key, min, max)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrangeByLexBatch  wait complete
func (r *RedisCluster) ZrangeByLexBatch(key, min, max string, offset, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrangeByLexBatch(key, min, max, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrevrangeByLex  wait complete
func (r *RedisCluster) ZrevrangeByLex(key, max, min string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeByLex(key, max, min)
	}
	return ToStringArrayReply(command.run(key))
}

//ZrevrangeByLexBatch  wait complete
func (r *RedisCluster) ZrevrangeByLexBatch(key, max, min string, offset, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZrevrangeByLexBatch(key, max, min, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZremrangeByLex  wait complete
func (r *RedisCluster) ZremrangeByLex(key, min, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZremrangeByLex(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//Linsert  wait complete
func (r *RedisCluster) Linsert(key string, where ListOption, pivot, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Linsert(key, where, pivot, value)
	}
	return ToInt64Reply(command.run(key))
}

//Lpushx  wait complete
func (r *RedisCluster) Lpushx(key string, strs ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Lpushx(key, strs...)
	}
	return ToInt64Reply(command.run(key))
}

//Rpushx  wait complete
func (r *RedisCluster) Rpushx(key string, strs ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rpushx(key, strs...)
	}
	return ToInt64Reply(command.run(key))
}

//Echo  wait complete
func (r *RedisCluster) Echo(str string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Echo(str)
	}
	return ToStringReply(command.run(str))
}

// Deprecated
func (r *RedisCluster) Move(key string, dbIndex int) (int64, error) {
	panic("make no sense")
}

//Bitcount  wait complete
func (r *RedisCluster) Bitcount(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Bitcount(key)
	}
	return ToInt64Reply(command.run(key))
}

//BitcountRange  wait complete
func (r *RedisCluster) BitcountRange(key string, start int64, end int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitcountRange(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//Bitpos  wait complete
func (r *RedisCluster) Bitpos(key string, value bool, params ...BitPosParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Bitpos(key, value, params...)
	}
	return ToInt64Reply(command.run(key))
}

//Hscan  wait complete
func (r *RedisCluster) Hscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Hscan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//Sscan  wait complete
func (r *RedisCluster) Sscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sscan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//Zscan  wait complete
func (r *RedisCluster) Zscan(key, cursor string, params ...ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zscan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//Pfadd  wait complete
func (r *RedisCluster) Pfadd(key string, elements ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Pfadd(key, elements...)
	}
	return ToInt64Reply(command.run(key))
}

//Geoadd  wait complete
func (r *RedisCluster) Geoadd(key string, longitude, latitude float64, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Geoadd(key, longitude, latitude, member)
	}
	return ToInt64Reply(command.run(key))
}

//GeoaddByMap  wait complete
func (r *RedisCluster) GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoaddByMap(key, memberCoordinateMap)
	}
	return ToInt64Reply(command.run(key))
}

//Geodist  wait complete
func (r *RedisCluster) Geodist(key string, member1, member2 string, unit ...GeoUnit) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Geodist(key, member1, member2, unit...)
	}
	return ToFloat64Reply(command.run(key))
}

//Geohash  wait complete
func (r *RedisCluster) Geohash(key string, members ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Geohash(key, members...)
	}
	return ToStringArrayReply(command.run(key))
}

//Geopos  wait complete
func (r *RedisCluster) Geopos(key string, members ...string) ([]*GeoCoordinate, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Geopos(key, members...)
	}
	return ToGeoArrayReply(command.run(key))
}

//Georadius  wait complete
func (r *RedisCluster) Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Georadius(key, longitude, latitude, radius, unit, param...)
	}
	return ToGeoArrayReply(command.run(key))
}

//GeoradiusByMember  wait complete
func (r *RedisCluster) GeoradiusByMember(key string, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]*GeoCoordinate, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoradiusByMember(key, member, radius, unit, param...)
	}
	return ToGeoArrayReply(command.run(key))
}

//Bitfield  wait complete
func (r *RedisCluster) Bitfield(key string, arguments ...string) ([]int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Bitfield(key, arguments...)
	}
	return ToInt64ArrayReply(command.run(key))
}

//</editor-fold>

//<editor-fold desc="multikeycommands">

//Del delete one or more keys
// return the number of deleted keys
func (r *RedisCluster) Del(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		defer redis.Close()
		return redis.Del(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//Exists  wait complete
func (r *RedisCluster) Exists(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Exists(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//BlpopTimout  wait complete
func (r *RedisCluster) BlpopTimout(timeout int, keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BlpopTimout(timeout, keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//BrpopTimout  wait complete
func (r *RedisCluster) BrpopTimout(timeout int, keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BrpopTimout(timeout, keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//Blpop  wait complete
func (r *RedisCluster) Blpop(args ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Blpop(args...)
	}
	return ToStringArrayReply(command.runBatch(len(args), args...))
}

//Brpop  wait complete
func (r *RedisCluster) Brpop(args ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Brpop(args...)
	}
	return ToStringArrayReply(command.runBatch(len(args), args...))
}

//Mget  wait complete
func (r *RedisCluster) Mget(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Mget(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//Mset  wait complete
func (r *RedisCluster) Mset(keysvalues ...string) (string, error) {
	keys := make([]string, 0)
	for i := 0; i < len(keysvalues); i++ {
		keys[i] = keysvalues[i*2]
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Mset(keysvalues...)
	}
	return ToStringReply(command.runBatch(len(keys), keys...))
}

//Msetnx  wait complete
func (r *RedisCluster) Msetnx(keysvalues ...string) (int64, error) {
	keys := make([]string, 0)
	for i := 0; i < len(keysvalues); i++ {
		keys[i] = keysvalues[i*2]
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Msetnx(keysvalues...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//Rename  wait complete
func (r *RedisCluster) Rename(oldkey, newkey string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rename(oldkey, newkey)
	}
	return ToStringReply(command.runBatch(2, oldkey, newkey))
}

//Renamenx  wait complete
func (r *RedisCluster) Renamenx(oldkey, newkey string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Renamenx(oldkey, newkey)
	}
	return ToInt64Reply(command.runBatch(2, oldkey, newkey))
}

//Rpoplpush  wait complete
func (r *RedisCluster) Rpoplpush(srckey, dstkey string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rpoplpush(srckey, dstkey)
	}
	return ToStringReply(command.runBatch(2, srckey, dstkey))
}

//Sdiff  wait complete
func (r *RedisCluster) Sdiff(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sdiff(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//Sdiffstore  wait complete
func (r *RedisCluster) Sdiffstore(dstkey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sdiffstore(dstkey, keys...)
	}
	arr := StringStringArrayToStringArray(dstkey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Sinter  wait complete
func (r *RedisCluster) Sinter(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sinter(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//Sinterstore  wait complete
func (r *RedisCluster) Sinterstore(dstkey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sinterstore(dstkey, keys...)
	}
	arr := StringStringArrayToStringArray(dstkey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Smove  wait complete
func (r *RedisCluster) Smove(srckey, dstkey, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Smove(srckey, dstkey, member)
	}
	return ToInt64Reply(command.runBatch(2, srckey, dstkey))
}

//SortMulti  wait complete
func (r *RedisCluster) SortMulti(key, dstkey string, sortingParameters ...SortingParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SortMulti(key, dstkey, sortingParameters...)
	}
	return ToInt64Reply(command.runBatch(2, key, dstkey))
}

//Sunion  wait complete
func (r *RedisCluster) Sunion(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sunion(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//Sunionstore  wait complete
func (r *RedisCluster) Sunionstore(dstkey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sunionstore(dstkey, keys...)
	}
	arr := StringStringArrayToStringArray(dstkey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

// Deprecated do not use
func (r *RedisCluster) Watch(keys ...string) (string, error) {
	panic("not implement")
}

// Deprecated do not use
func (r *RedisCluster) Unwatch() (string, error) {
	panic("not implement")
}

//Zinterstore  wait complete
func (r *RedisCluster) Zinterstore(dstkey string, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zinterstore(dstkey, sets...)
	}
	arr := StringStringArrayToStringArray(dstkey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZinterstoreWithParams ...
func (r *RedisCluster) ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZinterstoreWithParams(dstkey, params, sets...)
	}
	arr := StringStringArrayToStringArray(dstkey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Zunionstore ...
func (r *RedisCluster) Zunionstore(dstkey string, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Zunionstore(dstkey, sets...)
	}
	arr := StringStringArrayToStringArray(dstkey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZunionstoreWithParams ...
func (r *RedisCluster) ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZunionstoreWithParams(dstkey, params, sets...)
	}
	arr := StringStringArrayToStringArray(dstkey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Brpoplpush ...
func (r *RedisCluster) Brpoplpush(source, destination string, timeout int) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Brpoplpush(source, destination, timeout)
	}
	return ToStringReply(command.runBatch(2, source, destination))
}

//Publish ...
func (r *RedisCluster) Publish(channel, message string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Publish(channel, message)
	}
	return ToInt64Reply(command.runWithAnyNode())
}

//Subscribe ...
func (r *RedisCluster) Subscribe(redisPubSub *RedisPubSub, channels ...string) error {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		err := redis.Subscribe(redisPubSub, channels...)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	_, err := command.runWithAnyNode()
	if err != nil {
		return err
	}
	return nil
}

//Psubscribe ...
func (r *RedisCluster) Psubscribe(redisPubSub *RedisPubSub, patterns ...string) error {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		err := redis.Psubscribe(redisPubSub, patterns...)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	_, err := command.runWithAnyNode()
	if err != nil {
		return err
	}
	return nil
}

//Bitop ...
func (r *RedisCluster) Bitop(op BitOP, destKey string, srcKeys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Bitop(op, destKey, srcKeys...)
	}
	arr := StringStringArrayToStringArray(destKey, srcKeys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Scan ...
func (r *RedisCluster) Scan(cursor string, params ...ScanParams) (*ScanResult, error) {
	matchPattern := ""
	param := NewScanParams()
	if len(params) > 0 {
		param = &params[1]
	}
	matchPattern = param.Match()
	if matchPattern == "" {
		return nil, errors.New("only supports SCAN commands with non-empty MATCH patterns")
	}
	if !newRedisClusterHashTagUtil().isClusterCompliantMatchPattern(matchPattern) {
		return nil, errors.New("only supports SCAN commands with MATCH patterns containing hash-tags ( curly-brackets enclosed strings )")
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Scan(cursor, params...)
	}
	return ToScanResultReply(command.run(matchPattern))
}

//Pfmerge ...
func (r *RedisCluster) Pfmerge(destkey string, sourcekeys ...string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Pfmerge(destkey, sourcekeys...)
	}
	arr := StringStringArrayToStringArray(destkey, sourcekeys)
	return ToStringReply(command.runBatch(len(arr), arr...))
}

//Pfcount ...
func (r *RedisCluster) Pfcount(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Pfcount(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

// Deprecated do not use
func (r *RedisCluster) Keys(pattern string) ([]string, error) {
	panic("implement me")
}

// Deprecated do not use
func (r *RedisCluster) RandomKey() (string, error) {
	panic("implement me")
}

//</editor-fold>

//<editor-fold desc="scriptcommands">

//Eval ...
func (r *RedisCluster) Eval(script string, keyCount int, params ...string) (interface{}, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Eval(script, keyCount, params...)
	}
	return command.runBatch(keyCount, params...)
}

//Evalsha ...
func (r *RedisCluster) Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Evalsha(sha1, keyCount, params...)
	}
	return command.runBatch(keyCount, params...)
}

//ScriptExists ...
func (r *RedisCluster) ScriptExists(key string, sha1 ...string) ([]bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ScriptExists(sha1...)
	}
	return ToBoolArrayReply(command.run(key))
}

//ScriptLoad ...
func (r *RedisCluster) ScriptLoad(key, script string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.ConnectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ScriptLoad(script)
	}
	return ToStringReply(command.run(key))
}

//</editor-fold>

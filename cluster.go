package godis

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	masterNodeIndex = 2
)

type redisClusterInfoCache struct {
	nodes sync.Map
	slots sync.Map

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
		if size <= masterNodeIndex {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		for i := masterNodeIndex; i < size; i++ {
			hostInfos := slotInfo[i].([]interface{})
			if len(hostInfos) <= 0 {
				continue
			}
			host, port := r.generateHostAndPort(hostInfos)
			r.setupNodeIfNotExist(false, host, port)
			if i == masterNodeIndex {
				r.assignSlotsToNode(false, slotNums, host, port)
			}
		}
	}
	return nil
}

func (r *redisClusterInfoCache) renewClusterSlots(redis *Redis) error {
	r.wLock.Lock()
	if r.rediscovering {
		return nil
	}
	defer func() {
		r.rediscovering = false
		r.wLock.Unlock()
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

func (r *redisClusterInfoCache) discoverClusterSlots(redis *Redis) error {
	slots, err := redis.ClusterSlots()
	if err != nil {
		return err
	}
	r.slots.Range(func(key, value interface{}) bool {
		r.slots.Delete(key)
		return true
	})
	for _, s := range slots {
		slotInfo := s.([]interface{})
		size := len(slotInfo)
		if size <= masterNodeIndex {
			continue
		}
		slotNums := r.getAssignedSlotArray(slotInfo)
		hostInfos := slotInfo[masterNodeIndex].([]interface{})
		if len(hostInfos) == 0 {
			continue
		}
		host, port := r.generateHostAndPort(hostInfos)
		r.assignSlotsToNode(true, slotNums, host, port)
	}
	return nil
}

func (r *redisClusterInfoCache) reset(lock bool) {
	r.nodes.Range(func(key, value interface{}) bool {
		if value != nil {
			value.(*Pool).Destroy()
		}
		return true
	})
	r.nodes.Range(func(key, value interface{}) bool {
		r.nodes.Delete(key)
		return true
	})
	r.slots.Range(func(key, value interface{}) bool {
		r.slots.Delete(key)
		return true
	})
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
	nodeKey := host + ":" + strconv.Itoa(port)
	existingPool, ok := r.nodes.Load(nodeKey)
	if ok && existingPool != nil {
		return existingPool.(*Pool)
	}
	nodePool := NewPool(r.poolConfig, &Option{
		Host:              host,
		Port:              port,
		ConnectionTimeout: r.connectionTimeout,
		SoTimeout:         r.soTimeout,
		Password:          r.password,
	})
	r.nodes.Store(nodeKey, nodePool)
	return nodePool
}

func (r *redisClusterInfoCache) assignSlotToNode(slot int, host string, port int) {
	targetPool := r.setupNodeIfNotExist(false, host, port)
	r.slots.Store(slot, targetPool)
}

func (r *redisClusterInfoCache) assignSlotsToNode(lock bool, slots []int, host string, port int) {
	targetPool := r.setupNodeIfNotExist(false, host, port)
	for _, slot := range slots {
		r.slots.Store(slot, targetPool)
	}
}

func (r *redisClusterInfoCache) getShuffledNodesPool() []*Pool {
	pools := make([]*Pool, 0)
	r.nodes.Range(func(key, value interface{}) bool {
		if value != nil {
			pools = append(pools, value.(*Pool))
		}
		return true
	})
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
	if value, ok := r.nodes.Load(nodeKey); ok {
		return value.(*Pool)
	}
	return nil
}

func (r *redisClusterInfoCache) getNodes() map[string]*Pool {
	ret := make(map[string]*Pool)
	r.nodes.Range(func(key, value interface{}) bool {
		if value != nil {
			ret[key.(string)] = value.(*Pool)
		}
		return true
	})
	return ret
}

func (r *redisClusterInfoCache) getSlotPool(slot int) *Pool {
	if value, ok := r.slots.Load(slot); ok {
		return value.(*Pool)
	}
	return nil
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
		redis, err := pool.GetResource()
		if err != nil {
			continue
		}
		result, err := redis.Ping()
		if err != nil {
			continue
		}
		if strings.ToUpper(result) == keywordPong.Name {
			return redis, nil
		}
	}
	return nil, newNoReachableClusterNodeError("no reachable node in cluster")
}

func (r *redisClusterConnectionHandler) getConnectionFromSlot(slot int) (*Redis, error) {
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

func (r *redisClusterConnectionHandler) getConnectionFromNode(host string, port int) (*Redis, error) {
	return r.cache.setupNodeIfNotExist(true, host, port).GetResource()
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
		e := strings.Index(key, "}")
		if e > -1 && e != s+1 {
			return key[s+1 : e]
		}
	}
	if returnKeyOnAbsence {
		return key
	}
	return ""
}

type redisClusterCommand struct {
	maxAttempts       int
	connectionHandler *redisClusterConnectionHandler

	execute func(redis *Redis) (interface{}, error)
}

func newRedisClusterCommand(maxAttempts int, connectionHandler *redisClusterConnectionHandler) *redisClusterCommand {
	return &redisClusterCommand{maxAttempts: maxAttempts, connectionHandler: connectionHandler}
}

func (r *redisClusterCommand) run(key string) (interface{}, error) {
	if key == "" {
		return nil, newClusterOperationError("no way to dispatch this command to Redis cluster")
	}
	return r.runWithRetries([]byte(key), r.maxAttempts, false, nil)
}

func (r *redisClusterCommand) runBatch(keyCount int, keys ...string) (interface{}, error) {
	if len(keys) == 0 {
		return nil, newClusterOperationError("no way to dispatch this command to Redis cluster")
	}
	if len(keys) > 1 {
		crc16 := newCRC16()
		slot := crc16.getStringSlot(keys[0])
		for i := 1; i < keyCount; i++ {
			nextSlot := crc16.getStringSlot(keys[i])
			if nextSlot != slot {
				return nil, newClusterOperationError("no way to dispatch this command to Redis cluster,because keys have different slots")
			}
		}
	}
	return r.runWithRetries([]byte(keys[0]), r.maxAttempts, false, nil)
}

func (r *redisClusterCommand) runWithAnyNode() (interface{}, error) {
	connection, err := r.connectionHandler.getConnection()
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

func (r *redisClusterCommand) runWithRetries(key []byte, attempts int, tryRandomNode bool, redirect error) (interface{}, error) {
	if attempts <= 0 {
		return nil, newClusterMaxAttemptsError("too many cluster redirections")
	}
	var connection *Redis
	var err error
	if redirect != nil {
		if connection, err = r.processRedirect(redirect); err != nil {
			return nil, err
		}
	} else {
		if tryRandomNode {
			connection, err = r.connectionHandler.getConnection()
			if err != nil {
				return nil, err
			}
		} else {
			connection, err = r.connectionHandler.getConnectionFromSlot(int(newCRC16().getByteSlot(key)))
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
		if attempts <= 1 {
			r.connectionHandler.renewSlotCache()
			//return nil, err
		}
		return r.runWithRetries(key, attempts-1, tryRandomNode, redirect)
	case *MovedDataError:
		r.connectionHandler.renewSlotCache(connection)
		_ = r.releaseConnection(connection)
		return r.runWithRetries(key, attempts-1, false, err)
	}
	return nil, err
}

func (r *redisClusterCommand) processRedirect(redirect error) (*Redis, error) {
	switch redirect.(type) {
	case *MovedDataError:
		dataError := redirect.(*MovedDataError)
		connection, err := r.connectionHandler.getConnectionFromNode(dataError.Host, dataError.Port)
		if err != nil {
			return nil, err
		}
		return connection, nil
	case *AskDataError:
		dataError := redirect.(*AskDataError)
		connection, err := r.connectionHandler.getConnectionFromNode(dataError.Host, dataError.Port)
		if err != nil {
			return nil, err
		}
		_, err = connection.Asking()
		if err != nil {
			return nil, err
		}
		return connection, nil
	}
	return nil, newRedisError("wrong redirect error")
}

//ClusterOption when you create a new cluster instance ,then you need set some option
type ClusterOption struct {
	Nodes             []string      //cluster nodes, for example: []string{"localhost:7000","localhost:7001"}
	ConnectionTimeout time.Duration //redis connect timeout
	SoTimeout         time.Duration //redis read timeout
	MaxAttempts       int           //when operation or socket is not alright,then program will attempt retry
	Password          string        //cluster redis password
	PoolConfig        *PoolConfig   //redis connection pool config
}

//RedisCluster redis cluster tool
type RedisCluster struct {
	MaxAttempts       int
	connectionHandler *redisClusterConnectionHandler
}

//NewRedisCluster constructor
func NewRedisCluster(option *ClusterOption) *RedisCluster {
	if option.MaxAttempts <= 0 {
		option.MaxAttempts = 5
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
		connectionHandler: newRedisClusterConnectionHandler(option.Nodes, conTimeout, soTimeout, option.Password, option.PoolConfig),
	}
}

//<editor-fold desc="rediscommands">

//Set set key/value,without timeout
func (r *RedisCluster) Set(key, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Set(key, value)
	}
	return ToStringReply(command.run(key))
}

//SetWithParamsAndTime see redis command
func (r *RedisCluster) SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetWithParamsAndTime(key, value, nxxx, expx, time)
	}
	return ToStringReply(command.run(key))
}

//SetWithParams see redis command
func (r *RedisCluster) SetWithParams(key, value, nxxx string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetWithParams(key, value, nxxx)
	}
	return ToStringReply(command.run(key))
}

//Get see redis command
func (r *RedisCluster) Get(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Get(key)
	}
	return ToStringReply(command.run(key))
}

//Persist see redis command
func (r *RedisCluster) Persist(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Persist(key)
	}
	return ToInt64Reply(command.run(key))
}

//Type see redis command
func (r *RedisCluster) Type(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Type(key)
	}
	return ToStringReply(command.run(key))
}

//Expire see redis command
func (r *RedisCluster) Expire(key string, seconds int) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Expire(key, seconds)
	}
	return ToInt64Reply(command.run(key))
}

//PExpire see redis command
func (r *RedisCluster) PExpire(key string, milliseconds int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PExpire(key, milliseconds)
	}
	return ToInt64Reply(command.run(key))
}

//ExpireAt see redis command
func (r *RedisCluster) ExpireAt(key string, unixtime int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ExpireAt(key, unixtime)
	}
	return ToInt64Reply(command.run(key))
}

//PExpireAt see redis command
func (r *RedisCluster) PExpireAt(key string, millisecondsTimestamp int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PExpireAt(key, millisecondsTimestamp)
	}
	return ToInt64Reply(command.run(key))
}

//TTL see redis command
func (r *RedisCluster) TTL(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.TTL(key)
	}
	return ToInt64Reply(command.run(key))
}

//PTTL see redis command
func (r *RedisCluster) PTTL(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PTTL(key)
	}
	return ToInt64Reply(command.run(key))
}

//SetBitWithBool see redis command
func (r *RedisCluster) SetBitWithBool(key string, offset int64, value bool) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetBitWithBool(key, offset, value)
	}
	return ToBoolReply(command.run(key))
}

//SetBit see redis command
func (r *RedisCluster) SetBit(key string, offset int64, value string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetBit(key, offset, value)
	}
	return ToBoolReply(command.run(key))
}

//GetBit see redis command
func (r *RedisCluster) GetBit(key string, offset int64) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GetBit(key, offset)
	}
	return ToBoolReply(command.run(key))
}

//SetRange see redis command
func (r *RedisCluster) SetRange(key string, offset int64, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetRange(key, offset, value)
	}
	return ToInt64Reply(command.run(key))
}

//GetRange see redis command
func (r *RedisCluster) GetRange(key string, startOffset, endOffset int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GetRange(key, startOffset, endOffset)
	}
	return ToStringReply(command.run(key))
}

//GetSet see redis command
func (r *RedisCluster) GetSet(key, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GetSet(key, value)
	}
	return ToStringReply(command.run(key))
}

//SetNx see redis command
func (r *RedisCluster) SetNx(key, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetNx(key, value)
	}
	return ToInt64Reply(command.run(key))
}

//SetEx see redis command
func (r *RedisCluster) SetEx(key string, seconds int, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SetEx(key, seconds, value)
	}
	return ToStringReply(command.run(key))
}

//PSetEx see redis command
func (r *RedisCluster) PSetEx(key string, milliseconds int64, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PSetEx(key, milliseconds, value)
	}
	return ToStringReply(command.run(key))
}

//DecrBy see redis command
func (r *RedisCluster) DecrBy(key string, decrement int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.DecrBy(key, decrement)
	}
	return ToInt64Reply(command.run(key))
}

//Decr see redis command
func (r *RedisCluster) Decr(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Decr(key)
	}
	return ToInt64Reply(command.run(key))
}

//IncrBy see redis command
func (r *RedisCluster) IncrBy(key string, increment int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.IncrBy(key, increment)
	}
	return ToInt64Reply(command.run(key))
}

//IncrByFloat see redis command
func (r *RedisCluster) IncrByFloat(key string, increment float64) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.IncrByFloat(key, increment)
	}
	return ToFloat64Reply(command.run(key))
}

//Incr see redis command
func (r *RedisCluster) Incr(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Incr(key)
	}
	return ToInt64Reply(command.run(key))
}

//Append see redis command
func (r *RedisCluster) Append(key, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Append(key, value)
	}
	return ToInt64Reply(command.run(key))
}

//SubStr see redis command
func (r *RedisCluster) SubStr(key string, start, end int) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SubStr(key, start, end)
	}
	return ToStringReply(command.run(key))
}

//HSet see redis command
func (r *RedisCluster) HSet(key, field string, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HSet(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//HGet see redis command
func (r *RedisCluster) HGet(key, field string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HGet(key, field)
	}
	return ToStringReply(command.run(key))
}

//HSetNx see redis command
func (r *RedisCluster) HSetNx(key, field, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HSetNx(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//HMSet see redis command
func (r *RedisCluster) HMSet(key string, hash map[string]string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HMSet(key, hash)
	}
	return ToStringReply(command.run(key))
}

//HMGet see redis command
func (r *RedisCluster) HMGet(key string, fields ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HMGet(key, fields...)
	}
	return ToStringArrayReply(command.run(key))
}

//HIncrBy see redis command
func (r *RedisCluster) HIncrBy(key, field string, value int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HIncrBy(key, field, value)
	}
	return ToInt64Reply(command.run(key))
}

//HIncrByFloat see redis command
func (r *RedisCluster) HIncrByFloat(key, field string, value float64) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HIncrByFloat(key, field, value)
	}
	return ToFloat64Reply(command.run(key))
}

//HExists see redis command
func (r *RedisCluster) HExists(key, field string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HExists(key, field)
	}
	return ToBoolReply(command.run(key))
}

//HDel see redis command
func (r *RedisCluster) HDel(key string, fields ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HDel(key, fields...)
	}
	return ToInt64Reply(command.run(key))
}

//HLen see redis command
func (r *RedisCluster) HLen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HLen(key)
	}
	return ToInt64Reply(command.run(key))
}

//HKeys see redis command
func (r *RedisCluster) HKeys(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HKeys(key)
	}
	return ToStringArrayReply(command.run(key))
}

//HVals see redis command
func (r *RedisCluster) HVals(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HVals(key)
	}
	return ToStringArrayReply(command.run(key))
}

//HGetAll see redis command
func (r *RedisCluster) HGetAll(key string) (map[string]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HGetAll(key)
	}
	return ToMapReply(command.run(key))
}

//RPush see redis command
func (r *RedisCluster) RPush(key string, strings ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.RPush(key, strings...)
	}
	return ToInt64Reply(command.run(key))
}

//LPush see redis command
func (r *RedisCluster) LPush(key string, strings ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LPush(key, strings...)
	}
	return ToInt64Reply(command.run(key))
}

//LLen see redis command
func (r *RedisCluster) LLen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LLen(key)
	}
	return ToInt64Reply(command.run(key))
}

//LRange see redis command
func (r *RedisCluster) LRange(key string, start, stop int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LRange(key, start, stop)
	}
	return ToStringArrayReply(command.run(key))
}

//LTrim see redis command
func (r *RedisCluster) LTrim(key string, start, stop int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LTrim(key, start, stop)
	}
	return ToStringReply(command.run(key))
}

//LIndex see redis command
func (r *RedisCluster) LIndex(key string, index int64) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LIndex(key, index)
	}
	return ToStringReply(command.run(key))
}

//LSet see redis command
func (r *RedisCluster) LSet(key string, index int64, value string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LSet(key, index, value)
	}
	return ToStringReply(command.run(key))
}

//LRem see redis command
func (r *RedisCluster) LRem(key string, count int64, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LRem(key, count, value)
	}
	return ToInt64Reply(command.run(key))
}

//LPop see redis command
func (r *RedisCluster) LPop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LPop(key)
	}
	return ToStringReply(command.run(key))
}

//RPop see redis command
func (r *RedisCluster) RPop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.RPop(key)
	}
	return ToStringReply(command.run(key))
}

//SAdd see redis command
func (r *RedisCluster) SAdd(key string, members ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SAdd(key, members...)
	}
	return ToInt64Reply(command.run(key))
}

//SMembers see redis command
func (r *RedisCluster) SMembers(key string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SMembers(key)
	}
	return ToStringArrayReply(command.run(key))
}

//SRem see redis command
func (r *RedisCluster) SRem(key string, members ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SRem(key, members...)
	}
	return ToInt64Reply(command.run(key))
}

//SPop see redis command
func (r *RedisCluster) SPop(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SPop(key)
	}
	return ToStringReply(command.run(key))
}

//SPopBatch  see comment in redis.go
func (r *RedisCluster) SPopBatch(key string, count int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SPopBatch(key, count)
	}
	return ToStringArrayReply(command.run(key))
}

//SCard  see comment in redis.go
func (r *RedisCluster) SCard(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SCard(key)
	}
	return ToInt64Reply(command.run(key))
}

//SIsMember  see comment in redis.go
func (r *RedisCluster) SIsMember(key string, member string) (bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SIsMember(key, member)
	}
	return ToBoolReply(command.run(key))
}

//SRandMember  see comment in redis.go
func (r *RedisCluster) SRandMember(key string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SRandMember(key)
	}
	return ToStringReply(command.run(key))
}

//SRandMemberBatch  see comment in redis.go
func (r *RedisCluster) SRandMemberBatch(key string, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SRandMemberBatch(key, count)
	}
	return ToStringArrayReply(command.run(key))
}

//StrLen  see comment in redis.go
func (r *RedisCluster) StrLen(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.StrLen(key)
	}
	return ToInt64Reply(command.run(key))
}

//ZAdd  see comment in redis.go
func (r *RedisCluster) ZAdd(key string, score float64, member string, params ...*ZAddParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZAdd(key, score, member, params...)
	}
	return ToInt64Reply(command.run(key))
}

//ZAddByMap  see comment in redis.go
func (r *RedisCluster) ZAddByMap(key string, scoreMembers map[string]float64, params ...*ZAddParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZAddByMap(key, scoreMembers, params...)
	}
	return ToInt64Reply(command.run(key))
}

//ZRange  see comment in redis.go
func (r *RedisCluster) ZRange(key string, start, end int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRange(key, start, end)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRem  see comment in redis.go
func (r *RedisCluster) ZRem(key string, member ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRem(key, member...)
	}
	return ToInt64Reply(command.run(key))
}

//ZIncrBy  see comment in redis.go
func (r *RedisCluster) ZIncrBy(key string, score float64, member string, params ...*ZAddParams) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZIncrBy(key, score, member, params...)
	}
	return ToFloat64Reply(command.run(key))
}

//ZRank  see comment in redis.go
func (r *RedisCluster) ZRank(key, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRank(key, member)
	}
	return ToInt64Reply(command.run(key))
}

//ZRevRank  see comment in redis.go
func (r *RedisCluster) ZRevRank(key, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRank(key, member)
	}
	return ToInt64Reply(command.run(key))
}

//ZRevRange  see comment in redis.go
func (r *RedisCluster) ZRevRange(key string, start, end int64) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRange(key, start, end)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRangeWithScores  see comment in redis.go
func (r *RedisCluster) ZRangeWithScores(key string, start, end int64) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeWithScores(key, start, end)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZRevRangeWithScores  see comment in redis.go
func (r *RedisCluster) ZRevRangeWithScores(key string, start, end int64) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeWithScores(key, start, end)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZCard  see comment in redis.go
func (r *RedisCluster) ZCard(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZCard(key)
	}
	return ToInt64Reply(command.run(key))
}

//ZScore  see comment in redis.go
func (r *RedisCluster) ZScore(key, member string) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZScore(key, member)
	}
	return ToFloat64Reply(command.run(key))
}

//Sort  see comment in redis.go
func (r *RedisCluster) Sort(key string, sortingParameters ...SortingParams) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Sort(key, sortingParameters...)
	}
	return ToStringArrayReply(command.run(key))
}

//ZCount  see comment in redis.go
func (r *RedisCluster) ZCount(key string, min string, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZCount(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//ZRangeByScore  see comment in redis.go
func (r *RedisCluster) ZRangeByScore(key string, min string, max string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByScore(key, min, max)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRevRangeByScore  see comment in redis.go
func (r *RedisCluster) ZRevRangeByScore(key string, max string, min string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeByScore(key, max, min)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRangeByScoreBatch  see comment in redis.go
func (r *RedisCluster) ZRangeByScoreBatch(key string, min string, max string, offset int, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByScoreBatch(key, min, max, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRangeByScoreWithScores  see comment in redis.go
func (r *RedisCluster) ZRangeByScoreWithScores(key, min, max string) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByScoreWithScores(key, min, max)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZRevRangeByScoreWithScores  see comment in redis.go
func (r *RedisCluster) ZRevRangeByScoreWithScores(key, max, min string) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeByScoreWithScores(key, max, min)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZRangeByScoreWithScoresBatch  see comment in redis.go
func (r *RedisCluster) ZRangeByScoreWithScoresBatch(key, min, max string, offset, count int) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByScoreWithScoresBatch(key, min, max, offset, count)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZRevRangeByScoreWithScoresBatch  see comment in redis.go
func (r *RedisCluster) ZRevRangeByScoreWithScoresBatch(key, max, min string, offset, count int) ([]Tuple, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeByScoreWithScoresBatch(key, max, min, offset, count)
	}
	return ToTupleArrayReply(command.run(key))
}

//ZRemRangeByRank  see comment in redis.go
func (r *RedisCluster) ZRemRangeByRank(key string, start, end int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRemRangeByRank(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//ZRemRangeByScore  see comment in redis.go
func (r *RedisCluster) ZRemRangeByScore(key, start, end string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRemRangeByScore(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//ZLexCount  see comment in redis.go
func (r *RedisCluster) ZLexCount(key, min, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZLexCount(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//ZRangeByLex  see comment in redis.go
func (r *RedisCluster) ZRangeByLex(key, min, max string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByLex(key, min, max)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRangeByLexBatch  see comment in redis.go
func (r *RedisCluster) ZRangeByLexBatch(key, min, max string, offset, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRangeByLexBatch(key, min, max, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRevRangeByLex  see comment in redis.go
func (r *RedisCluster) ZRevRangeByLex(key, max, min string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeByLex(key, max, min)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRevRangeByLexBatch  see comment in redis.go
func (r *RedisCluster) ZRevRangeByLexBatch(key, max, min string, offset, count int) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRevRangeByLexBatch(key, max, min, offset, count)
	}
	return ToStringArrayReply(command.run(key))
}

//ZRemRangeByLex  see comment in redis.go
func (r *RedisCluster) ZRemRangeByLex(key, min, max string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZRemRangeByLex(key, min, max)
	}
	return ToInt64Reply(command.run(key))
}

//LInsert  see comment in redis.go
func (r *RedisCluster) LInsert(key string, where ListOption, pivot, value string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LInsert(key, where, pivot, value)
	}
	return ToInt64Reply(command.run(key))
}

//LPushX  see comment in redis.go
func (r *RedisCluster) LPushX(key string, strs ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.LPushX(key, strs...)
	}
	return ToInt64Reply(command.run(key))
}

//RPushX  see comment in redis.go
func (r *RedisCluster) RPushX(key string, strs ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.RPushX(key, strs...)
	}
	return ToInt64Reply(command.run(key))
}

//Echo  see comment in redis.go
func (r *RedisCluster) Echo(str string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Echo(str)
	}
	return ToStringReply(command.run(str))
}

//BitCount  see comment in redis.go
func (r *RedisCluster) BitCount(key string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitCount(key)
	}
	return ToInt64Reply(command.run(key))
}

//BitCountRange  see comment in redis.go
func (r *RedisCluster) BitCountRange(key string, start int64, end int64) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitCountRange(key, start, end)
	}
	return ToInt64Reply(command.run(key))
}

//BitPos  see comment in redis.go
func (r *RedisCluster) BitPos(key string, value bool, params ...BitPosParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitPos(key, value, params...)
	}
	return ToInt64Reply(command.run(key))
}

//HScan  see comment in redis.go
func (r *RedisCluster) HScan(key, cursor string, params ...*ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.HScan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//SScan  see comment in redis.go
func (r *RedisCluster) SScan(key, cursor string, params ...*ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SScan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//ZScan  see comment in redis.go
func (r *RedisCluster) ZScan(key, cursor string, params ...*ScanParams) (*ScanResult, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZScan(key, cursor, params...)
	}
	return ToScanResultReply(command.run(key))
}

//PfAdd  see comment in redis.go
func (r *RedisCluster) PfAdd(key string, elements ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PfAdd(key, elements...)
	}
	return ToInt64Reply(command.run(key))
}

//GeoAdd  see comment in redis.go
func (r *RedisCluster) GeoAdd(key string, longitude, latitude float64, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoAdd(key, longitude, latitude, member)
	}
	return ToInt64Reply(command.run(key))
}

//GeoAddByMap  see comment in redis.go
func (r *RedisCluster) GeoAddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoAddByMap(key, memberCoordinateMap)
	}
	return ToInt64Reply(command.run(key))
}

//GeoDist  see comment in redis.go
func (r *RedisCluster) GeoDist(key string, member1, member2 string, unit ...GeoUnit) (float64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoDist(key, member1, member2, unit...)
	}
	return ToFloat64Reply(command.run(key))
}

//GeoHash  see comment in redis.go
func (r *RedisCluster) GeoHash(key string, members ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoHash(key, members...)
	}
	return ToStringArrayReply(command.run(key))
}

//GeoPos  see comment in redis.go
func (r *RedisCluster) GeoPos(key string, members ...string) ([]*GeoCoordinate, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoPos(key, members...)
	}
	return ToGeoArrayReply(command.run(key))
}

//GeoRadius  see comment in redis.go
func (r *RedisCluster) GeoRadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...*GeoRadiusParam) ([]GeoRadiusResponse, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoRadius(key, longitude, latitude, radius, unit, param...)
	}
	return ToGeoRespArrayReply(command.run(key))
}

//GeoRadiusByMember  see comment in redis.go
func (r *RedisCluster) GeoRadiusByMember(key string, member string, radius float64, unit GeoUnit, param ...*GeoRadiusParam) ([]GeoRadiusResponse, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.GeoRadiusByMember(key, member, radius, unit, param...)
	}
	return ToGeoRespArrayReply(command.run(key))
}

//BitField  see comment in redis.go
func (r *RedisCluster) BitField(key string, arguments ...string) ([]int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitField(key, arguments...)
	}
	return ToInt64ArrayReply(command.run(key))
}

//</editor-fold>

//<editor-fold desc="multikeycommands">

//Del delete one or more keys
// return the number of deleted keys
func (r *RedisCluster) Del(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		//defer redis.Close()
		return redis.Del(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//Exists  see comment in redis.go
func (r *RedisCluster) Exists(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Exists(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//BLPopTimeout  see comment in redis.go
func (r *RedisCluster) BLPopTimeout(timeout int, keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BLPopTimeout(timeout, keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//BRPopTimeout  see comment in redis.go
func (r *RedisCluster) BRPopTimeout(timeout int, keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BRPopTimeout(timeout, keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//BLPop  see comment in redis.go
func (r *RedisCluster) BLPop(args ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BLPop(args...)
	}
	return ToStringArrayReply(command.runBatch(len(args), args...))
}

//BRPop  see comment in redis.go
func (r *RedisCluster) BRPop(args ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BRPop(args...)
	}
	return ToStringArrayReply(command.runBatch(len(args), args...))
}

//MGet  see comment in redis.go
func (r *RedisCluster) MGet(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.MGet(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//MSet  see comment in redis.go
func (r *RedisCluster) MSet(kvs ...string) (string, error) {
	keys := make([]string, 0)
	for i := 0; i < len(kvs)/2; i++ {
		keys = append(keys, kvs[i*2])
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.MSet(kvs...)
	}
	return ToStringReply(command.runBatch(len(keys), keys...))
}

//MSetNx  see comment in redis.go
func (r *RedisCluster) MSetNx(kvs ...string) (int64, error) {
	keys := make([]string, 0)
	for i := 0; i < len(kvs)/2; i++ {
		keys = append(keys, kvs[i*2])
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.MSetNx(kvs...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//Rename  see comment in redis.go
func (r *RedisCluster) Rename(oldKey, newKey string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Rename(oldKey, newKey)
	}
	return ToStringReply(command.runBatch(2, oldKey, newKey))
}

//RenameNx  see comment in redis.go
func (r *RedisCluster) RenameNx(oldKey, newKey string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.RenameNx(oldKey, newKey)
	}
	return ToInt64Reply(command.runBatch(2, oldKey, newKey))
}

//RPopLPush  see comment in redis.go
func (r *RedisCluster) RPopLPush(srcKey, destKey string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.RPopLPush(srcKey, destKey)
	}
	return ToStringReply(command.runBatch(2, srcKey, destKey))
}

//SDiff  see comment in redis.go
func (r *RedisCluster) SDiff(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SDiff(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//SDiffStore  see comment in redis.go
func (r *RedisCluster) SDiffStore(destKey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SDiffStore(destKey, keys...)
	}
	arr := StringStringArrayToStringArray(destKey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//SInter  see comment in redis.go
func (r *RedisCluster) SInter(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SInter(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//SInterStore  see comment in redis.go
func (r *RedisCluster) SInterStore(destKey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SInterStore(destKey, keys...)
	}
	arr := StringStringArrayToStringArray(destKey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//SMove  see comment in redis.go
func (r *RedisCluster) SMove(srcKey, destKey, member string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SMove(srcKey, destKey, member)
	}
	return ToInt64Reply(command.runBatch(2, srcKey, destKey))
}

//SortStore  see comment in redis.go
func (r *RedisCluster) SortStore(key, destKey string, sortingParameters ...SortingParams) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SortStore(key, destKey, sortingParameters...)
	}
	return ToInt64Reply(command.runBatch(2, key, destKey))
}

//SUnion  see comment in redis.go
func (r *RedisCluster) SUnion(keys ...string) ([]string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SUnion(keys...)
	}
	return ToStringArrayReply(command.runBatch(len(keys), keys...))
}

//SUnionStore  see comment in redis.go
func (r *RedisCluster) SUnionStore(destKey string, keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.SUnionStore(destKey, keys...)
	}
	arr := StringStringArrayToStringArray(destKey, keys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZInterStore  see comment in redis.go
func (r *RedisCluster) ZInterStore(destKey string, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZInterStore(destKey, sets...)
	}
	arr := StringStringArrayToStringArray(destKey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZInterStoreWithParams see redis command
func (r *RedisCluster) ZInterStoreWithParams(destKey string, params ZParams, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZInterStoreWithParams(destKey, params, sets...)
	}
	arr := StringStringArrayToStringArray(destKey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZUnionStore see redis command
func (r *RedisCluster) ZUnionStore(destKey string, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZUnionStore(destKey, sets...)
	}
	arr := StringStringArrayToStringArray(destKey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//ZUnionStoreWithParams see redis command
func (r *RedisCluster) ZUnionStoreWithParams(destKey string, params ZParams, sets ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ZUnionStoreWithParams(destKey, params, sets...)
	}
	arr := StringStringArrayToStringArray(destKey, sets)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//BRPopLPush see redis command
func (r *RedisCluster) BRPopLPush(source, destination string, timeout int) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BRPopLPush(source, destination, timeout)
	}
	return ToStringReply(command.runBatch(2, source, destination))
}

//Publish see redis command
func (r *RedisCluster) Publish(channel, message string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Publish(channel, message)
	}
	return ToInt64Reply(command.runWithAnyNode())
}

//Subscribe see redis command
func (r *RedisCluster) Subscribe(redisPubSub *RedisPubSub, channels ...string) error {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
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

//PSubscribe see redis command
func (r *RedisCluster) PSubscribe(redisPubSub *RedisPubSub, patterns ...string) error {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		err := redis.PSubscribe(redisPubSub, patterns...)
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

//BitOp see redis command
func (r *RedisCluster) BitOp(op BitOP, destKey string, srcKeys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.BitOp(op, destKey, srcKeys...)
	}
	arr := StringStringArrayToStringArray(destKey, srcKeys)
	return ToInt64Reply(command.runBatch(len(arr), arr...))
}

//Scan see redis command
func (r *RedisCluster) Scan(cursor string, params ...*ScanParams) (*ScanResult, error) {
	matchPattern := ""
	param := NewScanParams()
	if len(params) > 0 {
		param = params[0]
	}
	matchPattern = param.Match()
	if matchPattern == "" {
		return nil, errors.New("only supports SCAN commands with non-empty MATCH patterns")
	}
	if !newRedisClusterHashTagUtil().isClusterCompliantMatchPattern(matchPattern) {
		return nil, errors.New("only supports SCAN commands with MATCH patterns containing hash-tags ( curly-brackets enclosed strings )")
	}
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Scan(cursor, params...)
	}
	return ToScanResultReply(command.run(matchPattern))
}

//PfMerge see redis command
func (r *RedisCluster) PfMerge(destkey string, sourcekeys ...string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PfMerge(destkey, sourcekeys...)
	}
	arr := StringStringArrayToStringArray(destkey, sourcekeys)
	return ToStringReply(command.runBatch(len(arr), arr...))
}

//PfCount see redis command
func (r *RedisCluster) PfCount(keys ...string) (int64, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.PfCount(keys...)
	}
	return ToInt64Reply(command.runBatch(len(keys), keys...))
}

//</editor-fold>

//<editor-fold desc="scriptcommands">

//Eval see redis command
func (r *RedisCluster) Eval(script string, keyCount int, params ...string) (interface{}, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.Eval(script, keyCount, params...)
	}
	return command.runBatch(keyCount, params...)
}

//EvalSha see redis command
func (r *RedisCluster) EvalSha(sha1 string, keyCount int, params ...string) (interface{}, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.EvalSha(sha1, keyCount, params...)
	}
	return command.runBatch(keyCount, params...)
}

//ScriptExists see redis command
func (r *RedisCluster) ScriptExists(key string, sha1 ...string) ([]bool, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ScriptExists(sha1...)
	}
	return ToBoolArrayReply(command.run(key))
}

//ScriptLoad see redis command
func (r *RedisCluster) ScriptLoad(key, script string) (string, error) {
	command := newRedisClusterCommand(r.MaxAttempts, r.connectionHandler)
	command.execute = func(redis *Redis) (interface{}, error) {
		return redis.ScriptLoad(script)
	}
	return ToStringReply(command.run(key))
}

//</editor-fold>

package godis

import (
	"strconv"
)

//Client send command to redis, and receive data from redis
type client struct {
	*connection
	Password  string
	Db        int
	isInMulti bool
	isInWatch bool
}

//NewClient
func newClient(option *Option) *client {
	db := 0
	if option.Db != 0 {
		db = option.Db
	}
	client := &client{
		Password:  option.Password,
		Db:        db,
		isInMulti: false,
		isInWatch: false,
	}
	client.connection = newConnection(option.Host, option.Port, option.ConnectionTimeout, option.SoTimeout)
	return client
}

func (c *client) host() string {
	return c.connection.host
}

func (c *client) port() int {
	return c.connection.port
}

//Receive
func (c *client) receive() (interface{}, error) {
	return c.connection.getOne()
}

//Connect
func (c *client) connect() error {
	err := c.connection.connect()
	if err != nil {
		return err
	}
	if c.Password != "" {
		err = c.auth(c.Password)
		if err != nil {
			return err
		}
		_, err = c.getStatusCodeReply()
		if err != nil {
			return err
		}
	}
	if c.Db > 0 {
		err = c.selectDb(c.Db)
		if err != nil {
			return err
		}
		_, err = c.getStatusCodeReply()
		if err != nil {
			return err
		}
	}
	return nil
}

//Close
func (c *client) close() error {
	return c.connection.close()
}

//Ping
func (c *client) ping() error {
	return c.sendCommand(cmdPing)
}

//Quit
func (c *client) quit() error {
	return c.sendCommand(cmdQuit)
}

//Info
func (c *client) info(section ...string) error {
	return c.sendCommand(cmdInfo, StringArrayToByteArray(section)...)
}

//Auth
func (c *client) auth(password string) error {
	c.Password = password
	return c.sendCommand(cmdAuth, []byte(password))
}

//Select
func (c *client) selectDb(index int) error {
	return c.sendCommand(cmdSelect, IntToByteArray(index))
}

func (c *client) set(key, value string) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value))
}

func (c *client) setWithParamsAndTime(key, value, nxxx, expx string, time int64) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value), []byte(nxxx), []byte(expx), Int64ToByteArray(time))
}

func (c *client) setWithParams(key, value, nxxx string) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value), []byte(nxxx))
}

func (c *client) get(key string) error {
	return c.sendCommand(cmdGet, []byte(key))
}

func (c *client) del(keys ...string) error {
	return c.sendCommand(cmdDel, StringArrayToByteArray(keys)...)
}

func (c *client) exists(keys ...string) error {
	return c.sendCommand(cmdExists, StringArrayToByteArray(keys)...)
}

func (c *client) typeKey(key string) error {
	return c.sendCommand(cmdType, []byte(key))
}

func (c *client) keys(pattern string) error {
	return c.sendCommand(cmdKeys, []byte(pattern))
}

func (c *client) rename(oldKey, newKey string) error {
	return c.sendCommand(cmdRename, []byte(oldKey), []byte(newKey))
}

func (c *client) renamenx(oldKey, newKey string) error {
	return c.sendCommand(cmdRenamenx, []byte(oldKey), []byte(newKey))
}

func (c *client) expire(key string, seconds int) error {
	return c.sendCommand(cmdExpire, []byte(key), IntToByteArray(seconds))
}

func (c *client) expireAt(key string, unixTime int64) error {
	return c.sendCommand(cmdExpireat, []byte(key), Int64ToByteArray(unixTime))
}

func (c *client) pexpire(key string, milliseconds int64) error {
	return c.sendCommand(cmdPexpire, []byte(key), Int64ToByteArray(milliseconds))
}

func (c *client) pexpireAt(key string, unixTime int64) error {
	return c.sendCommand(cmdPexpireat, []byte(key), Int64ToByteArray(unixTime))
}

func (c *client) ttl(key string) error {
	return c.sendCommand(cmdTTL, []byte(key))
}

func (c *client) pttl(key string) error {
	return c.sendCommand(cmdPttl, []byte(key))
}

func (c *client) move(key string, dbIndex int) error {
	return c.sendCommand(cmdMove, []byte(key), IntToByteArray(dbIndex))
}

func (c *client) getSet(key, value string) error {
	return c.sendCommand(cmdGetset, []byte(key), []byte(value))
}

func (c *client) mget(keys ...string) error {
	return c.sendCommand(cmdMget, StringArrayToByteArray(keys)...)
}

func (c *client) setnx(key, value string) error {
	return c.sendCommand(cmdSetnx, []byte(key), []byte(value))
}

func (c *client) setex(key string, seconds int, value string) error {
	return c.sendCommand(cmdSetex, []byte(key), IntToByteArray(seconds), []byte(value))
}

func (c *client) psetex(key string, milliseconds int64, value string) error {
	return c.sendCommand(cmdSetex, []byte(key), Int64ToByteArray(milliseconds), []byte(value))
}

func (c *client) mset(keysvalues ...string) error {
	return c.sendCommand(cmdMset, StringArrayToByteArray(keysvalues)...)
}

func (c *client) msetnx(keysvalues ...string) error {
	return c.sendCommand(cmdMsetnx, StringArrayToByteArray(keysvalues)...)
}

func (c *client) decrBy(key string, decrement int64) error {
	return c.sendCommand(cmdDecrby, []byte(key), Int64ToByteArray(decrement))
}

func (c *client) decr(key string) error {
	return c.sendCommand(cmdDecr, []byte(key))
}

func (c *client) incrBy(key string, increment int64) error {
	return c.sendCommand(cmdIncrby, []byte(key), Int64ToByteArray(increment))
}

func (c *client) incr(key string) error {
	return c.sendCommand(cmdIncr, []byte(key))
}

func (c *client) append(key, value string) error {
	return c.sendCommand(cmdAppend, []byte(key), []byte(value))
}

func (c *client) substr(key string, start, end int) error {
	return c.sendCommand(cmdSubstr, []byte(key), IntToByteArray(start), IntToByteArray(end))
}

func (c *client) hset(key, field, value string) error {
	return c.sendCommand(cmdHset, []byte(key), []byte(field), []byte(value))
}

func (c *client) hget(key, field string) error {
	return c.sendCommand(cmdHget, []byte(key), []byte(field))
}

func (c *client) hsetnx(key, field, value string) error {
	return c.sendCommand(cmdHsetnx, []byte(key), []byte(field), []byte(value))
}

func (c *client) hmset(key string, hash map[string]string) error {
	params := make([][]byte, 0)
	params = append(params, []byte(key))
	for k, v := range hash {
		params = append(params, []byte(k))
		params = append(params, []byte(v))
	}
	return c.sendCommand(cmdHmset, params...)
}

func (c *client) hmget(key string, fields ...string) error {
	return c.sendCommand(cmdHmget, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) hincrBy(key, field string, increment int64) error {
	return c.sendCommand(cmdHincrby, []byte(key), []byte(field), Int64ToByteArray(increment))
}

func (c *client) hexists(key, field string) error {
	return c.sendCommand(cmdHexists, []byte(key), []byte(field))
}

func (c *client) hdel(key string, fields ...string) error {
	return c.sendCommand(cmdHdel, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) hlen(key string) error {
	return c.sendCommand(cmdHlen, []byte(key))
}

func (c *client) hkeys(key string) error {
	return c.sendCommand(cmdHkeys, []byte(key))
}

func (c *client) hvals(key string) error {
	return c.sendCommand(cmdHvals, []byte(key))
}

func (c *client) hgetAll(key string) error {
	return c.sendCommand(cmdHgetall, []byte(key))
}

func (c *client) rpush(key string, fields ...string) error {
	return c.sendCommand(cmdRpush, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) lpush(key string, fields ...string) error {
	return c.sendCommand(cmdRpush, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) llen(key string) error {
	return c.sendCommand(cmdLlen, []byte(key))
}

func (c *client) lrange(key string, start, end int64) error {
	return c.sendCommand(cmdLrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) ltrim(key string, start, end int64) error {
	return c.sendCommand(cmdLtrim, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) lindex(key string, index int64) error {
	return c.sendCommand(cmdLindex, []byte(key), Int64ToByteArray(index))
}

func (c *client) lset(key string, index int64, value string) error {
	return c.sendCommand(cmdLset, []byte(key), Int64ToByteArray(index), []byte(value))
}

func (c *client) lrem(key string, count int64, value string) error {
	return c.sendCommand(cmdLrem, []byte(key), Int64ToByteArray(count), []byte(value))
}

func (c *client) lpop(key string) error {
	return c.sendCommand(cmdLpop, []byte(key))
}

func (c *client) rpop(key string) error {
	return c.sendCommand(cmdRpop, []byte(key))
}

func (c *client) rpopLpush(srckey, dstkey string) error {
	return c.sendCommand(cmdRpoplpush, []byte(srckey), []byte(dstkey))
}

func (c *client) sadd(key string, members ...string) error {
	return c.sendCommand(cmdSadd, StringStringArrayToByteArray(key, members)...)
}

func (c *client) smembers(key string) error {
	return c.sendCommand(cmdSmembers, []byte(key))
}

func (c *client) srem(key string, members ...string) error {
	return c.sendCommand(cmdSrem, StringStringArrayToByteArray(key, members)...)
}

func (c *client) spop(key string) error {
	return c.sendCommand(cmdSpop, []byte(key))
}

func (c *client) spopBatch(key string, count int64) error {
	return c.sendCommand(cmdSpop, []byte(key), Int64ToByteArray(count))
}

func (c *client) smove(srckey, dstkey, member string) error {
	return c.sendCommand(cmdSmove, []byte(srckey), []byte(dstkey), []byte(member))
}

func (c *client) scard(key string) error {
	return c.sendCommand(cmdScard, []byte(key))
}

func (c *client) sismember(key, member string) error {
	return c.sendCommand(cmdSismember, []byte(key), []byte(member))
}

func (c *client) sinter(keys ...string) error {
	return c.sendCommand(cmdSinter, StringArrayToByteArray(keys)...)
}

func (c *client) sinterstore(dstkey string, keys ...string) error {
	return c.sendCommand(cmdSinterstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) sunion(keys ...string) error {
	return c.sendCommand(cmdSunion, StringArrayToByteArray(keys)...)
}

func (c *client) sunionstore(dstkey string, keys ...string) error {
	return c.sendCommand(cmdSunionstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) sdiff(keys ...string) error {
	return c.sendCommand(cmdSdiff, StringArrayToByteArray(keys)...)
}

func (c *client) sdiffstore(dstkey string, keys ...string) error {
	return c.sendCommand(cmdSdiffstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) srandmember(key string) error {
	return c.sendCommand(cmdSrandmember, []byte(key))
}

func (c *client) zadd(key string, score float64, member string, params ...*ZAddParams) error {
	newArr := make([][]byte, 0)
	if len(params) == 0 {
		return c.sendCommand(cmdZadd, []byte(key), Float64ToByteArray(score), []byte(member))
	}
	newArr = append(newArr, Float64ToByteArray(score))
	newArr = append(newArr, []byte(member))
	return c.sendCommand(cmdZadd, params[0].GetByteParams([]byte(key), newArr...)...)
}

func (c *client) zaddByMap(key string, scoreMembers map[string]float64, params ...*ZAddParams) error {
	newArr := make([][]byte, 0)
	if len(params) == 0 {
		newArr = append(newArr, []byte(key))
		for k, v := range scoreMembers {
			newArr = append(newArr, Float64ToByteArray(v))
			newArr = append(newArr, []byte(k))
		}
		return c.sendCommand(cmdZadd, newArr...)
	}
	for k, v := range scoreMembers {
		newArr = append(newArr, Float64ToByteArray(v))
		newArr = append(newArr, []byte(k))
	}
	return c.sendCommand(cmdZadd, params[0].GetByteParams([]byte(key), newArr...)...)
}

func (c *client) zrange(key string, start, end int64) error {
	return c.sendCommand(cmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zrem(key string, members ...string) error {
	return c.sendCommand(cmdZrem, StringStringArrayToByteArray(key, members)...)
}

func (c *client) zincrby(key string, score float64, member string) error {
	return c.sendCommand(cmdZincrby, []byte(key), Float64ToByteArray(score), []byte(member))
}

func (c *client) zrank(key, member string) error {
	return c.sendCommand(cmdZrank, []byte(key), []byte(member))
}

func (c *client) zrevrank(key, member string) error {
	return c.sendCommand(cmdZrevrank, []byte(key), []byte(member))
}

func (c *client) zrevrange(key string, start, end int64) error {
	return c.sendCommand(cmdZrevrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zrangeWithScores(key string, start, end int64) error {
	return c.sendCommand(cmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), keywordWithscores.GetRaw())
}

func (c *client) zrevrangeWithScores(key string, start, end int64) error {
	return c.sendCommand(cmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), keywordWithscores.GetRaw())
}

func (c *client) zcard(key string) error {
	return c.sendCommand(cmdZcard, []byte(key))
}

func (c *client) zscore(key, member string) error {
	return c.sendCommand(cmdZscore, []byte(key), []byte(member))
}

func (c *client) watch(keys ...string) error {
	return c.sendCommand(cmdWatch, StringArrayToByteArray(keys)...)
}

func (c *client) sort(key string, sortingParameters ...SortingParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.params...)
	}
	return c.sendCommand(cmdSort, newArr...)
}

func (c *client) sortMulti(key, dstkey string, sortingParameters ...SortingParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.params...)
	}
	newArr = append(newArr, keywordStore.GetRaw())
	newArr = append(newArr, []byte(dstkey))
	return c.sendCommand(cmdSort, newArr...)
}

func (c *client) blpop(args []string) error {
	return c.sendCommand(cmdBlpop, StringArrayToByteArray(args)...)
}

func (c *client) brpop(args []string) error {
	return c.sendCommand(cmdBrpop, StringArrayToByteArray(args)...)
}

func (c *client) zcount(key, min, max string) error {
	return c.sendCommand(cmdZcount, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByScore(key, min, max string) error {
	return c.sendCommand(cmdZrangebyscore, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByScoreWithScores(key, min, max string) error {
	return c.sendCommand(cmdZrangebyscore, []byte(key), []byte(min), []byte(max), keywordWithscores.GetRaw())
}

func (c *client) zrevrangeByScore(key, max, min string) error {
	return c.sendCommand(cmdZrevrangebyscore, []byte(key), []byte(max), []byte(min))
}

func (c *client) zrevrangeByScoreWithScores(key, max, min string) error {
	return c.sendCommand(cmdZrevrangebyscore, []byte(key), []byte(max), []byte(min), keywordWithscores.GetRaw())
}

func (c *client) zremrangeByRank(key string, start, end int64) error {
	return c.sendCommand(cmdZremrangebyrank, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zremrangeByScore(key, start, end string) error {
	return c.sendCommand(cmdZremrangebyscore, []byte(key), []byte(start), []byte(end))
}

func (c *client) zunionstore(dstkey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdZunionstore, arr...)
}

func (c *client) zunionstoreWithParams(dstkey string, params ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.GetParams()...)
	return c.sendCommand(cmdZunionstore, arr...)
}

func (c *client) zinterstore(dstkey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdZinterstore, arr...)
}

func (c *client) zinterstoreWithParams(dstkey string, params ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.GetParams()...)
	return c.sendCommand(cmdZinterstore, arr...)
}

func (c *client) zlexcount(key, min, max string) error {
	return c.sendCommand(cmdZlexcount, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLex(key, min, max string) error {
	return c.sendCommand(cmdZrangebylex, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLexBatch(key, min, max string, offset, count int) error {
	return c.sendCommand(cmdZrangebylex, []byte(key), []byte(min), []byte(max), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zrevrangeByLex(key, max, min string) error {
	return c.sendCommand(cmdZrevrangebylex, []byte(key), []byte(max), []byte(min))
}

func (c *client) zrevrangeByLexBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(cmdZrangebylex, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zremrangeByLex(key, min, max string) error {
	return c.sendCommand(cmdZremrangebylex, []byte(key), []byte(min), []byte(max))
}

func (c *client) strlen(key string) error {
	return c.sendCommand(cmdStrlen, []byte(key))
}

func (c *client) lpushx(key string, string ...string) error {
	return c.sendCommand(cmdLpushx, StringStringArrayToByteArray(key, string)...)
}

func (c *client) persist(key string) error {
	return c.sendCommand(cmdPersist, []byte(key))
}

func (c *client) rpushx(key string, string ...string) error {
	return c.sendCommand(cmdRpushx, StringStringArrayToByteArray(key, string)...)
}

func (c *client) echo(string string) error {
	return c.sendCommand(cmdEcho, []byte(string))
}

func (c *client) brpoplpush(source, destination string, timeout int) error {
	return c.sendCommand(cmdBrpoplpush, []byte(source), []byte(destination), IntToByteArray(timeout))
}

func (c *client) setbit(key string, offset int64, value string) error {
	return c.sendCommand(cmdSetbit, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *client) getbit(key string, offset int64) error {
	return c.sendCommand(cmdGetbit, []byte(key), Int64ToByteArray(offset))
}

func (c *client) setrange(key string, offset int64, value string) error {
	return c.sendCommand(cmdSetrange, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *client) getrange(key string, startOffset, endOffset int64) error {
	return c.sendCommand(cmdGetrange, []byte(key), Int64ToByteArray(startOffset), Int64ToByteArray(endOffset))
}

func (c *client) publish(channel, message string) error {
	return c.sendCommand(cmdPublish, []byte(channel), []byte(message))
}

func (c *client) unsubscribe(channels ...string) error {
	return c.sendCommand(cmdUnsubscribe, StringArrayToByteArray(channels)...)
}

func (c *client) psubscribe(patterns ...string) error {
	return c.sendCommand(cmdPsubscribe, StringArrayToByteArray(patterns)...)
}

func (c *client) punsubscribe(patterns ...string) error {
	return c.sendCommand(cmdPunsubscribe, StringArrayToByteArray(patterns)...)
}

func (c *client) subscribe(channels ...string) error {
	return c.sendCommand(cmdSubscribe, StringArrayToByteArray(channels)...)
}

func (c *client) pubsub(subcommand string, args ...string) error {
	return c.sendCommand(cmdPubsub, StringStringArrayToByteArray(subcommand, args)...)
}

func (c *client) configSet(parameter, value string) error {
	return c.sendCommand(cmdConfig, keywordSet.GetRaw(), []byte(parameter), []byte(value))
}

func (c *client) configGet(pattern string) error {
	return c.sendCommand(cmdConfig, keywordGet.GetRaw(), []byte(pattern))
}

func (c *client) eval(script string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(script))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.sendCommand(cmdEval, arr...)
}

func (c *client) evalsha(sha1 string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(sha1))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.sendCommand(cmdEvalsha, arr...)
}

func (c *client) scriptExists(sha1 ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, keywordExists.GetRaw())
	arr = append(arr, StringArrayToByteArray(sha1)...)
	return c.sendCommand(cmdScript, arr...)
}

func (c *client) scriptLoad(script string) error {
	return c.sendCommand(cmdScript, keywordLoad.GetRaw(), []byte(script))
}

func (c *client) sentinel(args ...string) error {
	return c.sendCommand(cmdSentinel, StringArrayToByteArray(args)...)
}

func (c *client) dump(key string) error {
	return c.sendCommand(cmdDump, []byte(key))
}

func (c *client) restore(key string, ttl int, serializedValue []byte) error {
	return c.sendCommand(cmdRestore, []byte(key), IntToByteArray(ttl), serializedValue)
}

func (c *client) incrByFloat(key string, increment float64) error {
	return c.sendCommand(cmdIncrbyfloat, []byte(key), Float64ToByteArray(increment))
}

func (c *client) srandmemberBatch(key string, count int) error {
	return c.sendCommand(cmdSrandmember, []byte(key), IntToByteArray(count))
}

func (c *client) clientKill(client string) error {
	return c.sendCommand(cmdClient, keywordKill.GetRaw(), []byte(client))
}

func (c *client) clientGetname() error {
	return c.sendCommand(cmdClient, keywordGetname.GetRaw())
}

func (c *client) clientList() error {
	return c.sendCommand(cmdClient, keywordList.GetRaw())
}

func (c *client) clientSetname(name string) error {
	return c.sendCommand(cmdClient, keywordSetname.GetRaw(), []byte(name))
}

func (c *client) time() error {
	return c.sendCommand(cmdTime)
}

func (c *client) migrate(host string, port int, key string, destinationDb int, timeout int) error {
	return c.sendCommand(cmdMigrate, []byte(host), IntToByteArray(port), []byte(key), IntToByteArray(destinationDb), IntToByteArray(timeout))
}

func (c *client) hincrByFloat(key, field string, increment float64) error {
	return c.sendCommand(cmdHincrbyfloat, []byte(key), []byte(field), Float64ToByteArray(increment))
}

func (c *client) waitReplicas(replicas int, timeout int64) error {
	return c.sendCommand(cmdWait, IntToByteArray(replicas), Int64ToByteArray(timeout))
}

func (c *client) cluster(args ...[]byte) error {
	return c.sendCommand(cmdCluster, args...)
}

func (c *client) asking() error {
	return c.sendCommand(cmdAsking)
}

func (c *client) readonly() error {
	return c.sendCommand(cmdReadonly)
}

func (c *client) geoadd(key string, longitude, latitude float64, member string) error {
	return c.sendCommand(cmdGeoadd, []byte(key), Float64ToByteArray(longitude), Float64ToByteArray(latitude), []byte(member))
}

func (c *client) geoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	for k, v := range memberCoordinateMap {
		arr = append(arr, Float64ToByteArray(v.longitude))
		arr = append(arr, Float64ToByteArray(v.latitude))
		arr = append(arr, []byte(k))
	}
	return c.sendCommand(cmdGeoadd, arr...)
}

func (c *client) geodist(key, member1, member2 string, unit ...GeoUnit) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member1))
	arr = append(arr, []byte(member2))
	for _, u := range unit {
		arr = append(arr, u.GetRaw())
	}
	return c.sendCommand(cmdGeodist, arr...)
}

func (c *client) geohash(key string, members ...string) error {
	return c.sendCommand(cmdGeohash, StringStringArrayToByteArray(key, members)...)
}

func (c *client) geopos(key string, members ...string) error {
	return c.sendCommand(cmdGeopos, StringStringArrayToByteArray(key, members)...)
}

func (c *client) flushDB() error {
	return c.sendCommand(cmdFlushdb)
}

func (c *client) dbSize() error {
	return c.sendCommand(cmdDbsize)
}

func (c *client) flushAll() error {
	return c.sendCommand(cmdFlushall)
}

func (c *client) save() error {
	return c.sendCommand(cmdSave)
}

func (c *client) bgsave() error {
	return c.sendCommand(cmdBgsave)
}

func (c *client) bgrewriteaof() error {
	return c.sendCommand(cmdBgrewriteaof)
}

func (c *client) lastsave() error {
	return c.sendCommand(cmdLastsave)
}

func (c *client) shutdown() error {
	return c.sendCommand(cmdShutdown)
}

func (c *client) slaveof(host string, port int) error {
	return c.sendCommand(cmdSlaveof, []byte(host), IntToByteArray(port))
}

func (c *client) slaveofNoOne() error {
	return c.sendCommand(cmdSlaveof, keywordNo.GetRaw(), keywordOne.GetRaw())
}

func (c *client) getDB() int {
	return c.Db
}

func (c *client) debug(params DebugParams) error {
	return c.sendCommand(cmdDebug, StringArrayToByteArray(params.command)...)
}

func (c *client) configResetStat() error {
	return c.sendCommand(cmdConfig, keywordResetstat.GetRaw())
}

func (c *client) zrangeByScoreBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(cmdZrangebyscore, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zrevrangeByScoreBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(cmdZrevrangebyscore, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) linsert(key string, where ListOption, pivot, value string) error {
	return c.sendCommand(cmdLinsert, []byte(key), where.GetRaw(), []byte(pivot), []byte(value))
}

func (c *client) bitcount(key string) error {
	return c.sendCommand(cmdBitcount, []byte(key))
}

func (c *client) bitcountRange(key string, start, end int64) error {
	return c.sendCommand(cmdBitcount, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) bitpos(key string, value bool, params ...BitPosParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, BoolToByteArray(value))
	for _, p := range params {
		arr = append(arr, p.params...)
	}
	return c.sendCommand(cmdBitpos, arr...)
}

func (c *client) scan(cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(cmdScan, arr...)
}

func (c *client) hscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(cmdHscan, arr...)
}

func (c *client) sscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(cmdSscan, arr...)
}

func (c *client) zscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(cmdZscan, arr...)
}

func (c *client) unwatch() error {
	return c.sendCommand(cmdUnwatch)
}

func (c *client) blpopTimout(timeout int, keys ...string) error {
	arr := make([]string, 0)
	for _, k := range keys {
		arr = append(arr, k)
	}
	arr = append(arr, strconv.Itoa(timeout))
	return c.blpop(arr)
}

func (c *client) brpopTimout(timeout int, keys ...string) error {
	arr := make([]string, 0)
	for _, k := range keys {
		arr = append(arr, k)
	}
	arr = append(arr, strconv.Itoa(timeout))
	return c.brpop(arr)
}

func (c *client) pfadd(key string, elements ...string) error {
	return c.sendCommand(cmdPfadd, StringStringArrayToByteArray(key, elements)...)
}

func (c *client) georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...*GeoRadiusParam) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, Float64ToByteArray(longitude))
	arr = append(arr, Float64ToByteArray(latitude))
	arr = append(arr, Float64ToByteArray(radius))
	arr = append(arr, unit.GetRaw())
	if len(param) == 0 {
		return c.sendCommand(cmdGeoradius, arr...)
	}
	return c.sendCommand(cmdGeoradius, param[0].GetParams(arr)...)
}

func (c *client) georadiusByMember(key, member string, radius float64, unit GeoUnit, param ...*GeoRadiusParam) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member))
	arr = append(arr, Float64ToByteArray(radius))
	arr = append(arr, unit.GetRaw())
	if len(param) == 0 {
		return c.sendCommand(cmdGeoradiusbymember, arr...)
	}
	return c.sendCommand(cmdGeoradiusbymember, param[0].GetParams(arr)...)
}

func (c *client) bitfield(key string, arguments ...string) error {
	return c.sendCommand(cmdBitfield, StringStringArrayToByteArray(key, arguments)...)
}

func (c *client) randomKey() error {
	return c.sendCommand(cmdRandomkey)
}

func (c *client) bitop(op BitOP, destKey string, srcKeys ...string) error {
	kw := BitOpAnd
	switch op.Name {
	case "AND":
		kw = BitOpAnd
	case "OR":
		kw = BitOpOr
	case "XOR":
		kw = BitOpXor
	case "NOT":
		kw = BitOpNot
	}
	arr := make([][]byte, 0)
	arr = append(arr, kw.GetRaw())
	arr = append(arr, []byte(destKey))
	for _, s := range srcKeys {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdBitop, arr...)
}

func (c *client) pfmerge(destkey string, sourcekeys ...string) error {
	return c.sendCommand(cmdPfmerge, StringStringArrayToByteArray(destkey, sourcekeys)...)
}

func (c *client) pfcount(keys ...string) error {
	return c.sendCommand(cmdPfcount, StringArrayToByteArray(keys)...)
}

func (c *client) slowlogReset() error {
	return c.sendCommand(cmdSlowlog, keywordReset.GetRaw())
}

func (c *client) slowlogLen() error {
	return c.sendCommand(cmdSlowlog, keywordLen.GetRaw())
}

func (c *client) slowlogGet(entries ...int64) error {
	arr := make([][]byte, 0)
	arr = append(arr, keywordGet.GetRaw())
	for _, e := range entries {
		arr = append(arr, Int64ToByteArray(e))
	}
	return c.sendCommand(cmdSlowlog, arr...)
}

func (c *client) objectRefcount(str string) error {
	return c.sendCommand(cmdObject, keywordRefcount.GetRaw(), []byte(str))
}

func (c *client) objectEncoding(str string) error {
	return c.sendCommand(cmdObject, keywordEncoding.GetRaw(), []byte(str))
}

func (c *client) objectIdletime(str string) error {
	return c.sendCommand(cmdObject, keywordIdletime.GetRaw(), []byte(str))
}

func (c *client) clusterNodes() error {
	return c.sendCommand(cmdCluster, []byte(clusterNodes))
}

func (c *client) clusterMeet(ip string, port int) error {
	return c.sendCommand(cmdCluster, []byte(clusterMeet), []byte(ip), IntToByteArray(port))
}

func (c *client) clusterAddSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(clusterAddslots))
	for _, s := range slots {
		arr = append(arr, IntToByteArray(s))
	}
	return c.sendCommand(cmdCluster, arr...)
}

func (c *client) clusterDelSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(clusterDelslots))
	for _, s := range slots {
		arr = append(arr, IntToByteArray(s))
	}
	return c.sendCommand(cmdCluster, arr...)
}

func (c *client) clusterInfo() error {
	return c.sendCommand(cmdCluster, []byte(clusterInfo))
}

func (c *client) clusterGetKeysInSlot(slot int, count int) error {
	return c.sendCommand(cmdCluster, []byte(clusterGetkeysinslot), IntToByteArray(slot), IntToByteArray(count))
}

func (c *client) clusterSetSlotNode(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetslotNode), IntToByteArray(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotMigrating(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetslotMigrating), IntToByteArray(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotImporting(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetslotImporting), IntToByteArray(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotStable(slot int) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetslotStable), IntToByteArray(slot))
}

func (c *client) clusterForget(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterForget), []byte(nodeID))
}

func (c *client) clusterFlushSlots() error {
	return c.sendCommand(cmdCluster, []byte(clusterFlushslot))
}

func (c *client) clusterKeySlot(key string) error {
	return c.sendCommand(cmdCluster, []byte(clusterKeyslot), []byte(key))
}

func (c *client) clusterCountKeysInSlot(slot int) error {
	return c.sendCommand(cmdCluster, []byte(clusterCountkeyinslot), IntToByteArray(slot))
}

func (c *client) clusterSaveConfig() error {
	return c.sendCommand(cmdCluster, []byte(clusterSaveconfig))
}

func (c *client) clusterReplicate(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterReplicate), []byte(nodeID))
}

func (c *client) clusterSlaves(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSlaves), []byte(nodeID))
}

func (c *client) clusterFailover() error {
	return c.sendCommand(cmdCluster, []byte(clusterFailover))
}

func (c *client) clusterSlots() error {
	return c.sendCommand(cmdCluster, []byte(clusterSlots))
}

func (c *client) clusterReset(resetType Reset) error {
	return c.sendCommand(cmdCluster, []byte(clusterReset), resetType.GetRaw())
}

func (c *client) sentinelMasters() error {
	return c.sendCommand(cmdSentinel, []byte(sentinelMasters))
}

func (c *client) sentinelGetMasterAddrByName(masterName string) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelGetMasterAddrByName), []byte(masterName))
}

func (c *client) sentinelReset(pattern string) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelReset), []byte(pattern))
}

func (c *client) sentinelSlaves(masterName string) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelSlaves), []byte(masterName))
}

func (c *client) sentinelFailover(masterName string) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelFailover), []byte(masterName))
}

func (c *client) sentinelMonitor(masterName, ip string, port, quorum int) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelMonitor), []byte(masterName), []byte(ip), IntToByteArray(port), IntToByteArray(quorum))
}

func (c *client) sentinelRemove(masterName string) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelRemove), []byte(masterName))
}

func (c *client) sentinelSet(masterName string, parameterMap map[string]string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(sentinelSet))
	arr = append(arr, []byte(masterName))
	for k, v := range parameterMap {
		arr = append(arr, []byte(k))
		arr = append(arr, []byte(v))
	}
	return c.sendCommandByStr(sentinelFailover, arr...)
}

func (c *client) pubsubChannels(pattern string) error {
	return c.sendCommand(cmdPubsub, []byte(pubsubChannels), []byte(pattern))
}

func (c *client) multi() error {
	err := c.sendCommand(cmdMulti)
	if err != nil {
		return err
	}
	c.isInMulti = true
	return nil
}

func (c *client) discard() error {
	err := c.sendCommand(cmdDiscard)
	if err != nil {
		return err
	}
	c.isInMulti = false
	c.isInWatch = false
	return nil
}

func (c *client) exec() error {
	err := c.sendCommand(cmdExec)
	if err != nil {
		return err
	}
	c.isInMulti = false
	c.isInWatch = false
	return nil
}

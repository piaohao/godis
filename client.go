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
	return c.sendCommand(CmdPing)
}

//Quit
func (c *client) quit() error {
	return c.sendCommand(CmdQuit)
}

//Info
func (c *client) info(section ...string) error {
	return c.sendCommand(CmdInfo, StringArrayToByteArray(section)...)
}

//Auth
func (c *client) auth(password string) error {
	c.Password = password
	return c.sendCommand(CmdAuth, []byte(password))
}

//Select
func (c *client) selectDb(index int) error {
	return c.sendCommand(CmdSelect, IntToByteArray(index))
}

func (c *client) set(key, value string) error {
	return c.sendCommand(CmdSet, []byte(key), []byte(value))
}

func (c *client) setWithParamsAndTime(key, value, nxxx, expx string, time int64) error {
	return c.sendCommand(CmdSet, []byte(key), []byte(value), []byte(nxxx), []byte(expx), Int64ToByteArray(time))
}

func (c *client) setWithParams(key, value, nxxx string) error {
	return c.sendCommand(CmdSet, []byte(key), []byte(value), []byte(nxxx))
}

func (c *client) get(key string) error {
	return c.sendCommand(CmdGet, []byte(key))
}

func (c *client) del(keys ...string) error {
	return c.sendCommand(CmdDel, StringArrayToByteArray(keys)...)
}

func (c *client) exists(keys ...string) error {
	return c.sendCommand(CmdExists, StringArrayToByteArray(keys)...)
}

func (c *client) typeKey(key string) error {
	return c.sendCommand(CmdType, []byte(key))
}

func (c *client) keys(pattern string) error {
	return c.sendCommand(CmdKeys, []byte(pattern))
}

func (c *client) rename(oldKey, newKey string) error {
	return c.sendCommand(CmdRename, []byte(oldKey), []byte(newKey))
}

func (c *client) renamenx(oldKey, newKey string) error {
	return c.sendCommand(CmdRenamex, []byte(oldKey), []byte(newKey))
}

func (c *client) expire(key string, seconds int) error {
	return c.sendCommand(CmdExpire, []byte(key), IntToByteArray(seconds))
}

func (c *client) expireAt(key string, unixTime int64) error {
	return c.sendCommand(CmdExpireat, []byte(key), Int64ToByteArray(unixTime))
}

func (c *client) pexpire(key string, milliseconds int64) error {
	return c.sendCommand(CmdPexpire, []byte(key), Int64ToByteArray(milliseconds))
}

func (c *client) pexpireAt(key string, unixTime int64) error {
	return c.sendCommand(CmdPexpireat, []byte(key), Int64ToByteArray(unixTime))
}

func (c *client) ttl(key string) error {
	return c.sendCommand(CmdTtl, []byte(key))
}

func (c *client) pttl(key string) error {
	return c.sendCommand(CmdPttl, []byte(key))
}

func (c *client) move(key string, dbIndex int) error {
	return c.sendCommand(CmdMove, []byte(key), IntToByteArray(dbIndex))
}

func (c *client) getSet(key, value string) error {
	return c.sendCommand(CmdGetset, []byte(key), []byte(value))
}

func (c *client) mget(keys ...string) error {
	return c.sendCommand(CmdMget, StringArrayToByteArray(keys)...)
}

func (c *client) setnx(key, value string) error {
	return c.sendCommand(CmdSetnx, []byte(key), []byte(value))
}

func (c *client) setex(key string, seconds int, value string) error {
	return c.sendCommand(CmdSetex, []byte(key), IntToByteArray(seconds), []byte(value))
}

func (c *client) psetex(key string, milliseconds int64, value string) error {
	return c.sendCommand(CmdSetex, []byte(key), Int64ToByteArray(milliseconds), []byte(value))
}

func (c *client) mset(keysvalues ...string) error {
	return c.sendCommand(CmdMset, StringArrayToByteArray(keysvalues)...)
}

func (c *client) msetnx(keysvalues ...string) error {
	return c.sendCommand(CmdMsetnx, StringArrayToByteArray(keysvalues)...)
}

func (c *client) decrBy(key string, decrement int64) error {
	return c.sendCommand(CmdDecrby, []byte(key), Int64ToByteArray(decrement))
}

func (c *client) decr(key string) error {
	return c.sendCommand(CmdDecr, []byte(key))
}

func (c *client) incrBy(key string, increment int64) error {
	return c.sendCommand(CmdIncrby, []byte(key), Int64ToByteArray(increment))
}

func (c *client) incr(key string) error {
	return c.sendCommand(CmdIncr, []byte(key))
}

func (c *client) append(key, value string) error {
	return c.sendCommand(CmdAppend, []byte(key), []byte(value))
}

func (c *client) substr(key string, start, end int) error {
	return c.sendCommand(CmdSubstr, []byte(key), IntToByteArray(start), IntToByteArray(end))
}

func (c *client) hset(key, field, value string) error {
	return c.sendCommand(CmdHset, []byte(key), []byte(field), []byte(value))
}

func (c *client) hget(key, field string) error {
	return c.sendCommand(CmdHget, []byte(key), []byte(field))
}

func (c *client) hsetnx(key, field, value string) error {
	return c.sendCommand(CmdHsetnx, []byte(key), []byte(field), []byte(value))
}

func (c *client) hmset(key string, hash map[string]string) error {
	params := make([][]byte, 0)
	params = append(params, []byte(key))
	for k, v := range hash {
		params = append(params, []byte(k))
		params = append(params, []byte(v))
	}
	return c.sendCommand(CmdHmset, params...)
}

func (c *client) hmget(key string, fields ...string) error {
	return c.sendCommand(CmdHmget, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) hincrBy(key, field string, increment int64) error {
	return c.sendCommand(CmdHincrby, []byte(key), []byte(field), Int64ToByteArray(increment))
}

func (c *client) hexists(key, field string) error {
	return c.sendCommand(CmdHexists, []byte(key), []byte(field))
}

func (c *client) hdel(key string, fields ...string) error {
	return c.sendCommand(CmdHdel, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) hlen(key string) error {
	return c.sendCommand(CmdHlen, []byte(key))
}

func (c *client) hkeys(key string) error {
	return c.sendCommand(CmdHkeys, []byte(key))
}

func (c *client) hvals(key string) error {
	return c.sendCommand(CmdHvals, []byte(key))
}

func (c *client) hgetAll(key string) error {
	return c.sendCommand(CmdHgetall, []byte(key))
}

func (c *client) rpush(key string, fields ...string) error {
	return c.sendCommand(CmdRpush, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) lpush(key string, fields ...string) error {
	return c.sendCommand(CmdRpush, StringStringArrayToByteArray(key, fields)...)
}

func (c *client) llen(key string) error {
	return c.sendCommand(CmdLlen, []byte(key))
}

func (c *client) lrange(key string, start, end int64) error {
	return c.sendCommand(CmdLrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) ltrim(key string, start, end int64) error {
	return c.sendCommand(CmdLtrim, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) lindex(key string, index int64) error {
	return c.sendCommand(CmdLindex, []byte(key), Int64ToByteArray(index))
}

func (c *client) lset(key string, index int64, value string) error {
	return c.sendCommand(CmdLset, []byte(key), Int64ToByteArray(index), []byte(value))
}

func (c *client) lrem(key string, count int64, value string) error {
	return c.sendCommand(CmdLrem, []byte(key), Int64ToByteArray(count), []byte(value))
}

func (c *client) lpop(key string) error {
	return c.sendCommand(CmdLpop, []byte(key))
}

func (c *client) rpop(key string) error {
	return c.sendCommand(CmdRpop, []byte(key))
}

func (c *client) rpopLpush(srckey, dstkey string) error {
	return c.sendCommand(CmdRpoplpush, []byte(srckey), []byte(dstkey))
}

func (c *client) sadd(key string, members ...string) error {
	return c.sendCommand(CmdSadd, StringStringArrayToByteArray(key, members)...)
}

func (c *client) smembers(key string) error {
	return c.sendCommand(CmdSmembers, []byte(key))
}

func (c *client) srem(key string, members ...string) error {
	return c.sendCommand(CmdSrem, StringStringArrayToByteArray(key, members)...)
}

func (c *client) spop(key string) error {
	return c.sendCommand(CmdSpop, []byte(key))
}

func (c *client) spopBatch(key string, count int64) error {
	return c.sendCommand(CmdSpop, []byte(key), Int64ToByteArray(count))
}

func (c *client) smove(srckey, dstkey, member string) error {
	return c.sendCommand(CmdSmove, []byte(srckey), []byte(dstkey), []byte(member))
}

func (c *client) scard(key string) error {
	return c.sendCommand(CmdScard, []byte(key))
}

func (c *client) sismember(key, member string) error {
	return c.sendCommand(CmdSismember, []byte(key), []byte(member))
}

func (c *client) sinter(keys ...string) error {
	return c.sendCommand(CmdSinter, StringArrayToByteArray(keys)...)
}

func (c *client) sinterstore(dstkey string, keys ...string) error {
	return c.sendCommand(CmdSinterstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) sunion(keys ...string) error {
	return c.sendCommand(CmdSunion, StringArrayToByteArray(keys)...)
}

func (c *client) sunionstore(dstkey string, keys ...string) error {
	return c.sendCommand(CmdSunionstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) sdiff(keys ...string) error {
	return c.sendCommand(CmdSdiff, StringArrayToByteArray(keys)...)
}

func (c *client) sdiffstore(dstkey string, keys ...string) error {
	return c.sendCommand(CmdSdiffstore, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *client) srandmember(key string) error {
	return c.sendCommand(CmdSrandmember, []byte(key))
}

func (c *client) zadd(key string, score float64, member string) error {
	return c.sendCommand(CmdZadd, []byte(key), Float64ToByteArray(score), []byte(member))
}

func (c *client) zaddByMap(key string, scoreMembers map[string]float64, params ...*ZAddParams) error {
	newArr := make([][]byte, 0)
	for k, v := range scoreMembers {
		newArr = append(newArr, Float64ToByteArray(v))
		newArr = append(newArr, []byte(k))
	}
	return c.sendCommand(CmdZadd, params[0].GetByteParams([]byte(key), newArr...)...)
}

func (c *client) zrange(key string, start, end int64) error {
	return c.sendCommand(CmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zrem(key string, members ...string) error {
	return c.sendCommand(CmdZrem, StringStringArrayToByteArray(key, members)...)
}

func (c *client) zincrby(key string, score float64, member string) error {
	return c.sendCommand(CmdZincrby, []byte(key), Float64ToByteArray(score), []byte(member))
}

func (c *client) zrank(key, member string) error {
	return c.sendCommand(CmdZrank, []byte(key), []byte(member))
}

func (c *client) zrevrank(key, member string) error {
	return c.sendCommand(CmdZrevrank, []byte(key), []byte(member))
}

func (c *client) zrevrange(key string, start, end int64) error {
	return c.sendCommand(CmdZrevrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zrangeWithScores(key string, start, end int64) error {
	return c.sendCommand(CmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), KeywordWithscores.GetRaw())
}

func (c *client) zrevrangeWithScores(key string, start, end int64) error {
	return c.sendCommand(CmdZrange, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), KeywordWithscores.GetRaw())
}

func (c *client) zcard(key string) error {
	return c.sendCommand(CmdZcard, []byte(key))
}

func (c *client) zscore(key, member string) error {
	return c.sendCommand(CmdZscore, []byte(key), []byte(member))
}

func (c *client) watch(keys ...string) error {
	return c.sendCommand(CmdWatch, StringArrayToByteArray(keys)...)
}

func (c *client) sort(key string, sortingParameters ...SortingParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.params...)
	}
	return c.sendCommand(CmdSort, newArr...)
}

func (c *client) sortMulti(key, dstkey string, sortingParameters ...SortingParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.params...)
	}
	newArr = append(newArr, []byte(dstkey))
	return c.sendCommand(CmdSort, newArr...)
}

func (c *client) blpop(args []string) error {
	return c.sendCommand(CmdBlpop, StringArrayToByteArray(args)...)
}

func (c *client) brpop(args []string) error {
	return c.sendCommand(CmdBrpop, StringArrayToByteArray(args)...)
}

func (c *client) zcount(key, min, max string) error {
	return c.sendCommand(CmdZcount, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByScore(key, min, max string) error {
	return c.sendCommand(CmdZrangebyscore, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByScoreWithScores(key, min, max string) error {
	return c.sendCommand(CmdZrangebyscore, []byte(key), []byte(min), []byte(max), KeywordWithscores.GetRaw())
}

func (c *client) zrevrangeByScore(key, max, min string) error {
	return c.sendCommand(CmdZrevrangebyscore, []byte(key), []byte(max), []byte(min))
}

func (c *client) zrevrangeByScoreWithScores(key, max, min string) error {
	return c.sendCommand(CmdZrevrangebyscore, []byte(key), []byte(max), []byte(min), KeywordWithscores.GetRaw())
}

func (c *client) zremrangeByRank(key string, start, end int64) error {
	return c.sendCommand(CmdZremrangebyrank, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) zremrangeByScore(key, start, end string) error {
	return c.sendCommand(CmdZremrangebyscore, []byte(key), []byte(start), []byte(end))
}

func (c *client) zunionstore(dstkey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(CmdZunionstore, arr...)
}

func (c *client) zunionstoreWithParams(dstkey string, params ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.GetParams()...)
	return c.sendCommand(CmdZunionstore, arr...)
}

func (c *client) zinterstore(dstkey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(CmdZinterstore, arr...)
}

func (c *client) zinterstoreWithParams(dstkey string, params ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(dstkey))
	arr = append(arr, IntToByteArray(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.GetParams()...)
	return c.sendCommand(CmdZinterstore, arr...)
}

func (c *client) zlexcount(key, min, max string) error {
	return c.sendCommand(CmdZlexcount, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLex(key, min, max string) error {
	return c.sendCommand(CmdZrangebylex, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLexBatch(key, min, max string, offset, count int) error {
	return c.sendCommand(CmdZrangebylex, []byte(key), []byte(min), []byte(max), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zrevrangeByLex(key, max, min string) error {
	return c.sendCommand(CmdZrevrangebylex, []byte(key), []byte(max), []byte(min))
}

func (c *client) zrevrangeByLexBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(CmdZrangebylex, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zremrangeByLex(key, min, max string) error {
	return c.sendCommand(CmdZremrangebylex, []byte(key), []byte(min), []byte(max))
}

func (c *client) strlen(key string) error {
	return c.sendCommand(CmdStrlen, []byte(key))
}

func (c *client) lpushx(key string, string ...string) error {
	return c.sendCommand(CmdLpushx, StringStringArrayToByteArray(key, string)...)
}

func (c *client) persist(key string) error {
	return c.sendCommand(CmdPersist, []byte(key))
}

func (c *client) rpushx(key string, string ...string) error {
	return c.sendCommand(CmdRpushx, StringStringArrayToByteArray(key, string)...)
}

func (c *client) echo(string string) error {
	return c.sendCommand(CmdEcho, []byte(string))
}

func (c *client) brpoplpush(source, destination string, timeout int) error {
	return c.sendCommand(CmdBrpoplpush, []byte(source), []byte(destination), IntToByteArray(timeout))
}

func (c *client) setbit(key string, offset int64, value string) error {
	return c.sendCommand(CmdSetbit, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *client) getbit(key string, offset int64) error {
	return c.sendCommand(CmdGetbit, []byte(key), Int64ToByteArray(offset))
}

func (c *client) setrange(key string, offset int64, value string) error {
	return c.sendCommand(CmdSetrange, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *client) getrange(key string, startOffset, endOffset int64) error {
	return c.sendCommand(CmdGetrange, []byte(key), Int64ToByteArray(startOffset), Int64ToByteArray(endOffset))
}

func (c *client) publish(channel, message string) error {
	return c.sendCommand(CmdPublish, []byte(channel), []byte(message))
}

func (c *client) unsubscribe(channels ...string) error {
	return c.sendCommand(CmdUnsubscribe, StringArrayToByteArray(channels)...)
}

func (c *client) psubscribe(patterns ...string) error {
	return c.sendCommand(CmdPsubscribe, StringArrayToByteArray(patterns)...)
}

func (c *client) punsubscribe(patterns ...string) error {
	return c.sendCommand(CmdPunsubscribe, StringArrayToByteArray(patterns)...)
}

func (c *client) subscribe(channels ...string) error {
	return c.sendCommand(CmdSubscribe, StringArrayToByteArray(channels)...)
}

func (c *client) pubsub(subcommand string, args ...string) error {
	return c.sendCommand(CmdPubsub, StringStringArrayToByteArray(subcommand, args)...)
}

func (c *client) configSet(parameter, value string) error {
	return c.sendCommand(CmdConfig, KeywordSet.GetRaw(), []byte(parameter), []byte(value))
}

func (c *client) configGet(pattern string) error {
	return c.sendCommand(CmdConfig, KeywordGet.GetRaw(), []byte(pattern))
}

func (c *client) eval(script string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(script))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.sendCommand(CmdEval, arr...)
}

func (c *client) evalsha(sha1 string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(sha1))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.sendCommand(CmdEvalsha, arr...)
}

func (c *client) scriptExists(sha1 ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, KeywordExists.GetRaw())
	arr = append(arr, StringArrayToByteArray(sha1)...)
	return c.sendCommand(CmdScript, arr...)
}

func (c *client) scriptLoad(script string) error {
	return c.sendCommand(CmdScript, KeywordLoad.GetRaw(), []byte(script))
}

func (c *client) sentinel(args ...string) error {
	return c.sendCommand(CmdSentinel, StringArrayToByteArray(args)...)
}

func (c *client) dump(key string) error {
	return c.sendCommand(CmdDump, []byte(key))
}

func (c *client) restore(key string, ttl int, serializedValue []byte) error {
	return c.sendCommand(CmdRestore, []byte(key), IntToByteArray(ttl), serializedValue)
}

func (c *client) incrByFloat(key string, increment float64) error {
	return c.sendCommand(CmdIncrbyfloat, []byte(key), Float64ToByteArray(increment))
}

func (c *client) srandmemberBatch(key string, count int) error {
	return c.sendCommand(CmdSrandmember, []byte(key), IntToByteArray(count))
}

func (c *client) clientKill(client string) error {
	return c.sendCommand(CmdClient, KeywordKill.GetRaw(), []byte(client))
}

func (c *client) clientGetname() error {
	return c.sendCommand(CmdClient, KeywordGetname.GetRaw())
}

func (c *client) clientList() error {
	return c.sendCommand(CmdClient, KeywordList.GetRaw())
}

func (c *client) clientSetname(name string) error {
	return c.sendCommand(CmdClient, KeywordSetname.GetRaw(), []byte(name))
}

func (c *client) time() error {
	return c.sendCommand(CmdTime)
}

func (c *client) migrate(host string, port int, key string, destinationDb int, timeout int) error {
	return c.sendCommand(CmdMigrate, []byte(host), IntToByteArray(port), []byte(key), IntToByteArray(destinationDb), IntToByteArray(timeout))
}

func (c *client) hincrByFloat(key, field string, increment float64) error {
	return c.sendCommand(CmdHincrbyfloat, []byte(key), []byte(field), Float64ToByteArray(increment))
}

func (c *client) waitReplicas(replicas int, timeout int64) error {
	return c.sendCommand(CmdWait, IntToByteArray(replicas), Int64ToByteArray(timeout))
}

func (c *client) cluster(args ...[]byte) error {
	return c.sendCommand(CmdCluster, args...)
}

func (c *client) asking() error {
	return c.sendCommand(CmdAsking)
}

func (c *client) readonly() error {
	return c.sendCommand(CmdReadonly)
}

func (c *client) geoadd(key string, longitude, latitude float64, member string) error {
	return c.sendCommand(CmdGeoadd, []byte(key), Float64ToByteArray(longitude), Float64ToByteArray(latitude), []byte(member))
}

func (c *client) geoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	for k, v := range memberCoordinateMap {
		arr = append(arr, Float64ToByteArray(v.longitude))
		arr = append(arr, Float64ToByteArray(v.latitude))
		arr = append(arr, []byte(k))
	}
	return c.sendCommand(CmdGeoadd, arr...)
}

func (c *client) geodist(key, member1, member2 string, unit ...GeoUnit) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member1))
	arr = append(arr, []byte(member2))
	for _, u := range unit {
		arr = append(arr, u.GetRaw())
	}
	return c.sendCommand(CmdGeodist, arr...)
}

func (c *client) geohash(key string, members ...string) error {
	return c.sendCommand(CmdGeohash, StringStringArrayToByteArray(key, members)...)
}

func (c *client) geopos(key string, members ...string) error {
	return c.sendCommand(CmdGeopos, StringStringArrayToByteArray(key, members)...)
}

func (c *client) flushDB() error {
	return c.sendCommand(CmdFlushdb)
}

func (c *client) dbSize() error {
	return c.sendCommand(CmdDbsize)
}

func (c *client) flushAll() error {
	return c.sendCommand(CmdFlushall)
}

func (c *client) save() error {
	return c.sendCommand(CmdSave)
}

func (c *client) bgsave() error {
	return c.sendCommand(CmdBgsave)
}

func (c *client) bgrewriteaof() error {
	return c.sendCommand(CmdBgrewriteaof)
}

func (c *client) lastsave() error {
	return c.sendCommand(CmdLastsave)
}

func (c *client) shutdown() error {
	return c.sendCommand(CmdShutdown)
}

func (c *client) slaveof(host string, port int) error {
	return c.sendCommand(CmdSlaveof, []byte(host), IntToByteArray(port))
}

func (c *client) slaveofNoOne() error {
	return c.sendCommand(CmdSlaveof, KeywordNo.GetRaw(), KeywordOne.GetRaw())
}

func (c *client) getDB() int {
	return c.Db
}

func (c *client) debug(params DebugParams) error {
	return c.sendCommand(CmdDebug, StringArrayToByteArray(params.command)...)
}

func (c *client) configResetStat() error {
	return c.sendCommand(CmdConfig, KeywordResetstat.GetRaw())
}

func (c *client) zrangeByScoreBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(CmdZrangebyscore, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) zrevrangeByScoreBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(CmdZrevrangebyscore, []byte(key), []byte(max), []byte(min), IntToByteArray(offset), IntToByteArray(count))
}

func (c *client) linsert(key string, where ListOption, pivot, value string) error {
	return c.sendCommand(CmdLinsert, []byte(key), where.GetRaw(), []byte(pivot), []byte(value))
}

func (c *client) bitcount(key string) error {
	return c.sendCommand(CmdBitcount, []byte(key))
}

func (c *client) bitcountRange(key string, start, end int64) error {
	return c.sendCommand(CmdBitcount, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *client) bitpos(key string, value bool, params ...BitPosParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, BoolToByteArray(value))
	for _, p := range params {
		arr = append(arr, p.params...)
	}
	return c.sendCommand(CmdBitpos, arr...)
}

func (c *client) scan(cursor string, params ...ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(CmdScan, arr...)
}

func (c *client) hscan(key, cursor string, params ...ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(CmdHscan, arr...)
}

func (c *client) sscan(key, cursor string, params ...ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(CmdHscan, arr...)
}

func (c *client) zscan(key, cursor string, params ...ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.GetParams()...)
	}
	return c.sendCommand(CmdHscan, arr...)
}

func (c *client) unwatch() error {
	return c.sendCommand(CmdUnwatch)
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
	return c.sendCommand(CmdPfadd, StringStringArrayToByteArray(key, elements)...)
}

func (c *client) georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...*GeoRadiusParam) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, Float64ToByteArray(longitude))
	arr = append(arr, Float64ToByteArray(latitude))
	arr = append(arr, Float64ToByteArray(radius))
	arr = append(arr, unit.GetRaw())
	return c.sendCommand(CmdGeoradius, param[0].GetParams(arr)...)
}

func (c *client) georadiusByMember(key, member string, radius float64, unit GeoUnit, param ...*GeoRadiusParam) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member))
	arr = append(arr, Float64ToByteArray(radius))
	arr = append(arr, unit.GetRaw())
	return c.sendCommand(CmdGeoradiusbymember, param[0].GetParams(arr)...)
}

func (c *client) bitfield(key string, arguments ...string) error {
	return c.sendCommand(CmdBitfield, StringStringArrayToByteArray(key, arguments)...)
}

func (c *client) randomKey() error {
	return c.sendCommand(CmdRandomkey)
}

func (c *client) bitop(op BitOP, destKey string, srcKeys ...string) error {
	kw := BitopAnd
	switch op.Name {
	case "AND":
		kw = BitopAnd
	case "OR":
		kw = BitopOr
	case "XOR":
		kw = BitopXor
	case "NOT":
		kw = BitopNot
	}
	arr := make([][]byte, 0)
	arr = append(arr, kw.GetRaw())
	arr = append(arr, []byte(destKey))
	for _, s := range srcKeys {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(CmdBitop, arr...)
}

func (c *client) pfmerge(destkey string, sourcekeys ...string) error {
	return c.sendCommand(CmdPfmerge, StringStringArrayToByteArray(destkey, sourcekeys)...)
}

func (c *client) pfcount(keys ...string) error {
	return c.sendCommand(CmdPfcount, StringArrayToByteArray(keys)...)
}

func (c *client) slowlogReset() error {
	return c.sendCommand(CmdSlowlog, KeywordReset.GetRaw())
}

func (c *client) slowlogLen() error {
	return c.sendCommand(CmdSlowlog, KeywordLen.GetRaw())
}

func (c *client) slowlogGet(entries ...int64) error {
	arr := make([][]byte, 0)
	arr = append(arr, KeywordGet.GetRaw())
	for _, e := range entries {
		arr = append(arr, Int64ToByteArray(e))
	}
	return c.sendCommand(CmdSlowlog, arr...)
}

func (c *client) objectRefcount(str string) error {
	return c.sendCommand(CmdObject, KeywordRefcount.GetRaw(), []byte(str))
}

func (c *client) objectEncoding(str string) error {
	return c.sendCommand(CmdObject, KeywordEncoding.GetRaw(), []byte(str))
}

func (c *client) objectIdletime(str string) error {
	return c.sendCommand(CmdObject, KeywordIdletime.GetRaw(), []byte(str))
}

func (c *client) clusterNodes() error {
	return c.sendCommand(CmdCluster, []byte(ClusterNodes))
}

func (c *client) clusterMeet(ip string, port int) error {
	return c.sendCommand(CmdCluster, []byte(ClusterMeet), []byte(ip), IntToByteArray(port))
}

func (c *client) clusterAddSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(ClusterAddslots))
	for _, s := range slots {
		arr = append(arr, IntToByteArray(s))
	}
	return c.sendCommand(CmdCluster, arr...)
}

func (c *client) clusterDelSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(ClusterDelslots))
	for _, s := range slots {
		arr = append(arr, IntToByteArray(s))
	}
	return c.sendCommand(CmdCluster, arr...)
}

func (c *client) clusterInfo() error {
	return c.sendCommand(CmdCluster, []byte(ClusterInfo))
}

func (c *client) clusterGetKeysInSlot(slot int, count int) error {
	return c.sendCommand(CmdCluster, []byte(ClusterGetkeysinslot), IntToByteArray(slot), IntToByteArray(count))
}

func (c *client) clusterSetSlotNode(slot int, nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterSetslotNode), IntToByteArray(slot), []byte(nodeId))
}

func (c *client) clusterSetSlotMigrating(slot int, nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterSetslotMigrating), IntToByteArray(slot), []byte(nodeId))
}

func (c *client) clusterSetSlotImporting(slot int, nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterSetslotImporting), IntToByteArray(slot), []byte(nodeId))
}

func (c *client) clusterSetSlotStable(slot int) error {
	return c.sendCommand(CmdCluster, []byte(ClusterSetslotStable), IntToByteArray(slot))
}

func (c *client) clusterForget(nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterForget), []byte(nodeId))
}

func (c *client) clusterFlushSlots() error {
	return c.sendCommand(CmdCluster, []byte(ClusterFlushslot))
}

func (c *client) clusterKeySlot(key string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterKeyslot), []byte(key))
}

func (c *client) clusterCountKeysInSlot(slot int) error {
	return c.sendCommand(CmdCluster, []byte(ClusterCountkeyinslot), IntToByteArray(slot))
}

func (c *client) clusterSaveConfig() error {
	return c.sendCommand(CmdCluster, []byte(ClusterSaveconfig))
}

func (c *client) clusterReplicate(nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterReplicate), []byte(nodeId))
}

func (c *client) clusterSlaves(nodeId string) error {
	return c.sendCommand(CmdCluster, []byte(ClusterSlaves), []byte(nodeId))
}

func (c *client) clusterFailover() error {
	return c.sendCommand(CmdCluster, []byte(ClusterFailover))
}

func (c *client) clusterSlots() error {
	return c.sendCommand(CmdCluster, []byte(ClusterSlots))
}

func (c *client) clusterReset(resetType Reset) error {
	return c.sendCommand(CmdCluster, []byte(ClusterReset), resetType.GetRaw())
}

func (c *client) sentinelMasters() error {
	return c.sendCommand(CmdSentinel, []byte(SentinelMasters))
}

func (c *client) sentinelGetMasterAddrByName(masterName string) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelGetMasterAddrByName), []byte(masterName))
}

func (c *client) sentinelReset(pattern string) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelReset), []byte(pattern))
}

func (c *client) sentinelSlaves(masterName string) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelSlaves), []byte(masterName))
}

func (c *client) sentinelFailover(masterName string) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelFailover), []byte(masterName))
}

func (c *client) sentinelMonitor(masterName, ip string, port, quorum int) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelMonitor), []byte(masterName), []byte(ip), IntToByteArray(port), IntToByteArray(quorum))
}

func (c *client) sentinelRemove(masterName string) error {
	return c.sendCommand(CmdSentinel, []byte(SentinelRemove), []byte(masterName))
}

func (c *client) sentinelSet(masterName string, parameterMap map[string]string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(SentinelSet))
	arr = append(arr, []byte(masterName))
	for k, v := range parameterMap {
		arr = append(arr, []byte(k))
		arr = append(arr, []byte(v))
	}
	return c.sendCommandByStr(SentinelFailover, arr...)
}

func (c *client) pubsubChannels(pattern string) error {
	return c.sendCommand(CmdPubsub, []byte(PubsubChannels), []byte(pattern))
}

func (c *client) multi() error {
	err := c.sendCommand(CmdMulti)
	if err != nil {
		return err
	}
	c.isInMulti = true
	return nil
}

func (c *client) discard() error {
	err := c.sendCommand(CmdDiscard)
	if err != nil {
		return err
	}
	c.isInMulti = false
	c.isInWatch = false
	return nil
}

func (c *client) exec() error {
	err := c.sendCommand(CmdExec)
	if err != nil {
		return err
	}
	c.isInMulti = false
	c.isInWatch = false
	return nil
}

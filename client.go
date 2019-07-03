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
	return c.sendCommand(cmdInfo, StrArrToByteArrArr(section)...)
}

//Auth
func (c *client) auth(password string) error {
	c.Password = password
	return c.sendCommand(cmdAuth, []byte(password))
}

//Select
func (c *client) selectDb(index int) error {
	return c.sendCommand(cmdSelect, IntToByteArr(index))
}

func (c *client) set(key, value string) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value))
}

func (c *client) setWithParamsAndTime(key, value, nxxx, expx string, time int64) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value), []byte(nxxx), []byte(expx), Int64ToByteArr(time))
}

func (c *client) setWithParams(key, value, nxxx string) error {
	return c.sendCommand(cmdSet, []byte(key), []byte(value), []byte(nxxx))
}

func (c *client) get(key string) error {
	return c.sendCommand(cmdGet, []byte(key))
}

func (c *client) del(keys ...string) error {
	return c.sendCommand(cmdDel, StrArrToByteArrArr(keys)...)
}

func (c *client) exists(keys ...string) error {
	return c.sendCommand(cmdExists, StrArrToByteArrArr(keys)...)
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
	return c.sendCommand(cmdRenameNx, []byte(oldKey), []byte(newKey))
}

func (c *client) expire(key string, seconds int) error {
	return c.sendCommand(cmdExpire, []byte(key), IntToByteArr(seconds))
}

func (c *client) expireAt(key string, unixTime int64) error {
	return c.sendCommand(cmdExpireAt, []byte(key), Int64ToByteArr(unixTime))
}

func (c *client) pExpire(key string, milliseconds int64) error {
	return c.sendCommand(cmdPExpire, []byte(key), Int64ToByteArr(milliseconds))
}

func (c *client) pExpireAt(key string, unixTime int64) error {
	return c.sendCommand(cmdPExpireAt, []byte(key), Int64ToByteArr(unixTime))
}

func (c *client) ttl(key string) error {
	return c.sendCommand(cmdTTL, []byte(key))
}

func (c *client) pttl(key string) error {
	return c.sendCommand(cmdPTTL, []byte(key))
}

func (c *client) move(key string, dbIndex int) error {
	return c.sendCommand(cmdMove, []byte(key), IntToByteArr(dbIndex))
}

func (c *client) getSet(key, value string) error {
	return c.sendCommand(cmdGetSet, []byte(key), []byte(value))
}

func (c *client) mget(keys ...string) error {
	return c.sendCommand(cmdMGet, StrArrToByteArrArr(keys)...)
}

func (c *client) setnx(key, value string) error {
	return c.sendCommand(cmdSetNx, []byte(key), []byte(value))
}

func (c *client) setex(key string, seconds int, value string) error {
	return c.sendCommand(cmdSetEx, []byte(key), IntToByteArr(seconds), []byte(value))
}

func (c *client) pSetEx(key string, milliseconds int64, value string) error {
	return c.sendCommand(cmdSetEx, []byte(key), Int64ToByteArr(milliseconds), []byte(value))
}

func (c *client) mset(keysvalues ...string) error {
	return c.sendCommand(cmdMSet, StrArrToByteArrArr(keysvalues)...)
}

func (c *client) msetnx(keysvalues ...string) error {
	return c.sendCommand(cmdMSetNx, StrArrToByteArrArr(keysvalues)...)
}

func (c *client) decrBy(key string, decrement int64) error {
	return c.sendCommand(cmdDecrBy, []byte(key), Int64ToByteArr(decrement))
}

func (c *client) decr(key string) error {
	return c.sendCommand(cmdDecr, []byte(key))
}

func (c *client) incrBy(key string, increment int64) error {
	return c.sendCommand(cmdIncrBy, []byte(key), Int64ToByteArr(increment))
}

func (c *client) incr(key string) error {
	return c.sendCommand(cmdIncr, []byte(key))
}

func (c *client) append(key, value string) error {
	return c.sendCommand(cmdAppend, []byte(key), []byte(value))
}

func (c *client) substr(key string, start, end int) error {
	return c.sendCommand(cmdSubstr, []byte(key), IntToByteArr(start), IntToByteArr(end))
}

func (c *client) hset(key, field, value string) error {
	return c.sendCommand(cmdHSet, []byte(key), []byte(field), []byte(value))
}

func (c *client) hget(key, field string) error {
	return c.sendCommand(cmdHGet, []byte(key), []byte(field))
}

func (c *client) hsetnx(key, field, value string) error {
	return c.sendCommand(cmdHSetNx, []byte(key), []byte(field), []byte(value))
}

func (c *client) hmset(key string, hash map[string]string) error {
	params := make([][]byte, 0)
	params = append(params, []byte(key))
	for k, v := range hash {
		params = append(params, []byte(k))
		params = append(params, []byte(v))
	}
	return c.sendCommand(cmdHMSet, params...)
}

func (c *client) hmget(key string, fields ...string) error {
	return c.sendCommand(cmdHMGet, StrStrArrToByteArrArr(key, fields)...)
}

func (c *client) hincrBy(key, field string, increment int64) error {
	return c.sendCommand(cmdHIncrBy, []byte(key), []byte(field), Int64ToByteArr(increment))
}

func (c *client) hexists(key, field string) error {
	return c.sendCommand(cmdHExists, []byte(key), []byte(field))
}

func (c *client) hdel(key string, fields ...string) error {
	return c.sendCommand(cmdHDel, StrStrArrToByteArrArr(key, fields)...)
}

func (c *client) hlen(key string) error {
	return c.sendCommand(cmdHLen, []byte(key))
}

func (c *client) hkeys(key string) error {
	return c.sendCommand(cmdHKeys, []byte(key))
}

func (c *client) hvals(key string) error {
	return c.sendCommand(cmdHVals, []byte(key))
}

func (c *client) hgetAll(key string) error {
	return c.sendCommand(cmdHGetAll, []byte(key))
}

func (c *client) rpush(key string, fields ...string) error {
	return c.sendCommand(cmdRPush, StrStrArrToByteArrArr(key, fields)...)
}

func (c *client) lpush(key string, fields ...string) error {
	return c.sendCommand(cmdRPush, StrStrArrToByteArrArr(key, fields)...)
}

func (c *client) llen(key string) error {
	return c.sendCommand(cmdLLen, []byte(key))
}

func (c *client) lrange(key string, start, end int64) error {
	return c.sendCommand(cmdLRange, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) ltrim(key string, start, end int64) error {
	return c.sendCommand(cmdLtrim, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) lindex(key string, index int64) error {
	return c.sendCommand(cmdLIndex, []byte(key), Int64ToByteArr(index))
}

func (c *client) lset(key string, index int64, value string) error {
	return c.sendCommand(cmdLSet, []byte(key), Int64ToByteArr(index), []byte(value))
}

func (c *client) lrem(key string, count int64, value string) error {
	return c.sendCommand(cmdLRem, []byte(key), Int64ToByteArr(count), []byte(value))
}

func (c *client) lpop(key string) error {
	return c.sendCommand(cmdLPop, []byte(key))
}

func (c *client) rPop(key string) error {
	return c.sendCommand(cmdRPop, []byte(key))
}

func (c *client) rpopLpush(srcKey, destKey string) error {
	return c.sendCommand(cmdRPopLPush, []byte(srcKey), []byte(destKey))
}

func (c *client) sAdd(key string, members ...string) error {
	return c.sendCommand(cmdSAdd, StrStrArrToByteArrArr(key, members)...)
}

func (c *client) sMembers(key string) error {
	return c.sendCommand(cmdSMembers, []byte(key))
}

func (c *client) sRem(key string, members ...string) error {
	return c.sendCommand(cmdSRem, StrStrArrToByteArrArr(key, members)...)
}

func (c *client) sPop(key string) error {
	return c.sendCommand(cmdSPop, []byte(key))
}

func (c *client) sPopBatch(key string, count int64) error {
	return c.sendCommand(cmdSPop, []byte(key), Int64ToByteArr(count))
}

func (c *client) smove(srcKey, destKey, member string) error {
	return c.sendCommand(cmdSMove, []byte(srcKey), []byte(destKey), []byte(member))
}

func (c *client) sCard(key string) error {
	return c.sendCommand(cmdSCard, []byte(key))
}

func (c *client) sIsMember(key, member string) error {
	return c.sendCommand(cmdSIsMember, []byte(key), []byte(member))
}

func (c *client) sInter(keys ...string) error {
	return c.sendCommand(cmdSInter, StrArrToByteArrArr(keys)...)
}

func (c *client) sInterStore(destKey string, keys ...string) error {
	return c.sendCommand(cmdSInterStore, StrStrArrToByteArrArr(destKey, keys)...)
}

func (c *client) sUnion(keys ...string) error {
	return c.sendCommand(cmdSUnion, StrArrToByteArrArr(keys)...)
}

func (c *client) sUnionStore(destKey string, keys ...string) error {
	return c.sendCommand(cmdSUnionStore, StrStrArrToByteArrArr(destKey, keys)...)
}

func (c *client) sDiff(keys ...string) error {
	return c.sendCommand(cmdSDiff, StrArrToByteArrArr(keys)...)
}

func (c *client) sDiffStore(destKey string, keys ...string) error {
	return c.sendCommand(cmdSDiffStore, StrStrArrToByteArrArr(destKey, keys)...)
}

func (c *client) sRandMember(key string) error {
	return c.sendCommand(cmdSRandMember, []byte(key))
}

func (c *client) zAdd(key string, score float64, member string, params ...*ZAddParams) error {
	newArr := make([][]byte, 0)
	if len(params) == 0 {
		return c.sendCommand(cmdZAdd, []byte(key), Float64ToByteArr(score), []byte(member))
	}
	newArr = append(newArr, Float64ToByteArr(score))
	newArr = append(newArr, []byte(member))
	return c.sendCommand(cmdZAdd, params[0].getByteParams([]byte(key), newArr...)...)
}

func (c *client) ZAddByMap(key string, scoreMembers map[string]float64, params ...*ZAddParams) error {
	newArr := make([][]byte, 0)
	if len(params) == 0 {
		newArr = append(newArr, []byte(key))
		for k, v := range scoreMembers {
			newArr = append(newArr, Float64ToByteArr(v))
			newArr = append(newArr, []byte(k))
		}
		return c.sendCommand(cmdZAdd, newArr...)
	}
	for k, v := range scoreMembers {
		newArr = append(newArr, Float64ToByteArr(v))
		newArr = append(newArr, []byte(k))
	}
	return c.sendCommand(cmdZAdd, params[0].getByteParams([]byte(key), newArr...)...)
}

func (c *client) zRange(key string, start, end int64) error {
	return c.sendCommand(cmdZRange, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) zRem(key string, members ...string) error {
	return c.sendCommand(cmdZRem, StrStrArrToByteArrArr(key, members)...)
}

func (c *client) zIncrBy(key string, score float64, member string) error {
	return c.sendCommand(cmdZIncrBy, []byte(key), Float64ToByteArr(score), []byte(member))
}

func (c *client) zRank(key, member string) error {
	return c.sendCommand(cmdZRank, []byte(key), []byte(member))
}

func (c *client) zRevRank(key, member string) error {
	return c.sendCommand(cmdZRevRank, []byte(key), []byte(member))
}

func (c *client) zRevRange(key string, start, end int64) error {
	return c.sendCommand(cmdZRevRange, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) ZRangeWithScores(key string, start, end int64) error {
	return c.sendCommand(cmdZRange, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end), keywordWithScores.getRaw())
}

func (c *client) ZRevRangeWithScores(key string, start, end int64) error {
	return c.sendCommand(cmdZRevRange, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end), keywordWithScores.getRaw())
}

func (c *client) zCard(key string) error {
	return c.sendCommand(cmdZCard, []byte(key))
}

func (c *client) zScore(key, member string) error {
	return c.sendCommand(cmdZScore, []byte(key), []byte(member))
}

func (c *client) watch(keys ...string) error {
	return c.sendCommand(cmdWatch, StrArrToByteArrArr(keys)...)
}

func (c *client) sort(key string, sortingParameters ...*SortParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.getParams()...)
	}
	return c.sendCommand(cmdSort, newArr...)
}

func (c *client) sortMulti(key, destKey string, sortingParameters ...*SortParams) error {
	newArr := make([][]byte, 0)
	newArr = append(newArr, []byte(key))
	for _, p := range sortingParameters {
		newArr = append(newArr, p.getParams()...)
	}
	newArr = append(newArr, keywordStore.getRaw())
	newArr = append(newArr, []byte(destKey))
	return c.sendCommand(cmdSort, newArr...)
}

func (c *client) blpop(args []string) error {
	return c.sendCommand(cmdBLPop, StrArrToByteArrArr(args)...)
}

func (c *client) brpop(args []string) error {
	return c.sendCommand(cmdBRPop, StrArrToByteArrArr(args)...)
}

func (c *client) zCount(key string, min, max float64) error {
	return c.sendCommand(cmdZCount, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max))
}

func (c *client) zRangeByScore(key string, min, max float64) error {
	return c.sendCommand(cmdZRangeByScore, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max))
}

func (c *client) zRangeByScoreWithScores(key string, min, max float64) error {
	return c.sendCommand(cmdZRangeByScore, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max), keywordWithScores.getRaw())
}

func (c *client) zRevRangeByScore(key string, max, min float64) error {
	return c.sendCommand(cmdZRevRangeByScore, []byte(key), Float64ToByteArr(max), Float64ToByteArr(min))
}

func (c *client) zRevRangeByScoreWithScores(key string, max, min float64) error {
	return c.sendCommand(cmdZRevRangeByScore, []byte(key), Float64ToByteArr(max), Float64ToByteArr(min), keywordWithScores.getRaw())
}

func (c *client) zRemRangeByRank(key string, start, end int64) error {
	return c.sendCommand(cmdZRemRangeByRank, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) ZRemRangeByScore(key string, min, max float64) error {
	return c.sendCommand(cmdZRemRangeByScore, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max))
}

func (c *client) zunionstore(destKey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(destKey))
	arr = append(arr, IntToByteArr(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdZUnionStore, arr...)
}

func (c *client) zunionstoreWithParams(destKey string, params *ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(destKey))
	arr = append(arr, IntToByteArr(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.getParams()...)
	return c.sendCommand(cmdZUnionStore, arr...)
}

func (c *client) zinterstore(destKey string, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(destKey))
	arr = append(arr, IntToByteArr(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdZInterStore, arr...)
}

func (c *client) zinterstoreWithParams(destKey string, params *ZParams, sets ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(destKey))
	arr = append(arr, IntToByteArr(len(sets)))
	for _, s := range sets {
		arr = append(arr, []byte(s))
	}
	arr = append(arr, params.getParams()...)
	return c.sendCommand(cmdZInterStore, arr...)
}

func (c *client) zlexcount(key, min, max string) error {
	return c.sendCommand(cmdZLexCount, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLex(key, min, max string) error {
	return c.sendCommand(cmdZRangeByLex, []byte(key), []byte(min), []byte(max))
}

func (c *client) zrangeByLexBatch(key, min, max string, offset, count int) error {
	return c.sendCommand(cmdZRangeByLex, []byte(key), []byte(min), []byte(max), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count))
}

func (c *client) zrevrangeByLex(key, max, min string) error {
	return c.sendCommand(cmdZRevRangeByLex, []byte(key), []byte(max), []byte(min))
}

func (c *client) zrevrangeByLexBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(cmdZRevRangeByLex, []byte(key), []byte(max), []byte(min), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count))
}

func (c *client) zremrangeByLex(key, min, max string) error {
	return c.sendCommand(cmdZRemRangeByLex, []byte(key), []byte(min), []byte(max))
}

func (c *client) strLen(key string) error {
	return c.sendCommand(cmdStrLen, []byte(key))
}

func (c *client) lPushX(key string, string ...string) error {
	return c.sendCommand(cmdLPushX, StrStrArrToByteArrArr(key, string)...)
}

func (c *client) persist(key string) error {
	return c.sendCommand(cmdPersist, []byte(key))
}

func (c *client) rPushX(key string, string ...string) error {
	return c.sendCommand(cmdRPushX, StrStrArrToByteArrArr(key, string)...)
}

func (c *client) echo(string string) error {
	return c.sendCommand(cmdEcho, []byte(string))
}

func (c *client) brpoplpush(source, destination string, timeout int) error {
	return c.sendCommand(cmdBRPopLPush, []byte(source), []byte(destination), IntToByteArr(timeout))
}

func (c *client) setBit(key string, offset int64, value string) error {
	return c.sendCommand(cmdSetBit, []byte(key), Int64ToByteArr(offset), []byte(value))
}

func (c *client) getBit(key string, offset int64) error {
	return c.sendCommand(cmdGetBit, []byte(key), Int64ToByteArr(offset))
}

func (c *client) setrange(key string, offset int64, value string) error {
	return c.sendCommand(cmdSetRange, []byte(key), Int64ToByteArr(offset), []byte(value))
}

func (c *client) getrange(key string, startOffset, endOffset int64) error {
	return c.sendCommand(cmdGetRange, []byte(key), Int64ToByteArr(startOffset), Int64ToByteArr(endOffset))
}

func (c *client) publish(channel, message string) error {
	return c.sendCommand(cmdPublish, []byte(channel), []byte(message))
}

func (c *client) unsubscribe(channels ...string) error {
	return c.sendCommand(cmdUnSubscribe, StrArrToByteArrArr(channels)...)
}

func (c *client) psubscribe(patterns ...string) error {
	return c.sendCommand(cmdPSubscribe, StrArrToByteArrArr(patterns)...)
}

func (c *client) punsubscribe(patterns ...string) error {
	return c.sendCommand(cmdPUnSubscribe, StrArrToByteArrArr(patterns)...)
}

func (c *client) subscribe(channels ...string) error {
	return c.sendCommand(cmdSubscribe, StrArrToByteArrArr(channels)...)
}

func (c *client) pubsub(subcommand string, args ...string) error {
	return c.sendCommand(cmdPubSub, StrStrArrToByteArrArr(subcommand, args)...)
}

func (c *client) configSet(parameter, value string) error {
	return c.sendCommand(cmdConfig, keywordSet.getRaw(), []byte(parameter), []byte(value))
}

func (c *client) configGet(pattern string) error {
	return c.sendCommand(cmdConfig, keywordGet.getRaw(), []byte(pattern))
}

func (c *client) eval(script string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(script))
	arr = append(arr, IntToByteArr(keyCount))
	arr = append(arr, StrArrToByteArrArr(params)...)
	return c.sendCommand(cmdEval, arr...)
}

func (c *client) evalsha(sha1 string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(sha1))
	arr = append(arr, IntToByteArr(keyCount))
	arr = append(arr, StrArrToByteArrArr(params)...)
	return c.sendCommand(cmdEvalSha, arr...)
}

func (c *client) scriptExists(sha1 ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, keywordExists.getRaw())
	arr = append(arr, StrArrToByteArrArr(sha1)...)
	return c.sendCommand(cmdScript, arr...)
}

func (c *client) scriptLoad(script string) error {
	return c.sendCommand(cmdScript, keywordLoad.getRaw(), []byte(script))
}

func (c *client) sentinel(args ...string) error {
	return c.sendCommand(cmdSentinel, StrArrToByteArrArr(args)...)
}

func (c *client) dump(key string) error {
	return c.sendCommand(cmdDump, []byte(key))
}

func (c *client) restore(key string, ttl int, serializedValue []byte) error {
	return c.sendCommand(cmdRestore, []byte(key), IntToByteArr(ttl), serializedValue)
}

func (c *client) incrByFloat(key string, increment float64) error {
	return c.sendCommand(cmdIncrByFloat, []byte(key), Float64ToByteArr(increment))
}

func (c *client) sRandMemberBatch(key string, count int) error {
	return c.sendCommand(cmdSRandMember, []byte(key), IntToByteArr(count))
}

func (c *client) clientKill(client string) error {
	return c.sendCommand(cmdClient, keywordKill.getRaw(), []byte(client))
}

func (c *client) clientGetname() error {
	return c.sendCommand(cmdClient, keywordGetName.getRaw())
}

func (c *client) clientList() error {
	return c.sendCommand(cmdClient, keywordList.getRaw())
}

func (c *client) clientSetname(name string) error {
	return c.sendCommand(cmdClient, keywordSetName.getRaw(), []byte(name))
}

func (c *client) time() error {
	return c.sendCommand(cmdTime)
}

func (c *client) migrate(host string, port int, key string, destinationDb int, timeout int) error {
	return c.sendCommand(cmdMigrate, []byte(host), IntToByteArr(port), []byte(key), IntToByteArr(destinationDb), IntToByteArr(timeout))
}

func (c *client) hincrByFloat(key, field string, increment float64) error {
	return c.sendCommand(cmdHIncrByFloat, []byte(key), []byte(field), Float64ToByteArr(increment))
}

func (c *client) waitReplicas(replicas int, timeout int64) error {
	return c.sendCommand(cmdWait, IntToByteArr(replicas), Int64ToByteArr(timeout))
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
	return c.sendCommand(cmdGeoAdd, []byte(key), Float64ToByteArr(longitude), Float64ToByteArr(latitude), []byte(member))
}

func (c *client) geoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	for k, v := range memberCoordinateMap {
		arr = append(arr, Float64ToByteArr(v.longitude))
		arr = append(arr, Float64ToByteArr(v.latitude))
		arr = append(arr, []byte(k))
	}
	return c.sendCommand(cmdGeoAdd, arr...)
}

func (c *client) geodist(key, member1, member2 string, unit ...*GeoUnit) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member1))
	arr = append(arr, []byte(member2))
	for _, u := range unit {
		arr = append(arr, u.getRaw())
	}
	return c.sendCommand(cmdGeoDist, arr...)
}

func (c *client) geohash(key string, members ...string) error {
	return c.sendCommand(cmdGeoHash, StrStrArrToByteArrArr(key, members)...)
}

func (c *client) geopos(key string, members ...string) error {
	return c.sendCommand(cmdGeoPos, StrStrArrToByteArrArr(key, members)...)
}

func (c *client) flushDB() error {
	return c.sendCommand(cmdFlushDB)
}

func (c *client) dbSize() error {
	return c.sendCommand(cmdDbSize)
}

func (c *client) flushAll() error {
	return c.sendCommand(cmdFlushAll)
}

func (c *client) save() error {
	return c.sendCommand(cmdSave)
}

func (c *client) bgsave() error {
	return c.sendCommand(cmdBgSave)
}

func (c *client) bgrewriteaof() error {
	return c.sendCommand(cmdBgRewriteAof)
}

func (c *client) lastsave() error {
	return c.sendCommand(cmdLastSave)
}

func (c *client) shutdown() error {
	return c.sendCommand(cmdShutdown)
}

func (c *client) slaveof(host string, port int) error {
	return c.sendCommand(cmdSlaveOf, []byte(host), IntToByteArr(port))
}

func (c *client) slaveofNoOne() error {
	return c.sendCommand(cmdSlaveOf, keywordNo.getRaw(), keywordOne.getRaw())
}

func (c *client) getDB() int {
	return c.Db
}

func (c *client) debug(params DebugParams) error {
	return c.sendCommand(cmdDebug, StrArrToByteArrArr(params.command)...)
}

func (c *client) configResetStat() error {
	return c.sendCommand(cmdConfig, keywordResetStat.getRaw())
}

func (c *client) zRangeByScoreBatch(key string, min, max float64, offset, count int) error {
	return c.sendCommand(cmdZRangeByScore, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count))
}

func (c *client) zRangeByScoreWithScoresBatch(key string, min, max float64, offset, count int) error {
	return c.sendCommand(cmdZRangeByScore, []byte(key), Float64ToByteArr(min), Float64ToByteArr(max), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count), keywordWithScores.getRaw())
}

func (c *client) zrevrangeByScoreBatch(key, max, min string, offset, count int) error {
	return c.sendCommand(cmdZRevRangeByScore, []byte(key), []byte(max), []byte(min), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count))
}

func (c *client) zRevRangeByScoreWithScoresBatch(key string, max, min float64, offset, count int) error {
	return c.sendCommand(cmdZRevRangeByScore, []byte(key), Float64ToByteArr(max), Float64ToByteArr(min), keywordLimit.getRaw(),
		IntToByteArr(offset), IntToByteArr(count), keywordWithScores.getRaw())
}

func (c *client) linsert(key string, where *ListOption, pivot, value string) error {
	return c.sendCommand(cmdLInsert, []byte(key), where.getRaw(), []byte(pivot), []byte(value))
}

func (c *client) bitcount(key string) error {
	return c.sendCommand(cmdBitCount, []byte(key))
}

func (c *client) bitcountRange(key string, start, end int64) error {
	return c.sendCommand(cmdBitCount, []byte(key), Int64ToByteArr(start), Int64ToByteArr(end))
}

func (c *client) bitpos(key string, value bool, params ...*BitPosParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, BoolToByteArr(value))
	for _, p := range params {
		arr = append(arr, p.params...)
	}
	return c.sendCommand(cmdBitPos, arr...)
}

func (c *client) scan(cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.getParams()...)
	}
	return c.sendCommand(cmdScan, arr...)
}

func (c *client) hscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.getParams()...)
	}
	return c.sendCommand(cmdHScan, arr...)
}

func (c *client) sscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.getParams()...)
	}
	return c.sendCommand(cmdSScan, arr...)
}

func (c *client) zscan(key, cursor string, params ...*ScanParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(cursor))
	for _, p := range params {
		arr = append(arr, p.getParams()...)
	}
	return c.sendCommand(cmdZScan, arr...)
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
	return c.sendCommand(cmdPfAdd, StrStrArrToByteArrArr(key, elements)...)
}

func (c *client) georadius(key string, longitude, latitude, radius float64, unit *GeoUnit, param ...*GeoRadiusParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, Float64ToByteArr(longitude))
	arr = append(arr, Float64ToByteArr(latitude))
	arr = append(arr, Float64ToByteArr(radius))
	arr = append(arr, unit.getRaw())
	if len(param) == 0 {
		return c.sendCommand(cmdGeoRadius, arr...)
	}
	return c.sendCommand(cmdGeoRadius, param[0].getParams(arr)...)
}

func (c *client) georadiusByMember(key, member string, radius float64, unit *GeoUnit, param ...*GeoRadiusParams) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(key))
	arr = append(arr, []byte(member))
	arr = append(arr, Float64ToByteArr(radius))
	arr = append(arr, unit.getRaw())
	if len(param) == 0 {
		return c.sendCommand(cmdGeoRadiusByMember, arr...)
	}
	return c.sendCommand(cmdGeoRadiusByMember, param[0].getParams(arr)...)
}

func (c *client) bitfield(key string, arguments ...string) error {
	return c.sendCommand(cmdBitField, StrStrArrToByteArrArr(key, arguments)...)
}

func (c *client) randomKey() error {
	return c.sendCommand(cmdRandomKey)
}

func (c *client) bitop(op BitOP, destKey string, srcKeys ...string) error {
	kw := BitOpAnd
	switch op.name {
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
	arr = append(arr, kw.getRaw())
	arr = append(arr, []byte(destKey))
	for _, s := range srcKeys {
		arr = append(arr, []byte(s))
	}
	return c.sendCommand(cmdBitOp, arr...)
}

func (c *client) pfmerge(destkey string, sourcekeys ...string) error {
	return c.sendCommand(cmdPfMerge, StrStrArrToByteArrArr(destkey, sourcekeys)...)
}

func (c *client) pfcount(keys ...string) error {
	return c.sendCommand(cmdPfCount, StrArrToByteArrArr(keys)...)
}

func (c *client) slowlogReset() error {
	return c.sendCommand(cmdSlowLog, keywordReset.getRaw())
}

func (c *client) slowlogLen() error {
	return c.sendCommand(cmdSlowLog, keywordLen.getRaw())
}

func (c *client) slowlogGet(entries ...int64) error {
	arr := make([][]byte, 0)
	arr = append(arr, keywordGet.getRaw())
	for _, e := range entries {
		arr = append(arr, Int64ToByteArr(e))
	}
	return c.sendCommand(cmdSlowLog, arr...)
}

func (c *client) objectRefcount(str string) error {
	return c.sendCommand(cmdObject, keywordRefCount.getRaw(), []byte(str))
}

func (c *client) objectEncoding(str string) error {
	return c.sendCommand(cmdObject, keywordEncoding.getRaw(), []byte(str))
}

func (c *client) objectIdletime(str string) error {
	return c.sendCommand(cmdObject, keywordIdleTime.getRaw(), []byte(str))
}

func (c *client) clusterNodes() error {
	return c.sendCommand(cmdCluster, []byte(clusterNodes))
}

func (c *client) clusterMeet(ip string, port int) error {
	return c.sendCommand(cmdCluster, []byte(clusterMeet), []byte(ip), IntToByteArr(port))
}

func (c *client) clusterAddSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(clusterAddSlots))
	for _, s := range slots {
		arr = append(arr, IntToByteArr(s))
	}
	return c.sendCommand(cmdCluster, arr...)
}

func (c *client) clusterDelSlots(slots ...int) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(clusterDelSlots))
	for _, s := range slots {
		arr = append(arr, IntToByteArr(s))
	}
	return c.sendCommand(cmdCluster, arr...)
}

func (c *client) clusterInfo() error {
	return c.sendCommand(cmdCluster, []byte(clusterInfo))
}

func (c *client) clusterGetKeysInSlot(slot int, count int) error {
	return c.sendCommand(cmdCluster, []byte(clusterGetKeysInSlot), IntToByteArr(slot), IntToByteArr(count))
}

func (c *client) clusterSetSlotNode(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetSlotNode), IntToByteArr(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotMigrating(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetSlotMigrating), IntToByteArr(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotImporting(slot int, nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetSlotImporting), IntToByteArr(slot), []byte(nodeID))
}

func (c *client) clusterSetSlotStable(slot int) error {
	return c.sendCommand(cmdCluster, []byte(clusterSetSlotStable), IntToByteArr(slot))
}

func (c *client) clusterForget(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterForget), []byte(nodeID))
}

func (c *client) clusterFlushSlots() error {
	return c.sendCommand(cmdCluster, []byte(clusterFlushSlot))
}

func (c *client) clusterKeySlot(key string) error {
	return c.sendCommand(cmdCluster, []byte(clusterKeySlot), []byte(key))
}

func (c *client) clusterCountKeysInSlot(slot int) error {
	return c.sendCommand(cmdCluster, []byte(clusterCountKeyInSlot), IntToByteArr(slot))
}

func (c *client) clusterSaveConfig() error {
	return c.sendCommand(cmdCluster, []byte(clusterSaveConfig))
}

func (c *client) clusterReplicate(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterReplicate), []byte(nodeID))
}

func (c *client) clusterSlaves(nodeID string) error {
	return c.sendCommand(cmdCluster, []byte(clusterSlaves), []byte(nodeID))
}

func (c *client) clusterFailover() error {
	return c.sendCommand(cmdCluster, []byte(clusterFailOver))
}

func (c *client) clusterSlots() error {
	return c.sendCommand(cmdCluster, []byte(clusterSlots))
}

func (c *client) clusterReset(resetType Reset) error {
	return c.sendCommand(cmdCluster, []byte(clusterReset), resetType.getRaw())
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
	return c.sendCommand(cmdSentinel, []byte(sentinelFailOver), []byte(masterName))
}

func (c *client) sentinelMonitor(masterName, ip string, port, quorum int) error {
	return c.sendCommand(cmdSentinel, []byte(sentinelMonitor), []byte(masterName), []byte(ip), IntToByteArr(port), IntToByteArr(quorum))
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
	return c.sendCommandByStr(sentinelFailOver, arr...)
}

func (c *client) pubsubChannels(pattern string) error {
	return c.sendCommand(cmdPubSub, []byte(pubSubChannels), []byte(pattern))
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

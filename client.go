package godis

type Client struct {
	*Connection
	//Host              string
	//Port              int
	//ConnectionTimeout int
	//SoTimeout         int
	Password string
	Db       int
	//IsInMulti         bool
	//IsInWatch         bool
	//Ssl               bool
}

func NewClient(shardInfo ShardInfo) *Client {
	db := 0
	if shardInfo.Db != 0 {
		db = shardInfo.Db
	}
	client := &Client{
		//Host:              options.Host,
		//Port:              options.Port,
		//ConnectionTimeout: options.ConnectionTimeout,
		//SoTimeout:         options.SoTimeout,
		Password: shardInfo.Password,
		Db:       db,
		//IsInMulti:         options.IsInMulti,
		//IsInWatch:         options.IsInWatch,
		//Ssl:               options.Ssl,
	}
	client.Connection = NewConnection(shardInfo.Host, shardInfo.Port, shardInfo.ConnectionTimeout, shardInfo.SoTimeout, shardInfo.Ssl)
	return client
}

func (c *Client) Host() string {
	return c.Connection.Host
}

func (c *Client) Port() int {
	return c.Connection.Port
}

func (c *Client) Connect() error {
	err := c.Connection.Connect()
	if err != nil {
		return err
	}
	if c.Password != "" {
		err = c.Auth(c.Password)
		if err != nil {
			return err
		}
		_, err = c.getStatusCodeReply()
		if err != nil {
			return err
		}
	}
	if c.Db > 0 {
		err = c.Select(c.Db)
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

func (c *Client) Close() error {
	return c.Connection.Close()
}

func (c *Client) Ping() error {
	return c.SendCommand(CMD_PING)
}

func (c *Client) Quit() error {
	return c.SendCommand(CMD_QUIT)
}

func (c *Client) Info() error {
	return c.SendCommand(CMD_INFO)
}

func (c *Client) Auth(password string) error {
	c.Password = password
	return c.SendCommand(CMD_AUTH, []byte(password))
}

func (c *Client) Select(index int) error {
	return c.SendCommand(CMD_SELECT, IntToByteArray(index))
}

func (c *Client) Set(key, value string) error {
	return c.SendCommand(CMD_SET, []byte(key), []byte(value))
}

func (c *Client) SetWithParams(key, value, nxxx, expx string, time int64) error {
	return c.SendCommand(CMD_SET, []byte(key), []byte(value), []byte(nxxx), []byte(expx), Int64ToByteArray(time))
}

func (c *Client) Get(key string) error {
	return c.SendCommand(CMD_GET, []byte(key))
}

func (c *Client) Del(keys ...string) error {
	return c.SendCommand(CMD_DEL, StringArrayToByteArray(keys)...)
}

func (c *Client) Exists(keys ...string) error {
	return c.SendCommand(CMD_EXISTS, StringArrayToByteArray(keys)...)
}

func (c *Client) Type(key string) error {
	return c.SendCommand(CMD_TYPE, []byte(key))
}

func (c *Client) Keys(pattern string) error {
	return c.SendCommand(CMD_KEYS, []byte(pattern))
}

func (c *Client) Rename(oldKey, newKey string) error {
	return c.SendCommand(CMD_RENAME, []byte(oldKey), []byte(newKey))
}

func (c *Client) Renamenx(oldKey, newKey string) error {
	return c.SendCommand(CMD_RENAMEX, []byte(oldKey), []byte(newKey))
}

func (c *Client) Expire(key string, seconds int) error {
	return c.SendCommand(CMD_EXPIRE, []byte(key), IntToByteArray(seconds))
}

func (c *Client) ExpireAt(key string, unixTime int64) error {
	return c.SendCommand(CMD_EXPIREAT, []byte(key), Int64ToByteArray(unixTime))
}

func (c *Client) Ttl(key string) error {
	return c.SendCommand(CMD_TTL, []byte(key))
}

func (c *Client) Move(key string, dbIndex int) error {
	return c.SendCommand(CMD_MOVE, []byte(key), IntToByteArray(dbIndex))
}

func (c *Client) GetSet(key, value string) error {
	return c.SendCommand(CMD_GETSET, []byte(key), []byte(value))
}

func (c *Client) Mget(keys ...string) error {
	return c.SendCommand(CMD_MGET, StringArrayToByteArray(keys)...)
}

func (c *Client) Setnx(key, value string) error {
	return c.SendCommand(CMD_SETNX, []byte(key), []byte(value))
}

func (c *Client) Setex(key string, seconds int, value string) error {
	return c.SendCommand(CMD_SETEX, []byte(key), IntToByteArray(seconds), []byte(value))
}

func (c *Client) Mset(keysvalues ...string) error {
	return c.SendCommand(CMD_MSET, StringArrayToByteArray(keysvalues)...)
}

func (c *Client) Msetnx(keysvalues ...string) error {
	return c.SendCommand(CMD_MSETNX, StringArrayToByteArray(keysvalues)...)
}

func (c *Client) DecrBy(key string, decrement int64) error {
	return c.SendCommand(CMD_DECRBY, []byte(key), Int64ToByteArray(decrement))
}

func (c *Client) Decr(key string) error {
	return c.SendCommand(CMD_DECR, []byte(key))
}

func (c *Client) IncrBy(key string, increment int64) error {
	return c.SendCommand(CMD_INCRBY, []byte(key), Int64ToByteArray(increment))
}

func (c *Client) Incr(key string) error {
	return c.SendCommand(CMD_INCR, []byte(key))
}

func (c *Client) Append(key, value string) error {
	return c.SendCommand(CMD_APPEND, []byte(key), []byte(value))
}

func (c *Client) Substr(key string, start, end int) error {
	return c.SendCommand(CMD_SUBSTR, []byte(key), IntToByteArray(start), IntToByteArray(end))
}

func (c *Client) Hset(key, field, value string) error {
	return c.SendCommand(CMD_HSET, []byte(key), []byte(field), []byte(value))
}

func (c *Client) Hget(key, field string) error {
	return c.SendCommand(CMD_HGET, []byte(key), []byte(field))
}

func (c *Client) Hsetnx(key, field, value string) error {
	return c.SendCommand(CMD_SETNX, []byte(key), []byte(field), []byte(value))
}

func (c *Client) Hmset(key string, hash map[string]string) error {
	params := make([][]byte, 0)
	params = append(params, []byte(key))
	for k, v := range hash {
		params = append(params, []byte(k))
		params = append(params, []byte(v))
	}
	return c.SendCommand(CMD_HMSET, params...)
}

func (c *Client) Hmget(key string, fields ...string) error {
	return c.SendCommand(CMD_HMGET, StringStringArrayToByteArray(key, fields)...)
}

func (c *Client) HincrBy(key, field string, increment int64) error {
	return c.SendCommand(CMD_HINCRBY, []byte(key), []byte(field), Int64ToByteArray(increment))
}

func (c *Client) Hexists(key, field string) error {
	return c.SendCommand(CMD_HEXISTS, []byte(key), []byte(field))
}

func (c *Client) Hdel(key string, fields ...string) error {
	return c.SendCommand(CMD_HDEL, StringStringArrayToByteArray(key, fields)...)
}

func (c *Client) Hlen(key string) error {
	return c.SendCommand(CMD_HLEN, []byte(key))
}

func (c *Client) Hkeys(key string) error {
	return c.SendCommand(CMD_HKEYS, []byte(key))
}

func (c *Client) Hvals(key string) error {
	return c.SendCommand(CMD_HVALS, []byte(key))
}

func (c *Client) HgetAll(key string) error {
	return c.SendCommand(CMD_HGETALL, []byte(key))
}

func (c *Client) Rpush(key string, fields ...string) error {
	return c.SendCommand(CMD_RPUSH, StringStringArrayToByteArray(key, fields)...)
}

func (c *Client) Lpush(key string, fields ...string) error {
	return c.SendCommand(CMD_RPUSH, StringStringArrayToByteArray(key, fields)...)
}

func (c *Client) Llen(key string) error {
	return c.SendCommand(CMD_LLEN, []byte(key))
}

func (c *Client) Lrange(key string, start, end int64) error {
	return c.SendCommand(CMD_LRANGE, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *Client) Ltrim(key string, start, end int64) error {
	return c.SendCommand(CMD_LTRIM, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *Client) Lindex(key string, index int64) error {
	return c.SendCommand(CMD_LINDEX, []byte(key), Int64ToByteArray(index))
}

func (c *Client) Lset(key string, index int64, value string) error {
	return c.SendCommand(CMD_LSET, []byte(key), Int64ToByteArray(index), []byte(value))
}

func (c *Client) Lrem(key string, count int64, value string) error {
	return c.SendCommand(CMD_LREM, []byte(key), Int64ToByteArray(count), []byte(value))
}

func (c *Client) Lpop(key string) error {
	return c.SendCommand(CMD_LPOP, []byte(key))
}

func (c *Client) Rpop(key string) error {
	return c.SendCommand(CMD_RPOP, []byte(key))
}

func (c *Client) RpopLpush(srckey, dstkey string) error {
	return c.SendCommand(CMD_RPOPLPUSH, []byte(srckey), []byte(dstkey))
}

func (c *Client) Sadd(key string, members ...string) error {
	return c.SendCommand(CMD_SADD, StringStringArrayToByteArray(key, members)...)
}

func (c *Client) Smembers(key string) error {
	return c.SendCommand(CMD_SMEMBERS, []byte(key))
}

func (c *Client) Srem(key string, members ...string) error {
	return c.SendCommand(CMD_DEL, StringStringArrayToByteArray(key, members)...)
}

func (c *Client) Spop(key string) error {
	return c.SendCommand(CMD_SPOP, []byte(key))
}

func (c *Client) SpopBatch(key string, count int64) error {
	return c.SendCommand(CMD_SPOP, []byte(key), Int64ToByteArray(count))
}

func (c *Client) Smove(srckey, dstkey, member string) error {
	return c.SendCommand(CMD_SMOVE, []byte(srckey), []byte(dstkey), []byte(member))
}

func (c *Client) Scard(key string) error {
	return c.SendCommand(CMD_SCARD, []byte(key))
}

func (c *Client) Sismember(key, member string) error {
	return c.SendCommand(CMD_SISMEMBER, []byte(key), []byte(member))
}

func (c *Client) Sinter(keys ...string) error {
	return c.SendCommand(CMD_SINTER, StringArrayToByteArray(keys)...)
}

func (c *Client) Sinterstore(dstkey string, keys ...string) error {
	return c.SendCommand(CMD_SINTERSTORE, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *Client) Sunion(keys ...string) error {
	return c.SendCommand(CMD_SUNION, StringArrayToByteArray(keys)...)
}

func (c *Client) Sunionstore(dstkey string, keys ...string) error {
	return c.SendCommand(CMD_SUNIONSTORE, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *Client) Sdiff(keys ...string) error {
	return c.SendCommand(CMD_SDIFF, StringArrayToByteArray(keys)...)
}

func (c *Client) Sdiffstore(dstkey string, keys ...string) error {
	return c.SendCommand(CMD_SDIFFSTORE, StringStringArrayToByteArray(dstkey, keys)...)
}

func (c *Client) Srandmember(key string) error {
	return c.SendCommand(CMD_SRANDMEMBER, []byte(key))
}

func (c *Client) Zadd(key string, score float64, member string) error {
	return c.SendCommand(CMD_ZADD, []byte(key), Float64ToByteArray(score), []byte(member))
}

func (c *Client) Zrange(key string, start, end int64) error {
	return c.SendCommand(CMD_ZRANGE, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *Client) Zrem(key string, members ...string) error {
	return c.SendCommand(CMD_ZREM, StringStringArrayToByteArray(key, members)...)
}

func (c *Client) Zincrby(key string, score float64, member string) error {
	return c.SendCommand(CMD_ZINCRBY, []byte(key), Float64ToByteArray(score), []byte(member))
}

func (c *Client) Zrank(key, member string) error {
	return c.SendCommand(CMD_ZRANK, []byte(key), []byte(member))
}

func (c *Client) Zrevrank(key, member string) error {
	return c.SendCommand(CMD_ZREVRANK, []byte(key), []byte(member))
}

func (c *Client) Zrevrange(key string, start, end int64) error {
	return c.SendCommand(CMD_ZREVRANGE, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *Client) ZrangeWithScores(key string, start, end int64) error {
	return c.SendCommand(CMD_ZRANGE, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), KEYWORD_WITHSCORES.GetRaw())
}

func (c *Client) ZrevrangeWithScores(key string, start, end int64) error {
	return c.SendCommand(CMD_ZRANGE, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end), KEYWORD_WITHSCORES.GetRaw())
}

func (c *Client) Zcard(key string) error {
	return c.SendCommand(CMD_ZCARD, []byte(key))
}

func (c *Client) Zscore(key, member string) error {
	return c.SendCommand(CMD_ZSCORE, []byte(key), []byte(member))
}

func (c *Client) Watch(keys ...string) error {
	return c.SendCommand(CMD_WATCH, StringArrayToByteArray(keys)...)
}

func (c *Client) Sort(key string) error {
	return c.SendCommand(CMD_SORT, []byte(key))
}

func (c *Client) Blpop(args []string) error {
	return c.SendCommand(CMD_BLPOP, StringArrayToByteArray(args)...)
}

func (c *Client) Brpop(args []string) error {
	return c.SendCommand(CMD_BRPOP, StringArrayToByteArray(args)...)
}

func (c *Client) Zcount(key, min, max string) error {
	return c.SendCommand(CMD_ZCOUNT, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZrangeByScore(key, min, max string) error {
	return c.SendCommand(CMD_ZRANGEBYSCORE, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZrevrangeByScore(key, min, max string) error {
	return c.SendCommand(CMD_ZREVRANGEBYSCORE, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZremrangeByRank(key string, start, end int64) error {
	return c.SendCommand(CMD_ZREMRANGEBYRANK, []byte(key), Int64ToByteArray(start), Int64ToByteArray(end))
}

func (c *Client) ZremrangeByScore(key, start, end string) error {
	return c.SendCommand(CMD_ZREMRANGEBYSCORE, []byte(key), []byte(start), []byte(end))
}

func (c *Client) Zunionstore(dstkey string, sets ...string) error {
	return c.SendCommand(CMD_ZUNIONSTORE, StringStringArrayToByteArray(dstkey, sets)...)
}

func (c *Client) Zinterstore(dstkey string, sets ...string) error {
	return c.SendCommand(CMD_ZINTERSTORE, StringStringArrayToByteArray(dstkey, sets)...)
}

func (c *Client) Zlexcount(key, min, max string) error {
	return c.SendCommand(CMD_ZLEXCOUNT, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZrangeByLex(key, min, max string) error {
	return c.SendCommand(CMD_ZRANGEBYLEX, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZrevrangeByLex(key, min, max string) error {
	return c.SendCommand(CMD_ZREVRANGEBYLEX, []byte(key), []byte(min), []byte(max))
}

func (c *Client) ZremrangeByLex(key, min, max string) error {
	return c.SendCommand(CMD_ZREMRANGEBYLEX, []byte(key), []byte(min), []byte(max))
}

func (c *Client) Strlen(key string) error {
	return c.SendCommand(CMD_STRLEN, []byte(key))
}

func (c *Client) Lpushx(key string, string ...string) error {
	return c.SendCommand(CMD_LPUSHX, StringStringArrayToByteArray(key, string)...)
}

func (c *Client) Persist(key string) error {
	return c.SendCommand(CMD_PERSIST, []byte(key))
}

func (c *Client) Rpushx(key string, string ...string) error {
	return c.SendCommand(CMD_RPUSHX, StringStringArrayToByteArray(key, string)...)
}

func (c *Client) Echo(string string) error {
	return c.SendCommand(CMD_ECHO, []byte(string))
}

//func (c *Client) Linsert(key string, final LIST_POSITION where, final String pivot, final String value)  error {
//	return c.SendCommand(CMD_LINSERT, []byte(key))
//}

func (c *Client) Brpoplpush(source, destination string, timeout int) error {
	return c.SendCommand(CMD_BRPOPLPUSH, []byte(source), []byte(destination), IntToByteArray(timeout))
}

func (c *Client) Setbit(key string, offset int64, value string) error {
	return c.SendCommand(CMD_SETBIT, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *Client) Getbit(key string, offset int64) error {
	return c.SendCommand(CMD_GETBIT, []byte(key), Int64ToByteArray(offset))
}

func (c *Client) Setrange(key string, offset int64, value string) error {
	return c.SendCommand(CMD_SETRANGE, []byte(key), Int64ToByteArray(offset), []byte(value))
}

func (c *Client) Getrange(key string, startOffset, endOffset int64) error {
	return c.SendCommand(CMD_GETRANGE, []byte(key), Int64ToByteArray(startOffset), Int64ToByteArray(endOffset))
}

func (c *Client) Publish(channel, message string) error {
	return c.SendCommand(CMD_PUBLISH, []byte(channel), []byte(message))
}

func (c *Client) Unsubscribe(channels ...string) error {
	return c.SendCommand(CMD_UNSUBSCRIBE, StringArrayToByteArray(channels)...)
}

func (c *Client) Psubscribe(patterns ...string) error {
	return c.SendCommand(CMD_PSUBSCRIBE, StringArrayToByteArray(patterns)...)
}

func (c *Client) Punsubscribe(patterns ...string) error {
	return c.SendCommand(CMD_PUNSUBSCRIBE, StringArrayToByteArray(patterns)...)
}

func (c *Client) Subscribe(channels ...string) error {
	return c.SendCommand(CMD_SUBSCRIBE, StringArrayToByteArray(channels)...)
}

func (c *Client) Pubsub(subcommand string, args ...string) error {
	return c.SendCommand(CMD_PUBSUB, StringStringArrayToByteArray(subcommand, args)...)
}

func (c *Client) ConfigSet(parameter, value string) error {
	return c.SendCommand(CMD_CONFIG, KEYWORD_SET.GetRaw(), []byte(parameter), []byte(value))
}

func (c *Client) ConfigGet(pattern string) error {
	return c.SendCommand(CMD_CONFIG, KEYWORD_GET.GetRaw(), []byte(pattern))
}

func (c *Client) Eval(script string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(script))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.SendCommand(CMD_EVAL, arr...)
}

func (c *Client) Evalsha(sha1 string, keyCount int, params ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, []byte(sha1))
	arr = append(arr, IntToByteArray(keyCount))
	arr = append(arr, StringArrayToByteArray(params)...)
	return c.SendCommand(CMD_EVALSHA, arr...)
}

func (c *Client) ScriptExists(sha1 ...string) error {
	arr := make([][]byte, 0)
	arr = append(arr, KEYWORD_EXISTS.GetRaw())
	arr = append(arr, StringArrayToByteArray(sha1)...)
	return c.SendCommand(CMD_SCRIPT, arr...)
}

func (c *Client) ScriptLoad(script string) error {
	return c.SendCommand(CMD_SCRIPT, KEYWORD_LOAD.GetRaw(), []byte(script))
}

func (c *Client) Sentinel(args ...string) error {
	return c.SendCommand(CMD_SENTINEL, StringArrayToByteArray(args)...)
}

func (c *Client) Dump(key string) error {
	return c.SendCommand(CMD_DUMP, []byte(key))
}

func (c *Client) Restore(key string, ttl int, serializedValue []byte) error {
	return c.SendCommand(CMD_RESTORE, []byte(key), IntToByteArray(ttl), serializedValue)
}

func (c *Client) IncrByFloat(key string, increment float64) error {
	return c.SendCommand(CMD_INCRBYFLOAT, []byte(key), Float64ToByteArray(increment))
}

func (c *Client) SrandmemberBatch(key string, count int) error {
	return c.SendCommand(CMD_SRANDMEMBER, []byte(key), IntToByteArray(count))
}

func (c *Client) ClientKill(client string) error {
	return c.SendCommand(CMD_CLIENT, KEYWORD_KILL.GetRaw(), []byte(client))
}

func (c *Client) ClientGetname() error {
	return c.SendCommand(CMD_CLIENT, KEYWORD_GETNAME.GetRaw())
}

func (c *Client) ClientList() error {
	return c.SendCommand(CMD_CLIENT, KEYWORD_LIST.GetRaw())
}

func (c *Client) ClientSetname(name string) error {
	return c.SendCommand(CMD_CLIENT, KEYWORD_SETNAME.GetRaw(), []byte(name))
}

func (c *Client) Time() error {
	return c.SendCommand(CMD_TIME)
}

func (c *Client) Migrate(host string, port int, key string, destinationDb int, timeout int) error {
	return c.SendCommand(CMD_MIGRATE, []byte(host), IntToByteArray(port), []byte(key), IntToByteArray(destinationDb), IntToByteArray(timeout))
}

func (c *Client) HincrByFloat(key, field string, increment float64) error {
	return c.SendCommand(CMD_HINCRBYFLOAT, []byte(key), []byte(field), Float64ToByteArray(increment))
}

//todo scan系列函数

func (c *Client) WaitReplicas(replicas int, timeout int64) error {
	return c.SendCommand(CMD_WAIT, IntToByteArray(replicas), Int64ToByteArray(timeout))
}

func (c *Client) Cluster(args ...[]byte) error {
	return c.SendCommand(CMD_CLUSTER, args...)
}

func (c *Client) Asking() error {
	return c.SendCommand(CMD_ASKING)
}

func (c *Client) Readonly() error {
	return c.SendCommand(CMD_READONLY)
}

func (c *Client) Geoadd(key []byte, longitude, latitude float64, member []byte) error {
	return c.SendCommand(CMD_GEOADD, key, Float64ToByteArray(longitude), Float64ToByteArray(latitude), member)
}

func (c *Client) Geodist(key, member1, member2 []byte) error {
	return c.SendCommand(CMD_GEODIST, key, member1, member2)
}

func (c *Client) Geohash(key []byte, members ...[]byte) error {
	arr := make([][]byte, 0)
	arr = append(arr, key)
	for _, m := range members {
		arr = append(arr, m)
	}
	return c.SendCommand(CMD_GEOHASH, arr...)
}

func (c *Client) Geopos(key []byte, members [][]byte) error {
	arr := make([][]byte, 0)
	arr = append(arr, key)
	for _, m := range members {
		arr = append(arr, m)
	}
	return c.SendCommand(CMD_GEOPOS, arr...)
}

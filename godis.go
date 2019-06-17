// godis
package godis

type ShardInfo struct {
	Host              string
	Port              int
	ConnectionTimeout int
	SoTimeout         int
	Password          string
	Db                int
	IsInMulti         bool
	IsInWatch         bool
	Ssl               bool
}

// Redis redis tool
type Redis struct {
	Client *Client
}

func NewRedis(shardInfo ShardInfo) *Redis {
	client := NewClient(shardInfo)
	return &Redis{Client: client}
}

func (r *Redis) Connect() error {
	return r.Client.Connect()
}

func (r *Redis) Close() error {
	if r != nil && r.Client != nil {
		return r.Client.Close()
	}
	return nil
}

//<editor-fold desc="rediscommands">
/**
 * Set the string value as value of the key. The string can't be longer than 1073741824 bytes (1
 * GB).
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param value
 * @return Status code reply
 */
func (r *Redis) Set(key, value string) (string, error) {
	err := r.Client.Set(key, value)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

/**
 * Set the string value as value of the key. The string can't be longer than 1073741824 bytes (1
 * GB).
 * @param key
 * @param value
 * @param nxxx NX|XX, NX -- Only set the key if it does not already exist. XX -- Only set the key
 *          if it already exist.
 * @param expx EX|PX, expire time units: EX = seconds; PX = milliseconds
 * @param time expire time in the units of <code>expx</code>
 * @return Status code reply
 */
func (r *Redis) SetWithParamsAndTime(key, value, nxxx, expx string, time int64) (string, error) {
	err := r.Client.SetWithParamsAndTime(key, value, nxxx, expx, time)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

/**
 * Get the value of the specified key. If the key does not exist null is returned. If the value
 * stored at key is not a string an error is returned because GET can only handle string values.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @return Bulk reply
 */
func (r *Redis) Get(key string) (string, error) {
	err := r.Client.Get(key)
	if err != nil {
		return "", err
	}
	return ByteToStringReply(r.Client.getBinaryBulkReply())
}

/**
 * Return the type of the value stored at key in form of a string. The type can be one of "none",
 * "string", "list", "set". "none" is returned if the key does not exist. Time complexity: O(1)
 * @param key
 * @return Status code reply, specifically: "none" if the key does not exist "string" if the key
 *         contains a String value "list" if the key contains a List value "set" if the key
 *         contains a Set value "zset" if the key contains a Sorted Set value "hash" if the key
 *         contains a Hash value
 */
func (r *Redis) Type(key string) (string, error) {
	err := r.Client.Get(key)
	if err != nil {
		return "", err
	}
	return ByteToStringReply(r.Client.getBinaryBulkReply())
}

/**
 * Set a timeout on the specified key. After the timeout the key will be automatically deleted by
 * the server. A key with an associated timeout is said to be volatile in Redis terminology.
 * <p>
 * Voltile keys are stored on disk like the other keys, the timeout is persistent too like all the
 * other aspects of the dataset. Saving a dataset containing expires and stopping the server does
 * not stop the flow of time as Redis stores on disk the time when the key will no longer be
 * available as Unix time, and not the remaining seconds.
 * <p>
 * Since Redis 2.1.3 you can update the value of the timeout of a key already having an expire
 * set. It is also possible to undo the expire at all turning the key into a normal key using the
 * {@link #persist(String) PERSIST} command.
 * <p>
 * Time complexity: O(1)
 * @see <a href="http://code.google.com/p/redis/wiki/ExpireCommand">ExpireCommand</a>
 * @param key
 * @param seconds
 * @return Integer reply, specifically: 1: the timeout was set. 0: the timeout was not set since
 *         the key already has an associated timeout (this may happen only in Redis versions &lt;
 *         2.1.3, Redis &gt;= 2.1.3 will happily update the timeout), or the key does not exist.
 */
func (r *Redis) Expire(key string, seconds int) (int64, error) {
	err := r.Client.Expire(key, seconds)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * EXPIREAT works exctly like {@link #expire(String, int) EXPIRE} but instead to get the number of
 * seconds representing the Time To Live of the key as a second argument (that is a relative way
 * of specifing the TTL), it takes an absolute one in the form of a UNIX timestamp (Number of
 * seconds elapsed since 1 Gen 1970).
 * <p>
 * EXPIREAT was introduced in order to implement the Append Only File persistence mode so that
 * EXPIRE commands are automatically translated into EXPIREAT commands for the append only file.
 * Of course EXPIREAT can also used by programmers that need a way to simply specify that a given
 * key should expire at a given time in the future.
 * <p>
 * Since Redis 2.1.3 you can update the value of the timeout of a key already having an expire
 * set. It is also possible to undo the expire at all turning the key into a normal key using the
 * {@link #persist(String) PERSIST} command.
 * <p>
 * Time complexity: O(1)
 * @see <a href="http://code.google.com/p/redis/wiki/ExpireCommand">ExpireCommand</a>
 * @param key
 * @param unixTime
 * @return Integer reply, specifically: 1: the timeout was set. 0: the timeout was not set since
 *         the key already has an associated timeout (this may happen only in Redis versions &lt;
 *         2.1.3, Redis &gt;= 2.1.3 will happily update the timeout), or the key does not exist.
 */
func (r *Redis) ExpireAt(key string, unixtime int64) (int64, error) {
	err := r.Client.ExpireAt(key, unixtime)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * The TTL command returns the remaining time to live in seconds of a key that has an
 * {@link #expire(String, int) EXPIRE} set. This introspection capability allows a Redis client to
 * check how many seconds a given key will continue to be part of the dataset.
 * @param key
 * @return Integer reply, returns the remaining time to live in seconds of a key that has an
 *         EXPIRE. In Redis 2.6 or older, if the Key does not exists or does not have an
 *         associated expire, -1 is returned. In Redis 2.8 or newer, if the Key does not have an
 *         associated expire, -1 is returned or if the Key does not exists, -2 is returned.
 */
func (r *Redis) Ttl(key string) (int64, error) {
	err := r.Client.Ttl(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Pttl Like TTL this command returns the remaining time to live of a key that has an expire set,
// with the sole difference that TTL returns the amount of remaining time in seconds while PTTL returns it in milliseconds.
//In Redis 2.6 or older the command returns -1 if the key does not exist or if the key exist but has no associated expire.
//Starting with Redis 2.8 the return value in case of error changed:
//The command returns -2 if the key does not exist.
//The command returns -1 if the key exists but has no associated expire.
//
//Integer reply: TTL in milliseconds, or a negative value in order to signal an error (see the description above).
func (r *Redis) Pttl(key string) (int64, error) {
	err := r.Client.Pttl(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Setrange Overwrites part of the string stored at key, starting at the specified offset,
// for the entire length of value. If the offset is larger than the current length of the string at key,
// the string is padded with zero-bytes to make offset fit. Non-existing keys are considered as empty strings,
// so this command will make sure it holds a string large enough to be able to set value at offset.
// Note that the maximum offset that you can set is 229 -1 (536870911), as Redis Strings are limited to 512 megabytes.
// If you need to grow beyond this size, you can use multiple keys.
//
// Warning: When setting the last possible byte and the string value stored at key does not yet hold a string value,
// or holds a small string value, Redis needs to allocate all intermediate memory which can block the server for some time.
// On a 2010 MacBook Pro, setting byte number 536870911 (512MB allocation) takes ~300ms,
// setting byte number 134217728 (128MB allocation) takes ~80ms,
// setting bit number 33554432 (32MB allocation) takes ~30ms and setting bit number 8388608 (8MB allocation) takes ~8ms.
// Note that once this first allocation is done,
// subsequent calls to SETRANGE for the same key will not have the allocation overhead.
//
// Patterns
// Thanks to SETRANGE and the analogous GETRANGE commands, you can use Redis strings as a linear array with O(1) random access. This is a very fast and efficient storage in many real world use cases.
//
// Return value
// Integer reply: the length of the string after it was modified by the command.
func (r *Redis) Setrange(key string, offset int64, value string) (int64, error) {
	err := r.Client.Setrange(key, offset, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Getrange Warning: this command was renamed to GETRANGE, it is called SUBSTR in Redis versions <= 2.0.
// Returns the substring of the string value stored at key,
// determined by the offsets start and end (both are inclusive).
// Negative offsets can be used in order to provide an offset starting from the end of the string.
// So -1 means the last character, -2 the penultimate and so forth.
//
// The function handles out of range requests by limiting the resulting range to the actual length of the string.
//
// Return value
// Bulk string reply
func (r *Redis) Getrange(key string, startOffset, endOffset int64) (string, error) {
	err := r.Client.Getrange(key, startOffset, endOffset)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * GETSET is an atomic set this value and return the old value command. Set key to the string
 * value and return the old value stored at key. The string can't be longer than 1073741824 bytes
 * (1 GB).
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param value
 * @return Bulk reply
 */
func (r *Redis) GetSet(key, value string) (string, error) {
	err := r.Client.GetSet(key, value)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * SETNX works exactly like {@link #set(String, String) SET} with the only difference that if the
 * key already exists no operation is performed. SETNX actually means "SET if Not eXists".
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param value
 * @return Integer reply, specifically: 1 if the key was set 0 if the key was not set
 */
func (r *Redis) Setnx(key, value string) (int64, error) {
	err := r.Client.Setnx(key, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * The command is exactly equivalent to the following group of commands:
 * {@link #set(String, String) SET} + {@link #expire(String, int) EXPIRE}. The operation is
 * atomic.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param seconds
 * @param value
 * @return Status code reply
 */
func (r *Redis) Setex(key string, seconds int, value string) (string, error) {
	err := r.Client.Setex(key, seconds, value)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * IDECRBY work just like {@link #decr(String) INCR} but instead to decrement by 1 the decrement
 * is integer.
 * <p>
 * INCR commands are limited to 64 bit signed integers.
 * <p>
 * Note: this is actually a string operation, that is, in Redis there are not "integer" types.
 * Simply the string stored at the key is parsed as a base 10 64 bit signed integer, incremented,
 * and then converted back as a string.
 * <p>
 * Time complexity: O(1)
 * @see #incr(String)
 * @see #decr(String)
 * @see #incrBy(String, long)
 * @param key
 * @param integer
 * @return Integer reply, this commands will reply with the new value of key after the increment.
 */
func (r *Redis) DecrBy(key string, decrement int64) (int64, error) {
	err := r.Client.DecrBy(key, decrement)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Decrement the number stored at key by one. If the key does not exist or contains a value of a
 * wrong type, set the key to the value of "0" before to perform the decrement operation.
 * <p>
 * INCR commands are limited to 64 bit signed integers.
 * <p>
 * Note: this is actually a string operation, that is, in Redis there are not "integer" types.
 * Simply the string stored at the key is parsed as a base 10 64 bit signed integer, incremented,
 * and then converted back as a string.
 * <p>
 * Time complexity: O(1)
 * @see #incr(String)
 * @see #incrBy(String, long)
 * @see #decrBy(String, long)
 * @param key
 * @return Integer reply, this commands will reply with the new value of key after the increment.
 */
func (r *Redis) Decr(key string) (int64, error) {
	err := r.Client.Decr(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * INCRBY work just like {@link #incr(String) INCR} but instead to increment by 1 the increment is
 * integer.
 * <p>
 * INCR commands are limited to 64 bit signed integers.
 * <p>
 * Note: this is actually a string operation, that is, in Redis there are not "integer" types.
 * Simply the string stored at the key is parsed as a base 10 64 bit signed integer, incremented,
 * and then converted back as a string.
 * <p>
 * Time complexity: O(1)
 * @see #incr(String)
 * @see #decr(String)
 * @see #decrBy(String, long)
 * @param key
 * @param integer
 * @return Integer reply, this commands will reply with the new value of key after the increment.
 */
func (r *Redis) IncrBy(key string, increment int64) (int64, error) {
	err := r.Client.IncrBy(key, increment)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * INCRBYFLOAT
 * <p>
 * INCRBYFLOAT commands are limited to double precision floating point values.
 * <p>
 * Note: this is actually a string operation, that is, in Redis there are not "double" types.
 * Simply the string stored at the key is parsed as a base double precision floating point value,
 * incremented, and then converted back as a string. There is no DECRYBYFLOAT but providing a
 * negative value will work as expected.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param value
 * @return Double reply, this commands will reply with the new value of key after the increment.
 */
func (r *Redis) IncrByFloat(key string, increment float64) (float64, error) {
	err := r.Client.IncrByFloat(key, increment)
	if err != nil {
		return 0, err
	}
	return StringToFloat64Reply(r.Client.getBulkReply())
}

/**
 * Increment the number stored at key by one. If the key does not exist or contains a value of a
 * wrong type, set the key to the value of "0" before to perform the increment operation.
 * <p>
 * INCR commands are limited to 64 bit signed integers.
 * <p>
 * Note: this is actually a string operation, that is, in Redis there are not "integer" types.
 * Simply the string stored at the key is parsed as a base 10 64 bit signed integer, incremented,
 * and then converted back as a string.
 * <p>
 * Time complexity: O(1)
 * @see #incrBy(String, long)
 * @see #decr(String)
 * @see #decrBy(String, long)
 * @param key
 * @return Integer reply, this commands will reply with the new value of key after the increment.
 */
func (r *Redis) Incr(key string) (int64, error) {
	err := r.Client.Incr(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * If the key already exists and is a string, this command appends the provided value at the end
 * of the string. If the key does not exist it is created and set as an empty string, so APPEND
 * will be very similar to SET in this special case.
 * <p>
 * Time complexity: O(1). The amortized time complexity is O(1) assuming the appended value is
 * small and the already present value is of any size, since the dynamic string library used by
 * Redis will double the free space available on every reallocation.
 * @param key
 * @param value
 * @return Integer reply, specifically the total length of the string after the append operation.
 */
func (r *Redis) Append(key, value string) (int64, error) {
	err := r.Client.Append(key, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return a subset of the string from offset start to offset end (both offsets are inclusive).
 * Negative offsets can be used in order to provide an offset starting from the end of the string.
 * So -1 means the last char, -2 the penultimate and so forth.
 * <p>
 * The function handles out of range requests without raising an error, but just limiting the
 * resulting range to the actual length of the string.
 * <p>
 * Time complexity: O(start+n) (with start being the start index and n the total length of the
 * requested range). Note that the lookup part of this command is O(1) so for small strings this
 * is actually an O(1) command.
 * @param key
 * @param start
 * @param end
 * @return Bulk reply
 */
func (r *Redis) Substr(key string, start, end int) (string, error) {
	err := r.Client.Substr(key, start, end)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Set the specified hash field to the specified value.
 * <p>
 * If key does not exist, a new key holding a hash is created.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param field
 * @param value
 * @return If the field already exists, and the HSET just produced an update of the value, 0 is
 *         returned, otherwise if a new field is created 1 is returned.
 */
func (r *Redis) Hset(key, field, value string) (int64, error) {
	err := r.Client.Hset(key, field, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * If key holds a hash, retrieve the value associated to the specified field.
 * <p>
 * If the field is not found or the key does not exist, a special 'nil' value is returned.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param field
 * @return Bulk reply
 */
func (r *Redis) Hget(key, field string) (string, error) {
	err := r.Client.Hget(key, field)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Set the specified hash field to the specified value if the field not exists. <b>Time
 * complexity:</b> O(1)
 * @param key
 * @param field
 * @param value
 * @return If the field already exists, 0 is returned, otherwise if a new field is created 1 is
 *         returned.
 */
func (r *Redis) Hsetnx(key, field, value string) (int64, error) {
	err := r.Client.Hsetnx(key, field, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Set the respective fields to the respective values. HMSET replaces old values with new values.
 * <p>
 * If key does not exist, a new key holding a hash is created.
 * <p>
 * <b>Time complexity:</b> O(N) (with N being the number of fields)
 * @param key
 * @param hash
 * @return Return OK or Exception if hash is empty
 */
func (r *Redis) Hmset(key string, hash map[string]string) (string, error) {
	err := r.Client.Hmset(key, hash)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Retrieve the values associated to the specified fields.
 * <p>
 * If some of the specified fields do not exist, nil values are returned. Non existing keys are
 * considered like empty hashes.
 * <p>
 * <b>Time complexity:</b> O(N) (with N being the number of fields)
 * @param key
 * @param fields
 * @return Multi Bulk Reply specifically a list of all the values associated with the specified
 *         fields, in the same order of the request.
 */
func (r *Redis) Hmget(key string, fields ...string) ([]string, error) {
	err := r.Client.Hmget(key, fields...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Increment the number stored at field in the hash at key by value. If key does not exist, a new
 * key holding a hash is created. If field does not exist or holds a string, the value is set to 0
 * before applying the operation. Since the value argument is signed you can use this command to
 * perform both increments and decrements.
 * <p>
 * The range of values supported by HINCRBY is limited to 64 bit signed integers.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param field
 * @param value
 * @return Integer reply The new value at field after the increment operation.
 */
func (r *Redis) HincrBy(key, field string, value int64) (int64, error) {
	err := r.Client.HincrBy(key, field, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Increment the number stored at field in the hash at key by a double precision floating point
 * value. If key does not exist, a new key holding a hash is created. If field does not exist or
 * holds a string, the value is set to 0 before applying the operation. Since the value argument
 * is signed you can use this command to perform both increments and decrements.
 * <p>
 * The range of values supported by HINCRBYFLOAT is limited to double precision floating point
 * values.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param field
 * @param value
 * @return Double precision floating point reply The new value at field after the increment
 *         operation.
 */
func (r *Redis) HincrByFloat(key, field string, value float64) (float64, error) {
	err := r.Client.HincrByFloat(key, field, value)
	if err != nil {
		return 0, err
	}
	return StringToFloat64Reply(r.Client.getBulkReply())
}

/**
 * Test for existence of a specified field in a hash. <b>Time complexity:</b> O(1)
 * @param key
 * @param field
 * @return Return 1 if the hash stored at key contains the specified field. Return 0 if the key is
 *         not found or the field is not present.
 */
func (r *Redis) Hexists(key, field string) (bool, error) {
	err := r.Client.Hexists(key, field)
	if err != nil {
		return false, err
	}
	return Int64ToBoolReply(r.Client.getIntegerReply())
}

/**
 * Remove the specified field from an hash stored at key.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param fields
 * @return If the field was present in the hash it is deleted and 1 is returned, otherwise 0 is
 *         returned and no operation is performed.
 */
func (r *Redis) Hdel(key string, fields ...string) (int64, error) {
	err := r.Client.Hdel(key, fields...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the number of items in a hash.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @return The number of entries (fields) contained in the hash stored at key. If the specified
 *         key does not exist, 0 is returned assuming an empty hash.
 */
func (r *Redis) Hlen(key string) (int64, error) {
	err := r.Client.Hlen(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return all the fields in a hash.
 * <p>
 * <b>Time complexity:</b> O(N), where N is the total number of entries
 * @param key
 * @return All the fields names contained into a hash.
 */
func (r *Redis) Hkeys(key string) ([]string, error) {
	err := r.Client.Hkeys(key)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Return all the values in a hash.
 * <p>
 * <b>Time complexity:</b> O(N), where N is the total number of entries
 * @param key
 * @return All the fields values contained into a hash.
 */
func (r *Redis) Hvals(key string) ([]string, error) {
	err := r.Client.Hvals(key)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Return all the fields and associated values in a hash.
 * <p>
 * <b>Time complexity:</b> O(N), where N is the total number of entries
 * @param key
 * @return All the fields and values contained into a hash.
 */
func (r *Redis) HgetAll(key string) (map[string]string, error) {
	err := r.Client.HgetAll(key)
	if err != nil {
		return nil, err
	}
	return StringArrayToMapReply(r.Client.getMultiBulkReply())
}

/**
 * Add the string value to the head (LPUSH) or tail (RPUSH) of the list stored at key. If the key
 * does not exist an empty list is created just before the append operation. If the key exists but
 * is not a List an error is returned.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param strings
 * @return Integer reply, specifically, the number of elements inside the list after the push
 *         operation.
 */
func (r *Redis) Rpush(key string, strings ...string) (int64, error) {
	err := r.Client.Rpush(key, strings...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Add the string value to the head (LPUSH) or tail (RPUSH) of the list stored at key. If the key
 * does not exist an empty list is created just before the append operation. If the key exists but
 * is not a List an error is returned.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @param strings
 * @return Integer reply, specifically, the number of elements inside the list after the push
 *         operation.
 */
func (r *Redis) Lpush(key string, strings ...string) (int64, error) {
	err := r.Client.Lpush(key, strings...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the length of the list stored at the specified key. If the key does not exist zero is
 * returned (the same behaviour as for empty lists). If the value stored at key is not a list an
 * error is returned.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @return The length of the list.
 */
func (r *Redis) Llen(key string) (int64, error) {
	err := r.Client.Llen(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the specified elements of the list stored at the specified key. Start and end are
 * zero-based indexes. 0 is the first element of the list (the list head), 1 the next element and
 * so on.
 * <p>
 * For example LRANGE foobar 0 2 will return the first three elements of the list.
 * <p>
 * start and end can also be negative numbers indicating offsets from the end of the list. For
 * example -1 is the last element of the list, -2 the penultimate element and so on.
 * <p>
 * <b>Consistency with range functions in various programming languages</b>
 * <p>
 * Note that if you have a list of numbers from 0 to 100, LRANGE 0 10 will return 11 elements,
 * that is, rightmost item is included. This may or may not be consistent with behavior of
 * range-related functions in your programming language of choice (think Ruby's Range.new,
 * Array#slice or Python's range() function).
 * <p>
 * LRANGE behavior is consistent with one of Tcl.
 * <p>
 * <b>Out-of-range indexes</b>
 * <p>
 * Indexes out of range will not produce an error: if start is over the end of the list, or start
 * &gt; end, an empty list is returned. If end is over the end of the list Redis will threat it
 * just like the last element of the list.
 * <p>
 * Time complexity: O(start+n) (with n being the length of the range and start being the start
 * offset)
 * @param key
 * @param start
 * @param end
 * @return Multi bulk reply, specifically a list of elements in the specified range.
 */
func (r *Redis) Lrange(key string, start, stop int64) ([]string, error) {
	err := r.Client.Lrange(key, start, stop)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Trim an existing list so that it will contain only the specified range of elements specified.
 * Start and end are zero-based indexes. 0 is the first element of the list (the list head), 1 the
 * next element and so on.
 * <p>
 * For example LTRIM foobar 0 2 will modify the list stored at foobar key so that only the first
 * three elements of the list will remain.
 * <p>
 * start and end can also be negative numbers indicating offsets from the end of the list. For
 * example -1 is the last element of the list, -2 the penultimate element and so on.
 * <p>
 * Indexes out of range will not produce an error: if start is over the end of the list, or start
 * &gt; end, an empty list is left as value. If end over the end of the list Redis will threat it
 * just like the last element of the list.
 * <p>
 * Hint: the obvious use of LTRIM is together with LPUSH/RPUSH. For example:
 * <p>
 * {@code lpush("mylist", "someelement"); ltrim("mylist", 0, 99); * }
 * <p>
 * The above two commands will push elements in the list taking care that the list will not grow
 * without limits. This is very useful when using Redis to store logs for example. It is important
 * to note that when used in this way LTRIM is an O(1) operation because in the average case just
 * one element is removed from the tail of the list.
 * <p>
 * Time complexity: O(n) (with n being len of list - len of range)
 * @param key
 * @param start
 * @param end
 * @return Status code reply
 */
func (r *Redis) Ltrim(key string, start, stop int64) (string, error) {
	err := r.Client.Ltrim(key, start, stop)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Return the specified element of the list stored at the specified key. 0 is the first element, 1
 * the second and so on. Negative indexes are supported, for example -1 is the last element, -2
 * the penultimate and so on.
 * <p>
 * If the value stored at key is not of list type an error is returned. If the index is out of
 * range a 'nil' reply is returned.
 * <p>
 * Note that even if the average time complexity is O(n) asking for the first or the last element
 * of the list is O(1).
 * <p>
 * Time complexity: O(n) (with n being the length of the list)
 * @param key
 * @param index
 * @return Bulk reply, specifically the requested element
 */
func (r *Redis) Lindex(key string, index int64) (string, error) {
	err := r.Client.Lindex(key, index)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Set a new value as the element at index position of the List at key.
 * <p>
 * Out of range indexes will generate an error.
 * <p>
 * Similarly to other list commands accepting indexes, the index can be negative to access
 * elements starting from the end of the list. So -1 is the last element, -2 is the penultimate,
 * and so forth.
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(N) (with N being the length of the list), setting the first or last elements of the list is
 * O(1).
 * @see #lindex(String, long)
 * @param key
 * @param index
 * @param value
 * @return Status code reply
 */
func (r *Redis) Lset(key string, index int64, value string) (string, error) {
	err := r.Client.Lset(key, index, value)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Remove the first count occurrences of the value element from the list. If count is zero all the
 * elements are removed. If count is negative elements are removed from tail to head, instead to
 * go from head to tail that is the normal behaviour. So for example LREM with count -2 and hello
 * as value to remove against the list (a,b,c,hello,x,hello,hello) will lave the list
 * (a,b,c,hello,x). The number of removed elements is returned as an integer, see below for more
 * information about the returned value. Note that non existing keys are considered like empty
 * lists by LREM, so LREM against non existing keys will always return 0.
 * <p>
 * Time complexity: O(N) (with N being the length of the list)
 * @param key
 * @param count
 * @param value
 * @return Integer Reply, specifically: The number of removed elements if the operation succeeded
 */
func (r *Redis) Lrem(key string, count int64, value string) (int64, error) {
	err := r.Client.Lrem(key, count, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Atomically return and remove the first (LPOP) or last (RPOP) element of the list. For example
 * if the list contains the elements "a","b","c" LPOP will return "a" and the list will become
 * "b","c".
 * <p>
 * If the key does not exist or the list is already empty the special value 'nil' is returned.
 * @see #rpop(String)
 * @param key
 * @return Bulk reply
 */
func (r *Redis) Lpop(key string) (string, error) {
	err := r.Client.Lpop(key)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Atomically return and remove the first (LPOP) or last (RPOP) element of the list. For example
 * if the list contains the elements "a","b","c" RPOP will return "c" and the list will become
 * "a","b".
 * <p>
 * If the key does not exist or the list is already empty the special value 'nil' is returned.
 * @see #lpop(String)
 * @param key
 * @return Bulk reply
 */
func (r *Redis) Rpop(key string) (string, error) {
	err := r.Client.Rpop(key)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Add the specified member to the set value stored at key. If member is already a member of the
 * set no operation is performed. If key does not exist a new set with the specified member as
 * sole member is created. If the key exists but does not hold a set value an error is returned.
 * <p>
 * Time complexity O(1)
 * @param key
 * @param members
 * @return Integer reply, specifically: 1 if the new element was added 0 if the element was
 *         already a member of the set
 */
func (r *Redis) Sadd(key string, members ...string) (int64, error) {
	err := r.Client.Sadd(key, members...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return all the members (elements) of the set value stored at key. This is just syntax glue for
 * {@link #sinter(String...) SINTER}.
 * <p>
 * Time complexity O(N)
 * @param key
 * @return Multi bulk reply
 */
func (r *Redis) Smembers(key string) ([]string, error) {
	err := r.Client.Smembers(key)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Remove the specified member from the set value stored at key. If member was not a member of the
 * set no operation is performed. If key does not hold a set value an error is returned.
 * <p>
 * Time complexity O(1)
 * @param key
 * @param members
 * @return Integer reply, specifically: 1 if the new element was removed 0 if the new element was
 *         not a member of the set
 */
func (r *Redis) Srem(key string, members ...string) (int64, error) {
	err := r.Client.Srem(key, members...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Remove a random element from a Set returning it as return value. If the Set is empty or the key
 * does not exist, a nil object is returned.
 * <p>
 * The {@link #srandmember(String)} command does a similar work but the returned element is not
 * removed from the Set.
 * <p>
 * Time complexity O(1)
 * @param key
 * @return Bulk reply
 */
func (r *Redis) Spop(key string) (string, error) {
	err := r.Client.Spop(key)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

func (r *Redis) SpopBatch(key string, count int64) ([]string, error) {
	err := r.Client.SpopBatch(key, count)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Return the set cardinality (number of elements). If the key does not exist 0 is returned, like
 * for empty sets.
 * @param key
 * @return Integer reply, specifically: the cardinality (number of elements) of the set as an
 *         integer.
 */
func (r *Redis) Scard(key string) (int64, error) {
	err := r.Client.Scard(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return 1 if member is a member of the set stored at key, otherwise 0 is returned.
 * <p>
 * Time complexity O(1)
 * @param key
 * @param member
 * @return Integer reply, specifically: 1 if the element is a member of the set 0 if the element
 *         is not a member of the set OR if the key does not exist
 */
func (r *Redis) Sismember(key, member string) (bool, error) {
	err := r.Client.Sismember(key, member)
	if err != nil {
		return false, err
	}
	return Int64ToBoolReply(r.Client.getIntegerReply())
}

/**
 * Return the members of a set resulting from the intersection of all the sets hold at the
 * specified keys. Like in {@link #lrange(String, long, long) LRANGE} the result is sent to the
 * client as a multi-bulk reply (see the protocol specification for more information). If just a
 * single key is specified, then this command produces the same result as
 * {@link #smembers(String) SMEMBERS}. Actually SMEMBERS is just syntax sugar for SINTER.
 * <p>
 * Non existing keys are considered like empty sets, so if one of the keys is missing an empty set
 * is returned (since the intersection with an empty set always is an empty set).
 * <p>
 * Time complexity O(N*M) worst case where N is the cardinality of the smallest set and M the
 * number of sets
 * @param keys
 * @return Multi bulk reply, specifically the list of common elements.
 */
func (r *Redis) Sinter(keys ...string) ([]string, error) {
	err := r.Client.Sinter(keys...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * This commnad works exactly like {@link #sinter(String...) SINTER} but instead of being returned
 * the resulting set is sotred as dstkey.
 * <p>
 * Time complexity O(N*M) worst case where N is the cardinality of the smallest set and M the
 * number of sets
 * @param dstkey
 * @param keys
 * @return Status code reply
 */
func (r *Redis) Sinterstore(dstkey string, keys ...string) (int64, error) {
	err := r.Client.Sinterstore(dstkey, keys...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the members of a set resulting from the union of all the sets hold at the specified
 * keys. Like in {@link #lrange(String, long, long) LRANGE} the result is sent to the client as a
 * multi-bulk reply (see the protocol specification for more information). If just a single key is
 * specified, then this command produces the same result as {@link #smembers(String) SMEMBERS}.
 * <p>
 * Non existing keys are considered like empty sets.
 * <p>
 * Time complexity O(N) where N is the total number of elements in all the provided sets
 * @param keys
 * @return Multi bulk reply, specifically the list of common elements.
 */
func (r *Redis) Sunion(keys ...string) ([]string, error) {
	err := r.Client.Sunion(keys...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * This command works exactly like {@link #sunion(String...) SUNION} but instead of being returned
 * the resulting set is stored as dstkey. Any existing value in dstkey will be over-written.
 * <p>
 * Time complexity O(N) where N is the total number of elements in all the provided sets
 * @param dstkey
 * @param keys
 * @return Status code reply
 */
func (r *Redis) Sunionstore(dstkey string, keys ...string) (int64, error) {
	err := r.Client.Sunionstore(dstkey, keys...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the difference between the Set stored at key1 and all the Sets key2, ..., keyN
 * <p>
 * <b>Example:</b>
 *
 * <pre>
 * key1 = [x, a, b, c]
 * key2 = [c]
 * key3 = [a, d]
 * SDIFF key1,key2,key3 =&gt; [x, b]
 * </pre>
 *
 * Non existing keys are considered like empty sets.
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(N) with N being the total number of elements of all the sets
 * @param keys
 * @return Return the members of a set resulting from the difference between the first set
 *         provided and all the successive sets.
 */
func (r *Redis) Sdiff(keys ...string) ([]string, error) {
	err := r.Client.Sdiff(keys...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * This command works exactly like {@link #sdiff(String...) SDIFF} but instead of being returned
 * the resulting set is stored in dstkey.
 * @param dstkey
 * @param keys
 * @return Status code reply
 */
func (r *Redis) Sdiffstore(dstkey string, keys ...string) (int64, error) {
	err := r.Client.Sdiffstore(dstkey, keys...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return a random element from a Set, without removing the element. If the Set is empty or the
 * key does not exist, a nil object is returned.
 * <p>
 * The SPOP command does a similar work but the returned element is popped (removed) from the Set.
 * <p>
 * Time complexity O(1)
 * @param key
 * @return Bulk reply
 */
func (r *Redis) Srandmember(key string) (string, error) {
	err := r.Client.Srandmember(key)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Add the specified member having the specifeid score to the sorted set stored at key. If member
 * is already a member of the sorted set the score is updated, and the element reinserted in the
 * right position to ensure sorting. If key does not exist a new sorted set with the specified
 * member as sole member is crated. If the key exists but does not hold a sorted set value an
 * error is returned.
 * <p>
 * The score value can be the string representation of a double precision floating point number.
 * <p>
 * Time complexity O(log(N)) with N being the number of elements in the sorted set
 * @param key
 * @param score
 * @param member
 * @return Integer reply, specifically: 1 if the new element was added 0 if the element was
 *         already a member of the sorted set and the score was updated
 */
func (r *Redis) Zadd(key string, score float64, member string, mparams ...ZAddParams) (int64, error) {
	err := r.Client.Zadd(key, score, member)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Zrange(key string, start, stop int64) ([]string, error) {
	err := r.Client.Zrange(key, start, stop)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Remove the specified member from the sorted set value stored at key. If member was not a member
 * of the set no operation is performed. If key does not not hold a set value an error is
 * returned.
 * <p>
 * Time complexity O(log(N)) with N being the number of elements in the sorted set
 * @param key
 * @param members
 * @return Integer reply, specifically: 1 if the new element was removed 0 if the new element was
 *         not a member of the set
 */
func (r *Redis) Zrem(key string, members ...string) (int64, error) {
	err := r.Client.Zrem(key, members...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * If member already exists in the sorted set adds the increment to its score and updates the
 * position of the element in the sorted set accordingly. If member does not already exist in the
 * sorted set it is added with increment as score (that is, like if the previous score was
 * virtually zero). If key does not exist a new sorted set with the specified member as sole
 * member is crated. If the key exists but does not hold a sorted set value an error is returned.
 * <p>
 * The score value can be the string representation of a double precision floating point number.
 * It's possible to provide a negative value to perform a decrement.
 * <p>
 * For an introduction to sorted sets check the Introduction to Redis data types page.
 * <p>
 * Time complexity O(log(N)) with N being the number of elements in the sorted set
 * @param key
 * @param score
 * @param member
 * @return The new score
 */
func (r *Redis) Zincrby(key string, increment float64, member string, params ...ZAddParams) (float64, error) {
	err := r.Client.Zincrby(key, increment, member)
	if err != nil {
		return 0, err
	}
	return StringToFloat64Reply(r.Client.getBulkReply())
}

/**
 * Return the rank (or index) or member in the sorted set at key, with scores being ordered from
 * low to high.
 * <p>
 * When the given member does not exist in the sorted set, the special value 'nil' is returned.
 * The returned rank (or index) of the member is 0-based for both commands.
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(log(N))
 * @see #zrevrank(String, String)
 * @param key
 * @param member
 * @return Integer reply or a nil bulk reply, specifically: the rank of the element as an integer
 *         reply if the element exists. A nil bulk reply if there is no such element.
 */
func (r *Redis) Zrank(key, member string) (int64, error) {
	err := r.Client.Zrank(key, member)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the rank (or index) or member in the sorted set at key, with scores being ordered from
 * high to low.
 * <p>
 * When the given member does not exist in the sorted set, the special value 'nil' is returned.
 * The returned rank (or index) of the member is 0-based for both commands.
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(log(N))
 * @see #zrank(String, String)
 * @param key
 * @param member
 * @return Integer reply or a nil bulk reply, specifically: the rank of the element as an integer
 *         reply if the element exists. A nil bulk reply if there is no such element.
 */
func (r *Redis) Zrevrank(key, member string) (int64, error) {
	err := r.Client.Zrevrank(key, member)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Zrevrange(key string, start, stop int64) ([]string, error) {
	err := r.Client.Zrevrange(key, start, stop)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Return the sorted set cardinality (number of elements). If the key does not exist 0 is
 * returned, like for empty sorted sets.
 * <p>
 * Time complexity O(1)
 * @param key
 * @return the cardinality (number of elements) of the set as an integer.
 */
func (r *Redis) Zcard(key string) (int64, error) {
	err := r.Client.Zcard(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the score of the specified element of the sorted set at key. If the specified element
 * does not exist in the sorted set, or the key does not exist at all, a special 'nil' value is
 * returned.
 * <p>
 * <b>Time complexity:</b> O(1)
 * @param key
 * @param member
 * @return the score
 */
func (r *Redis) Zscore(key, member string) (float64, error) {
	err := r.Client.Zscore(key, member)
	if err != nil {
		return 0, err
	}
	return StringToFloat64Reply(r.Client.getBulkReply())
}

func (r *Redis) Watch(keys ...string) (string, error) {
	err := r.Client.Watch(keys...)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

/**
 * Sort a Set or a List.
 * <p>
 * Sort the elements contained in the List, Set, or Sorted Set value at key. By default sorting is
 * numeric with elements being compared as double precision floating point numbers. This is the
 * simplest form of SORT.
 * @see #sort(String, String)
 * @see #sort(String, SortingParams)
 * @see #sort(String, SortingParams, String)
 * @param key
 * @return Assuming the Set/List at key contains a list of numbers, the return value will be the
 *         list of numbers ordered from the smallest to the biggest number.
 */
func (r *Redis) Sort(key string, sortingParameters ...SortingParams) ([]string, error) {
	err := r.Client.Sort(key, sortingParameters...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

// Zcount Returns the number of elements in the sorted set at key with a score between min and max.
// The min and max arguments have the same semantic as described for ZRANGEBYSCORE.
// Note: the command has a complexity of just O(log(N))
// because it uses elements ranks (see ZRANK) to get an idea of the range.
// Because of this there is no need to do a work proportional to the size of the range.
//
// Return value
// Integer reply: the number of elements in the specified score range.
func (r *Redis) Zcount(key, min, max string) (int64, error) {
	err := r.Client.Zcount(key, min, max)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Return the all the elements in the sorted set at key with a score between min and max
 * (including elements with score equal to min or max).
 * <p>
 * The elements having the same score are returned sorted lexicographically as ASCII strings (this
 * follows from a property of Redis sorted sets and does not involve further computation).
 * <p>
 * Using the optional {@link #zrangeByScore(String, double, double, int, int) LIMIT} it's possible
 * to get only a range of the matching elements in an SQL-alike way. Note that if offset is large
 * the commands needs to traverse the list for offset elements and this adds up to the O(M)
 * figure.
 * <p>
 * The {@link #zcount(String, double, double) ZCOUNT} command is similar to
 * {@link #zrangeByScore(String, double, double) ZRANGEBYSCORE} but instead of returning the
 * actual elements in the specified interval, it just returns the number of matching elements.
 * <p>
 * <b>Exclusive intervals and infinity</b>
 * <p>
 * min and max can be -inf and +inf, so that you are not required to know what's the greatest or
 * smallest element in order to take, for instance, elements "up to a given value".
 * <p>
 * Also while the interval is for default closed (inclusive) it's possible to specify open
 * intervals prefixing the score with a "(" character, so for instance:
 * <p>
 * {@code ZRANGEBYSCORE zset (1.3 5}
 * <p>
 * Will return all the values with score &gt; 1.3 and &lt;= 5, while for instance:
 * <p>
 * {@code ZRANGEBYSCORE zset (5 (10}
 * <p>
 * Will return all the values with score &gt; 5 and &lt; 10 (5 and 10 excluded).
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(log(N))+O(M) with N being the number of elements in the sorted set and M the number of
 * elements returned by the command, so if M is constant (for instance you always ask for the
 * first ten elements with LIMIT) you can consider it O(log(N))
 * @see #zrangeByScore(String, double, double)
 * @see #zrangeByScore(String, double, double, int, int)
 * @see #zrangeByScoreWithScores(String, double, double)
 * @see #zrangeByScoreWithScores(String, String, String)
 * @see #zrangeByScoreWithScores(String, double, double, int, int)
 * @see #zcount(String, double, double)
 * @param key
 * @param min a double or Double.MIN_VALUE for "-inf"
 * @param max a double or Double.MAX_VALUE for "+inf"
 * @return Multi bulk reply specifically a list of elements in the specified score range.
 */
func (r *Redis) ZrangeByScore(key, min, max string) ([]string, error) {
	err := r.Client.ZrangeByScore(key, min, max)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Return the all the elements in the sorted set at key with a score between min and max
 * (including elements with score equal to min or max).
 * <p>
 * The elements having the same score are returned sorted lexicographically as ASCII strings (this
 * follows from a property of Redis sorted sets and does not involve further computation).
 * <p>
 * Using the optional {@link #zrangeByScore(String, double, double, int, int) LIMIT} it's possible
 * to get only a range of the matching elements in an SQL-alike way. Note that if offset is large
 * the commands needs to traverse the list for offset elements and this adds up to the O(M)
 * figure.
 * <p>
 * The {@link #zcount(String, double, double) ZCOUNT} command is similar to
 * {@link #zrangeByScore(String, double, double) ZRANGEBYSCORE} but instead of returning the
 * actual elements in the specified interval, it just returns the number of matching elements.
 * <p>
 * <b>Exclusive intervals and infinity</b>
 * <p>
 * min and max can be -inf and +inf, so that you are not required to know what's the greatest or
 * smallest element in order to take, for instance, elements "up to a given value".
 * <p>
 * Also while the interval is for default closed (inclusive) it's possible to specify open
 * intervals prefixing the score with a "(" character, so for instance:
 * <p>
 * {@code ZRANGEBYSCORE zset (1.3 5}
 * <p>
 * Will return all the values with score &gt; 1.3 and &lt;= 5, while for instance:
 * <p>
 * {@code ZRANGEBYSCORE zset (5 (10}
 * <p>
 * Will return all the values with score &gt; 5 and &lt; 10 (5 and 10 excluded).
 * <p>
 * <b>Time complexity:</b>
 * <p>
 * O(log(N))+O(M) with N being the number of elements in the sorted set and M the number of
 * elements returned by the command, so if M is constant (for instance you always ask for the
 * first ten elements with LIMIT) you can consider it O(log(N))
 * @see #zrangeByScore(String, double, double)
 * @see #zrangeByScore(String, double, double, int, int)
 * @see #zrangeByScoreWithScores(String, double, double)
 * @see #zrangeByScoreWithScores(String, double, double, int, int)
 * @see #zcount(String, double, double)
 * @param key
 * @param min
 * @param max
 * @return Multi bulk reply specifically a list of elements in the specified score range.
 */
func (r *Redis) ZrangeByScoreWithScores(key, min, max string) ([]Tuple, error) {
	panic("not implement!")
}

// ZrevrangeByScore Returns all the elements in the sorted set at key with a score between max and min
// (including elements with score equal to max or min). In contrary to the default ordering of sorted sets,
// for this command the elements are considered to be ordered from high to low scores.
// The elements having the same score are returned in reverse lexicographical order.
// Apart from the reversed ordering, ZREVRANGEBYSCORE is similar to ZRANGEBYSCORE.
//
// Return value
// Array reply: list of elements in the specified score range (optionally with their scores).
func (r *Redis) ZrevrangeByScore(key, max, min string) ([]string, error) {
	err := r.Client.ZrevrangeByScore(key, max, min)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

func (r *Redis) ZrevrangeByScoreWithScores(key, max, min string) ([]Tuple, error) {
	err := r.Client.ZrevrangeByScore(key, max, min)
	if err != nil {
		return nil, err
	}
	return StringArrToTupleReply(r.Client.getMultiBulkReply())
}

/**
 * Remove all elements in the sorted set at key with rank between start and end. Start and end are
 * 0-based with rank 0 being the element with the lowest score. Both start and end can be negative
 * numbers, where they indicate offsets starting at the element with the highest rank. For
 * example: -1 is the element with the highest score, -2 the element with the second highest score
 * and so forth.
 * <p>
 * <b>Time complexity:</b> O(log(N))+O(M) with N being the number of elements in the sorted set
 * and M the number of elements removed by the operation
 */
func (r *Redis) ZremrangeByRank(key string, start, stop int64) (int64, error) {
	err := r.Client.ZremrangeByRank(key, start, stop)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Strlen Returns the length of the string value stored at key.
// An error is returned when key holds a non-string value.
// Return value
// Integer reply: the length of the string at key, or 0 when key does not exist.
func (r *Redis) Strlen(key string) (int64, error) {
	err := r.Client.Strlen(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Lpushx Inserts value at the head of the list stored at key,
// only if key already exists and holds a list. In contrary to LPUSH,
// no operation will be performed when key does not yet exist.
// Return value
// Integer reply: the length of the list after the push operation.
func (r *Redis) Lpushx(key string, string ...string) (int64, error) {
	err := r.Client.Lpushx(key, string...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Undo a {@link #expire(String, int) expire} at turning the expire key into a normal key.
 * <p>
 * Time complexity: O(1)
 * @param key
 * @return Integer reply, specifically: 1: the key is now persist. 0: the key is not persist (only
 *         happens when key not set).
 */
func (r *Redis) Persist(key string) (int64, error) {
	err := r.Client.Persist(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Rpushx Inserts value at the tail of the list stored at key,
// only if key already exists and holds a list. In contrary to RPUSH,
// no operation will be performed when key does not yet exist.
//
// Return value
// Integer reply: the length of the list after the push operation.
func (r *Redis) Rpushx(key string, string ...string) (int64, error) {
	err := r.Client.Rpushx(key, string...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

// Echo Returns message.
//
// Return value
// Bulk string reply
func (r *Redis) Echo(string string) (string, error) {
	err := r.Client.Echo(string)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

func (r *Redis) SetWithParams(key, value, nxxx string) (string, error) {
	err := r.Client.SetWithParams(key, value, nxxx)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

func (r *Redis) Pexpire(key string, milliseconds int64) (int64, error) {
	err := r.Client.Pexpire(key, milliseconds)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) PexpireAt(key string, millisecondsTimestamp int64) (int64, error) {
	err := r.Client.PexpireAt(key, millisecondsTimestamp)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SetbitWithBool(key string, offset int64, value bool) (bool, error) {
	valueByte := make([]byte, 0)
	if value {
		valueByte = BYTES_TRUE
	} else {
		valueByte = BYTES_FALSE
	}
	return r.Setbit(key, offset, string(valueByte))
}

func (r *Redis) Setbit(key string, offset int64, value string) (bool, error) {
	err := r.Client.Setbit(key, offset, value)
	if err != nil {
		return false, err
	}
	return Int64ToBoolReply(r.Client.getIntegerReply())
}

func (r *Redis) Getbit(key string, offset int64) (bool, error) {
	err := r.Client.Getbit(key, offset)
	if err != nil {
		return false, err
	}
	return Int64ToBoolReply(r.Client.getIntegerReply())
}

func (r *Redis) Psetex(key string, milliseconds int64, value string) (string, error) {
	err := r.Client.Psetex(key, milliseconds, value)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

func (r *Redis) SrandmemberBatch(key string, count int) ([]string, error) {
	err := r.Client.SrandmemberBatch(key, count)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

func (r *Redis) ZaddByMap(key string, scoreMembers map[string]float64, params ...ZAddParams) (int64, error) {
	err := r.Client.ZaddByMap(key, scoreMembers, params...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	err := r.Client.ZrangeWithScores(key, start, end)
	if err != nil {
		return nil, err
	}
	return StringArrToTupleReply(r.Client.getMultiBulkReply())
}

func (r *Redis) ZrevrangeWithScores(key string, start, end int64) ([]Tuple, error) {
	err := r.Client.ZrevrangeWithScores(key, start, end)
	if err != nil {
		return nil, err
	}
	return StringArrToTupleReply(r.Client.getMultiBulkReply())
}

func (r *Redis) ZrangeByScoreBatch(key, min, max string, offset, count int) ([]string, error) {
	err := r.Client.ZrangeByScoreBatch(key, min, max, offset, count)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

func (r *Redis) ZrangeByScoreWithScoresBatch(key, min, max string, offset, count int) ([]Tuple, error) {
	err := r.Client.ZrangeByScoreBatch(key, min, max, offset, count)
	if err != nil {
		return nil, err
	}
	return StringArrToTupleReply(r.Client.getMultiBulkReply())
}

func (r *Redis) ZrevrangeByScoreWithScoresBatch(key, max, min string, offset, count int) ([]Tuple, error) {
	err := r.Client.ZrevrangeByScoreBatch(key, max, min, offset, count)
	if err != nil {
		return nil, err
	}
	return StringArrToTupleReply(r.Client.getMultiBulkReply())
}

func (r *Redis) ZremrangeByScore(key, start, end string) (int64, error) {
	err := r.Client.ZremrangeByScore(key, start, end)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Zlexcount(key, min, max string) (int64, error) {
	err := r.Client.Zlexcount(key, min, max)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZrangeByLex(key, min, max string) ([]string, error) {
	err := r.Client.ZrangeByLex(key, min, max)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

func (r *Redis) ZrangeByLexBatch(key, min, max string, offset, count int) ([]string, error) {
	err := r.Client.ZrangeByLexBatch(key, min, max, offset, count)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZrevrangeByLex(key, max, min string) ([]string, error) {
	err := r.Client.ZrevrangeByLex(key, max, min)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZrevrangeByLexBatch(key, max, min string, offset, count int) ([]string, error) {
	err := r.Client.ZrevrangeByLexBatch(key, max, min, offset, count)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZremrangeByLex(key, min, max string) (int64, error) {
	err := r.Client.ZremrangeByLex(key, min, max)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Linsert(key string, where ListOption, pivot, value string) (int64, error) {
	err := r.Client.Linsert(key, where, pivot, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Move(key string, dbIndex int) (int64, error) {
	err := r.Client.Move(key, dbIndex)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Bitcount(key string) (int64, error) {
	err := r.Client.Bitcount(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) BitcountRange(key string, start, end int64) (int64, error) {
	err := r.Client.BitcountRange(key, start, end)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Bitpos(key string, value bool, params ...BitPosParams) (int64, error) {
	err := r.Client.Bitpos(key, value, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Hscan(key, cursor string, params ...ScanParams) (ScanResult, error) {
	err := r.Client.Hscan(key, cursor, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Sscan(key, cursor string, params ...ScanParams) (ScanResult, error) {
	err := r.Client.Sscan(key, cursor, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Zscan(key, cursor string, params ...ScanParams) (ScanResult, error) {
	err := r.Client.Zscan(key, cursor, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Pfadd(key string, elements ...string) (int64, error) {
	err := r.Client.Pfadd(key, elements)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Geoadd(key string, longitude, latitude float64, member string) (int64, error) {
	err := r.Client.Geoadd(key, longitude, latitude)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) GeoaddByMap(key string, memberCoordinateMap map[string]GeoCoordinate) (int64, error) {
	err := r.Client.GeoaddByMap(key, memberCoordinateMap)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Geodist(key, member1, member2 string, unit ...GeoUnit) (float64, error) {
	err := r.Client.Geodist(key, member1, member2, unit)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Geohash(key string, members ...string) ([]string, error) {
	err := r.Client.Geohash(key, members)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Geopos(key string, members ...string) ([]GeoCoordinate, error) {
	err := r.Client.Geopos(key, members)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Georadius(key string, longitude, latitude, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]GeoCoordinate, error) {
	err := r.Client.Georadius(key, longitude, latitude, radius, unit, param)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) GeoradiusByMember(key, member string, radius float64, unit GeoUnit, param ...GeoRadiusParam) ([]GeoCoordinate, error) {
	err := r.Client.GeoradiusByMember(key, member, radius, unit, param)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Bitfield(key string, arguments ...string) ([]int64, error) {
	err := r.Client.Bitfield(key, arguments)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="multikeycommands">
/**
 * Returns all the keys matching the glob-style pattern as space separated strings. For example if
 * you have in the database the keys "foo" and "foobar" the command "KEYS foo*" will return
 * "foo foobar".
 * <p>
 * Note that while the time complexity for this operation is O(n) the constant times are pretty
 * low. For example Redis running on an entry level laptop can scan a 1 million keys database in
 * 40 milliseconds. <b>Still it's better to consider this one of the slow commands that may ruin
 * the DB performance if not used with care.</b>
 * <p>
 * In other words this command is intended only for debugging and special operations like creating
 * a script to change the DB schema. Don't use it in your normal code. Use Redis Sets in order to
 * group together a subset of objects.
 * <p>
 * Glob style patterns examples:
 * <ul>
 * <li>h?llo will match hello hallo hhllo
 * <li>h*llo will match hllo heeeello
 * <li>h[ae]llo will match hello and hallo, but not hillo
 * </ul>
 * <p>
 * Use \ to escape special chars if you want to match them verbatim.
 * <p>
 * Time complexity: O(n) (with n being the number of keys in the DB, and assuming keys and pattern
 * of limited length)
 * @param pattern
 * @return Multi bulk reply
 */
func (r *Redis) Keys(pattern string) ([]string, error) {
	err := r.Client.Keys(pattern)
	if err != nil {
		return nil, err
	}
	return r.Client.getBinaryMultiBulkReply()
}

/**
 * Remove the specified keys. If a given key does not exist no operation is performed for this
 * key. The command returns the number of keys removed. Time complexity: O(1)
 * @param keys
 * @return Integer reply, specifically: an integer greater than 0 if one or more keys were removed
 *         0 if none of the specified key existed
 */
func (r *Redis) Del(key ...string) (int64, error) {
	err := r.Client.Del(key...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Test if the specified key exists. The command returns the number of keys existed Time
 * complexity: O(N)
 * @param keys
 * @return Integer reply, specifically: an integer greater than 0 if one or more keys were removed
 *         0 if none of the specified key existed
 */
func (r *Redis) Exists(keys ...string) (int64, error) {
	err := r.Client.Exists(keys...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Atomically renames the key oldkey to newkey. If the source and destination name are the same an
 * error is returned. If newkey already exists it is overwritten.
 * <p>
 * Time complexity: O(1)
 * @param oldkey
 * @param newkey
 * @return Status code repy
 */
func (r *Redis) Rename(oldkey, newkey string) (string, error) {
	err := r.Client.Rename(oldkey, newkey)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

/**
 * Rename oldkey into newkey but fails if the destination key newkey already exists.
 * <p>
 * Time complexity: O(1)
 * @param oldkey
 * @param newkey
 * @return Integer reply, specifically: 1 if the key was renamed 0 if the target key already exist
 */
func (r *Redis) Renamenx(oldkey, newkey string) (int64, error) {
	err := r.Client.Renamenx(oldkey, newkey)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Get the values of all the specified keys. If one or more keys dont exist or is not of type
 * String, a 'nil' value is returned instead of the value of the specified key, but the operation
 * never fails.
 * <p>
 * Time complexity: O(1) for every key
 * @param keys
 * @return Multi bulk reply
 */
func (r *Redis) Mget(keys ...string) ([]string, error) {
	err := r.Client.Mget(keys...)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

/**
 * Set the the respective keys to the respective values. MSET will replace old values with new
 * values, while {@link #msetnx(String...) MSETNX} will not perform any operation at all even if
 * just a single key already exists.
 * <p>
 * Because of this semantic MSETNX can be used in order to set different keys representing
 * different fields of an unique logic object in a way that ensures that either all the fields or
 * none at all are set.
 * <p>
 * Both MSET and MSETNX are atomic operations. This means that for instance if the keys A and B
 * are modified, another client talking to Redis can either see the changes to both A and B at
 * once, or no modification at all.
 * @see #msetnx(String...)
 * @param keysvalues
 * @return Status code reply Basically +OK as MSET can't fail
 */
func (r *Redis) Mset(keysvalues ...string) (string, error) {
	err := r.Client.Mset(keysvalues...)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Set the the respective keys to the respective values. {@link #mset(String...) MSET} will
 * replace old values with new values, while MSETNX will not perform any operation at all even if
 * just a single key already exists.
 * <p>
 * Because of this semantic MSETNX can be used in order to set different keys representing
 * different fields of an unique logic object in a way that ensures that either all the fields or
 * none at all are set.
 * <p>
 * Both MSET and MSETNX are atomic operations. This means that for instance if the keys A and B
 * are modified, another client talking to Redis can either see the changes to both A and B at
 * once, or no modification at all.
 * @see #mset(String...)
 * @param keysvalues
 * @return Integer reply, specifically: 1 if the all the keys were set 0 if no key was set (at
 *         least one key already existed)
 */
func (r *Redis) Msetnx(keysvalues ...string) (int64, error) {
	err := r.Client.Msetnx(keysvalues...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Atomically return and remove the last (tail) element of the srckey list, and push the element
 * as the first (head) element of the dstkey list. For example if the source list contains the
 * elements "a","b","c" and the destination list contains the elements "foo","bar" after an
 * RPOPLPUSH command the content of the two lists will be "a","b" and "c","foo","bar".
 * <p>
 * If the key does not exist or the list is already empty the special value 'nil' is returned. If
 * the srckey and dstkey are the same the operation is equivalent to removing the last element
 * from the list and pusing it as first element of the list, so it's a "list rotation" command.
 * <p>
 * Time complexity: O(1)
 * @param srckey
 * @param dstkey
 * @return Bulk reply
 */
func (r *Redis) Rpoplpush(srckey, dstkey string) (string, error) {
	err := r.Client.RpopLpush(srckey, dstkey)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

/**
 * Move the specifided member from the set at srckey to the set at dstkey. This operation is
 * atomic, in every given moment the element will appear to be in the source or destination set
 * for accessing clients.
 * <p>
 * If the source set does not exist or does not contain the specified element no operation is
 * performed and zero is returned, otherwise the element is removed from the source set and added
 * to the destination set. On success one is returned, even if the element was already present in
 * the destination set.
 * <p>
 * An error is raised if the source or destination keys contain a non Set value.
 * <p>
 * Time complexity O(1)
 * @param srckey
 * @param dstkey
 * @param member
 * @return Integer reply, specifically: 1 if the element was moved 0 if the element was not found
 *         on the first set and no operation was performed
 */
func (r *Redis) Smove(srckey, dstkey, member string) (int64, error) {
	err := r.Client.Smove(srckey, dstkey, member)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Creates a union or intersection of N sorted sets given by keys k1 through kN, and stores it at
 * dstkey. It is mandatory to provide the number of input keys N, before passing the input keys
 * and the other (optional) arguments.
 * <p>
 * As the terms imply, the {@link #zinterstore(String, String...) ZINTERSTORE} command requires an
 * element to be present in each of the given inputs to be inserted in the result. The
 * {@link #zunionstore(String, String...) ZUNIONSTORE} command inserts all elements across all
 * inputs.
 * <p>
 * Using the WEIGHTS option, it is possible to add weight to each input sorted set. This means
 * that the score of each element in the sorted set is first multiplied by this weight before
 * being passed to the aggregation. When this option is not given, all weights default to 1.
 * <p>
 * With the AGGREGATE option, it's possible to specify how the results of the union or
 * intersection are aggregated. This option defaults to SUM, where the score of an element is
 * summed across the inputs where it exists. When this option is set to be either MIN or MAX, the
 * resulting set will contain the minimum or maximum score of an element across the inputs where
 * it exists.
 * <p>
 * <b>Time complexity:</b> O(N) + O(M log(M)) with N being the sum of the sizes of the input
 * sorted sets, and M being the number of elements in the resulting sorted set
 * @see #zunionstore(String, String...)
 * @see #zunionstore(String, ZParams, String...)
 * @see #zinterstore(String, String...)
 * @see #zinterstore(String, ZParams, String...)
 * @param dstkey
 * @param sets
 * @return Integer reply, specifically the number of elements in the sorted set at dstkey
 */
func (r *Redis) Zunionstore(dstkey string, sets ...string) (int64, error) {
	err := r.Client.Zunionstore(dstkey, sets...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * Creates a union or intersection of N sorted sets given by keys k1 through kN, and stores it at
 * dstkey. It is mandatory to provide the number of input keys N, before passing the input keys
 * and the other (optional) arguments.
 * <p>
 * As the terms imply, the {@link #zinterstore(String, String...) ZINTERSTORE} command requires an
 * element to be present in each of the given inputs to be inserted in the result. The
 * {@link #zunionstore(String, String...) ZUNIONSTORE} command inserts all elements across all
 * inputs.
 * <p>
 * Using the WEIGHTS option, it is possible to add weight to each input sorted set. This means
 * that the score of each element in the sorted set is first multiplied by this weight before
 * being passed to the aggregation. When this option is not given, all weights default to 1.
 * <p>
 * With the AGGREGATE option, it's possible to specify how the results of the union or
 * intersection are aggregated. This option defaults to SUM, where the score of an element is
 * summed across the inputs where it exists. When this option is set to be either MIN or MAX, the
 * resulting set will contain the minimum or maximum score of an element across the inputs where
 * it exists.
 * <p>
 * <b>Time complexity:</b> O(N) + O(M log(M)) with N being the sum of the sizes of the input
 * sorted sets, and M being the number of elements in the resulting sorted set
 * @see #zunionstore(String, String...)
 * @see #zunionstore(String, ZParams, String...)
 * @see #zinterstore(String, String...)
 * @see #zinterstore(String, ZParams, String...)
 * @param dstkey
 * @param sets
 * @return Integer reply, specifically the number of elements in the sorted set at dstkey
 */
func (r *Redis) Zinterstore(dstkey string, sets ...string) (int64, error) {
	err := r.Client.Zinterstore(dstkey, sets...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) BlpopTimout(timeout int, keys ...string) ([]string, error) {
	err := r.Client.BlpopTimout(timeout, keys)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) BrpopTimout(timeout int, keys ...string) ([]string, error) {
	err := r.Client.BrpopTimout(timeout, keys)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

/**
 * BLPOP (and BRPOP) is a blocking list pop primitive. You can see this commands as blocking
 * versions of LPOP and RPOP able to block if the specified keys don't exist or contain empty
 * lists.
 * <p>
 * The following is a description of the exact semantic. We describe BLPOP but the two commands
 * are identical, the only difference is that BLPOP pops the element from the left (head) of the
 * list, and BRPOP pops from the right (tail).
 * <p>
 * <b>Non blocking behavior</b>
 * <p>
 * When BLPOP is called, if at least one of the specified keys contain a non empty list, an
 * element is popped from the head of the list and returned to the caller together with the name
 * of the key (BLPOP returns a two elements array, the first element is the key, the second the
 * popped value).
 * <p>
 * Keys are scanned from left to right, so for instance if you issue BLPOP list1 list2 list3 0
 * against a dataset where list1 does not exist but list2 and list3 contain non empty lists, BLPOP
 * guarantees to return an element from the list stored at list2 (since it is the first non empty
 * list starting from the left).
 * <p>
 * <b>Blocking behavior</b>
 * <p>
 * If none of the specified keys exist or contain non empty lists, BLPOP blocks until some other
 * client performs a LPUSH or an RPUSH operation against one of the lists.
 * <p>
 * Once new data is present on one of the lists, the client finally returns with the name of the
 * key unblocking it and the popped value.
 * <p>
 * When blocking, if a non-zero timeout is specified, the client will unblock returning a nil
 * special value if the specified amount of seconds passed without a push operation against at
 * least one of the specified keys.
 * <p>
 * The timeout argument is interpreted as an integer value. A timeout of zero means instead to
 * block forever.
 * <p>
 * <b>Multiple clients blocking for the same keys</b>
 * <p>
 * Multiple clients can block for the same key. They are put into a queue, so the first to be
 * served will be the one that started to wait earlier, in a first-blpopping first-served fashion.
 * <p>
 * <b>blocking POP inside a MULTI/EXEC transaction</b>
 * <p>
 * BLPOP and BRPOP can be used with pipelining (sending multiple commands and reading the replies
 * in batch), but it does not make sense to use BLPOP or BRPOP inside a MULTI/EXEC block (a Redis
 * transaction).
 * <p>
 * The behavior of BLPOP inside MULTI/EXEC when the list is empty is to return a multi-bulk nil
 * reply, exactly what happens when the timeout is reached. If you like science fiction, think at
 * it like if inside MULTI/EXEC the time will flow at infinite speed :)
 * <p>
 * Time complexity: O(1)
 * @see #brpop(int, String...)
 * @param timeout
 * @param keys
 * @return BLPOP returns a two-elements array via a multi bulk reply in order to return both the
 *         unblocking key and the popped value.
 *         <p>
 *         When a non-zero timeout is specified, and the BLPOP operation timed out, the return
 *         value is a nil multi bulk reply. Most client values will return false or nil
 *         accordingly to the programming language used.
 */
func (r *Redis) Blpop(args ...string) ([]string, error) {
	err := r.Client.Blpop(args)
	if err != nil {
		return nil, err
	}
	return r.Client.getMultiBulkReply()
}

func (r *Redis) Brpop(args ...string) ([]string, error) {
	err := r.Client.Brpop(args...)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SortMulti(key, dstkey string, sortingParameters ...SortingParams) (int64, error) {
	err := r.Client.SortMulti(key, dstkey, sortingParameters)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Unwatch() (string, error) {
	err := r.Client.Unwatch()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	err := r.Client.ZinterstoreWithParams(dstkey, params, sets)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (int64, error) {
	err := r.Client.ZunionstoreWithParams(dstkey, params, sets)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Brpoplpush(source, destination string, timeout int) (string, error) {
	err := r.Client.Brpoplpush(source, destination, timeout)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Publish(channel, message string) (int64, error) {
	err := r.Client.Publish(channel, message)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Subscribe(redisPubSub RedisPubSub, channels ...string) error {
	err := r.Client.Subscribe(redisPubSub, channels)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Psubscribe(redisPubSub RedisPubSub, patterns ...string) error {
	err := r.Client.Psubscribe(redisPubSub, patterns)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) RandomKey() (string, error) {
	err := r.Client.RandomKey()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Bitop(op BitOP, destKey string, srcKeys ...string) (int64, error) {
	err := r.Client.Bitop(op, destKey, srcKeys)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Scan(cursor string, params ...ScanParams) (ScanResult, error) {
	err := r.Client.Scan(cursor, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Pfmerge(destkey string, sourcekeys ...string) (string, error) {
	err := r.Client.Pfmerge(destkey, sourcekeys)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Pfcount(keys ...string) (int64, error) {
	err := r.Client.Pfcount(keys)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="advancedcommands">
func (r *Redis) ConfigGet(pattern string) ([]string, error) {
	err := r.Client.ConfigGet(pattern)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ConfigSet(parameter, value string) (string, error) {
	err := r.Client.ConfigSet(parameter, value)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SlowlogReset() (string, error) {
	err := r.Client.SlowlogReset()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SlowlogLen() (int64, error) {
	err := r.Client.SlowlogLen()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SlowlogGet(entries ...int64) ([]Slowlog, error) {
	err := r.Client.SlowlogGet(entries)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ObjectRefcount(str string) (int64, error) {
	err := r.Client.ObjectRefcount(str)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ObjectEncoding(str string) (string, error) {
	err := r.Client.ObjectEncoding(str)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ObjectIdletime(str string) (int64, error) {
	err := r.Client.ObjectIdletime(str)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="scriptcommands">
func (r *Redis) Eval(script string, keyCount int, params ...string) (interface{}, error) {
	err := r.Client.Eval(script, keyCount, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Evalsha(sha1 string, keyCount int, params ...string) (interface{}, error) {
	err := r.Client.Evalsha(sha1, keyCount, params)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ScriptExists(sha1 ...string) ([]bool, error) {
	err := r.Client.ScriptExists(sha1)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ScriptLoad(script string) (string, error) {
	err := r.Client.ScriptLoad(script)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="basiccommands">
func (r *Redis) Quit() (string, error) {
	err := r.Client.Quit()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Ping() (string, error) {
	err := r.Client.Ping()
	if err != nil {
		return "", err
	}
	return ByteToStringReply(r.Client.getBinaryBulkReply())
}

func (r *Redis) Select(index int) (string, error) {
	err := r.Client.Select(index)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) FlushDB() (string, error) {
	err := r.Client.FlushDB()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) DbSize() (int64, error) {
	err := r.Client.DbSize()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) FlushAll() (string, error) {
	err := r.Client.FlushAll()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Auth(password string) (string, error) {
	err := r.Client.Auth(password)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Save() (string, error) {
	err := r.Client.Save()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Bgsave() (string, error) {
	err := r.Client.Bgsave()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Bgrewriteaof() (string, error) {
	err := r.Client.Bgrewriteaof()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Lastsave() (int64, error) {
	err := r.Client.Lastsave()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Shutdown() (string, error) {
	err := r.Client.Shutdown()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) Info(section ...string) (string, error) {
	err := r.Client.Info(section...)
	if err != nil {
		return "", err
	}
	return r.Client.getBulkReply()
}

func (r *Redis) Slaveof(host string, port int) (string, error) {
	err := r.Client.Slaveof(host, port)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) SlaveofNoOne() (string, error) {
	err := r.Client.SlaveofNoOne()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) GetDB() int {
	return r.Client.Db
}

func (r *Redis) Debug(params DebugParams) (string, error) {
	err := r.Client.Debug(params)
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) ConfigResetStat() (string, error) {
	err := r.Client.ConfigResetStat()
	if err != nil {
		return "", err
	}
	return r.Client.getStatusCodeReply()
}

func (r *Redis) WaitReplicas(replicas int, timeout int64) (int64, error) {
	err := r.Client.WaitReplicas(replicas, timeout)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="clustercommands">
func (r *Redis) ClusterNodes() (string, error) {
	err := r.Client.ClusterNodes()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterMeet(ip string, port int) (string, error) {
	err := r.Client.ClusterMeet(ip, port)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterAddSlots(slots ...int) (string, error) {
	err := r.Client.ClusterAddSlots(slots)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterDelSlots(slots ...int) (string, error) {
	err := r.Client.ClusterDelSlots(slots)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterInfo() (string, error) {
	err := r.Client.ClusterInfo()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterGetKeysInSlot(slot int, count int) ([]string, error) {
	err := r.Client.ClusterGetKeysInSlot(slot, count)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSetSlotNode(slot int, nodeId string) (string, error) {
	err := r.Client.ClusterSetSlotNode(slot, nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSetSlotMigrating(slot int, nodeId string) (string, error) {
	err := r.Client.ClusterSetSlotMigrating(slot, nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSetSlotImporting(slot int, nodeId string) (string, error) {
	err := r.Client.ClusterSetSlotImporting(slot, nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSetSlotStable(slot int) (string, error) {
	err := r.Client.ClusterSetSlotStable(slot)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterForget(nodeId string) (string, error) {
	err := r.Client.ClusterForget(nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterFlushSlots() (string, error) {
	err := r.Client.ClusterFlushSlots()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterKeySlot(key string) (int64, error) {
	err := r.Client.ClusterKeySlot(key)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterCountKeysInSlot(slot int) (int64, error) {
	err := r.Client.ClusterCountKeysInSlot(slot)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSaveConfig() (string, error) {
	err := r.Client.ClusterSaveConfig()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterReplicate(nodeId string) (string, error) {
	err := r.Client.ClusterReplicate(nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSlaves(nodeId string) ([]string, error) {
	err := r.Client.ClusterSlaves(nodeId)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterFailover() (string, error) {
	err := r.Client.ClusterFailover()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterSlots() ([]interface{}, error) {
	err := r.Client.ClusterSlots()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) ClusterReset(resetType Reset) (string, error) {
	err := r.Client.ClusterReset(resetType)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) Readonly() (string, error) {
	err := r.Client.Readonly()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

//<editor-fold desc="sentinelcommands">
func (r *Redis) SentinelMasters() ([]map[string]string, error) {
	err := r.Client.SentinelMasters()
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelGetMasterAddrByName(masterName string) ([]string, error) {
	err := r.Client.SentinelGetMasterAddrByName(masterName)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelReset(pattern string) (int64, error) {
	err := r.Client.SentinelReset(pattern)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelSlaves(masterName string) ([]map[string]string, error) {
	err := r.Client.SentinelSlaves(masterName)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelFailover(masterName string) (string, error) {
	err := r.Client.SentinelFailover(masterName)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelMonitor(masterName, ip string, port, quorum int) (string, error) {
	err := r.Client.SentinelMonitor(masterName, ip, port, quorum)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelRemove(masterName string) (string, error) {
	err := r.Client.SentinelRemove(masterName)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

func (r *Redis) SentinelSet(masterName string, parameterMap map[string]string) (string, error) {
	err := r.Client.SentinelSet(masterName, parameterMap)
	if err != nil {
		return 0, err
	}
	return r.Client.getIntegerReply()
}

//</editor-fold>

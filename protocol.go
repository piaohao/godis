package godis

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

const (
	ASK_PREFIX         = "ASK "
	MOVED_PREFIX       = "MOVED "
	CLUSTERDOWN_PREFIX = "CLUSTERDOWN "
	BUSY_PREFIX        = "BUSY "
	NOSCRIPT_PREFIX    = "NOSCRIPT "

	DEFAULT_HOST          = "localhost"
	DEFAULT_PORT          = 6379
	DEFAULT_SENTINEL_PORT = 26379
	DEFAULT_TIMEOUT       = 2000
	DEFAULT_DATABASE      = 0

	CHARSET = "UTF-8"

	DOLLAR_BYTE   = '$'
	ASTERISK_BYTE = '*'
	PLUS_BYTE     = '+'
	MINUS_BYTE    = '-'
	COLON_BYTE    = ':'

	SENTINEL_MASTERS                 = "masters"
	SENTINEL_GET_MASTER_ADDR_BY_NAME = "get-master-addr-by-name"
	SENTINEL_RESET                   = "reset"
	SENTINEL_SLAVES                  = "slaves"
	SENTINEL_FAILOVER                = "failover"
	SENTINEL_MONITOR                 = "monitor"
	SENTINEL_REMOVE                  = "remove"
	SENTINEL_SET                     = "set"

	CLUSTER_NODES             = "nodes"
	CLUSTER_MEET              = "meet"
	CLUSTER_RESET             = "reset"
	CLUSTER_ADDSLOTS          = "addslots"
	CLUSTER_DELSLOTS          = "delslots"
	CLUSTER_INFO              = "info"
	CLUSTER_GETKEYSINSLOT     = "getkeysinslot"
	CLUSTER_SETSLOT           = "setslot"
	CLUSTER_SETSLOT_NODE      = "node"
	CLUSTER_SETSLOT_MIGRATING = "migrating"
	CLUSTER_SETSLOT_IMPORTING = "importing"
	CLUSTER_SETSLOT_STABLE    = "stable"
	CLUSTER_FORGET            = "forget"
	CLUSTER_FLUSHSLOT         = "flushslots"
	CLUSTER_KEYSLOT           = "keyslot"
	CLUSTER_COUNTKEYINSLOT    = "countkeysinslot"
	CLUSTER_SAVECONFIG        = "saveconfig"
	CLUSTER_REPLICATE         = "replicate"
	CLUSTER_SLAVES            = "slaves"
	CLUSTER_FAILOVER          = "failover"
	CLUSTER_SLOTS             = "slots"
	PUBSUB_CHANNELS           = "channels"
	PUBSUB_NUMSUB             = "numsub"
	PUBSUB_NUM_PAT            = "numpat"
)

var (
	BYTES_TRUE  = []byte{1}
	BYTES_FALSE = []byte{0}
	BYTES_TILDE = []byte("~")

	POSITIVE_INFINITY_BYTES = []byte("+inf")
	NEGATIVE_INFINITY_BYTES = []byte("-inf")
)

type redisOutputStream struct {
	*bufio.Writer
	buf   []byte
	count int
}

func newRedisOutputStream(bw *bufio.Writer) *redisOutputStream {
	return &redisOutputStream{
		Writer: bw,
		buf:    make([]byte, 0),
	}
}

func (r *redisOutputStream) writeIntCrLf(b int) (int, error) {
	_, err := r.Write(strconv.AppendInt(r.buf, int64(b), 10))
	if err != nil {
		return 0, err
	}
	return r.writeCrLf()
}

func (r *redisOutputStream) writeCrLf() (int, error) {
	return r.WriteString("\r\n")
}

type redisInputStream struct {
	*bufio.Reader
	buf   []byte
	count int
	limit int
}

func newRedisInputStream(br *bufio.Reader) *redisInputStream {
	return &redisInputStream{
		Reader: br,
		buf:    make([]byte, 8192),
	}
}

type protocol struct {
	os *redisOutputStream
	is *redisInputStream
}

func newProtocol(os *redisOutputStream, is *redisInputStream) *protocol {
	return &protocol{
		os: os,
		is: is,
	}
}

func (p *protocol) sendCommand(command []byte, args ...[]byte) error {
	if err := p.os.WriteByte(ASTERISK_BYTE); err != nil {
		return err
	}
	if _, err := p.os.writeIntCrLf(len(args) + 1); err != nil {
		return err
	}
	if err := p.os.WriteByte(DOLLAR_BYTE); err != nil {
		return err
	}
	if _, err := p.os.writeIntCrLf(len(command)); err != nil {
		return err
	}
	if _, err := p.os.Write(command); err != nil {
		return err
	}
	if _, err := p.os.writeCrLf(); err != nil {
		return err
	}
	for _, arg := range args {
		if err := p.os.WriteByte(DOLLAR_BYTE); err != nil {
			return err
		}
		if _, err := p.os.writeIntCrLf(len(arg)); err != nil {
			return err
		}
		if _, err := p.os.Write(arg); err != nil {
			return err
		}
		if _, err := p.os.writeCrLf(); err != nil {
			return err
		}
	}
	return nil
}

func (p *protocol) read() (interface{}, error) {
	return p.process()
}

func (p *protocol) process() (interface{}, error) {
	line, err := p.readLine()
	if err != nil {
		return nil, err
	}
	if len(line) == 0 {
		return nil, errors.New("short response line")
	}
	switch line[0] {
	case PLUS_BYTE:
		switch {
		case len(line) == 3 && line[1] == 'O' && line[2] == 'K':
			// Avoid allocation for frequent "+OK" response.
			return KEYWORD_OK, nil
		case len(line) == 5 && line[1] == 'P' && line[2] == 'O' && line[3] == 'N' && line[4] == 'G':
			// Avoid allocation in PING command benchmarks :)
			return KEYWORD_PONG, nil
		default:
			return string(line[1:]), nil
		}
	case DOLLAR_BYTE:
		n, err := p.parseLen(line[1:])
		if n < 0 || err != nil {
			return nil, err
		}
		arr := make([]byte, n)
		_, err = io.ReadFull(p.is, arr)
		if err != nil {
			return nil, err
		}
		if newLine, err := p.readLine(); err != nil {
			return nil, err
		} else if len(newLine) != 0 {
			return nil, errors.New("bad bulk string format")
		}
		return arr, nil
	case ASTERISK_BYTE:
		n, err := p.parseLen(line[1:])
		if n < 0 || err != nil {
			return nil, err
		}
		r := make([]interface{}, n)
		for i := range r {
			r[i], err = p.process()
			if err != nil {
				return nil, err
			}
		}
		return r, nil
	case COLON_BYTE:
		return p.parseInt(line[1:])
	case MINUS_BYTE:
		return nil, errors.New(string(line[1:]))
	default:
		return nil, errors.New(fmt.Sprint("Unknown reply: ", line[0]))
	}
}

func (p *protocol) readLine() ([]byte, error) {
	// To avoid allocations, attempt to read the line using ReadSlice. This
	// call typically succeeds. The known case where the call fails is when
	// reading the output from the MONITOR command.
	line, err := p.is.ReadSlice('\n')
	if err == bufio.ErrBufferFull {
		// The line does not fit in the bufio.Reader's buffer. Fall back to
		// allocating a buffer for the line.
		buf := append([]byte{}, line...)
		for err == bufio.ErrBufferFull {
			line, err = p.is.ReadSlice('\n')
			buf = append(buf, line...)
		}
		line = buf
	}
	if err != nil {
		return nil, err
	}
	i := len(line) - 2
	if i < 0 || line[i] != '\r' {
		return nil, errors.New("bad response line terminator")
	}
	return line[:i], nil
}

func (p *protocol) parseLen(b []byte) (int, error) {
	if len(b) == 0 {
		return -1, errors.New("malformed length")
	}

	if b[0] == '-' && len(b) == 2 && b[1] == '1' {
		// handle $-1 and $-1 null replies.
		return -1, nil
	}

	var n int
	for _, b := range b {
		n *= 10
		if b < '0' || b > '9' {
			return -1, errors.New("illegal bytes in length")
		}
		n += int(b - '0')
	}

	return n, nil
}

// parseInt parses an integer reply.
func (p *protocol) parseInt(b []byte) (interface{}, error) {
	if len(b) == 0 {
		return 0, errors.New("malformed integer")
	}

	var negate bool
	if b[0] == '-' {
		negate = true
		b = b[1:]
		if len(b) == 0 {
			return 0, errors.New("malformed integer")
		}
	}

	var n int64
	for _, a := range b {
		n *= 10
		if a < '0' || a > '9' {
			return 0, errors.New("illegal bytes in length")
		}
		n += int64(a - '0')
	}

	if negate {
		n = -n
	}
	return n, nil
}

type ProtocolCommand interface {
	GetRaw() []byte
}

type protocolCommand struct {
	Name string
}

func (p protocolCommand) GetRaw() []byte {
	return []byte(p.Name)
}

func newProtocolCommand(name string) protocolCommand {
	return protocolCommand{name}
}

var (
	CMD_PING                 = newProtocolCommand("PING")
	CMD_SET                  = newProtocolCommand("SET")
	CMD_GET                  = newProtocolCommand("GET")
	CMD_QUIT                 = newProtocolCommand("QUIT")
	CMD_EXISTS               = newProtocolCommand("EXISTS")
	CMD_DEL                  = newProtocolCommand("DEL")
	CMD_UNLINK               = newProtocolCommand("UNLINK")
	CMD_TYPE                 = newProtocolCommand("TYPE")
	CMD_FLUSHDB              = newProtocolCommand("FLUSHDB")
	CMD_KEYS                 = newProtocolCommand("KEYS")
	CMD_RANDOMKEY            = newProtocolCommand("RANDOMKEY")
	CMD_RENAME               = newProtocolCommand("RENAME")
	CMD_RENAMENX             = newProtocolCommand("RENAMENX")
	CMD_RENAMEX              = newProtocolCommand("RENAMEX")
	CMD_DBSIZE               = newProtocolCommand("DBSIZE")
	CMD_EXPIRE               = newProtocolCommand("EXPIRE")
	CMD_EXPIREAT             = newProtocolCommand("EXPIREAT")
	CMD_TTL                  = newProtocolCommand("TTL")
	CMD_SELECT               = newProtocolCommand("SELECT")
	CMD_MOVE                 = newProtocolCommand("MOVE")
	CMD_FLUSHALL             = newProtocolCommand("FLUSHALL")
	CMD_GETSET               = newProtocolCommand("GETSET")
	CMD_MGET                 = newProtocolCommand("MGET")
	CMD_SETNX                = newProtocolCommand("SETNX")
	CMD_SETEX                = newProtocolCommand("SETEX")
	CMD_MSET                 = newProtocolCommand("MSET")
	CMD_MSETNX               = newProtocolCommand("MSETNX")
	CMD_DECRBY               = newProtocolCommand("DECRBY")
	CMD_DECR                 = newProtocolCommand("DECR")
	CMD_INCRBY               = newProtocolCommand("INCRBY")
	CMD_INCR                 = newProtocolCommand("INCR")
	CMD_APPEND               = newProtocolCommand("APPEND")
	CMD_SUBSTR               = newProtocolCommand("SUBSTR")
	CMD_HSET                 = newProtocolCommand("HSET")
	CMD_HGET                 = newProtocolCommand("HGET")
	CMD_HSETNX               = newProtocolCommand("HSETNX")
	CMD_HMSET                = newProtocolCommand("HMSET")
	CMD_HMGET                = newProtocolCommand("HMGET")
	CMD_HINCRBY              = newProtocolCommand("HINCRBY")
	CMD_HEXISTS              = newProtocolCommand("HEXISTS")
	CMD_HDEL                 = newProtocolCommand("HDEL")
	CMD_HLEN                 = newProtocolCommand("HLEN")
	CMD_HKEYS                = newProtocolCommand("HKEYS")
	CMD_HVALS                = newProtocolCommand("HVALS")
	CMD_HGETALL              = newProtocolCommand("HGETALL")
	CMD_RPUSH                = newProtocolCommand("RPUSH")
	CMD_LPUSH                = newProtocolCommand("LPUSH")
	CMD_LLEN                 = newProtocolCommand("LLEN")
	CMD_LRANGE               = newProtocolCommand("LRANGE")
	CMD_LTRIM                = newProtocolCommand("LTRIM")
	CMD_LINDEX               = newProtocolCommand("LINDEX")
	CMD_LSET                 = newProtocolCommand("LSET")
	CMD_LREM                 = newProtocolCommand("LREM")
	CMD_LPOP                 = newProtocolCommand("LPOP")
	CMD_RPOP                 = newProtocolCommand("RPOP")
	CMD_RPOPLPUSH            = newProtocolCommand("RPOPLPUSH")
	CMD_SADD                 = newProtocolCommand("SADD")
	CMD_SMEMBERS             = newProtocolCommand("SMEMBERS")
	CMD_SREM                 = newProtocolCommand("SREM")
	CMD_SPOP                 = newProtocolCommand("SPOP")
	CMD_SMOVE                = newProtocolCommand("SMOVE")
	CMD_SCARD                = newProtocolCommand("SCARD")
	CMD_SISMEMBER            = newProtocolCommand("SISMEMBER")
	CMD_SINTER               = newProtocolCommand("SINTER")
	CMD_SINTERSTORE          = newProtocolCommand("SINTERSTORE")
	CMD_SUNION               = newProtocolCommand("SUNION")
	CMD_SUNIONSTORE          = newProtocolCommand("SUNIONSTORE")
	CMD_SDIFF                = newProtocolCommand("SDIFF")
	CMD_SDIFFSTORE           = newProtocolCommand("SDIFFSTORE")
	CMD_SRANDMEMBER          = newProtocolCommand("SRANDMEMBER")
	CMD_ZADD                 = newProtocolCommand("ZADD")
	CMD_ZRANGE               = newProtocolCommand("ZRANGE")
	CMD_ZREM                 = newProtocolCommand("ZREM")
	CMD_ZINCRBY              = newProtocolCommand("ZINCRBY")
	CMD_ZRANK                = newProtocolCommand("ZRANK")
	CMD_ZREVRANK             = newProtocolCommand("ZREVRANK")
	CMD_ZREVRANGE            = newProtocolCommand("ZREVRANGE")
	CMD_ZCARD                = newProtocolCommand("ZCARD")
	CMD_ZSCORE               = newProtocolCommand("ZSCORE")
	CMD_MULTI                = newProtocolCommand("MULTI")
	CMD_DISCARD              = newProtocolCommand("DISCARD")
	CMD_EXEC                 = newProtocolCommand("EXEC")
	CMD_WATCH                = newProtocolCommand("WATCH")
	CMD_UNWATCH              = newProtocolCommand("UNWATCH")
	CMD_SORT                 = newProtocolCommand("SORT")
	CMD_BLPOP                = newProtocolCommand("BLPOP")
	CMD_BRPOP                = newProtocolCommand("BRPOP")
	CMD_AUTH                 = newProtocolCommand("AUTH")
	CMD_SUBSCRIBE            = newProtocolCommand("SUBSCRIBE")
	CMD_PUBLISH              = newProtocolCommand("PUBLISH")
	CMD_UNSUBSCRIBE          = newProtocolCommand("UNSUBSCRIBE")
	CMD_PSUBSCRIBE           = newProtocolCommand("PSUBSCRIBE")
	CMD_PUNSUBSCRIBE         = newProtocolCommand("PUNSUBSCRIBE")
	CMD_PUBSUB               = newProtocolCommand("PUBSUB")
	CMD_ZCOUNT               = newProtocolCommand("ZCOUNT")
	CMD_ZRANGEBYSCORE        = newProtocolCommand("ZRANGEBYSCORE")
	CMD_ZREVRANGEBYSCORE     = newProtocolCommand("ZREVRANGEBYSCORE")
	CMD_ZREMRANGEBYRANK      = newProtocolCommand("ZREMRANGEBYRANK")
	CMD_ZREMRANGEBYSCORE     = newProtocolCommand("ZREMRANGEBYSCORE")
	CMD_ZUNIONSTORE          = newProtocolCommand("ZUNIONSTORE")
	CMD_ZINTERSTORE          = newProtocolCommand("ZINTERSTORE")
	CMD_ZLEXCOUNT            = newProtocolCommand("ZLEXCOUNT")
	CMD_ZRANGEBYLEX          = newProtocolCommand("ZRANGEBYLEX")
	CMD_ZREVRANGEBYLEX       = newProtocolCommand("ZREVRANGEBYLEX")
	CMD_ZREMRANGEBYLEX       = newProtocolCommand("ZREMRANGEBYLEX")
	CMD_SAVE                 = newProtocolCommand("SAVE")
	CMD_BGSAVE               = newProtocolCommand("BGSAVE")
	CMD_BGREWRITEAOF         = newProtocolCommand("BGREWRITEAOF")
	CMD_LASTSAVE             = newProtocolCommand("LASTSAVE")
	CMD_SHUTDOWN             = newProtocolCommand("SHUTDOWN")
	CMD_INFO                 = newProtocolCommand("INFO")
	CMD_MONITOR              = newProtocolCommand("MONITOR")
	CMD_SLAVEOF              = newProtocolCommand("SLAVEOF")
	CMD_CONFIG               = newProtocolCommand("CONFIG")
	CMD_STRLEN               = newProtocolCommand("STRLEN")
	CMD_SYNC                 = newProtocolCommand("SYNC")
	CMD_LPUSHX               = newProtocolCommand("LPUSHX")
	CMD_PERSIST              = newProtocolCommand("PERSIST")
	CMD_RPUSHX               = newProtocolCommand("RPUSHX")
	CMD_ECHO                 = newProtocolCommand("ECHO")
	CMD_LINSERT              = newProtocolCommand("LINSERT")
	CMD_DEBUG                = newProtocolCommand("DEBUG")
	CMD_BRPOPLPUSH           = newProtocolCommand("BRPOPLPUSH")
	CMD_SETBIT               = newProtocolCommand("SETBIT")
	CMD_GETBIT               = newProtocolCommand("GETBIT")
	CMD_BITPOS               = newProtocolCommand("BITPOS")
	CMD_SETRANGE             = newProtocolCommand("SETRANGE")
	CMD_GETRANGE             = newProtocolCommand("GETRANGE")
	CMD_EVAL                 = newProtocolCommand("EVAL")
	CMD_EVALSHA              = newProtocolCommand("EVALSHA")
	CMD_SCRIPT               = newProtocolCommand("SCRIPT")
	CMD_SLOWLOG              = newProtocolCommand("SLOWLOG")
	CMD_OBJECT               = newProtocolCommand("OBJECT")
	CMD_BITCOUNT             = newProtocolCommand("BITCOUNT")
	CMD_BITOP                = newProtocolCommand("BITOP")
	CMD_SENTINEL             = newProtocolCommand("SENTINEL")
	CMD_DUMP                 = newProtocolCommand("DUMP")
	CMD_RESTORE              = newProtocolCommand("RESTORE")
	CMD_PEXPIRE              = newProtocolCommand("PEXPIRE")
	CMD_PEXPIREAT            = newProtocolCommand("PEXPIREAT")
	CMD_PTTL                 = newProtocolCommand("PTTL")
	CMD_INCRBYFLOAT          = newProtocolCommand("INCRBYFLOAT")
	CMD_PSETEX               = newProtocolCommand("PSETEX")
	CMD_CLIENT               = newProtocolCommand("CLIENT")
	CMD_TIME                 = newProtocolCommand("TIME")
	CMD_MIGRATE              = newProtocolCommand("MIGRATE")
	CMD_HINCRBYFLOAT         = newProtocolCommand("HINCRBYFLOAT")
	CMD_SCAN                 = newProtocolCommand("SCAN")
	CMD_HSCAN                = newProtocolCommand("HSCAN")
	CMD_SSCAN                = newProtocolCommand("SSCAN")
	CMD_ZSCAN                = newProtocolCommand("ZSCAN")
	CMD_WAIT                 = newProtocolCommand("WAIT")
	CMD_CLUSTER              = newProtocolCommand("CLUSTER")
	CMD_ASKING               = newProtocolCommand("ASKING")
	CMD_PFADD                = newProtocolCommand("PFADD")
	CMD_PFCOUNT              = newProtocolCommand("PFCOUNT")
	CMD_PFMERGE              = newProtocolCommand("PFMERGE")
	CMD_READONLY             = newProtocolCommand("READONLY")
	CMD_GEOADD               = newProtocolCommand("GEOADD")
	CMD_GEODIST              = newProtocolCommand("GEODIST")
	CMD_GEOHASH              = newProtocolCommand("GEOHASH")
	CMD_GEOPOS               = newProtocolCommand("GEOPOS")
	CMD_GEORADIUS            = newProtocolCommand("GEORADIUS")
	CMD_GEORADIUS_RO         = newProtocolCommand("GEORADIUS_RO")
	CMD_GEORADIUSBYMEMBER    = newProtocolCommand("GEORADIUSBYMEMBER")
	CMD_GEORADIUSBYMEMBER_RO = newProtocolCommand("GEORADIUSBYMEMBER_RO")
	CMD_MODULE               = newProtocolCommand("MODULE")
	CMD_BITFIELD             = newProtocolCommand("BITFIELD")
	CMD_HSTRLEN              = newProtocolCommand("HSTRLEN")
	CMD_TOUCH                = newProtocolCommand("TOUCH")
	CMD_SWAPDB               = newProtocolCommand("SWAPDB")
	CMD_MEMORY               = newProtocolCommand("MEMORY")
	CMD_XADD                 = newProtocolCommand("XADD")
	CMD_XLEN                 = newProtocolCommand("XLEN")
	CMD_XDEL                 = newProtocolCommand("XDEL")
	CMD_XTRIM                = newProtocolCommand("XTRIM")
	CMD_XRANGE               = newProtocolCommand("XRANGE")
	CMD_XREVRANGE            = newProtocolCommand("XREVRANGE")
	CMD_XREAD                = newProtocolCommand("XREAD")
	CMD_XACK                 = newProtocolCommand("XACK")
	CMD_XGROUP               = newProtocolCommand("XGROUP")
	CMD_XREADGROUP           = newProtocolCommand("XREADGROUP")
	CMD_XPENDING             = newProtocolCommand("XPENDING")
	CMD_XCLAIM               = newProtocolCommand("XCLAIM")
)

type keyword struct {
	Name string
}

func (k keyword) GetRaw() []byte {
	return []byte(k.Name)
}

func newKeyword(name string) keyword {
	return keyword{name}
}

var (
	KEYWORD_AGGREGATE    = newKeyword("AGGREGATE")
	KEYWORD_ALPHA        = newKeyword("ALPHA")
	KEYWORD_ASC          = newKeyword("ASC")
	KEYWORD_BY           = newKeyword("BY")
	KEYWORD_DESC         = newKeyword("DESC")
	KEYWORD_GET          = newKeyword("GET")
	KEYWORD_LIMIT        = newKeyword("LIMIT")
	KEYWORD_MESSAGE      = newKeyword("MESSAGE")
	KEYWORD_NO           = newKeyword("NO")
	KEYWORD_NOSORT       = newKeyword("NOSORT")
	KEYWORD_PMESSAGE     = newKeyword("PMESSAGE")
	KEYWORD_PSUBSCRIBE   = newKeyword("PSUBSCRIBE")
	KEYWORD_PUNSUBSCRIBE = newKeyword("PUNSUBSCRIBE")
	KEYWORD_OK           = newKeyword("OK")
	KEYWORD_ONE          = newKeyword("ONE")
	KEYWORD_QUEUED       = newKeyword("QUEUED")
	KEYWORD_SET          = newKeyword("SET")
	KEYWORD_STORE        = newKeyword("STORE")
	KEYWORD_SUBSCRIBE    = newKeyword("SUBSCRIBE")
	KEYWORD_UNSUBSCRIBE  = newKeyword("UNSUBSCRIBE")
	KEYWORD_WEIGHTS      = newKeyword("WEIGHTS")
	KEYWORD_WITHSCORES   = newKeyword("WITHSCORES")
	KEYWORD_RESETSTAT    = newKeyword("RESETSTAT")
	KEYWORD_REWRITE      = newKeyword("REWRITE")
	KEYWORD_RESET        = newKeyword("RESET")
	KEYWORD_FLUSH        = newKeyword("FLUSH")
	KEYWORD_EXISTS       = newKeyword("EXISTS")
	KEYWORD_LOAD         = newKeyword("LOAD")
	KEYWORD_KILL         = newKeyword("KILL")
	KEYWORD_LEN          = newKeyword("LEN")
	KEYWORD_REFCOUNT     = newKeyword("REFCOUNT")
	KEYWORD_ENCODING     = newKeyword("ENCODING")
	KEYWORD_IDLETIME     = newKeyword("IDLETIME")
	KEYWORD_GETNAME      = newKeyword("GETNAME")
	KEYWORD_SETNAME      = newKeyword("SETNAME")
	KEYWORD_LIST         = newKeyword("LIST")
	KEYWORD_MATCH        = newKeyword("MATCH")
	KEYWORD_COUNT        = newKeyword("COUNT")
	KEYWORD_PING         = newKeyword("PING")
	KEYWORD_PONG         = newKeyword("PONG")
	KEYWORD_UNLOAD       = newKeyword("UNLOAD")
	KEYWORD_REPLACE      = newKeyword("REPLACE")
	KEYWORD_KEYS         = newKeyword("KEYS")
	KEYWORD_PAUSE        = newKeyword("PAUSE")
	KEYWORD_DOCTOR       = newKeyword("DOCTOR")
	KEYWORD_BLOCK        = newKeyword("BLOCK")
	KEYWORD_NOACK        = newKeyword("NOACK")
	KEYWORD_STREAMS      = newKeyword("STREAMS")
	KEYWORD_KEY          = newKeyword("KEY")
	KEYWORD_CREATE       = newKeyword("CREATE")
	KEYWORD_MKSTREAM     = newKeyword("MKSTREAM")
	KEYWORD_SETID        = newKeyword("SETID")
	KEYWORD_DESTROY      = newKeyword("DESTROY")
	KEYWORD_DELCONSUMER  = newKeyword("DELCONSUMER")
	KEYWORD_MAXLEN       = newKeyword("MAXLEN")
	KEYWORD_GROUP        = newKeyword("GROUP")
	KEYWORD_IDLE         = newKeyword("IDLE")
	KEYWORD_TIME         = newKeyword("TIME")
	KEYWORD_RETRYCOUNT   = newKeyword("RETRYCOUNT")
	KEYWORD_FORCE        = newKeyword("FORCE")
)

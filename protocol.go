package godis

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	AskPrefix         = "ASK "
	MovedPrefix       = "MOVED "
	ClusterdownPrefix = "CLUSTERDOWN "
	BusyPrefix        = "BUSY "
	NoscriptPrefix    = "NOSCRIPT "

	DefaultHost         = "localhost"
	DefaultPort         = 6379
	DefaultSentinelPort = 26379
	DefaultTimeout      = 5 * time.Second
	DefaultDatabase     = 2 * time.Second

	DollarByte   = '$'
	AsteriskByte = '*'
	PlusByte     = '+'
	MinusByte    = '-'
	ColonByte    = ':'

	SentinelMasters             = "masters"
	SentinelGetMasterAddrByName = "get-master-addr-by-name"
	SentinelReset               = "reset"
	SentinelSlaves              = "slaves"
	SentinelFailover            = "failover"
	SentinelMonitor             = "monitor"
	SentinelRemove              = "remove"
	SentinelSet                 = "set"

	ClusterNodes            = "nodes"
	ClusterMeet             = "meet"
	ClusterReset            = "reset"
	ClusterAddslots         = "addslots"
	ClusterDelslots         = "delslots"
	ClusterInfo             = "info"
	ClusterGetkeysinslot    = "getkeysinslot"
	ClusterSetslot          = "setslot"
	ClusterSetslotNode      = "node"
	ClusterSetslotMigrating = "migrating"
	ClusterSetslotImporting = "importing"
	ClusterSetslotStable    = "stable"
	ClusterForget           = "forget"
	ClusterFlushslot        = "flushslots"
	ClusterKeyslot          = "keyslot"
	ClusterCountkeyinslot   = "countkeysinslot"
	ClusterSaveconfig       = "saveconfig"
	ClusterReplicate        = "replicate"
	ClusterSlaves           = "slaves"
	ClusterFailover         = "failover"
	ClusterSlots            = "slots"
	PubsubChannels          = "channels"
	PubsubNumsub            = "numsub"
	PubsubNumPat            = "numpat"
)

var (
	BytesTrue  = IntToByteArray(1)
	BytesFalse = IntToByteArray(0)
	BytesTilde = []byte("~")

	PositiveInfinityBytes = []byte("+inf")
	NegativeInfinityBytes = []byte("-inf")
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

var (
	sizeTable = []int{9, 99, 999, 9999, 99999, 999999, 9999999, 99999999,
		999999999, MaxInt}

	DigitTens = []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1',
		'1', '1', '1', '1', '1', '1', '1', '1', '1', '2', '2', '2', '2', '2', '2', '2', '2', '2',
		'2', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4', '4', '4',
		'4', '4', '4', '5', '5', '5', '5', '5', '5', '5', '5', '5', '5', '6', '6', '6', '6', '6',
		'6', '6', '6', '6', '6', '7', '7', '7', '7', '7', '7', '7', '7', '7', '7', '8', '8', '8',
		'8', '8', '8', '8', '8', '8', '8', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9'}

	DigitOnes = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8',
		'9', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6',
		'7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4',
		'5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2',
		'3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a',
		'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

// send message to redis
type redisOutputStream struct {
	*bufio.Writer
	buf   []byte
	count int
	c     *connection
}

func newRedisOutputStream(bw *bufio.Writer, c *connection) *redisOutputStream {
	return &redisOutputStream{
		Writer: bw,
		buf:    make([]byte, 8192),
		c:      c,
	}
}

func (r *redisOutputStream) writeIntCrLf(b int) error {
	//_, err := r.Write(strconv.AppendInt(r.buf, int64(b), 10))
	//if err != nil {
	//	return 0, err
	//}
	//return r.writeCrLf()
	if b < 0 {
		if err := r.writeByte('-'); err != nil {
			return err
		}
		b = -b
	}
	size := 0
	for b > sizeTable[size] {
		size++
	}
	size++
	if size >= len(r.buf)-r.count {
		if err := r.flushBuffer(); err != nil {
			return err
		}
	}
	q, p := 0, 0
	charPos := r.count + size
	for b >= 65536 {
		q = b / 100
		p = b - ((q << 6) + (q << 5) + (q << 2))
		b = q
		charPos--
		r.buf[charPos] = DigitOnes[p]
		charPos--
		r.buf[charPos] = DigitTens[p]
	}
	for {
		q = (b * 52429) >> (16 + 3)
		p = b - ((q << 3) + (q << 1))
		charPos--
		r.buf[charPos] = digits[p]
		b = q
		if b == 0 {
			break
		}
	}
	r.count += size
	return r.writeCrLf()
}

func (r *redisOutputStream) writeCrLf() error {
	//return r.WriteString("\r\n")
	if 2 >= len(r.buf)-r.count {
		if err := r.flushBuffer(); err != nil {
			return err
		}
	}
	r.buf[r.count] = '\r'
	r.count++
	r.buf[r.count] = '\n'
	r.count++
	return nil
}

func (r *redisOutputStream) flushBuffer() error {
	if r.count > 0 {
		_, err := r.Write(r.buf[0:r.count])
		if err != nil {
			return err
		}
		r.count = 0
	}
	return nil
}

func (r *redisOutputStream) writeByte(b byte) error {
	if r.count == len(r.buf) {
		return r.flushBuffer()
	}
	r.buf[r.count] = b
	r.count++
	return nil
}

func (r *redisOutputStream) write(b []byte) error {
	return r.writeWithPos(b, 0, len(b))
}

func (r *redisOutputStream) writeWithPos(b []byte, off, size int) error {
	if size >= len(r.buf) {
		err := r.flushBuffer()
		if err != nil {
			return err
		}
		_, err = r.Write(b[off:size])
		return err
	} else {
		if size >= len(r.buf)-r.count {
			err := r.flushBuffer()
			if err != nil {
				return err
			}
		}
		for i := off; i < size; i++ {
			r.buf[r.count] = b[i]
			r.count++
		}
		return nil
	}
}

func (r *redisOutputStream) flush() error {
	r.flushBuffer()
	if err := r.Flush(); err != nil {
		return err
	}
	return r.c.socket.SetDeadline(time.Now().Add(r.c.soTimeout))
}

// receive message from redis
type redisInputStream struct {
	*bufio.Reader
	buf   []byte
	count int
	limit int
	c     *connection
}

func newRedisInputStream(br *bufio.Reader, c *connection) *redisInputStream {
	return &redisInputStream{
		Reader: br,
		buf:    make([]byte, 8192),
		c:      c,
	}
}

func (r *redisInputStream) readByte() (byte, error) {
	err := r.ensureFill()
	if err != nil {
		return 0, err
	}
	ret := r.buf[r.count]
	r.count++
	return ret, nil
}

func (r *redisInputStream) ensureFill() error {
	if r.count < r.limit {
		return nil
	}
	var err error
	r.limit, err = r.Read(r.buf)
	if err != nil {
		return newConnectError(err.Error())
	}
	err = r.c.socket.SetDeadline(time.Now().Add(r.c.soTimeout))
	if err != nil {
		return newConnectError(err.Error())
	}
	r.count = 0
	if r.limit == -1 {
		return newConnectError("Unexpected end of stream")
	}
	return nil
}

func (r *redisInputStream) readLine() (string, error) {
	buf := ""
	for {
		err := r.ensureFill()
		if err != nil {
			return "", err
		}
		b := r.buf[r.count]
		r.count++
		if b == '\r' {
			err := r.ensureFill()
			if err != nil {
				return "", err
			}
			c := r.buf[r.count]
			r.count++
			if c == '\n' {
				break
			}
			buf += string(b)
			buf += string(c)
		} else {
			buf += string(b)
		}
	}
	if buf == "" {
		return "", newConnectError("It seems like server has closed the connection.")
	}
	return buf, nil
}

func (r *redisInputStream) readLineBytes() ([]byte, error) {
	err := r.ensureFill()
	if err != nil {
		return nil, err
	}
	pos := r.count
	buf := r.buf
	for {
		if pos == r.limit {
			return r.readLineBytesSlowly()
		}
		p := buf[pos]
		pos++
		if p == '\r' {
			if pos == r.limit {
				return r.readLineBytesSlowly()
			}
			p := buf[pos]
			pos++
			if p == '\n' {
				break
			}
		}
	}
	N := pos - r.count - 2
	line := make([]byte, N)
	j := 0
	for i := r.count; i <= N; i++ {
		line[j] = buf[i]
		j++
	}
	r.count = pos
	return line, nil
}

func (r *redisInputStream) readLineBytesSlowly() ([]byte, error) {
	buf := make([]byte, 0)
	for {
		err := r.ensureFill()
		if err != nil {
			return nil, err
		}
		b := r.buf[r.count]
		r.count++
		if b == 'r' {
			err := r.ensureFill()
			if err != nil {
				return nil, err
			}
			c := r.buf[r.count]
			r.count++
			if c == '\n' {
				break
			}
			buf = append(buf, b)
			buf = append(buf, c)
		} else {
			buf = append(buf, b)
		}
	}
	return buf, nil
}

func (r *redisInputStream) readIntCrLf() (int64, error) {
	err := r.ensureFill()
	if err != nil {
		return 0, err
	}
	buf := r.buf
	isNeg := false
	if buf[r.count] == '-' {
		isNeg = true
	}
	if isNeg {
		r.count++
	}
	var value int64 = 0
	for {
		err := r.ensureFill()
		if err != nil {
			return 0, err
		}
		b := buf[r.count]
		r.count++
		if b == '\r' {
			err := r.ensureFill()
			if err != nil {
				return 0, err
			}
			c := buf[r.count]
			r.count++
			if c != '\n' {
				return 0, newConnectError("Unexpected character!")
			}
			break
		} else {
			value = value*10 + int64(b) - int64('0')
		}
	}
	if isNeg {
		return -value, nil
	}
	return value, nil
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
	if err := p.os.writeByte(AsteriskByte); err != nil {
		return err
	}
	if err := p.os.writeIntCrLf(len(args) + 1); err != nil {
		return err
	}
	if err := p.os.writeByte(DollarByte); err != nil {
		return err
	}
	if err := p.os.writeIntCrLf(len(command)); err != nil {
		return err
	}
	if err := p.os.write(command); err != nil {
		return err
	}
	if err := p.os.writeCrLf(); err != nil {
		return err
	}
	for _, arg := range args {
		if err := p.os.writeByte(DollarByte); err != nil {
			return err
		}
		if err := p.os.writeIntCrLf(len(arg)); err != nil {
			return err
		}
		if err := p.os.write(arg); err != nil {
			return err
		}
		if err := p.os.writeCrLf(); err != nil {
			return err
		}
	}
	return nil
}

func (p *protocol) read() (interface{}, error) {
	return p.process()
}

func (p *protocol) process() (interface{}, error) {
	b, err := p.is.readByte()
	if err != nil {
		return nil, newConnectError(err.Error())
	}
	switch b {
	case PlusByte:
		return p.processStatusCodeReply()
	case DollarByte:
		return p.processBulkReply()
	case AsteriskByte:
		return p.processMultiBulkReply()
	case ColonByte:
		return p.processInteger()
	case MinusByte:
		return p.processError()
	default:
		return nil, newConnectError(fmt.Sprintf("Unknown reply: %b", b))
	}
}

func (p *protocol) processStatusCodeReply() ([]byte, error) {
	return p.is.readLineBytes()
}

func (p *protocol) processBulkReply() ([]byte, error) {
	l, err := p.is.readIntCrLf()
	if err != nil {
		return nil, newConnectError(err.Error())
	}
	if l == -1 {
		return nil, nil
	}
	line := make([]byte, 0)
	for {
		err := p.is.ensureFill()
		if err != nil {
			return nil, err
		}
		b := p.is.buf[p.is.count]
		p.is.count++
		if b == '\r' {
			err := p.is.ensureFill()
			if err != nil {
				return nil, err
			}
			c := p.is.buf[p.is.count]
			p.is.count++
			if c != '\n' {
				return nil, newConnectError("Unexpected character!")
			}
			break
		} else {
			line = append(line, b)
		}
	}
	return line, nil
}

func (p *protocol) processMultiBulkReply() ([]interface{}, error) {
	l, err := p.is.readIntCrLf()
	if err != nil {
		return nil, newConnectError(err.Error())
	}
	if l == -1 {
		return nil, nil
	}
	ret := make([]interface{}, 0)
	for i := 0; i < int(l); i++ {
		if obj, err := p.process(); err != nil {
			ret = append(ret, newDataError(err.Error()))
		} else {
			ret = append(ret, obj)
		}
	}
	return ret, nil
}

func (p *protocol) processInteger() (int64, error) {
	return p.is.readIntCrLf()
}

func (p *protocol) processError() (interface{}, error) {
	msg, err := p.is.readLine()
	if err != nil {
		return nil, newConnectError(err.Error())
	}
	if strings.HasPrefix(msg, MovedPrefix) {
		host, port, slot := p.parseTargetHostAndSlot(msg)
		return nil, newMovedDataError(msg, host, port, slot)
	} else if strings.HasPrefix(msg, AskPrefix) {
		host, port, slot := p.parseTargetHostAndSlot(msg)
		return nil, newAskDataError(msg, host, port, slot)
	} else if strings.HasPrefix(msg, ClusterdownPrefix) {
		return nil, newClusterError(msg)
	} else if strings.HasPrefix(msg, BusyPrefix) {
		return nil, newBusyError(msg)
	} else if strings.HasPrefix(msg, NoscriptPrefix) {
		return nil, newNoScriptError(msg)
	}
	return nil, newDataError(msg)
}

func (p *protocol) parseTargetHostAndSlot(clusterRedirectResponse string) (string, int, int) {
	arr := strings.Split(clusterRedirectResponse, " ")
	host, port := p.extractParts(arr[2])
	slot, _ := strconv.Atoi(arr[1])
	po, _ := strconv.Atoi(port)
	return host, po, slot
}

func (p *protocol) extractParts(from string) (string, string) {
	idx := strings.LastIndex(from, ":")
	host := from
	if idx != -1 {
		host = from[0:idx]
	}
	port := ""
	if idx != -1 {
		port = from[idx+1:]
	}
	return host, port
}

// redis protocol command
type protocolCommand struct {
	Name string // name of command
}

// GetRaw get name byte array
func (p protocolCommand) GetRaw() []byte {
	return []byte(p.Name)
}

func newProtocolCommand(name string) protocolCommand {
	return protocolCommand{name}
}

var (
	CmdPing                = newProtocolCommand("PING")
	CmdSet                 = newProtocolCommand("SET")
	CmdGet                 = newProtocolCommand("GET")
	CmdQuit                = newProtocolCommand("QUIT")
	CmdExists              = newProtocolCommand("EXISTS")
	CmdDel                 = newProtocolCommand("DEL")
	CmdUnlink              = newProtocolCommand("UNLINK")
	CmdType                = newProtocolCommand("TYPE")
	CmdFlushdb             = newProtocolCommand("FLUSHDB")
	CmdKeys                = newProtocolCommand("KEYS")
	CmdRandomkey           = newProtocolCommand("RANDOMKEY")
	CmdRename              = newProtocolCommand("RENAME")
	CmdRenamenx            = newProtocolCommand("RENAMENX")
	CmdRenamex             = newProtocolCommand("RENAMEX")
	CmdDbsize              = newProtocolCommand("DBSIZE")
	CmdExpire              = newProtocolCommand("EXPIRE")
	CmdExpireat            = newProtocolCommand("EXPIREAT")
	CmdTtl                 = newProtocolCommand("TTL")
	CmdSelect              = newProtocolCommand("SELECT")
	CmdMove                = newProtocolCommand("MOVE")
	CmdFlushall            = newProtocolCommand("FLUSHALL")
	CmdGetset              = newProtocolCommand("GETSET")
	CmdMget                = newProtocolCommand("MGET")
	CmdSetnx               = newProtocolCommand("SETNX")
	CmdSetex               = newProtocolCommand("SETEX")
	CmdMset                = newProtocolCommand("MSET")
	CmdMsetnx              = newProtocolCommand("MSETNX")
	CmdDecrby              = newProtocolCommand("DECRBY")
	CmdDecr                = newProtocolCommand("DECR")
	CmdIncrby              = newProtocolCommand("INCRBY")
	CmdIncr                = newProtocolCommand("INCR")
	CmdAppend              = newProtocolCommand("APPEND")
	CmdSubstr              = newProtocolCommand("SUBSTR")
	CmdHset                = newProtocolCommand("HSET")
	CmdHget                = newProtocolCommand("HGET")
	CmdHsetnx              = newProtocolCommand("HSETNX")
	CmdHmset               = newProtocolCommand("HMSET")
	CmdHmget               = newProtocolCommand("HMGET")
	CmdHincrby             = newProtocolCommand("HINCRBY")
	CmdHexists             = newProtocolCommand("HEXISTS")
	CmdHdel                = newProtocolCommand("HDEL")
	CmdHlen                = newProtocolCommand("HLEN")
	CmdHkeys               = newProtocolCommand("HKEYS")
	CmdHvals               = newProtocolCommand("HVALS")
	CmdHgetall             = newProtocolCommand("HGETALL")
	CmdRpush               = newProtocolCommand("RPUSH")
	CmdLpush               = newProtocolCommand("LPUSH")
	CmdLlen                = newProtocolCommand("LLEN")
	CmdLrange              = newProtocolCommand("LRANGE")
	CmdLtrim               = newProtocolCommand("LTRIM")
	CmdLindex              = newProtocolCommand("LINDEX")
	CmdLset                = newProtocolCommand("LSET")
	CmdLrem                = newProtocolCommand("LREM")
	CmdLpop                = newProtocolCommand("LPOP")
	CmdRpop                = newProtocolCommand("RPOP")
	CmdRpoplpush           = newProtocolCommand("RPOPLPUSH")
	CmdSadd                = newProtocolCommand("SADD")
	CmdSmembers            = newProtocolCommand("SMEMBERS")
	CmdSrem                = newProtocolCommand("SREM")
	CmdSpop                = newProtocolCommand("SPOP")
	CmdSmove               = newProtocolCommand("SMOVE")
	CmdScard               = newProtocolCommand("SCARD")
	CmdSismember           = newProtocolCommand("SISMEMBER")
	CmdSinter              = newProtocolCommand("SINTER")
	CmdSinterstore         = newProtocolCommand("SINTERSTORE")
	CmdSunion              = newProtocolCommand("SUNION")
	CmdSunionstore         = newProtocolCommand("SUNIONSTORE")
	CmdSdiff               = newProtocolCommand("SDIFF")
	CmdSdiffstore          = newProtocolCommand("SDIFFSTORE")
	CmdSrandmember         = newProtocolCommand("SRANDMEMBER")
	CmdZadd                = newProtocolCommand("ZADD")
	CmdZrange              = newProtocolCommand("ZRANGE")
	CmdZrem                = newProtocolCommand("ZREM")
	CmdZincrby             = newProtocolCommand("ZINCRBY")
	CmdZrank               = newProtocolCommand("ZRANK")
	CmdZrevrank            = newProtocolCommand("ZREVRANK")
	CmdZrevrange           = newProtocolCommand("ZREVRANGE")
	CmdZcard               = newProtocolCommand("ZCARD")
	CmdZscore              = newProtocolCommand("ZSCORE")
	CmdMulti               = newProtocolCommand("MULTI")
	CmdDiscard             = newProtocolCommand("DISCARD")
	CmdExec                = newProtocolCommand("EXEC")
	CmdWatch               = newProtocolCommand("WATCH")
	CmdUnwatch             = newProtocolCommand("UNWATCH")
	CmdSort                = newProtocolCommand("SORT")
	CmdBlpop               = newProtocolCommand("BLPOP")
	CmdBrpop               = newProtocolCommand("BRPOP")
	CmdAuth                = newProtocolCommand("AUTH")
	CmdSubscribe           = newProtocolCommand("SUBSCRIBE")
	CmdPublish             = newProtocolCommand("PUBLISH")
	CmdUnsubscribe         = newProtocolCommand("UNSUBSCRIBE")
	CmdPsubscribe          = newProtocolCommand("PSUBSCRIBE")
	CmdPunsubscribe        = newProtocolCommand("PUNSUBSCRIBE")
	CmdPubsub              = newProtocolCommand("PUBSUB")
	CmdZcount              = newProtocolCommand("ZCOUNT")
	CmdZrangebyscore       = newProtocolCommand("ZRANGEBYSCORE")
	CmdZrevrangebyscore    = newProtocolCommand("ZREVRANGEBYSCORE")
	CmdZremrangebyrank     = newProtocolCommand("ZREMRANGEBYRANK")
	CmdZremrangebyscore    = newProtocolCommand("ZREMRANGEBYSCORE")
	CmdZunionstore         = newProtocolCommand("ZUNIONSTORE")
	CmdZinterstore         = newProtocolCommand("ZINTERSTORE")
	CmdZlexcount           = newProtocolCommand("ZLEXCOUNT")
	CmdZrangebylex         = newProtocolCommand("ZRANGEBYLEX")
	CmdZrevrangebylex      = newProtocolCommand("ZREVRANGEBYLEX")
	CmdZremrangebylex      = newProtocolCommand("ZREMRANGEBYLEX")
	CmdSave                = newProtocolCommand("SAVE")
	CmdBgsave              = newProtocolCommand("BGSAVE")
	CmdBgrewriteaof        = newProtocolCommand("BGREWRITEAOF")
	CmdLastsave            = newProtocolCommand("LASTSAVE")
	CmdShutdown            = newProtocolCommand("SHUTDOWN")
	CmdInfo                = newProtocolCommand("INFO")
	CmdMonitor             = newProtocolCommand("MONITOR")
	CmdSlaveof             = newProtocolCommand("SLAVEOF")
	CmdConfig              = newProtocolCommand("CONFIG")
	CmdStrlen              = newProtocolCommand("STRLEN")
	CmdSync                = newProtocolCommand("SYNC")
	CmdLpushx              = newProtocolCommand("LPUSHX")
	CmdPersist             = newProtocolCommand("PERSIST")
	CmdRpushx              = newProtocolCommand("RPUSHX")
	CmdEcho                = newProtocolCommand("ECHO")
	CmdLinsert             = newProtocolCommand("LINSERT")
	CmdDebug               = newProtocolCommand("DEBUG")
	CmdBrpoplpush          = newProtocolCommand("BRPOPLPUSH")
	CmdSetbit              = newProtocolCommand("SETBIT")
	CmdGetbit              = newProtocolCommand("GETBIT")
	CmdBitpos              = newProtocolCommand("BITPOS")
	CmdSetrange            = newProtocolCommand("SETRANGE")
	CmdGetrange            = newProtocolCommand("GETRANGE")
	CmdEval                = newProtocolCommand("EVAL")
	CmdEvalsha             = newProtocolCommand("EVALSHA")
	CmdScript              = newProtocolCommand("SCRIPT")
	CmdSlowlog             = newProtocolCommand("SLOWLOG")
	CmdObject              = newProtocolCommand("OBJECT")
	CmdBitcount            = newProtocolCommand("BITCOUNT")
	CmdBitop               = newProtocolCommand("BITOP")
	CmdSentinel            = newProtocolCommand("SENTINEL")
	CmdDump                = newProtocolCommand("DUMP")
	CmdRestore             = newProtocolCommand("RESTORE")
	CmdPexpire             = newProtocolCommand("PEXPIRE")
	CmdPexpireat           = newProtocolCommand("PEXPIREAT")
	CmdPttl                = newProtocolCommand("PTTL")
	CmdIncrbyfloat         = newProtocolCommand("INCRBYFLOAT")
	CmdPsetex              = newProtocolCommand("PSETEX")
	CmdClient              = newProtocolCommand("CLIENT")
	CmdTime                = newProtocolCommand("TIME")
	CmdMigrate             = newProtocolCommand("MIGRATE")
	CmdHincrbyfloat        = newProtocolCommand("HINCRBYFLOAT")
	CmdScan                = newProtocolCommand("SCAN")
	CmdHscan               = newProtocolCommand("HSCAN")
	CmdSscan               = newProtocolCommand("SSCAN")
	CmdZscan               = newProtocolCommand("ZSCAN")
	CmdWait                = newProtocolCommand("WAIT")
	CmdCluster             = newProtocolCommand("CLUSTER")
	CmdAsking              = newProtocolCommand("ASKING")
	CmdPfadd               = newProtocolCommand("PFADD")
	CmdPfcount             = newProtocolCommand("PFCOUNT")
	CmdPfmerge             = newProtocolCommand("PFMERGE")
	CmdReadonly            = newProtocolCommand("READONLY")
	CmdGeoadd              = newProtocolCommand("GEOADD")
	CmdGeodist             = newProtocolCommand("GEODIST")
	CmdGeohash             = newProtocolCommand("GEOHASH")
	CmdGeopos              = newProtocolCommand("GEOPOS")
	CmdGeoradius           = newProtocolCommand("GEORADIUS")
	CmdGeoradiusRo         = newProtocolCommand("GEORADIUS_RO")
	CmdGeoradiusbymember   = newProtocolCommand("GEORADIUSBYMEMBER")
	CmdGeoradiusbymemberRo = newProtocolCommand("GEORADIUSBYMEMBER_RO")
	CmdModule              = newProtocolCommand("MODULE")
	CmdBitfield            = newProtocolCommand("BITFIELD")
	CmdHstrlen             = newProtocolCommand("HSTRLEN")
	CmdTouch               = newProtocolCommand("TOUCH")
	CmdSwapdb              = newProtocolCommand("SWAPDB")
	CmdMemory              = newProtocolCommand("MEMORY")
	CmdXadd                = newProtocolCommand("XADD")
	CmdXlen                = newProtocolCommand("XLEN")
	CmdXdel                = newProtocolCommand("XDEL")
	CmdXtrim               = newProtocolCommand("XTRIM")
	CmdXrange              = newProtocolCommand("XRANGE")
	CmdXrevrange           = newProtocolCommand("XREVRANGE")
	CmdXread               = newProtocolCommand("XREAD")
	CmdXack                = newProtocolCommand("XACK")
	CmdXgroup              = newProtocolCommand("XGROUP")
	CmdXreadgroup          = newProtocolCommand("XREADGROUP")
	CmdXpending            = newProtocolCommand("XPENDING")
	CmdXclaim              = newProtocolCommand("XCLAIM")
)

// redis keyword
type keyword struct {
	Name string // name of keyword
}

// GetRaw byte array of name
func (k *keyword) GetRaw() []byte {
	return []byte(k.Name)
}

func newKeyword(name string) *keyword {
	return &keyword{name}
}

var (
	KeywordAggregate    = newKeyword("AGGREGATE")
	KeywordAlpha        = newKeyword("ALPHA")
	KeywordAsc          = newKeyword("ASC")
	KeywordBy           = newKeyword("BY")
	KeywordDesc         = newKeyword("DESC")
	KeywordGet          = newKeyword("GET")
	KeywordLimit        = newKeyword("LIMIT")
	KeywordMessage      = newKeyword("MESSAGE")
	KeywordNo           = newKeyword("NO")
	KeywordNosort       = newKeyword("NOSORT")
	KeywordPmessage     = newKeyword("PMESSAGE")
	KeywordPsubscribe   = newKeyword("PSUBSCRIBE")
	KeywordPunsubscribe = newKeyword("PUNSUBSCRIBE")
	KeywordOk           = newKeyword("OK")
	KeywordOne          = newKeyword("ONE")
	KeywordQueued       = newKeyword("QUEUED")
	KeywordSet          = newKeyword("SET")
	KeywordStore        = newKeyword("STORE")
	KeywordSubscribe    = newKeyword("SUBSCRIBE")
	KeywordUnsubscribe  = newKeyword("UNSUBSCRIBE")
	KeywordWeights      = newKeyword("WEIGHTS")
	KeywordWithscores   = newKeyword("WITHSCORES")
	KeywordResetstat    = newKeyword("RESETSTAT")
	KeywordRewrite      = newKeyword("REWRITE")
	KeywordReset        = newKeyword("RESET")
	KeywordFlush        = newKeyword("FLUSH")
	KeywordExists       = newKeyword("EXISTS")
	KeywordLoad         = newKeyword("LOAD")
	KeywordKill         = newKeyword("KILL")
	KeywordLen          = newKeyword("LEN")
	KeywordRefcount     = newKeyword("REFCOUNT")
	KeywordEncoding     = newKeyword("ENCODING")
	KeywordIdletime     = newKeyword("IDLETIME")
	KeywordGetname      = newKeyword("GETNAME")
	KeywordSetname      = newKeyword("SETNAME")
	KeywordList         = newKeyword("LIST")
	KeywordMatch        = newKeyword("MATCH")
	KeywordCount        = newKeyword("COUNT")
	KeywordPing         = newKeyword("PING")
	KeywordPong         = newKeyword("PONG")
	KeywordUnload       = newKeyword("UNLOAD")
	KeywordReplace      = newKeyword("REPLACE")
	KeywordKeys         = newKeyword("KEYS")
	KeywordPause        = newKeyword("PAUSE")
	KeywordDoctor       = newKeyword("DOCTOR")
	KeywordBlock        = newKeyword("BLOCK")
	KeywordNoack        = newKeyword("NOACK")
	KeywordStreams      = newKeyword("STREAMS")
	KeywordKey          = newKeyword("KEY")
	KeywordCreate       = newKeyword("CREATE")
	KeywordMkstream     = newKeyword("MKSTREAM")
	KeywordSetid        = newKeyword("SETID")
	KeywordDestroy      = newKeyword("DESTROY")
	KeywordDelconsumer  = newKeyword("DELCONSUMER")
	KeywordMaxlen       = newKeyword("MAXLEN")
	KeywordGroup        = newKeyword("GROUP")
	KeywordIdle         = newKeyword("IDLE")
	KeywordTime         = newKeyword("TIME")
	KeywordRetrycount   = newKeyword("RETRYCOUNT")
	KeywordForce        = newKeyword("FORCE")
)

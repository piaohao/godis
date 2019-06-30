package godis

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	askPrefix         = "ASK "
	movedPrefix       = "MOVED "
	clusterdownPrefix = "CLUSTERDOWN "
	busyPrefix        = "BUSY "
	noscriptPrefix    = "NOSCRIPT "

	defaultHost         = "localhost"
	defaultPort         = 6379
	defaultSentinelPort = 26379
	defaultTimeout      = 5 * time.Second
	defaultDatabase     = 2 * time.Second

	dollarByte   = '$'
	asteriskByte = '*'
	plusByte     = '+'
	minusByte    = '-'
	colonByte    = ':'

	sentinelMasters             = "masters"
	sentinelGetMasterAddrByName = "get-master-addr-by-name"
	sentinelReset               = "reset"
	sentinelSlaves              = "slaves"
	sentinelFailover            = "failover"
	sentinelMonitor             = "monitor"
	sentinelRemove              = "remove"
	sentinelSet                 = "set"

	clusterNodes            = "nodes"
	clusterMeet             = "meet"
	clusterReset            = "reset"
	clusterAddslots         = "addslots"
	clusterDelslots         = "delslots"
	clusterInfo             = "info"
	clusterGetkeysinslot    = "getkeysinslot"
	clusterSetslot          = "setslot"
	clusterSetslotNode      = "node"
	clusterSetslotMigrating = "migrating"
	clusterSetslotImporting = "importing"
	clusterSetslotStable    = "stable"
	clusterForget           = "forget"
	clusterFlushslot        = "flushslots"
	clusterKeyslot          = "keyslot"
	clusterCountkeyinslot   = "countkeysinslot"
	clusterSaveconfig       = "saveconfig"
	clusterReplicate        = "replicate"
	clusterSlaves           = "slaves"
	clusterFailover         = "failover"
	clusterSlots            = "slots"
	pubsubChannels          = "channels"
	pubsubNumsub            = "numsub"
	pubsubNumPat            = "numpat"
)

var (
	bytesTrue  = IntToByteArray(1)
	bytesFalse = IntToByteArray(0)
	bytesTilde = []byte("~")

	positiveInfinityBytes = []byte("+inf")
	negativeInfinityBytes = []byte("-inf")
)

const (
	maxUint = ^uint(0)
	minUint = 0
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

var (
	sizeTable = []int{9, 99, 999, 9999, 99999, 999999, 9999999, 99999999,
		999999999, maxInt}

	digitTens = []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1',
		'1', '1', '1', '1', '1', '1', '1', '1', '1', '2', '2', '2', '2', '2', '2', '2', '2', '2',
		'2', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4', '4', '4',
		'4', '4', '4', '5', '5', '5', '5', '5', '5', '5', '5', '5', '5', '6', '6', '6', '6', '6',
		'6', '6', '6', '6', '6', '7', '7', '7', '7', '7', '7', '7', '7', '7', '7', '8', '8', '8',
		'8', '8', '8', '8', '8', '8', '8', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9'}

	digitOnes = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
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
		r.buf[charPos] = digitOnes[p]
		charPos--
		r.buf[charPos] = digitTens[p]
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
	}

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
	value := int64(0)
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
	if err := p.os.writeByte(asteriskByte); err != nil {
		return err
	}
	if err := p.os.writeIntCrLf(len(args) + 1); err != nil {
		return err
	}
	if err := p.os.writeByte(dollarByte); err != nil {
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
		if err := p.os.writeByte(dollarByte); err != nil {
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
	case plusByte:
		return p.processStatusCodeReply()
	case dollarByte:
		return p.processBulkReply()
	case asteriskByte:
		return p.processMultiBulkReply()
	case colonByte:
		return p.processInteger()
	case minusByte:
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
	if strings.HasPrefix(msg, movedPrefix) {
		host, port, slot := p.parseTargetHostAndSlot(msg)
		return nil, newMovedDataError(msg, host, port, slot)
	} else if strings.HasPrefix(msg, askPrefix) {
		host, port, slot := p.parseTargetHostAndSlot(msg)
		return nil, newAskDataError(msg, host, port, slot)
	} else if strings.HasPrefix(msg, clusterdownPrefix) {
		return nil, newClusterError(msg)
	} else if strings.HasPrefix(msg, busyPrefix) {
		return nil, newBusyError(msg)
	} else if strings.HasPrefix(msg, noscriptPrefix) {
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
	cmdPing                = newProtocolCommand("PING")
	cmdSet                 = newProtocolCommand("SET")
	cmdGet                 = newProtocolCommand("GET")
	cmdQuit                = newProtocolCommand("QUIT")
	cmdExists              = newProtocolCommand("EXISTS")
	cmdDel                 = newProtocolCommand("DEL")
	cmdUnlink              = newProtocolCommand("UNLINK")
	cmdType                = newProtocolCommand("TYPE")
	cmdFlushdb             = newProtocolCommand("FLUSHDB")
	cmdKeys                = newProtocolCommand("KEYS")
	cmdRandomkey           = newProtocolCommand("RANDOMKEY")
	cmdRename              = newProtocolCommand("RENAME")
	cmdRenamenx            = newProtocolCommand("RENAMENX")
	cmdRenamex             = newProtocolCommand("RENAMEX")
	cmdDbsize              = newProtocolCommand("DBSIZE")
	cmdExpire              = newProtocolCommand("EXPIRE")
	cmdExpireat            = newProtocolCommand("EXPIREAT")
	cmdTTL                 = newProtocolCommand("TTL")
	cmdSelect              = newProtocolCommand("SELECT")
	cmdMove                = newProtocolCommand("MOVE")
	cmdFlushall            = newProtocolCommand("FLUSHALL")
	cmdGetset              = newProtocolCommand("GETSET")
	cmdMget                = newProtocolCommand("MGET")
	cmdSetnx               = newProtocolCommand("SETNX")
	cmdSetex               = newProtocolCommand("SETEX")
	cmdMset                = newProtocolCommand("MSET")
	cmdMsetnx              = newProtocolCommand("MSETNX")
	cmdDecrby              = newProtocolCommand("DECRBY")
	cmdDecr                = newProtocolCommand("DECR")
	cmdIncrby              = newProtocolCommand("INCRBY")
	cmdIncr                = newProtocolCommand("INCR")
	cmdAppend              = newProtocolCommand("APPEND")
	cmdSubstr              = newProtocolCommand("SUBSTR")
	cmdHset                = newProtocolCommand("HSET")
	cmdHget                = newProtocolCommand("HGET")
	cmdHsetnx              = newProtocolCommand("HSETNX")
	cmdHmset               = newProtocolCommand("HMSET")
	cmdHmget               = newProtocolCommand("HMGET")
	cmdHincrby             = newProtocolCommand("HINCRBY")
	cmdHexists             = newProtocolCommand("HEXISTS")
	cmdHdel                = newProtocolCommand("HDEL")
	cmdHlen                = newProtocolCommand("HLEN")
	cmdHkeys               = newProtocolCommand("HKEYS")
	cmdHvals               = newProtocolCommand("HVALS")
	cmdHgetall             = newProtocolCommand("HGETALL")
	cmdRpush               = newProtocolCommand("RPUSH")
	cmdLpush               = newProtocolCommand("LPUSH")
	cmdLlen                = newProtocolCommand("LLEN")
	cmdLrange              = newProtocolCommand("LRANGE")
	cmdLtrim               = newProtocolCommand("LTRIM")
	cmdLindex              = newProtocolCommand("LINDEX")
	cmdLset                = newProtocolCommand("LSET")
	cmdLrem                = newProtocolCommand("LREM")
	cmdLpop                = newProtocolCommand("LPOP")
	cmdRpop                = newProtocolCommand("RPOP")
	cmdRpoplpush           = newProtocolCommand("RPOPLPUSH")
	cmdSadd                = newProtocolCommand("SADD")
	cmdSmembers            = newProtocolCommand("SMEMBERS")
	cmdSrem                = newProtocolCommand("SREM")
	cmdSpop                = newProtocolCommand("SPOP")
	cmdSmove               = newProtocolCommand("SMOVE")
	cmdScard               = newProtocolCommand("SCARD")
	cmdSismember           = newProtocolCommand("SISMEMBER")
	cmdSinter              = newProtocolCommand("SINTER")
	cmdSinterstore         = newProtocolCommand("SINTERSTORE")
	cmdSunion              = newProtocolCommand("SUNION")
	cmdSunionstore         = newProtocolCommand("SUNIONSTORE")
	cmdSdiff               = newProtocolCommand("SDIFF")
	cmdSdiffstore          = newProtocolCommand("SDIFFSTORE")
	cmdSrandmember         = newProtocolCommand("SRANDMEMBER")
	cmdZadd                = newProtocolCommand("ZADD")
	cmdZrange              = newProtocolCommand("ZRANGE")
	cmdZrem                = newProtocolCommand("ZREM")
	cmdZincrby             = newProtocolCommand("ZINCRBY")
	cmdZrank               = newProtocolCommand("ZRANK")
	cmdZrevrank            = newProtocolCommand("ZREVRANK")
	cmdZrevrange           = newProtocolCommand("ZREVRANGE")
	cmdZcard               = newProtocolCommand("ZCARD")
	cmdZscore              = newProtocolCommand("ZSCORE")
	cmdMulti               = newProtocolCommand("MULTI")
	cmdDiscard             = newProtocolCommand("DISCARD")
	cmdExec                = newProtocolCommand("EXEC")
	cmdWatch               = newProtocolCommand("WATCH")
	cmdUnwatch             = newProtocolCommand("UNWATCH")
	cmdSort                = newProtocolCommand("SORT")
	cmdBlpop               = newProtocolCommand("BLPOP")
	cmdBrpop               = newProtocolCommand("BRPOP")
	cmdAuth                = newProtocolCommand("AUTH")
	cmdSubscribe           = newProtocolCommand("SUBSCRIBE")
	cmdPublish             = newProtocolCommand("PUBLISH")
	cmdUnsubscribe         = newProtocolCommand("UNSUBSCRIBE")
	cmdPsubscribe          = newProtocolCommand("PSUBSCRIBE")
	cmdPunsubscribe        = newProtocolCommand("PUNSUBSCRIBE")
	cmdPubsub              = newProtocolCommand("PUBSUB")
	cmdZcount              = newProtocolCommand("ZCOUNT")
	cmdZrangebyscore       = newProtocolCommand("ZRANGEBYSCORE")
	cmdZrevrangebyscore    = newProtocolCommand("ZREVRANGEBYSCORE")
	cmdZremrangebyrank     = newProtocolCommand("ZREMRANGEBYRANK")
	cmdZremrangebyscore    = newProtocolCommand("ZREMRANGEBYSCORE")
	cmdZunionstore         = newProtocolCommand("ZUNIONSTORE")
	cmdZinterstore         = newProtocolCommand("ZINTERSTORE")
	cmdZlexcount           = newProtocolCommand("ZLEXCOUNT")
	cmdZrangebylex         = newProtocolCommand("ZRANGEBYLEX")
	cmdZrevrangebylex      = newProtocolCommand("ZREVRANGEBYLEX")
	cmdZremrangebylex      = newProtocolCommand("ZREMRANGEBYLEX")
	cmdSave                = newProtocolCommand("SAVE")
	cmdBgsave              = newProtocolCommand("BGSAVE")
	cmdBgrewriteaof        = newProtocolCommand("BGREWRITEAOF")
	cmdLastsave            = newProtocolCommand("LASTSAVE")
	cmdShutdown            = newProtocolCommand("SHUTDOWN")
	cmdInfo                = newProtocolCommand("INFO")
	cmdMonitor             = newProtocolCommand("MONITOR")
	cmdSlaveof             = newProtocolCommand("SLAVEOF")
	cmdConfig              = newProtocolCommand("CONFIG")
	cmdStrlen              = newProtocolCommand("STRLEN")
	cmdSync                = newProtocolCommand("SYNC")
	cmdLpushx              = newProtocolCommand("LPUSHX")
	cmdPersist             = newProtocolCommand("PERSIST")
	cmdRpushx              = newProtocolCommand("RPUSHX")
	cmdEcho                = newProtocolCommand("ECHO")
	cmdLinsert             = newProtocolCommand("LINSERT")
	cmdDebug               = newProtocolCommand("DEBUG")
	cmdBrpoplpush          = newProtocolCommand("BRPOPLPUSH")
	cmdSetbit              = newProtocolCommand("SETBIT")
	cmdGetbit              = newProtocolCommand("GETBIT")
	cmdBitpos              = newProtocolCommand("BITPOS")
	cmdSetrange            = newProtocolCommand("SETRANGE")
	cmdGetrange            = newProtocolCommand("GETRANGE")
	cmdEval                = newProtocolCommand("EVAL")
	cmdEvalsha             = newProtocolCommand("EVALSHA")
	cmdScript              = newProtocolCommand("SCRIPT")
	cmdSlowlog             = newProtocolCommand("SLOWLOG")
	cmdObject              = newProtocolCommand("OBJECT")
	cmdBitcount            = newProtocolCommand("BITCOUNT")
	cmdBitop               = newProtocolCommand("BITOP")
	cmdSentinel            = newProtocolCommand("SENTINEL")
	cmdDump                = newProtocolCommand("DUMP")
	cmdRestore             = newProtocolCommand("RESTORE")
	cmdPexpire             = newProtocolCommand("PEXPIRE")
	cmdPexpireat           = newProtocolCommand("PEXPIREAT")
	cmdPttl                = newProtocolCommand("PTTL")
	cmdIncrbyfloat         = newProtocolCommand("INCRBYFLOAT")
	cmdPsetex              = newProtocolCommand("PSETEX")
	cmdClient              = newProtocolCommand("CLIENT")
	cmdTime                = newProtocolCommand("TIME")
	cmdMigrate             = newProtocolCommand("MIGRATE")
	cmdHincrbyfloat        = newProtocolCommand("HINCRBYFLOAT")
	cmdScan                = newProtocolCommand("SCAN")
	cmdHscan               = newProtocolCommand("HSCAN")
	cmdSscan               = newProtocolCommand("SSCAN")
	cmdZscan               = newProtocolCommand("ZSCAN")
	cmdWait                = newProtocolCommand("WAIT")
	cmdCluster             = newProtocolCommand("CLUSTER")
	cmdAsking              = newProtocolCommand("ASKING")
	cmdPfadd               = newProtocolCommand("PFADD")
	cmdPfcount             = newProtocolCommand("PFCOUNT")
	cmdPfmerge             = newProtocolCommand("PFMERGE")
	cmdReadonly            = newProtocolCommand("READONLY")
	cmdGeoadd              = newProtocolCommand("GEOADD")
	cmdGeodist             = newProtocolCommand("GEODIST")
	cmdGeohash             = newProtocolCommand("GEOHASH")
	cmdGeopos              = newProtocolCommand("GEOPOS")
	cmdGeoradius           = newProtocolCommand("GEORADIUS")
	cmdGeoradiusRo         = newProtocolCommand("GEORADIUS_RO")
	cmdGeoradiusbymember   = newProtocolCommand("GEORADIUSBYMEMBER")
	cmdGeoradiusbymemberRo = newProtocolCommand("GEORADIUSBYMEMBER_RO")
	cmdModule              = newProtocolCommand("MODULE")
	cmdBitfield            = newProtocolCommand("BITFIELD")
	cmdHstrlen             = newProtocolCommand("HSTRLEN")
	cmdTouch               = newProtocolCommand("TOUCH")
	cmdSwapdb              = newProtocolCommand("SWAPDB")
	cmdMemory              = newProtocolCommand("MEMORY")
	cmdXadd                = newProtocolCommand("XADD")
	cmdXlen                = newProtocolCommand("XLEN")
	cmdXdel                = newProtocolCommand("XDEL")
	cmdXtrim               = newProtocolCommand("XTRIM")
	cmdXrange              = newProtocolCommand("XRANGE")
	cmdXrevrange           = newProtocolCommand("XREVRANGE")
	cmdXread               = newProtocolCommand("XREAD")
	cmdXack                = newProtocolCommand("XACK")
	cmdXgroup              = newProtocolCommand("XGROUP")
	cmdXreadgroup          = newProtocolCommand("XREADGROUP")
	cmdXpending            = newProtocolCommand("XPENDING")
	cmdXclaim              = newProtocolCommand("XCLAIM")
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
	keywordAggregate    = newKeyword("AGGREGATE")
	keywordAlpha        = newKeyword("ALPHA")
	keywordAsc          = newKeyword("ASC")
	keywordBy           = newKeyword("BY")
	keywordDesc         = newKeyword("DESC")
	keywordGet          = newKeyword("GET")
	keywordLimit        = newKeyword("LIMIT")
	keywordMessage      = newKeyword("MESSAGE")
	keywordNo           = newKeyword("NO")
	keywordNosort       = newKeyword("NOSORT")
	keywordPmessage     = newKeyword("PMESSAGE")
	keywordPsubscribe   = newKeyword("PSUBSCRIBE")
	keywordPunsubscribe = newKeyword("PUNSUBSCRIBE")
	keywordOk           = newKeyword("OK")
	keywordOne          = newKeyword("ONE")
	keywordQueued       = newKeyword("QUEUED")
	keywordSet          = newKeyword("SET")
	keywordStore        = newKeyword("STORE")
	keywordSubscribe    = newKeyword("SUBSCRIBE")
	keywordUnsubscribe  = newKeyword("UNSUBSCRIBE")
	keywordWeights      = newKeyword("WEIGHTS")
	keywordWithscores   = newKeyword("WITHSCORES")
	keywordResetstat    = newKeyword("RESETSTAT")
	keywordRewrite      = newKeyword("REWRITE")
	keywordReset        = newKeyword("RESET")
	keywordFlush        = newKeyword("FLUSH")
	keywordExists       = newKeyword("EXISTS")
	keywordLoad         = newKeyword("LOAD")
	keywordKill         = newKeyword("KILL")
	keywordLen          = newKeyword("LEN")
	keywordRefcount     = newKeyword("REFCOUNT")
	keywordEncoding     = newKeyword("ENCODING")
	keywordIdletime     = newKeyword("IDLETIME")
	keywordGetname      = newKeyword("GETNAME")
	keywordSetname      = newKeyword("SETNAME")
	keywordList         = newKeyword("LIST")
	keywordMatch        = newKeyword("MATCH")
	keywordCount        = newKeyword("COUNT")
	keywordPing         = newKeyword("PING")
	keywordPong         = newKeyword("PONG")
	keywordUnload       = newKeyword("UNLOAD")
	keywordReplace      = newKeyword("REPLACE")
	keywordKeys         = newKeyword("KEYS")
	keywordPause        = newKeyword("PAUSE")
	keywordDoctor       = newKeyword("DOCTOR")
	keywordBlock        = newKeyword("BLOCK")
	keywordNoack        = newKeyword("NOACK")
	keywordStreams      = newKeyword("STREAMS")
	keywordKey          = newKeyword("KEY")
	keywordCreate       = newKeyword("CREATE")
	keywordMkstream     = newKeyword("MKSTREAM")
	keywordSetid        = newKeyword("SETID")
	keywordDestroy      = newKeyword("DESTROY")
	keywordDelconsumer  = newKeyword("DELCONSUMER")
	keywordMaxlen       = newKeyword("MAXLEN")
	keywordGroup        = newKeyword("GROUP")
	keywordIdle         = newKeyword("IDLE")
	keywordTime         = newKeyword("TIME")
	keywordRetrycount   = newKeyword("RETRYCOUNT")
	keywordForce        = newKeyword("FORCE")
)

package godis

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	Host              string
	Port              int
	Socket            net.Conn
	Protocol          *Protocol
	ConnectionTimeout int
	SoTimeout         int
	Broken            bool
	Ssl               bool

	pipelinedCommands int
}

func NewConnection(host string, port, connectionTimeout, soTimeout int, ssl bool) *Connection {
	if host == "" {
		host = DEFAULT_HOST
	}
	if port == 0 {
		port = DEFAULT_PORT
	}
	if connectionTimeout == 0 {
		connectionTimeout = DEFAULT_TIMEOUT
	}
	if soTimeout == 0 {
		soTimeout = DEFAULT_TIMEOUT
	}
	return &Connection{
		Host:              host,
		Port:              port,
		ConnectionTimeout: connectionTimeout,
		SoTimeout:         soTimeout,
		Broken:            false,
		Ssl:               ssl,
	}
}

func (c *Connection) setTimeoutInfinite() error {
	err := c.Socket.SetDeadline(time.Time{})
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *Connection) rollbackTimeout() error {
	err := c.Socket.SetDeadline(time.Now().Add(time.Duration(c.ConnectionTimeout) * time.Second))
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *Connection) resetPipelinedCount() {
	c.pipelinedCommands = 0
}

func (c *Connection) SendCommand(cmd protocolCommand, args ...[]byte) error {
	if err := c.Protocol.sendCommand(cmd.GetRaw(), args...); err != nil {
		return err
	}
	return nil
}

func (c *Connection) readProtocolWithCheckingBroken() (interface{}, error) {
	if c.Broken {
		return nil, errors.New("attempting to read from a broken connection")
	}
	read, err := c.Protocol.read()
	//todo	need distinguish error, when error is redis connection exception ,then set broken with true
	if err != nil {
		c.Broken = true
	}
	return read, err
}

func (c *Connection) getStatusCodeReply() (string, error) {
	reply, err := c.getOne()
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", nil
	}
	switch t := reply.(type) {
	case keyword:
		return string(t.GetRaw()), nil
	case string:
		return t, nil
	default:
		return "", errors.New("internal error")
	}
}

func (c *Connection) getBulkReply() (string, error) {
	result, err := c.getBinaryBulkReply()
	if nil != err {
		return string(result), nil
	} else {
		return "", err
	}
}

func (c *Connection) getBinaryBulkReply() ([]byte, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []byte{}, nil
	}
	resp := reply.([]interface{})
	respArr := make([]byte, 0)
	for _, r := range resp {
		respArr = append(respArr, r.(byte))
	}
	return respArr, nil
}

func (c *Connection) getIntegerReply() (int64, error) {
	reply, err := c.getOne()
	if err != nil {
		return 0, err
	}
	if reply == nil {
		return 0, nil
	}
	resp := reply.(int64)
	return resp, nil
}

func (c *Connection) getMultiBulkReply() ([]string, error) {
	reply, err := c.getBinaryMultiBulkReply()
	if err != nil {
		return nil, err
	}
	resp := make([]string, 0)
	for _, r := range reply {
		resp = append(resp, string(r))
	}
	return resp, nil
}

func (c *Connection) getBinaryMultiBulkReply() ([][]byte, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return [][]byte{}, nil
	}
	resp := reply.([]interface{})
	arr := make([][]byte, 0)
	for _, res := range resp {
		arr = append(arr, res.([]byte))
	}
	return arr, nil
}

func (c *Connection) getUnflushedObjectMultiBulkReply() ([]interface{}, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []interface{}{}, nil
	}
	return reply.([]interface{}), nil
}

func (c *Connection) getRawObjectMultiBulkReply() ([]interface{}, error) {
	reply, err := c.readProtocolWithCheckingBroken()
	return reply.([]interface{}), err
}

func (c *Connection) getObjectMultiBulkReply() ([]interface{}, error) {
	return c.getUnflushedObjectMultiBulkReply()
}

func (c *Connection) getIntegerMultiBulkReply() ([]int64, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []int64{}, nil
	}
	return reply.([]int64), nil
}

func (c *Connection) getOne() (interface{}, error) {
	if err := c.flush(); err != nil {
		return "", err
	}
	return c.readProtocolWithCheckingBroken()
}

func (c *Connection) flush() error {
	err := c.Protocol.os.Flush()
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *Connection) Connect() error {
	conn, err := net.Dial("tcp", fmt.Sprint(c.Host, ":", c.Port))
	if err != nil {
		return err
	}
	err = conn.SetDeadline(time.Now().Add(time.Duration(c.ConnectionTimeout) * time.Second))
	if err != nil {
		return err
	}
	c.Socket = conn
	c.Protocol = NewProtocol(NewRedisOutputStream(bufio.NewWriter(conn)), NewRedisInputStream(bufio.NewReader(conn)))
	return nil
}

func (c *Connection) IsConnected() {
}

func (c *Connection) Close() error {
	return c.Socket.Close()
}

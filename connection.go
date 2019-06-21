package godis

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

type connection struct {
	Host              string
	Port              int
	Socket            net.Conn
	Protocol          *protocol
	ConnectionTimeout int
	SoTimeout         int
	Broken            bool
	Ssl               bool

	pipelinedCommands int
}

func newConnection(host string, port, connectionTimeout, soTimeout int) *connection {
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
	return &connection{
		Host:              host,
		Port:              port,
		ConnectionTimeout: connectionTimeout,
		SoTimeout:         soTimeout,
		Broken:            false,
	}
}

func (c *connection) setTimeoutInfinite() error {
	err := c.Socket.SetDeadline(time.Time{})
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *connection) rollbackTimeout() error {
	err := c.Socket.SetDeadline(time.Now().Add(time.Duration(c.ConnectionTimeout) * time.Second))
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *connection) resetPipelinedCount() {
	c.pipelinedCommands = 0
}

func (c *connection) SendCommand(cmd protocolCommand, args ...[]byte) error {
	err := c.connect()
	if err != nil {
		return err
	}
	if err := c.Protocol.sendCommand(cmd.GetRaw(), args...); err != nil {
		return err
	}
	c.pipelinedCommands++
	return nil
}

func (c *connection) readProtocolWithCheckingBroken() (interface{}, error) {
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

func (c *connection) getStatusCodeReply() (string, error) {
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

func (c *connection) getBulkReply() (string, error) {
	result, err := c.getBinaryBulkReply()
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (c *connection) getBinaryBulkReply() ([]byte, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []byte{}, nil
	}
	resp := reply.([]byte)
	return resp, nil
}

func (c *connection) getIntegerReply() (int64, error) {
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

func (c *connection) getMultiBulkReply() ([]string, error) {
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

func (c *connection) getBinaryMultiBulkReply() ([][]byte, error) {
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

func (c *connection) getUnflushedObjectMultiBulkReply() ([]interface{}, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []interface{}{}, nil
	}
	return reply.([]interface{}), nil
}

func (c *connection) getRawObjectMultiBulkReply() ([]interface{}, error) {
	reply, err := c.readProtocolWithCheckingBroken()
	return reply.([]interface{}), err
}

func (c *connection) getObjectMultiBulkReply() ([]interface{}, error) {
	return c.getUnflushedObjectMultiBulkReply()
}

func (c *connection) getIntegerMultiBulkReply() ([]int64, error) {
	reply, err := c.getOne()
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return []int64{}, nil
	}
	return reply.([]int64), nil
}

func (c *connection) getOne() (interface{}, error) {
	if err := c.flush(); err != nil {
		return "", err
	}
	c.pipelinedCommands--
	return c.readProtocolWithCheckingBroken()
}

func (c *connection) getAll(expect ...int) (interface{}, error) {
	num := 0
	if len(expect) > 0 {
		num = expect[0]
	}
	if err := c.flush(); err != nil {
		return nil, err
	}
	all := make([]interface{}, 0)
	for c.pipelinedCommands > num {
		obj, err := c.readProtocolWithCheckingBroken()
		if err != nil {
			all = append(all, err)
		} else {
			all = append(all, obj)
		}
		c.pipelinedCommands--
	}
	return all, nil
}

func (c *connection) flush() error {
	err := c.Protocol.os.Flush()
	if err != nil {
		c.Broken = true
		return err
	}
	return nil
}

func (c *connection) connect() error {
	if c.isConnected() {
		return nil
	}
	conn, err := net.Dial("tcp", fmt.Sprint(c.Host, ":", c.Port))
	if err != nil {
		return err
	}
	err = conn.SetDeadline(time.Now().Add(time.Duration(c.ConnectionTimeout) * time.Second))
	if err != nil {
		return err
	}
	c.Socket = conn
	c.Protocol = newProtocol(newRedisOutputStream(bufio.NewWriter(c.Socket)), newRedisInputStream(bufio.NewReader(c.Socket)))
	return nil
}

func (c *connection) isConnected() bool {
	if c.Socket == nil {
		return false
	}
	return true
}

func (c *connection) close() error {
	if c.Socket == nil {
		return nil
	}
	return c.Socket.Close()
}

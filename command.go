package godis

import (
	"errors"
	"fmt"
	"strings"
)

// RedisError basic redis error
type RedisError struct {
	Message string
}

func NewRedisError(message string) *RedisError {
	return &RedisError{Message: message}
}

func (e *RedisError) Error() string {
	return e.Message
}

type NoReachableClusterNodeError struct {
	Message string
}

func NewNoReachableClusterNodeError(message string) *NoReachableClusterNodeError {
	return &NoReachableClusterNodeError{Message: message}
}

func (e *NoReachableClusterNodeError) Error() string {
	return e.Message
}

type MovedDataError struct {
	Message string
	Host    string
	Port    int
	Slot    int
}

func NewMovedDataError(message string, host string, port int, slot int) *MovedDataError {
	return &MovedDataError{Message: message, Host: host, Port: port, Slot: slot}
}

func (e *MovedDataError) Error() string {
	return e.Message
}

type AskDataError struct {
	Message string
	Host    string
	Port    int
	Slot    int
}

func NewAskDataError(message string, host string, port int, slot int) *AskDataError {
	return &AskDataError{Message: message, Host: host, Port: port, Slot: slot}
}

func (e *AskDataError) Error() string {
	return e.Message
}

type ClusterError struct {
	Message string
}

func NewClusterError(message string) *ClusterError {
	return &ClusterError{Message: message}
}

func (e *ClusterError) Error() string {
	return e.Message
}

type BusyError struct {
	Message string
}

func NewBusyError(message string) *BusyError {
	return &BusyError{Message: message}
}

func (e *BusyError) Error() string {
	return e.Message
}

type NoScriptError struct {
	Message string
}

func NewNoScriptError(message string) *NoScriptError {
	return &NoScriptError{Message: message}
}

func (e *NoScriptError) Error() string {
	return e.Message
}

type DataError struct {
	Message string
}

func NewDataError(message string) *DataError {
	return &DataError{Message: message}
}

func (e *DataError) Error() string {
	return e.Message
}

type ConnectError struct {
	Message string
}

func NewConnectError(message string) *ConnectError {
	return &ConnectError{Message: message}
}

func (e *ConnectError) Error() string {
	return e.Message
}

//ZAddParams ...
type ZAddParams struct {
	XX     bool
	NX     bool
	CH     bool
	params map[string]string
}

//GetByteParams ...
func (p ZAddParams) GetByteParams(key []byte, args ...[]byte) [][]byte {
	arr := make([][]byte, 0)
	arr = append(arr, key)
	if p.Contains("XX") {
		arr = append(arr, []byte("XX"))
	}
	if p.Contains("NX") {
		arr = append(arr, []byte("NX"))
	}
	if p.Contains("CH") {
		arr = append(arr, []byte("CH"))
	}
	for _, a := range args {
		arr = append(arr, a)
	}
	return arr
}

//Contains ...
func (p ZAddParams) Contains(key string) bool {
	_, ok := p.params[key]
	return ok
}

//BitPosParams
type BitPosParams struct {
	params [][]byte
}

//SortingParams
type SortingParams struct {
	params [][]byte
}

//ScanParams
type ScanParams struct {
	params map[keyword][]byte
}

//NewScanParams ...
func NewScanParams() *ScanParams {
	return &ScanParams{params: make(map[keyword][]byte)}
}

//GetParams ...
func (s ScanParams) GetParams() [][]byte {
	arr := make([][]byte, 0)
	for k, v := range s.params {
		arr = append(arr, k.GetRaw())
		arr = append(arr, []byte(v))
	}
	return arr
}

//Match ...
func (s ScanParams) Match() string {
	if v, ok := s.params[KeywordMatch]; !ok {
		return ""
	} else {
		return string(v)
	}
}

//Count ...
func (s ScanParams) Count() int {
	if v, ok := s.params[KeywordCount]; !ok {
		return 0
	} else {
		return int(ByteArrayToInt64(v))
	}
}

//ListOption ...
type ListOption struct {
	// name  ...
	Name string
}

//GetRaw ...
func (l ListOption) GetRaw() []byte {
	return []byte(l.Name)
}

//NewListOption ...
func NewListOption(name string) ListOption {
	return ListOption{name}
}

var (
	ListoptionBefore = NewListOption("BEFORE")
	ListoptionAfter  = NewListOption("AFTER")
)

//GeoUnit
type GeoUnit struct {
	// name ...
	Name string
}

//GetRaw ...
func (g GeoUnit) GetRaw() []byte {
	return []byte(g.Name)
}

//NewGeoUnit ...
func NewGeoUnit(name string) GeoUnit {
	return GeoUnit{name}
}

var (
	GeounitMi = NewGeoUnit("MI")
	GeounitM  = NewGeoUnit("M")
	GeounitKm = NewGeoUnit("KM")
	GeounitFt = NewGeoUnit("FT")
)

//GeoRadiusParam
type GeoRadiusParam struct {
	params map[string]interface{}
}

//GetParams ...
func (g GeoRadiusParam) GetParams(args [][]byte) [][]byte {
	arr := make([][]byte, 0)
	for _, a := range args {
		arr = append(arr, a)
	}

	if g.Contains("WITHCOORD") {
		arr = append(arr, []byte("WITHCOORD"))
	}
	if g.Contains("WITHDIST") {
		arr = append(arr, []byte("WITHDIST"))
	}

	if g.Contains("COUNT") {
		arr = append(arr, []byte("COUNT"))
		arr = append(arr, IntToByteArray(g.params["COUNT"].(int)))
	}

	if g.Contains("ASC") {
		arr = append(arr, []byte("ASC"))
	} else if g.Contains("DESC") {
		arr = append(arr, []byte("DESC"))
	}

	return arr
}

//Contains ...
func (g GeoRadiusParam) Contains(key string) bool {
	_, ok := g.params[key]
	return ok
}

//Tuple ...
type Tuple struct {
	element []byte
	score   float64
}

//GeoRadiusResponse ...
type GeoRadiusResponse struct {
	member     []byte
	distance   float64
	coordinate GeoCoordinate
}

//GeoCoordinate ...
type GeoCoordinate struct {
	longitude float64
	latitude  float64
}

//ScanResult ...
type ScanResult struct {
	Cursor  string
	Results []string
}

//ZParams ...
type ZParams struct {
	//name  ...
	Name   string
	params [][]byte
}

//GetRaw ...
func (g ZParams) GetRaw() []byte {
	return []byte(g.Name)
}

//GetParams ...
func (g ZParams) GetParams() [][]byte {
	return g.params
}

//NewZParams ...
func NewZParams(name string) ZParams {
	return ZParams{Name: name}
}

var (
	ZparamsSum = NewZParams("SUM")
	ZparamsMin = NewZParams("MIN")
	ZparamsMax = NewZParams("MAX")
)

//RedisPubSub ...
type RedisPubSub struct {
	subscribedChannels int
	redis              *Redis
	OnMessage          func(channel, message string)
	OnPMessage         func(pattern string, channel, message string)
	OnSubscribe        func(channel string, subscribedChannels int)
	OnUnsubscribe      func(channel string, subscribedChannels int)
	OnPUnsubscribe     func(pattern string, subscribedChannels int)
	OnPSubscribe       func(pattern string, subscribedChannels int)
	OnPong             func(channel string)
}

//Subscribe ...
func (r *RedisPubSub) Subscribe(channels ...string) error {
	if r.redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.redis.client.subscribe(channels...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Unsubscribe ...
func (r *RedisPubSub) Unsubscribe(channels ...string) error {
	if r.redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.redis.client.unsubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Psubscribe ...
func (r *RedisPubSub) Psubscribe(channels ...string) error {
	if r.redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.redis.client.psubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

//Punsubscribe ...
func (r *RedisPubSub) Punsubscribe(channels ...string) error {
	if r.redis.client == nil {
		return errors.New("redisPubSub is not subscribed to a Redis instance")
	}
	err := r.redis.client.punsubscribe(channels...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisPubSub) proceed(redis *Redis, channels ...string) error {
	r.redis = redis
	err := r.redis.client.subscribe(channels...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return r.process(redis)
}

func (r *RedisPubSub) isSubscribed() bool {
	return r.subscribedChannels > 0
}

func (r *RedisPubSub) proceedWithPatterns(redis *Redis, patterns ...string) error {
	r.redis = redis
	err := r.redis.client.psubscribe(patterns...)
	if err != nil {
		return err
	}
	err = r.redis.client.flush()
	if err != nil {
		return err
	}
	return r.process(redis)
}

func (r *RedisPubSub) process(redis *Redis) error {
	for {
		reply, err := redis.client.connection.getRawObjectMultiBulkReply()
		if err != nil {
			return err
		}
		respUpper := strings.ToUpper(string(reply[0].([]byte)))
		switch respUpper {
		case KeywordSubscribe.Name:
			r.processSubscribe(reply)
		case KeywordUnsubscribe.Name:
			r.processUnSubscribe(reply)
		case KeywordMessage.Name:
			r.processMessage(reply)
		case KeywordPmessage.Name:
			r.processPmessage(reply)
		case KeywordPsubscribe.Name:
			r.processPsubscribe(reply)
		case CmdPunsubscribe.Name:
			r.processPunsubcribe(reply)
		case KeywordPong.Name:
			r.processPong(reply)
		default:
			return errors.New(fmt.Sprintf("Unknown message type: %v", reply))
		}
		if !r.isSubscribed() {
			break
		}
	}
	// Reset pipeline count because subscribe() calls would have increased it but nothing decremented it.
	redis.client.resetPipelinedCount()
	// Invalidate instance since this thread is no longer listening
	r.redis.client = nil
	return nil
}

func (r *RedisPubSub) processSubscribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bChannel := reply[1].([]byte)
	strChannel := ""
	if bChannel != nil {
		strChannel = string(bChannel)
	}
	r.OnSubscribe(strChannel, r.subscribedChannels)
}

func (r *RedisPubSub) processUnSubscribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bChannel := reply[1].([]byte)
	strChannel := ""
	if bChannel != nil {
		strChannel = string(bChannel)
	}
	r.OnUnsubscribe(strChannel, r.subscribedChannels)
}

func (r *RedisPubSub) processMessage(reply []interface{}) {
	bChannel := reply[1].([]byte)
	bMsg := reply[2].([]byte)
	strChannel := ""
	if bChannel != nil {
		strChannel = string(bChannel)
	}
	strMsg := ""
	if bChannel != nil {
		strMsg = string(bMsg)
	}
	r.OnMessage(strChannel, strMsg)
}

func (r *RedisPubSub) processPmessage(reply []interface{}) {
	bPattern := reply[1].([]byte)
	bChannel := reply[2].([]byte)
	bMsg := reply[31].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	strChannel := ""
	if bChannel != nil {
		strChannel = string(bChannel)
	}
	strMsg := ""
	if bChannel != nil {
		strMsg = string(bMsg)
	}
	r.OnPMessage(strPattern, strChannel, strMsg)
}

func (r *RedisPubSub) processPsubscribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPSubscribe(strPattern, r.subscribedChannels)
}

func (r *RedisPubSub) processPunsubcribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPUnsubscribe(strPattern, r.subscribedChannels)
}

func (r *RedisPubSub) processPong(reply []interface{}) {
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPong(strPattern)
}

//BitOP ...
type BitOP struct {
	//name  ...
	Name string
}

//GetRaw ...
func (g BitOP) GetRaw() []byte {
	return []byte(g.Name)
}

//NewBitOP ...
func NewBitOP(name string) BitOP {
	return BitOP{name}
}

var (
	BitopAnd = NewBitOP("AND")
	BitopOr  = NewBitOP("OR")
	BitopXor = NewBitOP("XOR")
	BitopNot = NewBitOP("NOT")
)

//Slowlog ...
type Slowlog struct {
	id            int64
	timeStamp     int64
	executionTime int64
	args          []string
}

//DebugParams ...
type DebugParams struct {
	command []string
}

func NewDebugParamsSegfault() *DebugParams {
	return &DebugParams{command: []string{"SEGFAULT"}}
}

func NewDebugParamsObject(key string) *DebugParams {
	return &DebugParams{command: []string{"OBJECT", key}}
}

func NewDebugParamsReload() *DebugParams {
	return &DebugParams{command: []string{"RELOAD"}}
}

//Reset ...
type Reset struct {
	//name  ...
	Name string
}

//GetRaw ...
func (g Reset) GetRaw() []byte {
	return []byte(g.Name)
}

//NewReset ...
func NewReset(name string) Reset {
	return Reset{name}
}

var (
	ResetSoft = NewReset("SOFT")
	ResetHard = NewReset("HARD")
)

package godis

import (
	"fmt"
	"strconv"
	"strings"
)

//ZAddParams ...
type ZAddParams struct {
	params map[string]string
}

//NewZAddParams constructor
func NewZAddParams() *ZAddParams {
	return &ZAddParams{params: make(map[string]string)}
}

//XX set XX parameter, Only update elements that already exist. Never add elements.
func (p *ZAddParams) XX() *ZAddParams {
	p.params["XX"] = "XX"
	return p
}

//NX set NX parameter, Don't update already existing elements. Always add new elements.
func (p *ZAddParams) NX() *ZAddParams {
	p.params["NX"] = "NX"
	return p
}

//CH set CH parameter, Modify the return value from the number of new elements added, to the total number of elements changed
func (p *ZAddParams) CH() *ZAddParams {
	p.params["CH"] = "CH"
	return p
}

//getByteParams get all params
func (p *ZAddParams) getByteParams(key []byte, args ...[]byte) [][]byte {
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

//Contains return params map contains the key
func (p *ZAddParams) Contains(key string) bool {
	_, ok := p.params[key]
	return ok
}

//BitPosParams bitpos params
type BitPosParams struct {
	params [][]byte
}

//SortParams sort params
type SortParams struct {
	params []string
}

//NewSortParams create new sort params instance
func NewSortParams() *SortParams {
	return &SortParams{params: make([]string, 0)}
}

func (p *SortParams) getParams() [][]byte {
	return StrArrToByteArrArr(p.params)
}

//By set by param with pattern
func (p *SortParams) By(pattern string) *SortParams {
	p.params = append(p.params, keywordBy.name)
	p.params = append(p.params, pattern)
	return p
}

//NoSort set by param with nosort
func (p *SortParams) NoSort() *SortParams {
	p.params = append(p.params, keywordBy.name)
	p.params = append(p.params, keywordNosort.name)
	return p
}

//Desc set desc param,then sort elements in descending order
func (p *SortParams) Desc() *SortParams {
	p.params = append(p.params, keywordDesc.name)
	return p
}

//Asc set asc param,then sort elements in ascending order
func (p *SortParams) Asc() *SortParams {
	p.params = append(p.params, keywordAsc.name)
	return p
}

//Limit limit the sort result,[x,y)
func (p *SortParams) Limit(start, count int) *SortParams {
	p.params = append(p.params, keywordLimit.name)
	p.params = append(p.params, strconv.Itoa(start))
	p.params = append(p.params, strconv.Itoa(count))
	return p
}

//Alpha sort elements in alpha order
func (p *SortParams) Alpha() *SortParams {
	p.params = append(p.params, keywordAlpha.name)
	return p
}

//Get set get param with patterns
func (p *SortParams) Get(patterns ...string) *SortParams {
	for _, pattern := range patterns {
		p.params = append(p.params, keywordGet.name)
		p.params = append(p.params, pattern)
	}
	return p
}

//ScanParams scan,hscan,sscan,zscan params
type ScanParams struct {
	//params map[*keyword][]byte
	params map[string]string
}

//NewScanParams create scan params instance
func NewScanParams() *ScanParams {
	return &ScanParams{params: make(map[string]string)}
}

//Match scan match pattern
func (s *ScanParams) Match(pattern string) *ScanParams {
	s.params[keywordMatch.name] = pattern
	return s
}

//Count scan result count
func (s *ScanParams) Count(count int) *ScanParams {
	s.params[keywordCount.name] = strconv.Itoa(count)
	return s
}

//getParams get all scan params
func (s ScanParams) getParams() [][]byte {
	arr := make([][]byte, 0)
	for k, v := range s.params {
		arr = append(arr, []byte(k))
		arr = append(arr, []byte(v))
	}
	return arr
}

//GetMatch get the match param value
func (s ScanParams) GetMatch() string {
	if v, ok := s.params[keywordMatch.name]; ok {
		return v
	}
	return ""
}

//ListOption  list option
type ListOption struct {
	name string // name  ...
}

//getRaw get the option name byte array
func (l *ListOption) getRaw() []byte {
	return []byte(l.name)
}

//NewListOption create new list option instance
func newListOption(name string) *ListOption {
	return &ListOption{name}
}

var (
	//ListOptionBefore insert an new element before designated element
	ListOptionBefore = newListOption("BEFORE")
	//ListOptionAfter insert an new element after designated element
	ListOptionAfter = newListOption("AFTER")
)

//GeoUnit geo unit,m|mi|km|ft
type GeoUnit struct {
	name string // name of geo unit
}

//getRaw get the name byte array
func (g *GeoUnit) getRaw() []byte {
	return []byte(g.name)
}

//NewGeoUnit create a new geounit instance
func newGeoUnit(name string) *GeoUnit {
	return &GeoUnit{name}
}

var (
	//GeoUnitMi calculate distance use mi unit
	GeoUnitMi = newGeoUnit("mi")
	//GeoUnitM calculate distance use m unit
	GeoUnitM = newGeoUnit("m")
	//GeoUnitKm calculate distance use km unit
	GeoUnitKm = newGeoUnit("km")
	//GeoUnitFt calculate distance use ft unit
	GeoUnitFt = newGeoUnit("ft")
)

//GeoRadiusParams geo radius param
type GeoRadiusParams struct {
	params map[string]string
}

//NewGeoRadiusParam create a new geo radius param instance
func NewGeoRadiusParam() *GeoRadiusParams {
	return &GeoRadiusParams{params: make(map[string]string)}
}

//WithCoord fill the geo result with coordinate
func (p *GeoRadiusParams) WithCoord() *GeoRadiusParams {
	p.params["withcoord"] = "withcoord"
	return p
}

//WithDist fill the geo result with distance
func (p *GeoRadiusParams) WithDist() *GeoRadiusParams {
	p.params["withdist"] = "withdist"
	return p
}

//SortAscending sort th geo result in ascending order
func (p *GeoRadiusParams) SortAscending() *GeoRadiusParams {
	p.params["asc"] = "asc"
	return p
}

//SortDescending sort the geo result in descending order
func (p *GeoRadiusParams) SortDescending() *GeoRadiusParams {
	p.params["desc"] = "desc"
	return p
}

//Count fill the geo result with count
func (p *GeoRadiusParams) Count(count int) *GeoRadiusParams {
	if count > 0 {
		p.params["count"] = strconv.Itoa(count)
	}
	return p
}

//getParams  get geo param byte array
func (p *GeoRadiusParams) getParams(args [][]byte) [][]byte {
	arr := make([][]byte, 0)
	for _, a := range args {
		arr = append(arr, a)
	}

	if p.Contains("withcoord") {
		arr = append(arr, []byte("withcoord"))
	}
	if p.Contains("withdist") {
		arr = append(arr, []byte("withdist"))
	}

	if p.Contains("count") {
		arr = append(arr, []byte("count"))
		count, _ := strconv.Atoi(p.params["count"])
		arr = append(arr, IntToByteArr(count))
	}

	if p.Contains("asc") {
		arr = append(arr, []byte("asc"))
	} else if p.Contains("desc") {
		arr = append(arr, []byte("desc"))
	}

	return arr
}

//Contains test geo param contains the key
func (p *GeoRadiusParams) Contains(key string) bool {
	_, ok := p.params[key]
	return ok
}

//Tuple zset tuple
type Tuple struct {
	element string
	score   float64
}

//GeoRadiusResponse geo radius response
type GeoRadiusResponse struct {
	member     string
	distance   float64
	coordinate GeoCoordinate
}

func newGeoRadiusResponse(member string) *GeoRadiusResponse {
	return &GeoRadiusResponse{member: member}
}

//GeoCoordinate geo coordinate struct
type GeoCoordinate struct {
	longitude float64
	latitude  float64
}

//ScanResult scan result struct
type ScanResult struct {
	Cursor  string
	Results []string
}

//ZParams zset operation params
type ZParams struct {
	params []string
}

//getParams get params byte array
func (g *ZParams) getParams() [][]byte {
	return StrArrToByteArrArr(g.params)
}

//WeightsByDouble Set weights.
func (g *ZParams) WeightsByDouble(weights ...float64) *ZParams {
	g.params = append(g.params, keywordWeights.name)
	for _, w := range weights {
		g.params = append(g.params, Float64ToStr(w))
	}
	return g
}

//Aggregate Set Aggregate.
func (g *ZParams) Aggregate(aggregate *Aggregate) *ZParams {
	g.params = append(g.params, keywordAggregate.name)
	g.params = append(g.params, aggregate.name)
	return g
}

//newZParams create a new zparams instance
func newZParams() *ZParams {
	return &ZParams{params: make([]string, 0)}
}

//Aggregate aggregate,sum|min|max
type Aggregate struct {
	name string // name of Aggregate
}

//getRaw get the name byte array
func (g *Aggregate) getRaw() []byte {
	return []byte(g.name)
}

//newAggregate create a new geounit instance
func newAggregate(name string) *Aggregate {
	return &Aggregate{name}
}

var (
	//AggregateSum aggregate result with sum operation
	AggregateSum = newAggregate("SUM")
	//AggregateMin aggregate result with min operation
	AggregateMin = newAggregate("MIN")
	//AggregateMax aggregate result with max operation
	AggregateMax = newAggregate("MAX")
)

//RedisPubSub redis pubsub struct
type RedisPubSub struct {
	subscribedChannels int
	redis              *Redis
	OnMessage          func(channel, message string)                 //receive message
	OnPMessage         func(pattern string, channel, message string) //receive pattern message
	OnSubscribe        func(channel string, subscribedChannels int)  //listen subscribe event
	OnUnSubscribe      func(channel string, subscribedChannels int)  //listen unsubscribe event
	OnPUnSubscribe     func(pattern string, subscribedChannels int)  //listen pattern unsubscribe event
	OnPSubscribe       func(pattern string, subscribedChannels int)  //listen pattern subscribe event
	OnPong             func(channel string)                          //listen heart beat event
}

//Subscribe subscribe some channels
func (r *RedisPubSub) Subscribe(channels ...string) error {
	r.redis.mu.RLock()
	defer r.redis.mu.RUnlock()
	if r.redis.client == nil {
		return newConnectError("redisPubSub is not subscribed to a Redis instance")
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

//UnSubscribe unsubscribe some channels
func (r *RedisPubSub) UnSubscribe(channels ...string) error {
	r.redis.mu.RLock()
	defer r.redis.mu.RUnlock()
	if r.redis.client == nil {
		return newConnectError("redisPubSub is not subscribed to a Redis instance")
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

//PSubscribe subscribe some pattern channels
func (r *RedisPubSub) PSubscribe(channels ...string) error {
	r.redis.mu.RLock()
	defer r.redis.mu.RUnlock()
	if r.redis.client == nil {
		return newConnectError("redisPubSub is not subscribed to a Redis instance")
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

//PUnSubscribe unsubscribe some pattern channels
func (r *RedisPubSub) PUnSubscribe(channels ...string) error {
	r.redis.mu.RLock()
	defer r.redis.mu.RUnlock()
	if r.redis.client == nil {
		return newConnectError("redisPubSub is not subscribed to a Redis instance")
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
		case keywordSubscribe.name:
			r.processSubscribe(reply)
		case keywordUnsubscribe.name:
			r.processUnSubscribe(reply)
		case keywordMessage.name:
			r.processMessage(reply)
		case keywordPMessage.name:
			r.processPMessage(reply)
		case keywordPSubscribe.name:
			r.processPSubscribe(reply)
		case cmdPUnSubscribe.name:
			r.processPUnSubscribe(reply)
		case keywordPong.name:
			r.processPong(reply)
		default:
			return fmt.Errorf("unknown message type: %v", reply)
		}
		if !r.isSubscribed() {
			break
		}
	}
	redis.mu.Lock()
	defer redis.mu.Unlock()
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
	r.OnUnSubscribe(strChannel, r.subscribedChannels)
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

func (r *RedisPubSub) processPMessage(reply []interface{}) {
	bPattern := reply[1].([]byte)
	bChannel := reply[2].([]byte)
	bMsg := reply[3].([]byte)
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

func (r *RedisPubSub) processPSubscribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPSubscribe(strPattern, r.subscribedChannels)
}

func (r *RedisPubSub) processPUnSubscribe(reply []interface{}) {
	r.subscribedChannels = int(reply[2].(int64))
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPUnSubscribe(strPattern, r.subscribedChannels)
}

func (r *RedisPubSub) processPong(reply []interface{}) {
	bPattern := reply[1].([]byte)
	strPattern := ""
	if bPattern != nil {
		strPattern = string(bPattern)
	}
	r.OnPong(strPattern)
}

//BitOP bit operation struct
type BitOP struct {
	name string //name if bit operation
}

//getRaw get the name byte array
func (g *BitOP) getRaw() []byte {
	return []byte(g.name)
}

//NewBitOP
func newBitOP(name string) *BitOP {
	return &BitOP{name}
}

var (
	//BitOpAnd 'and' bit operation,&
	BitOpAnd = newBitOP("AND")
	//BitOpOr 'or' bit operation,|
	BitOpOr = newBitOP("OR")
	//BitOpXor 'xor' bit operation,X xor Y -> (X || Y) && !(X && Y)
	BitOpXor = newBitOP("XOR")
	//BitOpNot 'not' bit operation,^
	BitOpNot = newBitOP("NOT")
)

//SlowLog redis slow log struct
type SlowLog struct {
	id            int64
	timeStamp     int64
	executionTime int64
	args          []string
}

//DebugParams debug params
type DebugParams struct {
	command []string
}

//NewDebugParamsSegfault create debug prams with segfault
func NewDebugParamsSegfault() *DebugParams {
	return &DebugParams{command: []string{"SEGFAULT"}}
}

//NewDebugParamsObject create debug paramas with key
func NewDebugParamsObject(key string) *DebugParams {
	return &DebugParams{command: []string{"OBJECT", key}}
}

//NewDebugParamsReload create debug params with reload
func NewDebugParamsReload() *DebugParams {
	return &DebugParams{command: []string{"RELOAD"}}
}

//Reset reset struct
type Reset struct {
	name string //name of reset
}

//getRaw get the name byte array
func (g *Reset) getRaw() []byte {
	return []byte(g.name)
}

func newReset(name string) *Reset {
	return &Reset{name}
}

var (
	//ResetSoft soft reset
	ResetSoft = newReset("SOFT")
	//ResetHard hard reset
	ResetHard = newReset("HARD")
)

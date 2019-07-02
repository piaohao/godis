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

//GetByteParams get all params
func (p *ZAddParams) GetByteParams(key []byte, args ...[]byte) [][]byte {
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

//SortingParams sort params
type SortingParams struct {
	params [][]byte
}

//NewSortingParams create new sort params instance
func NewSortingParams() *SortingParams {
	return &SortingParams{params: make([][]byte, 0)}
}

//By set by param with pattern
func (p *SortingParams) By(pattern string) *SortingParams {
	p.params = append(p.params, keywordBy.GetRaw())
	p.params = append(p.params, []byte(pattern))
	return p
}

//NoSort set by param with nosort
func (p *SortingParams) NoSort() *SortingParams {
	p.params = append(p.params, keywordBy.GetRaw())
	p.params = append(p.params, keywordNosort.GetRaw())
	return p
}

//Desc set desc param,then sort elements in descending order
func (p *SortingParams) Desc() *SortingParams {
	p.params = append(p.params, keywordDesc.GetRaw())
	return p
}

//Asc set asc param,then sort elements in ascending order
func (p *SortingParams) Asc() *SortingParams {
	p.params = append(p.params, keywordAsc.GetRaw())
	return p
}

//Limit limit the sort result,[x,y)
func (p *SortingParams) Limit(start, count int) *SortingParams {
	p.params = append(p.params, keywordLimit.GetRaw())
	p.params = append(p.params, IntToByteArray(start))
	p.params = append(p.params, IntToByteArray(count))
	return p
}

//Alpha sort elements in alpha order
func (p *SortingParams) Alpha() *SortingParams {
	p.params = append(p.params, keywordAlpha.GetRaw())
	return p
}

//Get set get param with patterns
func (p *SortingParams) Get(patterns ...string) *SortingParams {
	for _, pattern := range patterns {
		p.params = append(p.params, keywordGet.GetRaw())
		p.params = append(p.params, []byte(pattern))
	}
	return p
}

//ScanParams scan,hscan,sscan,zscan params
type ScanParams struct {
	params map[*keyword][]byte
}

//NewScanParams create scan params instance
func NewScanParams() *ScanParams {
	return &ScanParams{params: make(map[*keyword][]byte)}
}

//GetParams get all scan params
func (s ScanParams) GetParams() [][]byte {
	arr := make([][]byte, 0)
	for k, v := range s.params {
		arr = append(arr, k.GetRaw())
		arr = append(arr, []byte(v))
	}
	return arr
}

//Match get the match param value
func (s ScanParams) Match() string {
	if v, ok := s.params[keywordMatch]; ok {
		return string(v)
	}
	return ""
}

//ListOption  list option
type ListOption struct {
	Name string // name  ...
}

//GetRaw get the option name byte array
func (l ListOption) GetRaw() []byte {
	return []byte(l.Name)
}

//NewListOption create new list option instance
func newListOption(name string) ListOption {
	return ListOption{name}
}

var (
	//ListOptionBefore insert an new element before designated element
	ListOptionBefore = newListOption("BEFORE")
	//ListOptionAfter insert an new element after designated element
	ListOptionAfter = newListOption("AFTER")
)

//GeoUnit geo unit,m|mi|km|ft
type GeoUnit struct {
	Name string // name of geo unit
}

//GetRaw get the name byte array
func (g GeoUnit) GetRaw() []byte {
	return []byte(g.Name)
}

//NewGeoUnit create a new geounit instance
func newGeoUnit(name string) GeoUnit {
	return GeoUnit{name}
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

//GeoRadiusParam geo radius param
type GeoRadiusParam struct {
	params map[string]string
}

//NewGeoRadiusParam create a new geo radius param instance
func NewGeoRadiusParam() *GeoRadiusParam {
	return &GeoRadiusParam{params: make(map[string]string)}
}

//WithCoord fill the geo result with coordinate
func (p *GeoRadiusParam) WithCoord() *GeoRadiusParam {
	p.params["withcoord"] = "withcoord"
	return p
}

//WithDist fill the geo result with distance
func (p *GeoRadiusParam) WithDist() *GeoRadiusParam {
	p.params["withdist"] = "withdist"
	return p
}

//SortAscending sort th geo result in ascending order
func (p *GeoRadiusParam) SortAscending() *GeoRadiusParam {
	p.params["asc"] = "asc"
	return p
}

//SortDescending sort the geo result in descending order
func (p *GeoRadiusParam) SortDescending() *GeoRadiusParam {
	p.params["desc"] = "desc"
	return p
}

//Count fill the geo result with count
func (p *GeoRadiusParam) Count(count int) *GeoRadiusParam {
	if count > 0 {
		p.params["count"] = strconv.Itoa(count)
	}
	return p
}

//GetParams  get geo param byte array
func (p *GeoRadiusParam) GetParams(args [][]byte) [][]byte {
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
		arr = append(arr, IntToByteArray(count))
	}

	if p.Contains("asc") {
		arr = append(arr, []byte("asc"))
	} else if p.Contains("desc") {
		arr = append(arr, []byte("desc"))
	}

	return arr
}

//Contains test geo param contains the key
func (p *GeoRadiusParam) Contains(key string) bool {
	_, ok := p.params[key]
	return ok
}

//Tuple zset tuple
type Tuple struct {
	element []byte
	score   float64
}

//GeoRadiusResponse geo radius response
type GeoRadiusResponse struct {
	member     []byte
	distance   float64
	coordinate GeoCoordinate
}

func newGeoRadiusResponse(member []byte) *GeoRadiusResponse {
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
	params [][]byte
}

//GetParams get params byte array
func (g *ZParams) GetParams() [][]byte {
	return g.params
}

//WeightsByDouble Set weights.
func (g *ZParams) WeightsByDouble(weights ...float64) *ZParams {
	g.params = append(g.params, keywordWeights.GetRaw())
	for _, w := range weights {
		g.params = append(g.params, Float64ToByteArray(w))
	}
	return g
}

//Aggregate Set Aggregate.
func (g *ZParams) Aggregate(aggregate Aggregate) *ZParams {
	g.params = append(g.params, keywordAggregate.GetRaw())
	g.params = append(g.params, aggregate.GetRaw())
	return g
}

//newZParams create a new zparams instance
func newZParams() *ZParams {
	return &ZParams{params: make([][]byte, 0)}
}

//Aggregate aggregate,sum|min|max
type Aggregate struct {
	Name string // name of Aggregate
}

//GetRaw get the name byte array
func (g Aggregate) GetRaw() []byte {
	return []byte(g.Name)
}

//newAggregate create a new geounit instance
func newAggregate(name string) Aggregate {
	return Aggregate{name}
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
		case keywordSubscribe.Name:
			r.processSubscribe(reply)
		case keywordUnsubscribe.Name:
			r.processUnSubscribe(reply)
		case keywordMessage.Name:
			r.processMessage(reply)
		case keywordPMessage.Name:
			r.processPmessage(reply)
		case keywordPSubscribe.Name:
			r.processPsubscribe(reply)
		case cmdPUnSubscribe.Name:
			r.processPunsubcribe(reply)
		case keywordPong.Name:
			r.processPong(reply)
		default:
			return fmt.Errorf("unknown message type: %v", reply)
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

func (r *RedisPubSub) processPmessage(reply []interface{}) {
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
	Name string //name if bit operation
}

//GetRaw get the name byte array
func (g BitOP) GetRaw() []byte {
	return []byte(g.Name)
}

//NewBitOP
func newBitOP(name string) BitOP {
	return BitOP{name}
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
	Name string //name of reset
}

//GetRaw get the name byte array
func (g *Reset) GetRaw() []byte {
	return []byte(g.Name)
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

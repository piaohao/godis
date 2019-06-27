package godis

import "sync"

// pipleline and transaction response,include replys from redis
type response struct {
	response interface{} //store replys

	building bool //whether response is building
	built    bool //whether response is build done
	isSet    bool //whether response is set with data

	builder    Builder     //response data convert rule
	data       interface{} //real data
	dependency *response   //response cycle dependency
}

func newResponse() *response {
	return &response{
		building: false,
		built:    false,
		isSet:    false,
	}
}

func (r *response) set(data interface{}) {
	r.data = data
	r.isSet = true
}

//Get get real content of response
func (r *response) Get() (interface{}, error) {
	if r.dependency != nil && r.dependency.isSet && !r.dependency.built {
		err := r.dependency.build()
		if err != nil {
			return nil, err
		}
	}
	if !r.isSet {
		return nil, newDataError("please close pipeline or multi block before calling this method")
	}
	if !r.built {
		err := r.build()
		if err != nil {
			return nil, err
		}
	}
	return r.response, nil
}

func (r *response) setDependency(dependency *response) {
	r.dependency = dependency
}

func (r *response) build() error {
	if r.building {
		return nil
	}
	r.building = true
	defer func() {
		r.building = false
		r.built = true
	}()
	if r.data != nil {
		r.response = r.builder.build(r.data)
	}
	r.data = nil
	return nil
}

type transaction struct {
	*multiKeyPipelineBase
	inTransaction bool
}

func newTransaction(c *client) *transaction {
	base := newMultiKeyPipelineBase(c)
	base.getClient = func(key string) *client {
		return c
	}
	return &transaction{multiKeyPipelineBase: base}
}

//Clear  clear
func (t *transaction) Clear() (string, error) {
	if t.inTransaction {
		return t.Discard()
	}
	return "", nil
}

//Exec execute transaction
func (t *transaction) Exec() ([]interface{}, error) {
	err := t.client.exec()
	if err != nil {
		return nil, err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return nil, err
	}
	t.inTransaction = false
	reply, err := t.client.getObjectMultiBulkReply()
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

//ExecGetResponse ...
func (t *transaction) ExecGetResponse() ([]*response, error) {
	err := t.client.exec()
	if err != nil {
		return nil, err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return nil, err
	}
	t.inTransaction = false
	reply, err := t.client.getObjectMultiBulkReply()
	if err != nil {
		return nil, err
	}
	result := make([]*response, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

//Discard  ...
func (t *transaction) Discard() (string, error) {
	err := t.client.discard()
	if err != nil {
		return "", err
	}
	_, err = t.client.getAll(1)
	if err != nil {
		return "", err
	}
	t.inTransaction = false
	t.clean()
	return t.client.getStatusCodeReply()
}

func (t *transaction) clean() {
	t.pipelinedResponses = make([]*response, 0)
}

type pipeline struct {
	*multiKeyPipelineBase
}

func newPipeline(c *client) *pipeline {
	base := newMultiKeyPipelineBase(c)
	base.getClient = func(key string) *client {
		return c
	}
	return &pipeline{multiKeyPipelineBase: base}
}

//Sync  ...
func (p *pipeline) Sync() error {
	if len(p.pipelinedResponses) == 0 {
		return nil
	}
	all, err := p.client.connection.getAll()
	if err != nil {
		return err
	}
	for _, a := range all.([]interface{}) {
		p.generateResponse(a)
	}
	return nil
}

type queable struct {
	pipelinedResponses []*response
	mu                 sync.Mutex
}

func newQueable() *queable {
	return &queable{pipelinedResponses: make([]*response, 0)}
}

func (q *queable) clean() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.pipelinedResponses = make([]*response, 0)
}

func (q *queable) generateResponse(data interface{}) *response {
	q.mu.Lock()
	defer q.mu.Unlock()
	size := len(q.pipelinedResponses)
	if size == 0 {
		return nil
	}
	r := q.pipelinedResponses[0]
	r.set(data)
	if size == 1 {
		q.pipelinedResponses = make([]*response, 0)
	} else {
		q.pipelinedResponses = q.pipelinedResponses[1:]
	}
	return r
}

func (q *queable) getResponse(builder Builder) *response {
	q.mu.Lock()
	defer q.mu.Unlock()
	response := newResponse()
	response.builder = builder
	q.pipelinedResponses = append(q.pipelinedResponses, response)
	return response
}

func (q *queable) hasPipelinedResponse() bool {
	return q.getPipelinedResponseLength() > 0
}

func (q *queable) getPipelinedResponseLength() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.pipelinedResponses)
}

type multiKeyPipelineBase struct {
	*queable
	client *client

	getClient func(key string) *client
}

func newMultiKeyPipelineBase(client *client) *multiKeyPipelineBase {
	return &multiKeyPipelineBase{queable: newQueable(), client: client}
}

//<editor-fold desc="basicpipeline">

//Bgrewriteaof ...
func (p *multiKeyPipelineBase) Bgrewriteaof() (*response, error) {
	err := p.client.bgrewriteaof()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Bgsave  ...
func (p *multiKeyPipelineBase) Bgsave() (*response, error) {
	err := p.client.bgsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ConfigGet  ...
func (p *multiKeyPipelineBase) ConfigGet(pattern string) (*response, error) {
	err := p.client.configGet(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//ConfigSet  ...
func (p *multiKeyPipelineBase) ConfigSet(parameter, value string) (*response, error) {
	err := p.client.configSet(parameter, value)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ConfigResetStat  ...
func (p *multiKeyPipelineBase) ConfigResetStat() (*response, error) {
	err := p.client.configResetStat()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Save  ...
func (p *multiKeyPipelineBase) Save() (*response, error) {
	err := p.client.save()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Lastsave  ...
func (p *multiKeyPipelineBase) Lastsave() (*response, error) {
	err := p.client.lastsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//FlushDB  ...
func (p *multiKeyPipelineBase) FlushDB() (*response, error) {
	err := p.client.flushDB()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//FlushAll  ...
func (p *multiKeyPipelineBase) FlushAll() (*response, error) {
	err := p.client.flushAll()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Info  ...
func (p *multiKeyPipelineBase) Info() (*response, error) {
	err := p.client.info()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Time  ...
func (p *multiKeyPipelineBase) Time() (*response, error) {
	err := p.client.time()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//DbSize  ...
func (p *multiKeyPipelineBase) DbSize() (*response, error) {
	err := p.client.dbSize()
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Shutdown  ...
func (p *multiKeyPipelineBase) Shutdown() (*response, error) {
	err := p.client.shutdown()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Ping  ...
func (p *multiKeyPipelineBase) Ping() (*response, error) {
	err := p.client.ping()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Select  ...
func (p *multiKeyPipelineBase) Select(index int) (*response, error) {
	err := p.client.selectDb(index)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//</editor-fold>

//<editor-fold desc="multikeypipeline">

//Del ...
func (p *multiKeyPipelineBase) Del(keys ...string) (*response, error) {
	err := p.client.del(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Exists  ...
func (p *multiKeyPipelineBase) Exists(keys ...string) (*response, error) {
	err := p.client.exists(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//BlpopTimout  ...
func (p *multiKeyPipelineBase) BlpopTimout(timeout int, keys ...string) (*response, error) {
	err := p.client.blpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//BrpopTimout  ...
func (p *multiKeyPipelineBase) BrpopTimout(timeout int, keys ...string) (*response, error) {
	err := p.client.brpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Blpop  ...
func (p *multiKeyPipelineBase) Blpop(args ...string) (*response, error) {
	err := p.client.blpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Brpop  ...
func (p *multiKeyPipelineBase) Brpop(args ...string) (*response, error) {
	err := p.client.brpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Keys  ...
func (p *multiKeyPipelineBase) Keys(pattern string) (*response, error) {
	err := p.client.keys(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Mget  ...
func (p *multiKeyPipelineBase) Mget(keys ...string) (*response, error) {
	err := p.client.mget(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Mset  ...
func (p *multiKeyPipelineBase) Mset(keysvalues ...string) (*response, error) {
	err := p.client.mset(keysvalues...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Msetnx  ...
func (p *multiKeyPipelineBase) Msetnx(keysvalues ...string) (*response, error) {
	err := p.client.msetnx(keysvalues...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Rename  ...
func (p *multiKeyPipelineBase) Rename(oldkey, newkey string) (*response, error) {
	err := p.client.rename(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Renamenx  ...
func (p *multiKeyPipelineBase) Renamenx(oldkey, newkey string) (*response, error) {
	err := p.client.renamenx(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Rpoplpush  ...
func (p *multiKeyPipelineBase) Rpoplpush(srckey, dstkey string) (*response, error) {
	err := p.client.rpopLpush(srckey, dstkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Sdiff  ...
func (p *multiKeyPipelineBase) Sdiff(keys ...string) (*response, error) {
	err := p.client.sdiff(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Sdiffstore  ...
func (p *multiKeyPipelineBase) Sdiffstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.sdiffstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Sinter  ...
func (p *multiKeyPipelineBase) Sinter(keys ...string) (*response, error) {
	err := p.client.sinter(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Sinterstore  ...
func (p *multiKeyPipelineBase) Sinterstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.sinterstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Smove  ...
func (p *multiKeyPipelineBase) Smove(srckey, dstkey, member string) (*response, error) {
	err := p.client.smove(srckey, dstkey, member)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//SortMulti  ...
func (p *multiKeyPipelineBase) SortMulti(key string, dstkey string, sortingParameters ...SortingParams) (*response, error) {
	err := p.client.sortMulti(key, dstkey, sortingParameters...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Sunion  ...
func (p *multiKeyPipelineBase) Sunion(keys ...string) (*response, error) {
	err := p.client.sunion(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Sunionstore  ...
func (p *multiKeyPipelineBase) Sunionstore(dstkey string, keys ...string) (*response, error) {
	err := p.client.sunionstore(dstkey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Watch  ...
func (p *multiKeyPipelineBase) Watch(keys ...string) (*response, error) {
	err := p.client.watch(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Zinterstore  ...
func (p *multiKeyPipelineBase) Zinterstore(dstkey string, sets ...string) (*response, error) {
	err := p.client.zinterstore(dstkey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//ZinterstoreWithParams  ...
func (p *multiKeyPipelineBase) ZinterstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error) {
	err := p.client.zinterstoreWithParams(dstkey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Zunionstore  ...
func (p *multiKeyPipelineBase) Zunionstore(dstkey string, sets ...string) (*response, error) {
	err := p.client.zunionstore(dstkey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//ZunionstoreWithParams  ...
func (p *multiKeyPipelineBase) ZunionstoreWithParams(dstkey string, params ZParams, sets ...string) (*response, error) {
	err := p.client.zunionstoreWithParams(dstkey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Brpoplpush  ...
func (p *multiKeyPipelineBase) Brpoplpush(source, destination string, timeout int) (*response, error) {
	err := p.client.brpoplpush(source, destination, timeout)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//Publish  ...
func (p *multiKeyPipelineBase) Publish(channel, message string) (*response, error) {
	err := p.client.publish(channel, message)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//RandomKey  ...
func (p *multiKeyPipelineBase) RandomKey() (*response, error) {
	err := p.client.randomKey()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Bitop  ...
func (p *multiKeyPipelineBase) Bitop(op BitOP, destKey string, srcKeys ...string) (*response, error) {
	err := p.client.bitop(op, destKey, srcKeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Pfmerge  ...
func (p *multiKeyPipelineBase) Pfmerge(destkey string, sourcekeys ...string) (*response, error) {
	err := p.client.pfmerge(destkey, sourcekeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Pfcount  ...
func (p *multiKeyPipelineBase) Pfcount(keys ...string) (*response, error) {
	err := p.client.pfcount(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//</editor-fold>

//<editor-fold desc="cluster pipeline">

//ClusterNodes ...
func (p *multiKeyPipelineBase) ClusterNodes() (*response, error) {
	err := p.client.clusterNodes()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterMeet  ...
func (p *multiKeyPipelineBase) ClusterMeet(ip string, port int) (*response, error) {
	err := p.client.clusterMeet(ip, port)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterAddSlots  ...
func (p *multiKeyPipelineBase) ClusterAddSlots(slots ...int) (*response, error) {
	err := p.client.clusterAddSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterDelSlots  ...
func (p *multiKeyPipelineBase) ClusterDelSlots(slots ...int) (*response, error) {
	err := p.client.clusterDelSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterInfo  ...
func (p *multiKeyPipelineBase) ClusterInfo() (*response, error) {
	err := p.client.clusterInfo()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterGetKeysInSlot  ...
func (p *multiKeyPipelineBase) ClusterGetKeysInSlot(slot int, count int) (*response, error) {
	err := p.client.clusterGetKeysInSlot(slot, count)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringArrayBuilder), nil
}

//ClusterSetSlotNode  ...
func (p *multiKeyPipelineBase) ClusterSetSlotNode(slot int, nodeId string) (*response, error) {
	err := p.client.clusterSetSlotNode(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterSetSlotMigrating  ...
func (p *multiKeyPipelineBase) ClusterSetSlotMigrating(slot int, nodeId string) (*response, error) {
	err := p.client.clusterSetSlotMigrating(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//ClusterSetSlotImporting  ...
func (p *multiKeyPipelineBase) ClusterSetSlotImporting(slot int, nodeId string) (*response, error) {
	err := p.client.clusterSetSlotImporting(slot, nodeId)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//</editor-fold>

//<editor-fold desc="scripting pipeline">

//Eval ...
func (p *multiKeyPipelineBase) Eval(script string, keyCount int, params ...string) (*response, error) {
	err := p.getClient(script).eval(script, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//Evalsha  ...
func (p *multiKeyPipelineBase) Evalsha(sha1 string, keyCount int, params ...string) (*response, error) {
	err := p.getClient(sha1).evalsha(sha1, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StringBuilder), nil
}

//</editor-fold>

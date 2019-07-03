package godis

import "sync"

//Response pipeline and transaction response,include replies from redis
type Response struct {
	response  interface{} //store replies
	exception *DataError

	building bool //whether response is building
	built    bool //whether response is build done
	isSet    bool //whether response is set with data

	builder    Builder     //response data convert rule
	data       interface{} //real data
	dependency *Response   //response cycle dependency
}

func newResponse() *Response {
	return &Response{
		building: false,
		built:    false,
		isSet:    false,
	}
}

func (r *Response) set(data interface{}) {
	r.data = data
	r.isSet = true
}

//Get get real content of response
func (r *Response) Get() (interface{}, error) {
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
	if r.exception != nil {
		return nil, r.exception
	}
	return r.response, nil
}

func (r *Response) setDependency(dependency *Response) {
	r.dependency = dependency
}

func (r *Response) build() error {
	if r.building {
		return nil
	}
	r.building = true
	defer func() {
		r.building = false
		r.built = true
	}()
	if r.data != nil {
		switch r.data.(type) {
		case *DataError:
			r.exception = r.data.(*DataError)
			return nil
		}
		result, err := r.builder.build(r.data)
		if err != nil {
			return err
		}
		r.response = result
	}
	r.data = nil
	return nil
}

//Transaction redis transaction struct
type Transaction struct {
	*multiKeyPipelineBase
	inTransaction bool
}

func newTransaction(c *client) *Transaction {
	base := newMultiKeyPipelineBase(c)
	base.getClient = func(key string) *client {
		return c
	}
	return &Transaction{multiKeyPipelineBase: base}
}

//Clear  clear
func (t *Transaction) Clear() (string, error) {
	if t.inTransaction {
		return t.Discard()
	}
	return "", nil
}

//Exec execute transaction
func (t *Transaction) Exec() ([]interface{}, error) {
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
func (t *Transaction) ExecGetResponse() ([]*Response, error) {
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
	result := make([]*Response, 0)
	for _, r := range reply {
		result = append(result, t.generateResponse(r))
	}
	return result, nil
}

//Discard  see redis command
func (t *Transaction) Discard() (string, error) {
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

func (t *Transaction) clean() {
	t.pipelinedResponses = make([]*Response, 0)
}

//Pipeline redis pipeline struct
type Pipeline struct {
	*multiKeyPipelineBase
}

func newPipeline(c *client) *Pipeline {
	base := newMultiKeyPipelineBase(c)
	base.getClient = func(key string) *client {
		return c
	}
	return &Pipeline{multiKeyPipelineBase: base}
}

//Sync  see redis command
func (p *Pipeline) Sync() error {
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

type queue struct {
	pipelinedResponses []*Response
	mu                 sync.Mutex
}

func newQueue() *queue {
	return &queue{pipelinedResponses: make([]*Response, 0)}
}

func (q *queue) clean() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.pipelinedResponses = make([]*Response, 0)
}

func (q *queue) generateResponse(data interface{}) *Response {
	q.mu.Lock()
	defer q.mu.Unlock()
	size := len(q.pipelinedResponses)
	if size == 0 {
		return nil
	}
	r := q.pipelinedResponses[0]
	r.set(data)
	if size == 1 {
		q.pipelinedResponses = make([]*Response, 0)
	} else {
		q.pipelinedResponses = q.pipelinedResponses[1:]
	}
	return r
}

func (q *queue) getResponse(builder Builder) *Response {
	q.mu.Lock()
	defer q.mu.Unlock()
	response := newResponse()
	response.builder = builder
	q.pipelinedResponses = append(q.pipelinedResponses, response)
	return response
}

func (q *queue) hasPipelinedResponse() bool {
	return q.getPipelinedResponseLength() > 0
}

func (q *queue) getPipelinedResponseLength() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.pipelinedResponses)
}

type multiKeyPipelineBase struct {
	*queue
	client *client

	getClient func(key string) *client
}

func newMultiKeyPipelineBase(client *client) *multiKeyPipelineBase {
	return &multiKeyPipelineBase{queue: newQueue(), client: client}
}

//<editor-fold desc="basicpipeline">

//BgRewriteAof see redis command
func (p *multiKeyPipelineBase) BgRewriteAof() (*Response, error) {
	err := p.client.bgrewriteaof()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//BgSave  see redis command
func (p *multiKeyPipelineBase) BgSave() (*Response, error) {
	err := p.client.bgsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ConfigGet  see redis command
func (p *multiKeyPipelineBase) ConfigGet(pattern string) (*Response, error) {
	err := p.client.configGet(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//ConfigSet  see redis command
func (p *multiKeyPipelineBase) ConfigSet(parameter, value string) (*Response, error) {
	err := p.client.configSet(parameter, value)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ConfigResetStat  see redis command
func (p *multiKeyPipelineBase) ConfigResetStat() (*Response, error) {
	err := p.client.configResetStat()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//Save  see redis command
func (p *multiKeyPipelineBase) Save() (*Response, error) {
	err := p.client.save()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//LastSave  see redis command
func (p *multiKeyPipelineBase) LastSave() (*Response, error) {
	err := p.client.lastsave()
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//FlushDB  see redis command
func (p *multiKeyPipelineBase) FlushDB() (*Response, error) {
	err := p.client.flushDB()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//FlushAll  see redis command
func (p *multiKeyPipelineBase) FlushAll() (*Response, error) {
	err := p.client.flushAll()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//Info  see redis command
func (p *multiKeyPipelineBase) Info() (*Response, error) {
	err := p.client.info()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//Time  see redis command
func (p *multiKeyPipelineBase) Time() (*Response, error) {
	err := p.client.time()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//DbSize  see redis command
func (p *multiKeyPipelineBase) DbSize() (*Response, error) {
	err := p.client.dbSize()
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Shutdown  see redis command
func (p *multiKeyPipelineBase) Shutdown() (*Response, error) {
	err := p.client.shutdown()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//Ping  see redis command
func (p *multiKeyPipelineBase) Ping() (*Response, error) {
	err := p.client.ping()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//Select  see redis command
func (p *multiKeyPipelineBase) Select(index int) (*Response, error) {
	err := p.client.selectDb(index)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//</editor-fold>

//<editor-fold desc="multikeypipeline">

//Del see redis command
func (p *multiKeyPipelineBase) Del(keys ...string) (*Response, error) {
	err := p.client.del(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Exists  see redis command
func (p *multiKeyPipelineBase) Exists(keys ...string) (*Response, error) {
	err := p.client.exists(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//BLPopTimeout  see redis command
func (p *multiKeyPipelineBase) BLPopTimeout(timeout int, keys ...string) (*Response, error) {
	err := p.client.blpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//BRPopTimeout  see redis command
func (p *multiKeyPipelineBase) BRPopTimeout(timeout int, keys ...string) (*Response, error) {
	err := p.client.brpopTimout(timeout, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//BLPop  see redis command
func (p *multiKeyPipelineBase) BLPop(args ...string) (*Response, error) {
	err := p.client.blpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//BRPop  see redis command
func (p *multiKeyPipelineBase) BRPop(args ...string) (*Response, error) {
	err := p.client.brpop(args)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//Keys  see redis command
func (p *multiKeyPipelineBase) Keys(pattern string) (*Response, error) {
	err := p.client.keys(pattern)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//MGet  see redis command
func (p *multiKeyPipelineBase) MGet(keys ...string) (*Response, error) {
	err := p.client.mget(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//MSet  see redis command
func (p *multiKeyPipelineBase) MSet(kvs ...string) (*Response, error) {
	err := p.client.mset(kvs...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//MSetNx  see redis command
func (p *multiKeyPipelineBase) MSetNx(kvs ...string) (*Response, error) {
	err := p.client.msetnx(kvs...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Rename  see redis command
func (p *multiKeyPipelineBase) Rename(oldkey, newkey string) (*Response, error) {
	err := p.client.rename(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//RenameNx  see redis command
func (p *multiKeyPipelineBase) RenameNx(oldkey, newkey string) (*Response, error) {
	err := p.client.renamenx(oldkey, newkey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//RPopLPush  see redis command
func (p *multiKeyPipelineBase) RPopLPush(srcKey, destKey string) (*Response, error) {
	err := p.client.rpopLpush(srcKey, destKey)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//SDiff  see redis command
func (p *multiKeyPipelineBase) SDiff(keys ...string) (*Response, error) {
	err := p.client.sDiff(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//SDiffStore  see redis command
func (p *multiKeyPipelineBase) SDiffStore(destKey string, keys ...string) (*Response, error) {
	err := p.client.sDiffStore(destKey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//SInter  see redis command
func (p *multiKeyPipelineBase) SInter(keys ...string) (*Response, error) {
	err := p.client.sInter(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//SInterStore  see redis command
func (p *multiKeyPipelineBase) SInterStore(destKey string, keys ...string) (*Response, error) {
	err := p.client.sInterStore(destKey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//SMove  see redis command
func (p *multiKeyPipelineBase) SMove(srcKey, destKey, member string) (*Response, error) {
	err := p.client.smove(srcKey, destKey, member)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//SortMulti  see redis command
func (p *multiKeyPipelineBase) SortStore(key string, destKey string, params ...*SortParams) (*Response, error) {
	err := p.client.sortMulti(key, destKey, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//SUnion  see redis command
func (p *multiKeyPipelineBase) SUnion(keys ...string) (*Response, error) {
	err := p.client.sUnion(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//SUnionStore  see redis command
func (p *multiKeyPipelineBase) SUnionStore(destKey string, keys ...string) (*Response, error) {
	err := p.client.sUnionStore(destKey, keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//Watch  see redis command
func (p *multiKeyPipelineBase) Watch(keys ...string) (*Response, error) {
	err := p.client.watch(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ZInterStore  see redis command
func (p *multiKeyPipelineBase) ZInterStore(destKey string, sets ...string) (*Response, error) {
	err := p.client.zinterstore(destKey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//ZInterStoreWithParams  see redis command
func (p *multiKeyPipelineBase) ZInterStoreWithParams(destKey string, params *ZParams, sets ...string) (*Response, error) {
	err := p.client.zinterstoreWithParams(destKey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//ZUnionStore  see redis command
func (p *multiKeyPipelineBase) ZUnionStore(destKey string, sets ...string) (*Response, error) {
	err := p.client.zunionstore(destKey, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//ZUnionStoreWithParams  see redis command
func (p *multiKeyPipelineBase) ZUnionStoreWithParams(destKey string, params *ZParams, sets ...string) (*Response, error) {
	err := p.client.zunionstoreWithParams(destKey, params, sets...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//BRPopLPush  see redis command
func (p *multiKeyPipelineBase) BRPopLPush(source, destination string, timeout int) (*Response, error) {
	err := p.client.brpoplpush(source, destination, timeout)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//Publish  see redis command
func (p *multiKeyPipelineBase) Publish(channel, message string) (*Response, error) {
	err := p.client.publish(channel, message)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//RandomKey  see redis command
func (p *multiKeyPipelineBase) RandomKey() (*Response, error) {
	err := p.client.randomKey()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//BitOp  see redis command
func (p *multiKeyPipelineBase) BitOp(op BitOP, destKey string, srcKeys ...string) (*Response, error) {
	err := p.client.bitop(op, destKey, srcKeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//PfMerge  see redis command
func (p *multiKeyPipelineBase) PfMerge(destKey string, srcKeys ...string) (*Response, error) {
	err := p.client.pfmerge(destKey, srcKeys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//PfCount  see redis command
func (p *multiKeyPipelineBase) PfCount(keys ...string) (*Response, error) {
	err := p.client.pfcount(keys...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(Int64Builder), nil
}

//</editor-fold>

//<editor-fold desc="cluster pipeline">

//ClusterNodes see redis command
func (p *multiKeyPipelineBase) ClusterNodes() (*Response, error) {
	err := p.client.clusterNodes()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterMeet  see redis command
func (p *multiKeyPipelineBase) ClusterMeet(ip string, port int) (*Response, error) {
	err := p.client.clusterMeet(ip, port)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterAddSlots  see redis command
func (p *multiKeyPipelineBase) ClusterAddSlots(slots ...int) (*Response, error) {
	err := p.client.clusterAddSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterDelSlots  see redis command
func (p *multiKeyPipelineBase) ClusterDelSlots(slots ...int) (*Response, error) {
	err := p.client.clusterDelSlots(slots...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterInfo  see redis command
func (p *multiKeyPipelineBase) ClusterInfo() (*Response, error) {
	err := p.client.clusterInfo()
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterGetKeysInSlot  see redis command
func (p *multiKeyPipelineBase) ClusterGetKeysInSlot(slot int, count int) (*Response, error) {
	err := p.client.clusterGetKeysInSlot(slot, count)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrArrBuilder), nil
}

//ClusterSetSlotNode  see redis command
func (p *multiKeyPipelineBase) ClusterSetSlotNode(slot int, nodeID string) (*Response, error) {
	err := p.client.clusterSetSlotNode(slot, nodeID)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterSetSlotMigrating  see redis command
func (p *multiKeyPipelineBase) ClusterSetSlotMigrating(slot int, nodeID string) (*Response, error) {
	err := p.client.clusterSetSlotMigrating(slot, nodeID)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//ClusterSetSlotImporting  see redis command
func (p *multiKeyPipelineBase) ClusterSetSlotImporting(slot int, nodeID string) (*Response, error) {
	err := p.client.clusterSetSlotImporting(slot, nodeID)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//</editor-fold>

//<editor-fold desc="scripting pipeline">

//Eval see redis command
func (p *multiKeyPipelineBase) Eval(script string, keyCount int, params ...string) (*Response, error) {
	err := p.getClient(script).eval(script, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//EvalSha  see redis command
func (p *multiKeyPipelineBase) EvalSha(sha1 string, keyCount int, params ...string) (*Response, error) {
	err := p.getClient(sha1).evalsha(sha1, keyCount, params...)
	if err != nil {
		return nil, err
	}
	return p.getResponse(StrBuilder), nil
}

//</editor-fold>

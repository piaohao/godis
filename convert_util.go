package godis

import (
	"encoding/binary"
	"math"
	"strconv"
)

//BoolToByteArray ...
func BoolToByteArray(a bool) []byte {
	if a {
		return BytesTrue
	}
	return BytesFalse
}

//IntToByteArray ...
func IntToByteArray(a int) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, int64(a), 10)
}

//Int64ToByteArray ...
func Int64ToByteArray(a int64) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, a, 10)
}

//Float64ToByteArray ...
func Float64ToByteArray(a float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(a))
	return buf[:]
}

//ByteArrayToFloat64 ...
func ByteArrayToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//ByteArrayToInt64 ...
func ByteArrayToInt64(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

//StringStringArrayToByteArray ...
func StringStringArrayToByteArray(str string, arr []string) [][]byte {
	params := make([][]byte, 0)
	params = append(params, []byte(str))
	for _, v := range arr {
		params = append(params, []byte(v))
	}
	return params
}

//StringStringArrayToStringArray ...
func StringStringArrayToStringArray(str string, arr []string) []string {
	params := make([]string, 0)
	params = append(params, str)
	for _, v := range arr {
		params = append(params, v)
	}
	return params
}

//StringArrayToByteArray ...
func StringArrayToByteArray(arr []string) [][]byte {
	newArr := make([][]byte, 0)
	for _, a := range arr {
		newArr = append(newArr, []byte(a))
	}
	return newArr
}

//StringToFloat64Reply ...
func StringToFloat64Reply(reply string, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	f, e := strconv.ParseFloat(reply, 64)
	if e != nil {
		return 0, e
	}
	return f, nil
}

//StringArrayToMapReply ...
func StringArrayToMapReply(reply []string, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	newMap := make(map[string]string, len(reply)/2)
	for i := 0; i < len(reply); i += 2 {
		newMap[reply[i]] = reply[i+1]
	}
	return newMap, nil
}

//Int64ToBoolReply ...
func Int64ToBoolReply(reply int64, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

//ByteToStringReply ...
func ByteToStringReply(reply []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(reply), nil
}

//StringArrToTupleReply ...
func StringArrToTupleReply(reply []string, err error) ([]Tuple, error) {
	if len(reply) == 0 {
		return []Tuple{}, nil
	}
	newArr := make([]Tuple, len(reply)/2)
	for i := 0; i < len(reply); i += 2 {
		f, err := strconv.ParseFloat(reply[i+1], 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, Tuple{element: []byte(reply[i]), score: f})
	}
	return newArr, err
}

//ObjectArrToScanResultReply ...
func ObjectArrToScanResultReply(reply []interface{}, err error) (*ScanResult, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	nexCursor := string(reply[0].([]byte))
	result := make([]string, 0)
	for _, r := range reply[1].([][]byte) {
		result = append(result, string(r))
	}
	return &ScanResult{Cursor: nexCursor, Results: result}, err
}

//ObjectArrToGeoCoordinateReply ...
func ObjectArrToGeoCoordinateReply(reply []interface{}, err error) ([]*GeoCoordinate, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	arr := make([]*GeoCoordinate, 0)
	for _, r := range reply {
		if r == nil {
			arr = append(arr, nil)
		} else {
			rArr := r.([][]byte)
			arr = append(arr, &GeoCoordinate{
				longitude: ByteArrayToFloat64(rArr[0]),
				latitude:  ByteArrayToFloat64(rArr[1]),
			})
		}
	}
	return arr, err
}

//ObjectArrToMapArrayReply ...
func ObjectArrToMapArrayReply(reply []interface{}, err error) ([]map[string]string, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	masters := make([]map[string]string, 0)
	for _, re := range reply {
		m := make(map[string]string)
		arr := re.([][]byte)
		for i := 0; i < len(arr); i += 2 {
			m[string(arr[i])] = string(arr[i+1])
		}
		masters = append(masters, m)
	}
	return masters, nil
}

//ObjectToEvalResult ...
func ObjectToEvalResult(reply interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	//todo reply解析待完成
	return reply, err
}

//<editor-fold desc="cluster reply convert">

//ToStringReply ...
func ToStringReply(reply interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return reply.(string), nil
}

//ToInt64Reply ...
func ToInt64Reply(reply interface{}, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return reply.(int64), nil
}

//ToInt64ArrayReply ...
func ToInt64ArrayReply(reply interface{}, err error) ([]int64, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]int64), nil
}

//ToBoolReply ...
func ToBoolReply(reply interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return reply.(bool), nil
}

//ToFloat64Reply ...
func ToFloat64Reply(reply interface{}, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return reply.(float64), nil
}

//ToBoolArrayReply ...
func ToBoolArrayReply(reply interface{}, err error) ([]bool, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]bool), nil
}

//ToStringArrayReply ...
func ToStringArrayReply(reply interface{}, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]string), nil
}

//ToScanResultReply ...
func ToScanResultReply(reply interface{}, err error) (*ScanResult, error) {
	if err != nil {
		return nil, err
	}
	return reply.(*ScanResult), nil
}

//ToMapReply ...
func ToMapReply(reply interface{}, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	return reply.(map[string]string), nil
}

//ToTupleArrayReply ...
func ToTupleArrayReply(reply interface{}, err error) ([]Tuple, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]Tuple), nil
}

//ToGeoArrayReply ...
func ToGeoArrayReply(reply interface{}, err error) ([]*GeoCoordinate, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]*GeoCoordinate), nil
}

//</editor-fold>

//Builder convert pipeline|transaction response data
type Builder interface {
	build(data interface{}) interface{}
}

var (
	//convert interface to string
	StringBuilder = newStringBuilder()
	//convert interface to int64
	Int64Builder = newInt64Builder()
	//convert interface to string array
	StringArrayBuilder = newStringArrayBuilder()
)

type stringBuilder struct {
}

func newStringBuilder() *stringBuilder {
	return &stringBuilder{}
}

func (b *stringBuilder) build(data interface{}) interface{} {
	if data == nil {
		return ""
	}
	return string(data.([]byte))
}

type int64Builder struct {
}

func newInt64Builder() *int64Builder {
	return &int64Builder{}
}

func (b *int64Builder) build(data interface{}) interface{} {
	if data == nil {
		return ""
	}
	return ByteArrayToInt64(data.([]byte))
}

type stringArrayBuilder struct {
}

func newStringArrayBuilder() *stringArrayBuilder {
	return &stringArrayBuilder{}
}

func (b *stringArrayBuilder) build(data interface{}) interface{} {
	if data == nil {
		return []string{}
	}
	arr := make([]string, 0)
	for _, b := range data.([]interface{}) {
		if b == nil {
			arr = append(arr, "")
		} else {
			arr = append(arr, string(b.([]byte)))
		}
	}
	return arr
}

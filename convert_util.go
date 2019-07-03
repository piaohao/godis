package godis

import (
	"fmt"
	"math"
	"strconv"
)

//BoolToByteArr convert bool to byte array
func BoolToByteArr(a bool) []byte {
	if a {
		return bytesTrue
	}
	return bytesFalse
}

//IntToByteArr convert int to byte array
func IntToByteArr(a int) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, int64(a), 10)
}

//Int64ToByteArr  convert int64 to byte array
func Int64ToByteArr(a int64) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, a, 10)
}

//Float64ToStr convert float64  to string
func Float64ToStr(a float64) string {
	if math.IsInf(a, 1) {
		return "+inf"
	} else if math.IsInf(a, -1) {
		return "-inf"
	} else {
		return strconv.FormatFloat(a, 'f', -1, 64)
	}
}

//Float64ToByteArr convert float64 to byte array
func Float64ToByteArr(a float64) []byte {
	var incrBytes []byte
	if math.IsInf(a, 1) {
		incrBytes = []byte("+inf")
	} else if math.IsInf(a, -1) {
		incrBytes = []byte("-inf")
	} else {
		incrBytes = []byte(strconv.FormatFloat(a, 'f', -1, 64))
	}
	return incrBytes
}

//ByteArrToFloat64 convert byte array to float64
func ByteArrToFloat64(bytes []byte) float64 {
	f, _ := strconv.ParseFloat(string(bytes), 64)
	return f
}

//StrStrArrToByteArrArr convert string and string array to byte array
func StrStrArrToByteArrArr(str string, arr []string) [][]byte {
	params := make([][]byte, 0)
	params = append(params, []byte(str))
	for _, v := range arr {
		params = append(params, []byte(v))
	}
	return params
}

//StrStrArrToStrArr convert string and string array to string array
func StrStrArrToStrArr(str string, arr []string) []string {
	params := make([]string, 0)
	params = append(params, str)
	for _, v := range arr {
		params = append(params, v)
	}
	return params
}

//StrArrToByteArrArr convert string array to byte array list
func StrArrToByteArrArr(arr []string) [][]byte {
	newArr := make([][]byte, 0)
	for _, a := range arr {
		newArr = append(newArr, []byte(a))
	}
	return newArr
}

//StrToFloat64Reply convert string reply to float64 reply
func StrToFloat64Reply(reply string, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	f, e := strconv.ParseFloat(reply, 64)
	if e != nil {
		return 0, e
	}
	return f, nil
}

//StrArrToMapReply convert string array reply to map reply
func StrArrToMapReply(reply []string, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	newMap := make(map[string]string, len(reply)/2)
	for i := 0; i < len(reply); i += 2 {
		newMap[reply[i]] = reply[i+1]
	}
	return newMap, nil
}

//Int64ToBoolReply convert int64 reply to bool reply
func Int64ToBoolReply(reply int64, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

//ByteArrToStrReply convert byte array reply to string reply
func ByteArrToStrReply(reply []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(reply), nil
}

//StrArrToTupleReply convert string array reply to tuple array reply
func StrArrToTupleReply(reply []string, err error) ([]Tuple, error) {
	if len(reply) == 0 {
		return []Tuple{}, nil
	}
	newArr := make([]Tuple, 0)
	for i := 0; i < len(reply); i += 2 {
		f, err := strconv.ParseFloat(reply[i+1], 64)
		if err != nil {
			return nil, err
		}
		newArr = append(newArr, Tuple{element: reply[i], score: f})
	}
	return newArr, err
}

//ObjArrToScanResultReply convert object array reply to scanresult reply
func ObjArrToScanResultReply(reply []interface{}, err error) (*ScanResult, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	nexCursor := string(reply[0].([]byte))
	result := make([]string, 0)
	for _, r := range reply[1].([]interface{}) {
		result = append(result, string(r.([]byte)))
	}
	return &ScanResult{Cursor: nexCursor, Results: result}, err
}

//ObjArrToGeoCoordinateReply convert object array reply to GeoCoordinate reply
func ObjArrToGeoCoordinateReply(reply []interface{}, err error) ([]*GeoCoordinate, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	arr := make([]*GeoCoordinate, 0)
	for _, r := range reply {
		if r == nil {
			arr = append(arr, nil)
		} else {
			rArr := r.([]interface{})
			lng, err := strconv.ParseFloat(string(rArr[0].([]byte)), 64)
			if err != nil {
				return nil, err
			}
			lat, err := strconv.ParseFloat(string(rArr[1].([]byte)), 64)
			if err != nil {
				return nil, err
			}
			arr = append(arr, &GeoCoordinate{
				longitude: lng,
				latitude:  lat,
			})
		}
	}
	return arr, err
}

//ObjArrToGeoRadiusResponseReply convert object array reply to GeoRadiusResponse reply
func ObjArrToGeoRadiusResponseReply(reply []interface{}, err error) ([]GeoRadiusResponse, error) {
	if err != nil || len(reply) == 0 {
		return nil, err
	}
	arr := make([]GeoRadiusResponse, 0)
	switch reply[0].(type) {
	case []interface{}:
		var resp GeoRadiusResponse
		for _, r := range reply {
			informations := r.([]interface{})
			resp = *newGeoRadiusResponse(string(informations[0].([]byte)))
			size := len(informations)
			for idx := 1; idx < size; idx++ {
				info := informations[idx]
				switch info.(type) {
				case []interface{}:
					coord := info.([]interface{})
					resp.coordinate = GeoCoordinate{
						longitude: ByteArrToFloat64(coord[0].([]byte)),
						latitude:  ByteArrToFloat64(coord[1].([]byte)),
					}
				default:
					resp.distance = ByteArrToFloat64(info.([]byte))
				}
			}
			arr = append(arr, resp)
		}
	default:
		for _, r := range reply {
			arr = append(arr, *newGeoRadiusResponse(string(r.([]byte))))
		}
	}
	return arr, err
}

//ObjArrToMapArrayReply convert object array reply to map array reply
func ObjArrToMapArrayReply(reply []interface{}, err error) ([]map[string]string, error) {
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

//ObjToEvalResult resolve response data when use script command
func ObjToEvalResult(reply interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	switch reply.(type) {
	case []byte:
		return string(reply.([]byte)), nil
	case []interface{}:
		list := reply.([]interface{})
		result := make([]interface{}, 0)
		for _, l := range list {
			evalResult, err := ObjToEvalResult(l, nil)
			if err != nil {
				return nil, err
			}
			result = append(result, evalResult)
		}
		return result, nil
	}
	return reply, err
}

//<editor-fold desc="cluster reply convert">

//ToStrReply convert object reply to string reply
func ToStrReply(reply interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	switch reply.(type) {
	case []byte:
		return string(reply.([]byte)), nil
	}
	return reply.(string), nil
}

//ToInt64Reply convert object reply to int64 reply
func ToInt64Reply(reply interface{}, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return reply.(int64), nil
}

//ToInt64ArrReply convert object reply to int64 array reply
func ToInt64ArrReply(reply interface{}, err error) ([]int64, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]int64), nil
}

//ToBoolReply convert object reply to bool reply
func ToBoolReply(reply interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return reply.(bool), nil
}

//ToFloat64Reply convert object reply to float64 reply
func ToFloat64Reply(reply interface{}, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return reply.(float64), nil
}

//ToBoolArrReply convert object reply to bool array reply
func ToBoolArrReply(reply interface{}, err error) ([]bool, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]bool), nil
}

//ToStrArrReply convert object reply to string array reply
func ToStrArrReply(reply interface{}, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]string), nil
}

//ToScanResultReply convert object reply to scanresult reply
func ToScanResultReply(reply interface{}, err error) (*ScanResult, error) {
	if err != nil {
		return nil, err
	}
	return reply.(*ScanResult), nil
}

//ToMapReply convert object reply to map reply
func ToMapReply(reply interface{}, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	return reply.(map[string]string), nil
}

//ToTupleArrReply convert object reply to tuple array reply
func ToTupleArrReply(reply interface{}, err error) ([]Tuple, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]Tuple), nil
}

//ToGeoCoordArrReply convert object reply to geocoordinate array reply
func ToGeoCoordArrReply(reply interface{}, err error) ([]*GeoCoordinate, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]*GeoCoordinate), nil
}

//ToGeoRespArrReply convert object reply to GeoRadiusResponse array reply
func ToGeoRespArrReply(reply interface{}, err error) ([]GeoRadiusResponse, error) {
	if err != nil {
		return nil, err
	}
	return reply.([]GeoRadiusResponse), nil
}

//</editor-fold>

//Builder convert pipeline|transaction response data
type Builder interface {
	build(data interface{}) (interface{}, error)
}

var (
	//StringBuilder convert interface to string
	StrBuilder = newStrBuilder()
	//Int64Builder convert interface to int64
	Int64Builder = newInt64Builder()
	//StringArrayBuilder convert interface to string array
	StrArrBuilder = newStringArrayBuilder()
)

type strBuilder struct {
}

func newStrBuilder() *strBuilder {
	return &strBuilder{}
}

func (b *strBuilder) build(data interface{}) (interface{}, error) {
	if data == nil {
		return "", nil
	}
	switch data.(type) {
	case []byte:
		return string(data.([]byte)), nil
	case error:
		return "", data.(error)
	}
	return "", fmt.Errorf("unexpected type:%T", data)
}

type int64Builder struct {
}

func newInt64Builder() *int64Builder {
	return &int64Builder{}
}

func (b *int64Builder) build(data interface{}) (interface{}, error) {
	if data == nil {
		return 0, nil
	}
	switch data.(type) {
	case int64:
		return data.(int64), nil
	}
	return 0, fmt.Errorf("unexpected type:%T", data)
}

type strArrBuilder struct {
}

func newStringArrayBuilder() *strArrBuilder {
	return &strArrBuilder{}
}

func (b *strArrBuilder) build(data interface{}) (interface{}, error) {
	if data == nil {
		return []string{}, nil
	}
	switch data.(type) {
	case []interface{}:
		arr := make([]string, 0)
		for _, b := range data.([]interface{}) {
			if b == nil {
				arr = append(arr, "")
			} else {
				arr = append(arr, string(b.([]byte)))
			}
		}
		return arr, nil
	}
	return nil, fmt.Errorf("unexpected type:%T", data)
}

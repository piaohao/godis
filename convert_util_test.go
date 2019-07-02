package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestBoolToByteArray(t *testing.T) {
	arr := BoolToByteArray(true)
	assert.Equal(t, []byte{0x31}, arr)

	arr = BoolToByteArray(false)
	assert.Equal(t, []byte{0x30}, arr)
}

func TestByteArrToStringReply(t *testing.T) {
	arr := []byte("good")
	s, e := ByteArrToStringReply(arr, nil)
	assert.Nil(t, e)
	assert.Equal(t, "good", s)

	s, e = ByteArrToStringReply(arr, newClusterMaxAttemptsError("exceed max attempts"))
	assert.NotNil(t, e)
	assert.Equal(t, "", s)
}

func TestByteArrayToFloat64(t *testing.T) {
	arr := []byte("1.1")
	f := ByteArrayToFloat64(arr)
	assert.Equal(t, 1.1, f)
}

func TestFloat64ToByteArray(t *testing.T) {
	arr := []byte("1.1")
	arr1 := Float64ToByteArray(1.1)
	assert.Equal(t, arr1, arr)

	arr1 = Float64ToByteArray(math.Inf(1))
	assert.Equal(t, []byte("+inf"), arr1)

	arr1 = Float64ToByteArray(math.Inf(-1))
	assert.Equal(t, []byte("-inf"), arr1)
}

func TestInt64ToBoolReply(t *testing.T) {
	b, e := Int64ToBoolReply(1, nil)
	assert.Nil(t, e)
	assert.Equal(t, true, b)

	b, e = Int64ToBoolReply(0, nil)
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = Int64ToBoolReply(10, nil)
	assert.Nil(t, e)
	assert.Equal(t, false, b)

	b, e = Int64ToBoolReply(0, newRedirectError("redirect too many times"))
	assert.NotNil(t, e, e.Error())
	assert.Equal(t, false, b)
}

func TestInt64ToByteArray(t *testing.T) {
}

func TestIntToByteArray(t *testing.T) {
}

func TestObjectArrToGeoCoordinateReply(t *testing.T) {
	arr, e := ObjectArrToGeoCoordinateReply(nil, newMovedDataError("move error", "localhost", 7000, 1000))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, arr)

	arr0 := make([]interface{}, 0)
	for i := 0; i < 3; i++ {
		arr0 = append(arr0, nil)
	}
	arr, e = ObjectArrToGeoCoordinateReply(arr0, nil)
	assert.Nil(t, e)
	assert.Len(t, arr, 3)

	arr0 = make([]interface{}, 0)
	item := make([]interface{}, 0)
	item = append(item, []byte("1a"))
	item = append(item, []byte("2"))
	arr0 = append(arr0, item)
	arr, e = ObjectArrToGeoCoordinateReply(arr0, nil)
	assert.NotNil(t, e)
	assert.Nil(t, arr)

	arr0 = make([]interface{}, 0)
	item = make([]interface{}, 0)
	item = append(item, []byte("1"))
	item = append(item, []byte("2b"))
	arr0 = append(arr0, item)
	arr, e = ObjectArrToGeoCoordinateReply(arr0, nil)
	assert.NotNil(t, e)
	assert.Nil(t, arr)
}

func TestObjectArrToGeoRadiusResponseReply(t *testing.T) {
	arr, e := ObjectArrToGeoRadiusResponseReply(nil, newAskDataError("move error", "localhost", 7000, 1000))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, arr)

	arr0 := make([]interface{}, 0)
	arr0 = append(arr0, []byte("a"))
	arr0 = append(arr0, []byte("b"))
	arr, e = ObjectArrToGeoRadiusResponseReply(arr0, nil)
	assert.Nil(t, e)
	assert.Len(t, arr, 2)
}

func TestObjectArrToMapArrayReply(t *testing.T) {
	objs := make([]interface{}, 0)
	for i := 0; i < 4; i++ {
		objs = append(objs, [][]byte{[]byte(fmt.Sprintf("good%d", i)), []byte(fmt.Sprintf("good%d", i+1))})
	}
	arr, e := ObjectArrToMapArrayReply(objs, nil)
	assert.Nil(t, e)
	//keyArr := make([]string, 0)
	//for k, _ := range arr {
	//	keyArr = append(keyArr, k)
	//}
	assert.Len(t, arr, 4)
}

func TestObjectArrToScanResultReply(t *testing.T) {
	result, e := ObjectArrToScanResultReply(nil, newNoReachableClusterNodeError("no reachable server"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, result)
}

func TestObjectToEvalResult(t *testing.T) {
	arr, e := ObjectToEvalResult(nil, newClusterError("error"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, arr)

	arr0 := make([]interface{}, 0)
	arr0 = append(arr0, []byte("a"))
	arr0 = append(arr0, []byte("b"))
	arr, e = ObjectToEvalResult(arr0, nil)
	assert.Nil(t, e)
	assert.Len(t, arr, 2)

	arr, e = ObjectToEvalResult("a", nil)
	assert.Nil(t, e)
	assert.Equal(t, "a", arr)
}

func TestStringArrToTupleReply(t *testing.T) {
	arr := []string{"a", "1", "b", "2"}
	tuples, e := StringArrToTupleReply(arr, nil)
	assert.Nil(t, e)
	assert.Len(t, tuples, 2)

	arr = []string{}
	tuples, e = StringArrToTupleReply(arr, nil)
	assert.Nil(t, e)
	assert.Len(t, tuples, 0)

	arr = []string{"a", "1a", "b", "2"}
	tuples, e = StringArrToTupleReply(arr, nil)
	assert.NotNil(t, e) //convert failed
	assert.Nil(t, tuples)
}

func TestStringArrayToByteArray(t *testing.T) {
}

func TestStringArrayToMapReply(t *testing.T) {
	m, e := StringArrayToMapReply(nil, newRedisError("internal error"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, m)
}

func TestStringStringArrayToByteArray(t *testing.T) {
}

func TestStringStringArrayToStringArray(t *testing.T) {
}

func TestStringToFloat64Reply(t *testing.T) {
	f, e := StringToFloat64Reply("1.1", nil)
	assert.Nil(t, e)
	assert.Equal(t, 1.1, f)

	f, e = StringToFloat64Reply("1.1a", nil)
	assert.NotNil(t, e) //not a float
	assert.Equal(t, float64(0), f)

	f, e = StringToFloat64Reply("1.1", newDataError("error data format"))
	assert.NotNil(t, e, e.Error())
	assert.Equal(t, float64(0), f)
}

func TestToBoolArrayReply(t *testing.T) {
	b, e := ToBoolArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, b)
}

func TestToBoolReply(t *testing.T) {
	b, e := ToBoolReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
	assert.Equal(t, false, b)
}

func TestToFloat64Reply(t *testing.T) {
	b, e := ToFloat64Reply(nil, newRedisError("internal error"))
	assert.NotNil(t, e, e.Error())
	assert.Equal(t, float64(0), b)
}

func TestToGeoArrayReply(t *testing.T) {
	b, e := ToGeoArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, b)
}

func TestToGeoRespArrayReply(t *testing.T) {
	b, e := ToGeoRespArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
	assert.Nil(t, b)
}

func TestToInt64ArrayReply(t *testing.T) {
	var obj interface{}
	obj = []int64{1, 2, 3}
	arr, e := ToInt64ArrayReply(obj, nil)
	assert.Nil(t, e)
	assert.Equal(t, []int64{1, 2, 3}, arr)
}

func TestToInt64Reply(t *testing.T) {
	_, e := ToGeoArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func TestToMapReply(t *testing.T) {
	_, e := ToMapReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func TestToScanResultReply(t *testing.T) {
	_, e := ToScanResultReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func TestToStringArrayReply(t *testing.T) {
	_, e := ToStringArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func TestToStringReply(t *testing.T) {
	_, e := ToStringReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func TestToTupleArrayReply(t *testing.T) {
	_, e := ToTupleArrayReply(nil, newBusyError("is busy now"))
	assert.NotNil(t, e, e.Error())
}

func Test_int64Builder_build(t *testing.T) {
	b := newInt64Builder()
	r, e := b.build(nil)
	assert.Nil(t, e)
	assert.Equal(t, 0, r)

	r, e = b.build("a")
	assert.NotNil(t, e)
	assert.Equal(t, 0, r)
}

func Test_stringArrayBuilder_build(t *testing.T) {
	b := newStringArrayBuilder()
	r, e := b.build(nil)
	assert.Nil(t, e)
	assert.Empty(t, r)

	r, e = b.build(1)
	assert.NotNil(t, e)
	assert.Equal(t, nil, r)
}

func Test_stringBuilder_build(t *testing.T) {
	b := newStringBuilder()
	r, e := b.build(nil)
	assert.Nil(t, e)
	assert.Equal(t, "", r)

	r, e = b.build(1)
	assert.NotNil(t, e)
	assert.Equal(t, "", r)
}

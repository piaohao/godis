package godis

import "strconv"

func IntToByteArray(a int) []byte {
	buf := make([]byte, 0)
	return strconv.AppendInt(buf, int64(a), 10)
}

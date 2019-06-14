package godis

import (
	"bufio"
	"strconv"
)

type RedisOutputStream struct {
	*bufio.Writer
	buf   []byte
	count int
}

func NewRedisOutputStream(bw *bufio.Writer) *RedisOutputStream {
	return &RedisOutputStream{
		Writer: bw,
		buf:    make([]byte, 0),
	}
}

//func (r *RedisOutputStream) flushBuffer() {
//	if (r.count > 0) {
//		r.Write(r.buf)
//		r.count = 0;
//	}
//}
//
//func (r *RedisOutputStream) write(b byte) {
//	if (r.count == len(r.buf)) {
//		r.flushBuffer();
//	}
//	r.count++
//	r.buf[r.count] = b;
//}

func (r *RedisOutputStream) writeIntCrLf(b int) (int, error) {
	_, err := r.Write(strconv.AppendInt(r.buf, int64(b), 10))
	if err != nil {
		return 0, err
	}
	return r.writeCrLf()
}

func (r *RedisOutputStream) writeCrLf() (int, error) {
	return r.WriteString("\r\n")
}

//
//func (r *RedisOutputStream) write(final byte[] b, final int off, final int len) throws IOException {
//if (len >= buf.length) {
//flushBuffer();
//out.write(b, off, len);
//} else {
//if (len >= buf.length - count) {
//flushBuffer();
//}
//
//System.arraycopy(b, off, buf, count, len);
//count += len;
//}
//}
//
//func (r *RedisOutputStream) writeCrLf() throws IOException {
//if (2 >= buf.length - count) {
//flushBuffer();
//}
//
//buf[count++] = '\r';
//buf[count++] = '\n';
//}
//
//func (r *RedisOutputStream) writeIntCrLf(int value) throws IOException {
//if (value < 0) {
//write((byte) '-');
//value = -value;
//}
//
//int size = 0;
//while (value > sizeTable[size])
//size++;
//
//size++;
//if (size >= buf.length - count) {
//flushBuffer();
//}
//
//int q, r;
//int charPos = count + size;
//
//while (value >= 65536) {
//q = value / 100;
//r = value - ((q << 6) + (q << 5) + (q << 2));
//value = q;
//buf[--charPos] = DigitOnes[r];
//buf[--charPos] = DigitTens[r];
//}
//
//for (;; ) {
//q = (value * 52429) >>> (16 + 3);
//r = value - ((q << 3) + (q << 1));
//buf[--charPos] = digits[r];
//value = q;
//if (value == 0) break;
//}
//count += size;
//
//writeCrLf();
//}
//
//func (r *RedisOutputStream) flush() throws IOException {
//flushBuffer();
//out.flush();
//}
type RedisInputStream struct {
	*bufio.Reader
	buf   []byte
	count int
	limit int
}

func NewRedisInputStream(br *bufio.Reader) *RedisInputStream {
	return &RedisInputStream{
		Reader: br,
		buf:    make([]byte, 8192),
	}
}

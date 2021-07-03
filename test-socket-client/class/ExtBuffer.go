package class

import (
	"bytes"
	"encoding/binary"
)

// 读写流
type ExtBuffer struct {
	index  int
	lenght int
	buf    []byte
}

// 读 byte
func (eb *ExtBuffer) ReadByte() byte {
	start := eb.index
	eb.index = eb.index + 1
	end := eb.index
	// fmt.Println(eb.buf[start:end], start, end)
	return eb.buf[start:end][0]
}

// 读 short
func (eb *ExtBuffer) ReadShort() int {
	start := eb.index
	eb.index = eb.index + 2
	end := eb.index
	return int(eb.buf[start:end][1])
}

// 读 int
func (eb *ExtBuffer) ReadInt() int {
	start := eb.index
	eb.index = eb.index + 4
	end := eb.index
	bytesBuf := bytes.NewBuffer(eb.buf[start:end])
	var x int32
	binary.Read(bytesBuf, binary.BigEndian, &x)
	return int(x)
}

// 读 string
func (eb *ExtBuffer) ReadString() string {
	strLen := eb.ReadShort()
	start := eb.index
	eb.index = eb.index + strLen
	end := eb.index
	s := string(eb.buf[start:end])
	return s
}

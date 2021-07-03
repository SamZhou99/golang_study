package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func BytesToString(b []byte) string {
	n1 := b[0:2][1]
	fmt.Println(n1)
	s := string(b[n1:4])
	n2 := b[4:6][1]
	fmt.Println(n2)
	s2 := string(b[n2:12])
	return s + s2
}

func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuf, binary.BigEndian, &x)
	return int(x)
}

func IntToBytes(n int) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, n)
	if err != nil {
		fmt.Println("IntToBytes Err = ", err)
	}
}

func Init() {
	fmt.Println("client start socket 09")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	// go func() {
	// }()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		_len := buf[0:2][1]
		_byte := buf[2:3][0]
		_short := buf[3:5][1]
		_int := BytesToInt(buf[5:9])
		_str := BytesToString(buf[9:23])
		fmt.Println(_len, _byte, _short, _int, _str, n)
	}

	time.Sleep(1000 * 10)
}

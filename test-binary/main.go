package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Package struct {
	Data int16
}

func (p *Package) Pack(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.BigEndian, &p.Data)
	return err
}
func (p *Package) Unpack(r io.Reader) error {
	var err error
	err = binary.Read(r, binary.BigEndian, &p.Data)
	return err
}

func appendString(s string, n int) string {
	b := []byte("")
	for i := 0; i < n; i++ {
		b = append(b, s...)
	}
	return string(b)
}

func main() {
	fmt.Println("start")

	// 写
	buf := new(bytes.Buffer)
	fmt.Println("空buf：", buf)

	binary.Write(buf, binary.BigEndian, 123456) //[64 9 33 251 84 68 45 24]
	fmt.Println("写buf：", buf)

	// 读
	read_buf := new(bytes.Buffer)
	binary.Read(buf, binary.BigEndian, &read_buf)
	fmt.Println("读buf：", read_buf)

	buffer1 := new(bytes.Buffer)
	pk := new(Package)
	pk.Data = 123
	pk.Pack(buffer1)
	fmt.Println(buffer1)

	buffer2 := new(bytes.Buffer)
	pk.Unpack(buffer2)
	fmt.Println(buffer2)

	fmt.Println(appendString("1", 50))
	b := []byte("abc世界")
	fmt.Println(len(b))
	fmt.Printf("%#v\n", b)
	fmt.Println(string(b[3:9]))

	fmt.Println(appendString("2", 50))
	for i, c := range []byte("abc世界") {
		fmt.Println(i, c)
	}

	fmt.Println(appendString("3", 50))
	const s = "abc\xe4\xb8\x96\xe7\x95\x8c"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}
}

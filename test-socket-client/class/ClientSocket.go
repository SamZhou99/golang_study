package class

import (
	"fmt"
	"net"
)

func Init() {
	fmt.Println("Client Init ")
	conn, err := net.Dial(Protocol, Port)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	index := 0
	for {
		buf := make([]byte, 1024)
		totleLength, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		eb := ExtBuffer{buf: buf}
		packageLength := eb.ReadShort()

		index += 1
		fmt.Println("\n---------------> 总字节长度：", totleLength, "包长：", packageLength)
		fmt.Println(index, packageLength, eb.ReadByte(), eb.ReadShort(), eb.ReadInt(), eb.ReadString(), eb.ReadString())
		totleLength -= packageLength + 2

		if totleLength > 0 {
			for totleLength >= packageLength {
				index += 1
				totleLength -= packageLength
				// fmt.Println(totleLength, packageLength)
				fmt.Println(index, eb.ReadShort(), eb.ReadByte(), eb.ReadShort(), eb.ReadInt(), eb.ReadString(), eb.ReadString())
			}
		}

	}

}

package socket

import (
	"fmt"
	"net"
	"strconv"
)

func ClientConnect(ip string, port int) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("链接异常:", err)
		return
	}
	defer conn.Close()

	// data := img
	// conn.Write([]byte(data))
}

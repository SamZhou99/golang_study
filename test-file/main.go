package main

import (
	"fmt"
	"test-file/textrw"
)

func demo_read() {
	fmt.Println("开始")

	filename := "01_test.txt"
	// result, err := textrw.ReadFile(filename)
	// result, err := textrw.ReadScanner(filename)
	result, err := textrw.ReadBytes(filename, 1024)

	if err != nil {
		fmt.Println("textrw err = ", err)
		return
	}
	fmt.Println(string(result))

	fmt.Println("结束")
}

func demo_write() {
	fmt.Println("开始")

	filename := "01_test.txt"
	// textrw.WriteFile(filename, "Hello, 世界！\n这是通过 Go 写入的文本。")
	// textrw.AppendFile(filename, "\n这是追加的内容。")
	textrw.StreamWrite(filename, "\n这是追加的内容。")

	fmt.Println("结束")
}

func main() {
	// demo_read()

	demo_write()
}

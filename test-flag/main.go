package main

import (
	"flag"
	"fmt"
)

var aValue *string = flag.String("a", "aValue", "测试 a 参数")
var bValue *string = flag.String("b", "bValue", "测试 b 参数")

func main() {
	fmt.Println("命令行参数的使用")

	flag.Parse()
	if aValue != nil || bValue != nil {
		fmt.Println("a = ", *aValue)
		fmt.Println("b = ", *bValue)
	}
}

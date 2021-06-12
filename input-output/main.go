package main

import (
	"fmt"
	"strconv"
)

func input_number() uint64 {
	for true {
		var str string
		fmt.Print("\n请输入一个整数：")
		fmt.Scanln(&str)

		num, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			fmt.Println("请输入有效的整数", err)
			continue
		}
		return num
	}
	return 0
}

func main() {
	num := input_number()
	fmt.Println("输入的数：", num)
}

// 知识点：字符串 转 整数
// strconv.Atoi(str) // int
// strconv.ParseUint(str, 10, 32) // to uint64
// uint(strconv.ParseUint(str, 10, 32)) // to uint32

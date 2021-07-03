package main

import (
	"fmt"
	"strings"
)

func main() {
	// 单引号 是 字符，双引号 是 字符串
	// Go 字符有两种：1、uint8或者叫[]byte代表ascii一个字符。2、rune类型，代表一个UFT-8字符。

	s1 := "abc"
	s2 := "中国"
	s3 := s1 + s2

	// 字符串长度
	fmt.Println(len(s1), len(s2), len(s3))
	// 分割
	fmt.Println(strings.Split(s3, ""))
	// 拼接操作
	fmt.Println(strings.Join(strings.Split(s3, ""), ", "))
	// 是否包含
	fmt.Println(strings.Contains(s3, "9"), strings.Contains(s3, "中"))
	// 前缀，后缀
	fmt.Println(strings.HasPrefix(s3, "abc"), strings.HasSuffix(s3, "中国"))
	// 字符串出现的位置
	fmt.Println(strings.Index(s3, "a"), strings.LastIndex(s3, "国"))

	// len 和 range 迭代的长度不一样
	fmt.Println("=================")
	for i := 0; i < len(s3); i++ {
		fmt.Printf("索引:%d, ascii:%v, Unicode:%U\n", i, s3[i], s3[i])
	}

	fmt.Println("=================")
	for _, c := range s3 {
		fmt.Printf("%c\n", c)
	}

	// 修改字符串
	cs1 := "小白兔"
	cs2 := []rune(cs1) // 转成 []rune 切片
	cs2[1] = '灰'       // 注意：要用 单引号 是字符
	fmt.Println(cs1, string(cs2))
}

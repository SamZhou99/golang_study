package main

import "fmt"

func main() {
	// []string 是切片，直接赋值，得到的是 引用变量。
	var a = []string{"A", "B", "C"}
	var b = a
	b[0] = "D" // 会影响 a 的值
	fmt.Println(a, b)

	// [size]int{}是数组，直接赋值，得到的是，复制变量。
	c := [4]int{1, 2, 3, 4}
	d := c
	d[0] = 9 //不影响 c 的值
	fmt.Println(c, d)

	// 数组 引用变量 可以用"&,*"号，可以影响之前的赋值。
	e := &c
	e[0] = 999
	fmt.Println(c, *e)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

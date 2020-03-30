package main

import(
	"fmt"
	"runtime"
	"os"
)

func main()  {
	// go版本
	fmt.Println(runtime.Version())
	// 系统名
	fmt.Println(runtime.GOOS)
	// 环境变量
	fmt.Println(os.Getenv("PATH"))
}
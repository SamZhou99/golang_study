package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5) // 缓冲通道 3
	capNum := cap(c)
	go func() {
		defer fmt.Println("goroutine 结束")
		// 注意看打印结果，当读一个 i 后，这里的goroutine才会结束
		for i := 0; i < capNum; i++ {
			fmt.Println("goroutine 正在运行", len(c), capNum)
			c <- i // 将i写入通道
		}

		close(c) // 尝试关闭 通道 //会影响打印结果

	}()

	time.Sleep(1000000 * time.Microsecond)

	for i := 0; i < capNum; i++ {
		num := <-c // 从通道中读取i
		fmt.Printf("num = %d\n", num)
	}

	// // 上面可以简写成range迭代通道数据
	// for data := range c {
	// 	fmt.Println(data)
	// }

	// // 多路监控通道的数据
	// select {
	// case <-c1:
	// 	//
	// case <-c2:
	// 	//
	// default:
	// 	//
	// }

	fmt.Println("main goroutine 结束")
	// close(c)
}

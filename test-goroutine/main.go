package main

import (
	"fmt"
	"time"
)

var count int32

func say(n int32) {
	for i := 0; i < 600; i++ {
		count += n
		// fmt.Println(n, i, count)
		// time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go say(1)
	go say(2)
	say(3)
	time.Sleep(9 * time.Millisecond)
	fmt.Printf("结果：count=%d", count)
	time.Sleep(1000 * time.Millisecond)
}

package main

import (
	"fmt"
	"time"
)

var count int32

func say(n int32) {
	for i := 0; i < 1000; i++ {
		count += n
		// fmt.Println(n, i, count)
		// time.Sleep(5 * time.Millisecond)
	}
}

func main() {
	go say(1)
	go say(2)
	go say(2)
	say(5)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("结果：count=%d", count)
	time.Sleep(1000 * time.Millisecond)
}

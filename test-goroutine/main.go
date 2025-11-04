package main

import (
	"fmt"
	"time"
)

var global_count int = 0

func say(n int, c chan int) {
	var max int = 100000000
	var total int = n * max
	var count int = 0
	var i int = 0
	for i = 0; i < max; i++ {
		count += n
		// fmt.Println(n, i, count)
		// time.Sleep(2 * time.Millisecond)
	}
	global_count += count
	c <- count / total
}

func main() {
	c := make(chan int)

	var cInt []int
	iLen := 36
	for i := 1; i <= iLen; i++ {
		go say(i, c)
		cInt = append(cInt, <-c)
	}

	fmt.Println(cInt)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("global_count =", global_count)
	time.Sleep(100 * time.Millisecond)
}

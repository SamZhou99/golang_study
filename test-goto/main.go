package main

import "fmt"

func main() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 6 {
		goto HERE
	}
}

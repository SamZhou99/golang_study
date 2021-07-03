package main

import (
	"fmt"
	"test-function/math2"
)

func main() {
	r1, err1 := math2.Add(2, 4)
	r2, err2 := math2.Add(-2, 4)
	fmt.Println(r1, err1)
	fmt.Println(r2, err2)
}

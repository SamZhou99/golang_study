package main

import (
	"./pkg"
	"fmt"
)

func main() {
	fmt.Println(2 * pkg.PI)
	fmt.Println(pkg.RectArea(5, 7))
	fmt.Println(pkg.MathMax())
}

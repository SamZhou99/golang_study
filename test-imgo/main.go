package main

import (
	"fmt"

	"github.com/Comdex/imgo"
)

func main() {
	cos, err := imgo.CosineSimilarity("aa.png", "aa copy.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(cos)
}

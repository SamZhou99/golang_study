package main

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	mr "math/rand"
	"time"
)

func randomInt(min, max int) int {
	result, _ := cr.Int(cr.Reader, big.NewInt(int64(max-min)))
	return min + int(result.Int64())

}

func randomInt2(min, max, seen int) int {
	mr.Seed(time.Now().UnixNano() + int64(seen))
	return min + mr.Intn(max-min)
}

func sortInt(sa []int) []int {
	a := []int{}
	for i := 0; i < len(sa)-1; i++ {
		// a[i], a[i+1] = sa[i+1], sa[i]
		a = append(a, sa[i+1], sa[i])
	}
	return a
}

func main() {
	arr := []int{}
	arr2 := []int{}
	for i := 0; i < 5; i++ {
		arr = append(arr, randomInt(1, 999))
		arr2 = append(arr2, randomInt2(1, 999, i))
	}

	fmt.Println(arr, sortInt(arr))
	fmt.Println(arr2, sortInt(arr2))
}

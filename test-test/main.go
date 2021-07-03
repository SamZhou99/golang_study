package main

import (
	"fmt"
	"math/rand"
	"test-test/sort"
	"time"
)

func main() {
	fmt.Println("Sort")

	const MAX int = 999

	var values, values2 []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < MAX; i++ {
		r := rand.Intn(MAX)
		values = append(values, []int{r}...)
	}

	values2 = append(values, []int{}...)

	fmt.Println("\n原数据：", values)

	start := time.Now()
	sort.QuickSort(values)
	fmt.Println("\n快排结果：", values)
	elapsed := time.Since(start)
	fmt.Println("耗时：", elapsed)

	start2 := time.Now()
	sort.BubbleSort(values2)
	fmt.Println("\n冒泡结果：", values2)
	elapsed2 := time.Since(start2)
	fmt.Println("耗时：", elapsed2)
}

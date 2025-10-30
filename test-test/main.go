package main

import (
	"fmt"
	"math/rand"
	"strconv"
	compTime "test-test/computing-time"
	mathSort "test-test/sort"
	"time"
)

func randLen(MaxNum int) []int {
	var values []int
	rand.New(rand.NewSource(time.Now().Unix())) // rand.Seed(time.Now().Unix()) // 旧的放弃方法
	for i := 0; i < MaxNum; i++ {
		r := rand.Intn(MaxNum)
		values = append(values, []int{r}...)
	}
	return values
}

func arrToStr(values []int, a int) string {
	s1 := fmt.Sprintf("%v", values[:a])
	s2 := fmt.Sprintf("%v", values[len(values)-a:])
	s3 := strconv.Itoa(len(values) - a*2)
	return s1 + "[..." + s3 + "]" + s2
}

func main() {
	fmt.Println("排序算法 计算时间 比较")

	var val1 = randLen(9999)
	var val2 = make([]int, len(val1))
	copy(val2, val1)

	fmt.Println("\n原数据计算总长度：", len(val1))

	fmt.Println("快排数据：", arrToStr(val1, 5))
	compTime.Start()
	mathSort.QuickSort(val1)
	fmt.Println("快排耗时：", compTime.End())

	fmt.Println("冒泡数据：", arrToStr(val2, 5))
	compTime.Start()
	mathSort.BubbleSort(val2)
	fmt.Println("冒泡耗时：", compTime.End())
}

package img

import "math/rand"

// 随机数
func Random(start int, end int) int {
	num1 := start + rand.Intn(end)
	return num1
}

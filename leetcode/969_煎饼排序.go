/*
题目大意
给定一个数组，要求输出“煎饼排序”的步骤，使得最终数组是从小到大有序的。“煎饼排序”，每次排序都反转前 n 个数，n 小于数组的长度。

解题思路
这道题的思路是，每次找到当前数组中无序段中最大的值，（初始的时候，整个数组相当于都是无序段），将最大值的下标 i 进行“煎饼排序”，前 i 个元素都反转一遍。这样最大值就到了第一个位置了。然后紧接着再进行一次数组总长度 n 的“煎饼排序”，目的是使最大值到数组最后一位，这样它的位置就归位了。那么数组的无序段为 n-1 。然后用这个方法不断的循环，直至数组中每个元素都到了排序后最终的位置下标上了。最终数组就有序了。

这道题有一个特殊点在于，数组里面的元素都是自然整数，那么最终数组排序完成以后，数组的长度就是最大值。所以找最大值也不需要遍历一次数组了，直接取出长度就是最大值。
*/

package main

import "fmt"

func pancakeSort(A []int) []int {
	if len(A) == 0 {
		return []int{}
	}
	right := len(A)
	var (
		ans []int
	)
	for right > 0 {
		idx := find(A, right)
		if idx != right-1 {
			reverse969(A, 0, idx)
			reverse969(A, 0, right-1)
			ans = append(ans, idx+1, right)
		}
		right--
	}

	return ans
}

func reverse969(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func find(nums []int, t int) int {
	for i, num := range nums {
		if num == t {
			return i
		}
	}
	return -1
}

func main() {
	r := pancakeSort([]int{1, 2, 3, 9, 8, 7, 6, 5, 4})
	fmt.Println(r)
}

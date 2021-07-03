/*
1. 两数之和
题目大意
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

解题思路
这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。
*/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		// 计算结果，看MAP里有没有，没有就，把当前值和索引[值:索引]，存到MAP中，等待下次匹配计算结果。
		fmt.Println(target, nums[i], another, m[another], m)
		if _, ok := m[another]; ok {
			return []int{m[another], i} // 返回索引，并非值。
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	target := 16
	result := twoSum(nums, target)
	fmt.Println("结果：", nums[result[0]], nums[result[1]])
}

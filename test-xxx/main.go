package main

import (
	"fmt"
	"strconv"
)

func main() {
	dayCount := int(input_number())
	if dayCount >= 50 {
		dayCount = 50
	}
	currentMoneys := 0
	totalMoneys := 0
	for i := 1; i <= dayCount; i++ {
		if currentMoneys > 1 {
			currentMoneys += currentMoneys
		} else {
			currentMoneys = i
		}
		totalMoneys += currentMoneys
		println("第"+strconv.Itoa(i)+"天", "给"+strconv.Itoa(currentMoneys)+"元", "共计"+strconv.Itoa(totalMoneys))
	}
}

func input_number() uint64 {
	for true {
		var str string
		fmt.Print("\n请输入一个整数(1-50)：")
		fmt.Scanln(&str)

		num, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			fmt.Println("请输入有效的整数", err)
			continue
		}
		return num
	}
	return 0
}

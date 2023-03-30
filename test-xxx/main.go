package main

import "strconv"

func main() {
	dayCount := 20
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

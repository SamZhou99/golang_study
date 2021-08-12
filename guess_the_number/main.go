package main

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	mr "math/rand"
	"os"
	"strconv"
	"time"
)

type Int int

func (i_int Int) ToString() string {
	return strconv.Itoa(int(i_int))
}

func randomInt(min, max int) int {
	result, _ := cr.Int(cr.Reader, big.NewInt(int64(max-min)))
	return min + int(result.Int64())

}

func randomInt2(min, max, seen int) int {
	mr.Seed(time.Now().UnixNano() + int64(seen))
	return min + mr.Intn(max-min)
}

func inputNumber(label string) int64 {
	for {
		var str string
		fmt.Print(label)
		fmt.Scanln(&str)

		if str == "exit" {
			os.Exit(0)
		}

		num, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			fmt.Println("请输入有效的整数", err)
			continue
		}
		return num
	}
}

func main() {
	fmt.Println("1 main...")
	fmt.Println("随机1：", randomInt(100, 999), ",   随机2：", randomInt2(100, 999, 0))
	rand_num := int64(randomInt(100, 999))
	// fmt.Println("随机数是：", rand_num)
	for {
		num := inputNumber("请输入一个整数：")

		if rand_num > num {
			fmt.Println(Int(num).ToString() + " 小了")
		} else if rand_num < num {
			fmt.Println(Int(num).ToString() + " 大了")
		} else if rand_num == num {
			fmt.Println("game over")
			break
		}
	}
}

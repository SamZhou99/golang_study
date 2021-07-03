package math2

import (
	"errors"
	"strconv"
)

func Add(a int, b int) (res int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Add(" + strconv.Itoa(a) + ", " + strconv.Itoa(b) + ") 函数不支持负数计算！")
		return
	}
	return a + b, nil
}

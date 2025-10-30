package math2

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// 获取函数名称
func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)       // 获取f的reflect.Value对象
	if v.Kind() == reflect.Func { // 确保f是函数类型
		// 获取函数名称
		funName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		start := strings.LastIndex(funName, ".")
		return funName[start+1:]
	} else {
		// fmt.Println("Not a function")
		return "Not a function"
	}
}

// 函数传递函数作为参数
// 函数可以作为参数传递给其他函数
// 函数可以作为返回值从其他函数返回

func add(x int) int {
	return x + 1
}
func sub(x int) int {
	return x - 1
}
func mul(x int) int {
	return x * 2
}
func div(x int) int {
	return x / 2
}

func apply(f func(int) int, x int) {
	result := f(x)
	fmt.Println("函数", nameOf(f), "参数", x, "结果", result)
}

func Demo() {
	fmt.Println("函数传递函数作为参数")
	apply(add, 5)
	apply(sub, 5)
	apply(mul, 5)
	apply(div, 5)
}

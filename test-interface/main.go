package main

// 接口，多态现象

import (
	"fmt"
)

// 先定义一个接口，接口本身就是 指针 类型
type AnmalIF interface {
	Sleep()
	GetColor() string
}

// Cat类
type Cat struct {
	color string
}

// 实现AnmalIF接口 Sleep 和 GetColor
func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}

// Dog类
type Dog struct {
	color string
}

// 实现AnmalIF接口 Sleep 和 GetColor
func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}

// *AnmalIF没有加*这样写，因为他本身就是指针。
func showAnmal(anmal AnmalIF) {
	anmal.Sleep()
	fmt.Println("Color = ", anmal.GetColor())
}

// 另外interface空接口，是一种万能类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is Called")
	fmt.Println(arg)

	// 给 interface 提供 类型断言 的机制
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg type is %T", value)
	} else {
		fmt.Println("arg is not string type")
	}
}

func main() {
	// &是指针类型
	cat := &Cat{"Red"}
	dog := &Dog{"Yellow"}

	showAnmal(cat)
	showAnmal(dog)

	myFunc(showAnmal)
	myFunc("my is string")
	myFunc(3.141592)
}

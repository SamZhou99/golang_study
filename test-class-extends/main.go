package main

import "fmt"

// 一个 Integer 类
type Integer int

// Integer.Less(Integer) 方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 父类
type Human struct {
	name string
	sex  string
}

// 方法
func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Print() {
	fmt.Println("Human.Println()...")
}

// 子类
type SuperMan struct {
	Human //继承父类
	level int
}

// 子类方法 覆盖父类方法Print
func (this *SuperMan) Print() {
	this.Human.Print()
	fmt.Println("SuperMan.Print()...")
	fmt.Println(this.name)
	fmt.Println(this.sex)
	fmt.Println(this.level)
}

func main() {
	// 一个超简单的类
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	// 超人类继承了男人类
	s := SuperMan{Human{"Peter", "男"}, 900}
	s.Print()
}

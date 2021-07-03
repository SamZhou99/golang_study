package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	person := make(map[string]PersonInfo)
	person["1"] = PersonInfo{"001", "Peter", "room:a"}
	person["2"] = PersonInfo{"002", "Allen", "room:b"}
	person["3"] = PersonInfo{"003", "Sam", "room:c"}
	fmt.Println(person)

	// 删除 map["2"], 删除一个没有值, 也不会异常。
	delete(person, "2")
	fmt.Println(person)

	// 查找 特定的键值，是否存在。
	p, ok := person["4"]
	if ok {
		fmt.Println("p 的数据是：", p)
	} else {
		fmt.Println("p 没有数据")
	}

	// 迭代 map
	for key, value := range person {
		fmt.Println(key, value)
	}
}

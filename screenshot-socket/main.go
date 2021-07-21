package main

import (
	"fmt"
	"screenshot-socket/app/json"
)

func main() {
	data, err := json.Read("config.json")
	if err != nil {
		fmt.Println("JSON读取异常：", err)
		return
	}
	fmt.Println(data)
}

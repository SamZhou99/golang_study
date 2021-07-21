package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonInfo struct {
	Name    string
	Age     int32
	Sex     bool
	Hobbies []string
}

func writeFile(fileName string, data interface{}) bool {
	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件创建失败：", err.Error())
		return false
	}
	defer file.Close()

	// 创建Json编码器
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		fmt.Println("编码失败：", err.Error())
		return false
	}
	fmt.Println("编码成功")
	return true
}

func readFile(fileName string) []PersonInfo {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]", err.Error())
		return nil
	}
	defer file.Close()

	var person []PersonInfo

	// 创建json解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("解码失败：", err.Error())
	} else {
		fmt.Println("解码成功")
	}
	return person
}

func main() {
	fmt.Println("init")

	fileName := "test.json"

	_, err := os.Stat(fileName)
	if !(err == nil || os.IsExist(err)) {
		personInfo := []PersonInfo{
			{"David", 30, true, []string{"跑步", "读书", "看电影"}},
			{"Lee", 27, false, []string{"工作", "睡觉", "玩游戏"}}}
		writeFile(fileName, personInfo)
	}

	personInfoJson := readFile(fileName)
	fmt.Println(personInfoJson)
	fmt.Println(personInfoJson[0].Hobbies[0])
}

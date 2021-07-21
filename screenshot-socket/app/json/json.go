package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// 读取JSON数据
func Read(fileName string) (Config, error) {
	var data Config

	file, err := os.Open(fileName)
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

// 写入JSON数据
func Write(fileName string) bool {
	fmt.Println("fileName", fileName)
	return true
}

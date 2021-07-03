package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(fileName string) bool {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ioutil.ReadFile err = ", err)
		return false
	}
	fmt.Println(data)
	return true
}

func writeFile(fileName string, data []byte) bool {
	err := ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Println("ioutil.WriteFile err = ", err)
		return false
	}
	return true
}

func main() {
	fmt.Println("main init")
	// byteValue := make([]byte, 1)
	// byteValue[0] = 1
	// shortValue := make([]byte, 2)
	// shortValue[0] = 2

	// data := append(byteValue, shortValue...)

	// writeFile("01_test.txt", data)
	readFile("01_test.txt")
}

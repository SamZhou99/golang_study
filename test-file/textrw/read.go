package textrw

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件内容（返回 []byte）
func ReadFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("读取失败：%v\n", err)
		return nil, err
	}
	return content, nil
}

// 逐行读取（适合大文件，节省内存）
func ReadScanner(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("打开文件失败：%v\n", err)
		return "", err
	}
	defer file.Close()

	// 创建 scanner 用于逐行读取
	scanner := bufio.NewScanner(file)
	lineNum := 1

	// 循环读取每行
	var lineStr = ""
	for scanner.Scan() {
		lineStr = lineStr + scanner.Text() // 获取当前行文本（字符串）
		// 或用 scanner.Bytes() 获取字节切片
		// fmt.Printf("第 %d 行：%s\n", lineNum, line)
		lineNum++
	}

	// 检查扫描过程中是否出错
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取错误：%v\n", err)
	}

	return lineStr, nil
}

// 按字节 / 块读取（灵活控制读取大小）
func ReadBytes(fileName string, size int) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("打开文件失败：%v\n", err)
		return "", err
	}
	defer file.Close()

	// // 创建字节切片用于存储读取数据
	// buf := make([]byte, size)
	// n, err := file.Read(buf)
	// if err != nil {
	// 	fmt.Printf("读取失败：%v\n", err)
	// 	return nil, err
	// }
	// return buf[:n], nil

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	var n = 0
	for {
		n, err := reader.Read(buf) // 读取数据到缓冲区
		if err != nil {
			if err == io.EOF { // 读取到文件末尾
				break
			}
			return "", err
		}
		// 打印读取到的内容（注意只取前 n 字节，避免缓冲区残留数据）
		fmt.Print(string(buf[:n]))
	}
	return string(buf[:n]), nil
}

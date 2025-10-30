package textrw

import (
	"bufio"
	"fmt"
	"os"
)

// 一次性写入（覆盖原有内容）
func WriteFile(fileName string, text string) {
	// 写入文件（权限设为 0644，即所有者可读写，其他人可读）
	err := os.WriteFile(fileName, []byte(text), 0644)
	if err != nil {
		fmt.Printf("写入失败：%v\n", err)
		return
	}
	fmt.Println("写入成功！")
}

// 追加写入（在文件末尾添加内容）
func AppendFile(fileName string, text string) {
	// content := "这是追加写入的内容。"

	// // 追加写入文件（权限设为 0644，即所有者可读写，其他人可读）
	// err := os.WriteFile("output.txt", []byte(content), 0644)
	// if err != nil {
	// 	fmt.Printf("追加写入失败：%v\n", err)
	// 	return
	// }
	// fmt.Println("追加写入成功！")

	// 打开文件：追加模式（O_APPEND）、只写（O_WRONLY）、不存在则创建（O_CREATE）
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开文件失败：%v\n", err)
		return
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString(text) // 直接写入字符串（也可用 Write 写入字节切片）
	if err != nil {
		fmt.Printf("追加失败：%v\n", err)
		return
	}
	fmt.Println("追加成功！")
}

// 流式写入（逐行或分块写入）
func StreamWrite(fileName string, text string) {
	// // 打开文件：写入模式（O_WRONLY）、不存在则创建（O_CREATE）、如果文件存在则清空内容（O_TRUNC）
	// file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// if err != nil {
	// 	fmt.Printf("打开文件失败：%v\n", err)
	// 	return
	// }
	// defer file.Close()

	// // 写入内容（分块写入）
	// chunkSize := 1024 // 每次写入 1024 字节
	// for i := 0; i < 10; i++ {
	// 	data := fmt.Sprintf("这是第 %d 行内容。\n", i)
	// 	_, err := file.WriteString(data)
	// 	if err != nil {
	// 		fmt.Printf("写入失败：%v\n", err)
	// 		return
	// 	}
	// }
	// fmt.Println("流式写入成功！")

	// 打开文件（覆盖模式，若需追加可加 os.O_APPEND）
	file, err := os.Create(fileName) // Create 等价于 OpenFile(..., os.O_WRONLY|os.O_CREATE|os.O_TRUNC, ...)
	if err != nil {
		fmt.Printf("创建文件失败：%v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // 创建带缓冲的写入器

	// 逐行写入
	lines := []string{"第一行", "第二行", "第三行"}
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n") // 写入字符串并换行
		if err != nil {
			fmt.Printf("写入失败：%v\n", err)
			return
		}
	}

	// 手动刷新缓冲区（确保数据写入磁盘，否则可能残留于缓冲区）
	err = writer.Flush()
	if err != nil {
		fmt.Printf("刷新缓冲区失败：%v\n", err)
		return
	}
	fmt.Println("流式写入成功！")
}

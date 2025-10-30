// package main

// import (
// 	"archive/zip"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	fileName := "test.zip"
// 	r, err := zip.OpenReader(fileName)
// 	if err != nil {
// 		// fmt.Printf("Content %s:\n", f.Name)
// 		fmt.Printf("打开文件失败 %s", fileName)
// 		log.Fatal(err)
// 	}
// 	defer r.Close()

// 	for _, f := range r.File {
// 		fmt.Printf("Content %s:\n", f.Name)
// 		rc, err := f.Open()
// 		if err != nil {
// 			fmt.Printf("文件解析失败")
// 			log.Fatal(err)
// 		}

// 		_, err = io.CopyN(os.Stdout, rc, 68)
// 		if err != nil {
// 			fmt.Printf("文件解析过程中失败")
// 			log.Fatal(err)
// 		}

// 		rc.Close()
// 		fmt.Println()
// 	}
// }

package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Zip 压缩功能：将源路径（文件或目录）压缩到目标 zip 文件
func Zip(srcPath string, destZip string) error {
	// 创建目标 zip 文件
	destFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 创建 zip 写入器
	zipWriter := zip.NewWriter(destFile)
	defer zipWriter.Close()

	// 遍历源路径，将所有文件添加到 zip 中
	return filepath.Walk(srcPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建 zip 中的文件头
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// 处理目录的路径格式（确保压缩后的目录结构正确）
		// 计算相对路径，避免压缩时包含源路径的上级目录
		relativePath, err := filepath.Rel(filepath.Dir(srcPath), filePath)
		if err != nil {
			return err
		}
		header.Name = relativePath

		// 如果是目录，需要标记为目录（否则解压时可能无法识别）
		if fileInfo.IsDir() {
			header.Name += "/"
		} else {
			// 设置压缩方法（默认不压缩，这里指定 DEFLATED 进行压缩）
			header.Method = zip.Deflate
		}

		// 创建 zip 中的文件写入器
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是目录，不需要写入内容
		if fileInfo.IsDir() {
			return nil
		}

		// 打开源文件
		srcFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// 将源文件内容写入 zip
		_, err = io.Copy(writer, srcFile)
		return err
	})
}

// Unzip 解压缩功能：将 zip 文件解压到目标目录
func Unzip(zipFile string, destDir string) error {
	// 打开 zip 文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 遍历 zip 中的所有文件
	for _, file := range reader.File {
		// 获取 zip 中文件的信息
		fileInfo := file.FileInfo()
		// 计算解压后的目标路径
		destPath := filepath.Join(destDir, file.Name)

		// 如果是目录，创建对应的目录
		if fileInfo.IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// 确保目标文件的父目录存在
		if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
			return err
		}

		// 打开 zip 中的文件
		srcFile, err := file.Open()
		if err != nil {
			return err
		}

		// 创建目标文件
		destFile, err := os.Create(destPath)
		if err != nil {
			srcFile.Close()
			return err
		}

		// 将 zip 中的文件内容写入目标文件
		_, err = io.Copy(destFile, srcFile)
		// 关闭文件（无论是否出错都需要关闭）
		srcFile.Close()
		destFile.Close()
		if err != nil {
			return err
		}

		// 保留源文件的权限
		if err := os.Chmod(destPath, fileInfo.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 示例：压缩单个文件或目录
	srcPath := "./testdir"    // 要压缩的源路径（可以是文件或目录）
	zipPath := "./output.zip" // 压缩后的 zip 文件路径
	if err := Zip(srcPath, zipPath); err != nil {
		panic("压缩失败：" + err.Error())
	}
	println("压缩成功，zip 文件路径：", zipPath)

	// 示例：解压缩 zip 包
	unzipDir := "./unzipped" // 解压后的目标目录
	if err := Unzip(zipPath, unzipDir); err != nil {
		panic("解压失败：" + err.Error())
	}
	println("解压成功，目标目录：", unzipDir)
}

package main

import (
	"fmt"
	"test-image/img"
)

func main() {

	// img.Demo()

	img_path := []string{"./assets/001 (1).jpg", "./assets/001 (1).png"}
	for i, v := range img_path {
		fmt.Println(i, v)
		b, f, _ := img.GetFileByte(v)

		img_fmt, _ := img.GetImageFormatByHeader(b)
		fmt.Println("-------格式", img_fmt)
		imgd, e := img.ImgDecode(f)
		if e != nil {
			fmt.Println("-------无法解码图片", e)
			f.Close()
			return
		}
		img.ImgInfo(imgd)

		f.Close()
	}

}

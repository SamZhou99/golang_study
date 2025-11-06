package main

import (
	"fmt"
	"test-image/img"
)

func main() {

	// img.Demo()

	img_path := []string{"./assets/001.jpeg", "./assets/001.png", "./assets/001.webp"}
	for i, v := range img_path {
		fmt.Println(i, v)
		b, f, e := img.GetFileByte(v)
		fmt.Println(b, f, e)
		img_fmt, err := img.GetImageFormatByHeader(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-------", img_fmt)
	}
	// b, f, e := img.GetFileByte(img_path)
	// fmt.Println(b, f, e)
	// file := img.OpenFile(img_path)
	// local_img := img.ReadImg(file)
	// img.ImgInfo(local_img)

}

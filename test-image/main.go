package main

import "test-image/img"

func main() {

	// img.Demo()

	imgPath := "./assets/000.png"
	file := img.OpenFile(imgPath)
	local_img := img.ReadImg(file)
	img.ImgInfo(local_img)

}

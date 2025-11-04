package img

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func OpenFile(imgPath string) *os.File {
	file, err := os.Open(imgPath)
	if err != nil {
		panic("无法打开图片：" + imgPath + "，错误信息：" + err.Error())
	}
	defer file.Close()
	return file
}

func ReadImg(file *os.File) image.Image {
	a, format, err := image.DecodeConfig(file)
	fmt.Println("【", a, "|", format, "|", err, "】")
	if err != nil {
		panic("1图片解码失败：" + err.Error())
	}

	var img image.Image
	switch format {
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		panic("不支持的图片格式：" + format)
	}

	if err != nil {
		panic("图片解码失败：" + err.Error())
	}
	return img
}

func ImgInfo(img image.Image) {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	fmt.Printf("图片，宽度：%d，高度：%d\n", width, height)
}

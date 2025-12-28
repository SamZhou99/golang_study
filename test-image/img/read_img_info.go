package img

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)

func ReadImageInfo(index int, imgPath string) image.Image {
	file, err := os.Open(imgPath)
	if err != nil {
		panic("无法打开图片：" + imgPath + "，错误信息：" + err.Error())
	}
	defer file.Close()

	img_data, format, err := image.Decode(file)
	if err != nil {
		panic("图片解析异常：" + err.Error())
	}
	fmt.Println("索引："+strconv.Itoa(index), "图片路径："+imgPath, "图片格式："+format)
	return img_data
}

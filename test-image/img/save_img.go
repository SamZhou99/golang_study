package img

import (
	"image"
	"image/png"
	"os"
	"strconv"
	"time"
)

// 保存图片
func SaveImg(imgRgba image.Image, index int) {
	now := time.Now()
	nowTime := now.Format("20060102_150405")
	outImgFileName := "./red_" + nowTime + "_" + strconv.Itoa(index) + ".png"

	file, err := os.Create(outImgFileName)
	if err != nil {
		panic("os.Create failed: " + err.Error())
	}
	defer file.Close()

	err = png.Encode(file, imgRgba)
	if err != nil {
		panic("png.Encode failed: " + err.Error())
	}

	println("图片保存为 " + outImgFileName)
}

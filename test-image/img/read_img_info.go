package img

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func ReadImageInfo(index int, imgPath string) image.Image {
	file, err := os.Open(imgPath)
	if err != nil {
		showError("请检查图片路径和文件名：", err)
	}
	defer file.Close()

	// img_data, format, err := image.Decode(file)
	img_data, _, err := image.Decode(file)
	if err != nil {
		showError("图片解析异常：", err)
	}
	// fmt.Println("索引："+strconv.Itoa(index), "图片路径："+imgPath, "格式："+format)
	return img_data
}

func showError(tipstr string, err error) {
	panic("\n💔💔💔\n" + tipstr + err.Error() + "\n")
}

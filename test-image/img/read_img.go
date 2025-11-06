package img

import (
	"bufio"
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

func ReadImg(filePath string) image.Image {
	// a, format, err := image.DecodeConfig(file)
	// imge, format, err := image.Decode(file)
	b, file, e := GetFileByte(filePath)
	if e != nil {
		panic("获取字节失败：" + e.Error())
	}
	format, err := GetImageFormatByHeader(b)
	fmt.Println("【", "|", format, "|", err, "】")
	if err != nil {
		panic("1图片解码失败：" + err.Error())
	}

	var imgD image.Image
	switch format {
	case "jpeg":
		imgD, err = jpeg.Decode(file)
	case "png":
		imgD, err = png.Decode(file)
	default:
		panic("不支持的图片格式：" + format)
	}

	if err != nil {
		panic("图片解码失败：" + err.Error())
	}
	return imgD
}

func ImgInfo(img image.Image) {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	fmt.Printf("图片，宽度：%d，高度：%d\n", width, height)
}

func GetFileByte(filePath string) ([]byte, *os.File, error) {
	header := make([]byte, 16)
	file, err := os.Open(filePath)
	if err != nil {
		return header, file, err
	}
	defer file.Close()

	// 读取文件头前 16 字节（足够判断大部分格式）
	n, err := bufio.NewReader(file).Read(header)
	if err != nil || n < 8 { // 至少需要 8 字节判断常见格式
		return header, file, fmt.Errorf("无法读取文件头")
	}
	return header, file, err
}

// 判断图片格式（通过文件头）
func GetImageFormatByHeader(header []byte) (string, error) {
	// 检查 PNG（文件头前 8 字节）
	if header[0] == 0x89 && header[1] == 0x50 && header[2] == 0x4E && header[3] == 0x47 &&
		header[4] == 0x0D && header[5] == 0x0A && header[6] == 0x1A && header[7] == 0x0A {
		return "png", nil
	}

	// 检查 JPEG（文件头前 3 字节）
	if header[0] == 0xFF && header[1] == 0xD8 && header[2] == 0xFF {
		return "jpeg", nil
	}

	// 检查 GIF（文件头前 6 字节）
	if (header[0] == 0x47 && header[1] == 0x49 && header[2] == 0x46 &&
		header[3] == 0x38 && header[4] == 0x39 && header[5] == 0x61) || // GIF89a
		(header[0] == 0x47 && header[1] == 0x49 && header[2] == 0x46 &&
			header[3] == 0x38 && header[4] == 0x37 && header[5] == 0x61) { // GIF87a
		return "gif", nil
	}

	// 检查 BMP（文件头前 2 字节）
	if header[0] == 0x42 && header[1] == 0x4D {
		return "bmp", nil
	}

	// 检查 WebP（文件头前 12 字节：RIFF + 4字节大小 + WEBP）
	if header[0] == 0x52 && header[1] == 0x49 && header[2] == 0x46 && header[3] == 0x46 &&
		header[8] == 0x57 && header[9] == 0x45 && header[10] == 0x42 && header[11] == 0x50 {
		return "webp", nil
	}

	return "unknown", nil
}

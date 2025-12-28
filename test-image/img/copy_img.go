package img

import "image"

func CopyImg(bounds image.Rectangle, img_data image.Image) *image.RGBA {
	// 复制原图像素
	rgbaImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x, y, img_data.At(x, y))
		}
	}
	return rgbaImg
}

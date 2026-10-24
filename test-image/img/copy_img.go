package img

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

// 旧方法
// CopyImg 复制指定区域的图像数据，返回新的 RGBA 图像对象
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

// CropImage 从源图片中裁剪指定坐标和宽高的区域，返回独立的新图片
// 参数说明：
//
//	src: 源图片
//	x, y: 裁剪区域左上角坐标（像素，原点在图片左上角）
//	w, h: 裁剪区域的宽度和高度（像素）
//
// 返回值：裁剪后的独立图片对象，错误信息
func CropImage(src image.Image, x, y, w, h int) (image.Image, error) {
	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	// 1. 前置边界校验，提前拦截非法参数避免运行时panic
	if x < 0 || y < 0 || w <= 0 || h <= 0 {
		return nil, fmt.Errorf("裁剪参数非法: x=%d, y=%d, w=%d, h=%d，坐标不能为负，宽高必须大于0", x, y, w, h)
	}
	if x+w > srcW || y+h > srcH {
		return nil, fmt.Errorf("裁剪区域超出原图范围: 原图尺寸%dx%d，裁剪区域%dx%d 位置(%d,%d)", srcW, srcH, w, h, x, y)
	}

	// 2. 高性能快速路径：针对 RGBA 格式直接内存拷贝（适配绝大多数场景）
	if rgbaSrc, ok := src.(*image.RGBA); ok {
		dst := image.NewRGBA(image.Rect(0, 0, w, h))
		srcStride := rgbaSrc.Stride
		dstStride := dst.Stride
		// 单行像素的总字节数：宽度 * 4 通道（R/G/B/A）
		rowBytes := w * 4

		for i := 0; i < h; i++ {
			// 计算源图当前行的起始字节偏移
			srcOffset := (y+i)*srcStride + x*4
			// 计算目标图当前行的起始字节偏移
			dstOffset := i * dstStride
			// 内置 copy 执行内存级批量拷贝，性能远高于逐像素读写
			copy(dst.Pix[dstOffset:dstOffset+rowBytes], rgbaSrc.Pix[srcOffset:srcOffset+rowBytes])
		}

		if w > h {
			fmt.Println("1 需要旋转图片。。。。。。。。。。。。。。。。。。。。。。。。", w, h)
			// 旋转后宽高互换：新宽 = 原高，新高 = 原宽
			rotated := image.NewRGBA(image.Rect(0, 0, h, w))
			rotStride := rotated.Stride

			for cy := 0; cy < h; cy++ {
				for cx := 0; cx < w; cx++ {
					// 顺时针 90 度坐标变换公式：newX = 原高-1 - y, newY = x
					dstX := h - 1 - cy
					dstY := cx

					srcIdx := cy*dstStride + cx*4
					dstIdx := dstY*rotStride + dstX*4

					// 直接操作字节数组拷贝像素，跳过方法调用开销
					rotated.Pix[dstIdx] = dst.Pix[srcIdx]
					rotated.Pix[dstIdx+1] = dst.Pix[srcIdx+1]
					rotated.Pix[dstIdx+2] = dst.Pix[srcIdx+2]
					rotated.Pix[dstIdx+3] = dst.Pix[srcIdx+3]
				}
			}
			return rotated, nil
		}

		return dst, nil
	}

	// 3. 通用兼容路径：支持所有 image.Image 类型的兜底方案
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	dstStride := dst.Stride
	// 源图上的裁剪矩形
	srcCropRect := image.Rect(x, y, x+w, y+h)
	// 将源图指定区域绘制到新图片的原点位置
	draw.Draw(dst, dst.Bounds(), src, srcCropRect.Min, draw.Src)

	// 需要旋转的图片
	if w > h {
		// 旋转后宽高互换：新宽 = 原高，新高 = 原宽
		rotated := image.NewRGBA(image.Rect(0, 0, h, w))
		rotStride := rotated.Stride

		for cy := 0; cy < h; cy++ {
			for cx := 0; cx < w; cx++ {
				// 顺时针 90 度坐标变换公式：newX = 原高-1 - y, newY = x
				dstX := h - 1 - cy
				dstY := cx

				srcIdx := cy*dstStride + cx*4
				dstIdx := dstY*rotStride + dstX*4

				// 直接操作字节数组拷贝像素，跳过方法调用开销
				rotated.Pix[dstIdx] = dst.Pix[srcIdx]
				rotated.Pix[dstIdx+1] = dst.Pix[srcIdx+1]
				rotated.Pix[dstIdx+2] = dst.Pix[srcIdx+2]
				rotated.Pix[dstIdx+3] = dst.Pix[srcIdx+3]
			}
		}
		return rotated, nil
	}

	return dst, nil
}

// 使用示例
func CopyImageDemo() {
	// 打开源图片文件
	file, err := os.Open("./demo-2.jpg")
	if err != nil {
		panic(fmt.Sprintf("打开图片失败: %v", err))
	}
	defer file.Close()

	// 解码图片
	srcImg, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Sprintf("解码图片失败: %v", err))
	}

	// 裁剪：左上角(100, 50)，宽200px，高150px的区域
	croppedImg, err := CropImage(srcImg, 100, 50, 200, 150)
	if err != nil {
		panic(fmt.Sprintf("裁剪失败: %v", err))
	}

	// 保存裁剪后的图片
	outFile, err := os.Create("cropped_result.png")
	if err != nil {
		panic(fmt.Sprintf("创建输出文件失败: %v", err))
	}
	defer outFile.Close()

	if err := png.Encode(outFile, croppedImg); err != nil {
		panic(fmt.Sprintf("保存图片失败: %v", err))
	}

	fmt.Println("裁剪完成，结果已保存为 cropped_result.png")
	fmt.Printf("裁剪后尺寸: 宽=%dpx 高=%dpx\n", croppedImg.Bounds().Dx(), croppedImg.Bounds().Dy())
}

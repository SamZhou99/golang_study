package img

import (
	"fmt"
	"image"
	"image/draw"
)

// StretchImage 将图片强制拉伸到指定宽高（不保持原比例）
// 参数：
//
//	src: 源图片
//	targetWidth: 目标宽度（像素）
//	targetHeight: 目标高度（像素）
//
// 返回：拉伸后的 RGBA 图片，以及错误信息
func StretchImage(src image.Image, targetWidth, targetHeight int) (*image.RGBA, error) {
	// 1. 参数合法性校验
	if targetWidth <= 0 || targetHeight <= 0 {
		return nil, fmt.Errorf("目标宽高必须大于0，当前 宽=%d 高=%d", targetWidth, targetHeight)
	}

	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	if srcW == 0 || srcH == 0 {
		return nil, fmt.Errorf("源图片尺寸无效")
	}

	// 2. 源图统一转为 RGBA 格式，复用已有类型避免冗余转换
	var srcRGBA *image.RGBA
	if rgba, ok := src.(*image.RGBA); ok {
		srcRGBA = rgba
	} else {
		srcRGBA = image.NewRGBA(srcBounds)
		draw.Draw(srcRGBA, srcBounds, src, srcBounds.Min, draw.Src)
	}
	srcStride := srcRGBA.Stride

	// 3. 预分配目标图片内存
	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	dstStride := dst.Stride

	// 计算缩放比例
	scaleX := float64(srcW) / float64(targetWidth)
	scaleY := float64(srcH) / float64(targetHeight)

	// 4. 双线性插值拉伸，按行遍历优化CPU缓存命中率
	for dy := 0; dy < targetHeight; dy++ {
		// 计算源图Y方向浮点坐标与权重
		srcY := float64(dy) * scaleY
		y0 := int(srcY)
		y1 := y0 + 1
		yWeight := srcY - float64(y0)

		// 边界钳位，防止数组越界
		if y0 < 0 {
			y0 = 0
		}
		if y1 >= srcH {
			y1 = srcH - 1
		}

		// 预计算两行的像素起始偏移，减少循环内重复计算
		y0RowOffset := y0 * srcStride
		y1RowOffset := y1 * srcStride
		dstRowOffset := dy * dstStride

		for dx := 0; dx < targetWidth; dx++ {
			// 计算源图X方向浮点坐标与权重
			srcX := float64(dx) * scaleX
			x0 := int(srcX)
			x1 := x0 + 1
			xWeight := srcX - float64(x0)

			// 边界钳位
			if x0 < 0 {
				x0 = 0
			}
			if x1 >= srcW {
				x1 = srcW - 1
			}

			// 四个相邻像素的内存偏移
			p00 := y0RowOffset + x0*4 // 左上
			p10 := y0RowOffset + x1*4 // 右上
			p01 := y1RowOffset + x0*4 // 左下
			p11 := y1RowOffset + x1*4 // 右下

			// 逐通道执行双线性插值
			for channel := 0; channel < 4; channel++ {
				v00 := float64(srcRGBA.Pix[p00+channel])
				v10 := float64(srcRGBA.Pix[p10+channel])
				v01 := float64(srcRGBA.Pix[p01+channel])
				v11 := float64(srcRGBA.Pix[p11+channel])

				// X方向两次线性插值
				top := v00*(1-xWeight) + v10*xWeight
				bottom := v01*(1-xWeight) + v11*xWeight
				// Y方向一次线性插值
				val := top*(1-yWeight) + bottom*yWeight

				// 四舍五入后写入目标像素
				dst.Pix[dstRowOffset+dx*4+channel] = uint8(val + 0.5)
			}
		}
	}

	return dst, nil
}

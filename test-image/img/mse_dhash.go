package img

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

// ====================== 算法1：感知哈希(dHash) 快速相似度 ======================

// ImageSimilarity 计算两张图片的感知相似率，返回值范围 0.0~1.0，1.0 表示完全相同
// 基于差异感知哈希算法，抗缩放、压缩、轻微裁剪和亮度变化，适合快速相似图检索
func ImageSimilarity(src, target image.Image) (float64, error) {
	hash1, err := dHash(src)
	if err != nil {
		return 0, err
	}
	hash2, err := dHash(target)
	if err != nil {
		return 0, err
	}

	distance := hammingDistance(hash1, hash2)
	// 64位哈希，相似率 = 1 - 差异位数 / 总位数
	similarity := 1.0 - float64(distance)/64.0
	return similarity, nil
}

// dHash 计算图片的64位差异感知哈希值
func dHash(img image.Image) (uint64, error) {
	// 统一缩放到 9x8 像素：宽9高8，横向相邻像素对比产生8*8=64位特征
	const hashWidth, hashHeight = 9, 8
	grayPixels := resizeToGray(img, hashWidth, hashHeight)
	if grayPixels == nil {
		return 0, fmt.Errorf("图片尺寸无效")
	}

	var hash uint64
	idx := 0
	for y := 0; y < hashHeight; y++ {
		rowOffset := y * hashWidth
		for x := 0; x < hashWidth-1; x++ {
			// 左侧像素 > 右侧像素则标记为1
			if grayPixels[rowOffset+x] > grayPixels[rowOffset+x+1] {
				hash |= 1 << idx
			}
			idx++
		}
	}
	return hash, nil
}

// resizeToGray 快速缩放图片并转为灰度图，返回行优先的灰度像素数组
// 采用最近邻采样，哈希场景下速度优先，精度足够
func resizeToGray(img image.Image, targetW, targetH int) []uint8 {
	srcBounds := img.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	if srcW == 0 || srcH == 0 {
		return nil
	}

	// 源图统一转RGBA，直接操作像素内存
	var srcRGBA *image.RGBA
	if rgba, ok := img.(*image.RGBA); ok {
		srcRGBA = rgba
	} else {
		srcRGBA = image.NewRGBA(srcBounds)
		draw.Draw(srcRGBA, srcBounds, img, srcBounds.Min, draw.Src)
	}
	srcStride := srcRGBA.Stride

	gray := make([]uint8, targetW*targetH)
	scaleX := float64(srcW) / float64(targetW)
	scaleY := float64(srcH) / float64(targetH)

	for dy := 0; dy < targetH; dy++ {
		srcY := int(float64(dy) * scaleY)
		if srcY >= srcH {
			srcY = srcH - 1
		}
		srcRowOff := srcY * srcStride
		dstRowOff := dy * targetW

		for dx := 0; dx < targetW; dx++ {
			srcX := int(float64(dx) * scaleX)
			if srcX >= srcW {
				srcX = srcW - 1
			}
			off := srcRowOff + srcX*4
			r := srcRGBA.Pix[off]
			g := srcRGBA.Pix[off+1]
			b := srcRGBA.Pix[off+2]
			// 整数灰度公式：0.299R + 0.587G + 0.114B，规避浮点运算
			gray[dstRowOff+dx] = uint8((int(r)*299 + int(g)*587 + int(b)*114) / 1000)
		}
	}
	return gray
}

// hammingDistance 计算两个64位整数的汉明距离（二进制位不同的数量）
func hammingDistance(a, b uint64) int {
	xor := a ^ b
	count := 0
	for xor != 0 {
		count += int(xor & 1)
		xor >>= 1
	}
	return count
}

// ====================== 算法2：像素级MSE 精确相似度 ======================

// ImageSimilarityMSE 基于均方误差计算像素级相似率，返回 0.0~1.0
// 要求两张图片尺寸完全一致，适合精确校验、像素级差异检测
func ImageSimilarityMSE(src, target image.Image) (float64, error) {
	srcBounds := src.Bounds()
	tgtBounds := target.Bounds()
	srcW, srcH := srcBounds.Dx(), srcBounds.Dy()
	tgtW, tgtH := tgtBounds.Dx(), tgtBounds.Dy()

	if srcW != tgtW || srcH != tgtH {
		return 0, fmt.Errorf("尺寸不一致，MSE对比需同尺寸：源(%dx%d) 目标(%dx%d)", srcW, srcH, tgtW, tgtH)
	}

	// 统一转RGBA格式
	var srcRGBA, tgtRGBA *image.RGBA
	if rgba, ok := src.(*image.RGBA); ok {
		srcRGBA = rgba
	} else {
		srcRGBA = image.NewRGBA(srcBounds)
		draw.Draw(srcRGBA, srcBounds, src, srcBounds.Min, draw.Src)
	}
	if rgba, ok := target.(*image.RGBA); ok {
		tgtRGBA = rgba
	} else {
		tgtRGBA = image.NewRGBA(tgtBounds)
		draw.Draw(tgtRGBA, tgtBounds, target, tgtBounds.Min, draw.Src)
	}

	totalPixels := srcW * srcH
	var mse float64

	// 逐像素计算RGB三通道均方误差，忽略Alpha通道
	for i := 0; i < totalPixels*4; i += 4 {
		sr := int(srcRGBA.Pix[i])
		sg := int(srcRGBA.Pix[i+1])
		sb := int(srcRGBA.Pix[i+2])
		tr := int(tgtRGBA.Pix[i])
		tg := int(tgtRGBA.Pix[i+1])
		tb := int(tgtRGBA.Pix[i+2])

		dr := sr - tr
		dg := sg - tg
		db := sb - tb
		mse += float64(dr*dr + dg*dg + db*db)
	}

	mse /= float64(totalPixels * 3)
	// 均方根误差归一化到 0~1，转为相似率
	similarity := 1.0 - math.Sqrt(mse)/255.0
	if similarity < 0 {
		similarity = 0
	}
	return similarity, nil
}

// ====================== 使用示例 ======================

func main2() {
	// 加载两张测试图片
	img1, err := loadImage("img1.png")
	if err != nil {
		panic(err)
	}
	img2, err := loadImage("img2.png")
	if err != nil {
		panic(err)
	}

	// 1. 快速感知相似度
	simHash, err := ImageSimilarity(img1, img2)
	if err != nil {
		fmt.Println("感知哈希对比失败:", err)
	} else {
		fmt.Printf("感知哈希相似率：%.2f%%\n", simHash*100)
	}

	// 2. 像素级精确相似度（需同尺寸）
	simMSE, err := ImageSimilarityMSE(img1, img2)
	if err != nil {
		fmt.Println("像素对比提示:", err)
	} else {
		fmt.Printf("像素级相似率：%.2f%%\n", simMSE*100)
	}
}

// loadImage 辅助函数：加载并解码图片
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("打开图片失败: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("解码图片失败: %v", err)
	}
	return img, nil
}

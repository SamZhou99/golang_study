package main

import (
	"fmt"
	"image"
	"strconv"
	"test-image/img"
	"time"
)

type PokerInfo struct {
	Right float64
	Name  string
}

func pokerImgs() map[string]image.Image {
	pk_img := map[string]image.Image{}

	// 读取识别模板图片
	img_index := 0
	for _, k := range []string{"a", "b", "c", "d"} {
		for n := 1; n <= 13; n++ {
			key := k + strconv.Itoa(n)
			img_path := "./assets/pk/" + key + ".png"
			img_data := img.ReadImageInfo(img_index, img_path)
			pk_img[key] = img_data
			img_index++
			// fmt.Println("img_index:", img_index, "key:", key, "img_path:", img_path)
		}
	}

	return pk_img
}

func main() {
	// 开始时间
	startTime := time.Now()

	// 扑克默认拉伸宽高
	PKW := 70
	PKH := 100

	// // // 生成渐变色图片
	// // img.Demo()

	// 读取扑克目录图片，并修改大小为 70x100
	pkImgs := pokerImgs()
	for k, v := range pkImgs {
		// 将裁剪后的图片强制拉伸到指定宽高（不保持原比例）
		stretchImg, err := img.StretchImage(v, PKW, PKH)
		if err != nil {
			continue
		}
		pkImgs[k] = stretchImg
	}

	// 读取需要解析的图片
	srcImg := img.ReadImageInfo(1, "./assets/source/001 (4).png")

	// 处理图片为灰度图
	srcImg = img.Filter(srcImg)

	// 找到符合条件的白色区域
	rects := img.FindRect(srcImg)

	// 裁剪每个白色区域并保存（横图变竖图）
	for i, rect := range rects {
		cropImg, err := img.CropImage(srcImg, rect.X, rect.Y, rect.W, rect.H)
		if err != nil {
			fmt.Printf("裁剪第%d个区域失败: %v\n", i+1, err)
			continue
		}

		// 比较结果
		temp := PokerInfo{Right: 0.0, Name: ""}
		for k, v := range pkImgs {
			// // 计算两张图片的感知相似率
			// res, err := img.ImageSimilarity(cropImg, v)

			// 将图片拉伸到一样的大小再进行比较
			stretchedImg, err := img.StretchImage(cropImg, PKW, PKH)
			if err != nil {
				fmt.Printf("拉伸第%d个区域失败: %v\n", i+1, err)
				continue
			}

			// 基于均方误差计算像素级相似率
			res, err := img.ImageSimilarityMSE(v, stretchedImg)
			if err != nil {
				fmt.Println("对比失败:", err)
			} else if res > temp.Right {
				temp.Right = res
				temp.Name = k
			}

		}
		fmt.Printf("相似率：%.4f%% (%s)\n", temp.Right*100, img.ConvertCardString(temp.Name))

		// // 保存
		// img.SaveImg(cropImg, i)
	}

	// 结束时间
	duration := time.Since(startTime)
	fmt.Printf("微秒：%d µs\n", duration.Milliseconds())

}

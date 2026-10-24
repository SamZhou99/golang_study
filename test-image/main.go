package main

import (
	"fmt"
	"image"
	"strconv"
	"test-image/img"
	"time"
)

func showArray2d() {
	s := [10][10]int{}
	for y := 0; y < 10; y += 2 {
		for x := 0; x < 10; x += 2 {
			s[y][x] = 1
			fmt.Print("|", y, x, "|")
		}
		fmt.Println("")
	}
	fmt.Println(s)
}

func readImages() {
	img_pk := map[string]image.Image{}

	// 读取识别模板图片
	img_index := 0
	for _, k := range []string{"a", "b", "c", "d"} {
		for n := 1; n <= 13; n++ {
			key := k + strconv.Itoa(n)
			img_path := "./assets/pk/" + key + ".png"
			img_data := img.ReadImageInfo(img_index, img_path)
			img_pk[key] = img_data
			img_index++
			// fmt.Println("img_index:", img_index, "key:", key, "img_path:", img_path)
		}
	}

	// 读取需要解析的图片
	img_path := []string{
		"001 (1).jpg",
		// "001 (1).png",
		// "001 (2).jpg",
		// "001 (2).png",
	}
	for i, v := range img_path {
		img_data := img.ReadImageInfo(i, "./assets/source/"+v)
		img.WriteImageInfo(i, img_data, img_pk)
	}
}

func main() {
	// 开始时间
	startTime := time.Now()
	fmt.Println(startTime)

	// // 二维数据示例
	// showArray2d()

	// // // 生成渐变色图片
	// // img.Demo()

	// // 目录图片读取
	// readImages()

	// var reaultArr = img.FindRect("./assets/source/001 (1).png")

	// 读取需要解析的图片
	srcImg := img.ReadImageInfo(1, "./assets/source/001 (1).jpg")

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
		// 后续可对 cropImg 进行保存、识别等处理
		img.SaveImg(cropImg, i)
	}

	// 结束时间
	duration := time.Since(startTime)
	fmt.Printf("微秒：%d µs\n", duration.Milliseconds())

}

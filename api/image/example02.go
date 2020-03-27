package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const (
	dx = 500
	dy = 200
)

func main() {
	file, err := os.Create("example02.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			rgba.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}

	fmt.Println(rgba.At(400, 100))    //144在指定位置的像素
	fmt.Println(rgba.Bounds())        //图片边界
	fmt.Println(rgba.Opaque())        //是否完全透明
	fmt.Println(rgba.PixOffset(1, 1)) //指定点相对于第一个点的距离
	fmt.Println(rgba.Stride)          //两个垂直像素之间的距离

	jpeg.Encode(file, rgba, nil) //将image信息写入文件中
}

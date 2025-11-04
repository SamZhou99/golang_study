package img

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

const (
	DX = 256 * 2
	DY = 256
)

func imgRGBA() *image.RGBA {
	alpha := image.NewRGBA(image.Rect(0, 0, DX, DY))
	for x := 0; x < DX; x++ {
		for y := 0; y < DY; y++ {
			alpha.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), uint8(x % 256), uint8(x % 256)}) //设定alpha图片的透明度 color.Alpha{uint8(x % 256)
		}
	}
	fmt.Println(alpha.At(400, 100))    //144 在指定位置的像素
	fmt.Println(alpha.Bounds())        //(0,0)-(500,200)，图片边界
	fmt.Println(alpha.Opaque())        //false，是否图片完全透明
	fmt.Println(alpha.PixOffset(1, 1)) //501，指定点相对于第一个点的距离
	fmt.Println(alpha.Stride)          //500，两个垂直像素之间的距离
	return alpha
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	alpha := imgRGBA()
	jpeg.Encode(file, alpha, nil) //将image信息写入文件中
}

func Demo() {
	createFile("./test.jpeg")
}

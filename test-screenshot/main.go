package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func saveToPng(fileName string, img *image.RGBA) {
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
}

func main() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		// bounds := screenshot.GetDisplayBounds(i) // 全屏
		bounds := image.Rect(300, 300, 300+800, 300+500) // 截图范围
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		saveToPng(fileName, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

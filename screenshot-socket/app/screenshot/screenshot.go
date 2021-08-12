package screenshot

import (
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func Screenshot(x1 int, y1 int, x2 int, y2 int) (*image.RGBA, error) {
	bounds := image.Rect(x1, y1, x2, y2)
	return screenshot.CaptureRect(bounds)
}

func SaveToPng(fileName string, img *image.RGBA) bool {
	file, err := os.Create(fileName)
	if err != nil {
		return false
	}
	defer file.Close()
	png.Encode(file, img)
	return true
}

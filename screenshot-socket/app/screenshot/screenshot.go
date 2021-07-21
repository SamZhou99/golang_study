package screenshot

import (
	"image"

	"github.com/kbinani/screenshot"
)

func Screenshot(x1 int, y1 int, x2 int, y2 int) (image.Image, error) {
	bounds := image.Rect(x1, y1, x2, y2)
	return screenshot.CaptureRect(bounds)

}

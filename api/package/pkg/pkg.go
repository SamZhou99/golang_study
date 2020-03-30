package pkg

import (
	"math"
)

var PI float64

func init() {
	PI = 4 * math.Atan(1)
}

func RectArea(w int, h int) int {
	return w * h
}

func MathMax() int {
	u8 := math.MaxUint8
	return u8
}

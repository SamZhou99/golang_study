package img

import (
	"image"

	"github.com/disintegration/gift"
)

func Filter(img_data image.Image) image.Image {
	filters := map[string]gift.Filter{
		"rotate_90":         gift.Rotate90(),           //旋转90
		"contrast_increase": gift.Contrast(60),         //对比度
		"grayscale":         gift.Grayscale(),          //灰度
		"unsharp_mask":      gift.UnsharpMask(1, 1, 0), //锐度
		"gamma":             gift.Gamma(0.45),          //伽马
		"threshold":         gift.Threshold(30),        //阈值
	}

	giftA := gift.New(
		filters["unsharp_mask"],
		filters["gamma"],
		filters["threshold"],
		// filters["contrast_increase"],
		// filters["grayscale"],
		// filters["contrast_increase"],
	)

	giftDst := image.NewNRGBA(img_data.Bounds())
	giftA.Draw(giftDst, img_data)

	return giftDst
}

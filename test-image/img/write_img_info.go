package img

import (
	"image"
)

func WriteImageInfo(index int, img_data image.Image, img_pk map[string]image.Image) {
	// 加上滤镜
	giftDst := Filter(img_data)

	// 区域打上红点
	RedDot(index, giftDst)

}

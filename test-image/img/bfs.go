package img

import (
	"image"
	"image/color"
)

// // 广度优先搜索，找到白色区域的边界
// func Bfs(img image.Image, x, y int, visited [][]bool) (int, int, int, int) {
// 	bounds := img.Bounds()
// 	minX, minY, maxX, maxY := x, y, x, y
// 	queue := [][2]int{{x, y}}
// 	visited[y][x] = true

// 	for len(queue) > 0 {
// 		pos := queue[0]
// 		queue = queue[1:]
// 		curX, curY := pos[0], pos[1]

// 		// 更新边界
// 		if curX < minX {
// 			minX = curX
// 		}
// 		if curY < minY {
// 			minY = curY
// 		}
// 		if curX > maxX {
// 			maxX = curX
// 		}
// 		if curY > maxY {
// 			maxY = curY
// 		}

// 		// 遍历四个方向
// 		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 		for _, d := range dirs {
// 			nx, ny := curX+d[0], curY+d[1]
// 			if nx >= 0 && nx < bounds.Max.X && ny >= 0 && ny < bounds.Max.Y && !visited[ny][nx] {
// 				c := img.At(nx, ny)
// 				r, g, b, _ := c.RGBA()
// 				if isWhite(r, g, b) {
// 					visited[ny][nx] = true
// 					queue = append(queue, [][2]int{{nx, ny}}...)
// 				}
// 			}
// 		}
// 	}
// 	return minX, minY, maxX, maxY
// }

func Bfs2(img image.Image, startX, startY int, visited [][]bool) (int, int, int, int) {
	bounds := img.Bounds()
	maxW, maxH := bounds.Max.X, bounds.Max.Y
	// minX, minY, maxX, maxY := x, y, x, y
	rgbaImg := CopyImg(bounds, img)
	var x, y = startX, startY
	for y = startY; y < maxH; y++ {
		startX = repStartX(img, startX, y, maxW)
		for x = startX; x < maxW; x++ {
			if isWhite(img, x, y) {
				setImgBlue(rgbaImg, x, y)
			} else {
				x = maxW + 1
			}
		}
	}
	SaveImg(rgbaImg, 20)
	return 0, 0, 0, 0
}

func setImgBlue(rgbaImg *image.RGBA, x, y int) {
	colorImg := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	rgbaImg.Set(x, y, colorImg)
}

func isWhite(img image.Image, x, y int) bool {
	c := img.At(x, y)
	r, g, b, _ := c.RGBA()
	r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
	return r8 > 200 && g8 > 200 && b8 > 200
}

func repStartX(img image.Image, startX, y, maxW int) int {
	if !isWhite(img, startX, y) {
		return startX
	}
	for startX > 0 {
		x := startX - 1
		if !isWhite(img, startX, y) {
			return startX
		}
		startX = x
	}
	return startX
}

// // 判断是否为白色（RGB 通道接近 255）
// func isWhite(r, g, b uint32) bool {
// 	r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
// 	return math.Abs(float64(r8-255)) < 10 &&
// 		math.Abs(float64(g8-255)) < 10 &&
// 		math.Abs(float64(b8-255)) < 10
// }

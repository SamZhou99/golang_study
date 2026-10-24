package img

import (
	"fmt"
	"image"
	"image/color"
)

// 打上红点
func RedDot(index int, img_data image.Image) {
	bounds := img_data.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	// 标记已访问的像素
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	v := make([][]uint, height)
	for i := range v {
		v[i] = make([]uint, width)
	}

	img_rgba := CopyImg(bounds, img_data)

	// ix, iy, mx, my := 0, 0, 0, 0
	imgColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	stepLength := 2
	for y := 0; y < height; y += stepLength {
		for x := 0; x < width; x += stepLength {
			r, g, b, _ := img_data.At(x, y).RGBA()
			r, g, b = r>>8, g>>8, b>>8
			if r >= 250 && g >= 250 && b >= 250 {
				// 打红点🔴
				img_rgba.Set(x, y, imgColor)
				v[y][x] = 1
			} else {
				v[y][x] = 0
			}
		}
	}

	// 查找一个区域起点
	startX, startY := findDot(v)
	img_rgba.Set(startY, startX, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img_rgba.Set(startY+1, startX, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img_rgba.Set(startY, startX+1, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img_rgba.Set(startY+1, startX+1, color.RGBA{R: 0, G: 255, B: 0, A: 255})

	fmt.Println(">>>", startX, startY, "<<<")
	minX, minY, maxX, maxY := Bfs2(img_data, startX, startY, visited)
	fmt.Println(">>>", minX, minY, maxX, maxY, "<<<")
	// fmt.Printf(">>>>>>%T<<<<<<", visited)

	// 保存修改后的图片
	SaveImg(img_data, index)
	SaveImg(img_rgba, index+10)
}

func findDot(arr [][]uint) (int, int) {
	stepLength := 1
	height := len(arr)
	width := len(arr[0])
	var y, x int
	for y = 0; y < height; y += stepLength {
		for x = 0; x < width; x += stepLength {
			u := arr[y][x]
			if u == 1 {
				return x, y
			}
		}
	}
	return x, y
}

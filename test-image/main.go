package main

import (
	"fmt"
	"image"
	"strconv"
	"test-image/img"
	"time"
)

func main() {
	// s := [10][10]int{}
	// for y := 0; y < 10; y += 2 {
	// 	for x := 0; x < 10; x += 2 {
	// 		s[y][x] = 1
	// 		fmt.Print("|", y, x, "|")
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Println(s)

	startTime := time.Now()
	// img.Demo()

	img_pk := map[string]image.Image{}
	img_index := 0
	for _, k := range []string{"a", "b", "c", "d"} {
		for n := 1; n <= 13; n++ {
			key := k + strconv.Itoa(n)
			img_data := img.ReadImageInfo(img_index, "./assets/pk/"+key+".png")
			img_pk[key] = img_data
			img_index++
		}
	}

	img_path := []string{
		"001 (1).jpg",
		// "001 (1).png",
		// "001 (2).jpg",
		// "001 (2).png",
	}
	for i, v := range img_path {
		img_data := img.ReadImageInfo(i, "./assets/source/"+v)
		img.WriteImageInfo(i, img_data, img_pk)
	}

	duration := time.Since(startTime)
	fmt.Printf("微秒：%d µs\n", duration.Milliseconds())
}

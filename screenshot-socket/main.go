// package main

// import (
// 	"fmt"
// 	"screenshot-socket/app/json"
// 	"screenshot-socket/app/screenshot"
// )

// func readConfig(configFile string) json.Config {
// 	// 读取 json
// 	data, err := json.Read(configFile)
// 	if err != nil {
// 		fmt.Println("JSON读取异常：", err)
// 		return data
// 	}
// 	fmt.Println(data)
// 	return data
// }

// func screenshotImg(data json.Config, imgFile string) {
// 	// 屏幕截图
// 	left := data.Screenshot.Left
// 	top := data.Screenshot.Top
// 	width := data.Screenshot.Width
// 	height := data.Screenshot.Height
// 	imgRGBA, err := screenshot.Screenshot(left, top, left+width, top+height)
// 	if err != nil {
// 		fmt.Println("screenshot.Screenshot err", err)
// 		return
// 	}
// 	// 保存图片
// 	saveResult := screenshot.SaveToPng(imgFile, imgRGBA)
// 	fmt.Println("保存图片", saveResult)
// }

// func main() {
// 	screenshotImg(readConfig("config.json"), "./test-output.png")
// }

package main

import (
	"fmt"
	"syscall"
)

const (
	SM_CXSCREEN = uintptr(0) // X Size of screen
	SM_CYSCREEN = uintptr(1) // Y Size of screen
	// ZOOM        = uintptr(2)
	DESKTOPHORZRES = 0
	DESKTOPVERTRES = 1
)

func main() {
	user32 := syscall.NewLazyDLL(`User32.dll`)
	w, _, _ := user32.NewProc(`GetSystemMetrics`).Call(SM_CXSCREEN)
	h, _, _ := user32.NewProc(`GetSystemMetrics`).Call(SM_CYSCREEN)
	dc, _, _ := user32.NewProc(`GetDC`).Call(0)
	defer user32.NewProc(`ReleaseDC`).Call(dc)

	gdi32 := syscall.NewLazyDLL(`gdi32.dll`)
	dw, _, _ := gdi32.NewProc(`GetDeviceCaps`).Call(dc, uintptr(118))
	dh, _, _ := gdi32.NewProc(`GetDeviceCaps`).Call(dc, uintptr(117))
	fmt.Println(int(w), int(h), dw, dh)
}

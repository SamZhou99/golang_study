package main

import (
	"fmt"
	"screenshot-socket/app/json"
	"screenshot-socket/app/screenshot"
)

func readConfig(configFile string) json.Config {
	// 读取 json
	data, err := json.Read(configFile)
	if err != nil {
		fmt.Println("JSON读取异常：", err)
		return data
	}
	fmt.Println(data)
	return data
}

func screenshotImg(data json.Config, imgFile string) {
	// 屏幕截图
	left := data.Screenshot.Left
	top := data.Screenshot.Top
	width := data.Screenshot.Width
	height := data.Screenshot.Height
	imgRGBA, err := screenshot.Screenshot(left, top, left+width, top+height)
	if err != nil {
		fmt.Println("screenshot.Screenshot err", err)
		return
	}
	// 保存图片
	saveResult := screenshot.SaveToPng(imgFile, imgRGBA)
	fmt.Println("保存图片", saveResult)
}

func main() {
	screenshotImg(readConfig("config.json"), "./test-output.png")
}

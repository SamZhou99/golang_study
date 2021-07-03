package main

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

func main() {
	fmt.Println("start")
	req := HttpRequest.NewRequest()
	result, err := req.Get("https://wo99.oss-cn-qingdao.aliyuncs.com/temp/test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json, err := result.Body()

	fmt.Println("%s", string(json))
}

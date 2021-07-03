package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	// "fmt"
	"strconv"
	"strings"
)

var markdownContent = []byte(`==# 文字样式==
**这是加粗的文字**
*这是倾斜的文字*
***这是斜体加粗的文字***
~~这是加删除线的文字~~`)

// 客户端信息
func clientInfo(ctx iris.Context) {
	ip := ctx.RemoteAddr()
	referrer := ctx.GetReferrer()
	ctx.Writef("IP:%s\r\nreferrer:%s", ip, referrer)
}

func fbnq(ctx iris.Context) {
	// num := ctx.Params().GetUint64Default("num", 0)
	num := ctx.URLParamDefault("num", "0")
	n, _ := strconv.Atoi(num)

	n1 := 1
	n2 := 0
	n3 := 0
	var sb strings.Builder
	for i := 0; i < n; i++ {
		n3 = n1 + n2
		n1, n2 = n2, n3
		sb.WriteString(strconv.Itoa(n3) + ", ")
	}
	// ctx.WriteString(sb.String())
	ctx.ViewData("fbnqStr", sb.String())
	ctx.View("fbnq.html")
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.HandleDir("/", "./static")

	app.Use(recover.New())
	app.Use(logger.New())

	app.RegisterView(iris.Django("./views", ".html"))

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown(markdownContent)
	})

	// 斐波那契 数列
	app.Get("/fbnq", fbnq)

	app.Get("/info", clientInfo)

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

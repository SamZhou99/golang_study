package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
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

// 斐波那契 计算
func fbnq(ctx iris.Context) {
	// num := ctx.Params().GetUint64Default("num", 0)
	num := ctx.URLParamDefault("num", "0")
	if num == "0" {
		ctx.HTML("<h1>请输入正确的参数 如:num=5</h1>")
		return
	}
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

func before(ctx iris.Context) {
	fmt.Println("before...............")
	ctx.Next()
}

func mainHandler(ctx iris.Context) {
	fmt.Println("main.................")
	ctx.Next()
}

func after(ctx iris.Context) {
	name := ctx.Params().Get("name")
	fmt.Println(name)
	ctx.WriteString("username: " + name)
}

// 这里开始运行
func main() {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel("debug")

	// 静态文件
	app.HandleDir("/", "./static")

	// 中间件
	app.Use(recover.New())

	// 记录日志
	app.Use(logger.New())

	// html模板
	app.RegisterView(iris.Django("./views", ".html"))

	// 路由传参
	app.Get("/username/{name}", before, mainHandler, after)

	// 设置参数
	app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
		i, e := ctx.Params().GetInt("id")
		if e != nil {
			ctx.WriteString("error you input")
		}
		ctx.WriteString(strconv.Itoa(i))
	})

	// 首页
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// 简单的ping-pong
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// json 返回
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// markdown 渲染
	app.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown(markdownContent)
	})

	// 异步的方式，不等待结果返回
	app.Get("long_async", func(ctx iris.Context) {
		// ctxCopy := ctx.Clone()
		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("Done! in path: %s , 这里可操作数据库", ctx.Path())
		}()
		ctx.WriteString("long_async done")
	})

	// 同步的方式，等待结果返回
	app.Get("/long_sync", func(ctx iris.Context) {
		time.Sleep(5 * time.Second)
		log.Printf("Done! in path: %s", ctx.Path())
		ctx.WriteString("long_sync done")
	})

	// 斐波那契 数列
	app.Get("/fbnq", fbnq)

	// 客户端信息
	app.Get("/info", clientInfo)

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

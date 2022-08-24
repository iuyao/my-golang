package main

import (
	"fmt"
	"my-golang/go-web/web-panic/gee"
	"net/http"
)

func test_recover() {
	//defer func() {
	//	fmt.Println("defer func")
	//	if err := recover(); err != nil {
	//		fmt.Println("recover success")
	//	}
	//}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}

func main() {
	engine := gee.NewEngine()
	//engine.Use(gee.Recovery())
	engine.Get("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>index Gee</h1>")
	})

	v1 := engine.NewGroup("/v1")
	{
		v1.Get("/login", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>v1 login Gee</h1>")
		})
		v1.Get("/panic", func(ctx *gee.Context) {
			names := []string{"geektutu"}
			ctx.String(http.StatusOK, names[100])
		})
	}
	v2 := engine.NewGroup("/v2")
	v2.Use()
	{
		v2.Get("/login", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>v2 login Gee</h1>")
		})
		v2.Get("/find", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>v2 find Gee</h1>")
		})
	}
	engine.Run(":9999")
}

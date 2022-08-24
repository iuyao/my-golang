package main

import (
	"my-golang/go-web/web-context/gee"
	"net/http"
)

func main() {
	g := gee.NewEngine()
	g.Get("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	g.Get("/getName", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	g.Post("/login", func(ctx *gee.Context) {
		ctx.ToJson(http.StatusOK, gee.H{
			"name":     ctx.PostForm("name"),
			"password": ctx.PostForm("password"),
		})
	})
	g.Run(":9999")
}

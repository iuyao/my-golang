package main

import (
	"my-golang/go-web/web-group/gee"
	"net/http"
)

func main() {

	engine := gee.NewEngine()
	engine.Get("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>index Gee</h1>")
	})
	group := engine.NewGroup("/v1")
	groupApi := group.NewGroup("/api")
	groupApi.Get("/login", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>v1 Gee</h1>")
	})
	groupAdmin := group.NewGroup("/admin")
	groupAdmin.Get("/find", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>v2 Gee</h1>")
	})
	engine.Run(":9999")
}

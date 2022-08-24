package main

import (
	"log"
	"my-golang/go-web/web-middleware/gee"
	"net/http"
	"time"
)

func onlyForV2() gee.HandleFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	engine := gee.NewEngine()
	engine.Use(gee.Logger())
	engine.Get("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>index Gee</h1>")
	})

	v1 := engine.NewGroup("/v1")
	{
		v1.Get("/login", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>v1 login Gee</h1>")
		})
		v1.Get("/find", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>v1 find Gee</h1>")
		})
	}
	v2 := engine.NewGroup("/v2")
	v2.Use(onlyForV2())
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

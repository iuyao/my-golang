package gee

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Router struct {
	handlers map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *Router) addRouter(method, patter string, handleFunc HandleFunc) {
	router := method + "-" + patter
	r.handlers[router] = handleFunc
}

func (r *Router) Handle(ctx *Context) {
	router := ctx.Method + "-" + ctx.Path
	if f, ok := r.handlers[router]; ok {
		// 将应该执行的func追加到上下文函数队列中
		ctx.handlers = append(ctx.handlers, f)
	} else {
		http.Error(ctx.Response, "not found handle", 500)
	}

	ctx.Next()
}

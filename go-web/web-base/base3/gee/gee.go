package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	Router map[string]HandleFunc
}

type HandleFunc func(w http.ResponseWriter, req *http.Request)

func (e *Engine) Get(patter string, handleFunc HandleFunc) {
	e.addRouter("GET", patter, handleFunc)
}
func (e *Engine) Post(patter string, handleFunc HandleFunc) {
	e.addRouter("POST", patter, handleFunc)
}

func (e *Engine) addRouter(method, patter string, handleFunc HandleFunc) {
	router := method + "-" + patter
	e.Router[router] = handleFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router := req.Method + "-" + req.URL.Path
	if f, ok := e.Router[router]; ok {
		f(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func NewEngine() *Engine {
	return &Engine{Router: make(map[string]HandleFunc)}
}

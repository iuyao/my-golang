package gee

import (
	"net/http"
)

type Engine struct {
	*Router
}

func (e *Engine) Get(patter string, handleFunc HandleFunc) {
	e.addRouter("GET", patter, handleFunc)
}
func (e *Engine) Post(patter string, handleFunc HandleFunc) {
	e.addRouter("POST", patter, handleFunc)
}

func (e *Engine) addRouter(method, patter string, handleFunc HandleFunc) {
	e.Router.addRouter(method, patter, handleFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(req, w)
	e.Router.Handle(context)
}
func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func NewEngine() *Engine {
	return &Engine{Router: NewRouter()}
}

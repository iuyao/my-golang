package gee

import (
	"log"
	"net/http"
)

type (
	Engine struct {
		*Router
		*RouterGroup
		groups []*RouterGroup
	}

	RouterGroup struct {
		prefix   string
		handlers map[string]HandleFunc
		parent   *RouterGroup // support nesting
		engine   *Engine      // all groups share a Engine instance
	}
)

func (group *RouterGroup) Get(patter string, handleFunc HandleFunc) {
	group.addRouter("GET", patter, handleFunc)
}
func (group *RouterGroup) Post(patter string, handleFunc HandleFunc) {
	group.addRouter("POST", patter, handleFunc)
}

func (group *RouterGroup) addRouter(method, patter string, handleFunc HandleFunc) {
	patter = group.prefix + patter
	log.Printf("Route %4s - %s", method, patter)
	group.engine.Router.addRouter(method, patter, handleFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(req, w)
	e.Router.Handle(context)
}
func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func NewEngine() *Engine {
	engine := &Engine{Router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine

}

func (group *RouterGroup) NewGroup(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

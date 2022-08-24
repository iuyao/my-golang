package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}
type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	Path       string
	Method     string
	StatusCode int
	handlers   []HandleFunc
	index      int
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	data := c.Request.PostFormValue(key)
	return data
}

func (c *Context) Status(status int) int {
	c.StatusCode = status
	c.Response.WriteHeader(status)
	return c.StatusCode
}

func (c *Context) SetHeader(key, value string) {
	c.Request.Header.Set(key, value)
}

func (c *Context) HTML(code int, data string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Response.Write([]byte(data))
}

func (c *Context) ToJson(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	if err := json.NewEncoder(c.Response).Encode(data); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data string) {
	c.Status(code)
	c.Response.Write([]byte(data))
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	c := &Context{
		Path:     r.URL.Path,
		Method:   r.Method,
		Response: w,
		Request:  r,
		index:    -1,
	}
	return c
}

func (c *Context) Next() {
	c.index++
	length := len(c.handlers)
	for ; c.index < length; c.index++ {
		c.handlers[c.index](c)
	}
}
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.ToJson(code, H{"message": err})
}

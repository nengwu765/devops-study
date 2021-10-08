package frame

import (
	"github.com/golang/glog"
	"net/http"
)

// HandleFunc 自定义一个HandleFunc，替代原生的 http.HandlerFunc
type HandleFunc func(c *Context)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandleFunc) {
	glog.V(LogInfo).Infof("Route %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.FormatString(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}

func (r *router) GET(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodGet, pattern, handler)
}

func (r *router) POST(pattern string, handler HandleFunc) {
	r.addRoute(http.MethodPost, pattern, handler)
}
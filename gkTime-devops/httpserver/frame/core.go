package frame

import (
	"net/http"
)

// 检测Core继承了 http.Handler
var _ http.Handler = &Core{}

type Core struct {
	Router *router
}

func NewCore() *Core {
	return &Core{newRouter()}
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	context := newContext(writer, request)
	c.Router.handle(context)
}

func (c *Core) Run(addr string) error {
	return http.ListenAndServe(addr, c)
}

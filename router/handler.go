package router

import (
	"github.com/next-frmework/apollo/config"
	"net/http"
)

type Handler interface {
	Handle(c *Context)
}

type HandlerRouterMapping interface {
	Add(router config.Router, handler Handler)
	Resolve(r http.Request) (Handler, *Context, error)
}

type DefaultHandlerRouterMapping struct {
	// todo 构建一个用于path查找的语法树
	Handlers map[string]Handler
	Routers  map[string]*config.Router
}

func (hrm *DefaultHandlerRouterMapping) Add(router config.Router, handler Handler) {

}

func (hrm *DefaultHandlerRouterMapping) Resolve(r http.Request) (Handler, *Context, error) {
	return nil, nil, nil
}

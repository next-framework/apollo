package router

import (
	"github.com/next-frmework/apollo/config"
	"net/http"
	"strings"
)

const (
	GET     string = "get"
	HEAD    string = "head"
	POST    string = "post"
	PUT     string = "put"
	PATCH   string = "patch"
	DELETE  string = "delete"
	OPTIONS string = "options"
	TRACE   string = "trace"
)

type Handler interface {
	Handle(c *Context)
}

type HandlerRouterMapping interface {
	Add(router *config.Router, handler Handler)
	Resolve(r *http.Request) (Handler, *Context, error)
}

type DefaultHandlerRouterMapping struct {
	// todo 构建一个用于url匹配的语法树
	Handlers map[string]Handler
	Routers  map[string]*config.Router
}

func (hrm *DefaultHandlerRouterMapping) Add(router *config.Router, handler Handler) {
	methods := router.Methods
	if len(methods) == 0 {
		methods = getAllMethods()
	}

	if hrm.Routers == nil {
		hrm.Routers = make(map[string]*config.Router)
	}
	for _, v := range methods {
		key := buildRouterKey(router.Handler, strings.ToLower(v))
		hrm.Routers[key] = router

		// todo 构建一个用于url匹配的语法树
	}

	if hrm.Handlers == nil {
		hrm.Handlers = make(map[string]Handler)
	}
	hrm.Handlers[router.Handler] = handler
}

func (hrm *DefaultHandlerRouterMapping) Resolve(r *http.Request) (Handler, *Context, error) {
	return nil, nil, nil
}

func getAllMethods() []string {
	methods := []string{GET, HEAD, POST, PUT, PATCH, DELETE, OPTIONS, TRACE}
	return methods
}

func buildRouterKey(handlerName, method string) string {
	key := "[" + handlerName + "]" + ":" + "[" + method + "]"
	return key
}

package router

import "github.com/next-frmework/apollo/config"

type Handler interface {
	Handle(c *Context)
}

type HandlerMapping interface {
	Add(router config.Router, handler Handler)
}

type HandlerMappingAdapter struct {
}

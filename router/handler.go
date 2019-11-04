package router

import "github.com/next-frmework/apollo/config"

type Handler interface {
	Handle(c *Context)
}

type Entry struct {
	Key     string
	Router  *config.Router
	Handler *Handler
}

package router

type Handler interface {
	Handle(c *Context)
}

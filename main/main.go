package main

import (
	"fmt"
	apollo2 "github.com/next-frmework/apollo"
)

type HelloHandler struct {
}

func (h *HelloHandler) Handle(c *apollo2.Context) {
	r := c.Request
	fmt.Println("hello handler")
	fmt.Println(r.URL.Path)
}

type RootHandler struct {
}

func (h *RootHandler) Handle(c *apollo2.Context) {
	r := c.Request
	fmt.Println("root handler")
	fmt.Println(r.URL.Path)
}

func main() {
	a := apollo2.NewApollo()
	h := &HelloHandler{}
	h1 := &RootHandler{}
	a.RegisterHandler("x.HelloMux", h)
	a.RegisterHandler("x.Mux", h1)
	a.Run()
}

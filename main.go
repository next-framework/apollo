package main

import (
	"fmt"
	"github.com/next-frmework/apollo/apollo"
)

type HelloHandler struct {
}

func (h *HelloHandler) Handle(c *apollo.Context) {
	r := c.Request
	fmt.Println(r.URL.Path)
}

type RootHandler struct {
}

func (h *RootHandler) Handle(c *apollo.Context) {
	r := c.Request
	fmt.Println(r.URL.Path)
}

func main() {
	a := apollo.NewApollo()
	h := &HelloHandler{}
	h1 := &RootHandler{}
	a.RegisterHandler("x.HelloMux", h)
	a.RegisterHandler("x.Mux", h1)
	a.Run()
}

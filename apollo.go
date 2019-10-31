package main

import (
	"fmt"
	"github.com/next-frmework/apollo/config"
	"net/http"
)

type Apollo struct {
	ConfigPath string
	App        *config.Application
}

func NewApollo() *Apollo {
	return new(Apollo)
}

func (a *Apollo) Run() {
	http.ListenAndServe("127.0.0.1:8080", a)
}

func (a *Apollo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

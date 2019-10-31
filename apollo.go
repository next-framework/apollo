package main

import (
	"fmt"
	"github.com/next-frmework/apollo/config"
	"net/http"
	"os"
)

type Apollo struct {
	ConfigPath string
	App        *config.Application
}

func NewApollo() *Apollo {
	return new(Apollo)
}

func (a *Apollo) Run() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(path)
	http.ListenAndServe("127.0.0.1:8080", a)
}

func (a *Apollo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

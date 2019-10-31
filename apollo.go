package main

import (
	"fmt"
	"github.com/next-frmework/apollo/config"
	"github.com/next-frmework/apollo/utils"
	"net/http"
	"os"
)

type Apollo struct {
	Filename string
	App      *config.Application
}

func NewApollo() *Apollo {
	a := &Apollo{}
	return a
}

func NewApolloWithFilename(filename string) *Apollo {
	a := &Apollo{Filename: filename}
	return a
}

func (a *Apollo) Run() {
	filename := a.Filename
	if len(filename) == 0 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf(path)
		// todo find apollo-application.* config file
	}

	suffix := utils.GetSuffix(a.Filename)
	var p config.Parser
	switch {
	case suffix == ".yml" || suffix == ".yaml":
		p = &config.Yaml{}
	case suffix == ".toml":
		p = &config.Toml{}
	default:
		return
	}

	app, err := p.Parse(a.Filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	a.App = app

	http.ListenAndServe("127.0.0.1:8080", a)
}

func (a *Apollo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

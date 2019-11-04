package apollo

import (
	"fmt"
	"github.com/next-frmework/apollo/config"
	"github.com/next-frmework/apollo/router"
	"github.com/next-frmework/apollo/utils"
	"net/http"
	"os"
	"path/filepath"
)

type Apollo struct {
	Filename           string
	ApplicationConfig  *config.Application
	RegisteredHandlers map[string]router.Handler
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
	if len(a.RegisteredHandlers) == 0 {
		// todo 打印日志
		return
	}

	filename := a.Filename
	if len(filename) == 0 {
		path, err := os.Getwd()
		if err != nil {
			// todo 打印日志
			return
		}
		filename, err = utils.Find(path, "apollo-application.*", a.match)
		if len(filename) == 0 || err != nil {
			// todo 打印日志
			return
		}

		a.Filename = filename
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
		// todo 打印日志
		return
	}
	a.ApplicationConfig = app

	if len(a.ApplicationConfig.Routers) == 0 {
		// todo 打印日志
		return
	}

	for _, v := range a.ApplicationConfig.Routers {
		handler, existed := a.RegisteredHandlers[v.Handler]
		if !existed {
			// todo 打印日志
			return
		}

		path := v.Path
	}

	http.ListenAndServe("127.0.0.1:8080", a)
}

func (a *Apollo) RegisterHandler(name string, handler *router.Handler) {
	if a.RegisteredHandlers == nil {
		a.RegisteredHandlers = make(map[string]router.Handler)
	}

	a.RegisteredHandlers[name] = handler
}

func (a *Apollo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func (a *Apollo) match(pattern, filename string) bool {
	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		return false
	}
	return matched
}

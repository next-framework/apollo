package apollo

import (
	"github.com/next-frmework/apollo/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Apollo struct {
	Filename             string
	ApplicationConfig    *Application
	RegisteredHandlers   map[string]Handler
	HandlerRouterMapping HandlerRouterMapping
}

func NewApollo() *Apollo {
	a := &Apollo{RegisteredHandlers: make(map[string]Handler), HandlerRouterMapping: new(DefaultHandlerRouterMapping)}
	return a
}

func NewApolloWithFilename(filename string) *Apollo {
	a := &Apollo{Filename: filename, RegisteredHandlers: make(map[string]Handler), HandlerRouterMapping: new(DefaultHandlerRouterMapping)}
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
	var p Parser
	switch {
	case suffix == ".yml" || suffix == ".yaml":
		p = &Yaml{}
	case suffix == ".toml":
		p = &Toml{}
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

	if a.HandlerRouterMapping == nil {
		a.HandlerRouterMapping = new(DefaultHandlerRouterMapping)
	}

	for _, v := range a.ApplicationConfig.Routers {
		handler, existed := a.RegisteredHandlers[v.Handler]
		if !existed {
			// todo 打印日志
			return
		}

		a.HandlerRouterMapping.Add(&v, handler)
	}

	server := a.ApplicationConfig.Server
	ipAndPort := server.Address + ":" + strconv.Itoa(server.Port)

	http.ListenAndServe(ipAndPort, a)
}

func (a *Apollo) RegisterHandler(name string, handler Handler) {
	if a.RegisteredHandlers == nil {
		a.RegisteredHandlers = make(map[string]Handler)
	}

	a.RegisteredHandlers[name] = handler
}

func (a *Apollo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := a.HandlerRouterMapping.Resolve(r)
	if h == nil {
		// todo 添加日志，并且根据不同的错误产生不同的响应信息
		return
	}

	c := &Context{Response: w, Request: r, Apollo: a}
	h.Handle(c)
}

func (a *Apollo) match(pattern, filename string) bool {
	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		return false
	}
	return matched
}

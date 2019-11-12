package apollo

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Parser interface {
	Parse(filename string) (*Application, error)
}

type Yaml struct {
	Name   string
	Server struct {
		Address string
		Port    int
	}
	Routers []struct {
		Name    string
		Path    string
		Handler string
		Methods []string
	}
}

func (y *Yaml) Parse(filename string) (*Application, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, y)
	if err != nil {
		return nil, err
	}

	app := &Application{}
	app.Name = y.Name

	app.Server = Server{Address: y.Server.Address, Port: y.Server.Port}

	if len(y.Routers) != 0 {
		app.Routers = make([]Router, 0, len(y.Routers))
		for _, v := range y.Routers {
			r := Router{Name: v.Name, Path: v.Path, Handler: v.Handler}
			if len(v.Methods) != 0 {
				r.Methods = make([]string, 0, len(v.Methods))
				for _, v1 := range v.Methods {
					r.Methods = append(r.Methods, v1)
				}
			}
			app.Routers = append(app.Routers, r)
		}
	}

	return app, nil
}

type Toml struct {
}

func (t *Toml) Parse(filename string) (*Application, error) {
	return nil, nil
}

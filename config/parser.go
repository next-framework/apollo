package config

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
	RequestMappings []struct {
		Name    string
		Path    string
		Handler string
		Methods []string
	} `yaml:"request-mappings"`
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

	if len(y.RequestMappings) != 0 {
		app.RequestMappings = make([]RequestMapping, 0, len(y.RequestMappings))
		for _, v := range y.RequestMappings {
			r := RequestMapping{Name: v.Name, Path: v.Path, Handler: v.Handler}
			if len(v.Methods) != 0 {
				r.Methods = make([]string, 0, len(v.Methods))
				for _, v1 := range v.Methods {
					r.Methods = append(r.Methods, v1)
				}
			}
			app.RequestMappings = append(app.RequestMappings, r)
		}
	}

	return app, nil
}

type Toml struct {
}

func (t *Toml) Parse(filename string) (*Application, error) {
	return nil, nil
}

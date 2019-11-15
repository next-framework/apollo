package apollo

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Name    string   `yaml:"name" toml:"name"`
	Server  Server   `yaml:"server" toml:"server"`
	Routers []Router `yaml:"routers" toml:"routers"`
}

type Server struct {
	HostName string `yaml:"host-name" tome:"host-name"`
	Port     uint16 `yaml:"port" toml:"port"`
}

type Router struct {
	Name       string   `yaml:"name" toml:"name"`
	Path       string   `yaml:"path" toml:"path"`
	Controller string   `yaml:"controller" toml:"controller"`
	Methods    []string `yaml:"methods" toml:"methods"`
}

type ReadOnlyConfiguration interface {
	GetName() string
	GetServer() Server
	GetRouters() []Router
}

func (c Configuration) GetName() string {
	return c.Name
}

func (c Configuration) GetServer() Server {
	return c.Server
}

func (c Configuration) GetRouters() []Router {
	return nil
}

type ReadOnlyServer interface {
	GetHostName() string
	GetPort() uint16
}

func (s Server) GetHostName() string {
	return s.HostName
}

func (s Server) GetPort() uint16 {
	return s.Port
}

type ReadOnlyRouter interface {
	GetName() string
	GetPath() string
	GetController() string
	GetMethods() []string
}

func (r Router) GetName() string {
	return r.Name
}

func (r Router) GetPath() string {
	return r.Path
}

func (r Router) GetController() string {
	return r.Controller
}

func (r Router) GetMethods() []string {
	return r.Methods
}

type ConfigurationReader interface {
	Read(resource string) (*Configuration, error)
}

type ConfigurationReaderFactory interface {
	Create() ConfigurationReader
}

type YamlConfigurationReader struct{}

func (r *YamlConfigurationReader) Read(resource string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(resource)
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type YamlConfigurationReaderFactory struct{}

func (f *YamlConfigurationReaderFactory) Create() ConfigurationReader {
	return &YamlConfigurationReader{}
}

type TomlConfigurationReader struct{}

func (r *TomlConfigurationReader) Read(resource string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(resource)
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	err = toml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type TomlConfigurationReaderFactory struct{}

func (f *TomlConfigurationReaderFactory) Create() ConfigurationReader {
	return &TomlConfigurationReader{}
}

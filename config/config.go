package config

var File = "apollo-application.yml"

type Config struct {
	Name            string           `yaml:"name" toml:"name"`
	Server          Server           `yaml:"server" toml:"server"`
	HandlerMappings []HandlerMapping `yaml:"handler-mappings" toml:"handler-mappings"`
}

type Server struct {
	Host string `yaml:"host" tome:"host"`
	Port uint16 `yaml:"port" toml:"port"`
}

type HandlerMapping struct {
	Name    string   `yaml:"name" toml:"name"`
	Path    string   `yaml:"path" toml:"path"`
	Handler string   `yaml:"handler" toml:"handler"`
	Methods []string `yaml:"methods" toml:"methods"`
}

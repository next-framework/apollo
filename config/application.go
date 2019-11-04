package config

type Application struct {
	Name            string
	Server          Server
	RequestMappings []RequestMapping
}

type Server struct {
	Address string
	Port    int
}

type RequestMapping struct {
	Name    string
	Path    string
	Handler string
	Methods []string
}

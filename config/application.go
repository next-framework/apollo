package config

type Application struct {
	Name            string
	Servers         []Server
	RequestMappings []RequestMapping
}

type Server struct {
	Name    string
	Address string
	Port    int
}

type RequestMapping struct {
	Name    string
	Path    string
	Handler string
	Methods []string
}

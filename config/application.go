package config

type Application struct {
	Name    string
	Server  Server
	Routers []Router
}

type Server struct {
	Address string
	Port    int
}

type Router struct {
	Name    string
	Path    string
	Handler string
	Methods []string
}

package config

type RequestMapping struct {
	Name    string
	Path    string
	Handler string
	Methods []string
}

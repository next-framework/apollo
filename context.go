package apollo

type emptyStruct struct{}

var (
	void        emptyStruct
	usedConfigs = make(map[string]emptyStruct)
)

type (
	Context interface {
		ConfigPath() string

		Configuration() ReadOnlyConfiguration

		Attributes() Storage
	}

	context struct {
		configPath    string
		configuration Configuration
		attributes    Storage
	}
)

func NewContext() Context {
	return nil
}

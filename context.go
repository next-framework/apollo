package apollo

import (
	"errors"
	"os"
	"path/filepath"
)

type emptyStruct struct{}

var (
	void        emptyStruct
	usedConfigs = make(map[string]emptyStruct)
)

type (
	Context interface {
		Filename() string

		Configuration() Configuration

		Attributes() Storage
	}

	context struct {
		filename      string
		configuration Configuration
		attributes    Attributes
	}

	Attributes struct {
		Storage
	}
)

func NewContext() (Context, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filename, err := FindFile(workDir, "apollo-application*.*", fileMatch)
	if err != nil {
		return nil, err
	}
	suffix := GetFileSuffix(filename)

	var reader ConfigurationReader
	var readerFactory ConfigurationReaderFactory
	switch {
	case suffix == ".yml" || suffix == ".yaml":
		readerFactory = &YamlConfigurationReaderFactory{}
		reader = readerFactory.Create()
	case suffix == ".tml" || suffix == ".toml":
		readerFactory = &TomlConfigurationReaderFactory{}
		reader = readerFactory.Create()
	default:
		return nil, errors.New("unsupported config file format")
	}

	config, err := reader.Read(filename)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewContextWithConfig(filename string) (Context, error) {
	return nil, nil
}

func (ctx *context) Filename() string {
	return ctx.filename
}

func (ctx *context) Configuration() Configuration {
	return ctx.configuration
}

func fileMatch(pattern, filename string) bool {
	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		return false
	}

	if _, existed := usedConfigs[filename]; !existed && matched {
		return true
	}

	return false
}

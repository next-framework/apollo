package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/next-frmework/apollo/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
)

type Factory struct {
}

func (f *Factory) Create() (*Config, error) {
	if File == "" {
		return nil, fmt.Errorf("empty config file")
	}

	matched, err := regexp.MatchString("apollo-application*", File)
	if err != nil {
		return nil, fmt.Errorf("regexp match error%w", err)
	}
	if !matched {
		return nil, fmt.Errorf("invalid file: %v", File)
	}

	suffix := util.GetSuffix(File)
	var parser Parser
	switch suffix {
	case ".yml":
		fallthrough
	case ".yaml":
		parser = new(YamlParser)
	case ".toml":
		parser = new(TomlParser)
	default:
		return nil, fmt.Errorf("unsupported file: %v", File)
	}

	buffer, err := ioutil.ReadFile(File)
	if err != nil {
		return nil, fmt.Errorf("read file error%w", err)
	}

	c, err := parser.Parse(buffer)
	if err != nil {
		return nil, fmt.Errorf("%T parse error%w", parser, err)
	}

	return c, nil
}

type Parser interface {
	Parse(buffer []byte) (*Config, error)
}

type YamlParser struct {
}

func (yp *YamlParser) Parse(buffer []byte) (*Config, error) {
	c := new(Config)

	err := yaml.Unmarshal(buffer, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type TomlParser struct {
}

func (tp *TomlParser) Parse(buffer []byte) (*Config, error) {
	c := new(Config)

	err := toml.Unmarshal(buffer, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

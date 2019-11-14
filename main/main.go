package main

import (
	"fmt"
	"github.com/next-frmework/apollo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Name  string `yaml:"name"`
	Hello Hello  `yaml:"hello"`
}

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) GetHello() ReadOnlyHello {
	return c.Hello
}

type Hello struct {
	Name string `yaml:"name"`
}

func (h *Hello) GetName() string {
	return h.Name
}

type ReadOnlyConfig interface {
	GetName() string
	GetHello() ReadOnlyHello
}

type ReadOnlyHello interface {
	GetName() string
}

type TestMap map[string]string

func (t *TestMap) test() {
	*t = TestMap{}
}

func main() {
	src, err := ioutil.ReadFile("main/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c := &Config{}
	err = yaml.Unmarshal(src, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := c.GetHello()
	fmt.Printf("%+v\n", r)

	fmt.Printf("%+v", c)

	fmt.Println(filepath.Clean("../"))

	m := apollo.Storage{}
	m.Put("a", 1)

	t := TestMap{}
	t["a"] = "1"

	t.test()

	m1 := apollo.Storage{}
	m1.Put("a", 1)

	var m2 map[string]string = map[string]string{}
	m2["a"] = "1"
}

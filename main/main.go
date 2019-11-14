package main

import (
	"fmt"
	"github.com/next-frmework/apollo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Name string `yaml:"name"`
}

type TestMap map[string]string

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

	fmt.Printf("%+v", c)

	fmt.Println(filepath.Clean("../"))

	m := apollo.Storage{}
	m.Put("a", 1)

	t := TestMap{}
	t["a"] = "1"

	m1 := apollo.Storage{}
	m1.Put("a", 1)

	var m2 map[string]string = map[string]string{}
	m2["a"] = "1"
}

package main

import (
	"fmt"
	"github.com/next-frmework/apollo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	User  string `yaml:"user"`
	Hello Hello  `yaml:"hello"`
	Attrs apollo.Storage
}

func (c Config) GetUser() string {
	return c.User
}

func (c Config) GetHello() Hello {
	fmt.Printf("%p\n", &c.Hello)
	return c.Hello
}

type Hello struct {
	To string `yaml:"to"`
}

func (h Hello) GetTo() string {
	return h.To
}

type ReadOnlyConfig interface {
	GetUser() string
	GetHello() Hello
}

type ReadOnlyHello interface {
	GetTo() string
}

type TestMap map[string]string

func (t *TestMap) test() {
	*t = TestMap{}
}

func main() {
	f, err1 := filepath.Abs("config.yaml")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	fmt.Println(f)

	fmt.Println(filepath.Match("app*.*", "app.c"))

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

	fmt.Printf("%p\n", &c.Hello)

	oo := c.Hello
	fmt.Printf("%p\n", &oo)
	oo.To = "A"

	fmt.Println(c.Hello.To)

	ooo := &c.Hello
	fmt.Printf("%p\n", ooo)
	ooo.To = "B"
	fmt.Println(c.Hello.To)

	r := c.GetHello()
	fmt.Printf("%p\n", &r)

	fmt.Printf("%+v\n", r)

	fmt.Printf("%#v\n", c)

	fmt.Printf("c.Attrs=====%p\n", &c.Attrs)

	a := c.Attrs
	if a == nil {
		c.Attrs = apollo.Storage{}
		a = c.Attrs
		fmt.Printf("c.Attrs=====%p\n", &c.Attrs)
	}
	a.Put("a", 1)

	fmt.Println(a["a"].IntDefault(-1))
	fmt.Println(c.Attrs["a"].IntDefault(-1))

	fmt.Printf("&a=====%p\n", &(a))

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

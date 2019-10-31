package main

import (
	"fmt"
	"github.com/next-frmework/apollo/config"
)

func main() {
	y := config.Yaml{}
	app, err := y.Parse("apollo-application.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", *app)
}

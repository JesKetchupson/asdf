package main

import (
	config "github.com/JesKetchupson/asdf/configs"
	"github.com/JesKetchupson/asdf/internal/asdf"
)

func main() {
	conf := config.ParceConfig()
	if err := asdf.Run(conf); err != nil {
		panic(err)
	}
}

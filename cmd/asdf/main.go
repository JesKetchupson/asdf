package main

import "github.com/JesKetchupson/asdf/internal/asdf"

func main() {
	if err := asdf.Run(); err != nil {
		panic(err)
	}
}

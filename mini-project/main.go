package main

import (
	"mini-project/config"
	"mini-project/route"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}

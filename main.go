package main

import (
	"github.com/binod210/go-inventory-management/app"
)

const (
	server = ":9090"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(server)

}

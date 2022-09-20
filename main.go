package main

import (
	"shrinkly/model"
	"shrinkly/server"
)

func main() {
	model.Setup()
	server.Setup()
}

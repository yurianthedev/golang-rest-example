package main

import (
	"github.com/yurianxdev/rest-example/app"
	"github.com/yurianxdev/rest-example/config"
)

func main() {
	config.InitConfiguration()
	app.StartApp()
}

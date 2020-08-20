package main

import (
	"github.com/yurianxdev/rest-example/app"
	"github.com/yurianxdev/rest-example/config"
	"github.com/yurianxdev/rest-example/database"
)

func main() {
	config.InitConfiguration()
	database.InitDatabase()
	app.StartApp()
}

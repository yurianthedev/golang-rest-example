package database

import (
	"log"

	"github.com/yurianxdev/rest-example/config"
)

func InitDatabase() {
	switch config.Database.Dialect {
	case "postgres":
		initPostgres()
	default:
		log.Fatalf("Selected dialect is not implemented\n")
	}
}

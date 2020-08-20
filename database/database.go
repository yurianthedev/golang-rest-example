package database

import (
	"database/sql"
	"log"

	"github.com/yurianxdev/rest-example/config"
)

var DB *sql.DB

func InitDatabase() {
	switch config.Database.Dialect {
	case "postgres":
		initPostgres()
	default:
		log.Fatalf("Selected dialect is not implemented\n")
	}
}

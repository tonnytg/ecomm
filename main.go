package main

import (
	"ecomm/api"
	"ecomm/database"
	"os"
)

func main() {
	// start database connection
	//database.InitDatabase()

	if len(os.Args) > 1 {
		// fist running to build tables
		// you need call with 1 argument for example: go run main.go create-db
		database.Migrate()
	} else {
		// start server to listen
		api.StartApi()
	}
}

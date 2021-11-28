package main

import (
	"ecomm/api"
	"ecomm/database"
	"os"
)

func main(){
	// start database connection
	database.InitDatabase()

	// start api handler
	if len(os.Args) > 1 {
		database.Migrate()
	} else {
		api.StartApi()
	}
}

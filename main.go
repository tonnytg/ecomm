package main

import (
	"ecomm/api"
	"ecomm/database"
)

func main(){
	// start database connection
	database.InitDatabase()

	// start api handler
	api.StartApi()
}

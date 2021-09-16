package database

import (
	log "ecomm/pkg/loggin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=ecomm port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Msg("CRITICAL", "Unable to open database")
		os.Exit(1)
	}
	DB = database
}

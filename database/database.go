package database

import (
	log "ecomm/pkg/loggin"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

const (
	host     = "127.0.0.1"
	port     = "5432"
	dbname   = "ecomm"
	user     = "postgres"
	password = "postgres"
	sslMode  = "disable"
)

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", host, port, dbname, user, password, sslMode)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Unable to open database %s:%s - dbname: %s", host, port, dbname)
		log.Msg("CRITICAL", msg)
		os.Exit(1)
	}
	DB = database
}

package database

import (
	log "ecomm/pkg/loggin"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type DBConf struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
	sslMode  string
}

var db DBConf

func init() {
	db.host = os.Getenv("DB_HOST")
	db.port = os.Getenv("DB_PORT")
	db.dbname = os.Getenv("DB_DBNAME")
	db.user = os.Getenv("DB_USER")
	db.password = os.Getenv("DB_PASSWORD")
	db.sslMode = os.Getenv("DB_SSL")
}

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		db.host, db.port, db.dbname, db.user, db.password, db.sslMode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to database %s:%s/%s", db.host, db.port, db.dbname)
		log.Msg("CRITICAL", msg)
		os.Exit(1)
	}
	DB = database
}

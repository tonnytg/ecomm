package database

import (
	client "ecomm/entity"
	log "ecomm/pkg/loggin"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

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

func Migrate() {
	Client := &client.Client{}
	DB.AutoMigrate(&Client)

	createClient()
	fmt.Println("Migration complete")
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func createClient() {
	genId := uuid.New()
	clients := &[2]client.Client{
		{
			Address    : "test",
			Email      : "teste@gmail.com",
			ValidEmail : true,
			FirstName  : "tonny",
			LastName   : "tg",
			Password   : "tonnytg",
			Username   : "tonnytg",
			Status     : "ok",
			ID         : genId,
		},
	}
	fmt.Println(clients)

	for i := 0; i < len(clients); i++ {
		generatePassword := HashAndSalt([]byte(clients[i].Username))
		user := &client.Client{Username: clients[i].Username, Email: clients[i].Email, Password: generatePassword}
		DB.Create(&user)
	}
}

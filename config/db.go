package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
)

// if err := godotenv.Load(); err != nil {
// 	log.Fatalf("Error loading .env file")
// }

var (
	host     = EnvFile["DB_HOST"]
	port     = EnvFile["DB_PORT"]
	user     = EnvFile["DB_USER"]
	password = EnvFile["DB_PASSWORD"]
	dbName   = EnvFile["DBNAME"]
)

func ConnectDB() *gorm.DB {
	dbConfig := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	fmt.Println(dbConfig, ";;;;;::::;;:::")
	helper.ErrorPanic(err)

	fmt.Println("DB Connection Successful")

	return db
}

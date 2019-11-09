package utils

import (
	"golang-template-api-authentication/modules/User/Shared/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
)

type DBController struct {
    DB *gorm.DB
}

var db = &DBController{}

//ConnectDB function: Make database connection
func InitDB() *gorm.DB {

	//Load environmenatal variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")
	databasePort := os.Getenv("databasePort")

	//Define DB connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, username, databaseName, password)

	//connect to db URI
	db.DB, err = gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	//defer db.DB.Close()

	// Migrate the schema
	db.DB.AutoMigrate(
		&models.User{})

	fmt.Println("Successfully connected!", db.DB)
	return db.DB
}

func GetDB() *gorm.DB {
	return db.DB
}
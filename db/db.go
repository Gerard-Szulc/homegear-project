package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitiateDB() {
	var dbhost, exists = os.LookupEnv("DBHOST")
	if !exists {
		dbhost = "localhost"
		fmt.Println(exists)
	}
	dbport, exists := os.LookupEnv("DBPORT")
	if !exists {
		fmt.Println(exists)
	}
	user, exists := os.LookupEnv("DBUSER")
	if !exists {
		fmt.Println(exists)
	}
	password, exists := os.LookupEnv("DBPASSWORD")
	if !exists {
		fmt.Println(exists)
	}
	dbname, exists := os.LookupEnv("DBNAME")
	if !exists {
		fmt.Println(exists)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbhost, dbport, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//db.DB().SetMaxIdleConns(20)
	//db.DB().SetMaxOpenConns(200)
	DB = db
}

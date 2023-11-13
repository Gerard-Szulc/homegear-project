package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitiateDB() {

	// sqlite for testing purposes
	sqliteFilePath, exists := os.LookupEnv("SQLITE_DB_DUST_PATH")
	if !exists {
		fmt.Println("sqlite env doesnt exist")
	}

	//dbport, exists := os.LookupEnv("DBPORT")
	//if !exists {
	//	fmt.Println(exists)
	//}
	//user, exists := os.LookupEnv("DBUSER")
	//if !exists {
	//	fmt.Println(exists)
	//}
	//password, exists := os.LookupEnv("DBPASSWORD")
	//if !exists {
	//	fmt.Println(exists)
	//}
	//dbname, exists := os.LookupEnv("DBNAME")
	//if !exists {
	//	fmt.Println(exists)
	//}
	db, err := gorm.Open(sqlite.Open(sqliteFilePath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//db.DB().SetMaxIdleConns(20)
	//db.DB().SetMaxOpenConns(200)
	DB = db

}

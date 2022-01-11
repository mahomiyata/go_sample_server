package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sample/server/entity"
)

var counter = 0
var dsn string = "host=todo_line_db user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=todo_line port=5432"

func Init() {
	db := GetDB()

	// Migrate the schema
	db.AutoMigrate(&entity.Note{})

	// Seeding
	// db.Create(&entity.Note{UserID: "123123", Content: "Note1!"})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()

}

func GetDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil && counter == 5 {
		fmt.Println("failed to connect database")
		panic(err)
	} else if err != nil {
		counter++
		GetDB()
	}
	return db
}

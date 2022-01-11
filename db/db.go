package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sample/server/entity"
)

var dbURL = os.Getenv("DATABASE_URL")
var dsn string = "host=localhost user=" + os.Getenv("USER") + " dbname=todo_line port=5432"

func Init() {
	db := GetDB()

	// Migrate the schema
	db.AutoMigrate(&entity.Note{})

	// Seeding
	// db.Create(&entity.Note{UserID: "123123", Content: "Note1!"})

}

func GetDB() *gorm.DB {
	if dbURL == "" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("failed to connect database")
			panic(err)
		}
		return db
	} else {
		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			fmt.Println("failed to connect database")
			panic(err)
		}
		return db
	}
}

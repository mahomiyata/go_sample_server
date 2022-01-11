package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"sample/server/db"
	"sample/server/entity"
)

type Note entity.Note

func main() {

	db.Init()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Get 5 notes
	r.GET("/notes", func(c *gin.Context) {
		db := db.GetDB()

		var notes []Note
		result := db.Limit(5).Find(&notes)

		if result.Error != nil {
			fmt.Println("Something Wrong...")
			log.Fatal(result.Error)
		}

		c.JSON(200, notes)
	})

	// Post new note
	r.POST("/notes", func(c *gin.Context) {
		db := db.GetDB()

		var note Note
		c.BindJSON(&note)
		db.Create(&entity.Note{UserID: note.UserID, Content: note.Content})

		fmt.Println(note.Content)
		c.String(http.StatusCreated, "Created!")
	})

	// Set up port number adn run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
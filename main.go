package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"sample/server/db"
	"sample/server/entity"

	"github.com/gin-gonic/gin"
)

type Note entity.Note

func main() {

	db.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This server is for Line Bot🦭",
		})
	})

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

	// Get the user's notes
	r.GET("/notes/:id", func(c *gin.Context) {
		id := c.Param("id")
		db := db.GetDB()

		var notes []Note
		result := db.Where("user_id = ?", id).Find(&notes)

		if result.Error != nil {
			fmt.Println("Something Wrong...")
			log.Fatal(result.Error)
		}

		c.JSON(http.StatusOK, notes)
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

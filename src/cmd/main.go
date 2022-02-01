package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/pkg"
)

func main() {
	db := pkg.ConnectToDatabase()
	createError := pkg.CreateNumberTable(db)
	if createError != nil {
		log.Fatalf("Cannot create Number table. %s Closing app.", createError.Error())
	}

	router := gin.Default()

	router.POST("/numbers/:number", func(c *gin.Context) { pkg.AddNumber(c, db) })

	router.Run()
}

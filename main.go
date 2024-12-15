// main.go
package main

import (
	"github.com/dimasmir03/rest-api-gin-gorm-logrocket/controllers"
	"github.com/dimasmir03/rest-api-gin-gorm-logrocket/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run()
	if err != nil {
		return
	}
}

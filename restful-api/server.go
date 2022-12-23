package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/siddhantprateek/ohmyback-end/restful-api/handler"
)

func main() {

	route := gin.Default()
	route.GET("/books", handler.GetBooks)
	route.GET("/books/:id", handler.BookById)
	route.POST("/books", handler.CreateBook)
	route.PATCH("/checkout", handler.CheckOutBook)
	route.PATCH("/return", handler.ReturnBook)

	route.Run("localhost:4000")
}

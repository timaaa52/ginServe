package main

import (
	"log"
	"simplehttpserve/database"
	"simplehttpserve/services"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	database.DB, err = database.Connection()
	if err != nil { 
		log.Fatalf("Bad connection to database: %v", err.Error())
	}

	defer database.DB.Close()
	
	route := gin.Default()
	route.GET("/", services.GetTodos)
	route.POST("/todos", services.CreateTodo)
	route.DELETE("/todos/:id", services.DeleteTodo)
	route.PUT("/todos/:id", services.UpdateTodos)
	route.Run("localhost:8080")
}

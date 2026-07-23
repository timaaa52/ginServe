package main

import (
	"simplehttpserve/services"
	"github.com/gin-gonic/gin"
)

func main() {
	
	
	route := gin.Default()
	route.GET("/", services.GetTodos)
	route.POST("/todos", services.CraeteTodo)
	route.DELETE("/todos/:id", services.DeleteTodo)
	route.PUT("/todos/:id", services.UpdateTodos)
	route.Run("localhost:8080")
}

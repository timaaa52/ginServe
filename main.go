package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isdone"`
	AddedDate string `json:"addeddate"`
}

var todos = []Todo{
	{ID: "1", Title: "buy bread", IsDone: false, AddedDate: "2026-07-05"},
	{ID: "2", Title: "feed cat", IsDone: true, AddedDate: "2026-07-03"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	todos = append(todos, todo)
	c.IndentedJSON(http.StatusCreated, todos)
}

func main() {
	route := gin.Default()
	route.GET("/", getTodos)
	route.POST("/todos", addTodo)
	route.Run("localhost:8080")
}

package main

import (
	"net/http"
	"slices"

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

func deletToDo(c *gin.Context) {
	id := c.Param("id")

	newTodos := slices.DeleteFunc(todos, func(t Todo) bool {
		return t.ID == id
	})
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "element with id:" + id + " was deleted",
	})

	todos = newTodos
}

func putToDo (c *gin.Context){
	type PutParam struct { 
		Title string `json:"title,omitempty"`
		IsDone bool `json:"isdone,omitempty"`
	}
	id := c.Param("id")
	var puttingParam PutParam

	if err := c.ShouldBindJSON(&puttingParam); err != nil { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todos[i].IsDone = puttingParam.IsDone
			todos[i].Title = puttingParam.Title
		} 
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":"toDo with id " + id + " was been update" ,
	})
}

func main() {
	route := gin.Default()
	route.GET("/", getTodos)
	route.POST("/todos", addTodo)
	route.DELETE("/todos/:id", deletToDo)
	route.PUT("/todos/:id", putToDo)
	route.Run("localhost:8080")
}

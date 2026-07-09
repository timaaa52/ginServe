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

func patchToDo(c *gin.Context) {
	type Param struct {
		Title  *string `json:"title"`
		IsDone *bool   `json:"isdone"`
	}
	id := c.Param("id")
	var params Param

	if err := c.ShouldBindJSON(&params); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	for i, r := range todos {
		if r.ID == id {
			if params.IsDone != nil {
				todos[i].IsDone = *params.IsDone
			}
			if params.Title != nil {
				todos[i].Title = *params.Title
			}
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "toDo with id " + id + " has been update",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})

}

func putToDo(c *gin.Context) {
	type PutParams struct {
		Title  string `json:"title" binding:"required"`
		IsDone bool   `json:"isdone" binding:"required"`
	}
	id := c.Param("id")
	var params PutParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	for i, r := range todos {
		if r.ID == id {
			todos[i].Title = params.Title
			todos[i].IsDone = params.IsDone
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "todo has been update",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "user with id " + id + " not found :(",
	})
}

func main() {
	route := gin.Default()
	route.GET("/", getTodos)
	route.POST("/todos", addTodo)
	route.DELETE("/todos/:id", deletToDo)
	route.PATCH("/todos/:id", patchToDo)
	route.PUT("/todos/:id", putToDo)
	route.Run("localhost:8080")
}

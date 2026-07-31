package services

import (
	"context"
	"net/http"
	"simplehttpserve/database"
	"simplehttpserve/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	var todo models.Todo 
	if err := c.BindJSON(&todo); err != nil { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{ 
			"err": err.Error(),
		})
		return
	}

	_, err := database.DB.Exec(context.Background(), "insert into tododb (title) values($1)", todo.Title)
	if err != nil { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "todo has been created",
	})
}
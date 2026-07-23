package services

import (
	"context"
	"net/http"
	"simplehttpserve/database"
	"simplehttpserve/models"

	"github.com/gin-gonic/gin"
)

func CraeteTodo(c *gin.Context) {

	db, err := database.Connection()
	if err != nil { 
		panic(err.Error())
	}
	defer db.Close()

	var todo models.Todo 
	if err := c.BindJSON(&todo); err != nil { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{ 
			"err": err.Error(),
		})
		return
	}

	_, err = db.Exec(context.Background(), "insert into tododb (title) values($1)", todo.Title)
	if err != nil { 
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "todo has been created",
	})
}
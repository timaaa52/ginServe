package services

import (
	"context"
	"fmt"
	"net/http"

	// "net/http"
	"simplehttpserve/database"
	"simplehttpserve/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	
	db, err := database.Connection()
	if err != nil  { 
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(context.Background(), "select id, title, isdone, added_date from tododb order by id asc")

	if err != nil {
		fmt.Println("error get query row", err.Error())
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo

		err := rows.Scan(&t.ID, &t.Title, &t.IsDone, &t.AddedDate)
		if err != nil {
			fmt.Println("err")
			continue
		}
		todos = append(todos, t)
	}
	c.IndentedJSON(http.StatusOK, todos)

}

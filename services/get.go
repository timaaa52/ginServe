package services

import (
	"context"
	"log"
	"net/http"
	"simplehttpserve/database"
	"simplehttpserve/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {

	rows, err := database.DB.Query(context.Background(), "select id, title, isdone, added_date from tododb order by id asc")

	if err != nil {
		log.Fatalf("error get query row %v", err.Error())
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo

		err := rows.Scan(&t.ID, &t.Title, &t.IsDone, &t.AddedDate)
		if err != nil {
			log.Fatalf("err reading row")
			continue
		}
		todos = append(todos, t)
	}
	c.IndentedJSON(http.StatusOK, todos)

}

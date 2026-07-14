package services

import (
	"context"
	"fmt"
	// "net/http"
	"simplehttpserve/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTodos(c *gin.Context) {
	url := "postgres://postgres:qwerty@localhost:5432/tododb"
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(context.Background(), "select id, title, isdone, added_date from tododb order by id asc")

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
	// c.IndentedJSON(http.StatusOK, todos)
	fmt.Println(todos)

}

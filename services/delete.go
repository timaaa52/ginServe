package services

import (
	"context"
	"net/http"
	"simplehttpserve/database"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	row, err := database.DB.Exec(context.Background(), "delete from tododb where id=$1", id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	rowAffected := row.RowsAffected()
	if rowAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Todo with id " + id + " not found",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Todo with id " + id + " was successfully deleted",
		})
	}
}

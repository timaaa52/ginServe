package services

import (
	"context"
	"net/http"
	"simplehttpserve/database"
	"simplehttpserve/models"
	"github.com/gin-gonic/gin"
)

func UpdateTodos(c *gin.Context) {

	db, err := database.Connection()
	if err != nil { 
		panic(err.Error())
	}

	defer db.Close()

	id := c.Param("id")
	var data models.RequestParams
	if err := c.ShouldBindJSON(&data); err != nil { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad params for request",
		})
		return
	}

	row, err := db.Exec(context.Background(), "update tododb set title=$1, isdone=$2 where id=$3", data.Title, data.IsDone, id)

	if err != nil { 
		panic(err.Error())
	}

	rowAffect := row.RowsAffected()
	if rowAffect == 0 { 
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": "Todo with id " + id + " not found",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{ 
			"succ": "Todo with id " + id + " was successfully updated",
		})
	}

}
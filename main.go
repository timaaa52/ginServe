package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Ping string `json:"ping" binding:"required"`
}

func main() {
	route := gin.Default()

	route.POST("/post", func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"pong": "0",
		})
	})
	route.Run("localhost:8080")

}

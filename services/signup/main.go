package main

import (
	"github.com/carmeloriolo/monorepo/commons/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/signup", func(c *gin.Context) {
		var json models.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Registered new user"})
	})
	r.Run() // listen and serve on 1.0.0.0:8080 (for windows "localhost:8080")
}

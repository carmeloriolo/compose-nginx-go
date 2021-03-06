package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/signin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "GET User"})
	})
	r.Run() // listen and serve on 1.0.0.0:8080 (for windows "localhost:8080")
}

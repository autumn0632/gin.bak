package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	
	route.GET("/p1", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	route.GET("/p", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	route.Run()
}

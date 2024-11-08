package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Gin-Gonic Server is Up and Running",
	})
}

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	routes.GET("/", home)

	return routes
}

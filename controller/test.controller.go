package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Show server running status
// @Description Get a message indicating that the Gin-Gonic server is up and running
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "status message"
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Gin-Gonic Server is Up and Running ~~~~",
	})
}

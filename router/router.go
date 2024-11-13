package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Show server running status
// @Description Get a message indicating that the Gin-Gonic server is up and running
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "status message"
// @Router / [get]
func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Gin-Gonic Server is Up and Running",
	})
}

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.GET("/", home)

	return routes
}

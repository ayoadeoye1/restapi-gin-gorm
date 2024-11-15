package router

import (
	"github.com/ayoadeoye1/restapi-gin-gorm/controller"
	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.GET("/", controller.Home)

	router := routes.Group("/api/v1")
	userRouter := router.Group("/user")

	userRouter.GET("/", controller.GetUser)
	return routes
}

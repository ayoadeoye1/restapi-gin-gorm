package router

import (
	"github.com/ayoadeoye1/restapi-gin-gorm/controller"
	usercontroller "github.com/ayoadeoye1/restapi-gin-gorm/controller/user_controller"
	"github.com/ayoadeoye1/restapi-gin-gorm/middleware"
	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userController *usercontroller.UserController) *gin.Engine {
	routes := gin.Default()

	// routes.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"}, github.com/gin-contrib/cors
	//     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//     AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	// }))

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.GET("/", controller.Home)

	router := routes.Group("/api/v1")
	userRouter := router.Group("/user")
	userRouter.POST("/signup", userController.CreateUser)
	userRouter.POST("/signin", userController.SignIn)

	userRouter.Use(middleware.UserAuth)
	{
		userRouter.GET("/", controller.GetUser)
	}

	return routes
}

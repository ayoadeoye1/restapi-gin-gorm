package usercontroller

import (
	"fmt"
	"net/http"

	"github.com/ayoadeoye1/restapi-gin-gorm/data/requests"
	"github.com/ayoadeoye1/restapi-gin-gorm/data/responses"
	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	userservice "github.com/ayoadeoye1/restapi-gin-gorm/services/user_service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService userservice.UserServiceImpl
}

func NewUserController(userService userservice.UserServiceImpl) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUsers godoc
// @Summary Sign-Up New User
// @Description An endpoint for a new user to sign-up
// @Param users body requests.CreateUserReq true "Create Users"
// @Accept json
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{}
// @Router /api/v1/user [post]
func (userController *UserController) Create(ctx *gin.Context) {
	// log.Fatal(">>> create user") //Avoid This!
	fmt.Println(">>> create user")
	createUserRequest := requests.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	userController.userService.Create(createUserRequest)
	jsonResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, jsonResponse)
}

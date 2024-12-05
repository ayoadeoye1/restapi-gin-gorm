package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ayoadeoye1/restapi-gin-gorm/data/requests"
	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	userservice "github.com/ayoadeoye1/restapi-gin-gorm/services/user_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(">>> create user")
	createUserRequest := requests.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		helper.SendError(ctx, int(http.StateIdle), "Internal Error, Please Try Again", err)
		return
	}

	validate := validator.New()

	err = validate.Struct(createUserRequest)
	if err != nil {
		// validationErrors := formatValidationErrors(err)
		helper.SendError(ctx, http.StatusBadRequest, "Validation Error", error(err))
		fmt.Print("EEEEEEEEÊ", err)
		return
	}

	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		helper.SendError(ctx, int(http.StateIdle), "Internal Error, Please Try Again", err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), salt)
	if err != nil {
		helper.SendError(ctx, int(http.StateIdle), "Internal Error, Please Try Again", err)
		return
	}

	createUserRequest.Password = string(hash)

	userController.userService.Create(createUserRequest)

	helper.SendSuccess(ctx, http.StatusOK, "Sign Up Successful", nil)
}

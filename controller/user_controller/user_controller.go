package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ayoadeoye1/restapi-gin-gorm/data/requests"
	"github.com/ayoadeoye1/restapi-gin-gorm/data/responses"
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
func (userController *UserController) CreateUser(ctx *gin.Context) {
	fmt.Println(">>> create user")
	createUserRequest := requests.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(createUserRequest)
	if err != nil {
		// validationErrors := helper.FormatValidationErrors(err)
		// helper.SendError(ctx, http.StatusBadRequest, "Validation Error", validationErrors)
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := helper.FormatValidationErrors(validationErrs)
			helper.SendError(ctx, http.StatusBadRequest, "Validation Error", formattedErrors)
		} else {
			helper.SendError(ctx, http.StatusBadRequest, "Invalid JSON input", err.Error())
		}
		return
	}

	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), salt)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	createUserRequest.Password = string(hash)

	userController.userService.Create(createUserRequest)

	helper.SendSuccess(ctx, http.StatusOK, "Sign Up Successful", nil)
}

// UserSignIn godoc
// @Summary Sign-In User
// @Description An endpoint for a user to sign-in
// @Param users body requests.LoginReq true "User SignIn"
// @Accept json
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{}
// @Router /api/v1/user [post]
func (userController *UserController) SignIn(ctx *gin.Context) {
	fmt.Println(">>> create user")
	LoginReq := requests.LoginReq{}
	err := ctx.ShouldBindJSON(&LoginReq)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(LoginReq)
	if err != nil {
		// validationErrors := helper.FormatValidationErrors(err)
		// helper.SendError(ctx, http.StatusBadRequest, "Validation Error", validationErrors)
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := helper.FormatValidationErrors(validationErrs)
			helper.SendError(ctx, http.StatusBadRequest, "Validation Error", formattedErrors)
		} else {
			helper.SendError(ctx, http.StatusBadRequest, "Invalid JSON input", err.Error())
		}
		return
	}

	//Find User
	loginUser, err := userController.userService.FindByEmail(LoginReq.Email)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Error occured while searching for user", err.Error())
		return
	}

	if loginUser == (responses.UserResponse{}) {
		helper.SendError(ctx, http.StatusBadRequest, "Account with email does not exist", nil)
		return
	}

	//Verify Password

	//Generate JWT Token

	helper.SendSuccess(ctx, http.StatusOK, "Sign Up Successful", nil)
}

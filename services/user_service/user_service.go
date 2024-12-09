package userservice

import (
	"errors"

	"github.com/ayoadeoye1/restapi-gin-gorm/data/requests"
	"github.com/ayoadeoye1/restapi-gin-gorm/data/responses"
	"gorm.io/gorm"

	// "github.com/ayoadeoye1/restapi-gin-gorm/helper"
	"github.com/ayoadeoye1/restapi-gin-gorm/models"
	userrepository "github.com/ayoadeoye1/restapi-gin-gorm/repository/user_repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository userrepository.UserRepo
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository userrepository.UserRepo, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Create(users requests.CreateUserReq) error {
	err := u.Validate.Struct(users)
	if err != nil {
		return err
	}

	userModel := models.Users{
		FirstName:  users.FirstName,
		LastName:   users.LastName,
		Email:      users.Email,
		Password:   users.Password,
		Occupation: users.Occupation,
		Address:    users.Address,
	}
	u.UserRepository.Add(userModel)
	return nil
}

// func (u *UserServiceImpl) FindByEmail(userEmail string) (user responses.UserResponse, err error) {
// 	user, err := u.UserRepository.FindByEmail(userEmail)
// 	if err != nil {
// 		return responses.UserResponse{}, err
// 	}

// 	userResponse := responses.UserResponse{
// 		ID:         user.ID,
// 		FirstName:  user.FirstName,
// 		LastName:   user.LastName,
// 		Email:      user.Email,
// 		Occupation: user.Occupation,
// 		Address:    user.Address,
// 	}

// 	return userResponse, nil
// }

func (u *UserServiceImpl) FindByEmail(userEmail string) (responses.UserResponse, error) {
	user, err := u.UserRepository.FindByEmail(userEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.UserResponse{}, errors.New("user not found")
		}
		return responses.UserResponse{}, err
	}

	userResponse := responses.UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Occupation: user.Occupation,
		Address:    user.Address,
	}

	return userResponse, nil
}

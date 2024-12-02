package userrepository

import "github.com/ayoadeoye1/restapi-gin-gorm/models"

type UserRepo interface {
	Add(users models.Users) error
	Edit(users models.Users) error
	FindAll(userId int) []models.Users
	FindById(userId int) (users models.Users, err error)
	Remove(userId int) error
}

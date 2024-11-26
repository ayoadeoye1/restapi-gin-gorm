package userrepository

import "github.com/ayoadeoye1/restapi-gin-gorm/models"

type UserRepo interface {
	Add(users models.Users)
	Edit(users models.Users)
	Remove(userId int)
	FindById(userId int) (users models.Users, err error)
	FindAll(userId int) []models.Users
}

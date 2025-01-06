package userrepository

import "github.com/ayoadeoye1/restapi-gin-gorm/models"

type UserRepo interface {
	Add(users models.Users) error
	Edit(users models.Users) error
	FindAll() ([]models.Users, error)
	FindByEmail(string) (*models.Users, error)
	FindById(userId int) (*models.Users, error)
	Remove(userId int) error
}

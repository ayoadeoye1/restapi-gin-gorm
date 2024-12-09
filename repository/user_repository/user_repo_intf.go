package userrepository

import "github.com/ayoadeoye1/restapi-gin-gorm/models"

//	type UserRepo interface {
//		Add(users models.Users) error
//		Edit(users models.Users) error
//		FindAll(userId int) []models.Users
//		FindById(userId int) (users models.Users, err error)
//		FindByEmail(userEmail string) (users models.Users, err error)
//		Remove(userId int) error
//	}
type UserRepo interface {
	Add(users models.Users) error
	Edit(users models.Users) error
	FindAll(userId int) ([]models.Users, error)
	FindByEmail(string) (*models.Users, error)
	FindById(userId int) (*models.Users, error)
	Remove(userId int) error
}

// type UserRepo interface {
// 	Add(users models.Users) error
// 	Edit(users models.Users) error
// 	FindAll(userId int) ([]models.Users, error)
// 	FindById(userId int) (*models.Users, error)
// 	FindByEmail(userEmail string) (*models.Users, error)
// 	Remove(userId int) error
// }

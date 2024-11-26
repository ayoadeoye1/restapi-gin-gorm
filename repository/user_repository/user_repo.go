package userrepository

import (
	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	"github.com/ayoadeoye1/restapi-gin-gorm/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepoImpl(Db *gorm.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) Add(users models.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Edit implements UserRepo.
func (u *UserRepoImpl) Edit(users models.Users) {
	panic("unimplemented")
}

// FindAll implements UserRepo.
func (u *UserRepoImpl) FindAll(userId int) []models.Users {
	panic("unimplemented")
}

// FindById implements UserRepo.
func (u *UserRepoImpl) FindById(userId int) (users models.Users, err error) {
	panic("unimplemented")
}

// Remove implements UserRepo.
func (u *UserRepoImpl) Remove(userId int) {
	panic("unimplemented")
}

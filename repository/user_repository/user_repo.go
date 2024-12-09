package userrepository

import (
	"github.com/ayoadeoye1/restapi-gin-gorm/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepoImpl(Db *gorm.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) Add(users models.Users) error {
	result := u.Db.Create(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepoImpl) Edit(users models.Users) error {
	result := u.Db.Save(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepoImpl) FindAll(userId int) []models.Users {
	var users []models.Users
	u.Db.Where("id = ?", userId).Find(&users)
	return users
}

func (u *UserRepoImpl) FindById(userId int) (users models.Users, err error) {
	result := u.Db.First(&users, userId)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (u *UserRepoImpl) FindByEmail(userEmail string) (*models.Users, error) {
	var users models.Users
	result := u.Db.Where("email = ?", userEmail).First(&users)
	if result.Error != nil {
		return &users, result.Error
	}
	return &users, nil
}

func (u *UserRepoImpl) Remove(userId int) error {
	result := u.Db.Delete(&models.Users{}, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package usermodel

import (
	"time"

	otpmodel "github.com/ayoadeoye1/restapi-gin-gorm/model/otp_model"
)

type Users struct {
	ID         uint   `gorm:"type:int;primaryKey;autoIncrement"`
	FirstName  string `gorm:"type:varchar(255);not null"`
	LastName   string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);unique;not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	Occupation string `gorm:"type:varchar(255)"`
	Address    string `gorm:"type:varchar(255)"`

	//time stamps
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:datetime;autoUpdateTime"`

	//associations
	Otp []otpmodel.Otp `gorm:"foreignKey:UserID"`
}

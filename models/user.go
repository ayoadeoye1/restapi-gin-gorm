package models

import (
	"time"
)

type Users struct {
	ID         int     `gorm:"type:int;primaryKey;autoIncrement"`
	FirstName  string  `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName   string  `gorm:"type:varchar(255);not null" json:"last_name"`
	Email      string  `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password   string  `gorm:"type:varchar(255);not null" json:"password"`
	Occupation *string `gorm:"type:varchar(255)" json:"occupation"`
	Address    *string `gorm:"type:varchar(255)" json:"address"`

	//time stamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	//associations
	Otp *Otp `gorm:"foreignKey:UserID"`
}

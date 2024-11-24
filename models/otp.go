package models

import (
	"time"
)

type Otp struct {
	ID            int       `gorm:"type:int;primaryKey;autoIncrement"`
	UserID        string    `gorm:"type:int;unique;not null"`
	Email         string    `gorm:"type:varchar(255);unique;;not null"`
	Otp           string    `gorm:"type:varchar(255);not null"`
	ExpiresAt     time.Time `gorm:"type:datetime;not null"`
	EmailVerified bool      `gorm:"type:boolean;default:false"`

	//time stamps
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:datetime;autoUpdateTime"`

	//associations
	Users *Users `gorm:"foreignKey:UserID"`
}

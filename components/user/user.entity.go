package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `gorm:"type:varchar(80);not null" json:"name"`
	LastName string    `gorm:"type:varchar(80); not null" json:"lastname"`
	BirtDay  time.Time `gorm:"type:date;not null" json:"birtday"`
	Email    string    `gorm:"type:varchar(256);not null" json:"email"`
	UserName string    `gorm:"type:varchar(100);not null" json:"username"`
}

func (User) TableName() string {
	return "users"
}

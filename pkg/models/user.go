package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("email = ?", email).First(&getUser)
	return &getUser, db
}

func CreateUser(user *User) {
	db.NewRecord(user)
	db.Create(&user)
}

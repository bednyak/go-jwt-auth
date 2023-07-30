package models

import (
	"github.com/bednyak/go-react-jwt-auth/pkg/app"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func ConnectToDb() {
	app.Connect()
	db = app.GetDB()
	db.AutoMigrate(&User{})
}

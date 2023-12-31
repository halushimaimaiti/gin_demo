package models

import (
	"gin_test/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id    uint
	Name  string
	Email string
}

func InitDataBase() {
	config.Db.AutoMigrate(&User{})
}

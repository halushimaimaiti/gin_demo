package models

import (
	"gin_test/config"
)

type Book struct {
	Id   uint
	Name string `gorm:"not null"`
}

type BookInfo struct {
	Id     uint
	BookId uint `gorm:"not null"`
	Info   string
}

func InitDataBase() {
	config.SqlLiteDB.AutoMigrate(&Book{}, &BookInfo{})
}

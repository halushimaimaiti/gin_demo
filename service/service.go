package service

import (
	"gin_test/dao"
	"gin_test/response"
)

func GetUser() response.User {
	var user response.User
	user = dao.GetUser()
	return user
}

func FindBookByInfo(s string, page int32) []response.BookInfo {
	var bookInfo []response.BookInfo
	bookInfo = dao.FindBookByInfo(s, page, 10)
	return bookInfo
}

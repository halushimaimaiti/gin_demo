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

package dao

import (
	"context"
	"gin_test/config"
	"gin_test/response"
)

func GetUser() response.User {
	var user response.User
	config.Db.First(&user, 1)
	config.Rdb.Set(context.Background(), "test", 1, 0)
	return user
}

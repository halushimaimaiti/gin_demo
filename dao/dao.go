package dao

import (
	"gin_test/config"
	"gin_test/response"
)

func GetUser() response.User {
	var user response.User
	config.Db.First(&user, 1)
	// config.Rdb.Set(context.Background(), "test", 1, 0)
	return user
}

func FindBookByInfo(s string, page int32, count int32) []response.BookInfo {
	var bookInfo []response.BookInfo
	sql := "select b.name as book_name,bi.info from books b inner join book_infos bi on b.id = bi.book_id where bi.info like ? limit ? offset ?"
	if err := config.SqlLiteDB.Raw(sql, "%"+s+"%", count, (page-1)*count).Scan(&bookInfo).Error; err != nil {
		config.Log.Error("Failed to execute raw query")
	}
	for i := 0; i < len(bookInfo); i++ {
		bookInfo[i].No = uint32((page-1)*count + int32(i) + 1)
	}
	return bookInfo
}

package response

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BookInfo struct {
	No       uint32 `json:"No"`
	BookName string `json:"BookName"`
	Info     string `json:"Info"`
}

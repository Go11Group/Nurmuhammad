package structs

type User struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfoRes struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

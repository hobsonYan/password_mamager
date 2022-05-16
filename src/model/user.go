package model

type RegisterForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Time     int64  `json:"time"`
}

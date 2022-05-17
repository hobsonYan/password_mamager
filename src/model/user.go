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

type UpdateForm struct {
	Id           string `form:"id"`
	New_Username string `form:"newUsername"`
	New_Password string `form:"newPassword"`
	New_Phone    string `form:"newPhone"`
}

type UserVo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

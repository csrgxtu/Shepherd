package models

type (
	User struct {
		Email  string `form:"email"`
		Passwd string `form:"passwd"`
	}
)

package controllers

import (
	"Shepherd/models"
	"fmt"

	"github.com/martini-contrib/render"
)

func UserLogin(r render.Render, user models.User) {
	if user.Email == "1066844321@qq.com" {
		r.HTML(200, "main", nil)
	} else {
		r.HTML(404, "404", nil)
	}
	fmt.Println(user)
}

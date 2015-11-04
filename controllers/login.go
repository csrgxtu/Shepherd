package controllers

import (
	"Shepherd/models"
	"Shepherd/services"
	"fmt"
	"github.com/martini-contrib/render"
)

func UserLogin(r render.Render, user models.User) {
	fmt.Println(user.Email, " ", user.Passwd)
	if user.Email == "1066844321@qq.com" {
		services.AddSession(user.Email)
		fmt.Println("if ", services.ReturnSession())
		r.HTML(200, "main", nil)
	} else {
		fmt.Println("else ", services.ReturnSession())
		r.HTML(404, "404", nil)
	}

}

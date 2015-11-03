package controllers

import (
	"fmt"
	"github.com/go-martini/martini"
	"strconv"
)

func UserLogin(params martini.Params) {
	var email, _ = strconv.Atoi(params["email"])
	var passwd, _ = strconv.Atoi(params["passwd"])
	fmt.Println(email)
	fmt.Println(passwd)
}

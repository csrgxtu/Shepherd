package controllers

import (
  // "time"
  // "strconv"

  // "labix.org/v2/mgo/bson"
  // "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "Shepherd/models"
)

func Auth(r render.Render, user models.User) {
  var rt models.Result

  if user.Username == "archer" && user.Password == "archer" {
    rt.Code = 200
    rt.Msg = "Successful"
  } else {
    rt.Code = 403
    rt.Msg = "Not Authorized"
  }

  r.JSON(200, rt)
}

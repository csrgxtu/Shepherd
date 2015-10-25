package controllers

import (
  "github.com/martini-contrib/render"

  "Shepherd/models"
)

func Create(r render.Render) {
  var rt models.Result

  rt.Code = 200
  rt.Msg = "Successful"
  rt.ResNum = 1
  
  r.JSON(200, rt)
}

func Read() string {
  return "Read GPS Data"
}

func Update() string {
  return "Update GPS Data"
}

func Delete() string {
  return "Delete GPS Data"
}

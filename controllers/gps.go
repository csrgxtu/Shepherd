package controllers

import (
  "fmt"
  
  "github.com/martini-contrib/render"

  "Shepherd/models"
)

func Create(r render.Render, gps models.Gps) {
  var rt models.Result

  fmt.Println(gps)

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

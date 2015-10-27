package controllers

import (
  "github.com/martini-contrib/render"

  "Shepherd/models"
  // "Shepherd/inits"
)

func Create(r render.Render) {
  var rt models.Result



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

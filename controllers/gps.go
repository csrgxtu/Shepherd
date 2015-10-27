package controllers

import (
  "time"

  "labix.org/v2/mgo/bson"
  "github.com/martini-contrib/render"

  "Shepherd/models"
  "Shepherd/services"
)

func Create(r render.Render, gps models.Gps) {
  gps.Id = bson.NewObjectId().Hex()
  gps.CreatedAt = time.Now()
  var rt models.Result

  gps, err := services.Create(gps)
  if err != nil {
    rt.Code = 500
    rt.Msg = "Server Internal Error"
  }

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

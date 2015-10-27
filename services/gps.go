package services

import (
  "Shepherd/models"
  "Shepherd/inits"
)

func Create(gps models.Gps)(ngps models.Gps, err error) {
  err = inits.Session.DB("shepherd").C("gps").Insert(&gps)
  ngps = gps
  return
}

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

func Read(gps models.Gps, start, offset int)(ngps []models.Gps, err error) {
  err = inits.Session.DB("shepherd").C("gps").Find(nil).Sort("-createdAt").Skip(start).Limit(offset).All(&ngps)
  return
}

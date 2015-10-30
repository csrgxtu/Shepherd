package services

import (
  "labix.org/v2/mgo/bson"
  "Shepherd/models"
  "Shepherd/inits"
)

func Create(gps models.Gps)(ngps models.Gps, err error) {
  err = inits.Session.DB("shepherd").C("gps").Insert(&gps)
  ngps = gps
  return
}

func Read(gps models.Gps, imei string, start, offset int)(ngps []models.Gps, err error) {
  if imei != "" {
    err = inits.Session.DB("shepherd").C("gps").Find(bson.M{"imei": imei}).Sort("-createdAt").Skip(start).Limit(offset).All(&ngps)
  } else {
    err = inits.Session.DB("shepherd").C("gps").Find(nil).Sort("-createdAt").Skip(start).Limit(offset).All(&ngps)
  }
  return
}

func GetDistance(imei1, imei2 string)(ngps []models.Gps, err error) {
  var gps models.Gps
  err = inits.Session.DB("shepherd").C("gps").Find(bson.M{"imei": imei1}).Sort("-createdAt").Limit(1).One(&gps)
  ngps = append(ngps, gps)
  err = inits.Session.DB("shepherd").C("gps").Find(bson.M{"imei": imei2}).Sort("-createdAt").Limit(1).One(&gps)
  ngps = append(ngps, gps)
  return
}

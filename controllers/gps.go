package controllers

import (
  "time"
  "strconv"

  "labix.org/v2/mgo/bson"
  "github.com/go-martini/martini"
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
  } else {
    rt.Code = 200
    rt.Msg = "Successful"
    rt.ResNum = 1
    rt.Data = make([]models.Recs, 1)
    rt.Data[0] = gps
  }

  r.JSON(200, rt)
}

// /gps/:start/:offset
func Read(r render.Render, params martini.Params) {
  var imei = params["imei"]
  var start, _ = strconv.Atoi(params["start"])
  var offset, _ = strconv.Atoi(params["offset"])
  var rt models.Result
  var gps models.Gps

  ngps, err := services.Read(gps, imei, start, offset)
  if err != nil {
    rt.Code = 500
    rt.Msg = "Server Internal Error"
  } else {
    rt.Code = 200
    rt.Msg = "Successful"
    rt.ResNum = int32(len(ngps))
    rt.Data = make([]models.Recs, rt.ResNum)
    for ix, value := range ngps {
			rt.Data[ix] = value
		}
  }

  r.JSON(200, rt)
}

func GetDistance(r render.Render, params martini.Params) {
  var rt models.Result
  r.JSON(200, rt)
}

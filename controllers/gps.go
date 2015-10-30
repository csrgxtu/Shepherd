package controllers

import (
  "time"
  "strconv"

  "labix.org/v2/mgo/bson"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "github.com/kellydunn/golang-geo"

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
  var imei1 = params["imei1"]
  var imei2 = params["imei2"]
  var rt models.Result

  ngps, err := services.GetDistance(imei1, imei2)
  if err != nil {
    rt.Code = 500
    rt.Msg = "Server Internal Error"
  } else {
    point1 := geo.NewPoint(float64(ngps[0].Loc.Longitude), float64(ngps[0].Loc.Latitude))
    point2 := geo.NewPoint(float64(ngps[1].Loc.Longitude), float64(ngps[1].Loc.Latitude))
    dist := point1.GreatCircleDistance(point2)

    rt.Code = 200
    rt.Msg = "Successful"
    rt.ResNum = int32(1)
    rt.Distance = dist
  }

  r.JSON(200, rt)
}

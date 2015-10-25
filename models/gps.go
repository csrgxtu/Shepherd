package models

import (
  "time"
)

type (
  gps struct {
    ID string `json:"id"`
    IMEI string `json:"imei"`
    CreatedAt time.Time `json:"createdAt"`
    Loc loc `json:"loc"`
  }

  loc struct {
    Longitude float32 `json:"longitude"`
    Latitude float32 `json:"latitude"`
  }
)

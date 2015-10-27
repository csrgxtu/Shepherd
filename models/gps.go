package models

import (
  "time"
)

type (
  Gps struct {
    ID string `json:"id"`
    IMEI string `json:"imei" binding:"required"`
    CreatedAt time.Time `json:"createdAt"`
    Loc loc `json:"loc" binding:"required"`
  }

  loc struct {
    Longitude float32 `json:"longitude"`
    Latitude float32 `json:"latitude"`
  }
)

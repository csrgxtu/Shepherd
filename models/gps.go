package models

import (
  "time"
)

type (
  Gps struct {
    // ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Id string `json:"_id" bson:"_id,omitempty"`
    IMEI string `json:"imei" bson:"imei" binding:"required"`
    CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
    Loc loc `json:"loc" bson:"loc" binding:"required"`
  }

  loc struct {
    Longitude float32 `json:"longitude" bson:"longitude"`
    Latitude float32 `json:"latitude" bson:"latitude"`
  }

  User struct {
    Id string `json:"_id" bson:"_id,omitempty"`
    Username string `json:"Username" bson:"Username"`
    Password string `json:"Password" bson"Password"`
  }
)

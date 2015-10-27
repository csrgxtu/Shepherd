package routers

import (
  "Shepherd/inits"
  "Shepherd/controllers"
)

func GpsRouters() {
  inits.Shepherd.Put("/gps", controllers.Create)
  inits.Shepherd.Get("/gps", controllers.Read)
  inits.Shepherd.Post("/gps", controllers.Update)
  inits.Shepherd.Delete("/gps", controllers.Delete)
}

package routers

import (
  "github.com/martini-contrib/binding"

  "Shepherd/inits"
  "Shepherd/controllers"
  "Shepherd/models"
)

func GpsRouters() {
  inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
  inits.Shepherd.Get("/gps", controllers.Read)
  inits.Shepherd.Post("/gps", controllers.Update)
  inits.Shepherd.Delete("/gps", controllers.Delete)
}

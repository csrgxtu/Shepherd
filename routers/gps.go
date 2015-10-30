package routers

import (
  "github.com/martini-contrib/binding"

  "Shepherd/inits"
  "Shepherd/controllers"
  "Shepherd/models"
)

func GpsRouters() {
  inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
  // inits.Shepherd.Get("/gps/:start/:offset", controllers.Read)
  inits.Shepherd.Get("/gps/:imei1/:imei2", controllers.GetDistance)
}

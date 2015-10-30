package routers

import (
	"github.com/martini-contrib/binding"

	"Shepherd/controllers"
	"Shepherd/inits"
	"Shepherd/models"
)

func GpsRouters() {
<<<<<<< HEAD
	inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
	inits.Shepherd.Get("/gps/:start/:offset", controllers.Read)
	inits.Shepherd.Get("/gps/:imei1/:imei2", controllers.GetDistance)
	inits.Shepherd.NotFound(func() string {
		return "this not router!!<br><h1>NOT FOUND<h1>"
	})
=======
  inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
  // inits.Shepherd.Get("/gps/:start/:offset", controllers.Read)
  inits.Shepherd.Get("/gps/:imei1/:imei2", controllers.GetDistance)
>>>>>>> fcd834bdbd7cbe3a7912ec87fce5e8a4aea7cea1
}

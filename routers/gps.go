package routers

import (
	"Shepherd/controllers"
	"Shepherd/inits"
	"Shepherd/models"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func GpsRouters() {
	inits.Shepherd.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})
	inits.Shepherd.Post("/userLogin", controllers.UserLogin)
	inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
	inits.Shepherd.Get("/gps/:start/:offset", controllers.Read)
	inits.Shepherd.Get("/gps/:imei1/:imei2", controllers.GetDistance)
	inits.Shepherd.NotFound(func() string {
		return "this not router!! NOT FOUND"
	})
	inits.Shepherd.Put("/gps", binding.Bind(models.Gps{}), controllers.Create)
	// inits.Shepherd.Get("/gps/:start/:offset", controllers.Read)
	inits.Shepherd.Get("/gps/:imei1/:imei2", controllers.GetDistance)
}

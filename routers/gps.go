package routers

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"

	"Shepherd/controllers"
	"Shepherd/inits"
	"Shepherd/models"
)

func GpsRouters() {
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

func AuthRouters() {
	allowCORSHandler := cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		// AllowMethods:     []string{"POST", "GET"},
		// AllowHeaders:     []string{"Origin"},
	})

	inits.Shepherd.Post("/auth", allowCORSHandler, binding.Bind(models.User{}), controllers.Auth)
}

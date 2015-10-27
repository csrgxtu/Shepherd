package main

import (
  // "github.com/go-martini/martini"
  // "github.com/martini-contrib/render"

  "Shepherd/controllers"
  "Shepherd/inits"
)

func main() {
  inits.Shepherd.Put("/gps", controllers.Create)
  inits.Shepherd.Get("/gps", controllers.Read)
  inits.Shepherd.Post("/gps", controllers.Update)
  inits.Shepherd.Delete("/gps", controllers.Delete)

  inits.Shepherd.Run()
}

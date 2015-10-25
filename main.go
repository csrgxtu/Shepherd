package main

import (
  "github.com/go-martini/martini"

  "Shepherd/controllers"
)

func main() {
  m := martini.Classic()

  m.Put("/gps", controllers.Create)
  m.Get("/gps", controllers.Read)
  m.Post("/gps", controllers.Update)
  m.Delete("/gps", controllers.Delete)

  m.Get("/", func() string {
    return "Welcome to Shepherd Project!"
  })

  m.Run()
}

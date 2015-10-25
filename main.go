package main

import (
  "fmt"
  
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "Shepherd/controllers"
  "Shepherd/confs"
)

func main() {
  fmt.Println(confs.ReadConfig())

  m := martini.Classic()

  m.Use(render.Renderer())

  m.Put("/gps", controllers.Create)
  m.Get("/gps", controllers.Read)
  m.Post("/gps", controllers.Update)
  m.Delete("/gps", controllers.Delete)

  m.Run()
}

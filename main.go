package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "Shepherd/controllers"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer())

  m.Put("/gps", controllers.Create)
  m.Get("/gps", controllers.Read)
  m.Post("/gps", controllers.Update)
  m.Delete("/gps", controllers.Delete)

  m.Get("/", func(r render.Render){
    // r.HTML(200, "hello", "jeremy")
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Run()
}

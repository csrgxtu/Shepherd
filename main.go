package main

import (
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()

  m.Put("/gps", func() string {
    return "create gps data"
  })

  m.Get("/gps", func() string {
    return "get gps data"
  })
  // m.Get("/gps/:id", func() string {
  //   return "get a specific gps data"
  // })
  // m.Get("/gps/:imei", func() string {
  //   return "get gps data related to a device"
  // })

  m.Post("/gps", func() string {
    return "update gps data"
  })

  m.Delete("/gps", func() string {
    return "delete gps data"
  })

  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Run()
}

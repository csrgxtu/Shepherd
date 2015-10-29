package main
//test
import (
  "Shepherd/inits"
  "Shepherd/routers"
)

func main() {
  routers.GpsRouters()

  inits.Shepherd.Run()
}

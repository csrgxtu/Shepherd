package main

//test
import (
	"Shepherd/inits"
	"Shepherd/routers"
)

func main() {
	routers.GpsRouters()
	routers.AuthRouters()

	inits.Shepherd.RunOnAddr(":8000")
}

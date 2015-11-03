package inits

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var Shepherd *martini.ClassicMartini

func init() {
	Shepherd = martini.Classic()
	Shepherd.Use(martini.Static("public"))
	Shepherd.Use(render.Renderer())
}

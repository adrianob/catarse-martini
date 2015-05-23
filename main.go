package main

import (
	"catarse/app/controllers"
	"catarse/app/models"
	"catarse/config"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	app := martini.Classic()
	app.Map(config.Db())
	app.Use(render.Renderer())

	app.Group("/users", func(router martini.Router) {
		router.Post("", binding.Json(models.User{}), controllers.UsersCreate)
		router.Get("/:id", controllers.UsersShow)
		router.Delete("/:id", controllers.UsersDelete)
		router.Put("/:id", binding.Json(models.User{}), controllers.UsersUpdate)
	})

	app.Run()
}

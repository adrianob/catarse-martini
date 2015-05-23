package controllers

import (
	"catarse/app/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/gorp.v1"
	"strconv"
)

func UsersCreate(user models.User, db *gorp.DbMap, render render.Render) {
	err := db.Insert(&user)
	if err == nil {
		render.JSON(201, &user)
	} else {
		render.JSON(422, map[string]string{"error": err.Error()})
	}
}

func UsersShow(dbmap *gorp.DbMap, params martini.Params, render render.Render) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	user, err := dbmap.Get(models.User{}, id)

	if err == nil {
		render.JSON(200, user)
	} else {
		render.JSON(404, map[string]string{"error": err.Error()})
	}

}

func UsersDelete(db *gorp.DbMap, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	_, err := db.Delete(&models.User{Id: id})

	if err == nil {
		render.JSON(204, "No content")
	} else {
		render.JSON(404, "Not found")
	}
}

func UsersUpdate(user models.User, db *gorp.DbMap, params martini.Params, render render.Render) {
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	user.Id = id
	_, err := db.Update(&user)

	if err == nil {
		render.JSON(200, user)
	} else {
		render.JSON(422, err.Error())
	}
}

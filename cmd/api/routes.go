package main

import (
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/v1/data", app.uploadDataHandler)
	return router
}

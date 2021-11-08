package main

import (
	"github.com/SonuKanwar/restapi"

	"github.com/julienschmidt/httprouter"
)

// addRouteHandlers adds routes for various APIs.
func addRouteHandlers(router *httprouter.Router) {

	//-----------------------API Routes-----------------------------------
	router.POST("/articles", restapi.AddClient)

	router.GET("/articles/:article_id", restapi.GetClient)
	router.GET("/articles/all", restapi.GetClient)
}

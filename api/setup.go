package api

import (
	"github.com/JuniorJDS/data-handler-api/api/route"
	"github.com/gorilla/mux"
)

func HttpHandler() *mux.Router {
	router := mux.NewRouter()

	// define routes
	helloRoute := route.NewHelloWorld()
	router.HandleFunc("/", helloRoute.GetHelloWorld).Methods("GET")

	return router
}

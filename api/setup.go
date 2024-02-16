package api

import (
	"github.com/JuniorJDS/data-handler-api/api/route"
	"github.com/gorilla/mux"
)

func HttpHandler() *mux.Router {
	router := mux.NewRouter()

	uploadRoute := route.NewFileUploadRoute()
	router.HandleFunc("/api/v1/upload", uploadRoute.UploadFile).Methods("POST")

	return router
}

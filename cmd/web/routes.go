package main

import (
	"net/http"

	_ "github.com/archit-p/MicroserviceTemplate/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// routes is used to setup routes for the application
func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.Use(app.secureHeaders, app.logRequest, app.recoverPanic)

	router.HandleFunc("/sample/{id}", app.getSample).Methods(http.MethodGet)
	router.HandleFunc("/sample/{id}", app.updateSample).Methods(http.MethodPut)
	router.HandleFunc("/sample/{id}", app.deleteSample).Methods(http.MethodDelete)
	router.HandleFunc("/sample/create", app.createSample).Methods(http.MethodPost)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}

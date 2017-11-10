package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func GetIndex() http.Handler {
	router := mux.NewRouter()
	router.Handle("/cities", CityHandler())
	return router
}

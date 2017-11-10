package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func GetCities(w http.ResponseWriter, r *http.Request) {
}

func CityHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", GetCities)
	return router
}

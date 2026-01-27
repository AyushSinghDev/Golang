package router

import (
	"github.com/AyushSinghDev/25_mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMoive).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMoives).Methods("DELETE")

	return router
}

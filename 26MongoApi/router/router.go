package router

import (
	"github.com/abu-Hadia/mongoapi/controller"
	"github.com/gorilla/mux"
)

// function router

func router() *mux.Router {
	router := mux.NewRouter()

	router.handleFunc("/api/movies", controller.GetMyALLMovvies).Methods("GET")
	router.handleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.handleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.handleFunc("/api/deleteallmovie{id}", controller.DeleteAllMovies).Methods("DELETE")
	//router.handleFunc("/api/movie/{id}",controller.markAswatched).Methods("PUST")

	return router
}

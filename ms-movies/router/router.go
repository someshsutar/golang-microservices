package router

import (
	"github.com/gorilla/mux"
	"github.com/someshsutar/golang-microservices/ms-movies/handlers"
)

func RegisterRoutes(movieHandler *handlers.MovieHandler) *mux.Router {

	// Initialize Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies", movieHandler.AddMovie).Methods("POST")

	return r
}

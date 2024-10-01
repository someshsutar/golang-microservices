package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/someshsutar/golang-microservices/ms-movies/models"
	"github.com/someshsutar/golang-microservices/ms-movies/service"
)

type MovieHandler struct {
	service service.MovieService
}

func NewMovieHandler(service service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddMovie(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.service.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

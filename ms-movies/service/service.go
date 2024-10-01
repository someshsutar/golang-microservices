package service

import (
	"github.com/someshsutar/golang-microservices/ms-movies/models"
	"github.com/someshsutar/golang-microservices/ms-movies/repository"
)

type MovieService interface {
	AddMovie(movie *models.Movie) error
	GetMovies() ([]models.Movie, error)
	GetMovieByID(id string) (*models.Movie, error)
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) AddMovie(movie *models.Movie) error {
	return s.repo.InsertMovie(movie)
}

func (s *movieService) GetMovies() ([]models.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *movieService) GetMovieByID(id string) (*models.Movie, error) {
	return s.repo.GetMovieByID(id)
}

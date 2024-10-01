package main

import (
	"log"
	"net/http"

	"github.com/someshsutar/golang-microservices/ms-movies/database"
	"github.com/someshsutar/golang-microservices/ms-movies/handlers"
	"github.com/someshsutar/golang-microservices/ms-movies/repository"
	"github.com/someshsutar/golang-microservices/ms-movies/router"
	"github.com/someshsutar/golang-microservices/ms-movies/service"
)

func main() {
	// Initialize MongoDB
	client, err := database.ConnectMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("mflix")

	// Initialize Repositories and Services
	movieRepo := repository.NewMongoMovieRepository(db, "movies")
	movieService := service.NewMovieService(movieRepo)
	movieHandler := handlers.NewMovieHandler(movieService)

	// Initialize Router
	r := router.RegisterRoutes(movieHandler)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

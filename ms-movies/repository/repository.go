package repository

import (
	"context"
	"time"

	"github.com/someshsutar/golang-microservices/ms-movies/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	InsertMovie(movie *models.Movie) error
	GetAllMovies() ([]models.Movie, error)
	GetMovieByID(id string) (*models.Movie, error)
}

type MongoMovieRepository struct {
	collection *mongo.Collection
}

func NewMongoMovieRepository(db *mongo.Database, collectionName string) *MongoMovieRepository {
	return &MongoMovieRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoMovieRepository) InsertMovie(movie *models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *MongoMovieRepository) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MongoMovieRepository) GetMovieByID(id string) (*models.Movie, error) {
	var movie models.Movie

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&movie)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return &movie, err
}

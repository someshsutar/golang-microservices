package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie represents a movie document in MongoDB
type Movie struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Plot             string             `bson:"plot" json:"plot"`
	Genres           []string           `bson:"genres" json:"genres"`
	Runtime          int                `bson:"runtime" json:"runtime"`
	Cast             []string           `bson:"cast" json:"cast"`
	NumMflixComments int                `bson:"num_mflix_comments" json:"num_mflix_comments"`
	Title            string             `bson:"title" json:"title"`
	FullPlot         string             `bson:"fullplot" json:"fullplot"`
	Countries        []string           `bson:"countries" json:"countries"`
	Released         primitive.DateTime `bson:"released" json:"released"`
	Directors        []string           `bson:"directors" json:"directors"`
	Rated            string             `bson:"rated" json:"rated"`
	Awards           Awards             `bson:"awards" json:"awards"`
	LastUpdated      string             `bson:"lastupdated" json:"lastupdated"`
	Year             int                `bson:"year" json:"year"`
	IMDB             IMDB               `bson:"imdb" json:"imdb"`
	Type             string             `bson:"type" json:"type"`
	Tomatoes         Tomatoes           `bson:"tomatoes" json:"tomatoes"`
}

// Awards represents the awards field in a movie document
type Awards struct {
	Wins        int    `bson:"wins" json:"wins"`
	Nominations int    `bson:"nominations" json:"nominations"`
	Text        string `bson:"text" json:"text"`
}

// IMDB represents the IMDb information for the movie
type IMDB struct {
	Rating float64 `bson:"rating" json:"rating"`
	Votes  int     `bson:"votes" json:"votes"`
	ID     int     `bson:"id" json:"id"`
}

// Tomatoes represents the Rotten Tomatoes rating for the movie
type Tomatoes struct {
	Viewer      Viewer             `bson:"viewer" json:"viewer"`
	LastUpdated primitive.DateTime `bson:"lastUpdated" json:"lastUpdated"`
}

// Viewer represents the viewer score from Rotten Tomatoes
type Viewer struct {
	Rating     float64 `bson:"rating" json:"rating"`
	NumReviews int     `bson:"numReviews" json:"numReviews"`
	Meter      int     `bson:"meter" json:"meter"`
}

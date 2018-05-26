package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeppyS/what-to-watch/scraper/imdb"
	"github.com/PeppyS/what-to-watch/scraper/rottentomatoes"
)

type APIClient struct {
	httpClient *http.Client
	url string
}

type IMDBMeta struct {
	Genre      string  `json:"genre"`
	MPAARating string  `json:"mpaa_rating"`
	Score      float64 `json:"score"`
}

type RottenTomatoesMeta struct {
	TomatoScore        int    `json:"tomato_score"`
	PopcornScore       int    `json:"popcorn_score"`
	TheaterReleaseDate string `json:"theater_release_date"`
	MpaaRating         string `json:"mpaa_rating"`
	Synopsis           string `json:"synopsis"`
	SynopsisType       string `json:"synopsis_type"`
	Runtime            string `json:"runtime"`
}

type Movie struct {
	Title              string `json:"title"`
	IMDBMeta           `json:"imdb_meta"`
	RottenTomatoesMeta `json:"rotten_tomatoes_meta"`
}

func NewClient(c *http.Client, u string) *APIClient {
	return &APIClient{c, u}
}

func (a *APIClient) NormalizeAndSend(i []imdb.Movie, r []rottentomatoes.Movie) error {
	movies := a.Normalize(i, r)
	fmt.Println(movies)

	return a.Send(movies)
}

func (a *APIClient) Normalize(i []imdb.Movie, r []rottentomatoes.Movie) []Movie {
	movies := make(map[string]Movie)

	// Add IMDB movies
	for _, movie := range i {
		if _, set := movies[movie.Title]; !set {
			movies[movie.Title] = Movie{
				movie.Title,
				IMDBMeta{
					movie.Genre,
					movie.Rating,
					movie.MovieRating,
				},
				RottenTomatoesMeta{},
			}
		}
	}

	// Add rotten tomatoes movies
	for _, movie := range r {
		if _, set := movies[movie.Title]; !set {
			movies[movie.Title] = Movie{
				movie.Title,
				IMDBMeta{},
				RottenTomatoesMeta{
					movie.TomatoScore,
					movie.PopcornScore,
					movie.TheaterReleaseDate,
					movie.MpaaRating,
					movie.Synopsis,
					movie.SynopsisType,
					movie.Runtime,
				},
			}
		}
	}


	// Parse and return values
	values := make([]Movie, 0, len(movies))
	for _, value := range movies {
		values = append(values, value)
	}

	return values
}

func (a *APIClient) Send(m []Movie) error {
	b, err := json.Marshal(map[string][]Movie{
		"movies": m,
	})
	if err != nil {
		return fmt.Errorf("Failed to encode payload: %v", err)
	}

	_, err = a.httpClient.Post("http://" + a.url + "/movies", "application/json", bytes.NewBuffer(b))

	return err
}

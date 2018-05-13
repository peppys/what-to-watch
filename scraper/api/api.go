package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeppyS/what-to-watch/scraper/imdb"
	"github.com/PeppyS/what-to-watch/scraper/rottentomatoes"
	"github.com/xrash/smetrics"
)

type IMDBMeta struct {
	Genre      string
	MPAARating string
	Score      float64
}

type RottenTomatoesMeta struct {
	TomatoScore        int
	PopcornScore       int
	TheaterReleaseDate string
	MpaaRating         string
	Synopsis           string
	SynopsisType       string
	Runtime            string
}

type Movie struct {
	Title string
	IMDBMeta
	RottenTomatoesMeta
}

const PostEndpoint = "localhost:8080/movies"

func NormalizeAndSend(i []imdb.Movie, r []rottentomatoes.Movie) error {
	movies := Normalize(i, r)

	return Send(movies)
}

func Normalize(i []imdb.Movie, r []rottentomatoes.Movie) []Movie {
	var movies []Movie
	const minimumWagnerFischerDistance = 3

	// Find potential duplicates
	for imdbIndex, imdbMovie := range i {
		for rottenIndex, rottenMovie := range r {
			score := smetrics.WagnerFischer(imdbMovie.Title, rottenMovie.Title, 1, 1, 2)

			if score <= minimumWagnerFischerDistance {
				// Combine movies and add to movie list
				movies = append(movies, Movie{
					imdbMovie.Title,
					IMDBMeta{
						imdbMovie.Genre,
						imdbMovie.Rating,
						imdbMovie.MovieRating,
					},
					RottenTomatoesMeta{
						rottenMovie.TomatoScore,
						rottenMovie.PopcornScore,
						rottenMovie.TheaterReleaseDate,
						rottenMovie.MpaaRating,
						rottenMovie.Synopsis,
						rottenMovie.SynopsisType,
						rottenMovie.Runtime,
					},
				})

				// Remove from pending list to normalize
				i = append(i[:imdbIndex], i[:imdbIndex+1]...)
				r = append(r[:rottenIndex], r[:rottenIndex+1]...)

			}
		}
	}

	// Add remaining IMDB movies
	for _, movie := range i {
		movies = append(movies, Movie{
			movie.Title,
			IMDBMeta{
				movie.Genre,
				movie.Rating,
				movie.MovieRating,
			},
			RottenTomatoesMeta{},
		})
	}

	// Add remaining rotten tomatoes movies
	for _, movie := range r {
		movies = append(movies, Movie{
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
		})
	}

	return movies
}

func Send(m []Movie) error {
	b, err := json.Marshal(map[string][]Movie{
		"movies": m,
	})
	if err != nil {
		return fmt.Errorf("Failed to encode payload: %v", err)
	}

	_, err = http.Post(PostEndpoint, "application/json", bytes.NewBuffer(b))

	return err
}

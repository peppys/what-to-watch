package imdb

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Movie struct {
	Title       string
	Rating      string
	MovieRating string
	Genre       string
}

const (
	moviesInTheatersPage = "https://www.imdb.com/showtimes/location"

	linkSelector = "div.lister-item-image > a[href]:nth-child(1)"

	infoSelector        = "td.overview-top"
	titleSelector       = "td.overview-top h4 a"
	ratingSelector      = "p.cert-runtime-genre > img"
	movieRatingSelector = "span.rating-rating"
	genreSelector       = "p.cert-runtime-genre > span[itemprop=genre]"
)

func Scrape() {
	var movies []Movie

	c := colly.NewCollector()

	// Crawl all movies in theaters
	c.OnHTML(linkSelector, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// Scrape movie info
	c.OnHTML(infoSelector, func(e *colly.HTMLElement) {
		movie := Movie{
			Title:       e.ChildAttr(titleSelector, "title"),
			Rating:      e.ChildAttr(ratingSelector, "title"),
			MovieRating: e.ChildText(movieRatingSelector),
			Genre:       e.ChildText(genreSelector),
		}

		fmt.Println(movie)

		movies = append(movies, movie)
	})

	c.Visit(moviesInTheatersPage)
}

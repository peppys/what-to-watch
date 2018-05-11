package imdb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Movie struct {
	Title       string
	Rating      string
	MovieRating float64
	Genre       string
}

const (
	moviesInTheatersPage = "https://www.imdb.com/showtimes/location"

	linkSelector = "div.lister-item-image > a[href]:nth-child(1)"

	moviePageSelector   = "div.title-overview"
	titleSelector       = "div.title_wrapper h1"
	ratingSelector      = "div.title_wrapper div meta[itemprop=contentRating]"
	movieRatingSelector = "div.imdbRating span[itemprop=ratingValue]"
	genreSelector       = "div.title_wrapper a span"
)

func Scrape() []Movie {
	var movies []Movie

	c := colly.NewCollector()

	// Crawl all movies in theaters
	c.OnHTML(linkSelector, func(e *colly.HTMLElement) {
		e.Request.Visit(strings.Replace(e.Attr("href"), "/showtimes", "", 1))
	})

	// Scrape movie info
	c.OnHTML(moviePageSelector, func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildText(titleSelector))
		rating := strings.TrimSpace(e.ChildAttr(ratingSelector, "content"))
		movieRating, _ := strconv.ParseFloat(strings.TrimSpace(e.ChildText(movieRatingSelector)), 64)
		genre := strings.TrimSpace(e.ChildText(genreSelector))

		movie := Movie{
			title,
			rating,
			movieRating,
			genre,
		}

		fmt.Println(movie)

		movies = append(movies, movie)
	})

	c.Visit(moviesInTheatersPage)

	c.Wait()

	return movies
}

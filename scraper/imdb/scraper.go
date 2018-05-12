package imdb

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	moviesInTheatersPage = "https://www.imdb.com/showtimes/location"

	linkSelector = "div.lister-item-image > a[href]:nth-child(1)"

	moviePageSelector   = "div.title-overview"
	titleSelector       = "div.title_wrapper h1"
	ratingSelector      = "div.title_wrapper div meta[itemprop=contentRating]"
	movieRatingSelector = "div.imdbRating span[itemprop=ratingValue]"
	genreSelector       = "div.title_wrapper a span"
)

type Movie struct {
	Title       string
	Rating      string
	MovieRating float64
	Genre       string
}

type Scraper struct {
	*colly.Collector
}

func NewScraper(c *colly.Collector) *Scraper {
	return &Scraper{c}
}

func (s *Scraper) Scrape() []Movie {
	var movies []Movie

	// Crawl all movies in theaters
	s.OnHTML(linkSelector, func(e *colly.HTMLElement) {
		go e.Request.Visit(strings.Replace(e.Attr("href"), "/showtimes", "", 1))
	})

	// Scrape movie info
	s.OnHTML(moviePageSelector, func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildText(titleSelector))
		rating := strings.TrimSpace(e.ChildAttr(ratingSelector, "content"))
		movieRating, _ := strconv.ParseFloat(strings.TrimSpace(e.ChildText(movieRatingSelector)), 64)
		genre := strings.TrimSpace(e.ChildText(genreSelector))

		movies = append(movies, Movie{
			title,
			rating,
			movieRating,
			genre,
		})
	})

	s.Visit(moviesInTheatersPage)

	s.Wait()

	return movies
}

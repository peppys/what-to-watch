package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PeppyS/what-to-watch/scraper/api"
	"github.com/PeppyS/what-to-watch/scraper/imdb"
	"github.com/PeppyS/what-to-watch/scraper/rottentomatoes"
	"github.com/gocolly/colly"
)

func main() {
	imdbScraper := imdb.NewScraper(colly.NewCollector())
	rottenTomatoesScraper := rottentomatoes.NewScraper(&http.Client{})

	imdbMovies := imdbScraper.Scrape()
	rottenTomatoesMovies, err := rottenTomatoesScraper.Scrape()
	if err != nil {
		log.Fatalln("Problem scraping for rotten tomatoes movies:", err)
	}

	fmt.Println("Successfully scraped movies from IMDB & Rotten Tomatoes")

	err = api.NormalizeAndSend(imdbMovies, rottenTomatoesMovies)
	if err != nil {
		log.Fatalln("Problem normalzing and sending movies to API:", err)
	}

	fmt.Println("Successfully finished scraping and sending movies to API")
}

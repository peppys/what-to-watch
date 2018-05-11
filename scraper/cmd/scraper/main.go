package main

import (
	"fmt"
	"github.com/PeppyS/what-to-watch/scraper/imdb"
)

func main() {
	movies := imdb.Scrape()

	fmt.Printf("Successfully scraped movies: %v", movies)
}

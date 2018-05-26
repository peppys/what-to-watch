package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeppyS/what-to-watch/proto"
)

const (
	movieIndex   = "movies"
	movieMapping = "movie"
)

type ElasticsearchClient struct {
	httpClient *http.Client
	url        string
}

func NewElasticsearchClient(c *http.Client, url string) *ElasticsearchClient {
	return &ElasticsearchClient{c, url}
}

func (c *ElasticsearchClient) BulkPostMovies(movies []*proto.PostMoviesPayload_Movie) error {
	for _, movie := range movies {
		b, err := json.Marshal(movie)
		if err != nil {
			return fmt.Errorf("Failed to encode payload: %v", err)
		}

		_, err = c.httpClient.Post(
			fmt.Sprintf("http://%s/%s/%s", c.url, movieIndex, movieMapping),
			"application/json",
			bytes.NewBuffer(b),
		)
		if err != nil {
			return fmt.Errorf("Failed sending request: %v", err)
		}
	}

	return nil
}

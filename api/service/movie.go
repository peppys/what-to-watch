package service

import (
	"github.com/PeppyS/what-to-watch/proto"
)

type elasticSearchClient interface {
	BulkPostMovies(movies []*proto.PostMoviesPayload_Movie) error
}

// MovieService defines service structure
type MovieService struct {
	esClient elasticSearchClient
}

// NewMovie instantiates MovieService
func NewMovie(esClient elasticSearchClient) *MovieService {
	return &MovieService{esClient}
}

func (s *MovieService) AddBulk(movies []*proto.PostMoviesPayload_Movie) error {
	return s.esClient.BulkPostMovies(movies)
}

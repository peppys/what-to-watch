package service

import (
	"github.com/PeppyS/what-to-watch/proto"
)

type elasticSearchClient interface {
	BulkIndexMovies(movies []*proto.MoviesList_Movie) error
}

// MovieService defines service structure
type MovieService struct {
	esClient elasticSearchClient
}

// NewMovie instantiates MovieService
func NewMovie(esClient elasticSearchClient) *MovieService {
	return &MovieService{esClient}
}

func (s *MovieService) BulkIndex(movies []*proto.MoviesList_Movie) error {
	return s.esClient.BulkIndexMovies(movies)
}

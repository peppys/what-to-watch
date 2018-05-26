package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/PeppyS/what-to-watch/proto"
	google_proto_empty "github.com/golang/protobuf/ptypes/empty"
)

type movieService interface {
	BulkIndex(movies []*proto.MoviesList_Movie) error
}

// MovieController defines controller structure
type MovieController struct {
	service movieService
}

// NewMovie instantiates MovieController, handling movie CRUD operations
func NewMovie(s movieService) *MovieController {
	return &MovieController{
		service: s,
	}
}

// BulkIndex adds movies to ES index
func (c *MovieController) BulkIndex(ctx context.Context, payload *proto.MoviesList) (*proto.PostMoviesResponse, error) {
	err := c.service.BulkIndex(payload.Movies)
	if err != nil {
		return nil, fmt.Errorf("Error bulk posting %v", err)
	}

	log.Println("Successfully uploaded", len(payload.Movies), "movies")

	return &proto.PostMoviesResponse{
		Success: true,
	}, nil
}

func (c *MovieController) Get(ctx context.Context, empty *google_proto_empty.Empty) (*proto.MoviesList, error) {
	return &proto.MoviesList{}, nil
}

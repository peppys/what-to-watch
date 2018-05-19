package controller

import (
	"context"
	"fmt"

	pb "github.com/PeppyS/what-to-watch/proto"
)

type movieService interface {
	Get() error
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

// Post adds movies to ES index
func (c *MovieController) Post(ctx context.Context, movies *pb.PostMoviesPayload) (*pb.PostMoviesResponse, error) {
	err := c.service.Get()
	if err != nil {
		return nil, err
	}

	fmt.Println("Movies:", movies)

	return &pb.PostMoviesResponse{
		Success: true,
	}, nil
}

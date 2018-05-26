package controller

import (
	"fmt"
	"log"
	"context"

	pb "github.com/PeppyS/what-to-watch/proto"
)

type movieService interface {
	AddBulk(movies []*pb.PostMoviesPayload_Movie) error
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
func (c *MovieController) Post(ctx context.Context, payload *pb.PostMoviesPayload) (*pb.PostMoviesResponse, error) {
	err := c.service.AddBulk(payload.Movies)
	if err != nil {
		return nil, fmt.Errorf("Error bulk posting %v", err)
	}

	log.Println("Successfully uploaded", len(payload.Movies), "movies")

	return &pb.PostMoviesResponse{
		Success: true,
	}, nil
}

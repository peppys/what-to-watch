package controller

import (
	"github.com/PeppyS/what-to-watch/proto"
	google_proto_empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

// HealthController - TODO
type HealthController struct {
}

// NewHealth - TODO
func NewHealth() *HealthController {
	return &HealthController{}
}

// Check - TODO
func (c *HealthController) Check(ctx context.Context, empty *google_proto_empty.Empty) (*proto.HealthResponse, error) {
	return &proto.HealthResponse{
		Status: "OK",
	}, nil
}

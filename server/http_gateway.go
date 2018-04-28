package server

import (
	"context"
	"net/http"

	pb "github.com/PeppyS/personal-site-api/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// ListenAndServeHTTPGateway TODO
func ListenAndServeHTTPGateway(grpcAddress, httpAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterResumeServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	err = pb.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpAddress, mux)
}

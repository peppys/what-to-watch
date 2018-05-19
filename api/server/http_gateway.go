package server

import (
	"context"
	"net/http"
	"log"

	pb "github.com/PeppyS/what-to-watch/proto"
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
	err := pb.RegisterMovieServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	err = pb.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	log.Println("Listening on", httpAddress)
	return http.ListenAndServe(httpAddress, mux)
}

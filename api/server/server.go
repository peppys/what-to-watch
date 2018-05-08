package server

import (
	"net"

	"github.com/PeppyS/what-to-watch/api/middleware"
	pb "github.com/PeppyS/what-to-watch/api/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// ListenAndServe - TODO
func ListenAndServe(address string, resumeAPI pb.ResumeServiceServer, healthAPI pb.HealthServiceServer) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc.UnaryServerInterceptor(middleware.Authentication),
			grpc.UnaryServerInterceptor(middleware.Logging),
		)),
	)
	pb.RegisterHealthServiceServer(server, healthAPI)
	pb.RegisterResumeServiceServer(server, resumeAPI)

	return server.Serve(listen)
}

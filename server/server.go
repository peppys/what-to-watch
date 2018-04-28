package server

import (
	"net"

	pb "github.com/PeppyS/personal-site-api/proto"
	"google.golang.org/grpc"
)

// ListenAndServe - TODO
func ListenAndServe(address string, resumeAPI pb.ResumeServiceServer, healthAPI pb.HealthServiceServer) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterHealthServiceServer(server, healthAPI)
	pb.RegisterResumeServiceServer(server, resumeAPI)
	
	return server.Serve(listen)
}

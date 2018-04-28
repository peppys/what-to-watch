package server

import (
	"net"

	pb "github.com/PeppyS/personal-site-api/proto"
	"google.golang.org/grpc"
)

// ListenAndServe - TODO
func ListenAndServe(address string, resumeAPI pb.ResumeServiceServer) error {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterResumeServiceServer(server, resumeAPI)
	
	return server.Serve(listen)
}

package main

import (
	"context"
	"net"
	"serverstatus/grpc/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnsafeServerStatusServer
}

func (s *Server) GetServerStatus(ctx context.Context, in *pb.Empty) (*pb.ServerStatusResponse, error) {
	return &pb.ServerStatusResponse{
		DiskUsage: &pb.DiskUsage{
			UsedSpace:  0,
			TotalSpace: 1,
		},
		CpuUsage: &pb.CPUUsage{
			UsagePercentage: 0,
		},
		MemoryUsage: &pb.MemoryUsage{
			UsedMemory:  0,
			TotalMemory: 0,
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	println("Server started")

	grpcServer := grpc.NewServer()
	pb.RegisterServerStatusServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

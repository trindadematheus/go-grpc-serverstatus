package main

import (
	"serverstatus/grpc/pb"
)

type Server struct {
	pb.UnsafeServerStatusServer
}

// func (s *server) GetServerStatus(ctx context.Context, in *pb.Empty) (*pb.ServerStatusResponse, error) {
// 	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
// }

func main() {
	println("Teste")
}

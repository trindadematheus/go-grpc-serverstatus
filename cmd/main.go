package main

import (
	"context"
	"net"
	"serverstatus/grpc/pb"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnsafeServerStatusServer
}

func (s *Server) GetServerStatus(ctx context.Context, in *pb.Empty) (*pb.ServerStatusResponse, error) {
	cpuPercentage, err := cpu.Percent(time.Second, false)
	diskUsage, err := disk.Usage("/")
	memUsage, err := mem.VirtualMemory()

	if err != nil {
		panic(err)
	}

	return &pb.ServerStatusResponse{
		DiskUsage: &pb.DiskUsage{
			UsedSpace:  float32(diskUsage.Used),
			TotalSpace: float32(diskUsage.Total),
		},
		CpuUsage: &pb.CPUUsage{
			UsagePercentage: float32(cpuPercentage[0]),
		},
		MemoryUsage: &pb.MemoryUsage{
			UsedMemory:  float32(memUsage.Used),
			TotalMemory: float32(memUsage.Total),
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

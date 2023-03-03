package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/Xacor/go-sysmon/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile = flag.String("cert_file", "", "The TLS cert file")
	// keyFile  = flag.String("key_file", "", "The TLS key file")
	port = flag.Int("port", 5000, "The server port")
)

type sysMonServer struct {
	pb.UnimplementedSysMonServer

	mu       sync.Mutex
	snapshot []pb.Snapshot
}

func (s *sysMonServer) GetSnapshot(req *pb.Request, stream pb.SysMon_GetSnapshotServer) error {
	log.Println("GetSnapshot:", req.RefreshRate.AsDuration(), req.RefreshInterval.AsDuration())
	ticker := time.NewTicker(req.RefreshRate.AsDuration())

	for range ticker.C {
		// Calculate snapshot and send it here
		err := stream.Send(&pb.Snapshot{
			LoadAverage: &pb.LoadAverage{Load1: 1.5, Load5: 0.9, Load15: 0.9},
			TimeCreated: timestamppb.New(time.Now()),
		})
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSysMonServer(grpcServer, &sysMonServer{})
	log.Println("Listening...")
	grpcServer.Serve(lis)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/Xacor/go-sysmon/sysmon"
	"google.golang.org/grpc"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 5000, "The server port")
)

type sysMonServer struct {
	pb.UnimplementedSysMonServer
	mu   sync.Mutex
	stat pb.Statistic
}

func (s *sysMonServer) GetStatistic(ctx context.Context, req *pb.Request) (*pb.Statistic, error) {
	log.Println("Received GetStatistic:", req.RefreshInterval, req.RefreshRate)
	return &pb.Statistic{SystemLoad: 2}, nil
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

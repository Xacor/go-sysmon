package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Xacor/go-sysmon/proto"
	"github.com/Xacor/go-sysmon/server/monitoring"
	"google.golang.org/grpc"
)

var (
	// tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile = flag.String("cert_file", "", "The TLS cert file")
	// keyFile  = flag.String("key_file", "", "The TLS key file")
	port = flag.Int("port", 5000, "The server port")
)

type sysMonServer struct {
	pb.UnimplementedSysMonServer
}

func (s *sysMonServer) GetSnapshot(req *pb.Request, stream pb.SysMon_GetSnapshotServer) error {
	log.Println("GetSnapshot:", req.RefreshRate.AsDuration(), req.RefreshInterval.AsDuration())
	ticker := time.NewTicker(req.RefreshRate.AsDuration())

	for range ticker.C {
		// Calculate snapshot and send it here

		snapshot, err := s.CreateSnapshot()
		if err != nil {
			return err
		}

		err = stream.Send(snapshot)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *sysMonServer) CreateSnapshot() (*pb.Snapshot, error) {

	avg, err := monitoring.LoadAvg("/proc/loadavg")
	if err != nil {
		return nil, fmt.Errorf("CreateSnapshot: %w", err)
	}

	avgpb, err := monitoring.MarshallLoadAvg(avg)
	if err != nil {
		return nil, fmt.Errorf("CreateSnapshot: %w", err)
	}

	snapshot := pb.Snapshot{LoadAverage: avgpb}
	return &snapshot, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSysMonServer(grpcServer, &sysMonServer{})

	log.Println("Listening...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

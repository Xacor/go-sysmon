package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
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

	mu        sync.Mutex
	snapshots []*pb.Snapshot
}

func (s *sysMonServer) GetSnapshot(req *pb.Request, stream pb.SysMon_GetSnapshotServer) error {
	log.Println("GetSnapshot:", req.RefreshRate.AsDuration(), req.RefreshInterval.AsDuration())
	ticker := time.NewTicker(req.RefreshRate.AsDuration())

	for range ticker.C {
		// Calculate snapshot and send it here

		s.mu.Lock()

		last := s.snapshots[len(s.snapshots)-1]

		s.mu.Unlock()

		err := stream.Send(last)

		if err != nil {
			return err
		}
	}
	return nil

}

func (s *sysMonServer) Run() {
	jobs := []job{
		job(LoadAvg),
		job(ProcStat),
	}

	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		//log.Println("Ticker: ", t)
		c := StartPipeline(jobs...)

		snapshot := <-c
		s.mu.Lock()

		log.Println("Appending snapshot:", snapshot)
		s.snapshots = append(s.snapshots, snapshot)

		s.mu.Unlock()

	}
}

type job func(out chan<- interface{})

func StartPipeline(jobs ...job) chan *pb.Snapshot {
	var wg sync.WaitGroup

	resIn := make(chan interface{}, len(jobs))
	resOut := make(chan *pb.Snapshot)

	go CombineSnapshot(resIn, resOut)

	wg.Add(len(jobs))
	for _, j := range jobs {

		go func(j job, out chan interface{}) {
			defer wg.Done()
			j(resIn)
		}(j, resIn)
	}

	go func() {
		wg.Wait()
		close(resIn)
	}()

	return resOut
}

func LoadAvg(out chan<- interface{}) {
	la, err := monitoring.GetLoadAvg("/proc/loadavg")
	if err != nil {
		log.Println(fmt.Errorf("Error LoadAvg: %w", err))
	}
	out <- *la
}

func ProcStat(out chan<- interface{}) {
	ps, err := monitoring.GetProcStat("/proc/stat")
	if err != nil {
		log.Println(fmt.Errorf("Error ProcStat: %w", err))
	}
	out <- *ps
}

func CombineSnapshot(in <-chan interface{}, out chan<- *pb.Snapshot) {
	snapshot := pb.Snapshot{}
	for field := range in {
		// log.Println("Combine result: ", field)
		switch t := field.(type) {

		case pb.LoadAverage:
			snapshot.LoadAverage = &t

		case pb.ProcStat:
			snapshot.ProcStat = &t

		default:
			log.Println(fmt.Errorf("combineSnapshot: Unknown type %v", t))
		}

	}
	// log.Println("New snap:", &snapshot)
	out <- &snapshot
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	server := sysMonServer{}
	go server.Run()

	pb.RegisterSysMonServer(grpcServer, &server)
	log.Println("Listening...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

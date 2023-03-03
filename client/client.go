package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/Xacor/go-sysmon/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	serverAddr = flag.String("addr", "localhost:5000", "The server address in the format of host:port")
)

func printSnaphots(client pb.SysMonClient, req *pb.Request) {
	log.Println("Getting snapshot...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	stream, err := client.GetSnapshot(ctx, req)
	if err != nil {
		log.Fatalf("client.GetSnapshot failed: %v", err)
	}
	for {
		snapshot, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.GetSnaphot failed: %v", err)
		}
		log.Printf("Received snapshot: %v", snapshot)
	}

}

func main() {
	//var opts []grpc.DialOption
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewSysMonClient(conn)

	printSnaphots(client, &pb.Request{RefreshRate: durationpb.New(time.Second * 3), RefreshInterval: durationpb.New(time.Second * 5)})
}

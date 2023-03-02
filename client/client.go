package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Xacor/go-sysmon/sysmon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	serverAddr = flag.String("addr", "localhost:5000", "The server address in the format of host:port")
)

func printStat(client pb.SysMonClient, req *pb.Request) {
	log.Println("Gettting system load...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stat, err := client.GetStatistic(ctx, req)
	if err != nil {
		log.Fatalf("client.etStatistic failed: %v", err)
	}
	log.Println(stat)

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

	printStat(client, &pb.Request{RefreshRate: durationpb.New(time.Second * 5), RefreshInterval: durationpb.New(time.Second * 10)})
}

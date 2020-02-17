package server

import (
	"flag"
	"fmt"
	"net"
	"os"

	bidirectional "grpc-bidirectional-stream/pkg"
	pb "grpc-bidirectional-stream/pkg/proto"

	"google.golang.org/grpc"
)

var (
	port int
)

func initializeFlags() {
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
}

func main() {
	initializeFlags()

	// start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	var s bidirectional.SampleServer
	pb.RegisterSampleServer(grpcServer, &s)

	defer grpcServer.GracefulStop()
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}
}

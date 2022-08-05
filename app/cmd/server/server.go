package main;

import ( 
	"net"
	"log"
	"google.golang.org/grpc"
)

func main () {
	listener, failure := net.Listen("tcp", "0.0.0.0:5555")

	if failure != nil {
		log.Fatalf("Could not connect: %v", failure)
	}

	grpcServer := grpc.NewServer()

	failure = grpcServer.Serve(listener)

	if failure != nil {
		log.Fatalf("Could not serve: %v", failure)
	}
}
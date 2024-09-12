package main

import "log"

func main() {
	// run grpcServer
	gRPCServer := NewGRPCServer(":9000")
	if err := gRPCServer.Run(); err != nil {
		log.Fatalf("error running grpc server : %v", err)
	}
}
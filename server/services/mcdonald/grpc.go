package main

import (
	"log"
	"net"

	"github.com/HenryGunadi/grpcs/services/mcdonald/handler"
	orders "github.com/HenryGunadi/grpcs/services/mcdonald/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{
		addr: addr,
	}
}

func (s *GRPCServer) Run() error {
	listener , err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	// register order service
	orderService := orders.NewOrdersService()
	handler.NewGRPCHandler(grpcServer, orderService)

	log.Println("grpc server running on", s.addr)

	return grpcServer.Serve(listener)
}
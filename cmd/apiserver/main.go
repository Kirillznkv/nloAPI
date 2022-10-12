package main

import (
	pb "github.com/Kirillznkv/nloAPI/pkg" //generated proto file
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := &GRPCServer{}
	pb.RegisterNloServer(grpcServer, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

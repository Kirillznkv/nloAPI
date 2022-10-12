package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Kirillznkv/nloAPI/pkg/api" //generated proto file
)

type GRPCServer struct {
} //del

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

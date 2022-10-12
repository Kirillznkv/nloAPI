package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/Kirillznkv/nloAPI/internal/app/apiserver"
	pb "github.com/Kirillznkv/nloAPI/pkg/api"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	grpcServer := grpc.NewServer()
	srv := &apiserver.ImplementedNloServer{}
	pb.RegisterNloServer(grpcServer, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

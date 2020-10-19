package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RohitNalePUCSD/Go-Pratics/gRPC/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

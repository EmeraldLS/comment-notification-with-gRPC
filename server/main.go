package main

import (
	"log"
	"net"

	pb "github.com/EmeraldLS/notification-with-grpc/proto"
	"github.com/EmeraldLS/notification-with-grpc/server/cmd"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":3031")
	if err != nil {
		log.Fatalf("An error occured while trying to create a connectio listener :: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, &cmd.Notification{})
	log.Printf("gRPC Server started @ %v", list.Addr())
	if err := s.Serve(list); err != nil {
		log.Fatalf("An error occured starting the gRPC server :: %v", err)
	}
}

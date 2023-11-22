package main

import (
	"context"
	"log"

	"github.com/EmeraldLS/notification-with-grpc/client/pkg"
	pb "github.com/EmeraldLS/notification-with-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":3031", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Unable to establish a connection to the server :: %v", err)
	}

	client := pb.NewNotificationClient(conn)
	log.Println("Connected to gRPC server")
	if err := pkg.Notify(client, context.TODO()); err != nil {
		log.Printf("An error occured while sending notification :: %v", err)
	}
}

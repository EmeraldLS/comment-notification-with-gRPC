package cmd

import (
	"fmt"
	"log"

	pb "github.com/EmeraldLS/notification-with-grpc/proto"
)

type Notification struct {
	pb.NotificationServer
}

// Notify implements proto.NotificationServer.
func (*Notification) Notify(stream pb.Notification_NotifyServer) error {
	ctx := stream.Context()
	var total_comment = 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:

		}

		comment, err := stream.Recv()
		if err != nil {
			return err
		}
		total_comment++
		log.Println("-----------")
		log.Printf("Recieved comment: '%v' from '%v'", comment.GetContent(), comment.User.GetName())
		log.Printf("User is following %v and has %v followers", comment.User.GetFollowing(), comment.User.GetFollowers())

		resp := &pb.Response{
			Content:      fmt.Sprintf("Hi %v commented on your photo", comment.User.GetName()),
			TotalComment: int32(total_comment),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}

	}

}

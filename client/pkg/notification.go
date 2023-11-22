package pkg

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/EmeraldLS/notification-with-grpc/proto"
)

/*
I'll establish 3 goroutines
1st goroutine will be for sending notifications
2nd goroutine will be for receiving responses from the server
3rd goroutine would be a channel for closing the stream when it's done
*/
func Notify(client pb.NotificationClient, ctx context.Context) error {

	stream, err := client.Notify(ctx)
	var done chan bool
	var errChan = make(chan error, 1)
	if err != nil {
		return err
	}
	// For sending comments
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:

				h := md5.New()

				comment1 := &pb.Comment{
					User: &pb.User{

						Name:      "Sanni Abdullah",
						Followers: 533,
						Following: 78,
					},
					Content:   "You're inspiring fr",
					Timestamp: time.Now().String(),
				}
				io.WriteString(h, comment1.User.GetName()+comment1.GetContent())
				ID := fmt.Sprintf("%x", h.Sum(nil))
				comment1.User.Id = ID

				comment2 := &pb.Comment{
					User: &pb.User{

						Name:      "Oluwasegun Lawrence",
						Followers: 290,
						Following: 1,
					},
					Content: "I'm gonna be great because I believe in myself",

					Timestamp: time.Now().Add(time.Duration(time.Second * 5)).String(),
				}
				h2 := md5.New()
				io.WriteString(h2, comment2.User.GetName()+comment2.GetContent())
				comment2.User.Id = fmt.Sprintf("%x", h2.Sum(nil))
				if err := stream.Send(comment1); err != nil {
					log.Printf("An error occured while trying to send comment :: %v", err)
					errChan <- err
					return
				}

				if err := stream.Send(comment2); err != nil {
					log.Printf("An error occured while trying to send comment :: %v", err)
					errChan <- err
					return
				}

				time.Sleep(time.Second * 60)

			}
		}
	}()

	// For receiving notifications

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("No more comments")
				errChan <- err
				return
			}
			if err != nil {
				log.Printf("An error occured while trying to recieved notifications :: %v", err)
				errChan <- err
				return
			}

			log.Printf("Notification: '%v'", resp.GetContent())
			log.Printf("Total comment: %v", resp.GetTotalComment())
		}
	}()

	// To close channel when syreaming is done

	go func() {
		<-stream.Context().Done()

		if err := stream.CloseSend(); err != nil {
			log.Panicln(err)
			errChan <- err

		}
	}()

	return <-errChan
}

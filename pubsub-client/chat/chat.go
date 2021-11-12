package chat

import (
	"fmt"
	"log"

	//"net/http"
	"cloud.google.com/go/pubsub"
	//"net/http"

	//"log"
	"golang.org/x/net/context"
)

var (
	topic *pubsub.Topic
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	//implementar pubsub
	fmt.Println("Received from grpc client: %s", message.Body)
	//codigo de pubsub aqui :)
	//ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "alpine-aspect-325702")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	topicName := "sopes-rpc"
	topic = client.Topic(topicName)

	msg := &pubsub.Message{
		Data: []byte(message.Body),
	}

	aaaa, errz := topic.Publish(ctx, msg).Get(ctx)

	if  errz != nil {
		fmt.Println(errz)
	}
	fmt.Println(aaaa)
	//fin pubsub


	return &Message{Body: message.Body}, nil
}

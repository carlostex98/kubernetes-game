package main

import (
	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"pubsub-client/chat"
	"golang.org/x/net/context"
)

var (
	topic *pubsub.Topic
)

func main() {
	ctx := context.Background()
	client, errx := pubsub.NewClient(ctx, "alpine-aspect-325702")
	if errx != nil {
		log.Fatal(errx)
	}
	defer client.Close()

	topicName := "sopes-rpc"
	topic = client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, errs := topic.Exists(ctx)
	if errs != nil {
		log.Fatal(errs)
	}
	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, errs = client.CreateTopic(ctx, topicName)
		if errs != nil {
			log.Fatal(errs)
		}
	}


	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}

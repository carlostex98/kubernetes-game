package main

import (
	"cloud.google.com/go/pubsub"
	//"fmt"
	"log"
	"golang.org/x/net/context"
)



func main() {

	ctx := context.Background()
	projID := "alpine-aspect-325702"

	client, err := pubsub.NewClient(ctx, projID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := "sopes-rpc"
	topic := client.Topic(topicName)

	sub, errz := client.CreateSubscription(context.Background(), "sopes-g", pubsub.SubscriptionConfig{Topic: topic})
	errz = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if errz != nil {
		// Handle error.
	}

}



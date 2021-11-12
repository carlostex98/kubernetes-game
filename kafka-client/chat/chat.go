package chat

import (
	"fmt"

	//"log"

	"golang.org/x/net/context"
	"github.com/Shopify/sarama"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	//implementar rabbitmq aqui y tambien el chungo para ganar
	fmt.Println("Received from grpc client: %s", message.Body)

	//kafka aqui
	PushCommentToQueue(
		"kafka-sopes",
		[]byte(message.Body),
	)


	return &Message{Body: "s"}, nil
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {

	brokersUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}



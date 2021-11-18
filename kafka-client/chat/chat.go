package chat

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	//"log"

	"github.com/Shopify/sarama"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	//implementar rabbitmq aqui y tambien el chungo para ganar
	fmt.Println("Received from grpc client: %s", message.Body)

	sx := strings.Split(message.Body, "|")
	x := setWiner(sx[0], sx[1], sx[2], sx[3]) //esto se envia al worker de rabbit mq

	//kafka aqui
	PushCommentToQueue(
		"kafka-sopes",
		[]byte(x),
	)


	return &Message{Body: "s"}, nil
}

func setWiner(name string, players string, game string, corr string) string {
	//count++
	//para el juego
	intVar, _ := strconv.Atoi(players)
	c := name + "|" + players + "|" + game + "|" + corr + "|" + calcGame(game, intVar)
	return c
}

func calcGame(game string, players int) string {
	switch game {
	case "1":
		return game1(players)
	case "2":
		return game2(players)
	case "3":
		return game3(players)
	default:
		return "1"

	}
	//return 1
}

func game1(players int) string {
	pl1 := rand.Intn(players-1) + 1
	pl2 := rand.Intn(players-1) + 1
	pl3 := rand.Intn(players-1) + 1

	mx := [3]int{pl1, pl2, pl3}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
}

func game2(players int) string {

	mx := [5]int{
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
	}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
}

func game3(players int) string {
	mx := [8]int{
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
		rand.Intn(players-1) + 1,
	}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
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

	brokersUrl := []string{"localhost:9092"}
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



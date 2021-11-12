package chat

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	//"log"

	amqp "github.com/streadway/amqp"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	//implementar rabbitmq aqui y tambien el chungo para ganar
	fmt.Println("Received from grpc client: %s", message.Body)

	sx := strings.Split(message.Body, "|")
	x := setWiner(sx[0], sx[1], sx[2], sx[3]) //esto se envia al worker de rabbit mq
	res, _ := winnerRPC(x)
	fmt.Println(res)

	return &Message{Body: x}, nil
}



func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func winnerRPC(m string) (res string, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(m),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			//res, err = strconv.Atoi(string(d.Body))
			res = string(d.Body)
			//failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
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

package main

import (
	"context"
	"fmt"
	amqp "github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const (
	URI = "mongodb://root:example@34.66.118.26:27017/"
)

type Data struct {
	NumberReq int    `bson:"number_req" json:"number_req"`
	Game     string `bson:"game" json:"game"`
	NameGame string `bson:"name_game" json:"name_game"`
	Winner  string `bson:"winner" json:"winner"`
	Players int    `bson:"players" json:"players"`
	Worker  string `bson:"worker" json:"worker"`
}

func main() {
	//mongo connect
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	//rabbit connect
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//implementar mongo aqui
			//n, err := strconv.Atoi(string(d.Body))
			fmt.Println(string(d.Body))
			//failOnError(err, "Failed to convert body to integer")
			al := string(d.Body)
			sx := strings.Split(al, "|")//valores spliteados
			xx := convertToMongo(sx[0], sx[1], sx[2], sx[3], sx[4])
			response := newDataMongo(xx, *client)

			//response := "1"

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.FormatBool(response)),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}


func convertToMongo(gameName string, players string, game string, numReq string, winner string) Data{
	intVar, _ := strconv.Atoi(players)
	x, _ := strconv.Atoi(numReq)
	c := Data{
		NumberReq: x,
		Game:      game,
		NameGame:  gameName,
		Players:   intVar,
		Worker:    "RabbitMQ",
		Winner:    winner,
	}

	return c
}

func newDataMongo(data Data, c mongo.Client) bool {

	col := c.Database("Proyecto2Sopes").Collection("Logs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//dt, _:= toDoc(data)
	_, insertErr := col.InsertOne(ctx, data)
	//fmt.Println(xx)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		defer cancel()
		return false
	} else {
		defer cancel()
		return true
	}
}
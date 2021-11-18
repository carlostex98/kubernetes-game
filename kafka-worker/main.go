package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Shopify/sarama"
	"context"
)

type Data struct {
	NumberReq int    `bson:"number_req" json:"number_req"`
	Game     string `bson:"game" json:"game"`
	NameGame string `bson:"name_game" json:"name_game"`
	Winner  string `bson:"winner" json:"winner"`
	Players int    `bson:"players" json:"players"`
	Worker  string `bson:"worker" json:"worker"`
}

const (
	URI = "mongodb://root:example@35.184.26.14:27017/"
)

func main() {

	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	topic := "kafka-sopes"
	worker, err := connectConsumer([]string{"localhost:29092"})
	if err != nil {
		panic(err)
	}

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	//msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				//insertar mongo
				al := string(msg.Value)
				sx := strings.Split(al, "|")//valores spliteados
				xx := convertToMongo(sx[0], sx[1], sx[2], sx[3], sx[4])
				newDataMongo(xx, *client)

				//fin mongo :P
				//fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	//fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

}


func convertToMongo(gameName string, players string, game string, numReq string, winner string) Data{
	intVar, _ := strconv.Atoi(players)
	x, _ := strconv.Atoi(numReq)
	c := Data{
		NumberReq: x,
		Game:      game,
		NameGame:  gameName,
		Players:   intVar,
		Worker:    "Kafka",
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

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

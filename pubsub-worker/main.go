package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"fmt"
	"log"

	"golang.org/x/net/context"
)

const (
	URI = "mongodb://root:example@35.184.26.14:27017/"
)

type Data struct {
	NumberReq int    `bson:"number_req" json:"number_req"`
	Game      string `bson:"game" json:"game"`
	NameGame  string `bson:"name_game" json:"name_game"`
	Winner    string `bson:"winner" json:"winner"`
	Players   int    `bson:"players" json:"players"`
	Worker    string `bson:"worker" json:"worker"`
}

func main() {
	clientM, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	ctx := context.Background()
	projID := "alpine-aspect-325702"

	client, err := pubsub.NewClient(ctx, projID)
	if err != nil {
		log.Fatal(err)
	}
	//	defer client.Close()

	topicName := "tema-chilero"
	client.Topic(topicName)	

	sub := client.Subscription("tema-chilero-sub")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data))
		al := string(m.Data)
		sx := strings.Split(al, "|") //valores spliteados
		xx := convertToMongo(sx[0], sx[1], sx[2], sx[3], sx[4])
		newDataMongo(xx, *clientM)
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Println(err)
	}

	/*sub, errz := client.CreateSubscription(ctx, "tema-chilero-sub", pubsub.SubscriptionConfig{Topic: topic})
	if errz != nil {
		// Handle error.
		fmt.Println(errz)
	}
	errz = sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		al := string(m.Data)
		sx := strings.Split(al, "|") //valores spliteados
		xx := convertToMongo(sx[0], sx[1], sx[2], sx[3], sx[4])
		newDataMongo(xx, *clientM)
		m.Ack()
	})*/
	

}

func convertToMongo(gameName string, players string, game string, numReq string, winner string) Data {
	intVar, _ := strconv.Atoi(players)
	x, _ := strconv.Atoi(numReq)
	c := Data{
		NumberReq: x,
		Game:      game,
		NameGame:  gameName,
		Players:   intVar,
		Worker:    "PubSub",
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

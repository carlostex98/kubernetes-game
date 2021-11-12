package chat

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

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
	sx := strings.Split(message.Body, "|")
	x := setWiner(sx[0], sx[1], sx[2], sx[3])
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
		Data: []byte(x),
	}

	aaaa, errz := topic.Publish(ctx, msg).Get(ctx)

	if  errz != nil {
		fmt.Println(errz)
	}
	fmt.Println(aaaa)
	//fin pubsub


	return &Message{Body: message.Body}, nil
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
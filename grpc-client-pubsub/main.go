package main

import (
	"grpc-client/chat"
	"strconv"

	"google.golang.org/grpc"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	//"fmt"
	"net/http"
	//"go.mongodb.org/mongo-driver/bson"
)

type Data struct {
	NumberReq int    `bson:"number_req" json:"number_req"`
	Game      string `bson:"game" json:"game"`
	NameGame  string `bson:"name_game" json:"name_game"`
	Winner    string `bson:"winner" json:"winner"`
	Players   int    `bson:"players" json:"players"`
	Worker    string `bson:"worker" json:"worker"`
}

/*const (
	URI = "mongodb://root:example@34.135.37.109:27017/"
)*/
var count = 0

//var client mongo.Client

func main() {
	//fmt.Println("go go power ranger")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": 2 + 2})
	})
	mx := chat.NewChatServiceClient(conn)

	r.GET("/game/:id/gamename/:name/players/:play", func(c *gin.Context) {
		//return the response from the server
		count++
		xy := c.Param("name") + "|" + c.Param("play") + "|" + c.Param("id") + "|" + strconv.Itoa(count)
		message := chat.Message{
			Body: xy,
		}

		response, err := mx.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		//log.Printf("Response from Server: %s", response.Body)

		c.JSON(http.StatusOK, gin.H{"data": response.Body})
	})

	r.Run(":3002")
}

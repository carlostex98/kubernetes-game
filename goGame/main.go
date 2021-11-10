package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"fmt"
	"net/http"
	//"go.mongodb.org/mongo-driver/bson"
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
	URI = "mongodb://root:example@34.135.37.109:27017/"
)
var count = 0

//var client mongo.Client

func main()  {
	//fmt.Println("go go power ranger")
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": 2+2})
	})

	r.GET("/game/:id/gamename/:name/players/:play", func(c *gin.Context) {
		/*ret := gin.H{
			"id": c.Param("id"),
			"gamename": c.Param("name"),
			"player":c.Param("play"),
		}*/

		x := setWiner(c.Param("name"), c.Param("play"), c.Param("id"))
		m:=newDataMongo(x, *client)
		c.JSON(http.StatusOK, gin.H{"sap":m})
	})

	r.Run()
}

func setWiner(name string, players string, game string) Data {
	count++
	//para el juego
	intVar, _ := strconv.Atoi(players)
	c := Data{
		NumberReq: count,
		Game:      game,
		NameGame:  name,
		Players:   intVar,
		Worker:    "local",
		Winner:    calcGame(game, intVar),
	}

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

func game1(players int) string {
	pl1 := rand.Intn(players-1)+1
	pl2 := rand.Intn(players-1)+1
	pl3 := rand.Intn(players-1)+1

	mx := [3]int{pl1, pl2, pl3}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
}

func game2(players int) string {

	mx := [5]int{
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
	}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
}


func game3(players int) string {
	mx := [8]int{
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
		rand.Intn(players-1)+1,
	}

	final := rand.Intn(len(mx))
	return strconv.Itoa(mx[final])
}


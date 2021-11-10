package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type game struct {
	gameId int
	name   string
}

type gameData struct {
	names       []game
	players     int
	rungames    int
	concurrence int
	timeout     string
}

func main() {
	host := "http://localhost:8080"
	argsWithProg := os.Args[1:]
	m := readConsole(argsWithProg)
	fmt.Println(m)
	xx := generateEndpoints(m.names, m.rungames, m.players, host)
	//usaremos un channel para manejar la concurrencia
	executeConcurence(xx, m)
}

func executeConcurence(xx []string, m gameData) {
	curr := 0
	//reqMake := 0
	for {
		c := make(chan int, m.concurrence)
		if curr == m.rungames {
			close(c)
			break
		}

		for i := 0; i < m.concurrence; i++ {
			go doRequest(xx[curr], c, i)
			curr = curr + 1
		}
		temp := 0
		for n := range c {

			if temp == m.concurrence-1 {
				close(c)
			}
			fmt.Println("Subproceso: ", n, " ejecutado de: ", m.rungames)
			temp++
		}

	}
}

func generateEndpoints(names []game, rungames int, players int, host string) []string {
	//generamos los endpoints
	var points []string
	mx := 0
	for i := 0; i < rungames; i++ {
		if mx == len(names) {
			mx = 0
		}
		res := host + "/game/" + strconv.Itoa(names[mx].gameId) + "/gamename/" + names[mx].name + "/players/" + strconv.Itoa(rand.Intn(players-10)+10)
		points = append(points, res)
		mx = mx + 1
	}
	return points
}

func doRequest(url string, c chan int, procId int) {
	resp, err := http.Get(url)
	time.Sleep(10 * time.Millisecond)
	if err != nil {
		fmt.Println(url, err)
	}
	defer resp.Body.Close()

	c <- procId
}

func readGameNamesConsole(cadena string) []game {
	var gsx []game
	s := strings.Split(cadena, "|")
	for i := 0; i < len(s)-1; i = i + 2 {
		intVal, _ := strconv.Atoi(strings.TrimSpace(s[i]))
		x := game{
			gameId: intVal,
			name:   strings.TrimSpace(s[i+1]),
		}
		gsx = append(gsx, x)
	}
	return gsx
}

func readConsole(mensajes []string) gameData {
	//leemos a partir del 1
	var arrayGames []game
	players, rungames, concurrence, timeout := 0, 0, 0, "0m"

	for i := 1; i < len(mensajes)-1; i = i + 2 {
		//just do it
		t := strings.Replace(mensajes[i], "-", "", -1)
		if t == "gamename" {
			arrayGames = readGameNamesConsole(mensajes[i+1])
		} else if t == "players" {
			intVal, _ := strconv.Atoi(strings.TrimSpace(mensajes[i+1]))
			players = intVal
		} else if t == "rungames" {
			intVal, _ := strconv.Atoi(strings.TrimSpace(mensajes[i+1]))
			rungames = intVal
		} else if t == "concurrence" {
			intVal, _ := strconv.Atoi(strings.TrimSpace(mensajes[i+1]))
			concurrence = intVal
		} else if t == "timeout" {
			timeout = strings.TrimSpace(mensajes[i+1])
		}
	}
	x := gameData{
		names:       arrayGames,
		players:     players,
		rungames:    rungames,
		concurrence: concurrence,
		timeout:     timeout,
	}

	return x
}

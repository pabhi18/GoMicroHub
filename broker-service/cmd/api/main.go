package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	rabbitConn, err := connect()
	if err != nil {
		panic(err)
	}
	app := Config{
		Rabbit: rabbitConn,
	}
	log.Printf("Starting broker service  the on port %s\n", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't connect until rebbit is not ready

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			fmt.Println("Rabbit is not ready yet")
			counts++
		} else {
			fmt.Println("Rabbit connection established \n", c.Properties)
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second

		log.Printf("backing of %d\n", backOff)
		continue

	}

	return connection, nil
}

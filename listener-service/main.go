package main

import (
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println(" Starting listener service...")

	// connect to rabbitmq

	rabbitConn, err := connect()
	if err != nil {
		fmt.Println("error during connect rabbit-mq ", err.Error())
		panic(err)
	}
	fmt.Println("Rabbit connection established", rabbitConn.Properties)

	// starting listening for msg : "sender"

	// create consumer

	// watch queue and events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't connect until rebbit is not ready

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			fmt.Println("Rabbit is not ready yet")
			counts++
		} else {
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

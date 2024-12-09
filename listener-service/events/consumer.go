package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type payload struct {
	Name string `json:"name"`
	Data string `json: "data"`
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil

}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)

}

func (consumer *Consumer) Listen(topics []string) error {
	chh, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	defer chh.Close()

	queue, err := declareQueue(chh)
	if err != nil {
		return err
	}

	for _, s := range topics {
		chh.QueueBind(
			queue.Name,
			s,
			"log_topic",
			false,
			nil,
		)

	}

	messages, err := chh.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	var forever = make(chan bool)

	for d := range messages {
		var payloadData payload
		_ = json.Unmarshal(d.Body, &payloadData)
		go handlePayload(payloadData)
	}

	fmt.Printf("Waiting for msg [Exchange, Queue] [logs_topic, %s] \n", queue.Name)
	<-forever
	return nil
}

func handlePayload(data payload) error {
	switch data.Name {
	case "log", "event":
		// log whatever we get
		err := logEvent(data)
		if err != nil {
			return err
		}
	case "auth":

	default:
		err := logEvent(data)
		if err != nil {
			return err
		}

	}
	return nil
}

func logEvent(entry payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	loggerRequestUrl := "http://logger-service:8082/log"

	request, err := http.NewRequest("POST", loggerRequestUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return errors.New("status code is not got correct during logging event ")
	}

	return nil
}

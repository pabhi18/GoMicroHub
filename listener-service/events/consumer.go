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
		return fmt.Errorf("failed to declare channel : %w", err)
	}
	err = declareExchange(channel)
	if err != nil {
		return fmt.Errorf("failed to declare exchange : %w", err)
	}

	return nil
}

func (consumer *Consumer) Listen(topics []string) error {
	chh, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	defer chh.Close()

	queue, err := declareQueue(chh)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	for _, topic := range topics {
		err := chh.QueueBind(
			queue.Name,
			topic,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to bind queue: %w", err)
		}
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
		return fmt.Errorf("failed to create a consume: %w", err)
	}

	var forever = make(chan bool)

	// Process messages
	for d := range messages {
		var payloadData payload
		err := json.Unmarshal(d.Body, &payloadData)
		if err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}
		go handlePayload(payloadData)
	}

	fmt.Printf("Waiting for messages [Exchange: logs_topic, Queue: %s] \n", queue.Name)
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

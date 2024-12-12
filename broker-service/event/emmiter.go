package events

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	Connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to declare channel : %w", err)
	}

	defer channel.Close()
	err = declareExchange(channel)
	if err != nil {
		return fmt.Errorf("failed to declare exchange : %w", err)
	}
	return nil
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	fmt.Println("pushing to the channel")

	err = channel.Publish(
		"log_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish event : %w", err)
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		Connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}

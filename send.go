package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func send(message string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"flight",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to create queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")

}

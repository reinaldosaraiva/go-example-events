package main

import (
	"log"

	"github.com/reinaldosaraiva/go-example-events/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out, "orders")
	for msg := range out {
		log.Printf("Received message with message: %s", msg.Body)
		msg.Ack(false)
	}
	
}

package main

import (
	"log"

	"github.com/reinaldosaraiva/go-example-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "Hellow World!", "amq.direct")}
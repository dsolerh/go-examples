package main

import (
	"context"
	"fmt"
	"log"
	rabbit "testings/rabitmq"
	"time"
)

func main() {
	// Connect to RabbitMQ server
	queue, err := rabbit.NewRabbitQueue("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer queue.Close()

	// Declare a topic exchange
	exchangeName := "logs_topic"

	// Publish messages with different routing keys
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var msg string
	type Msg struct {
		Msg string
	}
	for {
		_, err := fmt.Scan(&msg)
		if err != nil {
			log.Fatal(err)
		}

		err = queue.PublishWithContext(ctx, exchangeName, "test.some", &Msg{Msg: msg})
		if err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"log"
	rabbit "testings/rabitmq"
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

	trackingQueueName, err := queue.CreateQueue(exchangeName, "test_queue_consumer", "test.*", true)
	if err != nil {
		log.Fatal(err)
	}

	err = queue.ConsumeQueueRoute(trackingQueueName, func(route string, body []byte) {
		log.Printf("route: %s, body: %s\n", route, body)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Block and process messages
	forever := make(chan bool)

	log.Println("Consumer started. Press CTRL+C to exit.")
	<-forever
}

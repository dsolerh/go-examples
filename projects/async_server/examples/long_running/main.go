package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const queueName string = "jobQueue"
const hostString string = "0.0.0.0:8000"

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	jobServer := getServer(queueName)

	// Cleanup resources
	defer jobServer.Channel.Close()
	defer jobServer.Conn.Close()

	// Start Workers
	go func(conn *amqp091.Connection) {
		workerProcess := Workers{
			conn: conn,
		}
		workerProcess.run()
	}(jobServer.Conn)

	router := http.NewServeMux()
	// Attach handlers
	router.HandleFunc("/job/database", jobServer.asyncDBHandler)
	router.HandleFunc("/job/mail", jobServer.asyncMailHandler)
	router.HandleFunc("/job/callback", jobServer.asyncCallbackHandler)

	httpServer := &http.Server{
		Handler:      router,
		Addr:         hostString,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Run HTTP server
	log.Fatal(httpServer.ListenAndServe())
}

package main

import (
	"log"
	"net/http"
	"time"
)

const queueName string = "jobQueue"
const hostString string = "0.0.0.0:8000"

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	jobServer := NewServer(queueName)

	// Cleanup resources
	defer jobServer.Channel.Close()
	defer jobServer.Conn.Close()

	// Start Workers
	SpawnWorkers(jobServer.Conn, jobServer.redisClient)

	router := http.NewServeMux()
	// Attach handlers
	router.HandleFunc("/job/database", jobServer.asyncDBHandler)
	router.HandleFunc("/job/mail", jobServer.asyncMailHandler)
	router.HandleFunc("/job/callback", jobServer.asyncCallbackHandler)
	router.HandleFunc("/job/status", jobServer.statusHandler)

	httpServer := &http.Server{
		Handler:      router,
		Addr:         hostString,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Run HTTP server
	log.Fatal(httpServer.ListenAndServe())
}

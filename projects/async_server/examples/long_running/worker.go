package main

import (
	"async_server/models"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type Workers struct {
	conn *amqp091.Connection
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")

	channel, err := w.conn.Channel()
	handleError(err, "Fetching channel failed")
	defer channel.Close()

	jobQueue, err := channel.QueueDeclare(
		queueName, // Name of the queue
		false,     // Message is persisted or not
		false,     // Delete message when unused
		false,     // Exclusive
		false,     // No Waiting time
		nil,       // Extra args
	)
	handleError(err, "Job queue fetch failed")

	messages, err := channel.Consume(
		jobQueue.Name, // queue
		"",            // consumer
		true,          // auto-acknowledge
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)

	go func() {
		for message := range messages {
			log.Printf("message: %+v", message)

			switch message.Type {
			case "Log":
				job := models.Job[models.Log]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Log'")
				w.dbWork(job)
			case "Callback":
				job := models.Job[models.CallBack]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Callback'")
				w.callbackWork(job)
			case "Mail":
				job := models.Job[models.Mail]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Mail'")
				w.emailWork(job)
			default:
				log.Println("Invalid job type", message.Type)
			}
		}
	}()
	defer w.conn.Close()
	wait := make(chan bool)
	<-wait // Run long-running worker
}

func (w *Workers) dbWork(job models.Job[models.Log]) {
	result := job.ExtraData
	log.Printf("Worker %s: extracting data..., JOB: %+v", job.Type, result)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: saving data to database..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) callbackWork(job models.Job[models.CallBack]) {
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to the given callback..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) emailWork(job models.Job[models.Mail]) {
	log.Printf("Worker %s: sending the email..., JOB: %s", job.Type, job.ID)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
}

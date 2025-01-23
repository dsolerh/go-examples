package main

import (
	"async_server/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Workers struct {
	conn        *amqp091.Connection
	redisClient *redis.Client
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
			// log.Printf("message: %+v", message)
			ctx := context.TODO()

			switch message.Type {
			case "Log":
				job := models.Job[models.Log]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Log'")
				w.dbWork(ctx, job)
			case "Callback":
				job := models.Job[models.CallBack]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Callback'")
				w.callbackWork(ctx, job)
			case "Mail":
				job := models.Job[models.Mail]{}
				err = json.Unmarshal(message.Body, &job)
				handleError(err, "error parsing the job 'Mail'")
				w.emailWork(ctx, job)
			default:
				log.Println("Invalid job type", message.Type)
			}
		}
	}()
	defer w.conn.Close()
	wait := make(chan bool)
	<-wait // Run long-running worker
}

func (w *Workers) dbWork(ctx context.Context, job models.Job[models.Log]) {
	result := job.ExtraData
	w.redisClient.Set(ctx, job.ID.String(), "STARTED", 0)
	log.Printf("Worker %s: extracting data..., JOB: %+v", job.Type, result)
	w.redisClient.Set(ctx, job.ID.String(), "IN PROGRESS", 0)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: saving data to database..., JOB: %s", job.Type, job.ID)
	w.redisClient.Set(ctx, job.ID.String(), "DONE", 0)
}

func (w *Workers) callbackWork(ctx context.Context, job models.Job[models.CallBack]) {
	w.redisClient.Set(ctx, job.ID.String(), "STARTED", 0)
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	w.redisClient.Set(ctx, job.ID.String(), "IN PROGRESS", 0)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to the given callback..., JOB: %s", job.Type, job.ID)
	w.redisClient.Set(ctx, job.ID.String(), "DONE", 0)
}

func (w *Workers) emailWork(ctx context.Context, job models.Job[models.Mail]) {
	w.redisClient.Set(ctx, job.ID.String(), "STARTED", 0)
	log.Printf("Worker %s: sending the email..., JOB: %s", job.Type, job.ID)
	w.redisClient.Set(ctx, job.ID.String(), "IN PROGRESS", 0)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
	w.redisClient.Set(ctx, job.ID.String(), "DONE", 0)
}

func SpawnWorkers(conn *amqp091.Connection, rclient *redis.Client) {
	workerProcess := Workers{
		conn:        conn,
		redisClient: rclient,
	}
	go workerProcess.run()
}

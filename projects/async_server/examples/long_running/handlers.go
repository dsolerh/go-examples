package main

import (
	"async_server/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

// JobServer holds handler functions
type JobServer struct {
	Queue   amqp091.Queue
	Channel *amqp091.Channel
	Conn    *amqp091.Connection
}

func (s *JobServer) publish(msgType string, jsonBody []byte) error {
	message := amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonBody,
		Type:        msgType,
	}
	err := s.Channel.Publish(
		"",        // exchange
		queueName, // routing key(Queue)
		false,     // mandatory
		false,     // immediate
		message,
	)
	handleError(err, "Error while generating JobID")
	return err
}

func (s *JobServer) asyncDBHandler(w http.ResponseWriter, r *http.Request) {
	jobID := uuid.New()
	queryParams := r.URL.Query()
	// Ex: client_time: 1569174071
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	handleError(err, "Error while converting client time")
	clientTime := time.Unix(unixTime, 0)

	jobMsg := models.Job[models.Log]{
		ID:        jobID,
		Type:      models.LogType,
		ExtraData: models.Log{ClientTime: clientTime},
	}
	jsonBody, err := json.Marshal(jobMsg)
	handleError(err, "JSON body creation failed")

	if err := s.publish(jobMsg.Type.String(), jsonBody); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonBody)
	} else {
		log.Println("Some error happened:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncMailHandler(w http.ResponseWriter, r *http.Request) {
	jobID := uuid.New()
	queryParams := r.URL.Query()
	// Ex: client_time: 1569174071
	email := queryParams.Get("email")

	jobMsg := models.Job[models.Mail]{
		ID:        jobID,
		Type:      models.MailType,
		ExtraData: models.Mail{EmailAddress: email},
	}
	jsonBody, err := json.Marshal(jobMsg)
	handleError(err, "JSON body creation failed")

	if err := s.publish(jobMsg.Type.String(), jsonBody); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonBody)
	} else {
		log.Println("Some error happened:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncCallbackHandler(w http.ResponseWriter, r *http.Request) {
	jobID := uuid.New()
	queryParams := r.URL.Query()
	// Ex: client_time: 1569174071
	callbackURL := queryParams.Get("callback_url")

	jobMsg := models.Job[models.CallBack]{
		ID:        jobID,
		Type:      models.CallBackType,
		ExtraData: models.CallBack{CallBackURL: callbackURL},
	}
	jsonBody, err := json.Marshal(jobMsg)
	handleError(err, "JSON body creation failed")

	if err := s.publish(jobMsg.Type.String(), jsonBody); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonBody)
	} else {
		log.Println("Some error happened:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getServer(name string) JobServer {
	/*
		Creates a server object and initiates
		the Channel and Queue details to publish messages
	*/
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq-host:5672/")
	handleError(err, "Dialing failed to RabbitMQ broker")
	channel, err := conn.Channel()
	handleError(err, "Fetching channel failed")
	jobQueue, err := channel.QueueDeclare(
		name,  // Name of the queue
		false, // Message is persisted or not
		false, // Delete message when unused
		false, // Exclusive
		false, // No Waiting time
		nil,   // Extra args
	)
	handleError(err, "Job queue creation failed")
	return JobServer{Conn: conn, Channel: channel, Queue: jobQueue}
}

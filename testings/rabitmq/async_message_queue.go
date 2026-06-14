package rabbit

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type MessageQueue interface {
	CreateQueue(exchangeName, queueName, routingKey string, createExchange bool) (string, error)
	ConsumeQueue(name string, fn func(body []byte)) error
	ConsumeQueueRoute(name string, fn func(route string, body []byte)) error
	PublishWithContext(ctx context.Context, exchange, key string, data any) error
}

var _ MessageQueue = (*rabbitQueue)(nil)

type rabbitQueue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitQueue(url string) (*rabbitQueue, error) {
	// connect to rabbitmq
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &rabbitQueue{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *rabbitQueue) Close() {
	r.conn.Close()
	r.ch.Close()
}

func (r *rabbitQueue) CreateQueue(exchangeName, queueName, routingKey string, createExchange bool) (string, error) {
	if createExchange {
		err := r.ch.ExchangeDeclare(
			exchangeName, // name
			"topic",      // type
			true,         // durable
			false,        // auto-deleted
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			return "", err
		}
	}

	// Declare a queue (RabbitMQ will generate a random name if empty)
	q, err := r.ch.QueueDeclare(
		queueName, // name (empty = random name)
		true,      // durable
		true,      // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return "", err
	}

	err = r.ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key pattern
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return "", err
	}

	return q.Name, nil
}

func (r *rabbitQueue) ConsumeQueue(name string, fn func(body []byte)) error {
	msgsAdd, err := r.ch.Consume(
		name,  // queue
		"",    // consumer tag (empty = auto-generated)
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgsAdd {
			fn(msg.Body)
		}
	}()

	return nil
}

func (r *rabbitQueue) ConsumeQueueRoute(name string, fn func(route string, body []byte)) error {
	msgs, err := r.ch.Consume(
		name,  // queue
		"",    // consumer tag (empty = auto-generated)
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			fn(msg.RoutingKey, msg.Body)
		}
	}()

	return nil
}

func (r *rabbitQueue) PublishWithContext(ctx context.Context, exchange, key string, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	log.Printf("Calling PublishWithContext: body (%s)\n", body)
	return r.ch.PublishWithContext(
		ctx,
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
}

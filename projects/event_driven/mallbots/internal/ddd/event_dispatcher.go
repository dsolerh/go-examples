package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

var _ EventSubscriber = (*EventDispatcher)(nil)
var _ EventPublisher = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

// Publish implements EventPublisher.
func (e *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range e.handlers[event.EventName()] {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Subscribe implements EventSubscriber.
func (e *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.handlers[event.EventName()] = append(e.handlers[event.EventName()], handler)
}

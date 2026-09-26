package mediator

type Sender interface {
	Send(req any) (res any)
}

type Publisher interface {
	Publish(notification any)
}

type Mediator interface {
	Sender
	Publisher
}

type Handler interface {
}

type mediator struct {
	handlers []Handler
}

func (m *mediator) Send(req any) (res any) {
	//
	return req
}

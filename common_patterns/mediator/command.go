package mediator

import "fmt"

type UpdateXCommand struct {
	ID   string
	Name string
}

type UpdateXCommandHandler[R any] struct{}

func (h *UpdateXCommandHandler[R]) Handle(request R) {
	fmt.Printf("request: %v\n", request)
}

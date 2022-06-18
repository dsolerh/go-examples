package context

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initialize calls 3 functions to set up, then
// logs before terminating
func Initialize() {
	// set basic log up
	log.SetHandler(text.New(os.Stdout))
	// initialize our context
	ctx := context.Background()
	// create a logger and link it to the context
	ctx, e := FromContext(ctx, log.Log)

	// set a field
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	ctx = gatherName(ctx)
	e.Info("after gatherName")
	ctx = gatherLocation(ctx)
	e.Info("after gatherLocation")

	fmt.Printf("ctx: %v\n", ctx)
}

func gatherName(ctx context.Context) context.Context {
	return WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) context.Context {
	return WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
}

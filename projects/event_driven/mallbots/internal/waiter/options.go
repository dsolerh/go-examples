package waiter

import "context"

type config struct {
	parentCtx    context.Context
	catchSignals bool
}

type Option func(c *config)

func ParentContext(ctx context.Context) Option {
	return func(c *config) {
		c.parentCtx = ctx
	}
}

func CatchSignals() Option {
	return func(c *config) {
		c.catchSignals = true
	}
}

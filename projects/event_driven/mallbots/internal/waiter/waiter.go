package waiter

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type WaitFunc func(ctx context.Context) error

type Waiter interface {
	Add(fns ...WaitFunc)
	Wait() error
	Context() context.Context
	CancelFunc() context.CancelFunc
}

type waiter struct {
	ctx    context.Context
	fns    []WaitFunc
	cancel context.CancelFunc
}

func New(options ...Option) Waiter {
	cfg := &config{
		parentCtx:    context.Background(),
		catchSignals: false,
	}

	for _, option := range options {
		option(cfg)
	}

	w := &waiter{
		fns: []WaitFunc{},
	}

	if cfg.catchSignals {
		w.ctx, w.cancel = signal.NotifyContext(w.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	} else {
		w.ctx, w.cancel = context.WithCancel(cfg.parentCtx)
	}

	return w
}

func (w *waiter) Add(fns ...WaitFunc) {
	w.fns = append(w.fns, fns...)
}

func (w waiter) Wait() (err error) {
	g, ctx := errgroup.WithContext(w.ctx)

	g.Go(func() error {
		<-ctx.Done()
		w.cancel()
		return nil
	})

	for _, fn := range w.fns {
		g.Go(func() error { return fn(ctx) })
	}

	return g.Wait()
}

func (w waiter) Context() context.Context {
	return w.ctx
}

func (w waiter) CancelFunc() context.CancelFunc {
	return w.cancel
}

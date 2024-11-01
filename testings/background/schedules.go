package main

import (
	"time"
)

var _ schedule = (*recurringSchedule)(nil)

func newRecurringSchedule(duration, expiresIn time.Duration, immediate bool) *recurringSchedule {
	var expires *time.Timer
	if expiresIn != 0 {
		expires = time.NewTimer(expiresIn)
	}
	return &recurringSchedule{
		readyTicker:    time.NewTicker(duration),
		doneTimer:      expires,
		immediateStart: immediate,
	}
}

type recurringSchedule struct {
	readyTicker    *time.Ticker
	doneTimer      *time.Timer
	immediateStart bool
}

// close implements schedule.
func (d *recurringSchedule) close() {}

// done implements schedule.
func (d *recurringSchedule) done() <-chan time.Time {
	if d.doneTimer != nil {
		return d.doneTimer.C
	}
	return nil
}

// ready implements schedule.
func (d *recurringSchedule) ready() <-chan time.Time {
	if d.immediateStart {
		d.immediateStart = false
		ch := make(chan time.Time, 1)
		ch <- time.Now()
		return ch
	}
	return d.readyTicker.C
}

var _ schedule = (*oneOffSchedule)(nil)

func newOneOffSchedule(executeAt time.Time) *oneOffSchedule {
	expiration := time.Until(executeAt)
	s := &oneOffSchedule{
		readyChan: make(chan time.Time),
		doneChan:  make(chan empty),
		completed: false,
	}
	go s.processTicks(expiration)

	return s
}

type oneOffSchedule struct {
	readyChan chan time.Time
	doneChan  chan empty
	completed bool
}

func (d *oneOffSchedule) processTicks(expiration time.Duration) {
	ticks := time.After(expiration)
	for {
		select {
		case t := <-ticks:
			d.readyChan <- t
			d.completed = true
		case <-d.doneChan:
			return
		}
	}
}

// close implements schedule.
func (d *oneOffSchedule) close() {
	close(d.doneChan)
}

// done implements schedule.
func (d *oneOffSchedule) done() <-chan time.Time {
	if d.completed {
		ch := make(chan time.Time, 1)
		ch <- time.Now()
		return ch
	}
	return nil
}

// ready implements schedule.
func (d *oneOffSchedule) ready() <-chan time.Time {
	return d.readyChan
}

var _ schedule = neverSchedule{}

type neverSchedule struct{}

// close implements schedule.
func (n neverSchedule) close() {}

// done implements schedule.
func (n neverSchedule) done() <-chan time.Time {
	ch := make(chan time.Time, 1)
	ch <- time.Now()
	return ch
}

// ready implements schedule.
func (n neverSchedule) ready() <-chan time.Time {
	return nil
}

// Copyright © 2020-2021 The EVEN Solutions Developers Team

package context

import (
	ctx "context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type (
	// Context describes application's context interface.
	Context interface {
		// Embedded stdlib context interface.
		ctx.Context

		// Cancel calls cancel func of the context instance.
		// After the first call, subsequent calls to a Cancel do nothing.
		Cancel()

		// WgAdd adds delta to context's wait group.
		WgAdd(delta int)

		// WgDone decrements context's wait group counter by one.
		WgDone()

		// WgWait blocks until resolved context wait group counter is zero.
		WgWait()

		// WithCancel returns a copy of application context with a new done channel.
		// The returned context's Done channel is closed when the returned cancel
		// function is called or when the parent context's Done channel is closed,
		// whichever happens first.
		//
		// Canceling this context releases resources associated with it, so code should
		// call cancel as soon as the operations running in this Context complete.
		WithCancel() Context

		// WithCancelWait returns a copy of application context with a new done channel
		// and appends delta with value 1 to embedded application context wait group.
		// The returned context's Done channel is closed when the returned cancel
		// function is called or when the parent context's Done channel is closed,
		// whichever happens first.
		//
		// Canceling this context releases resources associated with it, so code should
		// call cancel as soon as the operations running in this Context complete.
		//
		// Call method WgDone of embedded application context wait group is required
		// when all tasks and processes finished for this copy of application context.
		WithCancelWait() Context

		// WithTimeout returns a copy of application context with timeout duration
		// specified with application environment into the network configuration.
		//
		// Canceling this context releases resources associated with it, so code should
		// call cancel as soon as the operations running in this Context complete.
		WithTimeout(time.Duration) Context
	}

	// context implements application's context interface.
	context struct {
		cancel    ctx.CancelFunc
		parent    ctx.Context
		waitGroup *sync.WaitGroup
	}
)

var (
	// Make sure context implements interface.
	_ Context = (*context)(nil)

	// Canceled is the error returned by Context.Err
	// when the context is canceled.
	Canceled = ctx.Canceled // nolint:gochecknoglobals

	// DeadlineExceeded is the error returned by Context.Err
	// when the context's deadline passes.
	DeadlineExceeded = ctx.DeadlineExceeded // nolint:gochecknoglobals
)

// NewContext constructs context of the application.
func NewContext() Context {
	signalsListener := func(cancel ctx.CancelFunc) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, // to stop channel
			syscall.SIGINT,  // interrupt
			syscall.SIGQUIT, // quit
			syscall.SIGABRT, // aborted
			syscall.SIGTERM, // terminated
		)

		sig := <-stop

		print("\r") // carriage return
		log.Printf("Received %s signal, shutting down...", sig)
		cancel()
	}

	cc, cancel := ctx.WithCancel(ctx.Background())
	go signalsListener(cancel)

	return &context{
		cancel:    cancel,
		parent:    cc,
		waitGroup: new(sync.WaitGroup),
	}
}

// Cancel implements Context.Cancel method of interface.
func (c *context) Cancel() {
	c.cancel()
}

// Deadline implements context interface.
func (c *context) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

// Done implements context interface.
func (c *context) Done() <-chan struct{} {
	return c.parent.Done()
}

// Err implements context interface.
func (c *context) Err() error {
	return c.parent.Err()
}

// Value implements context interface.
func (c *context) Value(key interface{}) interface{} {
	return c.parent.Value(key)
}

// WgAdd implements Context.WgAdd method of interface.
func (c *context) WgAdd(delta int) {
	c.waitGroup.Add(delta)
}

// WgDone implements Context.WgDone method of interface.
func (c *context) WgDone() {
	c.waitGroup.Done()
}

// WgWait implements Context.WgWait method of interface.
func (c *context) WgWait() {
	c.waitGroup.Wait()
}

// WithCancel implements Context.WithCancel method of interface.
func (c *context) WithCancel() Context {
	cc, cancel := ctx.WithCancel(c)

	return &context{
		cancel:    cancel,
		parent:    cc,
		waitGroup: c.waitGroup,
	}
}

// WithCancelWait implements Context.WithCancelWait method of interface.
func (c *context) WithCancelWait() Context {
	c.waitGroup.Add(1)

	return c.WithCancel()
}

// WithTimeout implements Context.WithTimeout method of interface.
func (c *context) WithTimeout(timeout time.Duration) Context {
	cc, cancel := ctx.WithTimeout(c, timeout)

	return &context{
		cancel:    cancel,
		parent:    cc,
		waitGroup: c.waitGroup,
	}
}

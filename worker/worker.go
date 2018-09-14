package worker

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Pool is a type to manage and limit goroutines
type Pool struct {
	workers  chan int // used to launch up to n goroutines at once
	errgroup *errgroup.Group
	context  context.Context
}

// NewPool creates a *Pool with the specified number of workers and set errgroup and context
func NewPool(workerCount int) *Pool {
	g, ctx := errgroup.WithContext(context.Background())
	return &Pool{
		workers:  make(chan int, workerCount),
		errgroup: g,
		context:  ctx,
	}
}

// Go starts the goroutine process and manages limiting it to the number of workers
func (w *Pool) Go(doWork func() error) {
	w.addWorker()
	w.errgroup.Go(func() error {
		defer w.removeWorker()
		select {
		case <-w.context.Done():
			return nil // returning early
		default:
			return doWork()
		}
	})
}

// Wait is a proxy to the errgroup.Wait function
func (w *Pool) Wait() error {
	return w.errgroup.Wait()
}

// addWorker increments the worker channel
func (w *Pool) addWorker() {
	w.workers <- 1
}

// removeWorker decrements the worker channel
func (w *Pool) removeWorker() {
	<-w.workers
}

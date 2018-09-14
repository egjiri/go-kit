package work

import "github.com/egjiri/go-kit/worker"

// Run does work concurrently using a goroutine manager
func Run(workerCount int) error {
	w := worker.NewPool(workerCount)
	for _, chunk := range chunks {
		chunk := chunk
		w.Go(func() error {
			return doWork(chunk)
		})
	}
	return w.Wait()
}

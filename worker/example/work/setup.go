package work

import (
	"errors"
	"flag"
	"math/rand"
	"time"

	"github.com/egjiri/go-kit/worker"
)

var chaos = false
var chunks = makeChunks(100000)

func init() {
	worker.ShowRunningGoroutines(10 * time.Millisecond)
	flag.BoolVar(&chaos, "chaos", false, "enables random errors")
	flag.Parse()
}

func makeChunks(n int) [][10]int {
	chunks := make([][10]int, n)
	for i := 0; i < n; i++ {
		var chunk [10]int
		for i := 0; i < 10; i++ {
			chunk[i] = rand.Intn(50) + 1 // generate random number between 1-50
		}
		chunks[i] = chunk
	}
	return chunks
}

func doWork(numbers [10]int) error {
	for range numbers {
		// ADD CHAOS. If the randomly generated number equals 42, we'll return an error.
		if chaos && rand.Intn(5000) == 42 {
			return errors.New("error 42")
		}
		time.Sleep(10 * time.Millisecond) // Time consumming task
	}
	return nil
}

package worker

import (
	"fmt"
	"runtime"
	"time"
)

// ShowRunningGoroutines prints out the number of running goroutines
// An interval of 10 * time.Millisecond seems to work quite well
func ShowRunningGoroutines(interval time.Duration) {
	go func() {
		for {
			<-time.After(interval)
			fmt.Println("Goroutine Count:", runtime.NumGoroutine())
		}
	}()
}

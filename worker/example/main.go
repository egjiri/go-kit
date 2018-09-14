package main

import (
	"log"

	"github.com/egjiri/go-kit/worker/example/work"
)

// go run main.go -chaos
func main() {
	if err := work.Run(200); err != nil {
		log.Fatalf("====> ERROR: errgroup received an error: %v\n", err)
	}
}

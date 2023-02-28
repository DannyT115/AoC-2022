package main

import (
	"log"
	"time"
)

// A function to track the time of execution (see how fast the solution was).
func executionTimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf(" %s took %s to run.", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "The Solution")

}

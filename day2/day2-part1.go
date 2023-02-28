package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// A function to track the time of execution (see how fast the solution was).
func executionTimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf(" %s took %s to run.", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "The Solution")

	fileInput, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: Could not open file.")
	}
	defer fileInput.Close()

	var totalScore int

	// A Map of all possible outcomes / scores.
	var strategyMap = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	scanner := bufio.NewScanner(fileInput)

	// Read the map, compare the input (file with options) to the map and add the score up.
	for scanner.Scan() {
		totalScore += strategyMap[scanner.Text()]
	}

	fmt.Println("\nThe answer to Day2-Part1 is:", totalScore)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// A function to track the time of execution (see how fast the solution was).
func executionTimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf(" %s took %s to run.", name, elapsed)
}

func main() {
	defer executionTimeTracker(time.Now(), "The Solution")

	fileInput, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error: Could not read in file")
	}
	defer fileInput.Close()

	highestCalories := 0
	elfCalories := 0

	scanner := bufio.NewScanner(fileInput)

	for scanner.Scan() {
		scran, err := strconv.Atoi(scanner.Text())
		elfCalories += scran

		if err != nil && highestCalories < elfCalories {
			if highestCalories < elfCalories {
				highestCalories = elfCalories
			}
			elfCalories = 0
		}
	}
	fmt.Println("\nThe answer to Day1-Part1 is:", highestCalories)
}

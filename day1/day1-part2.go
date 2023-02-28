package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var chadElves []int

	scanner := bufio.NewScanner(fileInput)

	for scanner.Scan() {
		scran, err := strconv.Atoi(scanner.Text())
		elfCalories += scran

		if err != nil {
			highestCalories = elfCalories
			chadElves = append(chadElves, highestCalories)
			elfCalories = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(chadElves)))
	sum := 0
	for _, v := range chadElves[:3] {
		sum += v
	}
	fmt.Println("\nThe answer to Day1-Part2 is:", sum)
}

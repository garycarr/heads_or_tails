package main

import (
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const minInRow = 20

func main() {
	// When to print to the console
	printAt := 100000000

	// The number of times printed to the console
	timesReachedPrintAt := 0

	counter := 0

	headsInRow := 0
	tailsInRow := 0
	heads := 0
	// tails := 1

	runningCounter := make(map[int]int)

	rand.Seed(time.Now().UnixNano())

	for {
		counter++
		newToss := rand.Intn(2)
		if newToss == heads {
			runningCounter = toinCoss(&headsInRow, &tailsInRow, runningCounter)
		} else {
			runningCounter = toinCoss(&tailsInRow, &headsInRow, runningCounter)
		}

		if counter%printAt == 0 {
			rand.Seed(time.Now().UnixNano())
			counter = 0
			timesReachedPrintAt++
			printStatus(&timesReachedPrintAt, runningCounter)

		}
	}
}

func toinCoss (winningFaceInRow *int, losingFaceInRow *int, runningCounter map[int]int) map[int]int {
	if *losingFaceInRow >= minInRow {
		runningCounter[*losingFaceInRow]++
	}
	*losingFaceInRow = 0
	*winningFaceInRow++
	return runningCounter
}

func printStatus (timesReachedPrintAt *int, runningCounter map[int]int) {
	log.Printf("\n%s hundred million tosses - ", strconv.Itoa(*timesReachedPrintAt))
	var keys []int
	for k := range runningCounter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		log.Printf("%s in a row occured %s times", strconv.Itoa(key), strconv.Itoa(runningCounter[key]))
	}
}

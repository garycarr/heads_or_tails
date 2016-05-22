package main

import (
	"flag"
	"os"
	"strconv"
	"sync"
)

// CounterTotal to count the overallTotal
type CounterTotal struct {
	billions      int
	count         int
	inARowCounter map[int]int
	mux           sync.Mutex
}

const billion = 1000000000
const defaultConcurrentThreads = 1
const defaultMaxTosses = billion
const defaultMinInRow = 15
const defaultNumSides = 2
const defaultPrintEvery = 1000000
const defaultVerbose = true

type Config struct {
	concurrentThreads int
	maxTosses         int
	minInRow          int
	numSides          int
	printEvery        int
	verbose           bool
}

var (
	concurrentThreads = flag.Int("concurrent-threads", defaultConcurrentThreads, "Num of concurrent threads")
	maxTosses         = flag.Int("max-tosses", defaultMaxTosses, "Maximum number of tosses. Input -1 for the highest possible")
	minInRow          = flag.Int("min-count", defaultMinInRow, "Min in a row to start counting")
	numSides          = flag.Int("num-sides", defaultNumSides, "How many sides. Example - 2 for a coin or 6 for a dice")
	printEvery        = flag.Int("print-every", defaultPrintEvery, "How many tosses to print at")
	verbose           = flag.Bool("verbose", defaultVerbose, "Prints out at every print-every")
	stop              = false
	testRun           = false
)
var cT CounterTotal

func init() {
	cT = CounterTotal{inARowCounter: make(map[int]int)}
}

func mustGetEnv(env string) int {
	envInt, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		panic(err)
	}
	return envInt
}

func getConfig() Config {
	// Should be able to make this neater
	var conf Config
	if os.Getenv("ENV_SET") != "" {
		// Find a package to do this
		// Need to fail if not found
		var verbose bool
		if os.Getenv("VERBOSE") == "true" {
			verbose = true
		} else {
			verbose = false
		}
		conf = Config{
			concurrentThreads: mustGetEnv("CONCURRENT_THREAD"),
			maxTosses:         mustGetEnv("MAX_TOSSES"),
			minInRow:          mustGetEnv("MIN_IN_ROW"),
			numSides:          mustGetEnv("NUM_SIDES"),
			printEvery:        mustGetEnv("PRINT_EVERY"),
			verbose:           verbose,
		}
	} else {
		flag.Parse()
		conf = Config{
			concurrentThreads: *concurrentThreads,
			maxTosses:         *maxTosses,
			minInRow:          *minInRow,
			numSides:          *numSides,
			printEvery:        *printEvery,
			verbose:           *verbose,
		}
	}
	return conf
}

func main() {
	app := newApp(getConfig())
	app.start()
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/divan/num2words"
)

type app struct {
	conf         Config
	counterTotal *CounterTotal
}

func newApp(c Config) *app {
	return &app{
		conf:         c,
		counterTotal: &cT,
	}
}

func (a app) start() {
	a.setupCoinTossers()
}

func (a app) setupCoinTossers() {
	var wg sync.WaitGroup
	inARowChan := make(chan map[int]int, a.conf.concurrentThreads)
	go func() {
		a.updateTotalCount(inARowChan)
	}()
	for i := 0; i < a.conf.concurrentThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.toss(inARowChan)
		}()
	}
	wg.Wait()
}

func (a app) toss(inARowChan chan map[int]int) {
	inARowCounter := make(map[int]int)
	sameSideInARow := 0
	tosses := 0
	currentToss := -1
	rand.Seed(time.Now().UnixNano())
	for {
		newToss := rand.Intn(a.conf.numSides)
		tosses++
		if newToss == currentToss {
			sameSideInARow++
		} else {
			if sameSideInARow >= a.conf.minInRow {
				inARowCounter[sameSideInARow]++
			}
			sameSideInARow = 0
		}
		currentToss = newToss
		if tosses%a.conf.printEvery == 0 {
			// To store the running total of where we are
			inARowChan <- inARowCounter
			tosses = 0
			inARowCounter = make(map[int]int)
		}
		if stop {
			return
		}
	}
}

func (a app) updateTotalCount(inARowChan chan map[int]int) {
	for inARowCounter := range inARowChan {
		a.counterTotal.mux.Lock()
		a.counterTotal.count += a.conf.printEvery
		if a.counterTotal.count >= billion {
			a.counterTotal.billions++
			a.counterTotal.count = cT.count % billion
		}
		billions := a.counterTotal.billions
		count := a.counterTotal.count
		for key, val := range inARowCounter {
			a.counterTotal.inARowCounter[key] += val
		}

		if a.conf.verbose {
			a.printStatus(a.counterTotal.inARowCounter, billions, count)
		}

		// if app.conf.maxTosses > 0 && count > app.conf.conf.maxTosses {
		if count > a.conf.maxTosses {
			if a.conf.verbose {
				a.printStatus(cT.inARowCounter, billions, count)
			}
			stop = true
		}
		a.counterTotal.mux.Unlock()
	}
}

func (a app) printStatus(inARowCounterCopy map[int]int, billions, count int) {
	var keys []int
	for k := range inARowCounterCopy {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("Tossed %d billion and %s\n", billions, num2words.Convert(count))
	for _, k := range keys {
		fmt.Printf("in a row:%d, times:%d\n", k, inARowCounterCopy[k])
	}
	fmt.Printf("\n")
}

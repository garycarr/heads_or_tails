// The following allows running `go generate` to bundle the static and template
// assets into binary form so that the web console can be a stand-alone binary.
// If you update anything in static/ or templates/ be sure to run `go generate`
// and commit the binary differences as well.
//go:generate go get github.com/jteeuwen/go-bindata/...
//go:generate go-bindata -o resources/resources.go -pkg resources -prefix resources -ignore resources/resources.go resources/...
//go:generate gofmt -w resources/resources.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/divan/num2words"
	"github.com/garycarr/heads_or_tails/api"
)

type app struct {
	server http.Server

	apiHandler api.APISomething

	conf Config

	counterTotal *CounterTotal
}

func newApp(c Config) *app {
	apiSomething := api.NewAPISomething()
	return &app{
		apiHandler:   apiSomething,
		conf:         c,
		counterTotal: &cT,
	}
}

func (a *app) start() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.setupCoinTossers()
	}()
	a.endpoints()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	wg.Wait()
}

func (a *app) endpoints() {
	http.HandleFunc("/", a.apiHandler.IndexHandlerGET)
	http.HandleFunc("/about", a.apiHandler.AboutHandlerGET)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources/static"))))
}

func (a *app) setupCoinTossers() {
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

func (a *app) toss(inARowChan chan map[int]int) {
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

func (a *app) updateTotalCount(inARowChan chan map[int]int) {
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
		if a.conf.maxTosses != -1 && count > a.conf.maxTosses {
			stop = true
		}
		a.counterTotal.mux.Unlock()
	}
}

func (a *app) printStatus(inARowCounterCopy map[int]int, billions, count int) {
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

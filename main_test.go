package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFixture struct {
	app *app

	// cleanup after the tests
	cleanup func()
}

const defaultTestingConcurrentThreads = 1
const defaultTestingMaxTosses = 10000
const defaultTestingMinInRow = 15
const defaultTestingNumSides = 2
const defaultTestingPrintEvery = 10001

func createTestFixture(c Config) testFixture {
	fixture := testFixture{}
	if c.concurrentThreads == 0 {
		c.concurrentThreads = defaultTestingConcurrentThreads
	}
	if c.maxTosses == 0 {
		c.maxTosses = defaultTestingMaxTosses
	}
	if c.minInRow == 0 {
		c.minInRow = defaultTestingMinInRow
	}
	if c.numSides == 0 {
		c.numSides = defaultTestingNumSides
	}
	if c.printEvery == 0 {
		c.printEvery = defaultTestingPrintEvery
	}
	fixture.app = newApp(c)
	fixture.cleanup = func() {
		stop = false
	}
	return fixture
}

func doBenchmarkCoinTossers(b *testing.B, tF testFixture) {
	for n := 0; n < b.N; n++ {
		stop = false
		cT.inARowCounter = make(map[int]int)
		cT.count = 0
		tF.app.start()
	}
}

func BenchmarkSetupCoinTossersConc1MaxTosses10000000(b *testing.B) {
	tF := createTestFixture(Config{concurrentThreads: 1, maxTosses: 10000000})
	tF.cleanup()
	doBenchmarkCoinTossers(b, tF)
}
func BenchmarkSetupCoinTossersConc5MaxTosses10000000(b *testing.B) {
	tF := createTestFixture(Config{concurrentThreads: 5, maxTosses: 10000000})
	tF.cleanup()
	doBenchmarkCoinTossers(b, tF)
}

// Just test the app can compile and start
func TestNewApp(t *testing.T) {
	tF := createTestFixture(Config{})
	tF.cleanup()
	tF.app.start()
	assert.Equal(t, 1, 1, "The app ran and completed")
}

// Make sure that the defaults are applied
func TestGetConfig(t *testing.T) {
	conf := getConfig()
	assert.Equal(t, conf.concurrentThreads, defaultConcurrentThreads)
	assert.Equal(t, conf.maxTosses, defaultMaxTosses)
	assert.Equal(t, conf.minInRow, defaultMinInRow)
	assert.Equal(t, conf.numSides, defaultNumSides)
	assert.Equal(t, conf.printEvery, defaultPrintEvery)
	assert.Equal(t, conf.verbose, defaultVerbose)
}

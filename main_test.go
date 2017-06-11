package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
		// Should be a better way to do this
		http.DefaultServeMux = new(http.ServeMux)
	}
	return fixture
}

func TestIndexHandler(t *testing.T) {
	tF := createTestFixture(Config{})
	defer tF.cleanup()
	tF.app.start()
	ts := httptest.NewServer(http.HandlerFunc(tF.app.apiHandler.IndexHandlerGET))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(body), "Heads or Tails")
}

func TestAboutHandler(t *testing.T) {
	tF := createTestFixture(Config{})
	defer tF.cleanup()
	tF.app.start()
	ts := httptest.NewServer(http.HandlerFunc(tF.app.apiHandler.AboutHandlerGET))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(body), "Heads or Tails about page")
}

// Make sure that the defaults are applied
func TestGetConfigNoArgs(t *testing.T) {
	conf := getConfig()
	assert.Equal(t, conf.concurrentThreads, defaultConcurrentThreads)
	assert.Equal(t, conf.maxTosses, defaultMaxTosses)
	assert.Equal(t, conf.minInRow, defaultMinInRow)
	assert.Equal(t, conf.numSides, defaultNumSides)
	assert.Equal(t, conf.printEvery, defaultPrintEvery)
	assert.Equal(t, conf.verbose, defaultVerbose)
}

// Make sure that the defaults are applied
func TestGetConfigEnvVars(t *testing.T) {
	defer clearEnvs()

	os.Setenv("ENV_SET", "TRUE")
	os.Setenv("CONCURRENT_THREAD", "5")
	os.Setenv("MAX_TOSSES", "100000")
	os.Setenv("MIN_IN_ROW", "5")
	os.Setenv("NUM_SIDES", "2")
	os.Setenv("PRINT_EVERY", "100")
	os.Setenv("VERBOSE", "false")

	conf := getConfig()
	assert.Equal(t, conf.concurrentThreads, 5)
	assert.Equal(t, conf.maxTosses, 100000)
	assert.Equal(t, conf.minInRow, 5)
	assert.Equal(t, conf.numSides, 2)
	assert.Equal(t, conf.printEvery, 100)
	assert.Equal(t, conf.verbose, false)
}

func clearEnvs() {
	os.Setenv("ENV_SET", "")
	os.Setenv("CONCURRENT_THREAD", "")
	os.Setenv("MAX_TOSSES", "")
	os.Setenv("MIN_IN_ROW", "")
	os.Setenv("NUM_SIDES", "")
	os.Setenv("PRINT_EVERY", "")
	os.Setenv("VERBOSE", "")
}

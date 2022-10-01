package main

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/alecthomas/kingpin"
)

var (
	requests    = kingpin.Flag("requests", "requests").Short('n').Default("100").Int()
	concurrency = kingpin.Flag("concurrency", "concurrency").Short('c').Default("2").Int()
	endpointURL = kingpin.Flag("endpoint-url", "endpoint-url").Short('e').String()

	apiKey string
)

func init() {
	kingpin.Parse()

	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())
}

func main() {
	StartWorkers(*concurrency)

	for i := 0; i < *requests; i++ {
		AddTask(i)
	}

	StopWorkers()
}

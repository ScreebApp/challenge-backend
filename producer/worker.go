package main

import (
	"sync"
	"sync/atomic"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

var jobs chan int
var wg sync.WaitGroup
var dup atomic.Pointer[EventBody]

func init() {
	dup.Store(lo.ToPtr(NewRequestIdentify()))
}

func StartWorkers(nbr int) {
	jobs = make(chan int, nbr*10)
	wg.Add(nbr)

	for w := 0; w < nbr; w++ {
		go func() {
			worker()
		}()
	}
}

func StopWorkers() {
	close(jobs)
	wg.Wait()
}

func worker() {
	defer wg.Done()

	// consume queue
	for taskNbr := range jobs {
		var event EventBody

		switch taskNbr % 20 {
		case 0, 6, 10, 14:
			event = NewRequestIdentify()
		case 1:
			event = NewRequestGroup()
		case 2, 7, 11, 15, 17, 19:
			event = NewRequestPage()
		case 3, 8, 12, 16, 18:
			event = NewRequestScreen()
		case 4, 9, 13:
			event = NewRequestTrack()
		case 5:
			// 5% of events will be send twice
			event = *dup.Load()
		}

		if taskNbr%3 == 0 {
			dup.Store(&event)
		}

		err := sendWebhookRequest(event)
		if err != nil {
			logrus.Errorf("Event %s: %s", event.Type, err.Error())
		} else {
			logrus.Infof("[%d] Event %s", taskNbr, event.Type)
		}
	}
}

func AddTask(i int) {
	jobs <- i
}

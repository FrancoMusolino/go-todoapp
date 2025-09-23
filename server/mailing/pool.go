package mailing

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MailingDispatcher struct {
	jobQueue       chan Message
	workerPool     chan chan ProcessMessageJob
	mailingService IMailService
	NumWorkers     int
	ResultsChan    chan error
}

func NewMailingDispatcher(jobQueue chan Message, mailingService IMailService, numWorkers int, resultsChan chan error) *MailingDispatcher {
	return &MailingDispatcher{
		jobQueue:       jobQueue,
		workerPool:     make(chan chan ProcessMessageJob, numWorkers),
		NumWorkers:     numWorkers,
		mailingService: mailingService,
		ResultsChan:    resultsChan,
	}
}

func (md *MailingDispatcher) Run() {
	fmt.Printf("Running dispatcher with %v workers\n", md.NumWorkers)
	for i := 0; i < md.NumWorkers; i++ {
		worker := NewWorker(i+1, md)
		go worker.Start()
	}

	go md.dispatch()
}

func (md *MailingDispatcher) dispatch() {
	for {
		msg := <-md.jobQueue

		go func() {
			workerJobQueue := <-md.workerPool
			job := ProcessMessageJob{
				CorrelationId: uuid.New(),
				Message:       msg,
				Timestamp:     time.Now(),
			}

			workerJobQueue <- job
		}()
	}
}

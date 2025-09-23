package mailing

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ProcessMessageJob struct {
	CorrelationId uuid.UUID
	Message       Message
	Timestamp     time.Time
}

type MailingWorker struct {
	ID          int
	workerQueue chan ProcessMessageJob
	dispatcher  *MailingDispatcher
}

func NewWorker(ID int, dispatcher *MailingDispatcher) *MailingWorker {
	return &MailingWorker{
		ID:          ID,
		workerQueue: make(chan ProcessMessageJob),
		dispatcher:  dispatcher,
	}
}

func (w *MailingWorker) Start() {
	fmt.Printf("Starting Worker with ID %v\n", w.ID)

	for {
		w.dispatcher.workerPool <- w.workerQueue
		job := <-w.workerQueue

		fmt.Printf("Worker with ID %v received job %s\n", w.ID, job.CorrelationId)
		w.process(&job)
	}
}

func (w *MailingWorker) process(job *ProcessMessageJob) {
	fmt.Printf("Worker with ID %v processing job\n", w.ID)
	err := w.dispatcher.mailingService.SendHTML(job.Message)
	w.dispatcher.ResultsChan <- err
}

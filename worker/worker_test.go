package worker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/worker"
)

func Test_worker(t *testing.T) {
	// init job queue
	job := worker.NewJob(10)
	// init dispatcher
	dispatcher := worker.NewDispatcher(3, job)
	dispatcher.Run()

	// init job
	for i := 0; i < 10; i++ {
		dispatcher.Job.Queue <- NewJob(i)
	}

	// wait for all jobs done
	time.Sleep(time.Second * 5)

	// close job queue
	job.Close()
}

// NewJob returns a new job
func NewJob(payload interface{}) worker.Job {
	return worker.JobFunc(func() {
		fmt.Println("Executing job with payload:", payload, " time:", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second * 1)
		// do something
	})
}

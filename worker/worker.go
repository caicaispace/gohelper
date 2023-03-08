package worker

// Worker is a worker that executes a job
type Worker struct {
	// WorkerPool is the pool of workers channel
	WorkerPool chan chan Job
	// JobChannel is the job channel
	JobChannel chan Job
	// quit is the quit channel
	quit chan bool
}

// NewWorker returns a new worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start starts the worker
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				job.Execute()
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop stops the worker
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// Dispatcher is a job dispatcher
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	MaxWorkers int
	Job        *job
}

// NewDispatcher returns a new dispatcher
func NewDispatcher(maxWorkers int, job *job) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers, Job: job}
}

// Run starts the dispatcher
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.Job.Queue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

// Job is a job to be run
type Job interface {
	Execute()
}

// // JobFunc is a function that implements the Job interface
type JobFunc func()

// Execute executes the job
func (f JobFunc) Execute() {
	f()
}

type job struct {
	// Func is the function to be executed
	Func JobFunc
	// Queue is the queue to be used
	Queue chan Job
}

// NewJob returns a new job
func NewJob(maxQueues int) *job {
	return &job{
		Queue: make(chan Job, maxQueues),
	}
}

// Execute executes the job
func (j *job) Execute() {
	j.Func()
}

// Close closes the job queue
func (j *job) Close() {
	close(j.Queue)
}

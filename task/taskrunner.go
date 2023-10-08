package task

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name     string
	Function func() error
}

type TaskRunner struct {
	tasks []Task
}

func NewTaskRunner() *TaskRunner {
	return &TaskRunner{
		tasks: []Task{},
	}
}

func (r *TaskRunner) Add(name string, function func() error) {
	task := Task{name, function}
	r.tasks = append(r.tasks, task)
}

func (r *TaskRunner) AddNoName(function func() error) {
	r.Add(fmt.Sprintf("Task %d", len(r.tasks)), function)
}

func (r *TaskRunner) Adds(tasks []Task) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *TaskRunner) AddsNoName(functions []func() error) {
	for i, function := range functions {
		r.Add(fmt.Sprintf("Task %d", i), function)
	}
}

func (r *TaskRunner) Run() {
	var wg sync.WaitGroup
	for _, task := range r.tasks {
		wg.Add(1)
		time.Sleep(10 * time.Millisecond)
		go func(task Task) {
			defer wg.Done()
			task.Function()
		}(task)
	}
	wg.Wait()
}

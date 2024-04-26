package main

import (
	"fmt"
	"sync"
	"time"
)

const MaxResources = 3

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	resourceProvider := NewResourceProvider(cond, MaxResources)

	wg.Add(10)

	for i := range 10 {
		go func(workerID int) {
			defer wg.Done()

			worker := NewWorker(workerID, cond, resourceProvider)
			worker.Run()
		}(i)
	}

	wg.Wait()
}

type ResourceProvider struct {
	maxResources       int
	availableResources int
	cond               *sync.Cond
}

func NewResourceProvider(cond *sync.Cond, maxResources int) *ResourceProvider {
	return &ResourceProvider{
		cond:               cond,
		availableResources: maxResources,
	}
}

func (rp *ResourceProvider) AvailableResources() int {
	return rp.availableResources
}

func (rp *ResourceProvider) AcquireResoirce() {
	rp.availableResources--
}

func (rp *ResourceProvider) ReleaseResource() {
	rp.availableResources++
}

type Worker struct {
	id   int
	cond *sync.Cond
	rp   *ResourceProvider
}

func NewWorker(workerID int, cond *sync.Cond, rp *ResourceProvider) *Worker {
	return &Worker{
		id:   workerID,
		cond: cond,
		rp:   rp,
	}
}

func (w *Worker) Run() {
	w.cond.L.Lock()

	for w.rp.AvailableResources() == 0 {
		fmt.Printf("Worker %d is waiting for resources\n", w.id)
		w.cond.Wait()
	}

	w.rp.AcquireResoirce()
	fmt.Printf("Worker %d acquired resource. Remaining resources: %d\n", w.id, w.rp.AvailableResources())
	w.cond.L.Unlock()

	time.Sleep(1 * time.Second) // Simulating work

	w.cond.L.Lock()
	defer w.cond.L.Unlock()

	w.rp.ReleaseResource()
	fmt.Printf("Worker %d released resource. Remaining resources: %d\n", w.id, w.rp.AvailableResources())
	w.cond.Signal()
}

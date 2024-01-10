package ethreads_test

import (
	"io"
	"sync"
	"testing"
	"time"

	"github.com/bylexus/go-stdlib/elog"
	"github.com/bylexus/go-stdlib/ethreads"
)

func TestThreadPoolWithJobFn(t *testing.T) {
	logger := elog.NewSeverityLogger(io.Discard)
	tp := ethreads.NewThreadPool(3, &logger)
	results := make([]int, 0)
	mu := sync.Mutex{}

	tp.Start()
	for i := 1; i <= 3; i++ {
		localI := i
		tp.AddJobFn(func(id ethreads.ThreadId) {
			mu.Lock()
			defer mu.Unlock()
			results = append(results, localI)
			time.Sleep(time.Millisecond * 100)
		})
	}
	tp.Shutdown()
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}
	if results[0] == results[1] || results[1] == results[2] {
		t.Errorf("Expected different results, got same")
	}
}

type TestJob struct {
	res *chan int
}

func (j *TestJob) Run(threadId ethreads.ThreadId) {
	*j.res <- int(threadId)
	time.Sleep(time.Millisecond * 100)
}

func TestThreadPoolWithJobStruct(t *testing.T) {
	logger := elog.NewSeverityLogger(io.Discard)
	tp := ethreads.NewThreadPool(3, &logger)
	resultChan := make(chan int, 3)
	results := make([]int, 0)

	tp.Start()
	for i := 1; i <= 3; i++ {
		job := TestJob{
			res: &resultChan,
		}
		tp.AddJob(&job)
	}
	tp.Shutdown()
	close(resultChan)
	for r := range resultChan {
		results = append(results, r)
	}

	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}
	if results[0] == results[1] || results[1] == results[2] {
		t.Errorf("Expected different results, got same")
	}
}

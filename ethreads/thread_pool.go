package ethreads

import (
	"io"
	"sync"

	"github.com/bylexus/go-stdlib/elog"
)

type ThreadId int
type JobFn func(threadId ThreadId)
type Job interface {
	Run(threadId ThreadId)
}

type ThreadPool struct {
	threads     int
	job_channel chan JobFn
	wait_group  *sync.WaitGroup
	logger      *elog.SeverityLogger
	isRunning   bool
	lock        *sync.Mutex
}

// NewThreadPool creates a new ThreadPool with the specified number of threads (goroutines).
// Optionally, a log.SeverityLogger can be passed in. If not provided, a null logger will be used.
func NewThreadPool(threads int, logger *elog.SeverityLogger) ThreadPool {
	if logger == nil {
		l := elog.NewSeverityLogger(io.Discard)
		logger = &l
	}
	tp := ThreadPool{
		threads:     threads,
		job_channel: make(chan JobFn, threads),
		wait_group:  &sync.WaitGroup{},
		logger:      logger,
		isRunning:   false,
		lock:        &sync.Mutex{},
	}
	return tp
}

// Starts the number of configured worker routines in the pool
//
// If the pool is not started, enqueued jobs are not processed.
func (tp *ThreadPool) Start() {
	tp.lock.Lock()
	defer tp.lock.Unlock()
	if !tp.isRunning {
		tp.logger.Debug("Starting thread pool with %d threads", tp.threads)
		tp.isRunning = true
		for i := 1; i <= tp.threads; i++ {
			tp.wait_group.Add(1)
			go tp.workerRun(ThreadId(i))
		}
	}
}

// Enqueues a job by providing a JobFn function.
func (tp *ThreadPool) AddJobFn(jobFn JobFn) {
	tp.lock.Lock()
	defer tp.lock.Unlock()
	if tp.isRunning {
		tp.job_channel <- jobFn
	}
}

func (tp *ThreadPool) AddJob(job Job) {
	tp.AddJobFn(job.Run)
}

func (tp *ThreadPool) Shutdown() {
	if tp.isRunning {
		tp.logger.Debug("Shutting down thread pool")
		tp.lock.Lock()
		tp.isRunning = false
		tp.lock.Unlock()
		close(tp.job_channel)
		tp.wait_group.Wait()
		tp.logger.Debug("Thread pool stopped")
	}
}

func (tp *ThreadPool) workerRun(threadId ThreadId) {
	tp.logger.Debug("Thread %d: Started", threadId)
	for job := range tp.job_channel {
		tp.logger.Debug("Thread %d: Working ...", threadId)
		job(threadId)
		tp.logger.Debug("Thread %d: Done with job ...", threadId)
	}
	tp.wait_group.Done()
	tp.logger.Debug("Thread %d: Stopped", threadId)
}

// Package main demonstrates a production-ready worker pool pattern
// This is an example file for the golang-pro skill
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

// =============================================================================
// Job and Result Types
// =============================================================================

// Job represents a unit of work to be processed
type Job struct {
	ID      int
	Payload string
}

// Result represents the outcome of processing a job
type Result struct {
	JobID    int
	Output   string
	Duration time.Duration
	Err      error
}

// =============================================================================
// Worker Pool Implementation
// =============================================================================

// WorkerPool manages a pool of workers for concurrent job processing
type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	logger     *slog.Logger
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers
func NewWorkerPool(numWorkers int, bufferSize int, logger *slog.Logger) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, bufferSize),
		results:    make(chan Result, bufferSize),
		logger:     logger,
	}
}

// Start begins the worker pool, spawning workers and waiting for jobs
func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(ctx, i)
	}

	// Close results channel when all workers are done
	go func() {
		wp.wg.Wait()
		close(wp.results)
	}()
}

// Submit adds a job to the pool for processing
// Returns false if the pool is shutting down
func (wp *WorkerPool) Submit(ctx context.Context, job Job) bool {
	select {
	case <-ctx.Done():
		return false
	case wp.jobs <- job:
		return true
	}
}

// Close signals that no more jobs will be submitted
func (wp *WorkerPool) Close() {
	close(wp.jobs)
}

// Results returns the results channel for consuming processed jobs
func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

// worker processes jobs from the jobs channel
func (wp *WorkerPool) worker(ctx context.Context, id int) {
	defer wp.wg.Done()

	wp.logger.Info("worker started", slog.Int("worker_id", id))

	for {
		select {
		case <-ctx.Done():
			wp.logger.Info("worker stopping due to context cancellation",
				slog.Int("worker_id", id))
			return
		case job, ok := <-wp.jobs:
			if !ok {
				wp.logger.Info("worker stopping, jobs channel closed",
					slog.Int("worker_id", id))
				return
			}
			result := wp.processJob(ctx, id, job)
			wp.results <- result
		}
	}
}

// processJob handles the actual work for a single job
func (wp *WorkerPool) processJob(ctx context.Context, workerID int, job Job) Result {
	start := time.Now()

	wp.logger.Info("processing job",
		slog.Int("worker_id", workerID),
		slog.Int("job_id", job.ID),
	)

	// Simulate work with context awareness
	select {
	case <-ctx.Done():
		return Result{
			JobID: job.ID,
			Err:   ctx.Err(),
		}
	case <-time.After(100 * time.Millisecond): // Simulated work
		// Process the job
		output := fmt.Sprintf("processed: %s", job.Payload)

		return Result{
			JobID:    job.ID,
			Output:   output,
			Duration: time.Since(start),
		}
	}
}

// =============================================================================
// Batch Processor (Higher-Level Abstraction)
// =============================================================================

// BatchProcessor processes jobs in batches with configurable concurrency
type BatchProcessor struct {
	pool   *WorkerPool
	logger *slog.Logger
}

func NewBatchProcessor(numWorkers int, logger *slog.Logger) *BatchProcessor {
	return &BatchProcessor{
		pool:   NewWorkerPool(numWorkers, numWorkers*2, logger),
		logger: logger,
	}
}

// ProcessBatch processes a slice of jobs and returns all results
func (bp *BatchProcessor) ProcessBatch(ctx context.Context, jobs []Job) ([]Result, error) {
	// Start the worker pool
	bp.pool.Start(ctx)

	// Submit all jobs
	go func() {
		for _, job := range jobs {
			if !bp.pool.Submit(ctx, job) {
				bp.logger.Warn("failed to submit job, context cancelled",
					slog.Int("job_id", job.ID))
				break
			}
		}
		bp.pool.Close()
	}()

	// Collect results
	var results []Result
	for result := range bp.pool.Results() {
		results = append(results, result)
	}

	// Check for context cancellation
	if ctx.Err() != nil {
		return results, ctx.Err()
	}

	return results, nil
}

// =============================================================================
// Fan-Out/Fan-In Pattern
// =============================================================================

// Pipeline stage function type
type StageFunc func(ctx context.Context, in <-chan int) <-chan int

// Generator creates a channel of integers from a slice
func Generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

// Square squares each number from the input channel
func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

// FanOut distributes work from input channel to multiple workers
func FanOut(ctx context.Context, in <-chan int, numWorkers int, stage StageFunc) []<-chan int {
	outputs := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		outputs[i] = stage(ctx, in)
	}
	return outputs
}

// FanIn merges multiple channels into a single channel
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Start a goroutine for each input channel
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Close output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// =============================================================================
// Rate Limited Worker
// =============================================================================

// RateLimitedProcessor limits the rate of job processing
type RateLimitedProcessor struct {
	ticker *time.Ticker
	logger *slog.Logger
}

func NewRateLimitedProcessor(ratePerSecond int, logger *slog.Logger) *RateLimitedProcessor {
	return &RateLimitedProcessor{
		ticker: time.NewTicker(time.Second / time.Duration(ratePerSecond)),
		logger: logger,
	}
}

func (rp *RateLimitedProcessor) Process(ctx context.Context, jobs <-chan Job, results chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-rp.ticker.C:
			select {
			case <-ctx.Done():
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				result := rp.processJob(job)
				select {
				case <-ctx.Done():
					return
				case results <- result:
				}
			}
		}
	}
}

func (rp *RateLimitedProcessor) processJob(job Job) Result {
	start := time.Now()
	output := fmt.Sprintf("rate-limited processed: %s", job.Payload)
	return Result{
		JobID:    job.ID,
		Output:   output,
		Duration: time.Since(start),
	}
}

func (rp *RateLimitedProcessor) Stop() {
	rp.ticker.Stop()
}

// =============================================================================
// Main Demonstration
// =============================================================================

func main() {
	// Setup logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Create context with cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Demo 1: Basic Worker Pool
	logger.Info("=== Demo 1: Basic Worker Pool ===")
	demoWorkerPool(ctx, logger)

	// Demo 2: Batch Processor
	logger.Info("=== Demo 2: Batch Processor ===")
	demoBatchProcessor(ctx, logger)

	// Demo 3: Fan-Out/Fan-In Pipeline
	logger.Info("=== Demo 3: Fan-Out/Fan-In Pipeline ===")
	demoPipeline(ctx, logger)

	logger.Info("All demos completed")
}

func demoWorkerPool(ctx context.Context, logger *slog.Logger) {
	pool := NewWorkerPool(3, 10, logger)
	pool.Start(ctx)

	// Submit jobs
	for i := 0; i < 5; i++ {
		pool.Submit(ctx, Job{ID: i, Payload: fmt.Sprintf("task-%d", i)})
	}
	pool.Close()

	// Collect results
	for result := range pool.Results() {
		logger.Info("result received",
			slog.Int("job_id", result.JobID),
			slog.String("output", result.Output),
			slog.Duration("duration", result.Duration),
		)
	}
}

func demoBatchProcessor(ctx context.Context, logger *slog.Logger) {
	processor := NewBatchProcessor(4, logger)

	jobs := make([]Job, 10)
	for i := range jobs {
		jobs[i] = Job{ID: i, Payload: fmt.Sprintf("batch-task-%d", i)}
	}

	results, err := processor.ProcessBatch(ctx, jobs)
	if err != nil {
		logger.Error("batch processing failed", slog.Any("error", err))
		return
	}

	logger.Info("batch completed", slog.Int("total_results", len(results)))
}

func demoPipeline(ctx context.Context, logger *slog.Logger) {
	// Create pipeline: generate -> fan-out to square workers -> fan-in results
	nums := Generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Fan-out to 3 square workers (note: each worker will get some of the input)
	squareChans := FanOut(ctx, nums, 3, Square)

	// Fan-in results
	results := FanIn(ctx, squareChans...)

	// Collect and print results
	var sum int
	for n := range results {
		sum += n
		logger.Info("pipeline result", slog.Int("value", n))
	}
	logger.Info("pipeline complete", slog.Int("sum", sum))
}

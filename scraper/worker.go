package scraper

import (
	"sync"
)

// WorkerPool manages a pool of worker goroutines
type WorkerPool struct {
	jobs    chan string    // Channel for distributing URLs
	results chan Result    // Channel for collecting results
	wg      sync.WaitGroup // WaitGroup for worker synchronization
	scraper *Scraper       // Reference to the scraper for processing URLs
}

// NewWorkerPool creates a new worker pool
// Input: int: Number of workers to create
// Output: *WorkerPool: Initialized worker pool
// Description:
//   - Creates channels for jobs and results
//   - Initializes WaitGroup
//   - Sets up worker goroutines
func NewWorkerPool(size int, scraper *Scraper) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan string, size),
		results: make(chan Result, size),
		wg:      sync.WaitGroup{},
		scraper: scraper,
	}
}

// Start begins processing the URL list
// Input: []string: List of URLs to process
// Output: []Result: Collected results from all workers
// Description:
//   - Distributes URLs to workers
//   - Collects results from all workers
//   - Handles worker synchronization
//   - Returns aggregated results
func (wp *WorkerPool) Start(urls []string) []Result {
	for _, url := range urls {
		wp.jobs <- url
	}
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)

	var allResults []Result
	for result := range wp.results {
		allResults = append(allResults, result)
	}

	return allResults
}

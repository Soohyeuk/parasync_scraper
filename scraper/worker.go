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
	wp := &WorkerPool{
		jobs:    make(chan string, size),
		results: make(chan Result, size),
		wg:      sync.WaitGroup{},
		scraper: scraper,
	}

	// Start worker goroutines
	for i := 0; i < size; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}

	return wp
}

// worker processes URLs from the jobs channel
// Description:
//   - Reads URLs from jobs channel
//   - Uses scraper to process each URL
//   - Sends results to results channel
//   - Handles channel closure
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for url := range wp.jobs {
		result := wp.scraper.ScrapeURL(url)
		wp.results <- result
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
	// Start a goroutine to collect results
	var allResults []Result
	var resultsWg sync.WaitGroup
	resultsWg.Add(1)
	go func() {
		defer resultsWg.Done()
		for result := range wp.results {
			allResults = append(allResults, result)
		}
	}()

	// Send all URLs to workers
	for _, url := range urls {
		wp.jobs <- url
	}
	close(wp.jobs)

	// Wait for all workers to finish
	wp.wg.Wait()
	close(wp.results)

	// Wait for results collection to finish
	resultsWg.Wait()

	return allResults
}

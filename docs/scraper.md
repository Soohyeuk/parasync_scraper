# Scraper Package

Core package containing scraping logic and worker pool implementation. This package is responsible for:
- Managing concurrent web scraping operations
- Implementing the worker pool pattern
- Handling HTTP requests with retries and timeouts
- Parsing HTML content using goquery
- Extracting required data from web pages
- Managing error handling and recovery
- Coordinating between workers and result collection

## Files

### scraper.go
```go
package scraper

import (
    "net/http"
    "sync"
    "time"
    "github.com/PuerkitoBio/goquery"
)

// Scraper represents the main scraping service
type Scraper struct {
    client     *http.Client  // HTTP client with configured timeout
    maxWorkers int          // Maximum number of concurrent workers
    maxRetries int          // Maximum number of retry attempts
}

// Result represents the scraped data from a single URL
type Result struct {
    URL         string   `json:"url"`          // The URL that was scraped
    Title       string   `json:"title"`        // Page title from <title> tag
    Description string   `json:"description"`  // Meta description content
    Headings    []string `json:"headings"`     // All H1 headings found
    Error       string   `json:"error,omitempty"` // Error message if scraping failed
}

// Function Headers

// NewScraper creates a new scraper instance
// Input: *Config: Configuration for the scraper
// Output: *Scraper: Initialized scraper instance
// Description:
//   - Creates HTTP client with configured timeout
//   - Sets up worker pool size
//   - Configures retry mechanism
//   - Initializes necessary channels and sync primitives
func NewScraper(config *Config) *Scraper

// Scrape processes a list of URLs concurrently
// Input: []string: List of URLs to scrape
// Output: 
//   - []Result: Slice of scraping results
//   - error: Any error encountered during scraping
// Description:
//   - Creates worker pool
//   - Distributes URLs among workers
//   - Collects and aggregates results
//   - Handles any errors during processing
func (s *Scraper) Scrape(urls []string) ([]Result, error)

// scrapeURL processes a single URL with retries
// Input: string: URL to scrape
// Output: Result: Scraping result for the URL
// Description:
//   - Attempts to fetch URL with retries
//   - Parses HTML content
//   - Extracts required data
//   - Handles any errors during processing
func (s *Scraper) scrapeURL(url string) Result
```

### worker.go
```go
package scraper

// WorkerPool manages a pool of worker goroutines
type WorkerPool struct {
    jobs    chan string      // Channel for distributing URLs
    results chan Result      // Channel for collecting results
    wg      sync.WaitGroup  // WaitGroup for worker synchronization
}

// Function Headers

// NewWorkerPool creates a new worker pool
// Input: int: Number of workers to create
// Output: *WorkerPool: Initialized worker pool
// Description:
//   - Creates channels for jobs and results
//   - Initializes WaitGroup
//   - Sets up worker goroutines
func NewWorkerPool(size int) *WorkerPool

// worker processes URLs from the jobs channel
// Input: int: Worker ID for logging
// Output: None
// Description:
//   - Continuously reads URLs from jobs channel
//   - Processes each URL
//   - Sends results to results channel
//   - Handles worker cleanup
func (wp *WorkerPool) worker(id int)

// Start begins processing the URL list
// Input: []string: List of URLs to process
// Output: []Result: Collected results from all workers
// Description:
//   - Distributes URLs to workers
//   - Collects results from all workers
//   - Handles worker synchronization
//   - Returns aggregated results
func (wp *WorkerPool) Start(urls []string) []Result
```

### utils.go
```go
package scraper

// Function Headers

// fetchURL retrieves and parses a URL
// Input: 
//   - string: URL to fetch
//   - *http.Client: Configured HTTP client
// Output:
//   - *goquery.Document: Parsed HTML document
//   - error: Any error during fetch or parsing
// Description:
//   - Makes HTTP GET request
//   - Handles response status codes
//   - Parses HTML into goquery document
//   - Manages request timeouts
func fetchURL(url string, client *http.Client) (*goquery.Document, error)

// extractData extracts required data from HTML document
// Input: *goquery.Document: Parsed HTML document
// Output:
//   - string: Page title
//   - string: Meta description
//   - []string: H1 headings
// Description:
//   - Extracts <title> tag content
//   - Finds meta description
//   - Collects all H1 headings
//   - Handles missing elements gracefully
func extractData(doc *goquery.Document) (title, description string, headings []string)

// writeResults writes results to JSON file
// Input:
//   - []Result: Scraping results to write
//   - string: Output file path
// Output: error if write fails
// Description:
//   - Marshals results to JSON
//   - Creates output file
//   - Writes formatted JSON
//   - Handles file I/O errors
func writeResults(results []Result, filename string) error

// readURLs reads URLs from input file
// Input: string: Input file path
// Output:
//   - []string: List of URLs
//   - error: Any error during file reading
// Description:
//   - Opens and reads input file
//   - Parses URLs (one per line)
//   - Validates URL format
//   - Handles file I/O errors
func readURLs(filename string) ([]string, error)
``` 
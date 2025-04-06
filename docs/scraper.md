# Scraper Package Documentation

The scraper package provides the core functionality for concurrent web scraping with rate limiting and error handling.

## Components

### Scraper

The `Scraper` struct is the main component that manages the scraping process:

```go
type Scraper struct {
    Client     *http.Client
    MaxWorkers int
    MaxRetries int
}
```

#### Methods

- `NewScraper(config ScraperConfig) *Scraper`: Creates a new scraper instance
- `Scrape(urls []string) ([]Result, error)`: Scrapes multiple URLs concurrently
- `ScrapeURL(url string) Result`: Scrapes a single URL with retries

### WorkerPool

The `WorkerPool` struct manages concurrent scraping:

```go
type WorkerPool struct {
    jobs    chan string
    results chan Result
    wg      sync.WaitGroup
    scraper *Scraper
}
```

#### Methods

- `NewWorkerPool(size int, scraper *Scraper) *WorkerPool`: Creates a new worker pool
- `Start(urls []string) []Result`: Processes URLs concurrently

### Result

The `Result` struct represents scraped data:

```go
type Result struct {
    URL         string
    Title       string
    Description string
    Headings    []string
    Error       error
}
```

## Rate Limiting

The scraper implements rate limiting to prevent overwhelming target servers:

1. Configurable delays between requests
2. Automatic retry on rate limit errors (429)
3. Worker pool size should be adjusted based on target server's rate limits

## Error Handling

The scraper handles various error scenarios:

1. HTTP errors (4xx, 5xx)
2. Rate limiting (429)
3. Network timeouts
4. HTML parsing errors

## Usage Example

```go
config := ScraperConfig{
    MaxWorkers: 10,
    MaxRetries: 3,
    Timeout:    30 * time.Second,
}

scraper := NewScraper(config)
results, err := scraper.Scrape(urls)
if err != nil {
    log.Fatal(err)
}
```

## Best Practices

1. Adjust worker count based on target server's rate limits
2. Set appropriate timeouts for your use case
3. Handle rate limiting errors gracefully
4. Monitor memory usage with large URL lists
5. Use appropriate retry counts for your needs

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
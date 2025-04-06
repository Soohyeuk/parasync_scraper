package scraper

import (
	"net/http"

	"github.com/Soohyeuk/parasync_scraper/pkg/config"
)

// Scraper represents the main scraping service
type Scraper struct {
	client     *http.Client // HTTP client with configured timeout
	maxWorkers int          // Maximum number of concurrent workers
	maxRetries int          // Maximum number of retry attempts
}

// Result represents the scraped data from a single URL
type Result struct {
	URL         string   `json:"url"`             // The URL that was scraped
	Title       string   `json:"title"`           // Page title from <title> tag
	Description string   `json:"description"`     // Meta description content
	Headings    []string `json:"headings"`        // All H1 headings found
	Error       string   `json:"error,omitempty"` // Error message if scraping failed
}

// NewScraper creates a new scraper instance
// Input: *config.ScraperConfig: Configuration for the scraper
// Output: *Scraper: Initialized scraper instance
// Description:
//   - Creates HTTP client with configured timeout
//   - Sets up worker pool size
//   - Configures retry mechanism
//   - Initializes necessary channels and sync primitives
func NewScraper(config *config.ScraperConfig) *Scraper {
	return &Scraper{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		maxWorkers: config.MaxWorkers,
		maxRetries: config.MaxRetries,
	}
}

// Scrape processes a list of URLs concurrently
// Input: []string: List of URLs to scrape
// Output:
//   - []Result: Slice of scraping results
//   - error: Any error encountered during scraping
//
// Description:
//   - Creates worker pool
//   - Distributes URLs among workers
//   - Collects and aggregates results
//   - Handles any errors during processing
func (s *Scraper) Scrape(urls []string) ([]Result, error) {
	// Create a worker pool
	pool := NewWorkerPool(s.maxWorkers, s)

	// Start the worker pool with the URLs
	results := pool.Start(urls)

	return results, nil
}

// ScrapeURL processes a single URL with retries
// Input: string: URL to scrape
// Output: Result: Scraping result for the URL
// Description:
//   - Attempts to fetch URL with retries
//   - Parses HTML content
//   - Extracts required data
//   - Handles any errors during processing
func (s *Scraper) ScrapeURL(url string) Result {
	query, err := FetchURL(url, s.client)
	if err != nil {
		return Result{
			URL:   url,
			Error: err.Error(),
		}
	}

	title, description, headings := ExtractData(query)
	return Result{
		URL:         url,
		Title:       title,
		Description: description,
		Headings:    headings,
		Error:       "",
	}
}

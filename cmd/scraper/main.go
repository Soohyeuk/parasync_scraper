package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Soohyeuk/parasync_scraper/pkg/config"
	"github.com/Soohyeuk/parasync_scraper/scraper"
)

// parseFlags parses and validates command line arguments
// Input: Command line arguments
// Output:
//   - *config.CLIConfig: Parsed configuration or nil if error
//   - error: Parsing or validation error if any
//
// Description:
//   - Parses -input, -output, -workers, -timeout, -retries flags
//   - Sets default values if not provided
//   - Validates flag values are within acceptable ranges
func parseFlags() (*config.CLIConfig, error) {
	cfg := &config.CLIConfig{
		MaxWorkers: 5,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	// Parse command line flags
	flag.StringVar(&cfg.InputFile, "input", "", "Path to input file with URLs (one per line)")
	flag.StringVar(&cfg.OutputFile, "output", "result/output.json", "Path to output file for results")
	flag.IntVar(&cfg.MaxWorkers, "workers", 5, "Maximum number of concurrent workers")
	flag.DurationVar(&cfg.Timeout, "timeout", 30*time.Second, "HTTP request timeout")
	flag.IntVar(&cfg.MaxRetries, "retries", 3, "Maximum number of retry attempts")
	flag.Parse()

	// Validate flags
	if cfg.InputFile == "" {
		return nil, fmt.Errorf("input file is required")
	}

	if cfg.MaxWorkers < 1 {
		return nil, fmt.Errorf("workers must be at least 1")
	}

	if cfg.Timeout < 1*time.Second {
		return nil, fmt.Errorf("timeout must be at least 1 second")
	}

	if cfg.MaxRetries < 0 {
		return nil, fmt.Errorf("retries must be non-negative")
	}

	return cfg, nil
}

// run executes the main program logic
// Input: *config.CLIConfig: Parsed configuration
// Output: error if any error occurs
// Description:
//   - Reads URLs from input file
//   - Creates and configures scraper
//   - Scrapes URLs in parallel
//   - Writes results to output file
func run(cfg *config.CLIConfig) error {
	// Read URLs from input file
	urls, err := scraper.ReadURLs(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("error reading URLs: %w", err)
	}

	// Create scraper with configuration
	s := scraper.NewScraper(cfg.ToScraperConfig())

	// Scrape URLs
	results, err := s.Scrape(urls)
	if err != nil {
		return fmt.Errorf("error scraping URLs: %w", err)
	}

	// Write results to output file
	err = scraper.WriteResults(results, cfg.OutputFile)
	if err != nil {
		return fmt.Errorf("error writing results: %w", err)
	}

	fmt.Printf("Successfully scraped %d URLs. Results written to %s\n", len(urls), cfg.OutputFile)
	return nil
}

func main() {
	// Parse command line flags
	cfg, err := parseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	// Run the program
	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

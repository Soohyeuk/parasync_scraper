package config

import "time"

// ScraperConfig holds scraper configuration
type ScraperConfig struct {
	MaxWorkers int           // Maximum number of concurrent workers
	Timeout    time.Duration // HTTP request timeout
	MaxRetries int           // Maximum number of retry attempts
}

// CLIConfig holds command-line interface configuration
type CLIConfig struct {
	InputFile  string        // Path to input file with URLs
	OutputFile string        // Path to output file for results
	MaxWorkers int           // Maximum number of concurrent workers
	Timeout    time.Duration // HTTP request timeout
	MaxRetries int           // Maximum number of retry attempts
}

// ToScraperConfig converts CLIConfig to ScraperConfig
func (c *CLIConfig) ToScraperConfig() *ScraperConfig {
	return &ScraperConfig{
		MaxWorkers: c.MaxWorkers,
		Timeout:    c.Timeout,
		MaxRetries: c.MaxRetries,
	}
}

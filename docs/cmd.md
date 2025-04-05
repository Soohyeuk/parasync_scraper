# Command Line Interface (cmd)

Directory containing the main application entry point and CLI handling. This directory is responsible for:
- Parsing and validating command-line arguments
- Initializing the scraper with user-provided configuration
- Managing the program's execution flow
- Handling errors and providing user feedback
- Coordinating between different components of the application

## Files

### main.go
```go
package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

// Config holds all CLI configuration
type Config struct {
    InputFile  string        // Path to file containing URLs (one per line)
    OutputFile string        // Path to save JSON results
    MaxWorkers int           // Maximum number of concurrent workers (default: 5)
    Timeout    time.Duration // HTTP request timeout (default: 30s)
    MaxRetries int           // Maximum number of retry attempts (default: 3)
}

// Function Headers

// main is the entry point of the application
// Input: Command line arguments
// Output: Exit code (0 for success, non-zero for errors)
// Description: Orchestrates the entire program flow, including:
//   - Parsing command line flags
//   - Validating configuration
//   - Initializing the scraper
//   - Running the scraping process
//   - Handling errors and cleanup
func main()

// parseFlags parses and validates command line arguments
// Input: Command line arguments
// Output: 
//   - *Config: Parsed configuration or nil if error
//   - error: Parsing or validation error if any
// Description: 
//   - Parses -input, -output, -workers, -timeout, -retries flags
//   - Sets default values if not provided
//   - Validates flag values are within acceptable ranges
func parseFlags() (*Config, error)

// validateConfig ensures all required fields are set
// Input: *Config: Configuration to validate
// Output: error if validation fails, nil otherwise
// Description:
//   - Checks if input file exists and is readable
//   - Verifies output file path is writable
//   - Ensures worker count is between 1 and 20
//   - Validates timeout is between 5s and 2m
//   - Confirms retry count is between 0 and 5
func validateConfig(cfg *Config) error

// run executes the main program logic
// Input: *Config: Validated configuration
// Output: error if execution fails, nil otherwise
// Description:
//   - Reads URLs from input file
//   - Initializes scraper with configuration
//   - Executes scraping process
//   - Writes results to output file
//   - Handles any errors during execution
func run(cfg *Config) error
``` 
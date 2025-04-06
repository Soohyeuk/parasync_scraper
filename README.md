# ParaSync Scraper

A high-performance, concurrent web scraper built in Go that supports parallel processing of URLs with configurable worker pools. The focus is in exploring concurrency in Go, rather than the information retrieved from the scraper. However, the scraper does work and retrieved all the tags I requested.

## Features

- Concurrent URL scraping with configurable worker pool size
- Rate limiting and retry mechanisms to handle HTTP errors
- Configurable timeouts and retry attempts
- JSON output format for scraped data
- Command-line interface for easy usage
- Robust error handling and logging

## Installation

```bash
git clone https://github.com/yourusername/parasync_scraper.git
cd parasync_scraper
go mod download
```

## Usage

### Basic Usage

```bash
go run cmd/scraper/main.go -input urls.txt -output results.json
```

### Command Line Options

- `-input`: Path to input file containing URLs (one per line)
- `-output`: Path to output file for results (default: result/output.json)
- `-workers`: Number of concurrent workers (default: 5)
- `-timeout`: HTTP request timeout in seconds (default: 30s)
- `-retries`: Maximum number of retry attempts (default: 3)

### Example

```bash
# Scrape 100 URLs with 10 workers
go run cmd/scraper/main.go -input test/data/urls.txt -output result/output.json -workers 10 -timeout 30s -retries 3
```

## Output Format

The scraper outputs results in JSON format:

```json
[
  {
    "url": "https://example.com",
    "title": "Page Title",
    "description": "Page description",
    "headings": ["Heading 1", "Heading 2"],
    "error": null
  }
]
```

## Rate Limiting

The scraper includes built-in rate limiting to prevent overwhelming target servers:
- Configurable delays between requests
- Automatic retry on rate limit errors (429)
- Worker pool size should be adjusted based on target server's rate limits

## Project Structure

```
parasync_scraper/
├── cmd/
│   └── scraper/         # Command-line interface
├── pkg/
│   └── config/          # Configuration management
├── scraper/             # Core scraping functionality
│   ├── scraper.go       # Main scraper implementation
│   ├── worker.go        # Worker pool implementation
│   └── utils.go         # Utility functions
├── test/                # Test files and fixtures
└── result/              # Output directory
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
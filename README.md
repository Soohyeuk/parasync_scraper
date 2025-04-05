# ParaSync Scraper

A high-performance, parallel web scraping CLI tool written in Go. This tool efficiently scrapes multiple URLs concurrently while maintaining controlled resource usage.

## Features

- Parallel web scraping with controlled concurrency
- Extracts page title, meta description, and H1 headings
- Configurable input via file or command-line arguments
- JSON output format
- Automatic retry mechanism for failed requests
- Request timeout handling
- Concurrent processing with worker pool pattern

## Project Structure

```
.
├── cmd/
│   └── scraper/         # Main application entry point
├── scraper/             # Core scraping logic and utilities
├── docs/                # Documentation
└── test/               # Test files and data
```

## Requirements

- Go 1.21 or higher
- Dependencies:
  - github.com/PuerkitoBio/goquery
  - Other dependencies will be listed in go.mod

## Installation

```bash
git clone https://github.com/soohyeuk/parasync_scraper.git
cd parasync_scraper
go mod download
```

## Usage

### Building

```bash
go build -o scraper cmd/scraper/main.go
```

### Running

```bash
# Using input file
./scraper -input urls.txt -output results.json

# Using command-line arguments
./scraper -urls "https://example.com,https://example.org" -output results.json
```

### Input File Format

Create a text file with one URL per line:

```
https://example.com
https://example.org
https://example.net
```

### Output Format

The tool generates a JSON file with the following structure:

```json
[
  {
    "url": "https://example.com",
    "title": "Example Title",
    "description": "Page description",
    "headings": ["Heading 1", "Another Heading"]
  }
]
```

## Features in Detail

### Concurrency Control
- Uses a worker pool pattern with 5 concurrent workers
- Implements WaitGroups for safe goroutine management
- Uses channels for task distribution and result collection

### Error Handling
- Automatic retry mechanism (up to 3 attempts)
- HTTP request timeout handling
- Graceful error reporting in JSON output

### Performance
- Parallel processing of URLs
- Controlled resource usage
- Efficient memory management

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 
# Command Line Interface Documentation

The CLI provides a user-friendly interface for the scraper with configurable options.

## Usage

```bash
go run cmd/scraper/main.go [options]
```

## Options

| Option    | Description                                    | Default           |
|-----------|------------------------------------------------|-------------------|
| -input    | Path to input file with URLs (one per line)    | (required)        |
| -output   | Path to output file for results                | result/output.json|
| -workers  | Number of concurrent workers                   | 5                 |
| -timeout  | HTTP request timeout in seconds                | 30s               |
| -retries  | Maximum number of retry attempts               | 3                 |

## Examples

### Basic Usage
```bash
go run cmd/scraper/main.go -input urls.txt
```

### Custom Configuration
```bash
go run cmd/scraper/main.go -input urls.txt -output results.json -workers 10 -timeout 60s -retries 5
```

### Rate-Limited Scraping
```bash
# Use fewer workers for rate-limited sites
go run cmd/scraper/main.go -input urls.txt -workers 3 -timeout 30s
```

## Input File Format

Create a text file with one URL per line:
```
https://example.com
https://example.org
https://example.net
```

## Output Format

Results are written to the specified output file in JSON format:
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

## Error Handling

The CLI handles various error scenarios:
1. Invalid input file
2. Invalid command-line options
3. Scraping errors (logged to output file)
4. Rate limiting errors (429)
5. Network timeouts

## Best Practices

1. Start with a small number of workers (3-5) and adjust based on target server's rate limits
2. Use appropriate timeouts for your network conditions
3. Monitor the output file for errors and adjust configuration accordingly
4. For large URL lists, consider splitting into smaller batches
5. Use appropriate retry counts based on target server reliability
``` 
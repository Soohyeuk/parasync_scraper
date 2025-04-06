# Test Package Documentation

The test package contains comprehensive tests for the scraper functionality, including worker pool behavior and rate limiting handling.

## Test Files

### scraper_test.go

Contains the main test suite for the scraper package:

```go
func TestNewScraper(t *testing.T)
func TestScrapeURL(t *testing.T)
func TestWorkerPool(t *testing.T)
func TestExtractData(t *testing.T)
```

## Test Fixtures

Located in `test/fixtures/`:

1. `basic.html`: Simple HTML page for basic scraping tests
2. `complex.html`: Complex HTML page with nested elements
3. `error.html`: Error page for testing error handling

## Test Data

Located in `test/data/`:

1. `urls.txt`: Sample URLs for testing (100 subreddits)

## Test Coverage

The test suite covers:

1. Scraper Initialization
   - Configuration validation
   - Default values
   - Custom settings

2. URL Scraping
   - Single URL scraping
   - Multiple URL concurrent scraping
   - Error handling
   - Rate limiting

3. Worker Pool
   - Concurrent processing
   - Result collection
   - Error propagation
   - Resource cleanup

4. Data Extraction
   - Title extraction
   - Description extraction
   - Heading extraction
   - Error handling

## Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific test
go test -v -run TestScrapeURL

# Run with race detector
go test -race ./...
```

## Test Best Practices

1. Use appropriate timeouts for network-dependent tests
2. Clean up test resources after each test
3. Test both success and error cases
4. Verify rate limiting behavior
5. Check concurrent processing correctness 
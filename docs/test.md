# Test Package

Directory containing test files and test data. This package is responsible for:
- Providing unit tests for all components
- Including integration tests for the complete workflow
- Maintaining test fixtures and sample data
- Ensuring code coverage and quality
- Validating error handling and edge cases
- Testing concurrent operations and race conditions

## Files

### scraper_test.go
```go
package scraper_test

import (
    "testing"
    "github.com/soohyeuk/parasync_scraper/scraper"
)

// Function Headers

// TestNewScraper tests scraper initialization
// Input: *testing.T: Testing context
// Output: None
// Description:
//   - Tests creation with valid configuration
//   - Verifies default values are set correctly
//   - Checks error handling with invalid config
//   - Validates HTTP client initialization
func TestNewScraper(t *testing.T)

// TestScrapeURL tests single URL scraping
// Input: *testing.T: Testing context
// Output: None
// Description:
//   - Tests successful URL scraping
//   - Verifies data extraction accuracy
//   - Tests error handling for invalid URLs
//   - Validates retry mechanism
//   - Checks timeout handling
func TestScrapeURL(t *testing.T)

// TestWorkerPool tests worker pool functionality
// Input: *testing.T: Testing context
// Output: None
// Description:
//   - Tests concurrent URL processing
//   - Verifies worker pool size limits
//   - Tests result collection
//   - Validates worker cleanup
//   - Checks race conditions
func TestWorkerPool(t *testing.T)

// TestExtractData tests HTML data extraction
// Input: *testing.T: Testing context
// Output: None
// Description:
//   - Tests title extraction
//   - Verifies description parsing
//   - Tests heading collection
//   - Validates handling of missing elements
//   - Checks HTML parsing edge cases
func TestExtractData(t *testing.T)
```

### test_data/
```
test/
├── data/
│   ├── urls.txt           # Sample URL list for testing
│   └── expected.json      # Expected output format and data
└── fixtures/
    ├── basic.html         # Basic HTML fixture with all required elements
    └── complex.html       # Complex HTML fixture with edge cases
```

### Example Test Data

#### urls.txt
```
https://example.com
https://example.org
https://example.net
```

#### expected.json
```json
[
  {
    "url": "https://example.com",
    "title": "Example Domain",
    "description": "Example website description",
    "headings": ["Welcome to Example.com"]
  }
]
```

### Test Fixtures

#### basic.html
```html
<!DOCTYPE html>
<html>
<head>
    <title>Example Domain</title>
    <meta name="description" content="Example website description">
</head>
<body>
    <h1>Welcome to Example.com</h1>
</body>
</html>
```

#### complex.html
```html
<!DOCTYPE html>
<html>
<head>
    <title>Complex Example</title>
    <meta name="description" content="A more complex example with multiple elements">
</head>
<body>
    <h1>Main Heading</h1>
    <h1>Another Heading</h1>
    <div>
        <h1>Nested Heading</h1>
    </div>
</body>
</html>
``` 
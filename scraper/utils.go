package scraper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// fetchURL retrieves and parses a URL
// Input:
//   - string: URL to fetch
//   - *http.Client: Configured HTTP client
//
// Output:
//   - *goquery.Document: Parsed HTML document
//   - error: Any error during fetch or parsing
//
// Description:
//   - Makes HTTP GET request
//   - Handles response status codes
//   - Parses HTML into goquery document
//   - Manages request timeouts
func fetchURL(url string, client *http.Client) (*goquery.Document, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	return doc, nil
}

// extractData extracts required data from HTML document
// Input: *goquery.Document: Parsed HTML document
// Output:
//   - string: Page title
//   - string: Meta description
//   - []string: H1 headings
//
// Description:
//   - Extracts <title> tag content
//   - Finds meta description
//   - Collects all H1 headings
//   - Handles missing elements gracefully
func extractData(doc *goquery.Document) (title, description string, headings []string) {
	title = doc.Find("title").Text()
	description, _ = doc.Find("meta[name=description]").First().Attr("content")
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		headings = append(headings, s.Text())
	})

	return title, description, headings
}

// writeResults writes results to JSON file
// Input:
//   - []Result: Scraping results to write
//   - string: Output file path
//
// Output: error if write fails
// Description:
//   - Marshals results to JSON
//   - Creates output file
//   - Writes formatted JSON
//   - Handles file I/O errors
func writeResults(results []Result, filename string) error {
	// Create result directory if it doesn't exist
	err := os.Mkdir("result", 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("error creating a result directory %w", err)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling results to JSON: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing results to file: %w", err)
	}

	return nil
}

// readURLs reads URLs from input file
// Input: string: Input file path
// Output:
//   - []string: List of URLs
//   - error: Any error during file reading
//
// Description:
//   - Opens and reads input file
//   - Parses URLs (one per line)
//   - Validates URL format
//   - Handles file I/O errors
func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error finding a file, please provide a correct filename: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}

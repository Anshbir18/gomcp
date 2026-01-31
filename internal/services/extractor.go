package services

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractTextFromHTML extracts readable text from raw HTML
func ExtractTextFromHTML(html string) (string, error) {
	reader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	// Remove unwanted elements
	doc.Find("script, style, noscript").Remove()

	// Extract text from body
	text := doc.Find("body").Text()

	// Clean up whitespace
	cleanText := cleanString(text)

	return cleanText, nil
}


func cleanString(input string) string {
	lines := strings.Split(input, "\n")

	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			cleaned = append(cleaned, line)
		}
	}

	return strings.Join(cleaned, " ")
}

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log/slog"
	"strings"
)

// ParseTitle extracts and returns a parsed title from an HTML document based on the specified program time and title.
func ParseTitle(r io.Reader, programTime, programTitle string) string {
	if r == nil {
		slog.Debug("empty reader passed to ParseTitle")
		return ""
	}

	// Parse HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		slog.Debug("unable to parse HTML document", "error", err)
		return ""
	}

	// Extract text from first li containing the text in program time and itself
	// containing a link with the text in program title
	title := ""
	liSelector := fmt.Sprintf(`li:contains("%s")`, programTime)
	aSelector := fmt.Sprintf(`a:contains("%s")`, programTitle)
	doc.Find(liSelector).Has(aSelector).First().Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.Replace(text, programTime, "", -1)
		text = strings.Replace(text, programTitle, "", -1)
		text = strings.TrimSpace(text)
		title = text
	})

	if title == "" {
		slog.Debug("unable to extract title from HTML document")
	}

	return title
}

package main

import (
	"log/slog"
	"net/http"
)
import "net/http/cookiejar"

// Client is a structure that holds an HTTP client and manages the referer header for HTTP requests.
type Client struct {
	referer    string
	httpClient *http.Client
}

// NewClient initializes and returns a new Client instance with an HTTP client that includes a cookie jar.
func NewClient() Client {
	httpClient := http.Client{}
	jar, err := cookiejar.New(nil)
	if err == nil {
		httpClient.Jar = jar
	}
	return Client{
		httpClient: &httpClient,
	}
}

// Get sends a GET request to the specified URL, optionally pretending to be an XMLHttpRequest if xhr is true,
// returning the response
func (c *Client) Get(url string, xhr ...bool) (*http.Response, error) {
	slog.Debug("attempting GET request", "url", url, "referer", c.referer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("error creating GET request", "url", url, "error", err)
		return nil, err
	}
	// Pretend we are a browser
	req.Header.Add("User-Agent", "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0")
	req.Header.Add("Accept", "text/html, */*; q=0.01")
	req.Header.Add("Accept-Language", "de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	if len(xhr) > 0 && xhr[0] {
		req.Header.Add("X-Requested-With", "XMLHttpRequest")
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("DNT", "1")
	req.Header.Add("Sec-GPC", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")

	// Add referer header if referer was set in previous call
	if c.referer != "" {
		req.Header.Add("Referer", c.referer)
	}

	// Store current URL as referer for the next request
	c.referer = url

	return c.httpClient.Do(req)
}

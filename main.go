package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

// Command line flag vars
var (
	debug        bool
	baseUrl      string
	programUri   string
	programTime  string
	programTitle string
)

const (
	DEFAULT_TIME  = "23:00"
	DEFAULT_TITLE = "x-fade die DJ Nacht"
)

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily

		fmt.Fprintf(w, "Usage of %s: A tiny tool to extract the title of tonight's %s\n", os.Args[0], DEFAULT_TITLE)

		flag.PrintDefaults()

	}

	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.StringVar(&baseUrl, "base", "http://radiox.de", "Base URL")
	flag.StringVar(&programUri, "program_uri", "/plus7/ajax/program_day", "Program URI")
	flag.StringVar(&programTime, "program_time", DEFAULT_TIME, "Program Time")
	flag.StringVar(&programTitle, "program_title", DEFAULT_TITLE, "Program Title")
	flag.Parse()

	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("debug mode enabled")
	}

	fallbackTitle := fmt.Sprintf("%s %s", programTitle, time.Now().Format(time.DateOnly))

	client := NewClient()
	_, err := client.Get(baseUrl + "/")
	if err != nil {
		fmt.Println(fallbackTitle)
		return
	}

	res, err := client.Get(baseUrl + programUri)
	if err != nil {
		fmt.Println(fallbackTitle)
		return
	}
	defer res.Body.Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			fmt.Println(fallbackTitle)
			slog.Error("unable to read gzipped response body", "error", err)
			return
		}
		defer reader.Close()
	default:
		reader = res.Body
	}

	title := ParseTitle(reader, programTime, programTitle)
	fmt.Println(title)
}

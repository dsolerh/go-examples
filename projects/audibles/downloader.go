package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// ProgressWriter wraps an io.Writer to track download progress
type ProgressWriter struct {
	writer     io.Writer
	total      int64
	downloaded int64
	lastUpdate time.Time
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.writer.Write(p)
	if err != nil {
		return n, err
	}

	pw.downloaded += int64(n)

	// Update progress every 500ms to avoid too frequent updates
	if time.Since(pw.lastUpdate) > 500*time.Millisecond {
		pw.printProgress()
		pw.lastUpdate = time.Now()
	}

	return n, err
}

func (pw *ProgressWriter) printProgress() {
	if pw.total > 0 {
		percentage := float64(pw.downloaded) / float64(pw.total) * 100
		fmt.Printf("\rDownloading... %.2f%% (%s / %s)",
			percentage,
			formatBytes(pw.downloaded),
			formatBytes(pw.total))
	} else {
		fmt.Printf("\rDownloading... %s", formatBytes(pw.downloaded))
	}
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

type downloadOpts struct {
	timeout time.Duration
}

func downloadFile(url, filename string, opts downloadOpts) error {
	fmt.Println("=== Starting Download ===")
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: opts.timeout,
	}

	// Make the request
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Get file size from Content-Length header
	var fileSize int64
	if contentLength := resp.Header.Get("Content-Length"); contentLength != "" {
		if size, err := strconv.ParseInt(contentLength, 10, 64); err == nil {
			fileSize = size
		}
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Create progress writer
	progressWriter := &ProgressWriter{
		writer:     file,
		total:      fileSize,
		lastUpdate: time.Now(),
	}

	fmt.Printf("Starting download of %s...\n", filename)
	start := time.Now()

	// Copy the response body to file with progress tracking
	_, err = io.Copy(progressWriter, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	// Final progress update
	progressWriter.printProgress()
	fmt.Printf("\nDownload completed in %v\n", time.Since(start))

	return nil
}

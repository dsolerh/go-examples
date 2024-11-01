package data

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const RequestTimeount = 30 * time.Second

func GetData(year, day int) (io.Reader, error) {
	fpath := getFilePath(year, day)
	if _, err := os.Stat(fpath); err == nil {
		f, err := os.Open(fpath)
		if err != nil {
			return nil, fmt.Errorf("error while opening the file: %w", err)
		}
		return f, nil
	}
	return RequestData(year, day)
}

func RequestData(year, day int) (io.Reader, error) {
	baseURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	session := os.Getenv("AOC_SESSION")

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeount)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/input", baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating the http request: %w", err)
	}

	req.Header.Add("Accept", "text/html")
	req.Header.Add("Accept", "application/xhtml+xml")
	req.Header.Add("Accept", "application/xml;q=0.9")
	req.Header.Add("Accept", "image/avif")
	req.Header.Add("Accept", "image/webp")
	req.Header.Add("Accept", "image/apng")
	req.Header.Add("Accept", "*/*;q=0.8")
	req.Header.Add("Accept", "application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))
	req.Header.Add("Referer", baseURL)

	client := &http.Client{Timeout: RequestTimeount}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while making the http request: %w", err)
	}
	defer resp.Body.Close()

	return WriteData(year, day, resp.Body)
}

func WriteData(year, day int, r io.Reader) (io.Reader, error) {
	buff := &bytes.Buffer{}
	f, err := os.Create(getFilePath(year, day))
	if err != nil {
		return nil, fmt.Errorf("error while creating the file: %w", err)
	}

	_, err = io.Copy(buff, r)
	if err != nil {
		return nil, fmt.Errorf("error while copying the data to memory: %w", err)
	}

	data := buff.Bytes()

	_, err = io.Copy(f, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error while writing the file: %w", err)
	}

	return bytes.NewBuffer(data), nil
}

func getFilePath(year, day int) string {
	return fmt.Sprintf("./%d_day_%d.txt", year, day)
}

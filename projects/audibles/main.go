package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var book_url = flag.String("bu", "", "book url")
var book_name = flag.String("bn", "", "book name")
var chapter_format = flag.String("cf", "Chapter %03d", "book name")
var timeout = flag.Duration("t", 2*time.Minute, "book name")

func main() {
	flag.Parse()
	chapters, err := getBookChapters(*book_url)
	if err != nil {
		fmt.Printf("could not open the book file err: %v\n", err)
		return
	}
	downloadAudioBook(*book_name, *chapter_format, *timeout, chapters)

}

func getBookChapters(book string) ([]chapterData, error) {
	cacheFile := fmt.Sprintf("cache/%s.json", book)
	// Check if the file exists in cache
	if _, err := os.Stat(cacheFile); err == nil {
		// File exists, read from cache
		fmt.Printf("Reading from cache: %s\n", cacheFile)
		data, err := os.ReadFile(cacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read cache file: %v", err)
		}
		return parseBookData(data)
	}

	data, err := getPageSrc(book)
	if err != nil {
		return nil, err
	}

	var markStart = []byte("tracks = ")
	var markSkip = []byte("},")
	var markEnd = []byte("]")
	var markLen = len(markStart)

	start := bytes.Index(data, markStart)
	if start == -1 {
		return nil, fmt.Errorf("could not find the tracks")
	}

	toSkip := bytes.Index(data[start+markLen:], markSkip)
	if toSkip == -1 {
		return nil, fmt.Errorf("could not find the tracks")
	}

	end := bytes.Index(data[start+markLen:], markEnd)
	if end == -1 {
		return nil, fmt.Errorf("could not find the tracks")
	}

	dataSections := data[start+markLen+toSkip+len(markSkip)-1 : start+markLen+end+1]
	dataSections[0] = '['
	fmt.Printf("dataSections: %s\n", dataSections)

	err = os.WriteFile(cacheFile, dataSections, 0755)
	if err != nil {
		return nil, err
	}

	return parseBookData(dataSections)
}

func getPageSrc(book string) ([]byte, error) {
	cacheFile := fmt.Sprintf("cache/%s", book)

	// Check if the file exists in cache
	if _, err := os.Stat(cacheFile); err == nil {
		// File exists, read from cache
		fmt.Printf("Reading from cache: %s\n", cacheFile)
		file, err := os.ReadFile(cacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read cache file: %v", err)
		}
		return file, nil
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("https://tokybook.com/%s", book))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, resp.Body)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(cacheFile), 0755); err != nil {
		return nil, err
	}

	err = os.WriteFile(cacheFile, buff.Bytes(), 0755)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

type chapterData struct {
	Name      string `json:"name"`
	ChapterId string `json:"chapter_id"`
}

func parseBookData(data []byte) ([]chapterData, error) {
	sections := []chapterData{}
	err := json.Unmarshal(data, &sections)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

type downloadData struct {
	url  string
	name string
}

func downloadAudioBook(name string, chapter_format string, timeout time.Duration, chapters []chapterData) {
	downloadChan := make(chan downloadData, 5)
	lr := newLinkRetriever()
	go func() {
		for idx, chap := range chapters {
			chapNo, err := strconv.Atoi(chap.ChapterId)
			if err != nil {
				fmt.Printf("could not parse the chapter err: %v\n", err)
				continue
			}
			link, err := lr.getMp3Link(context.Background(), chapNo)
			if err != nil {
				fmt.Printf("could not retrieve the chapter url err: %v\n", err)
			}
			chapName := chap.Name
			if chapter_format != "" {
				chapName = fmt.Sprintf(chapter_format, idx+1)
			}
			downloadChan <- downloadData{
				url:  link,
				name: fmt.Sprintf("downloads/%s/%s.mp3", name, chapName),
			}
		}
		close(downloadChan)
	}()

	opts := downloadOpts{
		timeout: timeout,
	}

	failed := []downloadData{}
	for downloadInfo := range downloadChan {
		err := downloadFile(downloadInfo.url, downloadInfo.name, opts)
		if err != nil {
			fmt.Printf("\nDownload failed for %s: %v\n", downloadInfo.url, err)
			failed = append(failed, downloadInfo)
		}
	}

	stillFailed := []downloadData{}
	// retry the rejected
	for _, downloadInfo := range failed {
		fmt.Println("retry failed downloads")
		err := downloadFile(downloadInfo.url, downloadInfo.name, opts)
		if err != nil {
			fmt.Printf("\nDownload failed for %s: %v\n", downloadInfo.url, err)
			stillFailed = append(stillFailed, downloadInfo)
		}
	}

	fmt.Println("failed chapters:")
	for _, dd := range stillFailed {
		fmt.Printf("url: %s\nname: %s\n", dd.url, dd.name)
	}
}

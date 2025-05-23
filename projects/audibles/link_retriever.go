package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type linkRetriever struct {
	client *http.Client
}

func newLinkRetriever() *linkRetriever {
	return &linkRetriever{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// https://api-web.tokybook.com/api-us/getMp3Link
// {"chapterId": 799072141, "serverType": 1}
// Access-Control-Allow-Origin: *
func (l *linkRetriever) getMp3Link(ctx context.Context, chap int) (string, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(map[string]int{
		"chapterId":  chap,
		"serverType": 1,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api-web.tokybook.com/api-us/getMp3Link",
		body,
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := l.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// buff := new(bytes.Buffer)
	// _, err = io.Copy(buff, resp.Body)
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Printf("buff.String(): %v\n", buff.String())

	type respData struct {
		LinkMp3 string `json:"link_mp3"`
	}
	var rdata = respData{}
	err = json.NewDecoder(resp.Body).Decode(&rdata)
	if err != nil {
		return "", err
	}

	return rdata.LinkMp3, nil
}

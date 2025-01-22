package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	err := request()
	fmt.Printf("err: %v\n", err)
}

func request() error {
	reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, "GET", "https://echo.free.beeceptor.com", nil)
	if err != nil {
		return err
	}

	// add the headers
	// req.Header.Set("Accept", "text/html")
	// req.Header.Add("Accept", "application/xhtml+xml")
	setHeaders(
		req,
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language: en-US,en;q=0.9",
		"Cache-Control: no-cache",
		"Connection: keep-alive",
		"Content-Type: application/x-www-form-urlencoded",
		"Cookie: PHPSESSID=a4cfdf069f228ced38ebf977bc39952e; wmsession=d0d1c396-284a-439d-915a-b2dd47c864ce-1732145711047; wm_lang_code=en",
		"DNT: 1",
		"Origin: https://www.cuballama.com",
		"Pragma: no-cache",
		"Referer: https://www.cuballama.com/viajes/activity/search",
		"Sec-Fetch-Dest: document",
		"Sec-Fetch-Mode: navigate",
		"Sec-Fetch-Site: same-origin",
		"Upgrade-Insecure-Requests: 1",
		"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		`sec-ch-ua: "Not;A=Brand";v="24", "Chromium";v="128"`,
		"sec-ch-ua-mobile: ?0",
		`sec-ch-ua-platform: "macOS"`,
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("body: %s\n", b)

	return nil
}

func setHeaders(rq *http.Request, headers ...string) {
	for _, headerPair := range headers {
		keyvalue := strings.Split(headerPair, ": ")
		rq.Header.Set(keyvalue[0], keyvalue[1])
	}
}

type RequestBuilder struct {
	headers http.Header
	client  *http.Client
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		headers: make(http.Header),
		client:  http.DefaultClient,
	}
}

func (r *RequestBuilder) Header(key, value string) *RequestBuilder {
	r.headers.Set(key, key)
	return r
}

func (r *RequestBuilder) HeaderStr(headerPair string) *RequestBuilder {
	keyvalue := strings.Split(headerPair, ": ")
	r.headers.Set(keyvalue[0], keyvalue[1])
	return r
}

func (r *RequestBuilder) Get() {

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jpillora/backoff"
)

type GetOptions struct {
	WaitForChange bool
	LastETag      string
	Recursive     bool
}

type GetResult struct {
	ETag string
}

type GCEMetadataClient struct {
	http *http.Client
}

func NewGCEMetadataClient() *GCEMetadataClient {
	return &GCEMetadataClient{
		http: &http.Client{},
	}
}

func (c *GCEMetadataClient) Get(
	key string,
	data interface{},
	opts GetOptions) (*GetResult, error) {
	boff := &backoff.Backoff{
		Jitter: true,
	}

	for {
		url := fmt.Sprintf("http://metadata.google.internal/computeMetadata/v1/%s", key)

		resp, err := c.doGET(url, opts)
		if err != nil {
			log.Print(err)
			time.Sleep(boff.Duration())
			continue
		}
		switch resp.StatusCode {
		case http.StatusOK:
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading body from response: %s", err)
				time.Sleep(boff.Duration())
				continue
			}

			err = json.Unmarshal(b, data)
			if err != nil {
				log.Printf("Error parsing body from response: %s. Body was: %s", err, string(b))
				time.Sleep(boff.Duration())
				continue
			}
			return &GetResult{
				ETag: resp.Header.Get("ETag"),
			}, nil

		case http.StatusBadGateway, http.StatusServiceUnavailable,
			http.StatusGatewayTimeout:
			log.Printf("Status %d from %s: %s", resp.StatusCode, url, resp.Status)
			time.Sleep(boff.Duration())
			continue

		default:
			return nil, fmt.Errorf("GET %s failed with status %d: %s",
				url, resp.StatusCode, resp.Status)
		}
	}
}

func (c *GCEMetadataClient) doGET(endpointURL string, opts GetOptions) (*http.Response, error) {
	u, err := url.Parse(endpointURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("alt", "json")
	if opts.WaitForChange {
		q.Set("wait_for_change", "true")
	}
	if opts.LastETag != "" {
		q.Set("last_etag", opts.LastETag)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Metadata-Flavor", "Google")

	log.Printf("GET %s", u)
	return c.http.Do(req)
}

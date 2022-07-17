package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	DefaultHost              = "http://192.168.99.2"
	DefaultPort              = 8700
	DefaultHttpClientTimeout = 10 // second
)

// Represents api client
type Client struct {
	host       string
	httpClient *http.Client
}

func NewClient() *Client {
	cl := &http.Client{
		Timeout: time.Second * DefaultHttpClientTimeout,
	}

	return &Client{
		host:       concatHostPort(DefaultHost, fmt.Sprint(DefaultPort)),
		httpClient: cl,
	}
}

func (c *Client) RunPing(ctx context.Context) error {
	url := fmt.Sprint(c.host, "/ping")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("ping status not 200. Got %d", resp.StatusCode)
	}

	return nil
}

func concatHostPort(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

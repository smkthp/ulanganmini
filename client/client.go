package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/smkthp/ulanganmini/system"
	"github.com/smkthp/ulanganmini/util"
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

func (c *Client) RunGetTasks(ctx context.Context) ([]system.Task, error) {
	url := fmt.Sprint(c.host, "/tasks")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tasks, err := util.UnmarshalTasks(body)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func concatHostPort(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

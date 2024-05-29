package provider

import (
	"love_knot/internal/pkg/client"
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func NewRequestClient(c *http.Client) *client.RequestClient {
	return client.NewRequestClient(c)
}

package gorest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
	apiUrl string
}

func NewClient(apiUrl string) (*Client, error) {
	if apiUrl == "" {
		return nil, errors.New("API URL is required")
	}

	return &Client{
		client: &http.Client{
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		apiUrl: apiUrl,
	}, nil
}

func (c Client) GetUsers() ([]User, error) {
	url := fmt.Sprintf("%s/users", c.apiUrl)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r usersResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Users, nil
}

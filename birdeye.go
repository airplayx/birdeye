package birdeye

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://public-api.birdeye.so/defi"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	var result = Client{
		apiKey: apiKey,
	}
	return &result
}

// birdEye defines the interface for all response birdEye
type birdEye interface {
	TradeData |
		TokenOverview |
		TokenListV1
}

// cli represents a Birdeye API cli with generic type support
type cli[m birdEye] struct {
	apiKey     string
	httpClient *http.Client
}

// newCli creates a new Birdeye API cli
func newCli[T birdEye](apiKey string) *cli[T] {
	var result = cli[T]{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
	return &result
}

// makeRequest is a generic helper function that handles HTTP requests and response parsing
func (c *cli[T]) makeRequest(method, url string, body io.Reader) (*T, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &result, nil
}

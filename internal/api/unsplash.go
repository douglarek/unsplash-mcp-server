package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/douglarek/unsplash-mcp-server/internal/config"
	"github.com/douglarek/unsplash-mcp-server/internal/models"
)

const (
	// UnsplashAPIURL is the base URL for Unsplash search API
	UnsplashAPIURL = "https://api.unsplash.com/search/photos"

	// UnsplashAPIVersion is the API version used
	UnsplashAPIVersion = "v1"
)

// Client provides methods to interact with the Unsplash API
type Client struct {
	httpClient *http.Client
	accessKey  string
}

// NewClient creates a new Unsplash API client
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
		accessKey: cfg.UnsplashAccessKey,
	}
}

// SearchPhotos searches for photos with the given parameters
func (c *Client) SearchPhotos(ctx context.Context, params url.Values) ([]models.Photo, error) {
	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, UnsplashAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Add("Accept-Version", UnsplashAPIVersion)
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", c.accessKey))
	req.URL.RawQuery = params.Encode()

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var searchResp models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return searchResp.Results, nil
}

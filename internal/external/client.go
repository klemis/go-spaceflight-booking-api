package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/klemis/go-spaceflight-booking-api/models"
)

// SpaceXAPIClient represents a client for interacting with the SpaceX API.
type SpaceXAPIClient struct {
	Client  *http.Client
	BaseURL string
}

// NewSpaceXAPIClient creates a new instance of SpaceXAPIClient with the specified base URL.
func NewSpaceXAPIClient(baseURL string) *SpaceXAPIClient {
	return &SpaceXAPIClient{
		Client:  http.DefaultClient,
		BaseURL: baseURL,
	}
}

// CheckScheduledLaunches checks if there are any launches scheduled for the given request body.
func (c *SpaceXAPIClient) CheckScheduledLaunches(body models.RequestBody) (models.FilteredResponse, error) {
	url := c.BaseURL + "launches/query"

	var result models.FilteredResponse
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return result, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode response body: %w", err)
	}

	return result, nil
}

// CheckLaunchpadState checks launchpad state.
func (c *SpaceXAPIClient) CheckLaunchpadState(id string) (string, error) {
	// FIXME: This endpoint does not support querying by status alone.
	// Additionally, the launchpads/query endpoint does not allow filtering by launchpad ID.
	// So this endpoint requests a single launchpad and returns its Status.
	url := c.BaseURL + fmt.Sprintf("launchpads/%s", id)

	resp, err := c.Client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result models.Launchpad
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	return result.Status, nil
}

// GetActiveLaunchpads gets launchpads in active state.
func (c *SpaceXAPIClient) GetActiveLaunchpads(body models.RequestBody) ([]models.Filtered, error) {
	url := c.BaseURL + "launchpads/query"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result models.FilteredResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return result.Docs, nil
}

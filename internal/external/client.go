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

// CheckLaunchpadAvailability checks if there are any launches available for the given request body.
func (c *SpaceXAPIClient) CheckLaunchpadAvailability(body models.RequestBody) (bool, error) {
	url := c.BaseURL + "launches/query"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return false, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result models.Launches
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode response body: %w", err)
	}

	return len(result.Docs) == 0, nil
}

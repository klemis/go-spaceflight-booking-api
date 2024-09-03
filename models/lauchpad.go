package models

// Launchpad represents the details of a launchpad.
type Launchpad struct {
	Images struct {
		Large []string `json:"large"`
	} `json:"images"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Locality        string   `json:"locality"`
	Region          string   `json:"region"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	LaunchAttempts  int      `json:"launch_attempts"`
	LaunchSuccesses int      `json:"launch_successes"`
	Rockets         []string `json:"rockets"`
	Timezone        string   `json:"timezone"`
	Launches        []string `json:"launches"`
	Status          string   `json:"status"`
	Details         string   `json:"details"`
	ID              string   `json:"id"`
}

// Filtered represents the id of the resource.
type Filtered struct {
	ID string `json:"id"`
}

// FilteredResponse represents the response containing a list of filtered ids.
type FilteredResponse struct {
	Docs []Filtered `json:"docs"`
}

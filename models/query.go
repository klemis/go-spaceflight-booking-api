package models

// PopulateOption represents the populate option for embedded documents.
type PopulateOption struct {
	Path   string         `json:"path"`
	Select map[string]int `json:"select"`
}

// Options specifies additional query options.
type Options struct {
	Select   map[string]int   `json:"select"`
	Populate []PopulateOption `json:"populate"`
}

type RequestBody struct {
	Query   map[string]interface{} `json:"query"`
	Options Options                `json:"options"`
}

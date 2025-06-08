package apimodel

type Board struct {
	ID             string                   `json:"id"`
	TeamID         string                   `json:"teamId"`
	Title          string                   `json:"title"`
	CardProperties []map[string]interface{} `json:"cardProperties"`
}

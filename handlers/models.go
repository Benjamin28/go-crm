package handlers

// Client .
type Client struct {
	ID           string   `json:"id"`
	Name         string   `json:"firstname"`
	Contact      string   `json:"lastname"`
	Stage        string   `json:"status"`
	Description  string   `json:"description"`
}
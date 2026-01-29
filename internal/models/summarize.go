package models

// SummarizeRequest represents input JSON
type SummarizeRequest struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
}



// SummarizeResponse represents output JSON
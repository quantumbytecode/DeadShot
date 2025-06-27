package models

import "time"

type LogModel struct {
	ID              int       `json:"id" db:"id"`
	Method          string    `json:"method" db:"method"`
	URL             string    `json:"url" db:"url"`
	Headers         string    `json:"headers" db:"headers"`
	QueryParams     string    `json:"query_params" db:"query_params"`
	Body            string    `json:"body" db:"body"`
	ReceivedAt      time.Time `json:"received_at" db:"received_at"`
	StatusCode      int       `json:"status_code" db:"status_code"`
	ResponseHeaders string    `json:"response_headers" db:"response_headers"`
	ResponseBody    string    `json:"response_body" db:"response_body"`
	Tags            string    `json:"tags" db:"tags"`         // Comma-separated or JSON
	Source          string    `json:"source" db:"source"`     // App/service name
	Replayed        bool      `json:"replayed" db:"replayed"` // Was this already replayed?
	Error           string    `json:"error" db:"error"`       // Error info, if any
}

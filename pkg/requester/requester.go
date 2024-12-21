package requester

import "context"

// Response represents a generic HTTP response
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}

// Requester defines the interface for making HTTP requests
type RequesterI interface {
	Get(ctx context.Context, url string, headers map[string]string) (*Response, error)
	Post(ctx context.Context, url string, headers map[string]string, body interface{}) (*Response, error)
}

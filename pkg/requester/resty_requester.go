package requester

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// RestyRequester implements Requester interface using resty client
type RestyRequester struct {
	client *resty.Client
}

// NewRestyRequester creates a new instance of RestyRequester
func NewRestyRequester() RequesterI {
	return &RestyRequester{
		client: resty.New(),
	}
}

// Get performs a GET request
func (r *RestyRequester) Get(ctx context.Context, url string, headers map[string]string) (*Response, error) {
	req := r.client.R().
		SetContext(ctx).
		SetHeaders(headers)

	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Headers:    resp.Header(),
	}, nil
}

// Post performs a POST request
func (r *RestyRequester) Post(ctx context.Context, url string, headers map[string]string, body interface{}) (*Response, error) {
	req := r.client.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(body)

	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Headers:    resp.Header(),
	}, nil
}

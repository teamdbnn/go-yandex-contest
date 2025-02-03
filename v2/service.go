package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSubmissionsQueueCapacity Get submissions queue capacity
type GetSubmissionsQueueCapacity struct {
	c *Client
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetSubmissionsQueueCapacity) Do(ctx context.Context, opts ...RequestOption) (*ServiceCapacity, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/service/capacity"),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ServiceCapacity)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAvailableScopes Get available scopes
type GetAvailableScopes struct {
	c *Client
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetAvailableScopes) Do(ctx context.Context, opts ...RequestOption) (*TokenInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/service/introspect"),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(TokenInfo)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

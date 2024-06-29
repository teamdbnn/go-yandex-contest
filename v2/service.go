package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAvailableScopes Get available scopes
type GetAvailableScopes struct {
	c *Client
}

// Do send req
func (s *GetAvailableScopes) Do(ctx context.Context, opts ...RequestOption) (*Service, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/service/introspect"),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Service)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

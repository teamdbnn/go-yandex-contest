package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUserTeams Get user teams
type GetUserTeams struct {
	c *Client
}

// Do Send GET request
func (s *GetUserTeams) Do(ctx context.Context, opts ...RequestOption) ([]*TeamView, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/teams"),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*TeamView, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

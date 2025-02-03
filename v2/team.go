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

// Do Get user teams
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/team/listTeamsUsingGET
// meta:operation GET /teams
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
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

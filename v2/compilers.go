package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCompilersList Get Compilers list
type GetCompilersList struct {
	c *Client
}

// Do Get compilers list
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/compilers/compilersListUsingGET
// meta:operation GET /compilers
func (s *GetCompilersList) Do(ctx context.Context, opts ...RequestOption) (*CompilerListResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/compilers"),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(CompilerListResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

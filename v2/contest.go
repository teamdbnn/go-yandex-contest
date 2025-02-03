package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestInfo Get information about contest by id
type GetContestInfo struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetContestInfo) ContestID(contestID int64) *GetContestInfo {
	s.contestID = contestID
	return s
}

func (s *GetContestInfo) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get information about contest
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/contest/getContestInfoUsingGET
// meta:operation GET /contests/{contestId}
func (s *GetContestInfo) Do(ctx context.Context, opts ...RequestOption) (*ContestDescription, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d", s.contestID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ContestDescription)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

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
	contestId int64
}

// ContestID Set contest id
func (s *GetContestInfo) ContestID(contestId int64) *GetContestInfo {
	s.contestId = contestId
	return s
}

func (s *GetContestInfo) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do Send request to GetContestService
func (s *GetContestInfo) Do(ctx context.Context, opts ...RequestOption) (*ContestDescription, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d", s.contestId),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(ContestDescription)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

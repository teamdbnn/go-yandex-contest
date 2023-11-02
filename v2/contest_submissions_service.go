package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSubmissionsOfContestService Get contest participants
type GetSubmissionsOfContestService struct {
	c        *Client
	contest  int64
	page     int32
	pageSize int32
}

// Contest Set contest
func (s *GetSubmissionsOfContestService) Contest(contest int64) *GetSubmissionsOfContestService {
	s.contest = contest
	return s
}

// Page Set page
func (s *GetSubmissionsOfContestService) Page(page int32) *GetSubmissionsOfContestService {
	s.page = page
	return s
}

// PageSize Set page size, default 100
func (s *GetSubmissionsOfContestService) PageSize(pageSize int32) *GetSubmissionsOfContestService {
	s.pageSize = pageSize
	return s
}

func (s *GetSubmissionsOfContestService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetSubmissionsOfContestService) Do(ctx context.Context, opts ...RequestOption) ([]*Submission, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions", s.contest),
	}
	if s.page != 0 {
		r.setParam("page", s.page)
	}
	if s.pageSize != 0 {
		r.setParam("pageSize", s.pageSize)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*Submission, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

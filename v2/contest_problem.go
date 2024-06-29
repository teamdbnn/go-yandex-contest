package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestProblems Get contest problems
type GetContestProblems struct {
	c       *Client
	contest int64
	locale  string
}

// Contest Set contest
func (s *GetContestProblems) Contest(contest int64) *GetContestProblems {
	s.contest = contest
	return s
}

// Locale Set locale
func (s *GetContestProblems) Locale(locale string) *GetContestProblems {
	s.locale = locale
	return s
}

func (s *GetContestProblems) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetContestProblems) Do(ctx context.Context, opts ...RequestOption) (*ContestProblems, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/problems", s.contest),
	}
	if s.locale != "" {
		r.setParam("locale", s.locale)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(ContestProblems)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestService Get Contest item info
type GetContestService struct {
	c       *Client
	contest int64
}

// Contest Set contest id
func (s *GetContestService) Contest(contest int64) *GetContestService {
	s.contest = contest
	return s
}

func (s *GetContestService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do Send request
func (s *GetContestService) Do(ctx context.Context, opts ...RequestOption) (*ContestDescription, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d", s.contest),
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

type GetClarificationsInContestService struct {
	c       *Client
	contest int64
	locale  string
}

func (s *GetClarificationsInContestService) Contest(contest int64) *GetClarificationsInContestService {
	s.contest = contest
	return s
}

func (s *GetClarificationsInContestService) Locale(locale string) *GetClarificationsInContestService {
	s.locale = locale
	return s
}

func (s *GetClarificationsInContestService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

func (s *GetClarificationsInContestService) Do(ctx context.Context, opts ...RequestOption) (*Clarifications, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d/clarifications", s.contest),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Clarifications)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetContestMessagesServie struct {
	c       *Client
	contest int64
}

func (s *GetContestMessagesServie) Contest(contest int64) *GetContestMessagesServie {
	s.contest = contest
	return s
}

func (s *GetContestMessagesServie) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

func (s *GetContestMessagesServie) Do(ctx context.Context, opts ...RequestOption) (*Messages, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d/messages", s.contest),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Messages)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

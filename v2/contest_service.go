package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestService Get information about contest by id
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

// Do Send request to GetContestService
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

// GetClarificationsInContestService Get clarifications in contest by contest id
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
	if s.locale != "" {
		r.setParam("locale", s.locale)
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

// GetContestMessagesService Get messages in contest by contest id
type GetContestMessagesService struct {
	c       *Client
	contest int64
}

func (s *GetContestMessagesService) Contest(contest int64) *GetContestMessagesService {
	s.contest = contest
	return s
}

func (s *GetContestMessagesService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

func (s *GetContestMessagesService) Do(ctx context.Context, opts ...RequestOption) (*Messages, error) {
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

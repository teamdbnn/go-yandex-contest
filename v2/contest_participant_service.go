package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetParticipantsOfContestService Get contest participants
type GetParticipantsOfContestService struct {
	c           *Client
	contest     int64
	displayName string
	login       string
}

// Contest Set contest
func (s *GetParticipantsOfContestService) Contest(contest int64) *GetParticipantsOfContestService {
	s.contest = contest
	return s
}

// DisplayName Set display name
func (s *GetParticipantsOfContestService) DisplayName(displayName string) *GetParticipantsOfContestService {
	s.displayName = displayName
	return s
}

// Login Set login
func (s *GetParticipantsOfContestService) Login(login string) *GetParticipantsOfContestService {
	s.login = login
	return s
}

func (s *GetParticipantsOfContestService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetParticipantsOfContestService) Do(ctx context.Context, opts ...RequestOption) ([]*Participant, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contest),
	}
	if s.displayName != "" {
		r.setParam("display_name", s.displayName)
	}
	if s.login != "" {
		r.setParam("login", s.login)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*Participant, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RegisterParticipantForContestService Register for contest
type RegisterParticipantForContestService struct {
	c       *Client
	contest int64
	login   string
	uid     int64
}

func (s *RegisterParticipantForContestService) Contest(contest int64) *RegisterParticipantForContestService {
	s.contest = contest
	return s
}

func (s *RegisterParticipantForContestService) Login(login string) *RegisterParticipantForContestService {
	s.login = login
	return s
}

func (s *RegisterParticipantForContestService) UID(uid int64) *RegisterParticipantForContestService {
	s.uid = uid
	return s
}

func (s *RegisterParticipantForContestService) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	if s.login == "" {
		return requiredError("login")
	}
	return nil
}

func (s *RegisterParticipantForContestService) Do(ctx context.Context, opts ...RequestOption) (*int64, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contest),
	}
	if s.login != "" {
		r.setParam("login", s.login)
	}

	if s.uid != 0 {
		r.setParam("uid", s.uid) // todo: Check usage
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := string(data)
	i, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

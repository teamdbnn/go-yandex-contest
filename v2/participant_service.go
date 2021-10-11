package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantListOfCompetitionService Get Participant list of competition
type GetParticipantListOfCompetitionService struct {
	c           *Client
	competition int64
}

func (s *GetParticipantListOfCompetitionService) validate() error {
	if s.competition == 0 {
		return newError().Required("competition")
	}
	return nil
}

// Competition Set competition
func (s *GetParticipantListOfCompetitionService) Competition(competition int64) *GetParticipantListOfCompetitionService {
	s.competition = competition
	return s
}

// Do Send GET request
func (s *GetParticipantListOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) (res *GrantResponse, err error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competition),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GrantResponse)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetParticipantListOfContest struct
type GetParticipantListOfContest struct {
	c           *Client
	contest     int64
	displayName string
	login       string
}

// Contest Set contest
func (s *GetParticipantListOfContest) Contest(contest int64) *GetParticipantListOfContest {
	s.contest = contest
	return s
}

// DisplayName Set display name
func (s *GetParticipantListOfContest) DisplayName(displayName string) *GetParticipantListOfContest {
	s.displayName = displayName
	return s
}

// Login Set login
func (s *GetParticipantListOfContest) Login(login string) *GetParticipantListOfContest {
	s.login = login
	return s
}

func (s *GetParticipantListOfContest) validate() error {
	if s.contest == 0 {
		return newError().Required("contest")
	}
	return nil
}

// Do send req
func (s *GetParticipantListOfContest) Do(ctx context.Context, opts ...RequestOption) (res []*Participant, err error) {
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

	res = make([]*Participant, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

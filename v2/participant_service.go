package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantListOfCompetition Get Participant list of competition
type GetParticipantListOfCompetition struct {
	c           *Client
	competition int64
}

// Competition Set competition
func (s *GetParticipantListOfCompetition) Competition(competition int64) *GetParticipantListOfCompetition {
	s.competition = competition
	return s
}

// Do Send GET request
func (s *GetParticipantListOfCompetition) Do(ctx context.Context, opts ...RequestOption) (res *GrantResponse, err error) {
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

// Do send req
func (s *GetParticipantListOfContest) Do(ctx context.Context, opts ...RequestOption) (res []*Participant, err error) {
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

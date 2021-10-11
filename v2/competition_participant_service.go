package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantsOfCompetitionService Get Participant list of competition
type GetParticipantsOfCompetitionService struct {
	c           *Client
	competition int64
}

func (s *GetParticipantsOfCompetitionService) validate() error {
	if s.competition == 0 {
		return requiredError("competition")
	}
	return nil
}

// Competition Set competition
func (s *GetParticipantsOfCompetitionService) Competition(competition int64) *GetParticipantsOfCompetitionService {
	s.competition = competition
	return s
}

// Do Send GET request
func (s *GetParticipantsOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*GrantResponse, error) {
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
	res := new(GrantResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type RegisterParticipantIntoCompetitionService struct {
	c           *Client
	competition int64
	users       []string
}

func (s *RegisterParticipantIntoCompetitionService) Competition(competition int64) *RegisterParticipantIntoCompetitionService {
	s.competition = competition
	return s
}

func (s *RegisterParticipantIntoCompetitionService) Users(users []string) *RegisterParticipantIntoCompetitionService {
	s.users = users
	return s
}

func (s *RegisterParticipantIntoCompetitionService) validate() error {
	if s.competition == 0 {
		return requiredError("competition")
	}
	if s.users == nil {
		return requiredError("users")
	}

	return nil
}

// Do Send GET request
func (s *RegisterParticipantIntoCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*GrantResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competition),
	}
	r.setFormParam("userUids", s.users)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(GrantResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

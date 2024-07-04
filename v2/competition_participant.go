package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantsOfCompetition Get Participant list of competition
type GetParticipantsOfCompetition struct {
	c             *Client
	competitionID int64
}

// CompetitionID Set competition id
func (s *GetParticipantsOfCompetition) CompetitionID(competitionID int64) *GetParticipantsOfCompetition {
	s.competitionID = competitionID
	return s
}

func (s *GetParticipantsOfCompetition) validate() error {
	if s.competitionID == 0 {
		return requiredError("competitionID")
	}
	return nil
}

// Do Send GET request
func (s *GetParticipantsOfCompetition) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	if err := s.validate(); err != nil {
		return "", err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competitionID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// RegisterParticipantIntoCompetition Register Participant into competition
type RegisterParticipantIntoCompetition struct {
	c             *Client
	competitionID int64
	body          struct {
		Users []string `json:"users"`
	}
}

// CompetitionID Set competition id
func (s *RegisterParticipantIntoCompetition) CompetitionID(competitionID int64) *RegisterParticipantIntoCompetition {
	s.competitionID = competitionID
	return s
}

// Users Set users
func (s *RegisterParticipantIntoCompetition) Users(users []string) *RegisterParticipantIntoCompetition {
	s.body.Users = users
	return s
}

func (s *RegisterParticipantIntoCompetition) validate() error {
	if s.competitionID == 0 {
		return requiredError("competitionID")
	}
	if s.body.Users == nil {
		return requiredError("Users")
	}

	return nil
}

// Do Send POST request
func (s *RegisterParticipantIntoCompetition) Do(ctx context.Context, opts ...RequestOption) (*CompetitionRegisterResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competitionID),
	}
	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(CompetitionRegisterResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

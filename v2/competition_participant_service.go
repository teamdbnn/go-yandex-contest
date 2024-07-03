package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantsOfCompetitionService Get Participant list of competition
type GetParticipantsOfCompetitionService struct {
	c             *Client
	competitionID int64
}

// CompetitionID Set competition id
func (s *GetParticipantsOfCompetitionService) CompetitionID(competitionID int64) *GetParticipantsOfCompetitionService {
	s.competitionID = competitionID
	return s
}

func (s *GetParticipantsOfCompetitionService) validate() error {
	if s.competitionID == 0 {
		return requiredError("competitionID")
	}
	return nil
}

// Do Send GET request
func (s *GetParticipantsOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competitionID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// RegisterParticipantIntoCompetitionService Register Participant into competition
type RegisterParticipantIntoCompetitionService struct {
	c             *Client
	competitionID int64
	body          struct {
		Users []string `json:"users"`
	}
}

// CompetitionID Set competition id
func (s *RegisterParticipantIntoCompetitionService) CompetitionID(competitionID int64) *RegisterParticipantIntoCompetitionService {
	s.competitionID = competitionID
	return s
}

// Users Set users
func (s *RegisterParticipantIntoCompetitionService) Users(Users []string) *RegisterParticipantIntoCompetitionService {
	s.body.Users = Users
	return s
}

func (s *RegisterParticipantIntoCompetitionService) validate() error {
	if s.competitionID == 0 {
		return requiredError("competitionID")
	}
	if s.body.Users == nil {
		return requiredError("Users")
	}

	return nil
}

// Do Send POST request
func (s *RegisterParticipantIntoCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*CompetitionRegisterResponse, error) {
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

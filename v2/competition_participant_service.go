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
	competitionId int64
}

// CompetitionID Set competition id
func (s *GetParticipantsOfCompetitionService) CompetitionID(competitionId int64) *GetParticipantsOfCompetitionService {
	s.competitionId = competitionId
	return s
}

func (s *GetParticipantsOfCompetitionService) validate() error {
	if s.competitionId == 0 {
		return requiredError("competitionId")
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
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competitionId),
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
	competitionId int64
	body          struct {
		users []string
	}
}

// CompetitionID Set competition id
func (s *RegisterParticipantIntoCompetitionService) CompetitionID(competitionId int64) *RegisterParticipantIntoCompetitionService {
	s.competitionId = competitionId
	return s
}

// Users Set users
func (s *RegisterParticipantIntoCompetitionService) Users(users []string) *RegisterParticipantIntoCompetitionService {
	s.body.users = users
	return s
}

func (s *RegisterParticipantIntoCompetitionService) validate() error {
	if s.competitionId == 0 {
		return requiredError("competitionId")
	}
	if s.body.users == nil {
		return requiredError("users")
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
		endpoint: fmt.Sprintf("/competitions/%v/participants", s.competitionId),
	}
	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)
	fmt.Println(data)
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

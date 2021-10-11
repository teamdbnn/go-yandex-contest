package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantsOfContestService struct
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

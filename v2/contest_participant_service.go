package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetParticipantsOfContest struct
type GetParticipantsOfContest struct {
	c           *Client
	contest     int64
	displayName string
	login       string
}

// Contest Set contest
func (s *GetParticipantsOfContest) Contest(contest int64) *GetParticipantsOfContest {
	s.contest = contest
	return s
}

// DisplayName Set display name
func (s *GetParticipantsOfContest) DisplayName(displayName string) *GetParticipantsOfContest {
	s.displayName = displayName
	return s
}

// Login Set login
func (s *GetParticipantsOfContest) Login(login string) *GetParticipantsOfContest {
	s.login = login
	return s
}

func (s *GetParticipantsOfContest) validate() error {
	if s.contest == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetParticipantsOfContest) Do(ctx context.Context, opts ...RequestOption) (res []*Participant, err error) {
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
